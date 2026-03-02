package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomUsageLogsFirehoseProps := &CfnBrowserCustomUsageLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnBrowserCustomUsageLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnBrowserCustomUsageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnBrowserCustomUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnBrowserCustomUsageLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnBrowserCustomUsageLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBrowserCustomUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

