package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerRuleSetApplicationLogsFirehoseProps := &CfnMailManagerRuleSetApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnMailManagerRuleSetApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_NAME,
//   	},
//   }
//
// Experimental.
type CfnMailManagerRuleSetApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerRuleSetApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

