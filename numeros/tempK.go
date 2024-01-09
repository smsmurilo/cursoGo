package main

import "fmt"

// valor constante para ebolição Kelvin
const ebulicaoK float64 = 373.0

func main() {
	// variáveis de valores temp kelvin e temp centígrados
	var tempK float64 = ebulicaoK
	// cálculo da conversão de kelvin para centígrados
	var tempC float64 = (tempK - 32) * 5 / 9

	fmt.Println("A temperatura de ebulição da água em graus K = ", tempK)
	fmt.Println("A temperatura de ebulição da água em graus C = ", tempC)

}
