package previewawsstepfunctionsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineExpressLogsLogGroupProps := &CfnStateMachineExpressLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnStateMachineExpressLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnStateMachineExpressLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnStateMachineExpressLogsRecordFields_DETAILS,
//   	},
//   }
//
// Experimental.
type CfnStateMachineExpressLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnStateMachineExpressLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnStateMachineExpressLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

