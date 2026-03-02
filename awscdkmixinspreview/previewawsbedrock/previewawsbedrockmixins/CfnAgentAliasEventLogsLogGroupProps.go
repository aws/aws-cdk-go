package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasEventLogsLogGroupProps := &CfnAgentAliasEventLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentAliasEventLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnAgentAliasEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentAliasEventLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnAgentAliasEventLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnAgentAliasEventLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentAliasEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

