package main

import (
	"flag"
	"fmt"

	"github.com/fxsjy/gonn/gonn"
)

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Система помощи принятия решений для персонажа. Статья с Хабр***************")
	fmt.Println(string(colorReset), "")

	hp := flag.Float64("hp", 0, "Здоровье персонажа в диапозоне: 0.1 - 1.0")
	weaponBool := flag.Bool("weapon", false, "Есть ли оружие у персонажа")
	enemyCount := flag.Float64("enemies", 0, "Колличество противников")
	flag.Parse()
	//CreateNN()

	nn := gonn.LoadNN("gonn")
	weapon := parseBoolToFloat64(*weaponBool)

	out := nn.Forward([]float64{*hp, weapon, *enemyCount})
	fmt.Printf("Здоровье: %.1f, Оружие: %t, Кол-во противников: %.0f\n", *hp, parseFloat64ToBool(weapon), *enemyCount)
	fmt.Printf("Действие: %s\n", GetResult(out))
}
func parseFloat64ToBool(f float64) bool {
	return f == 1.0
}

func parseBoolToFloat64(b bool) float64 {
	if b {
		return 1.0
	}
	return 0
}

func CreateNN() {
	nn := gonn.DefaultNetwork(3, 16, 4, false)

	input := [][]float64{
		[]float64{0.5, 1, 1}, []float64{0.9, 1, 2}, []float64{0.8, 0, 1},
		[]float64{0.3, 1, 1}, []float64{0.6, 1, 2}, []float64{0.4, 0, 1},
		[]float64{0.9, 1, 7}, []float64{0.6, 1, 4}, []float64{0.4, 0, 1},
		[]float64{0.6, 1, 0}, []float64{1, 0, 0},
	}
	target := [][]float64{
		[]float64{1, 0, 0, 0}, []float64{1, 0, 0, 0}, []float64{1, 0, 0, 0},
		[]float64{0, 1, 0, 0}, []float64{0, 1, 0, 0}, []float64{0, 1, 0, 0},
		[]float64{0, 0, 1, 0}, []float64{0, 0, 1, 0}, []float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1}, []float64{0, 0, 0, 1},
	}

	nn.Train(input, target, 100000)

	gonn.DumpNN("gonn", nn)
}

func GetResult(output []float64) string {
	max := float64(-99999)
	pos := -1

	for i, value := range output {
		if value > max {
			max = value
			pos = i
		}
	}

	switch pos {
	case 0:
		return "Атаковать"
	case 1:
		return "Красться"
	case 2:
		return "Убегать"
	case 3:
		return "Ничего не делать"
	}

	return ""
}
