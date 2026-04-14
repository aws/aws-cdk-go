package awsbedrockagentcore


// Contains information about a secret in AWS Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   clientSecretArnProperty := &ClientSecretArnProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-clientsecretarn.html
//
type CfnOAuth2CredentialProviderPropsMixin_ClientSecretArnProperty struct {
	// The ARN of the secret in AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-clientsecretarn.html#cfn-bedrockagentcore-oauth2credentialprovider-clientsecretarn-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

