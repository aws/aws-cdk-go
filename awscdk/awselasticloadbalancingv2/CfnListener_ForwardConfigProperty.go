package awselasticloadbalancingv2


// Information for creating an action that distributes requests among one or more target groups.
//
// For Network Load Balancers, you can specify a single target group. Specify only when `Type` is `forward` . If you specify both `ForwardConfig` and `TargetGroupArn` , you can specify only one target group using `ForwardConfig` and it must be the same target group specified in `TargetGroupArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forwardConfigProperty := &ForwardConfigProperty{
//   	TargetGroups: []interface{}{
//   		&TargetGroupTupleProperty{
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	TargetGroupStickinessConfig: &TargetGroupStickinessConfigProperty{
//   		DurationSeconds: jsii.Number(123),
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-forwardconfig.html
//
type CfnListener_ForwardConfigProperty struct {
	// Information about how traffic will be distributed between multiple target groups in a forward rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-forwardconfig.html#cfn-elasticloadbalancingv2-listener-forwardconfig-targetgroups
	//
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// Information about the target group stickiness for a rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-forwardconfig.html#cfn-elasticloadbalancingv2-listener-forwardconfig-targetgroupstickinessconfig
	//
	TargetGroupStickinessConfig interface{} `field:"optional" json:"targetGroupStickinessConfig" yaml:"targetGroupStickinessConfig"`
}

