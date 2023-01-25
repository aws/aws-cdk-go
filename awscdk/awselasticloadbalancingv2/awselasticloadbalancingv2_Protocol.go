package awselasticloadbalancingv2


// Backend protocol for network load balancers and health checks.
//
// Example:
//   listener := elbv2.networkListener.fromLookup(this, jsii.String("ALBListener"), &networkListenerLookupOptions{
//   	loadBalancerTags: map[string]*string{
//   		"Cluster": jsii.String("MyClusterName"),
//   	},
//   	listenerProtocol: elbv2.protocol_TCP,
//   	listenerPort: jsii.Number(12345),
//   })
//
// Experimental.
type Protocol string

const (
	// HTTP (ALB health checks and NLB health checks).
	// Experimental.
	Protocol_HTTP Protocol = "HTTP"
	// HTTPS (ALB health checks and NLB health checks).
	// Experimental.
	Protocol_HTTPS Protocol = "HTTPS"
	// TCP (NLB, NLB health checks).
	// Experimental.
	Protocol_TCP Protocol = "TCP"
	// TLS (NLB).
	// Experimental.
	Protocol_TLS Protocol = "TLS"
	// UDP (NLB).
	// Experimental.
	Protocol_UDP Protocol = "UDP"
	// Listen to both TCP and UDP on the same port (NLB).
	// Experimental.
	Protocol_TCP_UDP Protocol = "TCP_UDP"
)

