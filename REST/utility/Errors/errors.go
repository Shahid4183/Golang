package errors

// Error type - wrapper around Go's error type
// this is essential to create constants of error type to return
// in case of error of any api call
type Error string

// Our Error type must implement error interface by defining Error function
func (e Error) Error() string { return string(e) }

// InvalidRequest - this error is thrown when api request is invalid
// this error will occure while binding request body object to context of api call
const InvalidRequest = Error("Invalid Request")

// ValidationError - this error is thrown when there is validation error in api request body
// this error will occure when checking validation given to a strucutre
const ValidationError = Error("Validation Error")

// InvalidJWTToken - this error is throw when JWT token is invalid or Expired
// this error will occure when JWT token is expire or invalid
const InvalidJWTToken = Error("Invalid or Expired JWT token")
