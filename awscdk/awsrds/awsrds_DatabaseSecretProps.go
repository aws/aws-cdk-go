package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Construction properties for a DatabaseSecret.
//
// Example:
//   var instance databaseInstance
//
//   myUserSecret := rds.NewDatabaseSecret(this, jsii.String("MyUserSecret"), &databaseSecretProps{
//   	username: jsii.String("myuser"),
//   	secretName: jsii.String("my-user-secret"),
//   	 // optional, defaults to a CloudFormation-generated name
//   	masterSecret: instance.secret,
//   	excludeCharacters: jsii.String("{}[]()'\"/\\"),
//   })
//   myUserSecretAttached := myUserSecret.attach(instance) // Adds DB connections information in the secret
//
//   instance.addRotationMultiUser(jsii.String("MyUser"), &rotationMultiUserOptions{
//   	 // Add rotation using the multi user scheme
//   	secret: myUserSecretAttached,
//   })
//
type DatabaseSecretProps struct {
	// The username.
	Username *string `field:"required" json:"username" yaml:"username"`
	// The KMS key to use to encrypt the secret.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Characters to not include in the generated password.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The master secret which will be used to rotate this secret.
	MasterSecret awssecretsmanager.ISecret `field:"optional" json:"masterSecret" yaml:"masterSecret"`
	// Whether to replace this secret when the criteria for the password change.
	//
	// This is achieved by overriding the logical id of the AWS::SecretsManager::Secret
	// with a hash of the options that influence the password generation. This
	// way a new secret will be created when the password is regenerated and the
	// cluster or instance consuming this secret will have its credentials updated.
	ReplaceOnPasswordCriteriaChanges *bool `field:"optional" json:"replaceOnPasswordCriteriaChanges" yaml:"replaceOnPasswordCriteriaChanges"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// A name for the secret.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

