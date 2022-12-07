// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Properties for the CliCredentialsStackSynthesizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cliCredentialsStackSynthesizerProps := &cliCredentialsStackSynthesizerProps{
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	dockerTagPrefix: jsii.String("dockerTagPrefix"),
//   	fileAssetsBucketName: jsii.String("fileAssetsBucketName"),
//   	imageAssetsRepositoryName: jsii.String("imageAssetsRepositoryName"),
//   	qualifier: jsii.String("qualifier"),
//   }
//
type CliCredentialsStackSynthesizerProps struct {
	// bucketPrefix to use while storing S3 Assets.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// A prefix to use while tagging and uploading Docker images to ECR.
	//
	// This does not add any separators - the source hash will be appended to
	// this string directly.
	DockerTagPrefix *string `field:"optional" json:"dockerTagPrefix" yaml:"dockerTagPrefix"`
	// Name of the S3 bucket to hold file assets.
	//
	// You must supply this if you have given a non-standard name to the staging bucket.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	FileAssetsBucketName *string `field:"optional" json:"fileAssetsBucketName" yaml:"fileAssetsBucketName"`
	// Name of the ECR repository to hold Docker Image assets.
	//
	// You must supply this if you have given a non-standard name to the ECR repository.
	//
	// The placeholders `${Qualifier}`, `${AWS::AccountId}` and `${AWS::Region}` will
	// be replaced with the values of qualifier and the stack's account and region,
	// respectively.
	ImageAssetsRepositoryName *string `field:"optional" json:"imageAssetsRepositoryName" yaml:"imageAssetsRepositoryName"`
	// Qualifier to disambiguate multiple environments in the same account.
	//
	// You can use this and leave the other naming properties empty if you have deployed
	// the bootstrap environment with standard names but only differnet qualifiers.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

