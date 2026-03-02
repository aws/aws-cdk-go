package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerRuleSetApplicationLogsDestProps := &CfnMailManagerRuleSetApplicationLogsDestProps{
//   	RecordFields: []CfnMailManagerRuleSetApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_NAME,
//   	},
//   }
//
// Experimental.
type CfnMailManagerRuleSetApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerRuleSetApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

