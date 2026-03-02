package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationAdDecisionServerLogsLogGroupProps := &CfnPlaybackConfigurationAdDecisionServerLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnPlaybackConfigurationAdDecisionServerLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ORIGINID,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationAdDecisionServerLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationAdDecisionServerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

