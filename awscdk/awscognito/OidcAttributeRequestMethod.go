package awscognito


// The method to use to request attributes.
type OidcAttributeRequestMethod string

const (
	// GET.
	OidcAttributeRequestMethod_GET OidcAttributeRequestMethod = "GET"
	// POST.
	OidcAttributeRequestMethod_POST OidcAttributeRequestMethod = "POST"
)

