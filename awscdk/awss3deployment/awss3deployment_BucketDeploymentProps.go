package awss3deployment

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Properties for `BucketDeployment`.
//
// Example:
//   var websiteBucket bucket
//
//
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(path.join(__dirname, jsii.String("my-website"))),
//   	},
//   	destinationBucket: websiteBucket,
//   })
//
//   NewConstructThatReadsFromTheBucket(this, jsii.String("Consumer"), map[string]iBucket{
//   	// Use 'deployment.deployedBucket' instead of 'websiteBucket' here
//   	"bucket": deployment.deployedBucket,
//   })
//
// Experimental.
type BucketDeploymentProps struct {
	// The S3 bucket to sync the contents of the zip file to.
	// Experimental.
	DestinationBucket awss3.IBucket `field:"required" json:"destinationBucket" yaml:"destinationBucket"`
	// The sources from which to deploy the contents of this bucket.
	// Experimental.
	Sources *[]ISource `field:"required" json:"sources" yaml:"sources"`
	// System-defined x-amz-acl metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
	//
	// Experimental.
	AccessControl awss3.BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// System-defined cache-control metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	CacheControl *[]CacheControl `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// System-defined cache-disposition metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// System-defined content-encoding metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// System-defined content-language metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// System-defined content-type metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Key prefix in the destination bucket.
	//
	// Must be <=104 characters.
	// Experimental.
	DestinationKeyPrefix *string `field:"optional" json:"destinationKeyPrefix" yaml:"destinationKeyPrefix"`
	// The CloudFront distribution using the destination bucket as an origin.
	//
	// Files in the distribution's edge caches will be invalidated after
	// files are uploaded to the destination bucket.
	// Experimental.
	Distribution awscloudfront.IDistribution `field:"optional" json:"distribution" yaml:"distribution"`
	// The file paths to invalidate in the CloudFront distribution.
	// Experimental.
	DistributionPaths *[]*string `field:"optional" json:"distributionPaths" yaml:"distributionPaths"`
	// The size of the AWS Lambda function’s /tmp directory in MiB.
	// Experimental.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// If this is set, matching files or objects will be excluded from the deployment's sync command.
	//
	// This can be used to exclude a file from being pruned in the destination bucket.
	//
	// If you want to just exclude files from the deployment package (which excludes these files
	// evaluated when invalidating the asset), you should leverage the `exclude` property of
	// `AssetOptions` when defining your source.
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/index.html#use-of-exclude-and-include-filters
	//
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// System-defined expires metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	Expires awscdk.Expiration `field:"optional" json:"expires" yaml:"expires"`
	// If this is set, matching files or objects will be included with the deployment's sync command.
	//
	// Since all files from the deployment package are included by default, this property
	// is usually leveraged alongside an `exclude` filter.
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/index.html#use-of-exclude-and-include-filters
	//
	// Experimental.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
	// The number of days that the lambda function's log events are kept in CloudWatch Logs.
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The amount of memory (in MiB) to allocate to the AWS Lambda function which replicates the files from the CDK bucket to the destination bucket.
	//
	// If you are deploying large files, you will need to increase this number
	// accordingly.
	// Experimental.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// User-defined object metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#UserMetadata
	//
	// Experimental.
	Metadata *UserDefinedObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If this is set to false, files in the destination bucket that do not exist in the asset, will NOT be deleted during deployment (create/update).
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/sync.html
	//
	// Experimental.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// If this is set to "false", the destination files will be deleted when the resource is deleted or the destination is updated.
	//
	// NOTICE: Configuring this to "false" might have operational implications. Please
	// visit to the package documentation referred below to make sure you fully understand those implications.
	// See: https://github.com/aws/aws-cdk/tree/master/packages/%40aws-cdk/aws-s3-deployment#retain-on-delete
	//
	// Experimental.
	RetainOnDelete *bool `field:"optional" json:"retainOnDelete" yaml:"retainOnDelete"`
	// Execution role associated with this function.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// System-defined x-amz-server-side-encryption metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ServerSideEncryption ServerSideEncryption `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// System-defined x-amz-server-side-encryption-aws-kms-key-id metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ServerSideEncryptionAwsKmsKeyId *string `field:"optional" json:"serverSideEncryptionAwsKmsKeyId" yaml:"serverSideEncryptionAwsKmsKeyId"`
	// System-defined x-amz-server-side-encryption-customer-algorithm metadata to be set on all objects in the deployment.
	//
	// Warning: This is not a useful parameter until this bug is fixed: https://github.com/aws/aws-cdk/issues/6080
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html#sse-c-how-to-programmatically-intro
	//
	// Experimental.
	ServerSideEncryptionCustomerAlgorithm *string `field:"optional" json:"serverSideEncryptionCustomerAlgorithm" yaml:"serverSideEncryptionCustomerAlgorithm"`
	// System-defined x-amz-storage-class metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	StorageClass StorageClass `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Mount an EFS file system.
	//
	// Enable this if your assets are large and you encounter disk space errors.
	// Enabling this option will require a VPC to be specified.
	// Experimental.
	UseEfs *bool `field:"optional" json:"useEfs" yaml:"useEfs"`
	// The VPC network to place the deployment lambda handler in.
	//
	// This is required if `useEfs` is set.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where in the VPC to place the deployment lambda handler.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// System-defined x-amz-website-redirect-location metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	WebsiteRedirectLocation *string `field:"optional" json:"websiteRedirectLocation" yaml:"websiteRedirectLocation"`
}

