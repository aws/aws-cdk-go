package awss3deployment

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for `BucketDeployment`.
//
// Example:
//   var destinationBucket bucket
//
//
//   deployment := s3deploy.NewBucketDeployment(this, jsii.String("DeployFiles"), &BucketDeploymentProps{
//   	Sources: []iSource{
//   		s3deploy.Source_Asset(path.join(__dirname, jsii.String("source-files"))),
//   	},
//   	DestinationBucket: DestinationBucket,
//   })
//
//   deployment.HandlerRole.AddToPolicy(
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("kms:Decrypt"),
//   		jsii.String("kms:DescribeKey"),
//   	},
//   	Effect: iam.Effect_ALLOW,
//   	Resources: []*string{
//   		jsii.String("<encryption key ARN>"),
//   	},
//   }))
//
type BucketDeploymentProps struct {
	// The S3 bucket to sync the contents of the zip file to.
	DestinationBucket awss3.IBucket `field:"required" json:"destinationBucket" yaml:"destinationBucket"`
	// The sources from which to deploy the contents of this bucket.
	Sources *[]ISource `field:"required" json:"sources" yaml:"sources"`
	// System-defined x-amz-acl metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
	//
	// Default: - Not set.
	//
	AccessControl awss3.BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// System-defined cache-control metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Not set.
	//
	CacheControl *[]CacheControl `field:"optional" json:"cacheControl" yaml:"cacheControl"`
	// System-defined cache-disposition metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Not set.
	//
	ContentDisposition *string `field:"optional" json:"contentDisposition" yaml:"contentDisposition"`
	// System-defined content-encoding metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Not set.
	//
	ContentEncoding *string `field:"optional" json:"contentEncoding" yaml:"contentEncoding"`
	// System-defined content-language metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Not set.
	//
	ContentLanguage *string `field:"optional" json:"contentLanguage" yaml:"contentLanguage"`
	// System-defined content-type metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Not set.
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Key prefix in the destination bucket.
	//
	// Must be <=104 characters.
	// Default: "/" (unzip to root of the destination bucket).
	//
	DestinationKeyPrefix *string `field:"optional" json:"destinationKeyPrefix" yaml:"destinationKeyPrefix"`
	// The CloudFront distribution using the destination bucket as an origin.
	//
	// Files in the distribution's edge caches will be invalidated after
	// files are uploaded to the destination bucket.
	// Default: - No invalidation occurs.
	//
	Distribution awscloudfront.IDistribution `field:"optional" json:"distribution" yaml:"distribution"`
	// The file paths to invalidate in the CloudFront distribution.
	// Default: - All files under the destination bucket key prefix will be invalidated.
	//
	DistributionPaths *[]*string `field:"optional" json:"distributionPaths" yaml:"distributionPaths"`
	// The size of the AWS Lambda function’s /tmp directory in MiB.
	// Default: 512 MiB.
	//
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
	// Default: - No exclude filters are used.
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// System-defined expires metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - The objects in the distribution will not expire.
	//
	Expires awscdk.Expiration `field:"optional" json:"expires" yaml:"expires"`
	// If this is set, the zip file will be synced to the destination S3 bucket and extracted.
	//
	// If false, the file will remain zipped in the destination bucket.
	// Default: true.
	//
	Extract *bool `field:"optional" json:"extract" yaml:"extract"`
	// If this is set, matching files or objects will be included with the deployment's sync command.
	//
	// Since all files from the deployment package are included by default, this property
	// is usually leveraged alongside an `exclude` filter.
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/index.html#use-of-exclude-and-include-filters
	//
	// Default: - No include filters are used and all files are included with the sync command.
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
	// The Log Group used for logging of events emitted by the custom resource's lambda function.
	//
	// Providing a user-controlled log group was rolled out to commercial regions on 2023-11-16.
	// If you are deploying to another type of region, please check regional availability first.
	// Default: - a default log group created by AWS Lambda.
	//
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days that the lambda function's log events are kept in CloudWatch Logs.
	//
	// This is a legacy API and we strongly recommend you migrate to `logGroup` if you can.
	// `logGroup` allows you to create a fully customizable log group and instruct the Lambda function to send logs to it.
	// Default: logs.RetentionDays.INFINITE
	//
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The amount of memory (in MiB) to allocate to the AWS Lambda function which replicates the files from the CDK bucket to the destination bucket.
	//
	// If you are deploying large files, you will need to increase this number
	// accordingly.
	// Default: 128.
	//
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// User-defined object metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#UserMetadata
	//
	// Default: - No user metadata is set.
	//
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// If this is set to false, files in the destination bucket that do not exist in the asset, will NOT be deleted during deployment (create/update).
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/sync.html
	//
	// Default: true.
	//
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// If this is set to "false", the destination files will be deleted when the resource is deleted or the destination is updated.
	//
	// NOTICE: Configuring this to "false" might have operational implications. Please
	// visit to the package documentation referred below to make sure you fully understand those implications.
	// See: https://github.com/aws/aws-cdk/tree/main/packages/aws-cdk-lib/aws-s3-deployment#retain-on-delete
	//
	// Default: true - when resource is deleted/updated, files are retained.
	//
	RetainOnDelete *bool `field:"optional" json:"retainOnDelete" yaml:"retainOnDelete"`
	// Execution role associated with this function.
	// Default: - A role is automatically created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// System-defined x-amz-server-side-encryption metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Server side encryption is not used.
	//
	ServerSideEncryption ServerSideEncryption `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// System-defined x-amz-server-side-encryption-aws-kms-key-id metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Not set.
	//
	ServerSideEncryptionAwsKmsKeyId *string `field:"optional" json:"serverSideEncryptionAwsKmsKeyId" yaml:"serverSideEncryptionAwsKmsKeyId"`
	// System-defined x-amz-server-side-encryption-customer-algorithm metadata to be set on all objects in the deployment.
	//
	// Warning: This is not a useful parameter until this bug is fixed: https://github.com/aws/aws-cdk/issues/6080
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html#sse-c-how-to-programmatically-intro
	//
	// Default: - Not set.
	//
	ServerSideEncryptionCustomerAlgorithm *string `field:"optional" json:"serverSideEncryptionCustomerAlgorithm" yaml:"serverSideEncryptionCustomerAlgorithm"`
	// If set to true, uploads will precompute the value of `x-amz-content-sha256` and include it in the signed S3 request headers.
	// Default: - `x-amz-content-sha256` will not be computed.
	//
	SignContent *bool `field:"optional" json:"signContent" yaml:"signContent"`
	// System-defined x-amz-storage-class metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - Default storage-class for the bucket is used.
	//
	StorageClass StorageClass `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Mount an EFS file system.
	//
	// Enable this if your assets are large and you encounter disk space errors.
	// Enabling this option will require a VPC to be specified.
	// Default: - No EFS. Lambda has access only to 512MB of disk space.
	//
	UseEfs *bool `field:"optional" json:"useEfs" yaml:"useEfs"`
	// The VPC network to place the deployment lambda handler in.
	//
	// This is required if `useEfs` is set.
	// Default: None.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where in the VPC to place the deployment lambda handler.
	//
	// Only used if 'vpc' is supplied.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// System-defined x-amz-website-redirect-location metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Default: - No website redirection.
	//
	WebsiteRedirectLocation *string `field:"optional" json:"websiteRedirectLocation" yaml:"websiteRedirectLocation"`
}

