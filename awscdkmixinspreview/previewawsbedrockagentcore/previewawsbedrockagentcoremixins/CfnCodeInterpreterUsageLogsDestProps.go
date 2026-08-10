package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterUsageLogsDestProps := &CfnCodeInterpreterUsageLogsDestProps{
//   	RecordFields: []CfnCodeInterpreterUsageLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCodeInterpreterUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterUsageLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

