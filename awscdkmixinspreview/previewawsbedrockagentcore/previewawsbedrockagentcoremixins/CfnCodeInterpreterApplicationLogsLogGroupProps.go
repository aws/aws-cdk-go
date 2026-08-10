package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterApplicationLogsLogGroupProps := &CfnCodeInterpreterApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCodeInterpreterApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCodeInterpreterApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCodeInterpreterApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCodeInterpreterApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

