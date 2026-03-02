package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationManifestServiceLogsLogGroupProps := &CfnPlaybackConfigurationManifestServiceLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnPlaybackConfigurationManifestServiceLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnPlaybackConfigurationManifestServiceLogsRecordFields_CUSTOMERID,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationManifestServiceLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnPlaybackConfigurationManifestServiceLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationManifestServiceLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

