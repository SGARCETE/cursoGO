package main

import (
	"fmt"
	"math"
)

const iteraciones = 10;
const diferencia = 0.00000000000001

type ErrNegativeSqrt float64                   //Acá definimos un nuevo tipo ErrNegativeSqrt basado en float64

// Acá redifinimos el tipo ErrNegativeSqrt agregandole un método de tipo Error (Ahora también es Error)

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprint("It cant be a negative number")
}

func SqrtWithError(x float64) (float64, error) {

	if x < 0{
		return x, ErrNegativeSqrt(x)
	}else{

		var z= 1.0
		var zi=0.0

		for i:=0; i<iteraciones;i++{
			zi= z
			z= newton(z,x)

			if (math.Abs(z-zi)< diferencia){
				break
			}
		}
	return z, nil;
	}
}

func Sqrt(x float64) float64 {
	var z= 1.0
	var zi=0.0

	for i:=0; i<iteraciones;i++{
		zi= z
		z= newton(z,x)

		if (math.Abs(z-zi)< diferencia){
			break
		}
	}
	return z;
}


func newton (z float64, x float64) float64{
	return z - (z*z - x) / (2 * z)
}


func main() {
	fmt.Println(SqrtWithError(-2))
	fmt.Println("")
	fmt.Println(math.Sqrt(-2))
}
