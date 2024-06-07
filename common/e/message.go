package e

var MsgFlags = map[int]string{
	SUCCESS:        "Ok",
	CREATED:        "Ok",
	ERROR:          "Fail",
	INVALID_PARAMS: "Invalid parameters",
	UNAUTHORIZED:   "Unauthorized",
	FORBIDDEN:      "FORBIDDEN",
	NOTFOUND:       "Not found",
	CONFLICT:       "Conflict",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
