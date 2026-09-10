package previewawselementalinferencemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFeedApplicationLogsDestProps := &CfnFeedApplicationLogsDestProps{
//   	RecordFields: []CfnFeedApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnFeedApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFeedApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFeedApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

