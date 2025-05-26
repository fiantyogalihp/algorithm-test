package modules

import (
	"fmt"
	"log"
	"os"
	"testing"

	util "github.com/fiantyogalihp/algorithm-test"
)

// Test01 Print name and currency from each object
func Test01(t *testing.T) {
	fmt.Println("PRINT NAME AND CURRENCY FROM EACH OBJECT")

	jsonBytes, err := os.ReadFile("../data.json")
	if err != nil {
		log.Fatalf("Failed to read data.json: %v", err)
	}

	entries := TestAlgorithm01(jsonBytes)
	util.ExecutableMatchesLogic(t, entries, util.Test01)
}

// Test02 find total data from each object
func Test02(t *testing.T) {
	fmt.Println("FIND TOTAL DATA FROM EACH OBJECT")

	jsonBytes, err := os.ReadFile("../data.json")
	if err != nil {
		log.Fatalf("Failed to read data.json: %v", err)
	}

	entries := TestAlgorithm02(jsonBytes)
	util.ExecutableMatchesLogic(t, entries, util.Test02)
}

// Test03 find all the regions and print Total Regions
func Test03(t *testing.T) {
	fmt.Println("FIND ALL THE REGIONS AND PRINT IT")

	jsonBytes, err := os.ReadFile("../data.json")
	if err != nil {
		log.Fatalf("Failed to read data.json: %v", err)
	}

	entries := TestAlgorithm03(jsonBytes)
	util.ExecutableMatchesLogic(t, entries, util.Test03)
}

// Test04 find total currency from each regions
func Test04(t *testing.T) {
	fmt.Println("FIND TOTAL CURRENCY FROM EACH REGIONS")

	jsonBytes, err := os.ReadFile("../data.json")
	if err != nil {
		log.Fatalf("Failed to read data.json: %v", err)
	}

	entries := TestAlgorithm04(jsonBytes)
	util.ExecutableMatchesLogic(t, entries, util.Test04)
}
