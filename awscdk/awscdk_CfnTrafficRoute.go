// Version 2 of the AWS Cloud Development Kit library
package awscdk


// A traffic route, representing where the traffic is being directed to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoute := &cfnTrafficRoute{
//   	logicalId: jsii.String("logicalId"),
//   	type: jsii.String("type"),
//   }
//
type CfnTrafficRoute struct {
	// The logical id of the target resource.
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The resource type of the route.
	//
	// Today, the only allowed value is 'AWS::ElasticLoadBalancingV2::Listener'.
	Type *string `field:"required" json:"type" yaml:"type"`
}

