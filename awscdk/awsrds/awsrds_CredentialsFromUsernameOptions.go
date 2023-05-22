package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Options for creating Credentials from a username.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//   var secretValue secretValue
//
//   credentialsFromUsernameOptions := &CredentialsFromUsernameOptions{
//   	EncryptionKey: key,
//   	ExcludeCharacters: jsii.String("excludeCharacters"),
//   	Password: secretValue,
//   	ReplicaRegions: []replicaRegion{
//   		&replicaRegion{
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			EncryptionKey: key,
//   		},
//   	},
//   	SecretName: jsii.String("secretName"),
//   }
//
// Experimental.
type CredentialsFromUsernameOptions struct {
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The characters to exclude from the generated password.
	//
	// Has no effect if {@link password} has been provided.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// A list of regions where to replicate this secret.
	// Experimental.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// The name of the secret.
	// Experimental.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Experimental.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
}

