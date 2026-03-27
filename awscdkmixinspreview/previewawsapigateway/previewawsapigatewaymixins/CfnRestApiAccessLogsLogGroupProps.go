package previewawsapigatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiAccessLogsLogGroupProps := &CfnRestApiAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRestApiAccessLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnRestApiAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRestApiAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRestApiAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnRestApiAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRestApiAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

