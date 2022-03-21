/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"cloudfairy/cmd"
	"cloudfairy/templates"
)

func main() {
	templates.Load()
	cmd.Execute()
}
