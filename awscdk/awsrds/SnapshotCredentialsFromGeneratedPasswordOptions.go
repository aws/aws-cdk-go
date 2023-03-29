package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Options used in the `SnapshotCredentials.fromGeneratedPassword` method.
//
// Example:
//   var vpc vpc
//
//   engine := rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
//   	Version: rds.PostgresEngineVersion_VER_12_3(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("InstanceFromSnapshotWithCustomizedSecret"), &DatabaseInstanceFromSnapshotProps{
//   	Engine: Engine,
//   	Vpc: Vpc,
//   	SnapshotIdentifier: jsii.String("mySnapshot"),
//   	Credentials: rds.SnapshotCredentials_FromGeneratedSecret(jsii.String("username"), &SnapshotCredentialsFromGeneratedPasswordOptions{
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
type SnapshotCredentialsFromGeneratedPasswordOptions struct {
	// KMS encryption key to encrypt the generated secret.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The characters to exclude from the generated password.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// A list of regions where to replicate this secret.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
}

