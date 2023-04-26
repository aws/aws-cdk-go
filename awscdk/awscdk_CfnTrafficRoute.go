// An experiment to bundle the entire CDK into a single module
package awscdk


// A traffic route, representing where the traffic is being directed to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRoute := &CfnTrafficRoute{
//   	LogicalId: jsii.String("logicalId"),
//   	Type: jsii.String("type"),
//   }
//
// Experimental.
type CfnTrafficRoute struct {
	// The logical id of the target resource.
	// Experimental.
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The resource type of the route.
	//
	// Today, the only allowed value is 'AWS::ElasticLoadBalancingV2::Listener'.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

