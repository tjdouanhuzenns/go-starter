package gen

import (
	"embed"
	"html/template"
	"os"

	"github.com/fanchann/go-starter/pkg/utils"
)

type AppStructure struct {
	TmplFile   string
	OutputFile string
}

type StarterParam struct {
	PackageName string
	GoVersion   string
}

func (d AppStructure) GenerateFromTemplate(tmpl embed.FS, s StarterParam) {
	tmplByte, err := tmpl.ReadFile(d.TmplFile)
	utils.ErrorWithLog(err)

	file, err := os.Create(d.OutputFile)
	utils.ErrorWithLog(err)
	defer file.Close()

	// Memparsing template dari file sistem yang di-embed
	tmplParsed, err := template.New("").Parse(string(tmplByte))
	utils.ErrorWithLog(err)

	utils.ErrorWithLog(tmplParsed.Execute(file, s))
}
