package previewawsshieldmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProtectionFlowLogsFirehoseProps := &CfnProtectionFlowLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnProtectionFlowLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnProtectionFlowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnProtectionFlowLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnProtectionFlowLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnProtectionFlowLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnProtectionFlowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

