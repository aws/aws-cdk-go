package awss3


// All http request methods.
// Experimental.
type HttpMethods string

const (
	// The GET method requests a representation of the specified resource.
	// Experimental.
	HttpMethods_GET HttpMethods = "GET"
	// The PUT method replaces all current representations of the target resource with the request payload.
	// Experimental.
	HttpMethods_PUT HttpMethods = "PUT"
	// The HEAD method asks for a response identical to that of a GET request, but without the response body.
	// Experimental.
	HttpMethods_HEAD HttpMethods = "HEAD"
	// The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
	// Experimental.
	HttpMethods_POST HttpMethods = "POST"
	// The DELETE method deletes the specified resource.
	// Experimental.
	HttpMethods_DELETE HttpMethods = "DELETE"
)

