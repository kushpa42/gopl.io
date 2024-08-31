package weightconv

func LbToKg(lb Pound) Kilogram { return Kilogram(lb / 2.2) }
func KgToLb(kg Kilogram) Pound { return Pound(kg * 2.2) }
