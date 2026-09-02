package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasApplicationLogsFirehoseProps := &CfnAgentAliasApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentAliasApplicationLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnAgentAliasApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentAliasApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnAgentAliasApplicationLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnAgentAliasApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentAliasApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

