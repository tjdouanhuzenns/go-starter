package utils

var (
	DB = []string{"MongoDB", "MYSQL", "PostgreSQL"}
)

func DBSelected(s string) string {
	switch s {
	case DB[0]:
		return "mongodb"
	case DB[1]:
		return "mysql"
	case DB[2]:
		return "postgres"
	default:
		return "mysql"
	}
}
