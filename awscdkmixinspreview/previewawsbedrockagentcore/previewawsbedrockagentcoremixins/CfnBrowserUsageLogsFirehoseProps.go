package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserUsageLogsFirehoseProps := &CfnBrowserUsageLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnBrowserUsageLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnBrowserUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnBrowserUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnBrowserUsageLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnBrowserUsageLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBrowserUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

