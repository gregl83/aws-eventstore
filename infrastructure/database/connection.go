package database

import "fmt"

// GetConnectionUrl returns a MySQL formatted connection URL
func GetConnectionURL(host, port, database, username, password string) string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
}
