package main

import "fmt"

type Temp interface {
	conv2Cel() float32
	conv2Kel() float32
	conv2Fah() float32
}

type Celsius struct {
	degree float32
	unit   string
}

type Kelvin struct {
	degree float32
	unit   string
}

type Fahrenheit struct {
	degree float32
	unit   string
}

func (k Kelvin) conv2Cel() float32 {
	return k.degree - 273.15
}
func (f Fahrenheit) conv2Cel() float32 {
	return (f.degree - 32) * 5 / 9
}

func (k Kelvin) conv2Fah() float32 {
	return (k.degree-273.15)*9/5 + 32
}

func (c Celsius) conv2Fah() float32 {
	return (c.degree * 9 / 5) + 32
}

func (f Fahrenheit) conv2Kel() float32 {
	return (f.degree-32)*5/9 + 273.15
}

func (c Celsius) conv2Kel() float32 {
	return c.degree + 273.15
}

func main() {

	fmt.Println("1 for Celsius, 2 for Kelvin, 3 for Fahreheit")
	var unit int
	fmt.Scanln(&unit)
	fmt.Println("type the degree")
	var deg float32
	fmt.Scanln(&deg)
	fmt.Println("which type do you want to convert? 1 for Celsius, 2 for Kelvin, 3 for Fahreheit")
	var conv int
	fmt.Scanln(&conv)
	switch unit {
	case 1:
		c := Celsius{degree: deg, unit: "°C"}
		if conv == 2 {
			fmt.Println(c.conv2Kel())
		} else if conv == 3 {
			fmt.Println(c.conv2Fah())
		} else {
			fmt.Println("can be only converted to Fahrenheit or Kelvin")
		}
	case 2:
		k := Kelvin{degree: deg, unit: "K"}
		if conv == 1 {
			fmt.Println(k.conv2Cel())
		} else if conv == 3 {
			fmt.Println(k.conv2Fah())
		} else {
			fmt.Println("can be only converted to Fahrenheit or Celsius")
		}
	case 3:
		f := Fahrenheit{degree: deg, unit: "°F"}
		if conv == 1 {
			fmt.Println(f.conv2Cel())
		} else if conv == 2 {
			fmt.Println(f.conv2Kel())
		} else {
			fmt.Println("can be only converted to Celsius or Kelvin")
		}
	}

}
