package hasher

type Hasher interface {
	Hash(input string) (string, error)
}
