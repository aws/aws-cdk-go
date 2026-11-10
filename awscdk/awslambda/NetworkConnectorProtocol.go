package awslambda


// The IP protocol the connector's network interfaces use.
type NetworkConnectorProtocol string

const (
	// IPv4 only.
	NetworkConnectorProtocol_IPV4 NetworkConnectorProtocol = "IPV4"
	// IPv4 and IPv6.
	NetworkConnectorProtocol_DUAL_STACK NetworkConnectorProtocol = "DUAL_STACK"
)

