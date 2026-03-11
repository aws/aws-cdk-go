package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps := &CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

