package main

import (
	"context"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "phase"
)

type PhaseEqualibriumServer struct {
	pb.UnimplementedPhaseEqualibriumServer
}

type InitType struct {
	FLUID []*pb.OneComponent
	BIPs  []*pb.OneBIP
}

const R = 0.00831675

var N int

var z, Tkr, Pkr, Vkr, w, cpen []float64

//var c [][]float64

var InitData InitType

//var a []*pb.OneBIP
//var b []*pb.OneComponent

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPhaseEqualibriumServer(s, &PhaseEqualibriumServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

// func (s *PhaseEqualibriumServer) Init(ctx context.Context, req *pb.InitMessageRequest) (*pb.InitMessageResponse, error) {
func (s *PhaseEqualibriumServer) Init(ctx context.Context, req *pb.InitMessageRequest) (*emptypb.Empty, error) {
	N = len(req.FLUID)

	InitData.FLUID = req.FLUID
	InitData.BIPs = req.BIPs

	z = make([]float64, N)
	Tkr = make([]float64, N)
	Pkr = make([]float64, N)
	Vkr = make([]float64, N)
	w = make([]float64, N)
	cpen = make([]float64, N)
	var sum_z float64
	sum_z = 0.0
	for i := range InitData.FLUID {
		comp := InitData.FLUID[i]
		z[i] = comp.MoleFraction * 0.01
		Tkr[i] = comp.Tcr + 273.15
		Pkr[i] = comp.Pcr
		Vkr[i] = comp.Vcr
		w[i] = comp.WFact
		cpen[i] = comp.Pen * 0.001
		sum_z += z[i]
	}
	if sum_z != 1. {
		//fmt.Printf("Сумма не 100: sum_z: %3.5f\n", sum_z)
		for i := 0; i < len(z); i++ {
			z[i] = z[i] / sum_z
			//fmt.Printf("Z: i: %3.5f %d\n", z[i], i)
		}
	}

	//c := make([][]float64, N)
	//return &pb.InitMessageResponse{FLUID: req.FLUID, BIPs: req.BIPs}, nil
	return &emptypb.Empty{}, nil
	//return &pb.InitMessageResponse{}, nil
}

func (s *PhaseEqualibriumServer) Vle(ctx context.Context, req *pb.VleMessageRequest) (*pb.VleMessageResponse, error) {

	P := req.Pres
	T := req.Temp

	var W, Z_l, Z_v float64
	var Stable int

	//var Control_phase float64

	Vpkr_ar := multiply(z, Vkr)
	Tpkr_ar := multiply(z, Tkr)
	Vpkr := sum(Vpkr_ar)
	Tpkr := sum(Tpkr_ar)
	Control_phase_MF := Vpkr * Tpkr * Tpkr //the same method of liquid/vapor identification as in Multiflash

	c := make([][]float64, N) // массив
	for i := range c {
		c[i] = make([]float64, N)
	}
	for i := range InitData.BIPs {
		cc_temp := InitData.BIPs[i]
		for j := 0; j < N; j++ {
			c[i][j] = cc_temp.GetNum()[i]
			//c[i][j] = 0.0
		}
	}
	//ac_i := make([]float64, N)
	ac_i := make([]float64, N)
	psi_i := make([]float64, N)
	alpha_i := make([]float64, N)
	a_i := make([]float64, N)
	b_i := make([]float64, N)
	c_i := make([]float64, N)

	Biw := make([]float64, N)
	Ciw := make([]float64, N)

	Bil := make([]float64, N)
	Cil := make([]float64, N)

	K_i := make([]float64, N)
	K_il := make([]float64, N)
	K_iv := make([]float64, N)

	x_i := make([]float64, N)
	y_i := make([]float64, N)

	Yi_l := make([]float64, N)

	alll := make([]float64, N)
	fl_i := make([]float64, N)

	Ri := make([]float64, N)

	avvv := make([]float64, N)
	fw_i := make([]float64, N)

	df_lv := make([]float64, N)

	//fmt.Printf("ac_i: %3.5f\n", ac_i[i])

	for i := 0; i < N; i++ {
		ac_i[i] = 0.42747 * math.Pow(R, 2) * math.Pow(Tkr[i], 2) / Pkr[i]
		psi_i[i] = 0.48 + 1.574*w[i] - 0.176*math.Pow(w[i], 2)
		alpha_i[i] = math.Pow(1+psi_i[i]*(1-math.Sqrt(T/Tkr[i])), 2)
		a_i[i] = ac_i[i] * alpha_i[i]
		b_i[i] = 0.08664 * R * Tkr[i] / Pkr[i]
		c_i[i] = cpen[i]
	}

	for i := 0; i < N; i++ {
		K_i[i] = math.Pow(math.Exp(5.373*(1+w[i])*(1-Tkr[i]/T))*Pkr[i]/P, 1.0) //da
	}

	aw := 0.0
	bw := 0.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			aw += z[i] * z[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
		}
	}
	for i := 0; i < N; i++ {
		bw += z[i] * b_i[i]
	}
	Aw := aw * P / (math.Pow(R, 2) * math.Pow(T, 2))
	Bw := bw * P / (R * T)
	cw := 0.0
	for i := 0; i < N; i++ {
		cw += c_i[i] * z[i]
	}
	Cw := cw * P / (R * T)

	for i := 0; i < N; i++ {
		Biw[i] = b_i[i] * P / (R * T)
		Ciw[i] = c_i[i] * P / (R * T)
	}
	/*fmt.Printf("Aw: %.4f \n", Aw)
	fmt.Printf("Bw: %.4f \n", Bw)
	fmt.Printf("Cw: %.4f \n", Cw)*/

	coefficients := []float64{1, 3*Cw - 1, 3*math.Pow(Cw, 2) - math.Pow(Bw, 2) - 2*Cw - Bw + Aw, math.Pow(Cw, 3) - math.Pow(Bw, 2)*Cw - math.Pow(Cw, 2) - Bw*Cw + Aw*Cw - Aw*Bw}
	var cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])
	//fmt.Printf("coefficients: %.4f %.4f %.4f %.4f\n", coefficients[0], coefficients[1], coefficients[2], coefficients[3])
	Z_v = findMax(cubroot)
	Z_init := Z_v
	//fmt.Printf("Z_v: %.4f \n", Z_v)
	/*fmt.Printf("1 Roots: %.4f, %.4f, %.4f\n", cubroot[0], cubroot[1], cubroot[2])
	  fmt.Printf("1 Z_v %.4f\n", Z_v)*/

	//avvv := make([]float64, N)
	for i := 0; i < N; i++ {
		avv := 0.0
		for j := 0; j < N; j++ {
			avv += z[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
		}
		avvv[i] = avv
	}

	fz_i := make([]float64, N)
	for i := 0; i < N; i++ {
		fz_i[i] = math.Exp(math.Log(z[i]*P) - math.Log(Z_v+Cw-Bw) + (Biw[i]-Ciw[i])/(Z_v+Cw-Bw) - (Aw/Bw)*((2*avvv[i]/aw)-(b_i[i]/bw))*math.Log((Z_v+Bw+Cw)/(Z_v+Cw)) - (Aw/Bw)*(Biw[i]+Ciw[i])/(Z_v+Bw+Cw) + (Aw/Bw)*Ciw[i]/(Z_v+Cw))
	}

	m := 0

	Ri_v := 1.0
	TS_v_flag := 0
	TS_l_flag := 0
	var Sv, Sl float64

	// Часть 1 Проверка газовой фазы

	for m < 30 {

		Yi_v := make([]float64, N)
		Sv1 := 0.0
		for i := 0; i < N; i++ {
			Yi_v[i] = z[i] * K_i[i]
			Sv1 += Yi_v[i]
		}

		Sv = Sv1

		for i := 0; i < N; i++ {
			y_i[i] = Yi_v[i] / Sv
		}

		aw = 0.0
		bw = 0.0
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				aw += y_i[i] * y_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
			}
		}

		for i := 0; i < N; i++ {
			bw += y_i[i] * b_i[i]
		}

		Aw = aw * P / (math.Pow(R, 2) * math.Pow(T, 2))
		Bw = bw * P / (R * T)
		cw = 0.0
		for i := 0; i < N; i++ {
			cw += c_i[i] * y_i[i]
		}
		Cw = cw * P / (R * T)

		for i := 0; i < N; i++ {
			Biw[i] = b_i[i] * P / (R * T)
			Ciw[i] = c_i[i] * P / (R * T)
		}

		coefficients := []float64{1, 3*Cw - 1, 3*math.Pow(Cw, 2) - math.Pow(Bw, 2) - 2*Cw - Bw + Aw, math.Pow(Cw, 3) - math.Pow(Bw, 2)*Cw - math.Pow(Cw, 2) - Bw*Cw + Aw*Cw - Aw*Bw}
		var cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])
		Z_v = findMax(cubroot)

		// Расчет летучестей в паровой фазе

		avvv = make([]float64, N)
		for i := 0; i < N; i++ {
			avv := 0.0
			for j := 0; j < N; j++ {
				avv += y_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
			}
			avvv[i] = avv
		}

		fw_i := make([]float64, N)
		for i := 0; i < N; i++ {
			fw_i[i] = math.Exp(math.Log(y_i[i]*P) - math.Log(Z_v+Cw-Bw) + (Biw[i]-Ciw[i])/(Z_v+Cw-Bw) - (Aw/Bw)*((2*avvv[i]/aw)-(b_i[i]/bw))*math.Log((Z_v+Bw+Cw)/(Z_v+Cw)) - (Aw/Bw)*(Biw[i]+Ciw[i])/(Z_v+Bw+Cw) + (Aw/Bw)*Ciw[i]/(Z_v+Cw))
		}

		Ri := make([]float64, N)
		for i := 0; i < N; i++ {
			Ri[i] = fz_i[i] / (Sv * fw_i[i])
		}

		Ri_v = 0.0
		for i := 0; i < N; i++ {
			Ri_v += math.Pow((Ri[i] - 1), 2)
		}

		if Ri_v < math.Pow(10, -12) {
			m = 30
		}

		K_i = multiply(K_i, Ri)
		TS_v := 0.0

		for _, k := range K_i {
			TS_v += math.Pow(math.Log(k), 2)
		}

		if TS_v < math.Pow(10, -4) {
			TS_v_flag = 1
			m = 30
		}

		m++

	}

	copy(K_iv, K_i)

	for i := 0; i < N; i++ {
		K_i[i] = math.Pow(math.Exp(5.373*(1+w[i])*(1-Tkr[i]/T))*Pkr[i]/P, 1.0) //da
	}

	aw = 0.0
	bw = 0.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			aw += z[i] * z[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
		}
	}

	for i := 0; i < N; i++ {
		bw += z[i] * b_i[i]
	}

	Aw = aw * P / (math.Pow(R, 2) * math.Pow(T, 2))
	Bw = bw * P / (R * T)
	cw = 0.0
	for i := 0; i < N; i++ {
		cw += c_i[i] * z[i]
	}
	Cw = cw * P / (R * T)

	for i := 0; i < N; i++ {
		Biw[i] = b_i[i] * P / (R * T)
		Ciw[i] = c_i[i] * P / (R * T)
	}

	coefficients = []float64{1, 3*Cw - 1, 3*math.Pow(Cw, 2) - math.Pow(Bw, 2) - 2*Cw - Bw + Aw, math.Pow(Cw, 3) - math.Pow(Bw, 2)*Cw - math.Pow(Cw, 2) - Bw*Cw + Aw*Cw - Aw*Bw}
	cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])
	Z_v = findMax(cubroot)

	// Расчет летучестей в паровой фазе

	avvv = make([]float64, N)
	for i := 0; i < N; i++ {
		avv := 0.0
		for j := 0; j < N; j++ {
			avv += z[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
		}
		avvv[i] = avv
	}

	fz_i = make([]float64, N)
	for i := 0; i < N; i++ {
		fz_i[i] = math.Exp(math.Log(z[i]*P) - math.Log(Z_v+Cw-Bw) + (Biw[i]-Ciw[i])/(Z_v+Cw-Bw) - (Aw/Bw)*((2*avvv[i]/aw)-(b_i[i]/bw))*math.Log((Z_v+Bw+Cw)/(Z_v+Cw)) - (Aw/Bw)*(Biw[i]+Ciw[i])/(Z_v+Bw+Cw) + (Aw/Bw)*Ciw[i]/(Z_v+Cw))
	}

	// Часть 2 Проверка жидкой фазы 227 stroke matlab

	ml := 0
	Ri_l := 1.0

	for ml < 30 {

		Sl1 := 0.0
		for i := 0; i < N; i++ {
			Yi_l[i] = z[i] / K_i[i]
			Sl1 += Yi_l[i]
		}
		Sl = Sl1

		for i := 0; i < N; i++ {
			x_i[i] = Yi_l[i] / Sl1
		}

		al := 0.0
		bl := 0.0
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				al += x_i[i] * x_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
			}
		}

		for i := 0; i < N; i++ {
			bl += x_i[i] * b_i[i]
		}

		Al := al * P / (math.Pow(R, 2) * math.Pow(T, 2))
		Bl := bl * P / (R * T)
		cl := 0.0
		for i := 0; i < N; i++ {
			cl += c_i[i] * x_i[i]
		}
		Cl := cl * P / (R * T)

		for i := 0; i < N; i++ {
			Bil[i] = b_i[i] * P / (R * T)
			Cil[i] = c_i[i] * P / (R * T)
		}

		coefficients := []float64{1, (3*Cl - 1), (3*Cl*Cl - Bl*Bl - 2*Cl - Bl + Al), (Cl*Cl*Cl - Bl*Bl*Cl - Cl*Cl - Bl*Cl + Al*Cl - Al*Bl)}
		var cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])
		Z_l = findMin(cubroot)

		// Расчет летучестей в жидкой фазе

		for i := 0; i < N; i++ {
			all := 0.0
			for j := 0; j < N; j++ {
				all += x_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
			}
			alll[i] = all
		}
		//fl_i := make([]float64, N)
		for i := 0; i < N; i++ {
			fl_i[i] = math.Exp(math.Log(x_i[i]*P) - math.Log(Z_l+Cl-Bl) + (Bil[i]-Cil[i])/(Z_l+Cl-Bl) - (Al/Bl)*((2*alll[i]/al)-(b_i[i]/bl))*math.Log((Z_l+Bl+Cl)/(Z_l+Cl)) - (Al/Bl)*(Bil[i]+Cil[i])/(Z_l+Bl+Cl) + (Al/Bl)*Cil[i]/(Z_l+Cl))
		}
		//Ri := make([]float64, N)
		for i := 0; i < N; i++ {
			Ri[i] = Sl * fl_i[i] / fz_i[i]
		}
		Ri_l = 0.0
		for i := 0; i < N; i++ {
			Ri_l += math.Pow(-1, 2)
		}

		if Ri_l < math.Pow(10, -12) {
			m = 30
		}

		K_i = multiply(K_i, Ri)
		TS := 0.0

		for i := range K_i {
			TS += math.Pow(math.Log(K_i[i]), 2)
		}

		if TS < math.Pow(10, -4) {
			TS_l_flag = 1
			m = 30
		}
		ml++
	}

	copy(K_il, K_i)

	if (TS_l_flag == 1 && TS_v_flag == 1) || (Sv <= 1 && TS_l_flag == 1) || (Sl <= 1 && TS_v_flag == 1) || (Sv < 1 && Sl <= 1) {
		Stable = 1 //Stable
	} else {
		Stable = 0 //Unstable
	}

	if Stable == 0 {
		for i := 0; i < N; i++ {
			ac_i[i] = 0.42747 * math.Pow(R, 2) * math.Pow(Tkr[i], 2) / Pkr[i]
			psi_i[i] = 0.48 + 1.574*w[i] - 0.176*math.Pow(w[i], 2)
			alpha_i[i] = math.Pow(1+psi_i[i]*(1-math.Sqrt(T/Tkr[i])), 2)
			a_i[i] = ac_i[i] * alpha_i[i]
			b_i[i] = 0.08664 * R * Tkr[i] / Pkr[i]
			c_i[i] = cpen[i]
		}

		Kst_v := sum(square(subtract(K_iv, 1)))
		Kst_l := sum(square(subtract(K_il, 1)))

		if Kst_l > Kst_v {
			K_i = K_il
		} else {
			K_i = K_iv
		}

		m := 0
		eps_f := 1.0

		for eps_f > 0.000001 && m < 200 {

			// Шаг 1 Нахождение общей доли пара

			W = findRoot(z, K_i)

			// Шаг 2 Нахождение мольных долей xi, yi

			for i := 0; i < N; i++ {
				x_i[i] = z[i] / (1 + W*(K_i[i]-1))
			}

			for i := 0; i < N; i++ {
				y_i[i] = K_i[i] * x_i[i]
			}

			// Шаг 3 Нахождение z-фактора

			aw = 0.0
			bw = 0.0
			for i := 0; i < N; i++ {
				for j := 0; j < N; j++ {
					aw += y_i[i] * y_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
				}
			}

			for i := 0; i < N; i++ {
				bw += y_i[i] * b_i[i]
			}

			Aw = aw * P / (math.Pow(R, 2) * math.Pow(T, 2))
			Bw = bw * P / (R * T)
			cw = 0.0
			for i := 0; i < N; i++ {
				cw += c_i[i] * y_i[i]
			}
			Cw = cw * P / (R * T)

			for i := 0; i < N; i++ {
				Biw[i] = b_i[i] * P / (R * T)
				Ciw[i] = c_i[i] * P / (R * T)
			}

			coefficients := []float64{1, 3*Cw - 1, 3*math.Pow(Cw, 2) - math.Pow(Bw, 2) - 2*Cw - Bw + Aw, math.Pow(Cw, 3) - math.Pow(Bw, 2)*Cw - math.Pow(Cw, 2) - Bw*Cw + Aw*Cw - Aw*Bw}
			var cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])
			Z_v = findMax(cubroot)

			/* fmt.Printf("5 Roots: %.4f, %.4f, %.4f\n", cubroot[0], cubroot[1], cubroot[2])
			   fmt.Printf("5 Z_v %.4f\n", Z_v)*/

			//avvv = make([]float64, N)
			for i := 0; i < N; i++ {
				avv := 0.0
				for j := 0; j < N; j++ {
					avv += y_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
				}
				avvv[i] = avv
			}

			//fw_i := make([]float64, N)
			for i := 0; i < N; i++ {
				fw_i[i] = math.Exp(math.Log(y_i[i]*P) - math.Log(Z_v+Cw-Bw) + (Biw[i]-Ciw[i])/(Z_v+Cw-Bw) - (Aw/Bw)*((2*avvv[i]/aw)-(b_i[i]/bw))*math.Log((Z_v+Bw+Cw)/(Z_v+Cw)) - (Aw/Bw)*(Biw[i]+Ciw[i])/(Z_v+Bw+Cw) + (Aw/Bw)*Ciw[i]/(Z_v+Cw))
			}

			// Шаг 5 Нахождение z-фактора

			al := 0.0
			bl := 0.0
			for i := 0; i < N; i++ {
				for j := 0; j < N; j++ {
					al += x_i[i] * x_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
				}
			}

			for i := 0; i < N; i++ {
				bl += x_i[i] * b_i[i]
			}

			Al := al * P / (math.Pow(R, 2) * math.Pow(T, 2))
			Bl := bl * P / (R * T)
			cl := 0.0
			for i := 0; i < N; i++ {
				cl += c_i[i] * x_i[i]
			}
			Cl := cl * P / (R * T)

			for i := 0; i < N; i++ {
				Bil[i] = b_i[i] * P / (R * T)
				Cil[i] = c_i[i] * P / (R * T)
			}

			coefficients = []float64{1, (3*Cl - 1), (3*Cl*Cl - Bl*Bl - 2*Cl - Bl + Al), (Cl*Cl*Cl - Bl*Bl*Cl - Cl*Cl - Bl*Cl + Al*Cl - Al*Bl)}
			cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])
			Z_l = findMin(cubroot)

			for i := 0; i < N; i++ {
				all := 0.0
				for j := 0; j < N; j++ {
					all += x_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
				}
				alll[i] = all
			}

			for i := 0; i < N; i++ {
				fl_i[i] = math.Exp(math.Log(x_i[i]*P) - math.Log(Z_l+Cl-Bl) + (Bil[i]-Cil[i])/(Z_l+Cl-Bl) - (Al/Bl)*((2*alll[i]/al)-(b_i[i]/bl))*math.Log((Z_l+Bl+Cl)/(Z_l+Cl)) - (Al/Bl)*(Bil[i]+Cil[i])/(Z_l+Bl+Cl) + (Al/Bl)*Cil[i]/(Z_l+Cl))
			}

			// Корректировка распределения Ki

			for i := 0; i < N; i++ {
				if fl_i[i] != 0 {
					K_i[i] *= fl_i[i] / fw_i[i]
					df_lv[i] = fl_i[i]/fw_i[i] - 1
				}
			}

			eps_f = maxAbs(df_lv)

			if eps_f < 0.000001 {
				break
			}

			m++
		}

	}

	if Stable == 1 {
		Volume := 1000 * (Z_init * R * T / P)
		if Volume*T*T > Control_phase_MF { //тогда газ
			W = 1.0
			Z_l = -9999
			Z_v = Z_init
			copy(y_i, z)
			for i := 0; i < N; i++ {
				x_i[i] = 0.0
			}
		} else { //тогда жидкость
			W = 0.0
			//L:=1.0
			Z_v = -9999
			Z_l = Z_init
			copy(x_i, z)
			for i := 0; i < N; i++ {
				y_i[i] = 0.0
			}
		}

		/*for i := 0; i < N; i++ {
			K_i[i] = math.Pow(math.Exp(5.373*(1+w[i])*(1-Tkr[i]/T))*Pkr[i]/P, 1.0)
		}
		ZK_Mult := sum(multiply(z, K_i)) - 1

		// Проверка на однофазное состояние жидкости
		if ZK_Mult < 0 {
			W = 0.0
			//L:=1.0
			Z_v = -9999
			Z_l = Z_init
			copy(x_i, z)
			for i := 0; i < N; i++ {
				y_i[i] = 0.0
			}
		}

		// Проверка на однофазное состояние газ

		if ZK_Mult > 0 {
			W = 1.0
			Z_l = -9999
			Z_v = Z_init
			copy(y_i, z)
			for i := 0; i < N; i++ {
				x_i[i] = 0.0
			}
		}*/

	}
	return &pb.VleMessageResponse{L: 1 - W, ZL: Z_l, ZV: Z_v, XI: x_i, YI: y_i}, nil
}

