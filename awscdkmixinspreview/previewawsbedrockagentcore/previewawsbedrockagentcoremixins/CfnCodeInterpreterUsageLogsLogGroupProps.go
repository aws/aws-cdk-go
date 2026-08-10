package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterUsageLogsLogGroupProps := &CfnCodeInterpreterUsageLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCodeInterpreterUsageLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCodeInterpreterUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCodeInterpreterUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterUsageLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCodeInterpreterUsageLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

