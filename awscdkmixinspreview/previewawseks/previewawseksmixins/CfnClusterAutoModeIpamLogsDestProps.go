package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeIpamLogsDestProps := &CfnClusterAutoModeIpamLogsDestProps{
//   	RecordFields: []CfnClusterAutoModeIpamLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterAutoModeIpamLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeIpamLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeIpamLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

