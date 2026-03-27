package previewawsapigatewaymixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRestApiExecutionLogsDestProps := &CfnRestApiExecutionLogsDestProps{
//   	RecordFields: []CfnRestApiExecutionLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnRestApiExecutionLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnRestApiExecutionLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRestApiExecutionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

