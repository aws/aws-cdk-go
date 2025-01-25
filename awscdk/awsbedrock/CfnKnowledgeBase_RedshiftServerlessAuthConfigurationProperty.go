package awsbedrock


// Configurations for Redshift query engine serverless auth setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftServerlessAuthConfigurationProperty := &RedshiftServerlessAuthConfigurationProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	UsernamePasswordSecretArn: jsii.String("usernamePasswordSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessauthconfiguration.html
//
type CfnKnowledgeBase_RedshiftServerlessAuthConfigurationProperty struct {
	// Serverless Redshift auth type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftserverlessauthconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Arn of a SecretsManager Secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftserverlessauthconfiguration-usernamepasswordsecretarn
	//
	UsernamePasswordSecretArn *string `field:"optional" json:"usernamePasswordSecretArn" yaml:"usernamePasswordSecretArn"`
}

