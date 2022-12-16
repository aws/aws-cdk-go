package awselasticloadbalancingv2


// How the load balancer handles requests that might pose a security risk to your application.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &applicationLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//
//   	// Whether HTTP/2 is enabled
//   	http2Enabled: jsii.Boolean(false),
//
//   	// The idle timeout value, in seconds
//   	idleTimeout: cdk.duration_Seconds(jsii.Number(1000)),
//
//   	// Whether HTTP headers with header fields thatare not valid
//   	// are removed by the load balancer (true), or routed to targets
//   	dropInvalidHeaderFields: jsii.Boolean(true),
//
//   	// How the load balancer handles requests that might
//   	// pose a security risk to your application
//   	desyncMitigationMode: elbv2.desyncMitigationMode_DEFENSIVE,
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

