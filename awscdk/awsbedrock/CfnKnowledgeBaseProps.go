package awsbedrock


// Properties for defining a `CfnKnowledgeBase`.
//
// Example:
//   var collection ICollectionRef
//   var role Role
//
//
//   awscdk.Aws_bedrock.NewCfnKnowledgeBase(this, jsii.String("KnowledgeBase"), &CfnKnowledgeBaseProps{
//   	Name: jsii.String("product-documentation"),
//   	RoleArn: role.roleArn,
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		Type: jsii.String("VECTOR"),
//   		VectorKnowledgeBaseConfiguration: &VectorKnowledgeBaseConfigurationProperty{
//   			EmbeddingModelArn: bedrock.BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024().ModelArn,
//   		},
//   	},
//   }).With(mixins.NewKnowledgeBaseOpenSearchServerlessStorage(&KnowledgeBaseOpenSearchServerlessStorageProps{
//   	Collection: Collection,
//   	VectorIndexName: jsii.String("bedrock-knowledge-base-default-index"),
//   	VectorField: jsii.String("bedrock-knowledge-base-default-vector"),
//   	TextField: jsii.String("AMAZON_BEDROCK_TEXT_CHUNK"),
//   	MetadataField: jsii.String("AMAZON_BEDROCK_METADATA"),
//   }))
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html
//
type CfnKnowledgeBaseProps struct {
	// Contains details about the embeddings configuration of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-knowledgebaseconfiguration
	//
	KnowledgeBaseConfiguration interface{} `field:"required" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// The name of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the IAM role with permissions to invoke API operations on the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The description of the knowledge base associated with the inline agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Contains details about the storage configuration of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-storageconfiguration
	//
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// Metadata that you can assign to a resource as key-value pairs. For more information, see the following resources:.
	//
	// - [Tag naming limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)
	// - [Tagging best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebase.html#cfn-bedrock-knowledgebase-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

