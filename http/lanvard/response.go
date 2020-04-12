package lanvard

type Response struct {
	content string
}

func NewResponse() Response {
	return Response{}
}

func Json(content string) Response {
	return Response{content: content}
}

func (r Response) Content() string {
	return r.content
}

func (r Response) SetContent(content string) Response {
	r.content = content

	return r
}
