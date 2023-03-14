package awsglobalaccelerator


// The protocol for the connections from clients to the accelerator.
// Experimental.
type HealthCheckProtocol string

const (
	// TCP.
	// Experimental.
	HealthCheckProtocol_TCP HealthCheckProtocol = "TCP"
	// HTTP.
	// Experimental.
	HealthCheckProtocol_HTTP HealthCheckProtocol = "HTTP"
	// HTTPS.
	// Experimental.
	HealthCheckProtocol_HTTPS HealthCheckProtocol = "HTTPS"
)

