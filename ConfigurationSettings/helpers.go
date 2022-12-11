package ConfigurationSettings

type BoolExtended int

const (
	Unknown BoolExtended = iota
	True
	False
)

func (r BoolExtended) Bool() bool {
	switch r {
	case True:
		return true
	case False:
		return false
	case Unknown:
		return false
	}
	return false
}
