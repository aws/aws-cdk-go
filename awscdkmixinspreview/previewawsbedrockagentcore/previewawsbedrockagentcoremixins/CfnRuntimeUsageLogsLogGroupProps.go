package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeUsageLogsLogGroupProps := &CfnRuntimeUsageLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRuntimeUsageLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnRuntimeUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRuntimeUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRuntimeUsageLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnRuntimeUsageLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRuntimeUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

