package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (lb Pound) String() string    { return fmt.Sprintf("%g lbs", lb) }
func (kg Kilogram) String() string { return fmt.Sprintf("%g kg", kg) }
