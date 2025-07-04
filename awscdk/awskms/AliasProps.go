package awskms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a KMS Key Alias object.
//
// Example:
//   // Passing an encrypted replication bucket created in a different stack.
//   app := awscdk.NewApp()
//   replicationStack := awscdk.Newstack(app, jsii.String("ReplicationStack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-1"),
//   	},
//   })
//   key := kms.NewKey(replicationStack, jsii.String("ReplicationKey"))
//   alias := kms.NewAlias(replicationStack, jsii.String("ReplicationAlias"), &AliasProps{
//   	// aliasName is required
//   	AliasName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   	TargetKey: key,
//   })
//   replicationBucket := s3.NewBucket(replicationStack, jsii.String("ReplicationBucket"), &BucketProps{
//   	BucketName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   	EncryptionKey: alias,
//   })
//
type AliasProps struct {
	// The name of the alias.
	//
	// The name must start with alias followed by a
	// forward slash, such as alias/. You can't specify aliases that begin with
	// alias/AWS. These aliases are reserved.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// The ID of the key for which you are creating the alias.
	//
	// Specify the key's
	// globally unique identifier or Amazon Resource Name (ARN). You can't
	// specify another alias.
	TargetKey IKey `field:"required" json:"targetKey" yaml:"targetKey"`
	// Policy to apply when the alias is removed from this stack.
	// Default: - The alias will be deleted.
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

