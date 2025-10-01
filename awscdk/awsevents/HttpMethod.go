package awsevents


// Supported HTTP operations.
type HttpMethod string

const (
	// POST.
	HttpMethod_POST HttpMethod = "POST"
	// GET.
	HttpMethod_GET HttpMethod = "GET"
	// HEAD.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// OPTIONS.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// PUT.
	HttpMethod_PUT HttpMethod = "PUT"
	// PATCH.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// DELETE.
	HttpMethod_DELETE HttpMethod = "DELETE"
)

