package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   weightedTargetGroupProperty := &WeightedTargetGroupProperty{
//   	TargetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   	// the properties below are optional
//   	Weight: jsii.Number(123),
//   }
//
type CfnRule_WeightedTargetGroupProperty struct {
	// `CfnRule.WeightedTargetGroupProperty.TargetGroupIdentifier`.
	TargetGroupIdentifier *string `field:"required" json:"targetGroupIdentifier" yaml:"targetGroupIdentifier"`
	// `CfnRule.WeightedTargetGroupProperty.Weight`.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

