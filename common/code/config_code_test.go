package code_test

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/fanchann/go-starter/common/code"
)

// use 0644 instead of fs.ModePerm|fs.ModeAppend to avoid overly permissive file permissions
const configFileMode = fs.FileMode(0644)

func TestGenConfigMysql(t *testing.T) {
	m := code.ConfigurationFileGenerate("mysql")

	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := os.WriteFile(genFolder+"config.mysql.yaml", out, configFileMode)
	assert.Nil(t, err2)
}

func TestGenConfigMongoDB(t *testing.T) {

	m := code.ConfigurationFileGenerate("mongodb")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := os.WriteFile(genFolder+"config.mongodb.yaml", out, configFileMode)
	assert.Nil(t, err2)
}

func TestGenConfigPostgres(t *testing.T) {
	m := code.ConfigurationFileGenerate("postgres")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	err2 := os.WriteFile(genFolder+"config.postgres.yaml", out, configFileMode)
	assert.Nil(t, err2)
}
