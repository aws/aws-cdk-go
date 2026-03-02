package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseRuntimeLogsDestProps := &CfnKnowledgeBaseRuntimeLogsDestProps{
//   	RecordFields: []CfnKnowledgeBaseRuntimeLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnKnowledgeBaseRuntimeLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnKnowledgeBaseRuntimeLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnKnowledgeBaseRuntimeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

