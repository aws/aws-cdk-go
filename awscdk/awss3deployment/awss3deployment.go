package awss3deployment

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// `BucketDeployment` populates an S3 bucket with the contents of .zip files from other S3 buckets or from local disk.
// Experimental.
type BucketDeployment interface {
	constructs.Construct
	Node() constructs.Node
	ToString() *string
}

// The jsii proxy struct for BucketDeployment
type jsiiProxy_BucketDeployment struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BucketDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewBucketDeployment(scope constructs.Construct, id *string, props *BucketDeploymentProps) BucketDeployment {
	_init_.Initialize()

	j := jsiiProxy_BucketDeployment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBucketDeployment_Override(b BucketDeployment, scope constructs.Construct, id *string, props *BucketDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func BucketDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.BucketDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (b *jsiiProxy_BucketDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `BucketDeployment`.
// Experimental.
type BucketDeploymentProps struct {
	// The S3 bucket to sync the contents of the zip file to.
	// Experimental.
	DestinationBucket awss3.IBucket `json:"destinationBucket"`
	// The sources from which to deploy the contents of this bucket.
	// Experimental.
	Sources *[]ISource `json:"sources"`
	// System-defined cache-control metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	CacheControl *[]CacheControl `json:"cacheControl"`
	// System-defined cache-disposition metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentDisposition *string `json:"contentDisposition"`
	// System-defined content-encoding metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentEncoding *string `json:"contentEncoding"`
	// System-defined content-language metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentLanguage *string `json:"contentLanguage"`
	// System-defined content-type metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentType *string `json:"contentType"`
	// Key prefix in the destination bucket.
	// Experimental.
	DestinationKeyPrefix *string `json:"destinationKeyPrefix"`
	// The CloudFront distribution using the destination bucket as an origin.
	//
	// Files in the distribution's edge caches will be invalidated after
	// files are uploaded to the destination bucket.
	// Experimental.
	Distribution awscloudfront.IDistribution `json:"distribution"`
	// The file paths to invalidate in the CloudFront distribution.
	// Experimental.
	DistributionPaths *[]*string `json:"distributionPaths"`
	// System-defined expires metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	Expires awscdk.Expiration `json:"expires"`
	// The amount of memory (in MiB) to allocate to the AWS Lambda function which replicates the files from the CDK bucket to the destination bucket.
	//
	// If you are deploying large files, you will need to increase this number
	// accordingly.
	// Experimental.
	MemoryLimit *float64 `json:"memoryLimit"`
	// User-defined object metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#UserMetadata
	//
	// Experimental.
	Metadata *UserDefinedObjectMetadata `json:"metadata"`
	// If this is set to false, files in the destination bucket that do not exist in the asset, will NOT be deleted during deployment (create/update).
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/sync.html
	//
	// Experimental.
	Prune *bool `json:"prune"`
	// If this is set to "false", the destination files will be deleted when the resource is deleted or the destination is updated.
	//
	// NOTICE: Configuring this to "false" might have operational implications. Please
	// visit to the package documentation referred below to make sure you fully understand those implications.
	// See: https://github.com/aws/aws-cdk/tree/master/packages/%40aws-cdk/aws-s3-deployment#retain-on-delete
	//
	// Experimental.
	RetainOnDelete *bool `json:"retainOnDelete"`
	// Execution role associated with this function.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// System-defined x-amz-server-side-encryption metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ServerSideEncryption ServerSideEncryption `json:"serverSideEncryption"`
	// System-defined x-amz-server-side-encryption-aws-kms-key-id metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ServerSideEncryptionAwsKmsKeyId *string `json:"serverSideEncryptionAwsKmsKeyId"`
	// System-defined x-amz-server-side-encryption-customer-algorithm metadata to be set on all objects in the deployment.
	//
	// Warning: This is not a useful parameter until this bug is fixed: https://github.com/aws/aws-cdk/issues/6080
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html#sse-c-how-to-programmatically-intro
	//
	// Experimental.
	ServerSideEncryptionCustomerAlgorithm *string `json:"serverSideEncryptionCustomerAlgorithm"`
	// System-defined x-amz-storage-class metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	StorageClass StorageClass `json:"storageClass"`
	// The VPC network to place the deployment lambda handler in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Where in the VPC to place the deployment lambda handler.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// System-defined x-amz-website-redirect-location metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	WebsiteRedirectLocation *string `json:"websiteRedirectLocation"`
}

// Used for HTTP cache-control header, which influences downstream caches.
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type CacheControl interface {
	Value() interface{}
}

// The jsii proxy struct for CacheControl
type jsiiProxy_CacheControl struct {
	_ byte // padding
}

func (j *jsiiProxy_CacheControl) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Constructs a custom cache control key from the literal value.
// Experimental.
func CacheControl_FromString(s *string) CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Sets 'max-age=<duration-in-seconds>'.
// Experimental.
func CacheControl_MaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"maxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Sets 'must-revalidate'.
// Experimental.
func CacheControl_MustRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"mustRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'no-cache'.
// Experimental.
func CacheControl_NoCache() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"noCache",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'no-transform'.
// Experimental.
func CacheControl_NoTransform() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"noTransform",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'proxy-revalidate'.
// Experimental.
func CacheControl_ProxyRevalidate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"proxyRevalidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'private'.
// Experimental.
func CacheControl_SetPrivate() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"setPrivate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 'public'.
// Experimental.
func CacheControl_SetPublic() CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"setPublic",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Sets 's-maxage=<duration-in-seconds>'.
// Experimental.
func CacheControl_SMaxAge(t awscdk.Duration) CacheControl {
	_init_.Initialize()

	var returns CacheControl

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.CacheControl",
		"sMaxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Bind context for ISources.
// Experimental.
type DeploymentSourceContext struct {
	// The role for the handler.
	// Experimental.
	HandlerRole awsiam.IRole `json:"handlerRole"`
}

// Represents a source for bucket deployments.
// Experimental.
type ISource interface {
	// Binds the source to a bucket deployment.
	// Experimental.
	Bind(scope constructs.Construct, context *DeploymentSourceContext) *SourceConfig
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISource) Bind(scope constructs.Construct, context *DeploymentSourceContext) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

// Indicates whether server-side encryption is enabled for the object, and whether that encryption is from the AWS Key Management Service (AWS KMS) or from Amazon S3 managed encryption (SSE-S3).
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type ServerSideEncryption string

const (
	ServerSideEncryption_AES_256 ServerSideEncryption = "AES_256"
	ServerSideEncryption_AWS_KMS ServerSideEncryption = "AWS_KMS"
)

// Specifies bucket deployment source.
//
// Usage:
//
//      Source.bucket(bucket, key)
//      Source.asset('/local/path/to/directory')
//      Source.asset('/local/path/to/a/file.zip')
// Experimental.
type Source interface {
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	_ byte // padding
}

// Uses a local asset as the deployment source.
// Experimental.
func Source_Asset(path *string, options *awss3assets.AssetOptions) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.Source",
		"asset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Uses a .zip file stored in an S3 bucket as the source for the destination bucket contents.
// Experimental.
func Source_Bucket(bucket awss3.IBucket, zipObjectKey *string) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_deployment.Source",
		"bucket",
		[]interface{}{bucket, zipObjectKey},
		&returns,
	)

	return returns
}

// Source information.
// Experimental.
type SourceConfig struct {
	// The source bucket to deploy from.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
	// An S3 object key in the source bucket that points to a zip file.
	// Experimental.
	ZipObjectKey *string `json:"zipObjectKey"`
}

// Storage class used for storing the object.
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type StorageClass string

const (
	StorageClass_STANDARD StorageClass = "STANDARD"
	StorageClass_REDUCED_REDUNDANCY StorageClass = "REDUCED_REDUNDANCY"
	StorageClass_STANDARD_IA StorageClass = "STANDARD_IA"
	StorageClass_ONEZONE_IA StorageClass = "ONEZONE_IA"
	StorageClass_INTELLIGENT_TIERING StorageClass = "INTELLIGENT_TIERING"
	StorageClass_GLACIER StorageClass = "GLACIER"
	StorageClass_DEEP_ARCHIVE StorageClass = "DEEP_ARCHIVE"
)

// Custom user defined metadata.
// Experimental.
type UserDefinedObjectMetadata struct {
}

