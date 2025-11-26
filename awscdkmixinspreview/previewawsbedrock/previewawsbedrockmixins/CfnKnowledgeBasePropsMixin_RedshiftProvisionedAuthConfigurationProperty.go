package previewawsbedrockmixins


// Contains configurations for authentication to an Amazon Redshift provisioned data warehouse.
//
// Specify the type of authentication to use in the `type` field and include the corresponding field. If you specify IAM authentication, you don't need to include another field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftProvisionedAuthConfigurationProperty := &RedshiftProvisionedAuthConfigurationProperty{
//   	DatabaseUser: jsii.String("databaseUser"),
//   	Type: jsii.String("type"),
//   	UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html
//
type CfnKnowledgeBasePropsMixin_RedshiftProvisionedAuthConfigurationProperty struct {
	// The database username for authentication to an Amazon Redshift provisioned data warehouse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-databaseuser
	//
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	// The type of authentication to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The ARN of an Secrets Manager secret for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-usernamepasswordsecretarn
	//
	UsernamePasswordSecretArn *string `field:"optional" json:"usernamePasswordSecretArn" yaml:"usernamePasswordSecretArn"`
}

