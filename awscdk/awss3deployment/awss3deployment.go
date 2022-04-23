package awss3deployment

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/awss3deployment/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// `BucketDeployment` populates an S3 bucket with the contents of .zip files from other S3 buckets or from local disk.
//
// Example:
//   var websiteBucket bucket
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
type BucketDeployment interface {
	awscdk.Construct
	// The bucket after the deployment.
	//
	// If you want to reference the destination bucket in another construct and make sure the
	// bucket deployment has happened before the next operation is started, pass the other construct
	// a reference to `deployment.deployedBucket`.
	//
	// Doing this replaces calling `otherResource.node.addDependency(deployment)`.
	// Experimental.
	DeployedBucket() awss3.IBucket
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for BucketDeployment
type jsiiProxy_BucketDeployment struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_BucketDeployment) DeployedBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"deployedBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketDeployment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBucketDeployment_Override(b BucketDeployment, scope constructs.Construct, id *string, props *BucketDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_s3_deployment.BucketDeployment",
		[]interface{}{scope, id, props},
		b,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func BucketDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.BucketDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketDeployment) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BucketDeployment) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (b *jsiiProxy_BucketDeployment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketDeployment) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BucketDeployment) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (b *jsiiProxy_BucketDeployment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for `BucketDeployment`.
//
// Example:
//   var websiteBucket bucket
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
	DestinationBucket awss3.IBucket `json:"destinationBucket" yaml:"destinationBucket"`
	// The sources from which to deploy the contents of this bucket.
	// Experimental.
	Sources *[]ISource `json:"sources" yaml:"sources"`
	// System-defined x-amz-acl metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#canned-acl
	//
	// Experimental.
	AccessControl awss3.BucketAccessControl `json:"accessControl" yaml:"accessControl"`
	// System-defined cache-control metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	CacheControl *[]CacheControl `json:"cacheControl" yaml:"cacheControl"`
	// System-defined cache-disposition metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentDisposition *string `json:"contentDisposition" yaml:"contentDisposition"`
	// System-defined content-encoding metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentEncoding *string `json:"contentEncoding" yaml:"contentEncoding"`
	// System-defined content-language metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentLanguage *string `json:"contentLanguage" yaml:"contentLanguage"`
	// System-defined content-type metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ContentType *string `json:"contentType" yaml:"contentType"`
	// Key prefix in the destination bucket.
	//
	// Must be <=104 characters.
	// Experimental.
	DestinationKeyPrefix *string `json:"destinationKeyPrefix" yaml:"destinationKeyPrefix"`
	// The CloudFront distribution using the destination bucket as an origin.
	//
	// Files in the distribution's edge caches will be invalidated after
	// files are uploaded to the destination bucket.
	// Experimental.
	Distribution awscloudfront.IDistribution `json:"distribution" yaml:"distribution"`
	// The file paths to invalidate in the CloudFront distribution.
	// Experimental.
	DistributionPaths *[]*string `json:"distributionPaths" yaml:"distributionPaths"`
	// The size of the AWS Lambda function’s /tmp directory in MiB.
	// Experimental.
	EphemeralStorageSize awscdk.Size `json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
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
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// System-defined expires metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	Expires awscdk.Expiration `json:"expires" yaml:"expires"`
	// If this is set, matching files or objects will be included with the deployment's sync command.
	//
	// Since all files from the deployment package are included by default, this property
	// is usually leveraged alongside an `exclude` filter.
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/index.html#use-of-exclude-and-include-filters
	//
	// Experimental.
	Include *[]*string `json:"include" yaml:"include"`
	// The number of days that the lambda function's log events are kept in CloudWatch Logs.
	// Experimental.
	LogRetention awslogs.RetentionDays `json:"logRetention" yaml:"logRetention"`
	// The amount of memory (in MiB) to allocate to the AWS Lambda function which replicates the files from the CDK bucket to the destination bucket.
	//
	// If you are deploying large files, you will need to increase this number
	// accordingly.
	// Experimental.
	MemoryLimit *float64 `json:"memoryLimit" yaml:"memoryLimit"`
	// User-defined object metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#UserMetadata
	//
	// Experimental.
	Metadata *UserDefinedObjectMetadata `json:"metadata" yaml:"metadata"`
	// If this is set to false, files in the destination bucket that do not exist in the asset, will NOT be deleted during deployment (create/update).
	// See: https://docs.aws.amazon.com/cli/latest/reference/s3/sync.html
	//
	// Experimental.
	Prune *bool `json:"prune" yaml:"prune"`
	// If this is set to "false", the destination files will be deleted when the resource is deleted or the destination is updated.
	//
	// NOTICE: Configuring this to "false" might have operational implications. Please
	// visit to the package documentation referred below to make sure you fully understand those implications.
	// See: https://github.com/aws/aws-cdk/tree/master/packages/%40aws-cdk/aws-s3-deployment#retain-on-delete
	//
	// Experimental.
	RetainOnDelete *bool `json:"retainOnDelete" yaml:"retainOnDelete"`
	// Execution role associated with this function.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// System-defined x-amz-server-side-encryption metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ServerSideEncryption ServerSideEncryption `json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// System-defined x-amz-server-side-encryption-aws-kms-key-id metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	ServerSideEncryptionAwsKmsKeyId *string `json:"serverSideEncryptionAwsKmsKeyId" yaml:"serverSideEncryptionAwsKmsKeyId"`
	// System-defined x-amz-server-side-encryption-customer-algorithm metadata to be set on all objects in the deployment.
	//
	// Warning: This is not a useful parameter until this bug is fixed: https://github.com/aws/aws-cdk/issues/6080
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html#sse-c-how-to-programmatically-intro
	//
	// Experimental.
	ServerSideEncryptionCustomerAlgorithm *string `json:"serverSideEncryptionCustomerAlgorithm" yaml:"serverSideEncryptionCustomerAlgorithm"`
	// System-defined x-amz-storage-class metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	StorageClass StorageClass `json:"storageClass" yaml:"storageClass"`
	// Mount an EFS file system.
	//
	// Enable this if your assets are large and you encounter disk space errors.
	// Enabling this option will require a VPC to be specified.
	// Experimental.
	UseEfs *bool `json:"useEfs" yaml:"useEfs"`
	// The VPC network to place the deployment lambda handler in.
	//
	// This is required if `useEfs` is set.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Where in the VPC to place the deployment lambda handler.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// System-defined x-amz-website-redirect-location metadata to be set on all objects in the deployment.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
	//
	// Experimental.
	WebsiteRedirectLocation *string `json:"websiteRedirectLocation" yaml:"websiteRedirectLocation"`
}

