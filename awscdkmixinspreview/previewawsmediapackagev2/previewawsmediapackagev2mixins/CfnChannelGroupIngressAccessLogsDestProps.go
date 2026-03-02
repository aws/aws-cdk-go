package previewawsmediapackagev2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupIngressAccessLogsDestProps := &CfnChannelGroupIngressAccessLogsDestProps{
//   	RecordFields: []CfnChannelGroupIngressAccessLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnChannelGroupIngressAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnChannelGroupIngressAccessLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnChannelGroupIngressAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

