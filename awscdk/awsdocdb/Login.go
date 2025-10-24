package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Login credentials for a database cluster.
//
// Example:
//   var vpc Vpc
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	Vpc: Vpc,
//   	DeletionProtection: jsii.Boolean(true),
//   })
//
type Login struct {
	// Username.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Specifies characters to not include in generated passwords.
	// Default: "\"@/".
	//
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// KMS encryption key to encrypt the generated secret.
	// Default: default master key.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Default: a Secrets Manager generated password.
	//
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
	// The physical name of the secret, that will be generated.
	// Default: Secretsmanager will generate a physical name for the secret.
	//
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

