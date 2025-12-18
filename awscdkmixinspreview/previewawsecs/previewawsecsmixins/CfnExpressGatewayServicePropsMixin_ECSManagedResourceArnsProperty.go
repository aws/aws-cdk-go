package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eCSManagedResourceArnsProperty := &ECSManagedResourceArnsProperty{
//   	AutoScaling: &AutoScalingArnsProperty{
//   		ApplicationAutoScalingPolicies: []*string{
//   			jsii.String("applicationAutoScalingPolicies"),
//   		},
//   		ScalableTarget: jsii.String("scalableTarget"),
//   	},
//   	IngressPath: &IngressPathArnsProperty{
//   		CertificateArn: jsii.String("certificateArn"),
//   		ListenerArn: jsii.String("listenerArn"),
//   		ListenerRuleArn: jsii.String("listenerRuleArn"),
//   		LoadBalancerArn: jsii.String("loadBalancerArn"),
//   		LoadBalancerSecurityGroups: []*string{
//   			jsii.String("loadBalancerSecurityGroups"),
//   		},
//   		TargetGroupArns: []*string{
//   			jsii.String("targetGroupArns"),
//   		},
//   	},
//   	LogGroups: []*string{
//   		jsii.String("logGroups"),
//   	},
//   	MetricAlarms: []*string{
//   		jsii.String("metricAlarms"),
//   	},
//   	ServiceSecurityGroups: []*string{
//   		jsii.String("serviceSecurityGroups"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ecsmanagedresourcearns.html
//
type CfnExpressGatewayServicePropsMixin_ECSManagedResourceArnsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ecsmanagedresourcearns.html#cfn-ecs-expressgatewayservice-ecsmanagedresourcearns-autoscaling
	//
	AutoScaling interface{} `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ecsmanagedresourcearns.html#cfn-ecs-expressgatewayservice-ecsmanagedresourcearns-ingresspath
	//
	IngressPath interface{} `field:"optional" json:"ingressPath" yaml:"ingressPath"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ecsmanagedresourcearns.html#cfn-ecs-expressgatewayservice-ecsmanagedresourcearns-loggroups
	//
	LogGroups *[]*string `field:"optional" json:"logGroups" yaml:"logGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ecsmanagedresourcearns.html#cfn-ecs-expressgatewayservice-ecsmanagedresourcearns-metricalarms
	//
	MetricAlarms *[]*string `field:"optional" json:"metricAlarms" yaml:"metricAlarms"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-ecsmanagedresourcearns.html#cfn-ecs-expressgatewayservice-ecsmanagedresourcearns-servicesecuritygroups
	//
	ServiceSecurityGroups *[]*string `field:"optional" json:"serviceSecurityGroups" yaml:"serviceSecurityGroups"`
}

