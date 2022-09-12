package inp

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine(greeting string) string {
	fmt.Print(greeting)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	return string(text)
}
