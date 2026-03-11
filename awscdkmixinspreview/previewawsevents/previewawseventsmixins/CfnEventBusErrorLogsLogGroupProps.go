package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusErrorLogsLogGroupProps := &CfnEventBusErrorLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnEventBusErrorLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnEventBusErrorLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnEventBusErrorLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnEventBusErrorLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnEventBusErrorLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusErrorLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

