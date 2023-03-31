package awselasticloadbalancingv2


// Information about a source IP condition.
//
// You can use this condition to route based on the IP address of the source that connects to the load balancer. If a client is behind a proxy, this is the IP address of the proxy not the IP address of the client.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceIpConfigProperty := &sourceIpConfigProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_SourceIpConfigProperty struct {
	// The source IP addresses, in CIDR format. You can use both IPv4 and IPv6 addresses. Wildcards are not supported.
	//
	// If you specify multiple addresses, the condition is satisfied if the source IP address of the request matches one of the CIDR blocks. This condition is not satisfied by the addresses in the X-Forwarded-For header.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

