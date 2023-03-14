package awselasticloadbalancing


// Specifies the connection draining settings for your Classic Load Balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionDrainingPolicyProperty := &connectionDrainingPolicyProperty{
//   	enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	timeout: jsii.Number(123),
//   }
//
type CfnLoadBalancer_ConnectionDrainingPolicyProperty struct {
	// Specifies whether connection draining is enabled for the load balancer.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The maximum time, in seconds, to keep the existing connections open before deregistering the instances.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

