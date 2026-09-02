package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerRuleSetApplicationLogsFirehoseProps := &CfnMailManagerRuleSetApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnMailManagerRuleSetApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMailManagerRuleSetApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnMailManagerRuleSetApplicationLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerRuleSetApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

