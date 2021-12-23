package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		}
		if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(a, b []string) []string {
	return append(b, a[len(a)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(a []float64, b int) []float64 {
	trash := make([]float64, len(a))
	for i, v := range a {
		trash[i] = v * float64(b) / 2
	}
	return trash
}
