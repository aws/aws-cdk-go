package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCodeInterpreterCustomUsageLogsDestProps := &CfnCodeInterpreterCustomUsageLogsDestProps{
//   	RecordFields: []CfnCodeInterpreterCustomUsageLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCodeInterpreterCustomUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnCodeInterpreterCustomUsageLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCodeInterpreterCustomUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

