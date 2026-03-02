package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryApplicationLogsDestProps := &CfnMemoryApplicationLogsDestProps{
//   	RecordFields: []CfnMemoryApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnMemoryApplicationLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnMemoryApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMemoryApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

