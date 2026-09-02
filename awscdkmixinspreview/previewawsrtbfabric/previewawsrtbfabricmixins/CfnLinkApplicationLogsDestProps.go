package previewawsrtbfabricmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkApplicationLogsDestProps := &CfnLinkApplicationLogsDestProps{
//   	RecordFields: []CfnLinkApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnLinkApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnLinkApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLinkApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

