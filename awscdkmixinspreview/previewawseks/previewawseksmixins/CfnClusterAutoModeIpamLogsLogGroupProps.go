package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeIpamLogsLogGroupProps := &CfnClusterAutoModeIpamLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterAutoModeIpamLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnClusterAutoModeIpamLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterAutoModeIpamLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeIpamLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterAutoModeIpamLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeIpamLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

