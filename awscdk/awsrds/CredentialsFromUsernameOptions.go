package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Options for creating Credentials from a username.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
//   		Version: rds.AuroraPostgresEngineVersion_VER_15_2(),
//   	}),
//   	Credentials: rds.Credentials_FromUsername(jsii.String("adminuser"), &CredentialsFromUsernameOptions{
//   		Password: awscdk.SecretValue_UnsafePlainText(jsii.String("7959866cacc02c2d243ecfe177464fe6")),
//   	}),
//   	InstanceProps: &InstanceProps{
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_X2G, ec2.InstanceSize_XLARGE),
//   		VpcSubnets: &SubnetSelection{
//   			SubnetType: ec2.SubnetType_PUBLIC,
//   		},
//   		Vpc: *Vpc,
//   	},
//   	StorageType: rds.DBClusterStorageType_AURORA_IOPT1,
//   })
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

