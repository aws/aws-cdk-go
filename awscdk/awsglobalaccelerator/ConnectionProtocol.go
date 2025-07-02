package awsglobalaccelerator


// The protocol for the connections from clients to the accelerator.
type ConnectionProtocol string

const (
	// TCP.
	ConnectionProtocol_TCP ConnectionProtocol = "TCP"
	// UDP.
	ConnectionProtocol_UDP ConnectionProtocol = "UDP"
)

