package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowApplicationLogsDestProps := &CfnFlowApplicationLogsDestProps{
//   	RecordFields: []CfnFlowApplicationLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnFlowApplicationLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFlowApplicationLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFlowApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

