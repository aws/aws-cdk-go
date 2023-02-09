package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// The properties required to create a new secret in AWS Secrets Manager.
//
// Example:
//   // Creates a new IAM user, access and secret keys, and stores the secret access key in a Secret.
//   user := iam.NewUser(this, jsii.String("User"))
//   accessKey := iam.NewAccessKey(this, jsii.String("AccessKey"), &accessKeyProps{
//   	user: user,
//   })
//   secretValue := secretsmanager.secretStringValueBeta1.fromToken(accessKey.secretAccessKey.toString())
//   secretsmanager.NewSecret(this, jsii.String("Secret"), &secretProps{
//   	secretStringBeta1: secretValue,
//   })
//
// Experimental.
type SecretProps struct {
	// An optional, human-friendly description of the secret.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The customer-managed encryption key to use for encrypting the secret value.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Configuration for how to generate a secret value.
	//
	// Only one of `secretString` and `generateSecretString` can be provided.
	// Experimental.
	GenerateSecretString *SecretStringGenerator `field:"optional" json:"generateSecretString" yaml:"generateSecretString"`
	// Policy to apply when the secret is removed from this stack.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// A list of regions where to replicate this secret.
	// Experimental.
	ReplicaRegions *[]*ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// A name for the secret.
	//
	// Note that deleting secrets from SecretsManager does not happen immediately, but after a 7 to
	// 30 days blackout period. During that period, it is not possible to create another secret that shares the same name.
	// Experimental.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Initial value for the secret.
	//
	// **NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
	// The secret string -- if provided -- will be included in the output of the cdk as part of synthesis,
	// and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
	// another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
	// to the CloudFormation template (via the AWS Console, SDKs, or CLI).
	//
	// Specifies text data that you want to encrypt and store in this new version of the secret.
	// May be a simple string value, or a string representation of a JSON structure.
	//
	// Only one of `secretStringBeta1`, `secretStringValue`, and `generateSecretString` can be provided.
	// Deprecated: Use `secretStringValue` instead.
	SecretStringBeta1 SecretStringValueBeta1 `field:"optional" json:"secretStringBeta1" yaml:"secretStringBeta1"`
	// Initial value for the secret.
	//
	// **NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
	// The secret string -- if provided -- will be included in the output of the cdk as part of synthesis,
	// and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
	// another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
	// to the CloudFormation template (via the AWS Console, SDKs, or CLI).
	//
	// Specifies text data that you want to encrypt and store in this new version of the secret.
	// May be a simple string value, or a string representation of a JSON structure.
	//
	// Only one of `secretStringBeta1`, `secretStringValue`, and `generateSecretString` can be provided.
	// Experimental.
	SecretStringValue awscdk.SecretValue `field:"optional" json:"secretStringValue" yaml:"secretStringValue"`
}

