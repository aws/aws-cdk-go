package appstagingsynthesizeralpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// User configurable options to the DefaultStagingStack.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultStagingStack := appstagingsynthesizeralpha.DefaultStagingStack_Factory(&DefaultStagingStackOptions{
//   	AppId: jsii.String("my-app-id"),
//   	StagingBucketEncryption: awscdk.BucketEncryption_S3_MANAGED,
//   })
//
// Experimental.
type DefaultStagingStackOptions struct {
	// A unique identifier for the application that the staging stack belongs to.
	//
	// This identifier will be used in the name of staging resources
	// created for this application, and should be unique across CDK apps.
	//
	// The identifier should include lowercase characters and dashes ('-') only
	// and have a maximum of 20 characters.
	// Experimental.
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// Encryption type for staging bucket.
	//
	// In future versions of this package, the default will be BucketEncryption.S3_MANAGED.
	//
	// In previous versions of this package, the default was to use KMS encryption for the staging bucket. KMS keys cost
	// $1/month, which could result in unexpected costs for users who are not aware of this. As we stabilize this module
	// we intend to make the default S3-managed encryption, which is free. However, the migration path from KMS to S3
	// managed encryption for existing buckets is not straightforward. Therefore, for now, this property is required.
	//
	// If you have an existing staging bucket encrypted with a KMS key, you will likely want to set this property to
	// BucketEncryption.KMS. If you are creating a new staging bucket, you can set this property to
	// BucketEncryption.S3_MANAGED to avoid the cost of a KMS key.
	// Experimental.
	StagingBucketEncryption awss3.BucketEncryption `field:"required" json:"stagingBucketEncryption" yaml:"stagingBucketEncryption"`
	// Auto deletes objects in the staging S3 bucket and images in the staging ECR repositories.
	// Default: true.
	//
	// Experimental.
	AutoDeleteStagingAssets *bool `field:"optional" json:"autoDeleteStagingAssets" yaml:"autoDeleteStagingAssets"`
	// The lifetime for deploy time file assets.
	//
	// Assets that are only necessary at deployment time (for instance,
	// CloudFormation templates and Lambda source code bundles) will be
	// automatically deleted after this many days. Assets that may be
	// read from the staging bucket during your application's run time
	// will not be deleted.
	//
	// Set this to the length of time you wish to be able to roll back to
	// previous versions of your application without having to do a new
	// `cdk synth` and re-upload of assets.
	// Default: - Duration.days(30)
	//
	// Experimental.
	DeployTimeFileAssetLifetime awscdk.Duration `field:"optional" json:"deployTimeFileAssetLifetime" yaml:"deployTimeFileAssetLifetime"`
	// Pass in an existing role to be used as the file publishing role.
	// Default: - a new role will be created.
	//
	// Experimental.
	FileAssetPublishingRole BootstrapRole `field:"optional" json:"fileAssetPublishingRole" yaml:"fileAssetPublishingRole"`
	// Pass in an existing role to be used as the image publishing role.
	// Default: - a new role will be created.
	//
	// Experimental.
	ImageAssetPublishingRole BootstrapRole `field:"optional" json:"imageAssetPublishingRole" yaml:"imageAssetPublishingRole"`
	// The maximum number of image versions to store in a repository.
	//
	// Previous versions of an image can be stored for rollback purposes.
	// Once a repository has more than 3 image versions stored, the oldest
	// version will be discarded. This allows for sensible garbage collection
	// while maintaining a few previous versions for rollback scenarios.
	// Default: - up to 3 versions stored.
	//
	// Experimental.
	ImageAssetVersionCount *float64 `field:"optional" json:"imageAssetVersionCount" yaml:"imageAssetVersionCount"`
	// Explicit name for the staging bucket.
	// Default: - a well-known name unique to this app/env.
	//
	// Experimental.
	StagingBucketName *string `field:"optional" json:"stagingBucketName" yaml:"stagingBucketName"`
	// Specify a custom prefix to be used as the staging stack name and construct ID.
	//
	// The prefix will be appended before the appId, which
	// is required to be part of the stack name and construct ID to
	// ensure uniqueness.
	// Default: 'StagingStack'.
	//
	// Experimental.
	StagingStackNamePrefix *string `field:"optional" json:"stagingStackNamePrefix" yaml:"stagingStackNamePrefix"`
}

