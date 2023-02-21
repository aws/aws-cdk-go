package awsec2


// Transport protocol for client VPN.
type TransportProtocol string

const (
	// Transmission Control Protocol (TCP).
	TransportProtocol_TCP TransportProtocol = "TCP"
	// User Datagram Protocol (UDP).
	TransportProtocol_UDP TransportProtocol = "UDP"
)

