package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantEventLogsFirehoseProps := &CfnAssistantEventLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnAssistantEventLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnAssistantEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnAssistantEventLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnAssistantEventLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnAssistantEventLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAssistantEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

