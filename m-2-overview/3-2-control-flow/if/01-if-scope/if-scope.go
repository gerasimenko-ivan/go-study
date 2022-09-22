package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)
	mmddyyyyFmt := "01-02-2006"
	if today.Before(tomorrow) {
		fmt.Printf("\nToday '%v' is before tomorrow '%v'\n", today.Format(mmddyyyyFmt), tomorrow.Format(mmddyyyyFmt))
	}
	fmt.Printf("Variable 'tomorrow' is seen here.\n\n")

	if yesterday := today.Add(-24 * time.Hour); yesterday.Before(today) {
		fmt.Printf("yesterday '%v' is before today '%v'\n", yesterday.Format(mmddyyyyFmt), today.Format(mmddyyyyFmt))
	}
	fmt.Printf("Variable 'yesterday' is NOT seen here (ONLY INSIDE IF).\n")
}
