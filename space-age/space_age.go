package space

// Planet is the planet name without abreviation
type Planet string

const earthYearSeconds = 31557600

// Age ...
func Age(seconds float64, planet Planet) float64 {

	if planet == "Earth" {
		return seconds / earthYearSeconds
	}

	planets := map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return seconds / (earthYearSeconds * planets[string(planet)])
}

/*

BEST SOLUTION
====================================================

type Planet string

const earthYearSeconds = 31557600

func (p Planet) earthYear() float64 {
	switch p {
	case "Earth":
		return earthYearSeconds
	case "Mercury":
		return earthYearSeconds * 0.2408467
	case "Venus":
		return earthYearSeconds * 0.61519726
	case "Mars":
		return earthYearSeconds * 1.8808158
	case "Jupiter":
		return earthYearSeconds * 11.862615
	case "Saturn":
		return earthYearSeconds * 29.447498
	case "Uranus":
		return earthYearSeconds * 84.016846
	case "Neptune":
		return earthYearSeconds * 164.79132
	}
	return 0
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / planet.earthYear()
}
*/
