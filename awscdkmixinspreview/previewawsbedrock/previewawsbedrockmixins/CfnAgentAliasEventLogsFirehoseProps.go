package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasEventLogsFirehoseProps := &CfnAgentAliasEventLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentAliasEventLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnAgentAliasEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentAliasEventLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnAgentAliasEventLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnAgentAliasEventLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentAliasEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

