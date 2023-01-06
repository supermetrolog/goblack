package response

type Response interface {
	Content() []byte
	StatusCode() int
	Headers() map[string]string
}

type ResponseWriter interface {
	SetContent(any)
	SetStatusCode(int)
	AddHeader(key string, value string)
	JsonResponse() (Response, error)
	HtmlResponse() (Response, error)
	Response() Response
}
