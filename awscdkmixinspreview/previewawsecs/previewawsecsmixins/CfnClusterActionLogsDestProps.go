package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterActionLogsDestProps := &CfnClusterActionLogsDestProps{
//   	RecordFields: []CfnClusterActionLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterActionLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterActionLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterActionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

