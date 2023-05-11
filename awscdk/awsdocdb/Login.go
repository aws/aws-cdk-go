package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Login credentials for a database cluster.
//
// Example:
//   var vpc vpc
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
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// KMS encryption key to encrypt the generated secret.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
	// The physical name of the secret, that will be generated.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

