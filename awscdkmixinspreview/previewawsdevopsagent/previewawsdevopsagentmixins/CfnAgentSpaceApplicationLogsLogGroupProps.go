package previewawsdevopsagentmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentSpaceApplicationLogsLogGroupProps := &CfnAgentSpaceApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentSpaceApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnAgentSpaceApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnAgentSpaceApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnAgentSpaceApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentSpaceApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

