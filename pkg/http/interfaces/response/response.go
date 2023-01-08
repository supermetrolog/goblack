package response

type Response interface {
	Content() []byte
	StatusCode() int
	Headers() map[string][]string
}

type ResponseWriter interface {
	SetContent(any) ResponseWriter
	SetStatusCode(int) ResponseWriter
	AddHeader(key string, value string) ResponseWriter
	JsonResponse() (Response, error)
	HtmlResponse() (Response, error)
	XmlResponse() (Response, error)
}
