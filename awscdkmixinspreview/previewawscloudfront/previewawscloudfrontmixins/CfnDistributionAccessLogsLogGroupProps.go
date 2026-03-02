package previewawscloudfrontmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionAccessLogsLogGroupProps := &CfnDistributionAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnDistributionAccessLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnDistributionAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnDistributionAccessLogsRecordFields_DATE,
//   	},
//   }
//
// Experimental.
type CfnDistributionAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnDistributionAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnDistributionAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

