package docs

//invalid json or get parameters
//swagger:response badrequest
type BadRequest struct {
	//in: body
	Body ErrorRequestWrap
}

type ErrorRequestWrap struct {
	Code int   `json:"code"`
	Err  error `json:"error,omitempty"`
}

//data doesn't exist
//swagger:response notfound
type NotFound struct {
	//in: body
	Body ErrorRequestWrap
}
