package previewawsstepfunctionsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineStandardLogsDestProps := &CfnStateMachineStandardLogsDestProps{
//   	RecordFields: []CfnStateMachineStandardLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnStateMachineStandardLogsRecordFields_DETAILS,
//   	},
//   }
//
// Experimental.
type CfnStateMachineStandardLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnStateMachineStandardLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

