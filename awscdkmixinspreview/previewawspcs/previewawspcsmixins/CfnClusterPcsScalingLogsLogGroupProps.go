package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsScalingLogsLogGroupProps := &CfnClusterPcsScalingLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsScalingLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnClusterPcsScalingLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsScalingLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsScalingLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterPcsScalingLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsScalingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

