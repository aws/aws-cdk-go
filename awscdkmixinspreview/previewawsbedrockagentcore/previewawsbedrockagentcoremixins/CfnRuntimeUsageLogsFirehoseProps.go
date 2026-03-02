package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeUsageLogsFirehoseProps := &CfnRuntimeUsageLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRuntimeUsageLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnRuntimeUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRuntimeUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRuntimeUsageLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnRuntimeUsageLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRuntimeUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

