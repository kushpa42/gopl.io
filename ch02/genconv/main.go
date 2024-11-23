// genconv is a generic converter
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/lenconv"
	"gopl.io/ch2/tempconv"
	"gopl.io/ch2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))

		fmt.Printf("%s = %s\n", f, tempconv.FToK(f))
		fmt.Printf("%s = %s\n", k, tempconv.KToF(k))

		fmt.Printf("%s = %s\n", c, tempconv.CToK(c))
		fmt.Printf("%s = %s\n", k, tempconv.KToC(k))

		lb := weightconv.Pound(t)
		kg := weightconv.Kilogram(t)

		fmt.Printf("%s = %s\n", lb, weightconv.LbToKg(lb))
		fmt.Printf("%s = %s\n", kg, weightconv.KgToLb(kg))

		ft := lenconv.Feet(t)
		m := lenconv.Meter(t)

		fmt.Printf("%s = %s\n", ft, lenconv.FtToM(ft))
		fmt.Printf("%s = %s\n", m, lenconv.MToFt(m))
	}
}
