package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusInfoLogsLogGroupProps := &CfnEventBusInfoLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnEventBusInfoLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnEventBusInfoLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnEventBusInfoLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnEventBusInfoLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnEventBusInfoLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusInfoLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

