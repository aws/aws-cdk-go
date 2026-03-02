package previewawsstepfunctionsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineExpressLogsDestProps := &CfnStateMachineExpressLogsDestProps{
//   	RecordFields: []CfnStateMachineExpressLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnStateMachineExpressLogsRecordFields_DETAILS,
//   	},
//   }
//
// Experimental.
type CfnStateMachineExpressLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnStateMachineExpressLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

