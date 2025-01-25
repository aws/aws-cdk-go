package awsbedrock


// Configurations for Redshift query engine provisioned auth setup.
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
	// Provisioned Redshift auth type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Redshift database user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-databaseuser
	//
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	// Arn of a SecretsManager Secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftprovisionedauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftprovisionedauthconfiguration-usernamepasswordsecretarn
	//
	UsernamePasswordSecretArn *string `field:"optional" json:"usernamePasswordSecretArn" yaml:"usernamePasswordSecretArn"`
}