// Used for HTTP cache-control header, which influences downstream caches.
//
// Example:
//   var destinationBucket bucket
//   s3deploy.NewBucketDeployment(this, jsii.String("BucketDeployment"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(jsii.String("./website"), &assetOptions{
//   			exclude: []*string{
//   				jsii.String("index.html"),
//   			},
//   		}),
//   	},
//   	destinationBucket: destinationBucket,
//   	cacheControl: []cacheControl{
//   		s3deploy.*cacheControl.fromString(jsii.String("max-age=31536000,public,immutable")),
//   	},
//   	prune: jsii.Boolean(false),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("HTMLBucketDeployment"), &bucketDeploymentProps{
//   	sources: []*iSource{
//   		s3deploy.*source.asset(jsii.String("./website"), &assetOptions{
//   			exclude: []*string{
//   				jsii.String("*"),
//   				jsii.String("!index.html"),
//   			},
//   		}),
//   	},
//   	destinationBucket: destinationBucket,
//   	cacheControl: []*cacheControl{
//   		s3deploy.*cacheControl.fromString(jsii.String("max-age=0,no-cache,no-store,must-revalidate")),
//   	},
//   	prune: jsii.Boolean(false),
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type CacheControl interface {
	// The raw cache control setting.
	// Experimental.
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
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
		"monocdk.aws_s3_deployment.CacheControl",
		"sMaxAge",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Bind context for ISources.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3_deployment "github.com/aws/aws-cdk-go/awscdk/aws_s3_deployment"
//
//   var role role
//   deploymentSourceContext := &deploymentSourceContext{
//   	handlerRole: role,
//   }
//
// Experimental.
type DeploymentSourceContext struct {
	// The role for the handler.
	// Experimental.
	HandlerRole awsiam.IRole `json:"handlerRole" yaml:"handlerRole"`
}

// Used for HTTP expires header, which influences downstream caches.
//
// Does NOT influence deletion of the object.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3_deployment "github.com/aws/aws-cdk-go/awscdk/aws_s3_deployment"
//
//   var duration duration
//   expires := s3_deployment.expires.after(duration)
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Deprecated: use core.Expiration
type Expires interface {
	// The raw expiration date expression.
	// Deprecated: use core.Expiration
	Value() interface{}
}

// The jsii proxy struct for Expires
type jsiiProxy_Expires struct {
	_ byte // padding
}

func (j *jsiiProxy_Expires) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Expire once the specified duration has passed since deployment time.
// Deprecated: use core.Expiration
func Expires_After(t awscdk.Duration) Expires {
	_init_.Initialize()

	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"after",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Expire at the specified date.
// Deprecated: use core.Expiration
func Expires_AtDate(d *time.Time) Expires {
	_init_.Initialize()

	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"atDate",
		[]interface{}{d},
		&returns,
	)

	return returns
}

// Expire at the specified timestamp.
// Deprecated: use core.Expiration
func Expires_AtTimestamp(t *float64) Expires {
	_init_.Initialize()

	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"atTimestamp",
		[]interface{}{t},
		&returns,
	)

	return returns
}

