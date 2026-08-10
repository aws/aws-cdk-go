package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterApplicationLogsFirehoseProps := &CfnCodeInterpreterApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCodeInterpreterApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnCodeInterpreterApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCodeInterpreterApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnCodeInterpreterApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

