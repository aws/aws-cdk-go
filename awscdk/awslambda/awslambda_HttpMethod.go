package awslambda


// All http request methods.
// Experimental.
type HttpMethod string

const (
	// The GET method requests a representation of the specified resource.
	// Experimental.
	HttpMethod_GET HttpMethod = "GET"
	// The PUT method replaces all current representations of the target resource with the request payload.
	// Experimental.
	HttpMethod_PUT HttpMethod = "PUT"
	// The HEAD method asks for a response identical to that of a GET request, but without the response body.
	// Experimental.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
	// Experimental.
	HttpMethod_POST HttpMethod = "POST"
	// The DELETE method deletes the specified resource.
	// Experimental.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// The PATCH method applies partial modifications to a resource.
	// Experimental.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// The OPTIONS method describes the communication options for the target resource.
	// Experimental.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// The wildcard entry to allow all methods.
	// Experimental.
	HttpMethod_ALL HttpMethod = "ALL"
)

