package filestore

import "fmt"

// GetStorageUrl returns a S3 storage URL
func GetStorageURL(path string) string {
	return fmt.Sprintf("s3://%s", path)
}