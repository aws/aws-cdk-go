package awsecr

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   ecr.NewRepository(this, jsii.String("Repo"), &RepositoryProps{
//   	ImageTagMutability: ecr.TagMutability_IMMUTABLE,
//   })
//
type RepositoryProps struct {
	// Whether all images should be automatically deleted when the repository is removed from the stack or when the stack is deleted.
	//
	// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
	// Default: false.
	//
	// Deprecated: Use `emptyOnDelete` instead.
	AutoDeleteImages *bool `field:"optional" json:"autoDeleteImages" yaml:"autoDeleteImages"`
	// If true, deleting the repository force deletes the contents of the repository.
	//
	// If false, the repository must be empty before attempting to delete it.
	// Default: false.
	//
	EmptyOnDelete *bool `field:"optional" json:"emptyOnDelete" yaml:"emptyOnDelete"`
	// The kind of server-side encryption to apply to this repository.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryptionKey is not specified, an AWS managed KMS key is used.
	// Default: - `KMS` if `encryptionKey` is specified, or `AES256` otherwise.
	//
	Encryption RepositoryEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for repository encryption.
	//
	// The 'encryption' property must be either not specified or set to "KMS".
	// An error will be emitted if encryption is set to "AES256".
	// Default: - If encryption is set to `KMS` and this property is undefined,
	// an AWS managed KMS key is used.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Enable the scan on push when creating the repository.
	// Default: false.
	//
	ImageScanOnPush *bool `field:"optional" json:"imageScanOnPush" yaml:"imageScanOnPush"`
	// The tag mutability setting for the repository.
	//
	// If this parameter is omitted, the default setting of MUTABLE will be used which will allow image tags to be overwritten.
	// Default: TagMutability.MUTABLE
	//
	ImageTagMutability TagMutability `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// The AWS account ID associated with the registry that contains the repository.
	// See: https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_PutLifecyclePolicy.html
	//
	// Default: The default registry is assumed.
	//
	LifecycleRegistryId *string `field:"optional" json:"lifecycleRegistryId" yaml:"lifecycleRegistryId"`
	// Life cycle rules to apply to this registry.
	// Default: No life cycle rules.
	//
	LifecycleRules *[]*LifecycleRule `field:"optional" json:"lifecycleRules" yaml:"lifecycleRules"`
	// Determine what happens to the repository when the resource/stack is deleted.
	// Default: RemovalPolicy.Retain
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Name for this repository.
	//
	// The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// Default: Automatically generated name.
	//
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

