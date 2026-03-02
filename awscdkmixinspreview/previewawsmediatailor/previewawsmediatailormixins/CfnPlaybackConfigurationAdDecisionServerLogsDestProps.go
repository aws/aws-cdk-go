package previewawsmediatailormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackConfigurationAdDecisionServerLogsDestProps := &CfnPlaybackConfigurationAdDecisionServerLogsDestProps{
//   	RecordFields: []CfnPlaybackConfigurationAdDecisionServerLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnPlaybackConfigurationAdDecisionServerLogsRecordFields_ORIGINID,
//   	},
//   }
//
// Experimental.
type CfnPlaybackConfigurationAdDecisionServerLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnPlaybackConfigurationAdDecisionServerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

