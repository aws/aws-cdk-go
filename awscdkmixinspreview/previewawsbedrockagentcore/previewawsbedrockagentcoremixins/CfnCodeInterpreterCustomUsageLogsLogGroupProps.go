package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomUsageLogsLogGroupProps := &CfnCodeInterpreterCustomUsageLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCodeInterpreterCustomUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCodeInterpreterCustomUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterCustomUsageLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCodeInterpreterCustomUsageLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterCustomUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

