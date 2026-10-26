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
	MTH_Pi
	MTH_E
	MTH_Pow
	MTH_Floor
	MTH_Ceil
	MTH_Trunc
	MTH_Mod
	MTH_Exp
	MTH_Log
	MTH_Log10
	MTH_Log2
	MTH_Arcsin
	MTH_Arccos
	MTH_Arctan2
	MTH_Sign
	MTH_Hypot
	MTH_Cbrt
	MTH_ToRadians
	MTH_ToDegrees
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
	HTTP_Ping
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
	CRY_Hash
	CRY_HashFile
	CRY_Hmac
	CRY_HmacVerify
	CRY_Base64Encode
	CRY_Base64Decode
	CRY_RandomToken
	CRY_Encrypt
	CRY_Decrypt
	CSV_Open
	CSV_SetHeaders
	CSV_Headers
	CSV_RowCount
	CSV_ReadAll
	CSV_ReadRows
	CSV_AppendRow
	CSV_AppendRows
	CSV_WriteAll
	CSV_Close
	CSV_Parse
	CSV_ToString
	CSV_OpenWithEncoding
	CSV_Save
	CSV_ToJson
	CSV_ToExcel
	CSV_ReadRow
	ST_Open
	ST_Get
	ST_Set
	ST_SetWithExpiry
	ST_SetField
	ST_SetIfAbsent
	ST_Exists
	ST_Delete
	ST_Increment
	ST_CompareAndSet
	ST_Keys
	ST_Clear
	TBL_Select
	TBL_Where
	TBL_SortBy
	TBL_GroupBy
	TBL_Aggregate
	TBL_Join
	TBL_Distinct
	TBL_Diff
	TBL_AddColumn
	TBL_Pivot
	HTM_Select
	HTM_SelectFirst
	HTM_Text
	HTM_Texts
	HTM_Attr
	HTM_Attrs
	HTM_Count
	HTM_Exists
	HTM_Table
	HTM_Links
	HTM_ToText
	HTM_Meta
	FIL_Append
	FIL_ReadLines
	FIL_WriteLines
	FIL_Info
	FIL_IsFolder
	FIL_List
	FIL_CreateFolder
	FIL_RemoveFolder
	FIL_Copy
	FIL_Move
	FIL_Delete
	FIL_Zip
	FIL_Unzip
	FIL_WaitFor
	FIL_Join
	FIL_Name
	FIL_Extension
	FIL_ParentFolder
	FIL_ResourcesFolder
	FIL_OutputsFolder
	FIL_Print
	FIL_Permissions
	SKT_Create
	SKT_Connect
	SKT_Send
	SKT_Receive
	SKT_IsConnected
	SKT_Close
	PDF_Template
	PDF_TemplateText
	PDF_AddToContext
	PDF_Render
	PDF_Open
	PDF_Info
	PDF_Text
	PDF_Merge
	PDF_Split
	PDF_SelectPages
	PDF_Watermark
	PDF_Close
	XLS_Create
	XLS_Open
	XLS_Save
	XLS_Close
	XLS_ToJson
	XLS_ToCsv
	XLS_Sheets
	XLS_AddSheet
	XLS_RemoveSheet
	XLS_ReadSheet
	XLS_ReadRange
	XLS_ReadCell
	XLS_WriteSheet
	XLS_AppendRows
	XLS_WriteCell
	XLS_WriteFormula
	JSN_ToCsv
	JSN_ToExcel
	TXT_Length
	TXT_Lower
	TXT_Upper
	TXT_Trim
	TXT_RemoveAccents
	TXT_Replace
	TXT_Contains
	TXT_StartsWith
	TXT_EndsWith
	TXT_IndexOf
	TXT_Substring
	TXT_Split
	TXT_Join
	TXT_Lines
	TXT_PadLeft
	TXT_PadRight
	TXT_Repeat
	TXT_ToNumber
	TXT_FromNumber
	RGX_Matches
	RGX_Find
	RGX_FindAll
	RGX_Extract
	RGX_ExtractAll
	RGX_Replace
	RGX_Split
	RGX_Escape
	SSH_HostKey
	SSH_Connect
	SSH_ConnectWithKey
	SSH_Run
	SSH_Upload
	SSH_Download
	SSH_IsConnected
	SSH_Close
	RTY_SetSshConnect
	RTY_GetSshConnect
	RTY_SetSshTransfer
	RTY_GetSshTransfer
	RTY_SetSshRun
	RTY_GetSshRun
	RTY_SetSocketConnect
	RTY_GetSocketConnect
	RTY_SetHttpGet
	RTY_GetHttpGet
	RTY_SetHttpPut
	RTY_GetHttpPut
	RTY_SetHttpDelete
	RTY_GetHttpDelete
	RTY_SetHttpPost
	RTY_GetHttpPost
	RTY_SetHttpPatch
	RTY_GetHttpPatch
	RTY_SetPdfRender
	RTY_GetPdfRender
	RTY_SetFileWrite
	RTY_GetFileWrite
	RTY_SetEmailSend
	RTY_GetEmailSend
	RTY_SetDatabaseConnect
	RTY_GetDatabaseConnect
	RTY_ResetAll
	RTY_DisableAll
	RTY_List
	RTY_Last
	DEC_Add
	DEC_Subtract
	DEC_Multiply
	DEC_Divide
	DEC_Round
	DEC_Sum
	DEC_Compare
	DEC_Percent
	DEC_Allocate
	DEC_Parse
	DEC_Format
	DEC_ToNumber
	DEC_IsValid
	DEC_Abs
	UNT_Convert
	UNT_Kind
	UNT_Kinds
	UNT_List
	UNT_FormatBytes
	UNT_ParseBytes
	UNT_FormatDuration
	UNT_ParseDuration
	XML_Parse
	XML_Open
	XML_Create
	XML_Save
	XML_ToString
	XML_Select
	XML_SelectFirst
	XML_Count
	XML_Exists
	XML_SelectRows
	XML_Set
	XML_AddElement
	XML_Remove
	XML_Namespace
	XML_ToJson
	XML_FromRows
	XML_ToCsv
	JSN_ToXml
	GIT_Create
	GIT_Clone
	GIT_Pull
	GIT_Status
	GIT_Log
	GIT_Tag
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
	CSV
	STATE
	SOCKET
	PDF
	EXCEL
	SSH
	XML
	GIT
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
	case "csv":
		return CSV
	case "state":
		return STATE
	case "socket":
		return SOCKET
	case "pdf":
		return PDF
	case "excel":
		return EXCEL
	case "ssh":
		return SSH
	case "xml":
		return XML
	case "git":
		return GIT
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
	case CSV:
		return "CSV"
	case STATE:
		return "State"
	case SOCKET:
		return "Socket"
	case PDF:
		return "PDF"
	case EXCEL:
		return "Excel"
	case SSH:
		return "Ssh"
	case XML:
		return "Xml"
	case GIT:
		return "Git"
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
	case MTH_Pi:
		return "math::pi"
	case MTH_E:
		return "math::e"
	case MTH_Pow:
		return "math::pow"
	case MTH_Floor:
		return "math::floor"
	case MTH_Ceil:
		return "math::ceil"
	case MTH_Trunc:
		return "math::trunc"
	case MTH_Mod:
		return "math::mod"
	case MTH_Exp:
		return "math::exp"
	case MTH_Log:
		return "math::log"
	case MTH_Log10:
		return "math::log10"
	case MTH_Log2:
		return "math::log2"
	case MTH_Arcsin:
		return "math::arcsin"
	case MTH_Arccos:
		return "math::arccos"
	case MTH_Arctan2:
		return "math::arctan2"
	case MTH_Sign:
		return "math::sign"
	case MTH_Hypot:
		return "math::hypot"
	case MTH_Cbrt:
		return "math::cbrt"
	case MTH_ToRadians:
		return "math::toRadians"
	case MTH_ToDegrees:
		return "math::toDegrees"
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
	case HTTP_Ping:
		return "http::ping"
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
	case CRY_Hash:
		return "crypto::hash"
	case CRY_HashFile:
		return "crypto::hashFile"
	case CRY_Hmac:
		return "crypto::hmac"
	case CRY_HmacVerify:
		return "crypto::hmacVerify"
	case CRY_Base64Encode:
		return "crypto::base64Encode"
	case CRY_Base64Decode:
		return "crypto::base64Decode"
	case CRY_RandomToken:
		return "crypto::randomToken"
	case CRY_Encrypt:
		return "crypto::encrypt"
	case CRY_Decrypt:
		return "crypto::decrypt"
	case CSV_Open:
		return "csv::open"
	case CSV_SetHeaders:
		return "csv::setHeaders"
	case CSV_Headers:
		return "csv::headers"
	case CSV_RowCount:
		return "csv::rowCount"
	case CSV_ReadAll:
		return "csv::readAll"
	case CSV_ReadRows:
		return "csv::readRows"
	case CSV_AppendRow:
		return "csv::appendRow"
	case CSV_AppendRows:
		return "csv::appendRows"
	case CSV_WriteAll:
		return "csv::writeAll"
	case CSV_Close:
		return "csv::close"
	case CSV_Parse:
		return "csv::parse"
	case CSV_ToString:
		return "csv::toString"
	case CSV_OpenWithEncoding:
		return "csv::openWithEncoding"
	case CSV_Save:
		return "csv::save"
	case CSV_ToJson:
		return "csv::toJson"
	case CSV_ToExcel:
		return "csv::toExcel"
	case CSV_ReadRow:
		return "csv::readRow"
	case ST_Open:
		return "state::open"
	case ST_Get:
		return "state::get"
	case ST_Set:
		return "state::set"
	case ST_SetWithExpiry:
		return "state::setWithExpiry"
	case ST_SetField:
		return "state::setField"
	case ST_SetIfAbsent:
		return "state::setIfAbsent"
	case ST_Exists:
		return "state::exists"
	case ST_Delete:
		return "state::delete"
	case ST_Increment:
		return "state::increment"
	case ST_CompareAndSet:
		return "state::compareAndSet"
	case ST_Keys:
		return "state::keys"
	case ST_Clear:
		return "state::clear"
	case TBL_Select:
		return "table::select"
	case TBL_Where:
		return "table::where"
	case TBL_SortBy:
		return "table::sortBy"
	case TBL_GroupBy:
		return "table::groupBy"
	case TBL_Aggregate:
		return "table::aggregate"
	case TBL_Join:
		return "table::join"
	case TBL_Distinct:
		return "table::distinct"
	case TBL_Diff:
		return "table::diff"
	case TBL_AddColumn:
		return "table::addColumn"
	case TBL_Pivot:
		return "table::pivot"
	case HTM_Select:
		return "html::select"
	case HTM_SelectFirst:
		return "html::selectFirst"
	case HTM_Text:
		return "html::text"
	case HTM_Texts:
		return "html::texts"
	case HTM_Attr:
		return "html::attr"
	case HTM_Attrs:
		return "html::attrs"
	case HTM_Count:
		return "html::count"
	case HTM_Exists:
		return "html::exists"
	case HTM_Table:
		return "html::table"
	case HTM_Links:
		return "html::links"
	case HTM_ToText:
		return "html::toText"
	case HTM_Meta:
		return "html::meta"
	case FIL_Append:
		return "file::append"
	case FIL_ReadLines:
		return "file::readLines"
	case FIL_WriteLines:
		return "file::writeLines"
	case FIL_Info:
		return "file::info"
	case FIL_IsFolder:
		return "file::isFolder"
	case FIL_List:
		return "file::list"
	case FIL_CreateFolder:
		return "file::createFolder"
	case FIL_RemoveFolder:
		return "file::removeFolder"
	case FIL_Copy:
		return "file::copy"
	case FIL_Move:
		return "file::move"
	case FIL_Delete:
		return "file::delete"
	case FIL_Zip:
		return "file::zip"
	case FIL_Unzip:
		return "file::unzip"
	case FIL_WaitFor:
		return "file::waitFor"
	case FIL_Join:
		return "file::join"
	case FIL_Name:
		return "file::name"
	case FIL_Extension:
		return "file::extension"
	case FIL_ParentFolder:
		return "file::parentFolder"
	case FIL_ResourcesFolder:
		return "file::resourcesFolder"
	case FIL_OutputsFolder:
		return "file::outputsFolder"
	case FIL_Print:
		return "file::print"
	case FIL_Permissions:
		return "file::permissions"
	case SKT_Create:
		return "socket::create"
	case SKT_Connect:
		return "socket::connect"
	case SKT_Send:
		return "socket::send"
	case SKT_Receive:
		return "socket::receive"
	case SKT_IsConnected:
		return "socket::isConnected"
	case SKT_Close:
		return "socket::close"
	case PDF_Template:
		return "pdf::template"
	case PDF_TemplateText:
		return "pdf::templateText"
	case PDF_AddToContext:
		return "pdf::addToContext"
	case PDF_Render:
		return "pdf::render"
	case PDF_Open:
		return "pdf::open"
	case PDF_Info:
		return "pdf::info"
	case PDF_Text:
		return "pdf::text"
	case PDF_Merge:
		return "pdf::merge"
	case PDF_Split:
		return "pdf::split"
	case PDF_SelectPages:
		return "pdf::selectPages"
	case PDF_Watermark:
		return "pdf::watermark"
	case PDF_Close:
		return "pdf::close"
	case XLS_Create:
		return "excel::create"
	case XLS_Open:
		return "excel::open"
	case XLS_Save:
		return "excel::save"
	case XLS_Close:
		return "excel::close"
	case XLS_ToJson:
		return "excel::toJson"
	case XLS_ToCsv:
		return "excel::toCsv"
	case XLS_Sheets:
		return "excel::sheets"
	case XLS_AddSheet:
		return "excel::addSheet"
	case XLS_RemoveSheet:
		return "excel::removeSheet"
	case XLS_ReadSheet:
		return "excel::readSheet"
	case XLS_ReadRange:
		return "excel::readRange"
	case XLS_ReadCell:
		return "excel::readCell"
	case XLS_WriteSheet:
		return "excel::writeSheet"
	case XLS_AppendRows:
		return "excel::appendRows"
	case XLS_WriteCell:
		return "excel::writeCell"
	case XLS_WriteFormula:
		return "excel::writeFormula"
	case JSN_ToCsv:
		return "json::toCsv"
	case JSN_ToExcel:
		return "json::toExcel"
	case TXT_Length:
		return "text::length"
	case TXT_Lower:
		return "text::lower"
	case TXT_Upper:
		return "text::upper"
	case TXT_Trim:
		return "text::trim"
	case TXT_RemoveAccents:
		return "text::removeAccents"
	case TXT_Replace:
		return "text::replace"
	case TXT_Contains:
		return "text::contains"
	case TXT_StartsWith:
		return "text::startsWith"
	case TXT_EndsWith:
		return "text::endsWith"
	case TXT_IndexOf:
		return "text::indexOf"
	case TXT_Substring:
		return "text::substring"
	case TXT_Split:
		return "text::split"
	case TXT_Join:
		return "text::join"
	case TXT_Lines:
		return "text::lines"
	case TXT_PadLeft:
		return "text::padLeft"
	case TXT_PadRight:
		return "text::padRight"
	case TXT_Repeat:
		return "text::repeat"
	case TXT_ToNumber:
		return "text::toNumber"
	case TXT_FromNumber:
		return "text::fromNumber"
	case RGX_Matches:
		return "regex::matches"
	case RGX_Find:
		return "regex::find"
	case RGX_FindAll:
		return "regex::findAll"
	case RGX_Extract:
		return "regex::extract"
	case RGX_ExtractAll:
		return "regex::extractAll"
	case RGX_Replace:
		return "regex::replace"
	case RGX_Split:
		return "regex::split"
	case RGX_Escape:
		return "regex::escape"
	case SSH_HostKey:
		return "ssh::hostKey"
	case SSH_Connect:
		return "ssh::connect"
	case SSH_ConnectWithKey:
		return "ssh::connectWithKey"
	case SSH_Run:
		return "ssh::run"
	case SSH_Upload:
		return "ssh::upload"
	case SSH_Download:
		return "ssh::download"
	case SSH_IsConnected:
		return "ssh::isConnected"
	case SSH_Close:
		return "ssh::close"
	case RTY_SetSshConnect:
		return "retry::setSshConnect"
	case RTY_GetSshConnect:
		return "retry::getSshConnect"
	case RTY_SetSshTransfer:
		return "retry::setSshTransfer"
	case RTY_GetSshTransfer:
		return "retry::getSshTransfer"
	case RTY_SetSshRun:
		return "retry::setSshRun"
	case RTY_GetSshRun:
		return "retry::getSshRun"
	case RTY_SetSocketConnect:
		return "retry::setSocketConnect"
	case RTY_GetSocketConnect:
		return "retry::getSocketConnect"
	case RTY_SetHttpGet:
		return "retry::setHttpGet"
	case RTY_GetHttpGet:
		return "retry::getHttpGet"
	case RTY_SetHttpPut:
		return "retry::setHttpPut"
	case RTY_GetHttpPut:
		return "retry::getHttpPut"
	case RTY_SetHttpDelete:
		return "retry::setHttpDelete"
	case RTY_GetHttpDelete:
		return "retry::getHttpDelete"
	case RTY_SetHttpPost:
		return "retry::setHttpPost"
	case RTY_GetHttpPost:
		return "retry::getHttpPost"
	case RTY_SetHttpPatch:
		return "retry::setHttpPatch"
	case RTY_GetHttpPatch:
		return "retry::getHttpPatch"
	case RTY_SetPdfRender:
		return "retry::setPdfRender"
	case RTY_GetPdfRender:
		return "retry::getPdfRender"
	case RTY_SetFileWrite:
		return "retry::setFileWrite"
	case RTY_GetFileWrite:
		return "retry::getFileWrite"
	case RTY_SetEmailSend:
		return "retry::setEmailSend"
	case RTY_GetEmailSend:
		return "retry::getEmailSend"
	case RTY_SetDatabaseConnect:
		return "retry::setDatabaseConnect"
	case RTY_GetDatabaseConnect:
		return "retry::getDatabaseConnect"
	case RTY_ResetAll:
		return "retry::resetAll"
	case RTY_DisableAll:
		return "retry::disableAll"
	case RTY_List:
		return "retry::list"
	case RTY_Last:
		return "retry::last"
	case DEC_Add:
		return "decimal::add"
	case DEC_Subtract:
		return "decimal::subtract"
	case DEC_Multiply:
		return "decimal::multiply"
	case DEC_Divide:
		return "decimal::divide"
	case DEC_Round:
		return "decimal::round"
	case DEC_Sum:
		return "decimal::sum"
	case DEC_Compare:
		return "decimal::compare"
	case DEC_Percent:
		return "decimal::percent"
	case DEC_Allocate:
		return "decimal::allocate"
	case DEC_Parse:
		return "decimal::parse"
	case DEC_Format:
		return "decimal::format"
	case DEC_ToNumber:
		return "decimal::toNumber"
	case DEC_IsValid:
		return "decimal::isValid"
	case DEC_Abs:
		return "decimal::abs"
	case UNT_Convert:
		return "unit::convert"
	case UNT_Kind:
		return "unit::kind"
	case UNT_Kinds:
		return "unit::kinds"
	case UNT_List:
		return "unit::list"
	case UNT_FormatBytes:
		return "unit::formatBytes"
	case UNT_ParseBytes:
		return "unit::parseBytes"
	case UNT_FormatDuration:
		return "unit::formatDuration"
	case UNT_ParseDuration:
		return "unit::parseDuration"
	case XML_Parse:
		return "xml::parse"
	case XML_Open:
		return "xml::open"
	case XML_Create:
		return "xml::create"
	case XML_Save:
		return "xml::save"
	case XML_ToString:
		return "xml::toString"
	case XML_Select:
		return "xml::select"
	case XML_SelectFirst:
		return "xml::selectFirst"
	case XML_Count:
		return "xml::count"
	case XML_Exists:
		return "xml::exists"
	case XML_SelectRows:
		return "xml::selectRows"
	case XML_Set:
		return "xml::set"
	case XML_AddElement:
		return "xml::addElement"
	case XML_Remove:
		return "xml::remove"
	case XML_Namespace:
		return "xml::namespace"
	case XML_ToJson:
		return "xml::toJson"
	case XML_FromRows:
		return "xml::fromRows"
	case XML_ToCsv:
		return "xml::toCsv"
	case JSN_ToXml:
		return "json::toXml"
	case GIT_Create:
		return "git::create"
	case GIT_Clone:
		return "git::clone"
	case GIT_Pull:
		return "git::pull"
	case GIT_Status:
		return "git::status"
	case GIT_Log:
		return "git::log"
	case GIT_Tag:
		return "git::tag"
		/*	case SLK_ReadMessage:
			return "slack::readMessage"*/
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

/*var MAX_ARGUMENTS = 1024*/
