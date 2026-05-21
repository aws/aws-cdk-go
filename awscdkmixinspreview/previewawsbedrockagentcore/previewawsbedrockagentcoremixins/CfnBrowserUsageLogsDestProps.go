package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserUsageLogsDestProps := &CfnBrowserUsageLogsDestProps{
//   	RecordFields: []CfnBrowserUsageLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnBrowserUsageLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnBrowserUsageLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBrowserUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

