package mixinsawscodedeploy


// Information about a listener.
//
// The listener contains the path used to route traffic that is received from the load balancer to a target group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   trafficRouteProperty := &TrafficRouteProperty{
//   	ListenerArns: []*string{
//   		jsii.String("listenerArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-trafficroute.html
//
type CfnDeploymentGroupPropsMixin_TrafficRouteProperty struct {
	// The Amazon Resource Name (ARN) of one listener.
	//
	// The listener identifies the route between a target group and a load balancer. This is an array of strings with a maximum size of one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-trafficroute.html#cfn-codedeploy-deploymentgroup-trafficroute-listenerarns
	//
	ListenerArns *[]*string `field:"optional" json:"listenerArns" yaml:"listenerArns"`
}

