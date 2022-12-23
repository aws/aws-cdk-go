package awsec2


// The VPN connection type.
// Experimental.
type VpnConnectionType string

const (
	// The IPsec 1 VPN connection type.
	// Experimental.
	VpnConnectionType_IPSEC_1 VpnConnectionType = "IPSEC_1"
	// Dummy member TODO: remove once https://github.com/aws/jsii/issues/231 is fixed.
	// Experimental.
	VpnConnectionType_DUMMY VpnConnectionType = "DUMMY"
)

