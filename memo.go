package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// Memo is for storing memo data.
// Is is very simple struct that only have content and created date.
type Memo struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// PrintMemos prints out of passed memo list to console with colored text.
func PrintMemos(memos []Memo) {
	if len(memos) > 0 {
		green := color.New(color.FgGreen).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		for i, memo := range memos {
			fmt.Printf("\n%s", memo.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Printf("\n[%s] %s\n", green(i+1), bold(memo.Content))
		}
		fmt.Println()
	} else {
		fmt.Println("There is no memo")
	}
}
