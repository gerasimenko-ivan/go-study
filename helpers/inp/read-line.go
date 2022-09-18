package inp

import (
	"bufio"
	"fmt"
	"os"
)

// ReadLine returns string entered by user
func ReadLine(greeting string) string {
	fmt.Print(greeting)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	return string(text)
}
