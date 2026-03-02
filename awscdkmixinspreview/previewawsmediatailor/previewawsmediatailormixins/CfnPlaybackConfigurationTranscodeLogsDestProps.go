package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationTranscodeLogsDestProps := &CfnPlaybackConfigurationTranscodeLogsDestProps{
//   	RecordFields: []CfnPlaybackConfigurationTranscodeLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnPlaybackConfigurationTranscodeLogsRecordFields_EVENTTYPE,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationTranscodeLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationTranscodeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

