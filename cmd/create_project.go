package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fanchann/go-starter/pkg/gen"
	"github.com/fanchann/go-starter/pkg/utils"
	"github.com/fanchann/go-starter/templates"
)

type SParam struct {
	PackageName string
	DBOpts      string
}

func CreateProject(p SParam) {
	for _, p := range projectDirectoryLists {
		if err := createAppPath(p); err != nil {
			panic(err)
		}
	}

	gover := utils.GetGoVersion()
	config := tmplStr("config", p.DBOpts)
	compose := tmplStr("compose", p.DBOpts)
	main := tmplStr("main", p.DBOpts)
	deps := tmplStr("deps", p.DBOpts)
	appConfig := tmplStr("app_config", p.DBOpts)

	starter := []gen.AppStructure{
		{TmplFile: config, OutputFile: fmt.Sprintf("./internals/config/%s.go", p.DBOpts)},
		{TmplFile: compose, OutputFile: "./compose.yaml"},
		{TmplFile: main, OutputFile: "./cmd/main.go"},
		{TmplFile: deps, OutputFile: "./go.mod"},
		{TmplFile: "viper.tmpl", OutputFile: "./internals/config/viper.go"},
		{TmplFile: "helpers.tmpl", OutputFile: "./internals/helpers/helpers.go"},
		{TmplFile: appConfig, OutputFile: "./config.dev.yaml"},
	}

	for _, s := range starter {
		s.GenerateFromTemplate(templates.TemplateEmbed, gen.StarterParam{GoVersion: gover, PackageName: p.PackageName})
	}
}

func tmplStr(tmpl, name string) string {
	return fmt.Sprintf("%s_%s.tmpl", tmpl, name)
}

func createAppPath(path string) error {
	absPath := filepath.Join("./", path)
	return os.MkdirAll(absPath, os.ModePerm)
}
