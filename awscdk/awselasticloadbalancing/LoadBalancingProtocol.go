package awselasticloadbalancing


type LoadBalancingProtocol string

const (
	LoadBalancingProtocol_TCP LoadBalancingProtocol = "TCP"
	LoadBalancingProtocol_SSL LoadBalancingProtocol = "SSL"
	LoadBalancingProtocol_HTTP LoadBalancingProtocol = "HTTP"
	LoadBalancingProtocol_HTTPS LoadBalancingProtocol = "HTTPS"
)

