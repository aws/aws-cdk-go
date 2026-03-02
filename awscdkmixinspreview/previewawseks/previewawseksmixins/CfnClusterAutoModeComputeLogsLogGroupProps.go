package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeComputeLogsLogGroupProps := &CfnClusterAutoModeComputeLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterAutoModeComputeLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnClusterAutoModeComputeLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterAutoModeComputeLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeComputeLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterAutoModeComputeLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeComputeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

