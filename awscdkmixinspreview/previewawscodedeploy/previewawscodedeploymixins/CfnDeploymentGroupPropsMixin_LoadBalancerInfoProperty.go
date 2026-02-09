package previewawscodedeploymixins


// The `LoadBalancerInfo` property type specifies information about the load balancer or target group used for an AWS CodeDeploy deployment group.
//
// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
//
// For CloudFormation to use the properties specified in `LoadBalancerInfo` , the `DeploymentStyle.DeploymentOption` property must be set to `WITH_TRAFFIC_CONTROL` . If `DeploymentStyle.DeploymentOption` is not set to `WITH_TRAFFIC_CONTROL` , CloudFormation ignores any settings specified in `LoadBalancerInfo` .
//
// > CloudFormation supports blue/green deployments on the AWS Lambda compute platform only.
//
// `LoadBalancerInfo` is a property of the [DeploymentGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   loadBalancerInfoProperty := &LoadBalancerInfoProperty{
//   	ElbInfoList: []interface{}{
//   		&ELBInfoProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TargetGroupInfoList: []interface{}{
//   		&TargetGroupInfoProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	TargetGroupPairInfoList: []interface{}{
//   		&TargetGroupPairInfoProperty{
//   			ProdTrafficRoute: &TrafficRouteProperty{
//   				ListenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   			TargetGroups: []interface{}{
//   				&TargetGroupInfoProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			TestTrafficRoute: &TrafficRouteProperty{
//   				ListenerArns: []*string{
//   					jsii.String("listenerArns"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html
//
type CfnDeploymentGroupPropsMixin_LoadBalancerInfoProperty struct {
	// An array that contains information about the load balancers to use for load balancing in a deployment.
	//
	// If you're using Classic Load Balancers, specify those load balancers in this array.
	//
	// > You can add up to 10 load balancers to the array. > If you're using Application Load Balancers or Network Load Balancers, use the `targetGroupInfoList` array instead of this one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-elbinfolist
	//
	ElbInfoList interface{} `field:"optional" json:"elbInfoList" yaml:"elbInfoList"`
	// An array that contains information about the target groups to use for load balancing in a deployment.
	//
	// If you're using Application Load Balancers and Network Load Balancers, specify their associated target groups in this array.
	//
	// > You can add up to 10 target groups to the array. > If you're using Classic Load Balancers, use the `elbInfoList` array instead of this one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-targetgroupinfolist
	//
	TargetGroupInfoList interface{} `field:"optional" json:"targetGroupInfoList" yaml:"targetGroupInfoList"`
	// The target group pair information.
	//
	// This is an array of `TargeGroupPairInfo` objects with a maximum size of one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-loadbalancerinfo.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo-targetgrouppairinfolist
	//
	TargetGroupPairInfoList interface{} `field:"optional" json:"targetGroupPairInfoList" yaml:"targetGroupPairInfoList"`
}

