package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeUsageLogsDestProps := &CfnRuntimeUsageLogsDestProps{
//   	RecordFields: []CfnRuntimeUsageLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnRuntimeUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRuntimeUsageLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRuntimeUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

