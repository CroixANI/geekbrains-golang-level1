package main

import (
	"fmt"
	"math"
)

func getCircleDiameterByArea(circleAreaSize float64) float64 {
	return math.Sqrt(circleAreaSize/math.Pi) * 2.0
}

func getCircleCircumferenceByDiameter(circleDiameter float64) float64 {
	return circleDiameter * math.Pi
}

func main() {
	var circleAreaSize float64

	fmt.Println("This app calculate diameter and circumference by given area of a circle")
	fmt.Print("Enter number (circle area size): ")
	fmt.Scanln(&circleAreaSize)

	circleDiameter := getCircleDiameterByArea(circleAreaSize)
	fmt.Printf("Circle diameter is equal to: %f\n", circleDiameter)
	fmt.Printf("Circle circumference is equal to: %f\n", getCircleCircumferenceByDiameter(circleDiameter))
}
