package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Construction properties for a DatabaseSecret.
//
// Example:
//   var cluster DatabaseCluster
//
//   myUserSecret := docdb.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &DatabaseSecretProps{
//   	Username: jsii.String("myuser"),
//   	MasterSecret: cluster.Secret,
//   })
//   myUserSecretAttached := myUserSecret.attach(cluster) // Adds DB connections information in the secret
//
//   cluster.AddRotationMultiUser(jsii.String("MyUser"), &RotationMultiUserOptions{
//   	 // Add rotation using the multi user scheme
//   	Secret: myUserSecretAttached,
//   })
//
type DatabaseSecretProps struct {
	// The username.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	// Default: default master key.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Characters to not include in the generated password.
	// Default: "\"@/".
	//
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret which will be used to rotate this secret.
	// Default: - no master secret information will be included.
	//
	MasterSecret awssecretsmanager.ISecret `field:"optional" json:"masterSecret" yaml:"masterSecret"`
	// The physical name of the secret.
	// Default: Secretsmanager will generate a physical name for the secret.
	//
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

