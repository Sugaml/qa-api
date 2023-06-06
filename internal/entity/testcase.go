package entity

type TestCase struct {
	ID             string    // Unique identifier for the test case
	Name           string    // Name or title of the test case
	Description    string    // Description or summary of the test case
	Steps          []string  // List of steps to execute for the test case
	ExpectedResult string    // Expected result of the test case
	ActualResult   string    // Actual result obtained during test execution
	Status         string    // Current status of the test case (e.g., pass, fail, blocked)
	Comments       []Comment // List of comments or notes related to the test case
}