// Create an expiration date from a raw date string.
// Deprecated: use core.Expiration
func Expires_FromString(s *string) Expires {
	_init_.Initialize()

	var returns Expires

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Expires",
		"fromString",
		[]interface{}{s},
		&returns,
	)

	return returns
}

// Represents a source for bucket deployments.
// Experimental.
type ISource interface {
	// Binds the source to a bucket deployment.
	// Experimental.
	Bind(scope awscdk.Construct, context *DeploymentSourceContext) *SourceConfig
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISource) Bind(scope awscdk.Construct, context *DeploymentSourceContext) *SourceConfig {
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
//
// Example:
//   websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &bucketProps{
//   	websiteIndexDocument: jsii.String("index.html"),
//   	publicReadAccess: jsii.Boolean(true),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(jsii.String("./website-dist")),
//   	},
//   	destinationBucket: websiteBucket,
//   	destinationKeyPrefix: jsii.String("web/static"),
//   	 // optional prefix in destination bucket
//   	metadata: &userDefinedObjectMetadata{
//   		a: jsii.String("1"),
//   		b: jsii.String("2"),
//   	},
//   	 // user-defined metadata
//
//   	// system-defined metadata
//   	contentType: jsii.String("text/html"),
//   	contentLanguage: jsii.String("en"),
//   	storageClass: s3deploy.storageClass_INTELLIGENT_TIERING,
//   	serverSideEncryption: s3deploy.serverSideEncryption_AES_256,
//   	cacheControl: []cacheControl{
//   		s3deploy.*cacheControl.setPublic(),
//   		s3deploy.*cacheControl.maxAge(duration.hours(jsii.Number(1))),
//   	},
//   	accessControl: s3.bucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type ServerSideEncryption string

const (
	// 'AES256'.
	// Experimental.
	ServerSideEncryption_AES_256 ServerSideEncryption = "AES_256"
	// 'aws:kms'.
	// Experimental.
	ServerSideEncryption_AWS_KMS ServerSideEncryption = "AWS_KMS"
)

// Specifies bucket deployment source.
//
// Usage:
//
//      Source.bucket(bucket, key)
//      Source.asset('/local/path/to/directory')
//      Source.asset('/local/path/to/a/file.zip')
//      Source.data('hello/world/file.txt', 'Hello, world!')
//      Source.data('config.json', { baz: topic.topicArn })
//
// Example:
//   var websiteBucket bucket
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
type Source interface {
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	_ byte // padding
}

// Uses a local asset as the deployment source.
//
// If the local asset is a .zip archive, make sure you trust the
// producer of the archive.
// Experimental.
func Source_Asset(path *string, options *awss3assets.AssetOptions) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"asset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Uses a .zip file stored in an S3 bucket as the source for the destination bucket contents.
//
// Make sure you trust the producer of the archive.
// Experimental.
func Source_Bucket(bucket awss3.IBucket, zipObjectKey *string) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"bucket",
		[]interface{}{bucket, zipObjectKey},
		&returns,
	)

	return returns
}

// Deploys an object with the specified string contents into the bucket.
//
// The
// content can include deploy-time values (such as `snsTopic.topicArn`) that
// will get resolved only during deployment.
//
// To store a JSON object use `Source.jsonData()`.
// Experimental.
func Source_Data(objectKey *string, data *string) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"data",
		[]interface{}{objectKey, data},
		&returns,
	)

	return returns
}

// Deploys an object with the specified JSON object into the bucket.
//
// The
// object can include deploy-time values (such as `snsTopic.topicArn`) that
// will get resolved only during deployment.
// Experimental.
func Source_JsonData(objectKey *string, obj interface{}) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"monocdk.aws_s3_deployment.Source",
		"jsonData",
		[]interface{}{objectKey, obj},
		&returns,
	)

	return returns
}

