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
type CfnDeploymentGroup_TrafficRouteProperty struct {
	// `CfnDeploymentGroup.TrafficRouteProperty.ListenerArns`.
	ListenerArns *[]*string `field:"optional" json:"listenerArns" yaml:"listenerArns"`
}

