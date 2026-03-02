package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeApplicationLogsLogGroupProps := &CfnRuntimeApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRuntimeApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnRuntimeApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRuntimeApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnRuntimeApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnRuntimeApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRuntimeApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

