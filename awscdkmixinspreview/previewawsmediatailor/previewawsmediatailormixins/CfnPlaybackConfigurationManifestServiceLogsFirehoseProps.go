package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationManifestServiceLogsFirehoseProps := &CfnPlaybackConfigurationManifestServiceLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPlaybackConfigurationManifestServiceLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnPlaybackConfigurationManifestServiceLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnPlaybackConfigurationManifestServiceLogsRecordFields_CUSTOMERID,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationManifestServiceLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnPlaybackConfigurationManifestServiceLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationManifestServiceLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

