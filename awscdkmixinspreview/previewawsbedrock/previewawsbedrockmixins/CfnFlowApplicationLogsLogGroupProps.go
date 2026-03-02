package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowApplicationLogsLogGroupProps := &CfnFlowApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFlowApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnFlowApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFlowApplicationLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFlowApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnFlowApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFlowApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

