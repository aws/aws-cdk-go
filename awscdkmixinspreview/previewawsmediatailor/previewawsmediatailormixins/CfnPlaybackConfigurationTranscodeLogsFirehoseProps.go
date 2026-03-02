package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationTranscodeLogsFirehoseProps := &CfnPlaybackConfigurationTranscodeLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnPlaybackConfigurationTranscodeLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnPlaybackConfigurationTranscodeLogsRecordFields_EVENTTYPE,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationTranscodeLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnPlaybackConfigurationTranscodeLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationTranscodeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

