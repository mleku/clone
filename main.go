package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"lol.mleku.dev/chk"
)

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(
			os.Stderr,
			`usage: clone <go-style import url>

clones using SSH in standard style for github repos, eg:

    clone github.com/user/repo

will clone using git@github.com:user/repo.git' and place it in $GOPATH/src/github.com/user/repo

* requires git to be installed
`,
		)
		os.Exit(1)
	}
	split := strings.Split(os.Args[1], "/")
	path := fmt.Sprintf("git@%s:%s/%s", split[0], split[1], split[2])
	outPath := fmt.Sprintf(
		"%s/src/%s", os.Getenv("HOME"), os.Args[1],
	)
	command := []string{"git", "clone", path, outPath}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); chk.E(err) {
		os.Exit(1)
	}
}
