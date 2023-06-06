package entity

import "time"

type TestSheet struct {
	ID          string     // Unique identifier for the test sheet
	Name        string     // Name or title of the test sheet
	Description string     // Description or summary of the test sheet
	ProjectID   string     // Identifier of the project the test sheet belongs to
	ModuleID    string     // Identifier of the module the test sheet belongs to
	CreatedBy   string     // Name or identifier of the user who created the test sheet
	CreatedAt   time.Time  // Timestamp indicating when the test sheet was created
	UpdatedAt   time.Time  // Timestamp indicating the last update of the test sheet
	TestCases   []TestCase // List of test cases associated with the test sheet
}
