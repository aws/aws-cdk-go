package awsglobalaccelerator


// Client affinity gives you control over whether to always route each client to the same specific endpoint.
// See: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-listeners.html#about-listeners-client-affinity
//
// Experimental.
type ClientAffinity string

const (
	// Route traffic based on the 5-tuple `(source IP, source port, destination IP, destination port, protocol)`.
	// Experimental.
	ClientAffinity_NONE ClientAffinity = "NONE"
	// Route traffic based on the 2-tuple `(source IP, destination IP)`.
	//
	// The result is that multiple connections from the same client will be routed the same.
	// Experimental.
	ClientAffinity_SOURCE_IP ClientAffinity = "SOURCE_IP"
)

