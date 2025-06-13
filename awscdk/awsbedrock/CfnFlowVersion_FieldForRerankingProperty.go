package awsbedrock


// Contains information for a metadata field to include in or exclude from consideration when reranking.
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
	// The name of a metadata field to include in or exclude from consideration when reranking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-fieldforreranking.html#cfn-bedrock-flowversion-fieldforreranking-fieldname
	//
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
}

