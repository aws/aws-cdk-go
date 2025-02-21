package awsec2alpha


// Refers to two possible scope types under IPAM.
// Experimental.
type IpamScopeType string

const (
	// Default scopes created by IPAM.
	// Experimental.
	IpamScopeType_DEFAULT IpamScopeType = "DEFAULT"
	// Custom scope created using method.
	// Experimental.
	IpamScopeType_CUSTOM IpamScopeType = "CUSTOM"
)

