package main

import "github.com/fanchann/go-starter/cmd"

func main() {
	cmd.CreateProject(cmd.SParam{PackageName: "go-starter/3.5", DBOpts: "mysql"})
}
