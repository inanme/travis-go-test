package version

import (
	"fmt"
	"github.com/thoas/go-funk"
	"os"
)

//nolint:gochecknoglobals
var (
	version  string
	commit   string
	date     string
	builtBy  string
	snapshot string
)

func PrintVersion() {
	if snapshot == "true" {
		fmt.Println("This is a SNAPSHOT build")
	}
	if funk.ContainsString(os.Args, "simple") {
		fmt.Print(version) //nolint
	} else {
		fmt.Printf("version: %s\n", version) //nolint
		fmt.Printf("commit : %s\n", commit)  //nolint
		fmt.Printf("date: %s\n", date)       //nolint
		fmt.Printf("builtBy: %s\n", builtBy) //nolint
	}
}
