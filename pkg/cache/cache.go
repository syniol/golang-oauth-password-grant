package cache

type Cache interface {
	Persist()
	LookUp()
}
