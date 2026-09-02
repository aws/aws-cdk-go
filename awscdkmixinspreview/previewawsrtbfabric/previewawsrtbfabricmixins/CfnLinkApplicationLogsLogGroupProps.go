package previewawsrtbfabricmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkApplicationLogsLogGroupProps := &CfnLinkApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLinkApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnLinkApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLinkApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnLinkApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnLinkApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLinkApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

