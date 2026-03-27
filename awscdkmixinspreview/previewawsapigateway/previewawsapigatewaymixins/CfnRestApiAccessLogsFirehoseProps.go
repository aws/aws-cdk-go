package previewawsapigatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiAccessLogsFirehoseProps := &CfnRestApiAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRestApiAccessLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnRestApiAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRestApiAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRestApiAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnRestApiAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRestApiAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

