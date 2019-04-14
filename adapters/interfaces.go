package adapters

// KeyStore reads keys from secrets manager
type KeyStore interface {
	ReadKey(key string) (string, error)
}