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
//   forwardConfigProperty := &forwardConfigProperty{
//   	targetGroups: []interface{}{
//   		&targetGroupTupleProperty{
//   			targetGroupArn: jsii.String("targetGroupArn"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	targetGroupStickinessConfig: &targetGroupStickinessConfigProperty{
//   		durationSeconds: jsii.Number(123),
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnListener_ForwardConfigProperty struct {
	// Information about how traffic will be distributed between multiple target groups in a forward rule.
	TargetGroups interface{} `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// Information about the target group stickiness for a rule.
	TargetGroupStickinessConfig interface{} `field:"optional" json:"targetGroupStickinessConfig" yaml:"targetGroupStickinessConfig"`
}

