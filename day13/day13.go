package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

func solve(input string) {
	layers := parseLayers(input)

	severity := 0
	for _, layer := range layers {
		if layer.depth%(2*layer.drift-2) == 0 {
			severity += layer.depth * layer.drift
		}
	}

	delay := 0
	for ; isFirewalled(layers, delay); delay++ {
	}

	fmt.Printf("Severity: %d, Delay: %d\n", severity, delay)
}

type layer struct {
	depth int
	drift int
}

func parseLayers(input string) []layer {
	layers := []layer{}
	for _, layerString := range strings.Split(input, "\n") {
		layerStringSplit := strings.Split(layerString, ": ")
		depth, errDepth := strconv.Atoi(layerStringSplit[0])
		drift, errDrift := strconv.Atoi(layerStringSplit[1])
		_, _ = errDepth, errDrift
		layers = append(layers, layer{depth, drift})
	}
	return layers
}

func isFirewalled(layers []layer, delay int) bool {
	for _, layer := range layers {
		if (layer.depth+delay)%(2*layer.drift-2) == 0 {
			return true
		}
	}
	return false
}

//go:embed input.txt
var myInput string

func main() {
	solve(myInput)
}
