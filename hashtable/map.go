package hashtable

type Map interface {
	Set(key, value any) error
	Get(key any) *Pair
	Remove()
}
