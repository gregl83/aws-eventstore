package database

import "fmt"

// GetConnectionUrl returns a MySQL formatted connection URL
func GetConnectionURL(username, password, host, database, port string) string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
}
