package previewawsbedrockmixins


// Specifies a field to be used during the reranking process in a Knowledge Base vector search.
//
// This structure identifies metadata fields that should be considered when reordering search results to improve relevance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fieldForRerankingProperty := &FieldForRerankingProperty{
//   	FieldName: jsii.String("fieldName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-fieldforreranking.html
//
type CfnFlowPropsMixin_FieldForRerankingProperty struct {
	// The name of the metadata field to be used during the reranking process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-fieldforreranking.html#cfn-bedrock-flow-fieldforreranking-fieldname
	//
	FieldName *string `field:"optional" json:"fieldName" yaml:"fieldName"`
}

