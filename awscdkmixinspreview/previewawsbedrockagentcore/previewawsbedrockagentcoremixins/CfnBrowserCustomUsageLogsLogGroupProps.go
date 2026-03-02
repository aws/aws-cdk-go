package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomUsageLogsLogGroupProps := &CfnBrowserCustomUsageLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnBrowserCustomUsageLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnBrowserCustomUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnBrowserCustomUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnBrowserCustomUsageLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnBrowserCustomUsageLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBrowserCustomUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

