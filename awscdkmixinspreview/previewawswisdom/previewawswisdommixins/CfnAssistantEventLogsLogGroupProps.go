package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantEventLogsLogGroupProps := &CfnAssistantEventLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAssistantEventLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnAssistantEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAssistantEventLogsRecordFields_ASSISTANT_ID,
//   	},
//   }
//
// Experimental.
type CfnAssistantEventLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnAssistantEventLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAssistantEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