/*func checkSum(arr []float64) bool {
	sum := sum(arr)
	return math.Abs(sum-1.0) < math.Pow(10, -12)
}*/

// Сумма элементов одномерного массива
func sum(arr []float64) float64 {
	sum := 0.0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// Вычесть число из каждого элемента массива
func subtract(arr []float64, val float64) []float64 {
	result := make([]float64, len(arr))
	for i, v := range arr {
		result[i] = v - val
	}
	return result
}

// Возвести в квадрат каждый элемент массива
func square(arr []float64) []float64 {
	result := make([]float64, len(arr))
	for i, v := range arr {
		result[i] = math.Pow(v, 2)
	}
	return result
}

// Перемножить элементы двух массивов
func multiply(arr1, arr2 []float64) []float64 {
	result := make([]float64, len(arr1))
	for i := range arr1 {
		result[i] = arr1[i] * arr2[i]
	}
	return result
}

// Максимальный по модулю элемент массива
func maxAbs(arr []float64) float64 {
	max := math.Abs(arr[0])
	for _, val := range arr {
		if math.Abs(val) > max {
			max = math.Abs(val)
		}
	}
	return max
}

// Решение уравнения Ричфорда-Райса
func findRoot(z []float64, K []float64) float64 {

	FvMin := 1 / (1 - findMax(K))
	FvMax := 1 / (1 - findMin(K))

	a := FvMin + 0.00001
	b := FvMax - 0.00001

	X := (a + b) / 2
	fa := 0.0
	fb := 0.0
	fX := 0.0
	for i := range z {
		fa += z[i] * (K[i] - 1) / (1 + a*(K[i]-1))
		fb += z[i] * (K[i] - 1) / (1 + b*(K[i]-1))
		fX += z[i] * (K[i] - 1) / (1 + X*(K[i]-1))
	}

	for math.Abs(a-b) > 0.0000001 {
		fa := 0.0
		fb := 0.0
		fX := 0.0
		for i := range z {
			fa += z[i] * (K[i] - 1) / (1 + a*(K[i]-1))
			fb += z[i] * (K[i] - 1) / (1 + b*(K[i]-1))
			fX += z[i] * (K[i] - 1) / (1 + X*(K[i]-1))
		}

		if fa*fX < 0 {
			b = X
		} else if fb*fX < 0 {
			a = X
		}
		X = (a + b) / 2
	}
	return X
}

// Кубический корень из числа
func cubeRoot(x float64) float64 {
	if x >= 0 {
		return math.Pow(x, 1./3.)
	} else {
		return -math.Pow(-x, 1./3.)
	}
}

// Решение кубического уравнения на z
func cubicEquationSolver(A, B, C, D float64) []float64 {
	x := make([]float64, 3)
	d := 18*A*B*C*D - 4*math.Pow(B, 3)*D + math.Pow(B, 2)*math.Pow(C, 2) - 4*A*math.Pow(C, 3) - 27*math.Pow(A, 2)*math.Pow(D, 2)
	P := math.Pow(B, 2) - 3*A*C
	Q := 9*A*B*C - 2*math.Pow(B, 3) - 27*math.Pow(A, 2)*D
	if d > 0 {
		D1 := (2*math.Pow(B/A, 3) - 9*(B/A)*(C/A) + 27*(D/A)) / 54
		D2 := (math.Pow(B/A, 2) - 3*(C/A)) / 9
		theta := math.Acos(D1 / math.Sqrt(math.Pow(D2, 3)))
		x[0] = -2*math.Sqrt(D2)*math.Cos(theta/3) - B/(3*A)
		x[1] = -2*math.Sqrt(D2)*math.Cos((theta+2*math.Pi)/3) - B/(3*A)
		x[2] = -2*math.Sqrt(D2)*math.Cos((theta-2*math.Pi)/3) - B/(3*A)
		//fmt.Printf("The cubic equation has three distinct real roots: x1 = %.4f, x2 = %.4f, and x3 = %.4f\n", x1, x2, x3)
	} else if d < 0 {
		N := cubeRoot(Q/2+math.Sqrt(math.Pow(Q, 2)/4-math.Pow(P, 3))) + cubeRoot(Q/2-math.Sqrt(math.Pow(Q, 2)/4-math.Pow(P, 3)))
		x[0] = -B/(3*A) + N/(3*A)
		x[1] = x[0]
		x[2] = x[0]
		//fmt.Printf("The cubic equation has one real root x = %.4f and two (non-real) complex roots z1 = %.4f and z2 = %.4f\n", x1, x2, x3)
	} else {
		// fmt.Println("The cubic equation has a multiple root, and all of its roots are real!")
		if P == 0 {
			x[0] = -B / (3 * A)
			x[1] = -B / (3 * A)
			x[2] = -B / (3 * A)
			// fmt.Printf("In this case the cubic equation has one triple root x = %.2f\n", x1)
		} else {
			x[0] = (9*A*D - B*C) / (2 * P)
			x[1] = (9*A*D - B*C) / (2 * P)
			x[2] = (4*A*B*C - 9*A*A*D - B*B*B) / (A * P)
			//fmt.Printf("In this case the cubic equation has a double root xd = %.2f and a single root xs = %.4f\n", x1, x3)
		}
	}
	return x
}

// Поиск максимального значения массива
func findMax(a []float64) (max float64) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

// Поиск минимального значения массива
func findMin(a []float64) (min float64) {
	min = a[0]
	for _, value := range a {
		if value < min && value >= 0 {
			min = value
		}
	}
	return min
}
