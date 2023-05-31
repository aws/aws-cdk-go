package awsec2


// Port for client VPN.
// Experimental.
type VpnPort string

const (
	// HTTPS.
	// Experimental.
	VpnPort_HTTPS VpnPort = "HTTPS"
	// OpenVPN.
	// Experimental.
	VpnPort_OPENVPN VpnPort = "OPENVPN"
)

