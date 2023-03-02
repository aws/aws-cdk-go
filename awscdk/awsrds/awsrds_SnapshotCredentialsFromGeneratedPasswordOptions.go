package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Options used in the {@link SnapshotCredentials.fromGeneratedPassword} method.
//
// Example:
//   var vpc vpc
//
//   engine := rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   	version: rds.postgresEngineVersion_VER_12_3(),
//   })
//   myKey := kms.NewKey(this, jsii.String("MyKey"))
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("InstanceFromSnapshotWithCustomizedSecret"), &databaseInstanceFromSnapshotProps{
//   	engine: engine,
//   	vpc: vpc,
//   	snapshotIdentifier: jsii.String("mySnapshot"),
//   	credentials: rds.snapshotCredentials.fromGeneratedSecret(jsii.String("username"), &snapshotCredentialsFromGeneratedPasswordOptions{
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
type SnapshotCredentialsFromGeneratedPasswordOptions struct {
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The characters to exclude from the generated password.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// A list of regions where to replicate this secret.
	// Experimental.
	ReplicaRegions *[]*awssecretsmanager.ReplicaRegion `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
}

