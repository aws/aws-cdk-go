package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentApplicationLogsLogGroupProps := &CfnAgentApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnAgentApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentApplicationLogsRecordFields_AGENT_ALIAS_ARN,
//   	},
//   }
//
// Experimental.
type CfnAgentApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnAgentApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

