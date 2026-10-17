package globals

import (
	"fmt"
)

type StringContentType int

const (
	PLAIN_TEXT StringContentType = iota
	HTML
)

type StatementOp int

const (
	NOP StatementOp = iota
	BrowserOpen
	BrowserRefresh
	BrowserClose
	BrowserType
	BrowserClickElement
	BrowserReadElement
	BrowserFindElement
	BrowserSmartFindElement
	BrowserScreenshot
	BrowserSavePage
	BrowserScroll
	BrowserScrollElement
	BrowserMasterTabOpen
	BrowserMasterClose
	BrowserTestOperation
	Wait
	JsonValue
	AssertEquals
	AssertNotEquals
	AssertEmpty
	AssertNotEmpty
	AssignValue
	Print
	Warn
	Err
	Regex
	FindElements
	CheckHTMLElementAttribute
	GetHTMLElementAttribute
	UploadFile
	OpenTab
	SetCurrentTab
	DbRead
	DbWrite
	DbConnect
	DbPing
	RESTClientCreate
	RESTGET
	MetricsIndexSave
	Post
	DeleteFile
	FileRead
	FileWrite
	FileExists
	OpenFile
	TouchFile
	CompareFile
	JsonContains
	JsonNotContains
	Date
	CurrentUnknownOperation
	WS_Create
	WS_RunCommand
	WS_Stop
	WS_WriteFile
	WS_ArchiveFile
	WS_StartScript
	WS_KillScript
	JNK_Connect
	JNK_Execute
	VLT_Connect
	VLT_Get
	SLK_Connect
	SLK_SendMessage
	DIS_Connect
	DIS_SendMessage
	DIS_Stop
	EXT_Load
	EXT_Execute
	EXT_FreeStringFunctionRegister
	EXT_Unload
	TM_Now
	ENV_Get
	ENV_Set
	ENV_Unset
	ENV_All
	ENV_Exists
	ENV_Os
	ENV_Hostname
	ENV_Username
	ENV_Cwd
	ENV_Tempdir
	ENV_Homedir
	ENV_Executioncontext
	JSN_Parse
	JSN_Set
	JSN_Remove
	JSN_ArrayAppend
	JSN_ArrayRemove
	JSN_Exists
	JSN_GetOrDefault
	JSN_ArrayGet
	JSN_ArrayGetOrDefault
	JSN_ArrayContains
	JSN_Merge
	JSN_ArrayParse
	JSN_TestParse
	JSN_ArrayTestParse
	JSN_Keys
	JSN_FilterAttributes
	JSN_ArrayFilterAttributes
	JSN_Map
	JSN_ArrayMap
	JSN_ArrayIsEmpty
	JSN_ArrayPluck
	JSN_ArrayChunk
	JSN_ArraySlice
	JSN_ArrayInsert
	RND_GetIntInRange
	RND_GetDoubleInRange
	RND_GetString
	RND_GetBool
	RND_GetPick
	RND_GetShuffled
	RND_GetUuid
	RND_Seed
	MTH_Round
	MTH_Abs
	MTH_Clamp
	MTH_Min
	MTH_Max
	MTH_Sum
	MTH_Avg
	MTH_Stats
	MTH_Percent
	MTH_Sin
	MTH_Cos
	MTH_Tan
	MTH_Arctan
	MTH_Sqrt
	EML_Send
	EML_SendEach
	EML_Render
	EML_Validate
	EML_Connect
	EML_Close
	RND_GetStringInRange
	JSN_InferSchema
	JSN_MergeSchema
	JSN_Validate
	JSN_DiffSchema
	JSN_Transform
	JSN_ToGo
	JSN_ToTs
	JSN_Fake
	HTTP_Request
	HTTP_Get
	HTTP_GetAll
	HTTP_Each
	HTTP_Post
	HTTP_Put
	HTTP_Patch
	HTTP_Delete
	RESTPOST
	RESTPUT
	RESTPATCH
	RESTDELETE
	TM_NowDateTime
	TM_Today
	TM_ParseDate
	TM_ParseDateTime
	TM_Format
	TM_ToDate
	TM_ToDateTime
	TM_Add
	TM_Diff
	TM_Sleep
	//SLK_ReadMessage
)

type StatementParameterTypeBase int

