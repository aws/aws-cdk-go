package awselasticloadbalancingv2


// Information about how traffic will be distributed between multiple target groups in a forward rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupTupleProperty := &TargetGroupTupleProperty{
//   	TargetGroupArn: jsii.String("targetGroupArn"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-targetgrouptuple.html
//
type CfnListener_TargetGroupTupleProperty struct {
	// The Amazon Resource Name (ARN) of the target group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-targetgrouptuple.html#cfn-elasticloadbalancingv2-listener-targetgrouptuple-targetgrouparn
	//
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
	// The weight.
	//
	// The range is 0 to 999.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-targetgrouptuple.html#cfn-elasticloadbalancingv2-listener-targetgrouptuple-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

