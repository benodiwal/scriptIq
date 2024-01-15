package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/t01gyl0p/scriptIq/models"
)

func EvaluateCode(c *gin.Context) {

	var request struct {
		Code string `json:"code" binding:"required"`
		Problem  models.Problem `json:"problem" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H {"error": err.Error()})
		return
	}

	saveDir := "/home/user/test_go/"
	filePath := filepath.Join(saveDir, "main.go")
	fmt.Println(filePath)

	err := saveCodeToFile(request.Code, filePath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save code to file"})
		return
	}
	defer os.Remove(filePath)

	// Complile the code
	cmd := exec.Command("go", "build", filePath)
	err = cmd.Run()
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to compile code"})
		return
	}

	// Run the compile code
	result := runCompiledCode(filePath, request.Problem)
	c.JSON(http.StatusOK, result)
}

func saveCodeToFile(code, filePath string) error {
	return os.WriteFile(filePath, []byte(code), 0644)
}

func runCompiledCode(filePath string, problem models.Problem) map[string]interface{} {
	result := make(map[string]interface{})
	trimmedFilePath := strings.TrimSuffix(filePath, ".go")
	i := 0

	for _, testCase := range problem.TestCases {
		cmd := exec.Command(trimmedFilePath)
		
		// Pass the input to the stdin of the process
		stdin, err := cmd.StdinPipe()
		if err != nil {
			result["error"] = "Failed to create pipe for stdin"
			return result
		}

		stdin.Write([]byte(intSliceToString(testCase.Numbers) + "\n" + strconv.Itoa(testCase.Target) + "\n"))
		stdin.Close()

		// Capture the output and error
		output, err := cmd.CombinedOutput()
		if err != nil {
			result["error"] = "Failed to execute code"
			return result
		}

		// Parse the output and compare with exppected result
		actualResult, err := parseOutput(string(output))
		if err != nil {
			result["error"] = "Failed to parse output"
			return result
		}

		if reflect.DeepEqual(actualResult, testCase.ExpectedResult) {
			result["test_case_" + strconv.Itoa(i+1)] = "True"
		} else {
			result["test_case_" + strconv.Itoa(i+1)] = "Fail"
		}

		i++
	}

	return result
}

func parseOutput(output string) (result [2]int, err error) {
	parts := strings.Fields(output)
	
	if len(parts) != 2 {
		return result, errors.New("Output format is invalid")
	}

	result[0], err = strconv.Atoi(parts[0])
	if err != nil {
		return result, errors.New("Failed to parse first integer")
	}

	result[1], err = strconv.Atoi(parts[1])
	if err != nil {
		return result, errors.New("Failed to parse the second integer")
	}

	return result, nil
}

func intSliceToString(nums []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nums)), " "), "[]")
}