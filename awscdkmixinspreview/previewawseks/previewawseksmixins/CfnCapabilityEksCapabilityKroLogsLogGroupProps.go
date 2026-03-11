package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityKroLogsLogGroupProps := &CfnCapabilityEksCapabilityKroLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityKroLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityKroLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityKroLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityKroLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityKroLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

