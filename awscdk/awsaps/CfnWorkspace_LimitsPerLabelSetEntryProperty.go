package awsaps


// Limits that can be applied to a label set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitsPerLabelSetEntryProperty := &LimitsPerLabelSetEntryProperty{
//   	MaxSeries: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelsetentry.html
//
type CfnWorkspace_LimitsPerLabelSetEntryProperty struct {
	// The maximum number of active series that can be ingested for this label set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelsetentry.html#cfn-aps-workspace-limitsperlabelsetentry-maxseries
	//
	MaxSeries *float64 `field:"optional" json:"maxSeries" yaml:"maxSeries"`
}

