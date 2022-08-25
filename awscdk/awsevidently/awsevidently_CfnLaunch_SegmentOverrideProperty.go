package awsevidently


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   segmentOverrideProperty := &segmentOverrideProperty{
//   	evaluationOrder: jsii.Number(123),
//   	segment: jsii.String("segment"),
//   	weights: []interface{}{
//   		&groupToWeightProperty{
//   			groupName: jsii.String("groupName"),
//   			splitWeight: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnLaunch_SegmentOverrideProperty struct {
	// `CfnLaunch.SegmentOverrideProperty.EvaluationOrder`.
	EvaluationOrder *float64 `field:"required" json:"evaluationOrder" yaml:"evaluationOrder"`
	// `CfnLaunch.SegmentOverrideProperty.Segment`.
	Segment *string `field:"required" json:"segment" yaml:"segment"`
	// `CfnLaunch.SegmentOverrideProperty.Weights`.
	Weights interface{} `field:"required" json:"weights" yaml:"weights"`
}

