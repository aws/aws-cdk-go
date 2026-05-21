package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserUsageLogsLogGroupProps := &CfnBrowserUsageLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnBrowserUsageLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnBrowserUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnBrowserUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnBrowserUsageLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnBrowserUsageLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBrowserUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

