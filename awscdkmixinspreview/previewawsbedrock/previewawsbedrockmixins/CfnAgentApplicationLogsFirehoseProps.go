package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentApplicationLogsFirehoseProps := &CfnAgentApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnAgentApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentApplicationLogsRecordFields_AGENT_ALIAS_ARN,
//   	},
//   }
//
// Experimental.
type CfnAgentApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnAgentApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

