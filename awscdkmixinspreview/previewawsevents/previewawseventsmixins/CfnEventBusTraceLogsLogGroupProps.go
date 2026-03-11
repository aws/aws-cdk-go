package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusTraceLogsLogGroupProps := &CfnEventBusTraceLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnEventBusTraceLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnEventBusTraceLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnEventBusTraceLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnEventBusTraceLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnEventBusTraceLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusTraceLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

