package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomApplicationLogsDestProps := &CfnCodeInterpreterCustomApplicationLogsDestProps{
//   	RecordFields: []CfnCodeInterpreterCustomApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCodeInterpreterCustomApplicationLogsRecordFields_ACCOUNT_ID,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterCustomApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterCustomApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

