package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomApplicationLogsFirehoseProps := &CfnCodeInterpreterCustomApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCodeInterpreterCustomApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnCodeInterpreterCustomApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCodeInterpreterCustomApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterCustomApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnCodeInterpreterCustomApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterCustomApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

