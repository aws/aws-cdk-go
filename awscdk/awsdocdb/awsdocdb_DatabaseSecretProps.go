package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Construction properties for a DatabaseSecret.
//
// Example:
//   var cluster databaseCluster
//
//   myUserSecret := docdb.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &databaseSecretProps{
//   	username: jsii.String("myuser"),
//   	masterSecret: cluster.secret,
//   })
//   myUserSecretAttached := myUserSecret.attach(cluster) // Adds DB connections information in the secret
//
//   cluster.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
//   	 // Add rotation using the multi user scheme
//   	secret: myUserSecretAttached,
//   })
//
// Experimental.
type DatabaseSecretProps struct {
	// The username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Characters to not include in the generated password.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret which will be used to rotate this secret.
	// Experimental.
	MasterSecret awssecretsmanager.ISecret `field:"optional" json:"masterSecret" yaml:"masterSecret"`
	// The physical name of the secret.
	// Experimental.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

