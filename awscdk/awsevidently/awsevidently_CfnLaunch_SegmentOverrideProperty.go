package awsevidently


// Use this structure to specify different traffic splits for one or more audience *segments* .
//
// A segment is a portion of your audience that share one or more characteristics. Examples could be Chrome browser users, users in Europe, or Firefox browser users in Europe who also fit other criteria that your application collects, such as age.
//
// For more information, see [Use segments to focus your audience](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html) .
//
// This sructure is an array of up to six segment override objects. Each of these objects specifies a segment that you have already created, and defines the traffic split for that segment.
//
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
	// A number indicating the order to use to evaluate segment overrides, if there are more than one.
	//
	// Segment overrides with lower numbers are evaluated first.
	EvaluationOrder *float64 `field:"required" json:"evaluationOrder" yaml:"evaluationOrder"`
	// The ARN of the segment to use for this override.
	Segment *string `field:"required" json:"segment" yaml:"segment"`
	// The traffic allocation percentages among the feature variations to assign to this segment.
	//
	// This is a set of key-value pairs. The keys are variation names. The values represent the amount of traffic to allocate to that variation for this segment. This is expressed in thousandths of a percent, so a weight of 50000 represents 50% of traffic.
	Weights interface{} `field:"required" json:"weights" yaml:"weights"`
}

