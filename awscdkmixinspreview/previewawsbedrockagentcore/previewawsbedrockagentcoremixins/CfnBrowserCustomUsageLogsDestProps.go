package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomUsageLogsDestProps := &CfnBrowserCustomUsageLogsDestProps{
//   	RecordFields: []CfnBrowserCustomUsageLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnBrowserCustomUsageLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnBrowserCustomUsageLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBrowserCustomUsageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

