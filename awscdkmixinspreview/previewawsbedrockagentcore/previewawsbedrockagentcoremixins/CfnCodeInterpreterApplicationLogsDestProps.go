package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterApplicationLogsDestProps := &CfnCodeInterpreterApplicationLogsDestProps{
//   	RecordFields: []CfnCodeInterpreterApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCodeInterpreterApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

