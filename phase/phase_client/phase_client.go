package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "phase"
)

type InitType struct {
	FLUID []*pb.OneComponent
	BIPs  []*pb.OneBIP
}

func main() {
	serverAddress := "localhost:50051"

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}

	defer conn.Close()

	client := pb.NewPhaseEqualibriumClient(conn)

	var inData InitType
	//Mixture 3666.9
	inData = InitType{
		FLUID: []*pb.OneComponent{
			{Component: "N2", MoleFraction: 0.33, MolMass: 28.013514, Tcr: -146.95, Pcr: 3.394388, WFact: 0.04, Tb: -195.75, Vcr: 89.8, Pen: 0.92},
			{Component: "CO2", MoleFraction: 0.846, MolMass: 44.0098, Tcr: 31.04995, Pcr: 7.376459, WFact: 0.225, Tb: -78.5, Vcr: 94, Pen: 3.028},
			{Component: "C1", MoleFraction: 74.943, MolMass: 16.042879, Tcr: -82.55, Pcr: 4.600155, WFact: 0.008, Tb: -161.55, Vcr: 99.00002, Pen: 0.63},
			{Component: "C2", MoleFraction: 7.642, MolMass: 30.069818, Tcr: 32.24996, Pcr: 4.883864, WFact: 0.098, Tb: -88.55, Vcr: 148, Pen: 2.63},
			{Component: "C3", MoleFraction: 4.299, MolMass: 44.096748, Tcr: 96.64996, Pcr: 4.245517, WFact: 0.152, Tb: -42.05, Vcr: 203, Pen: 5.059999},
			{Component: "iC4", MoleFraction: 0.917, MolMass: 58.123703, Tcr: 134.9501, Pcr: 3.647701, WFact: 0.176, Tb: -11.75, Vcr: 263, Pen: 7.290001},
			{Component: "nC4", MoleFraction: 1.537, MolMass: 58.123711, Tcr: 152.0501, Pcr: 3.799688, WFact: 0.193, Tb: -0.44993, Vcr: 255, Pen: 7.86},
			{Component: "iC5", MoleFraction: 0.524, MolMass: 72.15065, Tcr: 187.2499, Pcr: 3.384255, WFact: 0.227, Tb: 27.85003, Vcr: 306, Pen: 10.93},
			{Component: "nC5", MoleFraction: 0.606, MolMass: 72.15065, Tcr: 196.45, Pcr: 3.374123, WFact: 0.251, Tb: 36.0501, Vcr: 304.0001, Pen: 12.18},
			{Component: "C6", MoleFraction: 0.699, MolMass: 86.177597, Tcr: 234.2501, Pcr: 2.968823, WFact: 0.296, Tb: 68.75005, Vcr: 370.0001, Pen: 17.98},
			{Component: "C7", MoleFraction: 1.228, MolMass: 91.924507, Tcr: 273.4901, Pcr: 3.574564, WFact: 0.436319, Tb: 91.95001, Vcr: 425.4445, Pen: 6.720626},
			{Component: "C8", MoleFraction: 1.453, MolMass: 105.71749, Tcr: 296.7907, Pcr: 3.122868, WFact: 0.472542, Tb: 116.75, Vcr: 471.2443, Pen: 13.03526},
			{Component: "C9", MoleFraction: 0.751, MolMass: 120.34306, Tcr: 318.7964, Pcr: 2.771944, WFact: 0.510256, Tb: 142.25, Vcr: 523.4991, Pen: 19.40622},
			{Component: "C10-C11", MoleFraction: 0.977, MolMass: 140.56026, Tcr: 327.7199, Pcr: 2.50968, WFact: 0.616047, Tb: 175.324, Vcr: 600.5144, Pen: 15.12194},
			{Component: "C12-C13", MoleFraction: 0.722, MolMass: 169.0862, Tcr: 360.6547, Pcr: 2.19141, WFact: 0.692018, Tb: 218.0582, Vcr: 712.9235, Pen: 19.6617},
			{Component: "C14-C15", MoleFraction: 0.598, MolMass: 199.09802, Tcr: 391.3671, Pcr: 1.972617, WFact: 0.768374, Tb: 256.4565, Vcr: 838.2478, Pen: 20.81166},
			{Component: "C16-C18", MoleFraction: 0.575, MolMass: 236.89189, Tcr: 426.3225, Pcr: 1.792949, WFact: 0.859284, Tb: 298.2033, Vcr: 1003.597, Pen: 17.46782},
			{Component: "C19-C20", MoleFraction: 0.302, MolMass: 269.75214, Tcr: 454.0104, Pcr: 1.694122, WFact: 0.931732, Tb: 330.9889, Vcr: 1147.239, Pen: 10.17628},
			{Component: "C21-C24", MoleFraction: 0.409, MolMass: 310.63672, Tcr: 494.3136, Pcr: 1.573101, WFact: 0.972036, Tb: 367.5937, Vcr: 1337.256, Pen: 9.111677},
			{Component: "C25-C29", MoleFraction: 0.321, MolMass: 372.29236, Tcr: 540.0753, Pcr: 1.488414, WFact: 1.076324, Tb: 416.9652, Vcr: 1629.407, Pen: -14.2148},
			{Component: "C30-C35", MoleFraction: 0.221, MolMass: 448.23227, Tcr: 591.9588, Pcr: 1.424685, WFact: 1.174966, Tb: 464.6909, Vcr: 1997.867, Pen: -49.8337},
			{Component: "C36-C80", MoleFraction: 0.102, MolMass: 701.12256, Tcr: 766.4369, Pcr: 1.351062, WFact: 1.109354, Tb: 596.6096, Vcr: 3509.115, Pen: -186.03}},

		BIPs: []*pb.OneBIP{
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
			{Num: []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}},
		},
	}

	if err != nil {
		log.Fatalf("Array RPC failed: %v", err)
	}

	client.Init(context.Background(), &pb.InitMessageRequest{
		FLUID: inData.FLUID,
		BIPs:  inData.BIPs,
	})

	Ttest := makeRange(223, 408, 1)
	Ptest := makeRange(0.1, 61., 0.1)

	// Создаем текстовый файл для параметров T, P и L
	file, err := os.Create("TP_L.txt")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()

	// Создаем текстовый файл для параметров T, P и SECONDS
	fileSeconds, err := os.Create("TP_SECONDS.txt")
	if err != nil {
		log.Fatalf("error creating fileSeconds: %v", err)
	}
	defer fileSeconds.Close()

	// начало алгоритма

	for _, temp := range Ttest {
		file.WriteString(fmt.Sprintf("%3.5f\t", temp))
		fileSeconds.WriteString(fmt.Sprintf("%3.5f\t", temp))
	}

	file.WriteString(fmt.Sprintf("\n"))
	fileSeconds.WriteString(fmt.Sprintf("\n"))

	for _, pres := range Ptest {
		file.WriteString(fmt.Sprintf("%3.5f\t", pres))
		fileSeconds.WriteString(fmt.Sprintf("%3.5f\t", pres))
		for _, temp := range Ttest {

			start := time.Now()
			result_vle, err := client.Vle(context.Background(), &pb.VleMessageRequest{Temp: temp, Pres: pres})
			if err != nil {
				log.Fatalf("VLE is not done: %v", err)
			}

			duration := time.Since(start)
			// вывод на консоль

			// вывод на консоль закончилось

			// создание файла началось
			// Записываем T, P и L в первый файл
			_, err = file.WriteString(fmt.Sprintf("%1.4f\t", 1-result_vle.L))
			if err != nil {
				log.Fatalf("error writing to file: %v", err)
			}

			// Записываем T, P и SECONDS во второй файл
			_, err = fileSeconds.WriteString(fmt.Sprintf("%1.7f\t", duration.Seconds()))
			if err != nil {
				log.Fatalf("error writing to fileSeconds: %v", err)
			}

			// создание файла закончилось

		}

		file.WriteString(fmt.Sprintf("\n"))
		fileSeconds.WriteString(fmt.Sprintf("\n"))
	}

}

func makeRange(min, max, step float64) []float64 {
	size := int((max-min)/step) + 1
	r := make([]float64, size)
	for i := range r {
		r[i] = min + float64(i)*step
	}
	return r
}
