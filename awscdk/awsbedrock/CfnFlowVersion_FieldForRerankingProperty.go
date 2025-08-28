package awsbedrock


// Specifies a field to be used during the reranking process in a Knowledge Base vector search.
//
// This structure identifies metadata fields that should be considered when reordering search results to improve relevance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldForRerankingProperty := &FieldForRerankingProperty{
//   	FieldName: jsii.String("fieldName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-fieldforreranking.html
//
type CfnFlowVersion_FieldForRerankingProperty struct {
	// The name of the metadata field to be used during the reranking process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-fieldforreranking.html#cfn-bedrock-flowversion-fieldforreranking-fieldname
	//
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
}

