package evalogger

type LoggerType int

const (
	REDIS LoggerType = iota
	STD
	JSON
)

type ILogger interface {
	PutSuccessMessage(id string, result bool, message string)
	Printf(id string, statementId string, format string, args ...interface{})
	Errorf(id string, statementId string, format string, args ...interface{})
	GetType() LoggerType
	Init(id string)
	IsOnError() bool
	Clear()
}