// Source information.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3 "github.com/aws/aws-cdk-go/awscdk/aws_s3"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3_deployment "github.com/aws/aws-cdk-go/awscdk/aws_s3_deployment"
//
//   var bucket bucket
//   var markers interface{}
//   sourceConfig := &sourceConfig{
//   	bucket: bucket,
//   	zipObjectKey: jsii.String("zipObjectKey"),
//
//   	// the properties below are optional
//   	markers: map[string]interface{}{
//   		"markersKey": markers,
//   	},
//   }
//
// Experimental.
type SourceConfig struct {
	// The source bucket to deploy from.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// An S3 object key in the source bucket that points to a zip file.
	// Experimental.
	ZipObjectKey *string `json:"zipObjectKey" yaml:"zipObjectKey"`
	// A set of markers to substitute in the source content.
	// Experimental.
	Markers *map[string]interface{} `json:"markers" yaml:"markers"`
}

// Storage class used for storing the object.
//
// Example:
//   websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &bucketProps{
//   	websiteIndexDocument: jsii.String("index.html"),
//   	publicReadAccess: jsii.Boolean(true),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(jsii.String("./website-dist")),
//   	},
//   	destinationBucket: websiteBucket,
//   	destinationKeyPrefix: jsii.String("web/static"),
//   	 // optional prefix in destination bucket
//   	metadata: &userDefinedObjectMetadata{
//   		a: jsii.String("1"),
//   		b: jsii.String("2"),
//   	},
//   	 // user-defined metadata
//
//   	// system-defined metadata
//   	contentType: jsii.String("text/html"),
//   	contentLanguage: jsii.String("en"),
//   	storageClass: s3deploy.storageClass_INTELLIGENT_TIERING,
//   	serverSideEncryption: s3deploy.serverSideEncryption_AES_256,
//   	cacheControl: []cacheControl{
//   		s3deploy.*cacheControl.setPublic(),
//   		s3deploy.*cacheControl.maxAge(duration.hours(jsii.Number(1))),
//   	},
//   	accessControl: s3.bucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#SysMetadata
//
// Experimental.
type StorageClass string

const (
	// 'STANDARD'.
	// Experimental.
	StorageClass_STANDARD StorageClass = "STANDARD"
	// 'REDUCED_REDUNDANCY'.
	// Experimental.
	StorageClass_REDUCED_REDUNDANCY StorageClass = "REDUCED_REDUNDANCY"
	// 'STANDARD_IA'.
	// Experimental.
	StorageClass_STANDARD_IA StorageClass = "STANDARD_IA"
	// 'ONEZONE_IA'.
	// Experimental.
	StorageClass_ONEZONE_IA StorageClass = "ONEZONE_IA"
	// 'INTELLIGENT_TIERING'.
	// Experimental.
	StorageClass_INTELLIGENT_TIERING StorageClass = "INTELLIGENT_TIERING"
	// 'GLACIER'.
	// Experimental.
	StorageClass_GLACIER StorageClass = "GLACIER"
	// 'DEEP_ARCHIVE'.
	// Experimental.
	StorageClass_DEEP_ARCHIVE StorageClass = "DEEP_ARCHIVE"
)

// Custom user defined metadata.
//
// Example:
//   websiteBucket := s3.NewBucket(this, jsii.String("WebsiteBucket"), &bucketProps{
//   	websiteIndexDocument: jsii.String("index.html"),
//   	publicReadAccess: jsii.Boolean(true),
//   })
//
//   s3deploy.NewBucketDeployment(this, jsii.String("DeployWebsite"), &bucketDeploymentProps{
//   	sources: []iSource{
//   		s3deploy.source.asset(jsii.String("./website-dist")),
//   	},
//   	destinationBucket: websiteBucket,
//   	destinationKeyPrefix: jsii.String("web/static"),
//   	 // optional prefix in destination bucket
//   	metadata: &userDefinedObjectMetadata{
//   		a: jsii.String("1"),
//   		b: jsii.String("2"),
//   	},
//   	 // user-defined metadata
//
//   	// system-defined metadata
//   	contentType: jsii.String("text/html"),
//   	contentLanguage: jsii.String("en"),
//   	storageClass: s3deploy.storageClass_INTELLIGENT_TIERING,
//   	serverSideEncryption: s3deploy.serverSideEncryption_AES_256,
//   	cacheControl: []cacheControl{
//   		s3deploy.*cacheControl.setPublic(),
//   		s3deploy.*cacheControl.maxAge(duration.hours(jsii.Number(1))),
//   	},
//   	accessControl: s3.bucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
//   })
//
// Experimental.
type UserDefinedObjectMetadata struct {
}

