package globals

type ISecret interface {
	IsHidden() bool
	SetIsHidden(bool)
}
