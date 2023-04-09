package services

type FormatCampusNetServiceError struct {  
	err    string  //error description
	other  string
}

func (e *FormatCampusNetServiceError) Error() string {  
	return e.err
}

func (e *FormatCampusNetServiceError) otherError() bool {  
	return  e.other != "ok"
}