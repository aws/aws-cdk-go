package awsglobalaccelerator


// The protocol for the connections from clients to the accelerator.
type HealthCheckProtocol string

const (
	// TCP.
	HealthCheckProtocol_TCP HealthCheckProtocol = "TCP"
	// HTTP.
	HealthCheckProtocol_HTTP HealthCheckProtocol = "HTTP"
	// HTTPS.
	HealthCheckProtocol_HTTPS HealthCheckProtocol = "HTTPS"
)

