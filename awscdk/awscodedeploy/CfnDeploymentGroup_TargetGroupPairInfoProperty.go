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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html
//
type CfnDeploymentGroup_TargetGroupPairInfoProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html#cfn-codedeploy-deploymentgroup-targetgrouppairinfo-prodtrafficroute
	//
	ProdTrafficRoute interface{} `field:"optional" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html#cfn-codedeploy-deploymentgroup-targetgrouppairinfo-targetgroups
	//
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-targetgrouppairinfo.html#cfn-codedeploy-deploymentgroup-targetgrouppairinfo-testtrafficroute
	//
	TestTrafficRoute interface{} `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

