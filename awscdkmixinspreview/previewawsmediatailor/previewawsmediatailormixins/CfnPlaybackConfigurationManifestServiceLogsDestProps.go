package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationManifestServiceLogsDestProps := &CfnPlaybackConfigurationManifestServiceLogsDestProps{
//   	RecordFields: []CfnPlaybackConfigurationManifestServiceLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnPlaybackConfigurationManifestServiceLogsRecordFields_CUSTOMERID,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationManifestServiceLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationManifestServiceLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

