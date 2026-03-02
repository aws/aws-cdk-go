package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationAdDecisionServerLogsFirehoseProps := &CfnPlaybackConfigurationAdDecisionServerLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnPlaybackConfigurationAdDecisionServerLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ORIGINID,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationAdDecisionServerLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnPlaybackConfigurationAdDecisionServerLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationAdDecisionServerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

