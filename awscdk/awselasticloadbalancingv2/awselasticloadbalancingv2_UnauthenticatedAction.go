package awselasticloadbalancingv2


// What to do with unauthenticated requests.
type UnauthenticatedAction string

const (
	// Return an HTTP 401 Unauthorized error.
	UnauthenticatedAction_DENY UnauthenticatedAction = "DENY"
	// Allow the request to be forwarded to the target.
	UnauthenticatedAction_ALLOW UnauthenticatedAction = "ALLOW"
	// Redirect the request to the IdP authorization endpoint.
	UnauthenticatedAction_AUTHENTICATE UnauthenticatedAction = "AUTHENTICATE"
)

