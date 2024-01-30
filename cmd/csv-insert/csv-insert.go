package main

import "github.com/sunglim/chart-server/pkg/cmd/csvinsert"

func main() {
	cmd := csvinsert.NewCommand()
	cmd.Execute()
}
