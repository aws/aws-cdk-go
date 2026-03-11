package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusInfoLogsDestProps := &CfnEventBusInfoLogsDestProps{
//   	RecordFields: []CfnEventBusInfoLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnEventBusInfoLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnEventBusInfoLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusInfoLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

