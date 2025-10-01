package awsbedrock


// Specifies configurations for authentication to a Redshift Serverless.
//
// Specify the type of authentication to use in the `type` field and include the corresponding field. If you specify IAM authentication, you don't need to include another field.
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
	// The type of authentication to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftserverlessauthconfiguration-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ARN of an Secrets Manager secret for authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-redshiftserverlessauthconfiguration.html#cfn-bedrock-knowledgebase-redshiftserverlessauthconfiguration-usernamepasswordsecretarn
	//
	UsernamePasswordSecretArn *string `field:"optional" json:"usernamePasswordSecretArn" yaml:"usernamePasswordSecretArn"`
}

