package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
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
type CredentialsBaseOptions struct {
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
}