const (
	STRING StatementParameterTypeBase = iota
	INTEGER
	DOUBLE
	JSON
	JSON_ARRAY
	HTML_ELEM
	DATE
	BOOLEAN
	VARIABLE
	VARIABLE_STORAGE
	DB
	HTTP_CLIENT
	BROWSER
	NUMBER
	NULL
	LIST
	CALCULATED
	EXPRESSION
	REFERENCE
	WORKSPACE
	JENKINS
	ENVIRONMENT
	JSON_PACKAGE
	VAULT
	SLACK
	DISCORD
	EXTERNAL
	USER_DEFINED
	VARIADIC
	WEB_ELEMENT
	WEB_ELEMENT_LIST
	MAILER
	DATETIME
)

type JSONStruct = map[string]interface{}
type JSONGenericArray = []interface{}
type JSONArray = []map[string]interface{}
type JSONArrayGen = []interface{}
type JSONObjectGen = interface{}
type JSONIntArray = []int64
type JSONDoubleArray = []float64
type JSONBoolArray = []bool
type JSONStringArray = []string

type JSONEmptyArray = []interface{}

type JSONArrayType int

const (
	JSON_EMPTY_ARRAY JSONArrayType = iota
	JSON_INT_ARRAY
	JSON_DOUBLE_ARRAY
	JSON_BOOL_ARRAY
	JSON_STRING_ARRAY
	JSON_OBJECT_ARRAY
	JSON_HTML_ARRAY
	JSON_OBJECT
)

type StatementType int

const (
	OPERATION StatementType = iota
	FOR_LOOP
	CONTROL
	RETURN
	CONTINUE
	BREAK
	IMPORT
	NEXT_SCRIPT
	BLOCK
	PROGRAM
	LABEL
	CALL_SCRIPT
	MODULE_DECLARATION
	MODULE_FUNCTION_EXPORT
	MODULE_FUNCTION_PARAMETER
	MODULE_FUNCTION_RESULT
	MODULE_FUNCTION_KEY
)

func GetUserFriendlyType(e string) StatementParameterTypeBase {
	switch e {
	case "expression":
		return EXPRESSION
	case "variable storage":
		return VARIABLE_STORAGE
	case "null":
		return NULL
	case "text":
		return STRING
	case "integer":
		return INTEGER
	case "double":
		return DOUBLE
	case "json":
		return JSON
	case "date":
		return DATE
	case "boolean":
		return BOOLEAN
	case "variable":
		return VARIABLE
	case "database object":
		return DB
	case "number":
		return NUMBER
	case "list":
		return LIST
	case "web element":
		return WEB_ELEMENT
	case "web element list":
		return WEB_ELEMENT_LIST
	case "html":
		return HTML_ELEM
	case "worskpace":
		return WORKSPACE
	case "jenkins":
		return JENKINS
	case "vault":
		return VAULT
	case "slack":
		return SLACK
	case "discord":
		return DISCORD
	case "external":
		return EXTERNAL
	case "user-defined":
		return USER_DEFINED
	case "env":
		return ENVIRONMENT
	case "mailer":
		return MAILER
	case "datetime":
		return DATETIME
	case "json package":
		{
			return JSON_PACKAGE
		}
	default:
		return NULL
	}

}

func (e StatementParameterTypeBase) String() string {
	switch e {
	case EXPRESSION:
		return "Expression"
	case VARIABLE_STORAGE:
		return "Variable Storage"
	case NULL:
		return "Null"
	case STRING:
		return "String"
	case INTEGER:
		return "Integer"
	case DOUBLE:
		return "Double"
	case JSON:
		return "JSON"
	case DATE:
		return "Date"
	case BOOLEAN:
		return "Boolean"
	case VARIABLE:
		return "Variable"
	case DB:
		return "Database Object"
	case NUMBER:
		return "Number"
	case LIST:
		return "List"
	case HTML_ELEM:
		return "HTML Element"
	case WEB_ELEMENT:
		return "Web Element"
	case WEB_ELEMENT_LIST:
		return "Web Element List"
	case WORKSPACE:
		return "Worskpace"
	case JENKINS:
		return "Jenkins"
	case VAULT:
		return "Vault"
	case SLACK:
		return "Slack"
	case DISCORD:
		return "Discord"
	case EXTERNAL:
		return "External"
	case USER_DEFINED:
		return "UserDefined"
	case ENVIRONMENT:
		return "Environment"
	case JSON_PACKAGE:
		return "Json"
	case MAILER:
		return "Mailer"
	case DATETIME:
		return "DateTime"
	case HTTP_CLIENT:
		return "HTTP Client"
	default:
		return "Unknown"
	}
}

