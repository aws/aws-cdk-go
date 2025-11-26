package previewawsbedrockmixins


// Contains the names of the fields to which to map information about the vector store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openSearchServerlessFieldMappingProperty := &OpenSearchServerlessFieldMappingProperty{
//   	MetadataField: jsii.String("metadataField"),
//   	TextField: jsii.String("textField"),
//   	VectorField: jsii.String("vectorField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessfieldmapping.html
//
type CfnKnowledgeBasePropsMixin_OpenSearchServerlessFieldMappingProperty struct {
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessfieldmapping.html#cfn-bedrock-knowledgebase-opensearchserverlessfieldmapping-metadatafield
	//
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	//
	// The text is split according to the chunking strategy you choose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessfieldmapping.html#cfn-bedrock-knowledgebase-opensearchserverlessfieldmapping-textfield
	//
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
	// The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchserverlessfieldmapping.html#cfn-bedrock-knowledgebase-opensearchserverlessfieldmapping-vectorfield
	//
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
}

