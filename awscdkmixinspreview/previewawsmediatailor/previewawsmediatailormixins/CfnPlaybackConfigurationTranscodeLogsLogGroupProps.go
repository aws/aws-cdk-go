package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationTranscodeLogsLogGroupProps := &CfnPlaybackConfigurationTranscodeLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPlaybackConfigurationTranscodeLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnPlaybackConfigurationTranscodeLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnPlaybackConfigurationTranscodeLogsRecordFields_EVENTTYPE,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationTranscodeLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnPlaybackConfigurationTranscodeLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationTranscodeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

