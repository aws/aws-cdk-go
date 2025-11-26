package previewawsbedrockmixins


// Contains the names of the fields to which to map information about the vector store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   openSearchManagedClusterFieldMappingProperty := &OpenSearchManagedClusterFieldMappingProperty{
//   	MetadataField: jsii.String("metadataField"),
//   	TextField: jsii.String("textField"),
//   	VectorField: jsii.String("vectorField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping.html
//
type CfnKnowledgeBasePropsMixin_OpenSearchManagedClusterFieldMappingProperty struct {
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping.html#cfn-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping-metadatafield
	//
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	//
	// The text is split according to the chunking strategy you choose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping.html#cfn-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping-textfield
	//
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
	// The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping.html#cfn-bedrock-knowledgebase-opensearchmanagedclusterfieldmapping-vectorfield
	//
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
}

