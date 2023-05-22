package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	ImageTagMutability: ecr.TagMutability_IMMUTABLE,
//   })
//
// Experimental.
type RepositoryProps struct {
	// The kind of server-side encryption to apply to this repository.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryptionKey is not specified, an AWS managed KMS key is used.
	// Experimental.
	Encryption RepositoryEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for repository encryption.
	//
	// The 'encryption' property must be either not specified or set to "KMS".
	// An error will be emitted if encryption is set to "AES256".
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Enable the scan on push when creating the repository.
	// Experimental.
	ImageScanOnPush *bool `field:"optional" json:"imageScanOnPush" yaml:"imageScanOnPush"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten.
	// Experimental.
	ImageTagMutability TagMutability `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// The AWS account ID associated with the registry that contains the repository.
	// See: https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_PutLifecyclePolicy.html
	//
	// Experimental.
	LifecycleRegistryId *string `field:"optional" json:"lifecycleRegistryId" yaml:"lifecycleRegistryId"`
	// Life cycle rules to apply to this registry.
	// Experimental.
	LifecycleRules *[]*LifecycleRule `field:"optional" json:"lifecycleRules" yaml:"lifecycleRules"`
	// Determine what happens to the repository when the resource/stack is deleted.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Name for this repository.
	// Experimental.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

