package cmd

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"

	"github.com/fanchann/go-starter/pkg/utils"
)

func StarterRootRun() {
	packageName := promptui.Prompt{
		Label: "Insert your package name",
		Validate: func(s string) error {
			if len(s) <= 3 {
				return errors.New("package name must > 3 character!")
			}
			return nil
		},
	}

	databaseList := promptui.Select{
		Label: "Select Database",
		Items: utils.DB,
	}

	pkg, err := packageName.Run()
	utils.ErrorWithLog(err)

	_, databaseSelect, err := databaseList.Run()
	utils.ErrorWithLog(err)

	CreateProject(SParam{PackageName: pkg, DBOpts: utils.DBSelected(databaseSelect)})

	log.Println("Success generate!")
}
