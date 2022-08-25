// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Type of the {@link CfnCodeDeployBlueGreenEcsAttributes.trafficRouting} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRouting := &cfnTrafficRouting{
//   	prodTrafficRoute: &cfnTrafficRoute{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   	targetGroups: []*string{
//   		jsii.String("targetGroups"),
//   	},
//   	testTrafficRoute: &cfnTrafficRoute{
//   		logicalId: jsii.String("logicalId"),
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnTrafficRouting struct {
	// The listener to be used by your load balancer to direct traffic to your target groups.
	ProdTrafficRoute *CfnTrafficRoute `field:"required" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// The logical IDs of the blue and green, respectively, AWS::ElasticLoadBalancingV2::TargetGroup target groups.
	TargetGroups *[]*string `field:"required" json:"targetGroups" yaml:"targetGroups"`
	// The listener to be used by your load balancer to direct traffic to your target groups.
	TestTrafficRoute *CfnTrafficRoute `field:"required" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

