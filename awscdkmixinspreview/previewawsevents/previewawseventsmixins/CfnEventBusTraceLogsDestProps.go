package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusTraceLogsDestProps := &CfnEventBusTraceLogsDestProps{
//   	RecordFields: []CfnEventBusTraceLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnEventBusTraceLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnEventBusTraceLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusTraceLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

