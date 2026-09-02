package previewawscloudfrontmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionConnectionLogsFirehoseProps := &CfnDistributionConnectionLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnDistributionConnectionLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnDistributionConnectionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnDistributionConnectionLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnDistributionConnectionLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnDistributionConnectionLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnDistributionConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

