
//Example of Method
//func (k kelvin) celsius() celsius {  // This is a method. Explanation of method structure: func = keyword,  (k kelvin) = reciever, celsius() = method name, celsius = result 
//	return celsius(k - 273.15)
//}

package main

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
    return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
    return kelvin(c + 273.15)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
    return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
    return f.celsius().kelvin()
}

type kelvin float64

func (k kelvin) celsius() celsius {
    return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
    return k.celsius().fahrenheit()
}

func main() {
    var k kelvin = 294.0
    c := k.celsius()
    fmt.Print(k, "° K is ", c, "° C")
}