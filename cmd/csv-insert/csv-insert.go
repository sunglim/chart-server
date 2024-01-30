package main

import "github.com/sunglim/chart-server/pkg/cmd/svcinsert"

func main() {
	cmd := svcinsert.NewCommand()
	cmd.Execute()
}
