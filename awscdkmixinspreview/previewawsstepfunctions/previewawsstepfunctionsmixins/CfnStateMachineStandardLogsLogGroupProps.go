package previewawsstepfunctionsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineStandardLogsLogGroupProps := &CfnStateMachineStandardLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnStateMachineStandardLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnStateMachineStandardLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnStateMachineStandardLogsRecordFields_DETAILS,
//   	},
//   }
//
// Experimental.
type CfnStateMachineStandardLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnStateMachineStandardLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnStateMachineStandardLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

