package awselasticloadbalancingv2


// What to do with unauthenticated requests.
// Experimental.
type UnauthenticatedAction string

const (
	// Return an HTTP 401 Unauthorized error.
	// Experimental.
	UnauthenticatedAction_DENY UnauthenticatedAction = "DENY"
	// Allow the request to be forwarded to the target.
	// Experimental.
	UnauthenticatedAction_ALLOW UnauthenticatedAction = "ALLOW"
	// Redirect the request to the IdP authorization endpoint.
	// Experimental.
	UnauthenticatedAction_AUTHENTICATE UnauthenticatedAction = "AUTHENTICATE"
)

