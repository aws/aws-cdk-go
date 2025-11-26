package previewawsapsmixins


// This defines a label set for the workspace, and defines the ingestion limit for active time series that match that label set.
//
// Each label name in a label set must be unique.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   limitsPerLabelSetProperty := &LimitsPerLabelSetProperty{
//   	LabelSet: []interface{}{
//   		&LabelProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Limits: &LimitsPerLabelSetEntryProperty{
//   		MaxSeries: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelset.html
//
type CfnWorkspacePropsMixin_LimitsPerLabelSetProperty struct {
	// This defines one label set that will have an enforced ingestion limit.
	//
	// You can set ingestion limits on time series that match defined label sets, to help prevent a workspace from being overwhelmed with unexpected spikes in time series ingestion.
	//
	// Label values accept all UTF-8 characters with one exception. If the label name is metric name label `__ *name* __` , then the *metric* part of the name must conform to the following pattern: `[a-zA-Z_:][a-zA-Z0-9_:]*`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelset.html#cfn-aps-workspace-limitsperlabelset-labelset
	//
	LabelSet interface{} `field:"optional" json:"labelSet" yaml:"labelSet"`
	// This structure contains the information about the limits that apply to time series that match this label set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelset.html#cfn-aps-workspace-limitsperlabelset-limits
	//
	Limits interface{} `field:"optional" json:"limits" yaml:"limits"`
}

