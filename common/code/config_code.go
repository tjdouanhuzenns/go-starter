package code

var (
	postgresYamlConfig = map[string]interface{}{
		"database": map[string]interface{}{
			"username":  "postgres",
			"password":  "postgres",
			"host":      "localhost",
			"port":      5432,
			"name":      "test",
			"sslmode":   "disable",
			"time_zone": "UTC",
		},
	}

	mysqlYamlConfig = map[string]interface{}{
		"database": map[string]interface{}{
			"username": "root",
			"password": "root",
			"host":     "localhost",
			"port":     3306,
			"name":     "information_schema",
			"pool": map[string]interface{}{
				"idle":     10,
				"max":      25,
				"lifetime": 300,
			},
		},
	}

	mongoYamlConfig = map[string]interface{}{
		"database": map[string]interface{}{
			"url":  "mongodb://localhost:27017",
			"name": "test",
		},
	}
)

// ConfigurationFileGenerate returns a default YAML config map for the given database driver.
// Supported values: "mongodb", "mysql", "postgres". Defaults to mysql if unrecognized.
func ConfigurationFileGenerate(using string) map[string]interface{} {
	switch using {
	case "mongodb":
		return mongoYamlConfig
	case "mysql":
		return mysqlYamlConfig
	case "postgres":
		return postgresYamlConfig
	default:
		return mysqlYamlConfig
	}
}
