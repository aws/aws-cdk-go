package awselasticloadbalancingv2


// How the load balancer handles requests that might pose a security risk to your application.
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
//   })
//
// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#desync-mitigation-mode
//
type DesyncMitigationMode string

const (
	// Allows all traffic.
	DesyncMitigationMode_MONITOR DesyncMitigationMode = "MONITOR"
	// Provides durable mitigation against HTTP desync while maintaining the availability of your application.
	DesyncMitigationMode_DEFENSIVE DesyncMitigationMode = "DEFENSIVE"
	// Receives only requests that comply with RFC 7230.
	DesyncMitigationMode_STRICTEST DesyncMitigationMode = "STRICTEST"
)

