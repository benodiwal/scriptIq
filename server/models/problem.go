package models

type TestCase struct {
	Numbers []int `json:"numbers"`
	Target int `json:"target"`
	ExpectedResult [2]int `json:"expected_result"`
}

type Problem struct {
	Description string `json:"description"`
	TestCases []TestCase `json:"test_cases"`
}