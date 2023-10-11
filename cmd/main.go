package main

import (
	"os"

	"github.com/taylormonacelli/bluemovement"
)

func main() {
	code := bluemovement.Main()
	os.Exit(code)
}
