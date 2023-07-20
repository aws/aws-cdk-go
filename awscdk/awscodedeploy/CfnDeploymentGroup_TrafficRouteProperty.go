package awscodedeploy


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trafficRouteProperty := &TrafficRouteProperty{
//   	ListenerArns: []*string{
//   		jsii.String("listenerArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-trafficroute.html
//
type CfnDeploymentGroup_TrafficRouteProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-trafficroute.html#cfn-codedeploy-deploymentgroup-trafficroute-listenerarns
	//
	ListenerArns *[]*string `field:"optional" json:"listenerArns" yaml:"listenerArns"`
}

