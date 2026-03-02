package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeApplicationLogsFirehoseProps := &CfnRuntimeApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRuntimeApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnRuntimeApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRuntimeApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnRuntimeApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnRuntimeApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRuntimeApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

