package awselasticloadbalancingv2


// Backend protocol for network load balancers and health checks.
//
// Example:
//   var vpc vpc
//
//
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
//   	Vpc: Vpc,
//   	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
//   	EnablePrefixForIpv6SourceNat: jsii.Boolean(true),
//   })
//
//   listener := lb.AddListener(jsii.String("Listener"), &BaseNetworkListenerProps{
//   	Port: jsii.Number(1229),
//   	Protocol: elbv2.Protocol_UDP,
//   })
//
type Protocol string

const (
	// HTTP (ALB health checks and NLB health checks).
	Protocol_HTTP Protocol = "HTTP"
	// HTTPS (ALB health checks and NLB health checks).
	Protocol_HTTPS Protocol = "HTTPS"
	// TCP (NLB, NLB health checks).
	Protocol_TCP Protocol = "TCP"
	// TLS (NLB).
	Protocol_TLS Protocol = "TLS"
	// UDP (NLB).
	Protocol_UDP Protocol = "UDP"
	// Listen to both TCP and UDP on the same port (NLB).
	Protocol_TCP_UDP Protocol = "TCP_UDP"
)

