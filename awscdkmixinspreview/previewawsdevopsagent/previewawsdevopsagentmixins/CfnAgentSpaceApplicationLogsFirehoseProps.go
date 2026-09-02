package previewawsdevopsagentmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentSpaceApplicationLogsFirehoseProps := &CfnAgentSpaceApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAgentSpaceApplicationLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnAgentSpaceApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAgentSpaceApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnAgentSpaceApplicationLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnAgentSpaceApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAgentSpaceApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

