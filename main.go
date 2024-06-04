package main // основной пакет программы, который является точкой входа. main - название

import ( // импорт необходимых библиотек для работы программы
	"fmt"
	"math"
	"os"
	"time"
)

func main() { // основная функция main, где написан основной код
	start := time.Now()     // функция отвечающая за замер времени выполнения кода
	var W, Z_l, Z_v float64 // объявление переменных
	var T, P float64        // объявление

	Ttest := makeRange(400, 400, 1) // Диапазон температур (начальная точка, конечная точка, шаг)
	Ptest := makeRange(1., 65., 1)  //  Диапазон давления (начальная точка, конечная точка, шаг)
	// mixture 1. 1 состав из ТЗ
	z := []float64{0.228, 0.605, 83.578, 7.4, 3.345, 0.755, 0.962, 0.338, 0.316, 0.356, 0.483, 0.593, 0.275, 0.297, 0.163, 0.112, 0.09, 0.035, 0.039, 0.021, 0.008, 0.001}
	// mixture 3. 3 состав из ТЗ
	// z := []float64{0.258, 0.457, 55.06, 6.964, 5.087, 0.808, 2.442, 0.614, 0.829, 0.675, 1.839, 1.933, 1.074, 1.536, 1.192, 1.108, 1.446, 0.927, 1.681, 2.21, 3.08, 8.78}

	z = normalizeZ(z) // переводим с процентов на число, путем деления всех чисел массива на 100

	N := len(z) // количество компонентов. len -длина массива

	// Наши данные по ТЗ у составов
	Pkr := []float64{3.394388, 7.376459, 4.600155, 4.883864, 4.245517, 3.647701, 3.799688, 3.384255, 3.374123, 2.968823, 3.574564, 3.122868, 2.771944, 2.50968, 2.19141, 1.972617, 1.792949, 1.694122, 1.573101, 1.488414, 1.424685, 1.351062}
	Tkr := []float64{126.2, 304.19995, 190.6, 305.39996, 369.79996, 408.1001, 425.2001, 460.3999, 469.6, 507.4001, 546.6401, 569.9407, 591.9464, 600.8699, 633.8047, 664.5171, 699.4725, 727.1604, 767.4636, 813.2253, 865.1088, 1039.5869}
	w := []float64{0.04, 0.225, 0.008, 0.098, 0.152, 0.176, 0.193, 0.227, 0.251, 0.296, 0.436319, 0.472542, 0.510256, 0.616047, 0.692018, 0.768374, 0.859284, 0.931732, 0.972036, 1.076324, 1.174966, 1.109354}
	cpen := []float64{0.92, 3.03, 0.63, 2.63, 5.06, 7.29, 7.86, 10.93, 12.18, 17.98, 6.72, 13.03, 19.41, 15.12, 19.66, 20.81, 17.47, 10.18, 9.11, -14.21, -49.83, -186.03}
	cpen = normalizeCpen(cpen) // переводим с процентов на число, путем деления всех чисел массива на 100

	// создание среза slice, массива данных для дальнейшей работы с массивами. N - какая длина массива, смотри выше строку с N
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

	c := make([][]float64, N) // 2-мерный массив

	for i := range c {
		c[i] = make([]float64, N)
	}

	R := 0.00831675 // константа

	fid5, err := os.Create("Mesh.txt") // создание текстового файла Mesh с нашими результатами T P Z_v Z_l W
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fid5.Close()

	fmt.Fprintln(fid5, "T P Z_v Z_l W") // Шапка

	fid6, err := os.Create("y_i.txt") // создание текстового файла y_i с нашими составами компонент
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fid6.Close()

	fmt.Fprintln(fid6, "T P W y_i") // Шапка

	step := 0 // начало алгоритма с массивами

	Mesh := make([][]float64, len(Ttest)*len(Ptest)) // создание наших P T в виде массива
	for i := range Mesh {
		Mesh[i] = make([]float64, 3)
	}

	for hT := 0; hT < len(Ttest); hT++ { // Цикл для всех составов который прогоняет по нашим данным которые мы задаем в Ttest
		T = Ttest[hT]
		for hP := 0; hP < len(Ptest); hP++ { // Цикл всех составов который прогоняет по нашим данным которые мы задаем в Ptest
			P = Ptest[hP]

			step++

			for i := 0; i < N; i++ { // формулу для коэффициентов
				ac_i[i] = 0.42747 * math.Pow(R, 2) * math.Pow(Tkr[i], 2) / Pkr[i]
				psi_i[i] = 0.48 + 1.574*w[i] - 0.176*math.Pow(w[i], 2)
				alpha_i[i] = math.Pow(1+psi_i[i]*(1-math.Sqrt(T/Tkr[i])), 2)
				a_i[i] = ac_i[i] * alpha_i[i]
				b_i[i] = 0.08664 * R * Tkr[i] / Pkr[i]
				c_i[i] = cpen[i]
			}

			for i := 0; i < N; i++ {
				K_i[i] = math.Pow(math.Exp(5.373*(1+w[i])*(1-Tkr[i]/T))*Pkr[i]/P, 1.0) // начальные значенения коэффициентов распределения
			}

			aw := 0.0 // начальное значение для работы цикла
			bw := 0.0
			for i := 0; i < N; i++ {
				for j := 0; j < N; j++ {
					aw += z[i] * z[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j]) // коэффициент для уравнения Aw
				}
			}
			for i := 0; i < N; i++ {
				bw += z[i] * b_i[i] // коэффициент для уравнения Bw
			}
			Aw := aw * P / (math.Pow(R, 2) * math.Pow(T, 2)) // коэффициент для кубического уравнения состояния
			Bw := bw * P / (R * T)                           // коэффициент для кубического уравнения состояния
			cw := 0.0
			for i := 0; i < N; i++ {
				cw += c_i[i] * z[i] // коэффициент для уравнения Cw
			}
			Cw := cw * P / (R * T) // коэффициент для кубического уравнения состояния

			for i := 0; i < N; i++ {
				Biw[i] = b_i[i] * P / (R * T) // коэффициент для уравнения летучести компонентов
				Ciw[i] = c_i[i] * P / (R * T) // коэффициент для уравнения летучести компонентов
			}

			coefficients := []float64{1, 3*Cw - 1, 3*math.Pow(Cw, 2) - math.Pow(Bw, 2) - 2*Cw - Bw + Aw, math.Pow(Cw, 3) - math.Pow(Bw, 2)*Cw - math.Pow(Cw, 2) - Bw*Cw + Aw*Cw - Aw*Bw} // расчет кубического уравнения состояния
			var cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])                                                                        // в конце кода расписана сама функция
			Z_v = findMax(cubroot)                                                                                                                                                       // максимальное значение по куб. ур. для пара

			for i := 0; i < N; i++ {
				avv := 0.0
				for j := 0; j < N; j++ {
					avv += z[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j]) // кусок уравнения из уравнения летучести компонентов, для удобного и правильного расчета
				}
				avvv[i] = avv
			}

			fz_i := make([]float64, N)
			for i := 0; i < N; i++ {
				fz_i[i] = math.Exp(math.Log(z[i]*P) - math.Log(Z_v+Cw-Bw) + (Biw[i]-Ciw[i])/(Z_v+Cw-Bw) - (Aw/Bw)*((2*avvv[i]/aw)-(b_i[i]/bw))*math.Log((Z_v+Bw+Cw)/(Z_v+Cw)) - (Aw/Bw)*(Biw[i]+Ciw[i])/(Z_v+Bw+Cw) + (Aw/Bw)*Ciw[i]/(Z_v+Cw)) // уравнение летучести компонентов в паровой фазе
			}

			// задаем для проверки и расчета стабильности

			m := 0

			Ri_v := 1.0
			TS_v_flag := 0
			TS_l_flag := 0
			var Sv, Sl float64

			// Часть 1 Проверка газовой фазы

			for m < 30 { // 30 итераций

				Yi_v := make([]float64, N) // компонент
				Sv1 := 0.0
				for i := 0; i < N; i++ {
					Yi_v[i] = z[i] * K_i[i]
					Sv1 += Yi_v[i]
				}

				Sv = Sv1

				for i := 0; i < N; i++ {
					y_i[i] = Yi_v[i] / Sv
				}

				aw = 0.0 // аналогичный расчет
				bw = 0.0
				for i := 0; i < N; i++ {
					for j := 0; j < N; j++ {
						aw += y_i[i] * y_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j]) // **
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

				// кубическое уравнение состояние, путем нахождения 3 корней
				coefficients := []float64{1, 3*Cw - 1, 3*math.Pow(Cw, 2) - math.Pow(Bw, 2) - 2*Cw - Bw + Aw, math.Pow(Cw, 3) - math.Pow(Bw, 2)*Cw - math.Pow(Cw, 2) - Bw*Cw + Aw*Cw - Aw*Bw}
				var cubroot = cubicEquationSolver(coefficients[0], coefficients[1], coefficients[2], coefficients[3])
				Z_v = findMax(cubroot) // максимальное значение по куб. ур. для пара

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
				TS_v := 0.0 // trivial solution

				for _, k := range K_i {
					TS_v += math.Pow(math.Log(k), 2)
				}

				if TS_v < math.Pow(10, -4) {
					TS_v_flag = 1
					m = 30
				}

				m++

			}

			copy(K_iv, K_i) // скопировать данные в K_i

			for i := 0; i < N; i++ {
				K_i[i] = math.Pow(math.Exp(5.373*(1+w[i])*(1-Tkr[i]/T))*Pkr[i]/P, 1.0) // **
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

			// Часть 2 Проверка жидкой фазы

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

				for i := 0; i < N; i++ {
					fl_i[i] = math.Exp(math.Log(x_i[i]*P) - math.Log(Z_l+Cl-Bl) + (Bil[i]-Cil[i])/(Z_l+Cl-Bl) - (Al/Bl)*((2*alll[i]/al)-(b_i[i]/bl))*math.Log((Z_l+Bl+Cl)/(Z_l+Cl)) - (Al/Bl)*(Bil[i]+Cil[i])/(Z_l+Bl+Cl) + (Al/Bl)*Cil[i]/(Z_l+Cl)) // уравнение летучести компонентов в жидкой фазе
				}

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

			// условие для проверки стабильности
			var Stable int

			if (TS_l_flag == 1 && TS_v_flag == 1) || (Sv <= 1 && TS_l_flag == 1) || (Sl <= 1 && TS_v_flag == 1) || (Sv < 1 && Sl <= 1) {
				Stable = 1 //Stable
			} else {
				Stable = 0 //Unstable
			}

			// если не стабильна
			if Stable == 0 {

				for i := 0; i < N; i++ {
					ac_i[i] = 0.42747 * math.Pow(R, 2) * math.Pow(Tkr[i], 2) / Pkr[i]
					psi_i[i] = 0.48 + 1.574*w[i] - 0.176*math.Pow(w[i], 2)
					alpha_i[i] = math.Pow(1+psi_i[i]*(1-math.Sqrt(T/Tkr[i])), 2)
					a_i[i] = ac_i[i] * alpha_i[i]
					b_i[i] = 0.08664 * R * Tkr[i] / Pkr[i]
					c_i[i] = cpen[i]
				}

				Kst_v := sum(square(subtract(K_iv, 1))) // **
				Kst_l := sum(square(subtract(K_il, 1))) // **

				if Kst_l > Kst_v { // **
					K_i = K_il
				} else {
					K_i = K_iv
				}

				m := 0
				eps_f := 1.0 // **

				for eps_f > 0.000001 && m < 50 { // условие

					// Шаг 1 Нахождение общей доли пара

					W = findRoot(z, K_i) // уравнение Ричарда-Райса

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
							aw += y_i[i] * y_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j]) // **
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

					for i := 0; i < N; i++ {
						avv := 0.0
						for j := 0; j < N; j++ {
							avv += y_i[j] * (1 - c[i][j]) * math.Sqrt(a_i[i]*a_i[j])
						}
						avvv[i] = avv
					}

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
						fl_i[i] = math.Exp(math.Log(x_i[i]*P) - math.Log(Z_l+Cl-Bl) + (Bil[i]-Cil[i])/(Z_l+Cl-Bl) - (Al/Bl)*((2*alll[i]/al)-(b_i[i]/bl))*math.Log((Z_l+Bl+Cl)/(Z_l+Cl)) - (Al/Bl)*(Bil[i]+Cil[i])/(Z_l+Bl+Cl) + (Al/Bl)*Cil[i]/(Z_l+Cl)) // уравнение летучести компонентов в жидкой фазе
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

			// если стабильна
			if Stable == 1 {
				for i := 0; i < N; i++ {
					K_i[i] = math.Pow(math.Exp(5.373*(1+w[i])*(1-Tkr[i]/T))*Pkr[i]/P, 1.0)
				}
				ZK_Mult := sum(multiply(z, K_i)) - 1

				// Проверка на однофазное состояние жидкости
				if ZK_Mult < 0 {
					W = 0.0
					Z_v = -9999
					copy(x_i, z)
				}

				// Проверка на однофазное состояние газ

				if ZK_Mult > 0 {
					W = 1.0
					Z_l = -9999
					copy(y_i, z)
				}

			}

			fmt.Fprintf(fid5, "%3.5f %3.5f %3.5f %3.5f %3.10f \t \n", T, P, Z_v, Z_l, W)
			fmt.Fprintf(fid6, "%4f \t \n", y_i)
		}

	}
	duration := time.Since(start)               // время работы кода
	fmt.Println("SECONDS:", duration.Seconds()) // вывод на консоль
}

