package services

import (
	"math/rand"
	"strconv"
)

func RandomWeather() string {
	return strconv.Itoa(rand.Intn(30))
}
