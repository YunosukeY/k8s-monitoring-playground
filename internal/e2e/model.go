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

type Query struct {
	DSID int    `json:"datasourceId"`
	Expr string `json:"expr"`
}
type Queries struct {
	Queries []Query `json:"queries"`
	From    string  `json:"from"`
	To      string  `json:"to"`
}

type Schema struct {
	Name string `json:"name"`
}
type Data struct {
	Values [][]int `json:"values"`
}
type Frame struct {
	Schema Schema `json:"schema"`
	Data   Data   `json:"data"`
}
type Result struct {
	Status int     `json:"status"`
	Frames []Frame `json:"frames"`
}
type Results struct {
	A Result
}
type QueryResponse struct {
	Results Results `json:"results"`
}
