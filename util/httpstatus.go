package util

type HTTPStatus struct {
	code int
	message string
	series Series
}

type Series string

func (status HTTPStatus) isInformational() bool {
	return status.series == INFORMATIONAL
}

func (status HTTPStatus) isSuccess() bool {
	return status.series == SUCCESS
}

func (status HTTPStatus) isRedirection() bool {
	return status.series == REDIRECTION
}

func (status HTTPStatus) isClientError() bool {
	return status.series == CLIENT
}

func (status HTTPStatus) isServerError() bool {
	return status.series == SERVER
}

func (status HTTPStatus) isError() bool {
	return status.isClientError() || status.isServerError()
}

// SERIES DECLARATION
//goland:noinspection ALL
const (
	INFORMATIONAL Series = "Informational Response"
	SUCCESS Series = "Success"
	REDIRECTION Series = "Redirection"
	CLIENT Series = "Client Error"
	SERVER Series = "Server Error"
)

// HTTP STATUS DECLARATION
//goland:noinspection ALL
var (
	CONTINUE = HTTPStatus{code: 100, message: "Continue", series: INFORMATIONAL}
	SWITCHING_PROTOCOLS = HTTPStatus{code: 101, message: "Switching Protocols", series: INFORMATIONAL}

	OK = HTTPStatus{code: 200, message: "OK", series: SUCCESS}
	CREATED = HTTPStatus{code: 201, message: "Created", series: SUCCESS}
	ACCEPTED = HTTPStatus{code: 202, message: "Accepted", series: SUCCESS}
	NON_AUTHORITATIVE_INFORMATION = HTTPStatus{code: 203, message: "Non-Authoritative Information", series: SUCCESS}
	NO_CONTENT = HTTPStatus{code: 204, message: "No Content", series: SUCCESS}
	RESET_CONTENT = HTTPStatus{code: 205, message: "Reset Content", series: SUCCESS}

	MULTIPLE_CHOICES = HTTPStatus{code: 300, message: "Multiple Choices", series: REDIRECTION}
	MOVED_PERMANENTLY = HTTPStatus{code: 301, message: "Moved Permanently", series: REDIRECTION}
	FOUND = HTTPStatus{code: 302, message: "Found", series: REDIRECTION}
	SEE_OTHER = HTTPStatus{code: 303, message: "See Other", series: REDIRECTION}
	NOT_MODIFIED = HTTPStatus{code: 304, message: "Not Modified", series: REDIRECTION}
	USE_PROXY = HTTPStatus{code: 305, message: "Use Proxy", series: REDIRECTION}
	TEMPORARY_REDIRECT = HTTPStatus{code: 307, message: "Temporary Redirect", series: REDIRECTION}
	PERMANENT_REDIRECT = HTTPStatus{code: 308, message: "Permanent Redirect", series: REDIRECTION}

	BAD_REQUEST = HTTPStatus{code: 400, message: "Bad Request", series: CLIENT}
	UNAUTHORIZED = HTTPStatus{code: 401, message: "Unauthorized", series: CLIENT}
	PAYMENT_REQUIRED = HTTPStatus{code: 402, message: "Payment Required", series: CLIENT}
	FORBIDDEN = HTTPStatus{code: 403, message: "Forbidden", series: CLIENT}
	NOT_FOUND = HTTPStatus{code: 404, message: "Not Found", series: CLIENT}
	METHOD_NOT_ALLOWED = HTTPStatus{code: 405, message: "Method Not Allowed", series: CLIENT}
	NOT_ACCEPTABLE = HTTPStatus{code: 406, message: "Not Acceptable", series: CLIENT}
	PROXY_AUTHENTICATION_REQUIRED = HTTPStatus{code: 407, message: "Proxy Authentication Required", series: CLIENT}
	REQUEST_TIMEOUT = HTTPStatus{code: 408, message: "Request Timeout", series: CLIENT}
	CONFLICT = HTTPStatus{code: 409, message: "Conflict", series: CLIENT}
	GONE = HTTPStatus{code: 410, message: "Gone", series: CLIENT}
	BLAZE_IT = HTTPStatus{code: 420, message: "Blaze It", series: CLIENT}

	INTERNAL_SERVER_ERROR = HTTPStatus{code: 500, message: "Internal Server Error", series: SERVER}
	NOT_IMPLEMENTED = HTTPStatus{code: 501, message: "Not Implemented", series: SERVER}
	BAD_GATEWAY = HTTPStatus{code: 502, message: "Bad Gateway", series: SERVER}
	SERVICE_UNAVAILABLE = HTTPStatus{code: 503, message: "Service Unavailable", series: SERVER}
	GATEWAY_TIMEOUT = HTTPStatus{code: 504, message: "Gateway Timeout", series: SERVER};
)
