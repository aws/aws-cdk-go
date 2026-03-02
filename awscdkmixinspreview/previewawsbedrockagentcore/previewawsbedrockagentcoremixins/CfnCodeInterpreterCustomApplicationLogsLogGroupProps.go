package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomApplicationLogsLogGroupProps := &CfnCodeInterpreterCustomApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCodeInterpreterCustomApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCodeInterpreterCustomApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCodeInterpreterCustomApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterCustomApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCodeInterpreterCustomApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterCustomApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

