package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeComputeLogsDestProps := &CfnClusterAutoModeComputeLogsDestProps{
//   	RecordFields: []CfnClusterAutoModeComputeLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterAutoModeComputeLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeComputeLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeComputeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

