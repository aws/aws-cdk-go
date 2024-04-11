package awselasticloadbalancingv2


// What kind of addresses to allocate to the load balancer.
//
// Example:
//   var vpc vpc
//
//
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//
//   	// Whether HTTP/2 is enabled
//   	Http2Enabled: jsii.Boolean(false),
//
//   	// The idle timeout value, in seconds
//   	IdleTimeout: awscdk.Duration_Seconds(jsii.Number(1000)),
//
//   	// Whether HTTP headers with header fields thatare not valid
//   	// are removed by the load balancer (true), or routed to targets
//   	DropInvalidHeaderFields: jsii.Boolean(true),
//
//   	// How the load balancer handles requests that might
//   	// pose a security risk to your application
//   	DesyncMitigationMode: elbv2.DesyncMitigationMode_DEFENSIVE,
//
//   	// The type of IP addresses to use.
//   	IpAddressType: elbv2.IpAddressType_IPV4,
//
//   	// The duration of client keep-alive connections
//   	ClientKeepAlive: awscdk.Duration_*Seconds(jsii.Number(500)),
//
//   	// Whether cross-zone load balancing is enabled.
//   	CrossZoneEnabled: jsii.Boolean(true),
//
//   	// Whether the load balancer blocks traffic through the Internet Gateway (IGW).
//   	DenyAllIgwTraffic: jsii.Boolean(false),
//
//   	// Whether to preserve host header in the request to the target
//   	PreserveHostHeader: jsii.Boolean(true),
//
//   	// Whether to add the TLS information header to the request
//   	XAmznTlsVersionAndCipherSuiteHeaders: jsii.Boolean(true),
//
//   	// Whether the X-Forwarded-For header should preserve the source port
//   	PreserveXffClientPort: jsii.Boolean(true),
//
//   	// The processing mode for X-Forwarded-For headers
//   	XffHeaderProcessingMode: elbv2.XffHeaderProcessingMode_APPEND,
//
//   	// Whether to allow a load balancer to route requests to targets if it is unable to forward the request to AWS WAF.
//   	WafFailOpen: jsii.Boolean(true),
//   })
//
type IpAddressType string

const (
	// Allocate IPv4 addresses.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// Allocate both IPv4 and IPv6 addresses.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

