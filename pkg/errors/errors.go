package errors

const (
	ErrSettingCache      = "Error happened while setting cache."
	ErrJSONMarshal       = "Error happened in JSON marshal."
	ErrParsingCache      = "Error happened while parsing cache data."
	ErrGettingCache      = "Error happened while getting cache."
	ErrUnsupportedType   = "Unsupported param type in url"
	ErrUnsupportedMethod = "Unsupported method"
	ErrGettingConfig     = "Error happend while getting config"
	ErrWrittingResponse  = "Error happend while writting http response"
)
