package awselasticloadbalancingv2


// Information about how traffic will be distributed between multiple target groups in a forward rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetGroupTupleProperty := &targetGroupTupleProperty{
//   	targetGroupArn: jsii.String("targetGroupArn"),
//   	weight: jsii.Number(123),
//   }
//
type CfnListenerRule_TargetGroupTupleProperty struct {
	// The Amazon Resource Name (ARN) of the target group.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
	// The weight.
	//
	// The range is 0 to 999.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

