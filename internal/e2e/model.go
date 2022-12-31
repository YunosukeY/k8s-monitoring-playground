package e2e

type TodoForPostRequest struct {
	Content string `json:"content" binding:"required"`
}

type TodoForResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type Mail struct {
	To  string  `json:"to" binding:"required,email"`
	Sub *string `json:"sub" binding:"required"`
	Msg *string `json:"msg" binding:"required"`
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RawMessage struct {
	From string
	To   []string
	Data string
}

type Message struct {
	Raw RawMessage
}

type Messages struct {
	Items []Message `json:"items"`
}

type Span struct {
	SpanID string `json:"spanID"`
}

type Trace struct {
	TraceID string `json:"traceID"`
	Spans   []Span `json:"spans"`
}

type Traces struct {
	Data []Trace `json:"data"`
}
