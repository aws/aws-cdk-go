package awsec2


// Port for client VPN.
type VpnPort string

const (
	// HTTPS.
	VpnPort_HTTPS VpnPort = "HTTPS"
	// OpenVPN.
	VpnPort_OPENVPN VpnPort = "OPENVPN"
)

