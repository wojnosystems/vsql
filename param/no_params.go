package param

// NewAppend creates a new appending Parameterer in which you can repeatedly append values to the parameter list as desired
func New( query string ) Appender {
	return &appender{
		query: query,
		parameters: make([]interface{},0),
	}
}
