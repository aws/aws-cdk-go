package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerLogsLogGroupProps := &CfnClusterPcsSchedulerLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsSchedulerLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnClusterPcsSchedulerLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsSchedulerLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsSchedulerLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterPcsSchedulerLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsSchedulerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

