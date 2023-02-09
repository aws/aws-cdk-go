package awsevents


// Supported HTTP operations.
// Experimental.
type HttpMethod string

const (
	// POST.
	// Experimental.
	HttpMethod_POST HttpMethod = "POST"
	// GET.
	// Experimental.
	HttpMethod_GET HttpMethod = "GET"
	// HEAD.
	// Experimental.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// OPTIONS.
	// Experimental.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// PUT.
	// Experimental.
	HttpMethod_PUT HttpMethod = "PUT"
	// PATCH.
	// Experimental.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// DELETE.
	// Experimental.
	HttpMethod_DELETE HttpMethod = "DELETE"
)

