package previewawsapsmixins


// This structure contains the limits that apply to time series that match one label set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   limitsPerLabelSetEntryProperty := &LimitsPerLabelSetEntryProperty{
//   	MaxSeries: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelsetentry.html
//
type CfnWorkspacePropsMixin_LimitsPerLabelSetEntryProperty struct {
	// The maximum number of active series that can be ingested that match this label set.
	//
	// Setting this to 0 causes no label set limit to be enforced, but it does cause Amazon Managed Service for Prometheus to vend label set metrics to CloudWatch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelsetentry.html#cfn-aps-workspace-limitsperlabelsetentry-maxseries
	//
	MaxSeries *float64 `field:"optional" json:"maxSeries" yaml:"maxSeries"`
}

