package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes required to import an existing secret into the Stack.
//
// One ARN format (`secretArn`, `secretCompleteArn`, `secretPartialArn`) must be provided.
//
// Example:
//   userpool := cognito.NewUserPool(this, jsii.String("Pool"))
//   secret := secretsmanager.Secret_FromSecretAttributes(this, jsii.String("CognitoClientSecret"), &SecretAttributes{
//   	SecretCompleteArn: jsii.String("arn:aws:secretsmanager:xxx:xxx:secret:xxx-xxx"),
//   }).SecretValue
//
//   provider := cognito.NewUserPoolIdentityProviderGoogle(this, jsii.String("Google"), &UserPoolIdentityProviderGoogleProps{
//   	ClientId: jsii.String("amzn-client-id"),
//   	ClientSecretValue: secret,
//   	UserPool: userpool,
//   })
//
type SecretAttributes struct {
	// The encryption key that is used to encrypt the secret, unless the default SecretsManager key is used.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The complete ARN of the secret in SecretsManager.
	//
	// This is the ARN including the Secrets Manager 6-character suffix.
	// Cannot be used with `secretArn` or `secretPartialArn`.
	SecretCompleteArn *string `field:"optional" json:"secretCompleteArn" yaml:"secretCompleteArn"`
	// The partial ARN of the secret in SecretsManager.
	//
	// This is the ARN without the Secrets Manager 6-character suffix.
	// Cannot be used with `secretArn` or `secretCompleteArn`.
	SecretPartialArn *string `field:"optional" json:"secretPartialArn" yaml:"secretPartialArn"`
}

