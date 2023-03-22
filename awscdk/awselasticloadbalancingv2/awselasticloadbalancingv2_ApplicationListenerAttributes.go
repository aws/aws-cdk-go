package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties to reference an existing listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   applicationListenerAttributes := &applicationListenerAttributes{
//   	listenerArn: jsii.String("listenerArn"),
//
//   	// the properties below are optional
//   	defaultPort: jsii.Number(123),
//   	securityGroup: securityGroup,
//   	securityGroupAllowsAllOutbound: jsii.Boolean(false),
//   	securityGroupId: jsii.String("securityGroupId"),
//   }
//
// Experimental.
type ApplicationListenerAttributes struct {
	// ARN of the listener.
	// Experimental.
	ListenerArn *string `field:"required" json:"listenerArn" yaml:"listenerArn"`
	// The default port on which this listener is listening.
	// Experimental.
	DefaultPort *float64 `field:"optional" json:"defaultPort" yaml:"defaultPort"`
	// Security group of the load balancer this listener is associated with.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Whether the imported security group allows all outbound traffic or not when imported using `securityGroupId`.
	//
	// Unless set to `false`, no egress rules will be added to the security group.
	// Deprecated: use `securityGroup` instead.
	SecurityGroupAllowsAllOutbound *bool `field:"optional" json:"securityGroupAllowsAllOutbound" yaml:"securityGroupAllowsAllOutbound"`
	// Security group ID of the load balancer this listener is associated with.
	// Deprecated: use `securityGroup` instead.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
}

