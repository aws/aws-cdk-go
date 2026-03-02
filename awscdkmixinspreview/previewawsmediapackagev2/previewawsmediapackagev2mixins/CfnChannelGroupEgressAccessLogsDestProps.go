package previewawsmediapackagev2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupEgressAccessLogsDestProps := &CfnChannelGroupEgressAccessLogsDestProps{
//   	RecordFields: []CfnChannelGroupEgressAccessLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnChannelGroupEgressAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnChannelGroupEgressAccessLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnChannelGroupEgressAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

