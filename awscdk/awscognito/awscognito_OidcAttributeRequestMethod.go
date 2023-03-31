package awscognito


// The method to use to request attributes.
// Experimental.
type OidcAttributeRequestMethod string

const (
	// GET.
	// Experimental.
	OidcAttributeRequestMethod_GET OidcAttributeRequestMethod = "GET"
	// POST.
	// Experimental.
	OidcAttributeRequestMethod_POST OidcAttributeRequestMethod = "POST"
)

