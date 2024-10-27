package main

import (
	"fmt"
	"lab2/classes"
)

func printShapeInfo(shape classes.Shape) {
	fmt.Println(shape)
}

func main() {
	n := 2.0
	rect := classes.NewRectangle(n, n, classes.NewColor("blue"))
	circ := classes.NewCircle(n, classes.NewColor("green"))
	sq := classes.NewSquare(n, classes.NewColor("red"))

	printShapeInfo(rect)
	printShapeInfo(circ)
	printShapeInfo(sq)
}
