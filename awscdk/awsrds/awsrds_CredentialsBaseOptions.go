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
//   engine := rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   	version: rds.postgresEngineVersion_VER_12_3(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstance(this, jsii.String("InstanceWithCustomizedSecret"), &databaseInstanceProps{
//   	engine: engine,
//   	vpc: vpc,
//   	credentials: rds.credentials.fromGeneratedSecret(jsii.String("postgres"), &credentialsBaseOptions{
//   		secretName: jsii.String("my-cool-name"),
//   		encryptionKey: myKey,
//   		excludeCharacters: jsii.String("!&*^#@()"),
//   		replicaRegions: []replicaRegion{
//   			&replicaRegion{
//   				region: jsii.String("eu-west-1"),
//   			},
//   			&replicaRegion{
//   				region: jsii.String("eu-west-2"),
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

