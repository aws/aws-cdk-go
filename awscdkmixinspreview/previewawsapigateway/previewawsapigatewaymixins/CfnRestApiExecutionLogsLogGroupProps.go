package previewawsapigatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiExecutionLogsLogGroupProps := &CfnRestApiExecutionLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRestApiExecutionLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnRestApiExecutionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRestApiExecutionLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRestApiExecutionLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnRestApiExecutionLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRestApiExecutionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

