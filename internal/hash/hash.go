package hash

type PasswordHash interface {
	ComparePasswordAndHash(password string, encodeHash string) (match bool, err error)
	GenerateFromPassword(password string) (encodedHash string, err error)
}
