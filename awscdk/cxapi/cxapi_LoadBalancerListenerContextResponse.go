package cxapi


// Properties of a discovered load balancer listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerListenerContextResponse := &loadBalancerListenerContextResponse{
//   	listenerArn: jsii.String("listenerArn"),
//   	listenerPort: jsii.Number(123),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
type LoadBalancerListenerContextResponse struct {
	// The ARN of the listener.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The port the listener is listening on.
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
	// The security groups of the load balancer.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
}

