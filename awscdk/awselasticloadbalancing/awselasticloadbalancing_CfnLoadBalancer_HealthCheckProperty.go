package awselasticloadbalancing


// Specifies health check settings for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckProperty := &healthCheckProperty{
//   	healthyThreshold: jsii.String("healthyThreshold"),
//   	interval: jsii.String("interval"),
//   	target: jsii.String("target"),
//   	timeout: jsii.String("timeout"),
//   	unhealthyThreshold: jsii.String("unhealthyThreshold"),
//   }
//
type CfnLoadBalancer_HealthCheckProperty struct {
	// The number of consecutive health checks successes required before moving the instance to the `Healthy` state.
	HealthyThreshold *string `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// The approximate interval, in seconds, between health checks of an individual instance.
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The instance being checked.
	//
	// The protocol is either TCP, HTTP, HTTPS, or SSL. The range of valid ports is one (1) through 65535.
	//
	// TCP is the default, specified as a TCP: port pair, for example "TCP:5000". In this case, a health check simply attempts to open a TCP connection to the instance on the specified port. Failure to connect within the configured timeout is considered unhealthy.
	//
	// SSL is also specified as SSL: port pair, for example, SSL:5000.
	//
	// For HTTP/HTTPS, you must include a ping path in the string. HTTP is specified as a HTTP:port;/;PathToPing; grouping, for example "HTTP:80/weather/us/wa/seattle". In this case, a HTTP GET request is issued to the instance on the given port and path. Any answer other than "200 OK" within the timeout period is considered unhealthy.
	//
	// The total length of the HTTP ping target must be 1024 16-bit Unicode characters or less.
	Target *string `field:"required" json:"target" yaml:"target"`
	// The amount of time, in seconds, during which no response means a failed health check.
	//
	// This value must be less than the `Interval` value.
	Timeout *string `field:"required" json:"timeout" yaml:"timeout"`
	// The number of consecutive health check failures required before moving the instance to the `Unhealthy` state.
	UnhealthyThreshold *string `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
}

