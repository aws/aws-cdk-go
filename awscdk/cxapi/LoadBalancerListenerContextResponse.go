package cxapi


// Properties of a discovered load balancer listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerListenerContextResponse := &LoadBalancerListenerContextResponse{
//   	ListenerArn: jsii.String("listenerArn"),
//   	ListenerPort: jsii.Number(123),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type LoadBalancerListenerContextResponse struct {
	// The ARN of the listener.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The port the listener is listening on.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
	// The security groups of the load balancer.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
}

