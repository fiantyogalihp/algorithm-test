package algorithm_test

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type DataTest struct {
	Name     string `json:"name"`
	Currency int    `json:"currency"`
	Country  string `json:"country"`
	Region   string `json:"region"`
}

const (
	Test01 string = "test01"
	Test02 string = "test02"
	Test03 string = "test03"
	Test04 string = "test04"
)

// @param "jsonData" only for result from each test
// @param "testName" only for from const at utils.go
func generateAnswerFromJSON(jsonData []DataTest, testName string) string {
	var builder strings.Builder

	if testName == Test02 {
		builder.WriteString(fmt.Sprintf("Total Data: %d", len(jsonData)))
		return builder.String()
	}

	// Optional: Sort entries by Name for consistent comparison
	sort.Slice(jsonData, func(i, j int) bool {
		switch testName {
		case Test01:
			return jsonData[i].Name < jsonData[j].Name

		case Test03, Test04:
			return jsonData[i].Region < jsonData[j].Region

		default:
			return jsonData[i].Name < jsonData[j].Name
		}
	})

	for _, entry := range jsonData {

		switch testName {
		case Test01:
			builder.WriteString(entry.Name)
			builder.WriteString(" : ")
			builder.WriteString(strings.TrimSpace(fmt.Sprintf("%d", entry.Currency)))
			builder.WriteString("\n")

		case Test03:
			builder.WriteString(entry.Region)
			builder.WriteString("\n")

		case Test04:
			builder.WriteString(fmt.Sprintf("%s %d\n", entry.Region, entry.Currency))
		}

	}

	// additional string build
	if testName == Test03 {
		builder.WriteString(fmt.Sprintf("Total Region: %d", len(jsonData)))
	}

	return builder.String()
}

func normalizeString(s string) string {
	// Remove trailing and leading whitespace from each line
	lines := strings.Split(strings.TrimSpace(s), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	// Join with a consistent newline
	return strings.Join(lines, "\n")
}

func ExecutableMatchesLogic(t *testing.T, entries []DataTest, testName string) {
	cmd := exec.Command("./answer_exec", "-"+testName)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run answer executable: %v", err)
	}
	executableOutput := strings.TrimSpace(out.String())

	// Step 3: Generate expected output from JSON
	expectedOutput := strings.TrimSpace(generateAnswerFromJSON(entries, testName))

	// fmt.Println(expectedOutput)

	// Step 4: Compare outputs
	require.Equal(t, normalizeString(expectedOutput), normalizeString(executableOutput), "Mismatch between executable output and expected logic")
}
