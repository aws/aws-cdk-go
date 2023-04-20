package awscodedeploy


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupPairInfoProperty := &TargetGroupPairInfoProperty{
//   	ProdTrafficRoute: &TrafficRouteProperty{
//   		ListenerArns: []*string{
//   			jsii.String("listenerArns"),
//   		},
//   	},
//   	TargetGroups: []interface{}{
//   		&TargetGroupInfoProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TestTrafficRoute: &TrafficRouteProperty{
//   		ListenerArns: []*string{
//   			jsii.String("listenerArns"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_TargetGroupPairInfoProperty struct {
	// `CfnDeploymentGroup.TargetGroupPairInfoProperty.ProdTrafficRoute`.
	ProdTrafficRoute interface{} `field:"optional" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// `CfnDeploymentGroup.TargetGroupPairInfoProperty.TargetGroups`.
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// `CfnDeploymentGroup.TargetGroupPairInfoProperty.TestTrafficRoute`.
	TestTrafficRoute interface{} `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

