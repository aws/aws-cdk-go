package awsglobalaccelerator


// The protocol for the connections from clients to the accelerator.
// Experimental.
type ConnectionProtocol string

const (
	// TCP.
	// Experimental.
	ConnectionProtocol_TCP ConnectionProtocol = "TCP"
	// UDP.
	// Experimental.
	ConnectionProtocol_UDP ConnectionProtocol = "UDP"
)

