package previewawsapigatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiAccessLogsDestProps := &CfnRestApiAccessLogsDestProps{
//   	RecordFields: []CfnRestApiAccessLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnRestApiAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRestApiAccessLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRestApiAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

