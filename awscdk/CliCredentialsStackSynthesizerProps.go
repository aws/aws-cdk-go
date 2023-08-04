package awscdk


// Properties for the CliCredentialsStackSynthesizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cliCredentialsStackSynthesizerProps := &CliCredentialsStackSynthesizerProps{
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	DockerTagPrefix: jsii.String("dockerTagPrefix"),
//   	FileAssetsBucketName: jsii.String("fileAssetsBucketName"),
//   	ImageAssetsRepositoryName: jsii.String("imageAssetsRepositoryName"),
//   	Qualifier: jsii.String("qualifier"),
//   }
//
type CliCredentialsStackSynthesizerProps struct {
	// bucketPrefix to use while storing S3 Assets.
	// Default: - DefaultStackSynthesizer.DEFAULT_FILE_ASSET_PREFIX
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// A prefix to use while tagging and uploading Docker images to ECR.
	//
	// This does not add any separators - the source hash will be appended to
	// this string directly.
	// Default: - DefaultStackSynthesizer.DEFAULT_DOCKER_ASSET_PREFIX
	//
	DockerTagPrefix *string `field:"optional" json:"dockerTagPrefix" yaml:"dockerTagPrefix"`
	// Name of the S3 bucket to hold file assets.
	//
	// You must supply this if you have given a non-standard name to the staging bucket.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_FILE_ASSETS_BUCKET_NAME
	//
	FileAssetsBucketName *string `field:"optional" json:"fileAssetsBucketName" yaml:"fileAssetsBucketName"`
	// Name of the ECR repository to hold Docker Image assets.
	//
	// You must supply this if you have given a non-standard name to the ECR repository.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	// Default: DefaultStackSynthesizer.DEFAULT_IMAGE_ASSETS_REPOSITORY_NAME
	//
	ImageAssetsRepositoryName *string `field:"optional" json:"imageAssetsRepositoryName" yaml:"imageAssetsRepositoryName"`
	// Qualifier to disambiguate multiple environments in the same account.
	//
	// You can use this and leave the other naming properties empty if you have deployed
	// the bootstrap environment with standard names but only different qualifiers.
	// Default: - Value of context key '@aws-cdk/core:bootstrapQualifier' if set, otherwise `DefaultStackSynthesizer.DEFAULT_QUALIFIER`
	//
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

