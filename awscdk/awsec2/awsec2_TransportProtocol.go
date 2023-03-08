package awsec2


// Transport protocol for client VPN.
// Experimental.
type TransportProtocol string

const (
	// Transmission Control Protocol (TCP).
	// Experimental.
	TransportProtocol_TCP TransportProtocol = "TCP"
	// User Datagram Protocol (UDP).
	// Experimental.
	TransportProtocol_UDP TransportProtocol = "UDP"
)

