package awsvpclattice


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   weightedTargetGroupProperty := &weightedTargetGroupProperty{
//   	targetGroupIdentifier: jsii.String("targetGroupIdentifier"),
//
//   	// the properties below are optional
//   	weight: jsii.Number(123),
//   }
//
type CfnListener_WeightedTargetGroupProperty struct {
	// `CfnListener.WeightedTargetGroupProperty.TargetGroupIdentifier`.
	TargetGroupIdentifier *string `field:"required" json:"targetGroupIdentifier" yaml:"targetGroupIdentifier"`
	// `CfnListener.WeightedTargetGroupProperty.Weight`.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

