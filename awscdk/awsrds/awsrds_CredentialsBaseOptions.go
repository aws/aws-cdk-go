package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Base options for creating Credentials.
//
// Example:
//   var vpc vpc
//
//   engine := rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
//   	Version: rds.PostgresEngineVersion_VER_12_3(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstance(this, jsii.String("InstanceWithCustomizedSecret"), &DatabaseInstanceProps{
//   	Engine: Engine,
//   	Vpc: Vpc,
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("postgres"), &CredentialsBaseOptions{
//   		SecretName: jsii.String("my-cool-name"),
//   		EncryptionKey: myKey,
//   		ExcludeCharacters: jsii.String("!&*^#@()"),
//   		ReplicaRegions: []replicaRegion{
//   			&replicaRegion{
//   				Region: jsii.String("eu-west-1"),
//   			},
//   			&replicaRegion{
//   				Region: jsii.String("eu-west-2"),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type CredentialsBaseOptions struct {
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
}

