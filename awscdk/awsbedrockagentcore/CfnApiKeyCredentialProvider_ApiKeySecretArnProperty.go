package awsbedrockagentcore


// Contains information about the API key secret in AWS Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   apiKeySecretArnProperty := &ApiKeySecretArnProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-apikeycredentialprovider-apikeysecretarn.html
//
type CfnApiKeyCredentialProvider_ApiKeySecretArnProperty struct {
	// The ARN of the secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-apikeycredentialprovider-apikeysecretarn.html#cfn-bedrockagentcore-apikeycredentialprovider-apikeysecretarn-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

