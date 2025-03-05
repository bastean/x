package release

type Doer interface {
	Do(list ...string) (result string, err error)
}
