package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantEventLogsDestProps := &CfnAssistantEventLogsDestProps{
//   	RecordFields: []CfnAssistantEventLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnAssistantEventLogsRecordFields_ASSISTANT_ID,
//   	},
//   }
//
// Experimental.
type CfnAssistantEventLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnAssistantEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

