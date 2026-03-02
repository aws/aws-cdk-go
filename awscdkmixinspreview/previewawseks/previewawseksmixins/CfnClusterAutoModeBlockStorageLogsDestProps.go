package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeBlockStorageLogsDestProps := &CfnClusterAutoModeBlockStorageLogsDestProps{
//   	RecordFields: []CfnClusterAutoModeBlockStorageLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterAutoModeBlockStorageLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeBlockStorageLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeBlockStorageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

