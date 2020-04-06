package main

import (
	"fmt"
)

//Converter represents the Converter
type Converter struct{}

//Centimeters is the unit in centimeters
type Centimeters float64

//Feet is the unit in feet
type Feet float64

//Minutes is the unit in minutes
type Minutes float64

//Seconds is the unit in seconds
type Seconds float64

//CentimetersToFeet converts centimeters to feet
func (cvr Converter) CentimetersToFeet(c Centimeters) Feet {
	//conversion code
	return Feet(float64(c) * 0.032)
}

//FeetToCentimeters converts feet to centimeters
func (cvr Converter) FeetToCentimeters(f Feet) Centimeters {
	//conversion code
	return Centimeters(float64(f) * 30.48)
}

//MinutesToSeconds converts Minutes to Seconds
func (cvr Converter) MinutesToSeconds(m Minutes) Seconds {
	//conversion code
	return Seconds(float64(m) * 60)
}

//SecondsToMinutes converts Minutes to Seconds
func (cvr Converter) SecondsToMinutes(s Seconds) Minutes {
	//conversion code
	return Minutes(float64(s) * 0.01666)
}

func main() {
	cvr := Converter{}

	fmt.Println((cvr.CentimetersToFeet(13.50)), "feet")
	fmt.Println((cvr.FeetToCentimeters(3)), "centimeters")
	fmt.Println((cvr.MinutesToSeconds(4)), "seconds")
	fmt.Println((cvr.SecondsToMinutes(120)), "minutes")
}
