package apis

type GetQueryStringError struct {  
	err    string  //error description
	other  string
}

func (e *GetQueryStringError) Error() string {  
	return e.err
}

func (e *GetQueryStringError) otherError() bool {  
	return  e.other != "ok"
}