package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusErrorLogsDestProps := &CfnEventBusErrorLogsDestProps{
//   	RecordFields: []CfnEventBusErrorLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnEventBusErrorLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnEventBusErrorLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusErrorLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