func (e StatementOp) String() string {
	switch e {
	case NOP:
		return "nop"
	case MetricsIndexSave:
		return "metricsIndexSave"
	case RESTClientCreate:
		return "RESTClientCreate"
	case CheckHTMLElementAttribute:
		return "checkHTMLElementAttribute"
	case GetHTMLElementAttribute:
		return "getHTMLElementAttribute"
	case BrowserOpen:
		return "browserOpen"
	case BrowserRefresh:
		return "browserRefresh"
	case Wait:
		return "wait"
	case FileRead:
		return "fileRead"
	case FileWrite:
		return "fileWrite"
	case FileExists:
		return "fileExists"
	case BrowserType:
		return "type"
	case BrowserClose:
		return "browserClose"
	case BrowserScroll:
		return "browserScroll"
	case BrowserScrollElement:
		return "browserScrollElement"
	case BrowserSavePage:
		return "browserSavePage"
	case JsonValue:
		return "jsonValue"
	case BrowserScreenshot:
		return "browserScreenshot"
	case BrowserFindElement:
		return "findElement"
	case BrowserClickElement:
		return "clickElement"
	case BrowserTestOperation:
		return "browserTestOperation"
	case BrowserSmartFindElement:
		return "smartFindElement"
	case AssertEquals:
		return "assertEquals"
	case AssertNotEquals:
		return "assertNotEquals"
	case AssertEmpty:
		return "assertEmpty"
	case AssertNotEmpty:
		return "assertNotEmpty"
	case FindElements:
		return "findElements"
	case AssignValue:
		return "assignValue"
	case Print:
		return "print"
	case Warn:
		return "warn"
	case Err:
		return "err"
	case Regex:
		return "regex"
	case BrowserMasterClose:
		return "browserMasterClose"
	case BrowserMasterTabOpen:
		return "browserMasterTabOpen"
	case BrowserReadElement:
		return "readElement"
	case UploadFile:
		return "uploadFile"
	case OpenTab:
		return "openTab"
	case SetCurrentTab:
		return "setCurrentTab"
	case DbRead:
		return "dbRead"
	case DbWrite:
		return "dbWrite"
	case DbConnect:
		return "dbConnect"
	case DbPing:
		return "dbPing"
	case RESTGET:
		return "get"
	case Post:
		return "post"
	case DeleteFile:
		return "deleteFile"
	case OpenFile:
		return "openFile"
	case TouchFile:
		return "touchFile"
	case CompareFile:
		return "compareFile"
	case JsonContains:
		return "jsonContains"
	case JsonNotContains:
		return "jsonNotContains"
	case WS_Create:
		return "workspace::create"
	case WS_Stop:
		return "workspace::stop"
	case WS_RunCommand:
		return "workspace::runCommand"
	case WS_ArchiveFile:
		return "workspace::archiveFile"
	case WS_StartScript:
		return "workspace::startScript"
	case WS_KillScript:
		return "workspace::killScript"

	case JNK_Connect:
		return "jenkins::connect"
	case JNK_Execute:
		return "jenkins::execute"
	case VLT_Connect:
		return "vault::connect"
	case VLT_Get:
		return "vault::get"
	case Date:
		return "date"
	case SLK_Connect:
		return "slack::connect"
	case SLK_SendMessage:
		return "slack::sendMessage"
	case DIS_Connect:
		return "discord::connect"
	case DIS_SendMessage:
		return "discord::sendMessage"
	case DIS_Stop:
		return "discord::stop"
	case EXT_FreeStringFunctionRegister:
		return "external::freeStringFunctionRegister"
	case EXT_Load:
		return "external::load"
	case EXT_Execute:
		return "external::execute"
	case EXT_Unload:
		return "external::unload"
	case TM_Now:
		return "time::now"
	case ENV_Get:
		return "environment::get"
	case ENV_Set:
		return "environment::set"
	case ENV_Unset:
		return "environment::unset"
	case ENV_All:
		return "environment::all"
	case ENV_Exists:
		return "environment::exists"
	case ENV_Os:
		return "environment::os"
	case ENV_Hostname:
		return "environment::hostname"
	case ENV_Username:
		return "environment::username"
	case ENV_Cwd:
		return "environment::cwd"
	case ENV_Tempdir:
		return "environment::tempDir"
	case ENV_Homedir:
		return "environment::homeDir"
	case ENV_Executioncontext:
		return "environment::executionContext"
	case JSN_Parse:
		return "json::parse"
	case JSN_Set:
		return "json::set"
	case JSN_Remove:
		return "json::remove"
	case JSN_ArrayAppend:
		return "json::arrayAppend"
	case JSN_ArrayRemove:
		return "json::arrayRemove"
	case JSN_Exists:
		return "json::exists"
	case JSN_GetOrDefault:
		return "json::getOrDefault"
	case JSN_Merge:
		return "json::merge"
	case JSN_ArrayGet:
		return "json::arrayGet"
	case JSN_ArrayGetOrDefault:
		return "json::arrayGetOrDefault"
	case JSN_ArrayContains:
		return "json::arrayContains"
	case JSN_ArrayParse:
		return "json::arrayParse"
	case JSN_TestParse:
		return "json::testParse"
	case JSN_ArrayTestParse:
		return "json::arrayTestParse"
	case JSN_Keys:
		return "json::keys"
	case JSN_FilterAttributes:
		return "json::filterAttributes"
	case JSN_ArrayFilterAttributes:
		return "json::arrayFilterAttributes"
	case JSN_Map:
		return "json::map"
	case JSN_ArrayMap:
		return "json::arrayMap"
	case JSN_ArrayIsEmpty:
		return "json::arrayIsEmpty"
	case JSN_ArrayPluck:
		return "json::arrayPluck"
	case JSN_ArrayChunk:
		return "json::arrayChunk"
	case JSN_ArraySlice:
		return "json::arraySlice"
	case JSN_ArrayInsert:
		return "json::arrayInsert"
	case RND_GetIntInRange:
		return "random::getIntInRange"
	case RND_GetDoubleInRange:
		return "random::getDoubleInRange"
	case RND_GetString:
		return "random::getString"
	case RND_GetBool:
		return "random::getBool"
	case RND_GetPick:
		return "random::getPick"
	case RND_GetShuffled:
		return "random::getShuffled"
	case RND_GetUuid:
		return "random::getUuid"
	case RND_Seed:
		return "random::seed"
	case RND_GetStringInRange:
		return "random::getStringInRange"
	case MTH_Round:
		return "math::round"
	case MTH_Abs:
		return "math::abs"
	case MTH_Clamp:
		return "math::clamp"
	case MTH_Min:
		return "math::min"
	case MTH_Max:
		return "math::max"
	case MTH_Sum:
		return "math::sum"
	case MTH_Avg:
		return "math::avg"
	case MTH_Stats:
		return "math::stats"
	case MTH_Percent:
		return "math::percent"
	case MTH_Sin:
		return "math::sin"
	case MTH_Cos:
		return "math::cos"
	case MTH_Tan:
		return "math::tan"
	case MTH_Arctan:
		return "math::arctan"
	case MTH_Sqrt:
		return "math::sqrt"
	case EML_Send:
		return "email::send"
	case EML_SendEach:
		return "email::sendEach"
	case EML_Render:
		return "email::render"
	case EML_Validate:
		return "email::validate"
	case EML_Connect:
		return "email::connect"
	case EML_Close:
		return "email::close"
	case JSN_InferSchema:
		return "json::inferSchema"
	case JSN_MergeSchema:
		return "json::mergeSchema"
	case JSN_Validate:
		return "json::validate"
	case JSN_DiffSchema:
		return "json::diffSchema"
	case JSN_Transform:
		return "json::transform"
	case JSN_ToGo:
		return "json::toGo"
	case JSN_ToTs:
		return "json::toTs"
	case JSN_Fake:
		return "json::fake"
	case HTTP_Request:
		return "http::request"
	case HTTP_Get:
		return "http::get"
	case HTTP_GetAll:
		return "http::getAll"
	case HTTP_Each:
		return "http::each"
	case HTTP_Post:
		return "http::post"
	case HTTP_Put:
		return "http::put"
	case HTTP_Patch:
		return "http::patch"
	case HTTP_Delete:
		return "http::delete"
	case RESTPOST:
		return "restPOSTRequest"
	case RESTPUT:
		return "restPUTRequest"
	case RESTPATCH:
		return "restPATCHRequest"
	case RESTDELETE:
		return "restDELETERequest"
	case TM_NowDateTime:
		return "time::nowDateTime"
	case TM_Today:
		return "time::today"
	case TM_ParseDate:
		return "time::parseDate"
	case TM_ParseDateTime:
		return "time::parseDateTime"
	case TM_Format:
		return "time::format"
	case TM_ToDate:
		return "time::toDate"
	case TM_ToDateTime:
		return "time::toDateTime"
	case TM_Add:
		return "time::add"
	case TM_Diff:
		return "time::diff"
	case TM_Sleep:
		return "time::sleep"
		/*	case SLK_ReadMessage:
			return "slack::readMessage"*/
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

/*var MAX_ARGUMENTS = 1024*/
