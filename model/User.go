// package declaration
package model

// type definitions
type UserRecord struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	Day      int    `json:"day"`
}

// data
var (
	UserRecords []UserRecord
)
