package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeApplicationLogsDestProps := &CfnRuntimeApplicationLogsDestProps{
//   	RecordFields: []CfnRuntimeApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnRuntimeApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnRuntimeApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRuntimeApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

