package previewawsapigatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiExecutionLogsFirehoseProps := &CfnRestApiExecutionLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRestApiExecutionLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnRestApiExecutionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRestApiExecutionLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRestApiExecutionLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnRestApiExecutionLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRestApiExecutionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

