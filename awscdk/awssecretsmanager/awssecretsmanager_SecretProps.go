package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// The properties required to create a new secret in AWS Secrets Manager.
//
// Example:
//   var stack stack
//   user := iam.NewUser(stack, jsii.String("User"))
//   accessKey := iam.NewAccessKey(stack, jsii.String("AccessKey"), &accessKeyProps{
//   	user: user,
//   })
//
//   secretsmanager.NewSecret(stack, jsii.String("Secret"), &secretProps{
//   	secretObjectValue: map[string]secretValue{
//   		"username": awscdk.SecretValue.unsafePlainText(user.userName),
//   		"database": awscdk.SecretValue.unsafePlainText(jsii.String("foo")),
//   		"password": accessKey.secretAccessKey,
//   	},
//   })
//
type SecretProps struct {
	// An optional, human-friendly description of the secret.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The customer-managed encryption key to use for encrypting the secret value.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Configuration for how to generate a secret value.
	//
	// Only one of `secretString` and `generateSecretString` can be provided.
	GenerateSecretString *SecretStringGenerator `field:"optional" json:"generateSecretString" yaml:"generateSecretString"`
	// Policy to apply when the secret is removed from this stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// A name for the secret.
	//
	// Note that deleting secrets from SecretsManager does not happen immediately, but after a 7 to
	// 30 days blackout period. During that period, it is not possible to create another secret that shares the same name.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Initial value for a JSON secret.
	//
	// **NOTE:** *It is **highly** encouraged to leave this field undefined and allow SecretsManager to create the secret value.
	// The secret object -- if provided -- will be included in the output of the cdk as part of synthesis,
	// and will appear in the CloudFormation template in the console. This can be secure(-ish) if that value is merely reference to
	// another resource (or one of its attributes), but if the value is a plaintext string, it will be visible to anyone with access
	// to the CloudFormation template (via the AWS Console, SDKs, or CLI).
	//
	// Specifies a JSON object that you want to encrypt and store in this new version of the secret.
	// To specify a simple string value instead, use `SecretProps.secretStringValue`
	//
	// Only one of `secretStringBeta1`, `secretStringValue`, 'secretObjectValue', and `generateSecretString` can be provided.
	//
	// Example:
	//   var user user
	//   var accessKey accessKey
	//   var stack stack
	//
	//   secretsmanager.NewSecret(stack, jsii.String("JSONSecret"), &secretProps{
	//   	secretObjectValue: map[string]secretValue{
	//   		"username": awscdk.SecretValue.unsafePlainText(user.userName),
	//   		 // intrinsic reference, not exposed as plaintext
	//   		"database": awscdk.SecretValue.unsafePlainText(jsii.String("foo")),
	//   		 // rendered as plain text, but not a secret
	//   		"password": accessKey.secretAccessKey,
	//   	},
	//   })
	//
	SecretObjectValue *map[string]awscdk.SecretValue `field:"optional" json:"secretObjectValue" yaml:"secretObjectValue"`
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
	// May be a simple string value. To provide a string representation of JSON structure, use `SecretProps.secretObjectValue` instead.
	//
	// Only one of `secretStringBeta1`, `secretStringValue`, 'secretObjectValue', and `generateSecretString` can be provided.
	SecretStringValue awscdk.SecretValue `field:"optional" json:"secretStringValue" yaml:"secretStringValue"`
}

