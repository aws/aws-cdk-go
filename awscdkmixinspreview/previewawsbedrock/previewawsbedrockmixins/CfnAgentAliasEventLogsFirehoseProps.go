package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentAliasEventLogsFirehoseProps := &CfnAgentAliasEventLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentAliasEventLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnAgentAliasEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentAliasEventLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnAgentAliasEventLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnAgentAliasEventLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentAliasEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