// Деление каждого числа массива на 100
func normalizeZ(arr []float64) []float64 {
	for i := range arr {
		arr[i] /= 100
	}
	return arr
}

// Деление каждого числа массива на 1000
func normalizeCpen(arr []float64) []float64 {
	for i := range arr {
		arr[i] /= 1000
	}
	return arr
}

// Сумма элементов одномерного массива
func sum(arr []float64) float64 {
	sum := 0.0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// Вычетание числа из каждого элемента массива
func subtract(arr []float64, val float64) []float64 {
	result := make([]float64, len(arr))
	for i, v := range arr {
		result[i] = v - val
	}
	return result
}

// Возведение в квадрат каждый элемент массива
func square(arr []float64) []float64 {
	result := make([]float64, len(arr))
	for i, v := range arr {
		result[i] = math.Pow(v, 2)
	}
	return result
}

// Ппроизведение элементов двух массивов
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

	FvMin := 1 / (1 - max(K))
	FvMax := 1 / (1 - min(K))

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

	for math.Abs(a-b) > 0.000001 {
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

// Максимальное значение для уравнения Ричарда-Райса
func max(values []float64) float64 {
	maxValue := values[0]
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// Минимальное значение для уравнения Ричарда-Райса
func min(values []float64) float64 {
	minValue := values[0]
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}

// Задать T P (начальное число и конечное число, шаг)
func makeRange(min, max, step float64) []float64 {
	size := int((max-min)/step) + 1
	r := make([]float64, size)
	for i := range r {
		r[i] = min + float64(i)*step
	}
	return r
}

// Кубический корень из числа
func cubeRoot(x float64) float64 {
	if x >= 0 {
		return math.Pow(x, 1./3.)
	} else {
		return -math.Pow(-x, 1./3.)
	}
}

// Решение кубического уравнение состояния
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

// Поиск максимального значения массива z
func findMax(a []float64) (max float64) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

// Поиск минимального значения массива z
func findMin(a []float64) (min float64) {
	min = a[0]
	for _, value := range a {
		if value < min && value >= 0 {
			min = value
		}
	}
	return min
}
