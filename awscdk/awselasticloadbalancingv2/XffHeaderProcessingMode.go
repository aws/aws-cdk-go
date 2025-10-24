package awselasticloadbalancingv2


// Processing mode of the X-Forwarded-For header in the HTTP request before the Application Load Balancer sends the request to the target.
//
// Example:
//   var vpc Vpc
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
//   	// Whether HTTP headers with header fields that are not valid
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
type XffHeaderProcessingMode string

const (
	// Application Load Balancer adds the client IP address (of the last hop) to the X-Forwarded-For header in the HTTP request before it sends it to targets.
	XffHeaderProcessingMode_APPEND XffHeaderProcessingMode = "APPEND"
	// Application Load Balancer preserves the X-Forwarded-For header in the HTTP request, and sends it to targets without any change.
	XffHeaderProcessingMode_PRESERVE XffHeaderProcessingMode = "PRESERVE"
	// Application Load Balancer removes the X-Forwarded-For header in the HTTP request before it sends it to targets.
	XffHeaderProcessingMode_REMOVE XffHeaderProcessingMode = "REMOVE"
)

