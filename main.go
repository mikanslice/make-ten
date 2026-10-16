package main

import "fmt"

type flaction struct{
	top, bottom uint
}

func mult(a flaction, b flaction) flaction{
	result := flaction{a.top*b.top, a.bottom*b.bottom}
	return clean(result)
}

func divi(a flaction, b flaction) flaction{
	result := flaction{a.top*b.bottom, a.bottom*b.top}
	return clean(result)	
}

func clean(a flaction) flaction{
	tmpGcm := gcm(a.top, a.bottom)
	return flaction{ a.top/tmpGcm, a.bottom/tmpGcm }
}

func gcm(a uint, b uint) uint{
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func main() {
	fmt.Println(divi(flaction{2,1},flaction{3,1}));
}