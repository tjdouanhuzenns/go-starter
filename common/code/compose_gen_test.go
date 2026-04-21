package code_test

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/fanchann/go-starter/common/code"
)

var (
	genFolder = "./gen/"
)

func TestGenMysqlCompose(t *testing.T) {
	m := code.ComposeCodeGenerate("mysql")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	// ensure gen directory exists before writing
	_ = os.MkdirAll(genFolder, fs.ModePerm)

	err2 := os.WriteFile(genFolder+"mysql.yaml", out, fs.ModePerm)
	assert.Nil(t, err2)
}

func TestGenPostgresCompose(t *testing.T) {
	m := code.ComposeCodeGenerate("postgres")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	// ensure gen directory exists before writing
	_ = os.MkdirAll(genFolder, fs.ModePerm)

	err2 := os.WriteFile(genFolder+"postgres.yaml", out, fs.ModePerm)
	assert.Nil(t, err2)
}

func TestGenMongoCompose(t *testing.T) {
	m := code.ComposeCodeGenerate("mongodb")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	// ensure gen directory exists before writing
	_ = os.MkdirAll(genFolder, fs.ModePerm)

	err2 := os.WriteFile(genFolder+"mongodb.yaml", out, fs.ModePerm)
	assert.Nil(t, err2)
}

// TestGenRedisCompose verifies that a Redis compose file can be generated.
// Added for personal use since I commonly use Redis for caching in my projects.
func TestGenRedisCompose(t *testing.T) {
	m := code.ComposeCodeGenerate("redis")
	out, err := yaml.Marshal(m)
	assert.Nil(t, err)

	// ensure gen directory exists before writing
	_ = os.MkdirAll(genFolder, fs.ModePerm)

	err2 := os.WriteFile(genFolder+"redis.yaml", out, fs.ModePerm)
	assert.Nil(t, err2)
}
