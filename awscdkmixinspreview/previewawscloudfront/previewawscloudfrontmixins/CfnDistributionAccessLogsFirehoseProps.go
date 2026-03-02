package previewawscloudfrontmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionAccessLogsFirehoseProps := &CfnDistributionAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnDistributionAccessLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnDistributionAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnDistributionAccessLogsRecordFields_DATE,
//   	},
//   }
//
// Experimental.
type CfnDistributionAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnDistributionAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnDistributionAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

