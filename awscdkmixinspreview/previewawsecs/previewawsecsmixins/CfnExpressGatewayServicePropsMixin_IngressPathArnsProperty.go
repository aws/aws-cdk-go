package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ingressPathArnsProperty := &IngressPathArnsProperty{
//   	CertificateArn: jsii.String("certificateArn"),
//   	ListenerArn: jsii.String("listenerArn"),
//   	ListenerRuleArn: jsii.String("listenerRuleArn"),
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	LoadBalancerSecurityGroups: []*string{
//   		jsii.String("loadBalancerSecurityGroups"),
//   	},
//   	TargetGroupArns: []*string{
//   		jsii.String("targetGroupArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspatharns.html
//
type CfnExpressGatewayServicePropsMixin_IngressPathArnsProperty struct {
	// The Certificate ARN associated with the express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspatharns.html#cfn-ecs-expressgatewayservice-ingresspatharns-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The ARN of the Load Balancer listener associated with the express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspatharns.html#cfn-ecs-expressgatewayservice-ingresspatharns-listenerarn
	//
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// The ARN of the Load Balancer listener rule associated with the express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspatharns.html#cfn-ecs-expressgatewayservice-ingresspatharns-listenerrulearn
	//
	ListenerRuleArn *string `field:"optional" json:"listenerRuleArn" yaml:"listenerRuleArn"`
	// The ARN of the Load Balancer associated with the express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspatharns.html#cfn-ecs-expressgatewayservice-ingresspatharns-loadbalancerarn
	//
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The list of Load Balancer Security Group ARNs associated with the express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspatharns.html#cfn-ecs-expressgatewayservice-ingresspatharns-loadbalancersecuritygroups
	//
	LoadBalancerSecurityGroups *[]*string `field:"optional" json:"loadBalancerSecurityGroups" yaml:"loadBalancerSecurityGroups"`
	// The list of Target Group ARNs associated with the express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ingresspatharns.html#cfn-ecs-expressgatewayservice-ingresspatharns-targetgrouparns
	//
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
}

