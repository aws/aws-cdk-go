package awscodedeploy


// The `LoadBalancerInfo` property type specifies information about the load balancer or target group used for an AWS CodeDeploy deployment group.
//
// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
//
// For AWS CloudFormation to use the properties specified in `LoadBalancerInfo` , the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` . If `DeploymentStyle.DeploymentOption` is not set to `WITH_TRAFFIC_CONTROL` , AWS CloudFormation ignores any settings specified in `LoadBalancerInfo` .
//
// > AWS CloudFormation supports blue/green deployments on the AWS Lambda compute platform only.
//
// `LoadBalancerInfo` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerInfoProperty := &loadBalancerInfoProperty{
//   	elbInfoList: []interface{}{
//   		&eLBInfoProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	targetGroupInfoList: []interface{}{
//   		&targetGroupInfoProperty{
//   			name: jsii.String("name"),
//   		},
//   	},
//   	targetGroupPairInfoList: []interface{}{
//   		&targetGroupPairInfoProperty{
//   			prodTrafficRoute: &trafficRouteProperty{
//   				listenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   			targetGroups: []interface{}{
//   				&targetGroupInfoProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			testTrafficRoute: &trafficRouteProperty{
//   				listenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_LoadBalancerInfoProperty struct {
	// An array that contains information about the load balancer to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing, load balancers are used with Classic Load Balancers.
	//
	// > Adding more than one load balancer to the array is not supported.
	ElbInfoList interface{} `field:"optional" json:"elbInfoList" yaml:"elbInfoList"`
	// An array that contains information about the target group to use for load balancing in a deployment.
	//
	// In Elastic Load Balancing , target groups are used with Application Load Balancers .
	//
	// > Adding more than one target group to the array is not supported.
	TargetGroupInfoList interface{} `field:"optional" json:"targetGroupInfoList" yaml:"targetGroupInfoList"`
	// The target group pair information.
	//
	// This is an array of `TargeGroupPairInfo` objects with a maximum size of one.
	TargetGroupPairInfoList interface{} `field:"optional" json:"targetGroupPairInfoList" yaml:"targetGroupPairInfoList"`
}

