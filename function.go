package main

import "fmt"
import "math"

func main() {
	fmt.Println("test...")
	//t126()
	t326()
}


type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func t126() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func t326() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
