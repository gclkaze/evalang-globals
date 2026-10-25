package globals

// Groupings of the random::, math::, email::, new json:: and http:: operations declared in
// types.go (RND_*, MTH_*, EML_*, JSN_InferSchema..JSN_Fake and HTTP_*, appended at the end of the StatementOp list so existing
// operation numbers do not shift).

// RandomOperations lists every random:: operation.
var RandomOperations = []StatementOp{
	RND_GetIntInRange, RND_GetDoubleInRange, RND_GetString, RND_GetStringInRange, RND_GetBool,
	RND_GetPick, RND_GetShuffled, RND_GetUuid, RND_Seed,
}

// MathOperations lists every math:: operation.
var MathOperations = []StatementOp{
	MTH_Round, MTH_Abs, MTH_Clamp, MTH_Min, MTH_Max, MTH_Sum, MTH_Avg,
	MTH_Stats, MTH_Percent, MTH_Sin, MTH_Cos, MTH_Tan, MTH_Arctan,
	MTH_Sqrt, MTH_Pi, MTH_E, MTH_Pow, MTH_Floor, MTH_Ceil,
	MTH_Trunc, MTH_Mod, MTH_Exp, MTH_Log, MTH_Log10, MTH_Log2,
	MTH_Arcsin, MTH_Arccos, MTH_Arctan2, MTH_Sign,
	MTH_Hypot, MTH_Cbrt, MTH_ToRadians, MTH_ToDegrees,
}

// EmailOperations lists every email:: operation.
var EmailOperations = []StatementOp{
	EML_Connect, EML_Close, EML_Send, EML_SendEach, EML_Render, EML_Validate,
}

// JsonSchemaOperations lists the json:: additions (schemas, transform, codegen, fake data).
var JsonSchemaOperations = []StatementOp{
	JSN_InferSchema, JSN_MergeSchema, JSN_Validate, JSN_DiffSchema,
	JSN_Transform, JSN_ToGo, JSN_ToTs, JSN_Fake,
}

// HttpOperations lists every http:: operation. restClientCreate and
// restGETRequest stay as they are; Http:: ops accept their client.
var HttpOperations = []StatementOp{
	HTTP_Request, HTTP_Get, HTTP_GetAll, HTTP_Each,
	HTTP_Post, HTTP_Put, HTTP_Patch, HTTP_Delete, HTTP_Ping,
}

// RESTMethodOperations are restPOSTRequest, restPUTRequest, restPATCHRequest
// and restDELETERequest: the restGETRequest form for the other methods.
var RESTMethodOperations = []StatementOp{RESTPOST, RESTPUT, RESTPATCH, RESTDELETE}

// TimeOperations are the time:: additions; time::now (Unix timestamp) is
// unchanged and not listed.
var TimeOperations = []StatementOp{
	TM_NowDateTime, TM_Today, TM_ParseDate, TM_ParseDateTime, TM_Format,
	TM_ToDate, TM_ToDateTime, TM_Add, TM_Diff, TM_Sleep,
}

// StatementOnlyOperations have no expression form, because they produce no
// value. Every other random:: and math:: operation can also be called inside
// an expression, e.g. $x = Math::round(Random::getDoubleInRange(0, 20), 2).
var StatementOnlyOperations = map[StatementOp]bool{
	RND_Seed:     true,
	EML_Send:     true, // side effect
	EML_SendEach: true, // side effect
	EML_Validate: true, // two outputs
	EML_Connect:  true, // returns a handle
	EML_Close:    true, // side effect
	JSN_Validate: true, // two outputs
	HTTP_Request: true, // network side effect, three outputs
	HTTP_Get:     true, // two outputs
	HTTP_GetAll:  true, // two outputs
	HTTP_Each:    true, // network side effects
	HTTP_Post:    true, // network side effect, two outputs
	HTTP_Put:     true, // network side effect, two outputs
	HTTP_Patch:   true, // network side effect, two outputs
	HTTP_Delete:  true, // network side effect, two outputs
	HTTP_Ping:    true, // network side effect
	TM_Sleep:     true, // no value
}
