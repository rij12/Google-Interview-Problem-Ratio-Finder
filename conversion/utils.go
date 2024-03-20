package conversion

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadConverstionFromTextFile(filepath string) []string {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Could not open file: [%s], with error: [%s]", filepath, err)
	}

	defer file.Close()

	nodes := make([]string, 20)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		nodes = append(nodes, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nodes
}

func CreateNodeFromString(str string) (string, string, float64) {
	strArray := strings.Split(str, ",")
	raito, err := strconv.ParseFloat(strings.TrimSpace(strArray[2]), 64)
	if err != nil {
		log.Fatal(err)
	}
	return strArray[0], strings.TrimSpace(strArray[1]), raito
}

// Simple Util to convert map keys to an slice
func ConvertMaptokeysToList(m map[string]float64) []string {
	keyArray := make([]string, 0)
	for key := range m {
		keyArray = append(keyArray, key)
	}
	return keyArray
}

// Simple Util to convert map keys to an slice
func ConvertMapOfMapstokeysToList(m map[string]map[string]float64) []string {
	keyArray := make([]string, 0)
	for k := range m {
		keyArray = append(keyArray, k)
	}
	return keyArray
}
