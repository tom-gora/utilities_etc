package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// helpers
func hexToDec(hex string) int {
	decimal, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		return 0
	}
	return int(decimal)
}

func hexValueSanitize(color string) string {
	if len(color) == 4 {
		return fmt.Sprintf("%c%c%c%c%c%c", color[1], color[1], color[2], color[2], color[3], color[3])
	}
	return color[1:]
}

func hexBrightness(hex string) float64 {
	hex = hexValueSanitize(hex)
	r := hexToDec(hex[:2])
	g := hexToDec(hex[2:4])
	b := hexToDec(hex[4:])
	return 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
}

// sorter
func sortColorsByBrightness(colors []string) []string {
	sort.SliceStable(colors, func(i, j int) bool {
		return hexBrightness(colors[i]) > hexBrightness(colors[j])
	})
	return colors
}

// color strings extractor
func extractHexColors(input string) []string {
	hexColorRegex := `#(?:[0-9a-fA-F]{3}){1,2}`

	re := regexp.MustCompile(hexColorRegex)

	matches := re.FindAllString(input, -1)

	return matches
}

// dedup util
func removeDuplicates(colors []string) []string {
	uniqueColors := make(map[string]struct{})
	for _, color := range colors {
		uniqueColors[color] = struct{}{}
	}

	var result []string
	for color := range uniqueColors {
		result = append(result, color)
	}

	return result
}

// save
func writeColorsToFile(colors []string, outputFilePath string) error {
	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, color := range colors {
		_, err := file.WriteString(color + ";\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Extract and sort by brightness hex color strings from svg, html, eps, md... and any other text files.\nUsage: color-extractor <input_file>")
		return
	}

	inputFilePath := os.Args[1]

	inputData, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	colors := extractHexColors(string(inputData))

	uniqueColors := removeDuplicates(colors)

	sortedColors := sortColorsByBrightness(uniqueColors)

	fileName := strings.TrimSuffix(inputFilePath, path.Ext(inputFilePath))

	outputFilePath := fmt.Sprintf("colors_%s.txt", fileName)

	err = writeColorsToFile(sortedColors, outputFilePath)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	fmt.Println("Colors written to", outputFilePath)
}
