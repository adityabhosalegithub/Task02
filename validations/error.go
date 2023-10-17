package validations

//built-in interface for representing an error condition
type error interface {
	Error() string //single method Error() returns a string representing the error message.
}

type CustomError struct { //CustomError type is custom error type that implements the error interface
	Message string //field message hold specific err
}

func (e CustomError) Error() string { //Err0r() method of customerror //defined to satisfy custom error
	return e.Message
}

// type ParseError struct {
// 	Field string
// 	Err   error
// }

// func (e ParseError) Error() string {
// 	return fmt.Sprintf("Field: %s, Error: %v", e.Field, e.Err)
// }
