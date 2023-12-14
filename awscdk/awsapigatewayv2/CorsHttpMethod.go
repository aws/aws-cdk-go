package awsapigatewayv2


// Supported CORS HTTP methods.
type CorsHttpMethod string

const (
	// HTTP ANY.
	CorsHttpMethod_ANY CorsHttpMethod = "ANY"
	// HTTP DELETE.
	CorsHttpMethod_DELETE CorsHttpMethod = "DELETE"
	// HTTP GET.
	CorsHttpMethod_GET CorsHttpMethod = "GET"
	// HTTP HEAD.
	CorsHttpMethod_HEAD CorsHttpMethod = "HEAD"
	// HTTP OPTIONS.
	CorsHttpMethod_OPTIONS CorsHttpMethod = "OPTIONS"
	// HTTP PATCH.
	CorsHttpMethod_PATCH CorsHttpMethod = "PATCH"
	// HTTP POST.
	CorsHttpMethod_POST CorsHttpMethod = "POST"
	// HTTP PUT.
	CorsHttpMethod_PUT CorsHttpMethod = "PUT"
)

