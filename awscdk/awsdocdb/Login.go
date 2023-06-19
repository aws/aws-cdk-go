package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Login credentials for a database cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   		 // NOTE: 'admin' is reserved by DocumentDB
//   		ExcludeCharacters: jsii.String("\"@/:"),
//   		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
//   		SecretName: jsii.String("/myapp/mydocdb/masteruser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_R5, ec2.InstanceSize_LARGE),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	Vpc: Vpc,
//   })
//
// Experimental.
type Login struct {
	// Username.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Specifies characters to not include in generated passwords.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// KMS encryption key to encrypt the generated secret.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	// Experimental.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
	// The physical name of the secret, that will be generated.
	// Experimental.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

