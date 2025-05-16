package awsbedrock


// Contains details about the storage configuration of the knowledge base in Pinecone.
//
// For more information, see [Create a vector index in Pinecone](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-pinecone.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pineconeConfigurationProperty := &PineconeConfigurationProperty{
//   	ConnectionString: jsii.String("connectionString"),
//   	CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   	FieldMapping: &PineconeFieldMappingProperty{
//   		MetadataField: jsii.String("metadataField"),
//   		TextField: jsii.String("textField"),
//   	},
//
//   	// the properties below are optional
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-pineconeconfiguration.html
//
type CfnKnowledgeBase_PineconeConfigurationProperty struct {
	// The endpoint URL for your index management page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-pineconeconfiguration.html#cfn-bedrock-knowledgebase-pineconeconfiguration-connectionstring
	//
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// The Amazon Resource Name (ARN) of the secret that you created in AWS Secrets Manager that is linked to your Pinecone API key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-pineconeconfiguration.html#cfn-bedrock-knowledgebase-pineconeconfiguration-credentialssecretarn
	//
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// Contains the names of the fields to which to map information about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-pineconeconfiguration.html#cfn-bedrock-knowledgebase-pineconeconfiguration-fieldmapping
	//
	FieldMapping interface{} `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// The namespace to be used to write new data to your database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-pineconeconfiguration.html#cfn-bedrock-knowledgebase-pineconeconfiguration-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

