package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryApplicationLogsFirehoseProps := &CfnMemoryApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMemoryApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnMemoryApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMemoryApplicationLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnMemoryApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnMemoryApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMemoryApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

