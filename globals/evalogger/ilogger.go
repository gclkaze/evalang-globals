package evalogger

type LoggerType int

const (
	REDIS LoggerType = iota
	STD
	JSON
	STRING
	COMPACT
)

type ILogger interface {
	PutSuccessMessage(id string, result bool, message string)
	Printf(id string, statementId string, format string, args ...interface{})
	Errorf(id string, statementId string, format string, args ...interface{})
	PrintErrorMessage(id string, message string)
	GetType() LoggerType
	Init(id string)
	IsOnError() bool
	Clear()
}
