package awsbedrock


// Contains details about the storage configuration of the knowledge base in Amazon RDS.
//
// For more information, see [Create a vector index in Amazon RDS](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-setup-rds.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rdsConfigurationProperty := &RdsConfigurationProperty{
//   	CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   	DatabaseName: jsii.String("databaseName"),
//   	FieldMapping: &RdsFieldMappingProperty{
//   		MetadataField: jsii.String("metadataField"),
//   		PrimaryKeyField: jsii.String("primaryKeyField"),
//   		TextField: jsii.String("textField"),
//   		VectorField: jsii.String("vectorField"),
//
//   		// the properties below are optional
//   		CustomMetadataField: jsii.String("customMetadataField"),
//   	},
//   	ResourceArn: jsii.String("resourceArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsconfiguration.html
//
type CfnKnowledgeBase_RdsConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the secret that you created in AWS Secrets Manager that is linked to your Amazon RDS database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsconfiguration.html#cfn-bedrock-knowledgebase-rdsconfiguration-credentialssecretarn
	//
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The name of your Amazon RDS database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsconfiguration.html#cfn-bedrock-knowledgebase-rdsconfiguration-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Contains the names of the fields to which to map information about the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsconfiguration.html#cfn-bedrock-knowledgebase-rdsconfiguration-fieldmapping
	//
	FieldMapping interface{} `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// The Amazon Resource Name (ARN) of the vector store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsconfiguration.html#cfn-bedrock-knowledgebase-rdsconfiguration-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The name of the table in the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-rdsconfiguration.html#cfn-bedrock-knowledgebase-rdsconfiguration-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

