package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"fmt"
	"math"
)

type sidesCountType int

func CalcSquare(sideLen float64, sidesNum sidesCountType) float64 {
	switch sidesNum {
	case 0:
		fmt.Println("Circle")
		return math.Pi * sideLen
	case 3:
		fmt.Println("Triangle")
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case 4:
		fmt.Println("Rectangle")
		return math.Pow(sideLen, 2)
	default:
		fmt.Println("Invalid number of sides")
		return 0.0
	}
}
