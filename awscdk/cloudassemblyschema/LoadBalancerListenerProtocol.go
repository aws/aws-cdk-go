package cloudassemblyschema


// The protocol for connections from clients to the load balancer.
type LoadBalancerListenerProtocol string

const (
	// HTTP protocol.
	LoadBalancerListenerProtocol_HTTP LoadBalancerListenerProtocol = "HTTP"
	// HTTPS protocol.
	LoadBalancerListenerProtocol_HTTPS LoadBalancerListenerProtocol = "HTTPS"
	// TCP protocol.
	LoadBalancerListenerProtocol_TCP LoadBalancerListenerProtocol = "TCP"
	// TLS protocol.
	LoadBalancerListenerProtocol_TLS LoadBalancerListenerProtocol = "TLS"
	// UDP protocol.
	LoadBalancerListenerProtocol_UDP LoadBalancerListenerProtocol = "UDP"
	// TCP and UDP protocol.
	LoadBalancerListenerProtocol_TCP_UDP LoadBalancerListenerProtocol = "TCP_UDP"
)

