package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomUsageLogsFirehoseProps := &CfnCodeInterpreterCustomUsageLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnCodeInterpreterCustomUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCodeInterpreterCustomUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterCustomUsageLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnCodeInterpreterCustomUsageLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterCustomUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

