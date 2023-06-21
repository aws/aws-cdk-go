package awscdk


// Type of the `CfnCodeDeployBlueGreenEcsAttributes.trafficRouting` property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficRouting := &CfnTrafficRouting{
//   	ProdTrafficRoute: &CfnTrafficRoute{
//   		LogicalId: jsii.String("logicalId"),
//   		Type: jsii.String("type"),
//   	},
//   	TargetGroups: []*string{
//   		jsii.String("targetGroups"),
//   	},
//   	TestTrafficRoute: &CfnTrafficRoute{
//   		LogicalId: jsii.String("logicalId"),
//   		Type: jsii.String("type"),
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

