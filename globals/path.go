package globals

type PathElemType int

const (
	VARIABLE_DECL PathElemType = iota
	PROPERTY
	INDEX
)

type PathElem struct {
	Value string
	Index int
	Type  PathElemType
}

func (p PathElem) Copy() PathElem {
	return PathElem{Value: p.Value, Index: p.Index, Type: p.Type}
}
