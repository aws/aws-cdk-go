package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseTracesDestProps := &CfnKnowledgeBaseTracesDestProps{
//   	RecordFields: []CfnKnowledgeBaseTracesRecordFields{
//   		awscdkmixinspreview.Mixins.CfnKnowledgeBaseTracesRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnKnowledgeBaseTracesDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnKnowledgeBaseTracesRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

