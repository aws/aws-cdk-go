package awsbedrock


// Contains configurations for authentication to an Amazon Redshift provisioned data warehouse.
//
// Specify the type of authentication to use in the `type` field and include the corresponding field. If you specify IAM authentication, you don't need to include another field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftProvisionedAuthConfigurationProperty := &RedshiftProvisionedAuthConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DatabaseUser: jsii.String("databaseUser"),
//   	UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html
//
type CfnKnowledgeBase_RedshiftProvisionedAuthConfigurationProperty struct {
	// The type of authentication to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The database username for authentication to an Amazon Redshift provisioned data warehouse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-databaseuser
	//
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	// The ARN of an Secrets Manager secret for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-usernamepasswordsecretarn
	//
	UsernamePasswordSecretArn *string `field:"optional" json:"usernamePasswordSecretArn" yaml:"usernamePasswordSecretArn"`
}

