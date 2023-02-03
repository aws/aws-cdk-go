package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Options for creating Credentials from a username.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//   var secretValue secretValue
//
//   credentialsFromUsernameOptions := &credentialsFromUsernameOptions{
//   	encryptionKey: key,
//   	excludeCharacters: jsii.String("excludeCharacters"),
//   	password: secretValue,
//   	replicaRegions: []replicaRegion{
//   		&replicaRegion{
//   			region: jsii.String("region"),
//
//   			// the properties below are optional
//   			encryptionKey: key,
//   		},
//   	},
//   	secretName: jsii.String("secretName"),
//   }
//
type CredentialsFromUsernameOptions struct {
	// KMS encryption key to encrypt the generated secret.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The characters to exclude from the generated password.
	//
	// Has no effect if `password` has been provided.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// The name of the secret.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
}

