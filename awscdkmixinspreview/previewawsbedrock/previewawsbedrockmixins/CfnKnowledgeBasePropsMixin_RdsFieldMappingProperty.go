package previewawsbedrockmixins


// Contains the names of the fields to which to map information about the vector store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rdsFieldMappingProperty := &RdsFieldMappingProperty{
//   	CustomMetadataField: jsii.String("customMetadataField"),
//   	MetadataField: jsii.String("metadataField"),
//   	PrimaryKeyField: jsii.String("primaryKeyField"),
//   	TextField: jsii.String("textField"),
//   	VectorField: jsii.String("vectorField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsfieldmapping.html
//
type CfnKnowledgeBasePropsMixin_RdsFieldMappingProperty struct {
	// Provide a name for the universal metadata field where Amazon Bedrock will store any custom metadata from your data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsfieldmapping.html#cfn-bedrock-knowledgebase-rdsfieldmapping-custommetadatafield
	//
	CustomMetadataField *string `field:"optional" json:"customMetadataField" yaml:"customMetadataField"`
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsfieldmapping.html#cfn-bedrock-knowledgebase-rdsfieldmapping-metadatafield
	//
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the ID for each entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsfieldmapping.html#cfn-bedrock-knowledgebase-rdsfieldmapping-primarykeyfield
	//
	PrimaryKeyField *string `field:"optional" json:"primaryKeyField" yaml:"primaryKeyField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	//
	// The text is split according to the chunking strategy you choose.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsfieldmapping.html#cfn-bedrock-knowledgebase-rdsfieldmapping-textfield
	//
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
	// The name of the field in which Amazon Bedrock stores the vector embeddings for your data sources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsfieldmapping.html#cfn-bedrock-knowledgebase-rdsfieldmapping-vectorfield
	//
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
}

