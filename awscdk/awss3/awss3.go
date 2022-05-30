package awss3

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &bucketProps{
//   	blockPublicAccess: s3.blockPublicAccess_BLOCK_ALL(),
//   })
//
type BlockPublicAccess interface {
	BlockPublicAcls() *bool
	SetBlockPublicAcls(val *bool)
	BlockPublicPolicy() *bool
	SetBlockPublicPolicy(val *bool)
	IgnorePublicAcls() *bool
	SetIgnorePublicAcls(val *bool)
	RestrictPublicBuckets() *bool
	SetRestrictPublicBuckets(val *bool)
}

// The jsii proxy struct for BlockPublicAccess
type jsiiProxy_BlockPublicAccess struct {
	_ byte // padding
}

func (j *jsiiProxy_BlockPublicAccess) BlockPublicAcls() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"blockPublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockPublicAccess) BlockPublicPolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"blockPublicPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockPublicAccess) IgnorePublicAcls() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"ignorePublicAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockPublicAccess) RestrictPublicBuckets() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"restrictPublicBuckets",
		&returns,
	)
	return returns
}


func NewBlockPublicAccess(options *BlockPublicAccessOptions) BlockPublicAccess {
	_init_.Initialize()

	j := jsiiProxy_BlockPublicAccess{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewBlockPublicAccess_Override(b BlockPublicAccess, options *BlockPublicAccessOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		[]interface{}{options},
		b,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetBlockPublicAcls(val *bool) {
	_jsii_.Set(
		j,
		"blockPublicAcls",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetBlockPublicPolicy(val *bool) {
	_jsii_.Set(
		j,
		"blockPublicPolicy",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetIgnorePublicAcls(val *bool) {
	_jsii_.Set(
		j,
		"ignorePublicAcls",
		val,
	)
}

func (j *jsiiProxy_BlockPublicAccess) SetRestrictPublicBuckets(val *bool) {
	_jsii_.Set(
		j,
		"restrictPublicBuckets",
		val,
	)
}

func BlockPublicAccess_BLOCK_ACLS() BlockPublicAccess {
	_init_.Initialize()
	var returns BlockPublicAccess
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		"BLOCK_ACLS",
		&returns,
	)
	return returns
}

func BlockPublicAccess_BLOCK_ALL() BlockPublicAccess {
	_init_.Initialize()
	var returns BlockPublicAccess
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.BlockPublicAccess",
		"BLOCK_ALL",
		&returns,
	)
	return returns
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBlockedBucket"), &bucketProps{
//   	blockPublicAccess: s3.NewBlockPublicAccess(&blockPublicAccessOptions{
//   		blockPublicPolicy: jsii.Boolean(true),
//   	}),
//   })
//
type BlockPublicAccessOptions struct {
	// Whether to block public ACLs.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	BlockPublicAcls *bool `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Whether to block public policy.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	BlockPublicPolicy *bool `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Whether to ignore public ACLs.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	IgnorePublicAcls *bool `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Whether to restrict public access.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-options
	//
	RestrictPublicBuckets *bool `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

// An S3 bucket with associated policy objects.
//
// This bucket does not yet have all features that exposed by the underlying
// BucketResource.
//
// Example:
//   var ecrRepository repository
//
//
//   codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	environment: &buildEnvironment{
//   		buildImage: codebuild.windowsBuildImage.fromEcrRepository(ecrRepository, jsii.String("v1.0"), codebuild.windowsImageType_SERVER_2019),
//   		// optional certificate to include in the build image
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("my-bucket")),
//   			objectKey: jsii.String("path/to/cert.pem"),
//   		},
//   	},
//   })
//
type Bucket interface {
	BucketBase
	// Indicates if a bucket resource policy should automatically created upon the first call to `addToResourcePolicy`.
	AutoCreatePolicy() *bool
	SetAutoCreatePolicy(val *bool)
	// The ARN of the bucket.
	BucketArn() *string
	// The IPv4 DNS name of the specified bucket.
	BucketDomainName() *string
	// The IPv6 DNS name of the specified bucket.
	BucketDualStackDomainName() *string
	// The name of the bucket.
	BucketName() *string
	// The regional domain name of the specified bucket.
	BucketRegionalDomainName() *string
	// The Domain name of the static website.
	BucketWebsiteDomainName() *string
	// The URL of the static website.
	BucketWebsiteUrl() *string
	// Whether to disallow public access.
	DisallowPublicAccess() *bool
	SetDisallowPublicAccess(val *bool)
	// Optional KMS encryption key associated with this bucket.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// If this bucket has been configured for static website hosting.
	IsWebsite() *bool
	// The tree node.
	Node() constructs.Node
	NotificationsHandlerRole() awsiam.IRole
	SetNotificationsHandlerRole(val awsiam.IRole)
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The resource policy associated with this bucket.
	//
	// If `autoCreatePolicy` is true, a `BucketPolicy` will be created upon the
	// first call to addToResourcePolicy(s).
	Policy() BucketPolicy
	SetPolicy(val BucketPolicy)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a cross-origin access configuration for objects in an Amazon S3 bucket.
	AddCorsRule(rule *CorsRule)
	// Adds a bucket notification event destination.
	//
	// Example:
	//      declare const myLambda: lambda.Function;
	//      const bucket = new s3.Bucket(this, 'MyBucket');
	//      bucket.addEventNotification(s3.EventType.OBJECT_CREATED, new s3n.LambdaDestination(myLambda), {prefix: 'home/myusername/*'});
	//
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
	//
	AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Add an inventory configuration.
	AddInventory(inventory *Inventory)
	// Add a lifecycle rule to the bucket.
	AddLifecycleRule(rule *LifecycleRule)
	// Adds a metrics configuration for the CloudWatch request metrics from the bucket.
	AddMetric(metric *BucketMetrics)
	// Subscribes a destination to receive notifications when an object is created in the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_CREATED)`.
	AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Subscribes a destination to receive notifications when an object is removed from the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_REMOVED)`.
	AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Adds a statement to the resource policy for a principal (i.e. account/role/service) to perform actions on this bucket and/or its contents. Use `bucketArn` and `arnForObjects(keys)` to obtain ARNs for this bucket or objects.
	//
	// Note that the policy statement may or may not be added to the policy.
	// For example, when an `IBucket` is created from an existing bucket,
	// it's not possible to tell whether the bucket already has a policy
	// attached, let alone to re-use that policy to add more statements to it.
	// So it's safest to do nothing in these cases.
	//
	// Returns: metadata about the execution of this method. If the policy
	// was not added, the value of `statementAdded` will be `false`. You
	// should always check this value to make sure that the operation was
	// actually carried out. Otherwise, synthesis and deploy will terminate
	// silently, which may be confusing.
	AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Returns an ARN that represents all objects within the bucket that match the key pattern specified.
	//
	// To represent all keys, specify ``"*"``.
	//
	// If you need to specify a keyPattern with multiple components, concatenate them into a single string, e.g.:
	//
	// arnForObjects(`home/${team}/${user}/*`).
	ArnForObjects(keyPattern *string) *string
	EnableEventBridgeNotification()
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants s3:DeleteObject* permission to an IAM principal for objects in this bucket.
	GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Allows unrestricted access to objects from this bucket.
	//
	// IMPORTANT: This permission allows anyone to perform actions on S3 objects
	// in this bucket, which is useful for when you configure your bucket as a
	// website and want everyone to be able to read objects in the bucket without
	// needing to authenticate.
	//
	// Without arguments, this method will grant read ("s3:GetObject") access to
	// all objects ("*") in the bucket.
	//
	// The method returns the `iam.Grant` object, which can then be modified
	// as needed. For example, you can add a condition that will restrict access only
	// to an IPv4 range like this:
	//
	//      const grant = bucket.grantPublicAccess();
	//      grant.resourceStatement!.addCondition(‘IpAddress’, { “aws:SourceIp”: “54.240.143.0/24” });
	//
	// Note that if this `IBucket` refers to an existing bucket, possibly not
	// managed by CloudFormation, this method will have no effect, since it's
	// impossible to modify the policy of an existing bucket.
	GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant
	// Grants s3:PutObject* and s3:Abort* permissions for this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant the given IAM identity permissions to modify the ACLs of objects in the given Bucket.
	//
	// If your application has the '@aws-cdk/aws-s3:grantWriteWithoutAcl' feature flag set,
	// calling {@link grantWrite} or {@link grantReadWrite} no longer grants permissions to modify the ACLs of the objects;
	// in this case, if you need to modify object ACLs, call this method explicitly.
	GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant
	// Grant read permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grants read/write permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to use the key for
	// encrypt/decrypt will also be granted.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant write permissions to this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object is uploaded to the specified paths (keys) in this bucket using the PutObject API call.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using `onCloudTrailWriteObject` may be preferable.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object at the specified paths (keys) in this bucket are written to.
	//
	// This includes
	// the events PutObject, CopyObject, and CompleteMultipartUpload.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using this method may be preferable to `onCloudTrailPutObject`.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// The S3 URL of an S3 object. For example:.
	//
	// - `s3://onlybucket`
	// - `s3://bucket/key`.
	//
	// Returns: an ObjectS3Url token.
	S3UrlForObject(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
	// The https Transfer Acceleration URL of an S3 object.
	//
	// Specify `dualStack: true` at the options
	// for dual-stack endpoint (connect to the bucket over IPv6). For example:
	//
	// - `https://bucket.s3-accelerate.amazonaws.com`
	// - `https://bucket.s3-accelerate.amazonaws.com/key`
	//
	// Returns: an TransferAccelerationUrl token.
	TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string
	// The https URL of an S3 object. Specify `regional: false` at the options for non-regional URLs. For example:.
	//
	// - `https://s3.us-west-1.amazonaws.com/onlybucket`
	// - `https://s3.us-west-1.amazonaws.com/bucket/key`
	// - `https://s3.cn-north-1.amazonaws.com.cn/china-bucket/mykey`
	//
	// Returns: an ObjectS3Url token.
	UrlForObject(key *string) *string
	// The virtual hosted-style URL of an S3 object. Specify `regional: false` at the options for non-regional URL. For example:.
	//
	// - `https://only-bucket.s3.us-west-1.amazonaws.com`
	// - `https://bucket.s3.us-west-1.amazonaws.com/key`
	// - `https://bucket.s3.amazonaws.com/key`
	// - `https://china-bucket.s3.cn-north-1.amazonaws.com.cn/mykey`
	//
	// Returns: an ObjectS3Url token.
	VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string
}

// The jsii proxy struct for Bucket
type jsiiProxy_Bucket struct {
	jsiiProxy_BucketBase
}

func (j *jsiiProxy_Bucket) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketDualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketWebsiteDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) BucketWebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) DisallowPublicAccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disallowPublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) IsWebsite() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isWebsite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) NotificationsHandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"notificationsHandlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Policy() BucketPolicy {
	var returns BucketPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Bucket) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBucket(scope constructs.Construct, id *string, props *BucketProps) Bucket {
	_init_.Initialize()

	j := jsiiProxy_Bucket{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.Bucket",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBucket_Override(b Bucket, scope constructs.Construct, id *string, props *BucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.Bucket",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_Bucket) SetAutoCreatePolicy(val *bool) {
	_jsii_.Set(
		j,
		"autoCreatePolicy",
		val,
	)
}

func (j *jsiiProxy_Bucket) SetDisallowPublicAccess(val *bool) {
	_jsii_.Set(
		j,
		"disallowPublicAccess",
		val,
	)
}

func (j *jsiiProxy_Bucket) SetNotificationsHandlerRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"notificationsHandlerRole",
		val,
	)
}

func (j *jsiiProxy_Bucket) SetPolicy(val BucketPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func Bucket_FromBucketArn(scope constructs.Construct, id *string, bucketArn *string) IBucket {
	_init_.Initialize()

	var returns IBucket

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"fromBucketArn",
		[]interface{}{scope, id, bucketArn},
		&returns,
	)

	return returns
}

// Creates a Bucket construct that represents an external bucket.
func Bucket_FromBucketAttributes(scope constructs.Construct, id *string, attrs *BucketAttributes) IBucket {
	_init_.Initialize()

	var returns IBucket

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"fromBucketAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

func Bucket_FromBucketName(scope constructs.Construct, id *string, bucketName *string) IBucket {
	_init_.Initialize()

	var returns IBucket

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"fromBucketName",
		[]interface{}{scope, id, bucketName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Bucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Bucket_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.Bucket",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Thrown an exception if the given bucket name is not valid.
func Bucket_ValidateBucketName(physicalName *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_s3.Bucket",
		"validateBucketName",
		[]interface{}{physicalName},
	)
}

func (b *jsiiProxy_Bucket) AddCorsRule(rule *CorsRule) {
	_jsii_.InvokeVoid(
		b,
		"addCorsRule",
		[]interface{}{rule},
	)
}

func (b *jsiiProxy_Bucket) AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{event, dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addEventNotification",
		args,
	)
}

func (b *jsiiProxy_Bucket) AddInventory(inventory *Inventory) {
	_jsii_.InvokeVoid(
		b,
		"addInventory",
		[]interface{}{inventory},
	)
}

func (b *jsiiProxy_Bucket) AddLifecycleRule(rule *LifecycleRule) {
	_jsii_.InvokeVoid(
		b,
		"addLifecycleRule",
		[]interface{}{rule},
	)
}

func (b *jsiiProxy_Bucket) AddMetric(metric *BucketMetrics) {
	_jsii_.InvokeVoid(
		b,
		"addMetric",
		[]interface{}{metric},
	)
}

func (b *jsiiProxy_Bucket) AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addObjectCreatedNotification",
		args,
	)
}

func (b *jsiiProxy_Bucket) AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addObjectRemovedNotification",
		args,
	)
}

func (b *jsiiProxy_Bucket) AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		b,
		"addToResourcePolicy",
		[]interface{}{permission},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_Bucket) ArnForObjects(keyPattern *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"arnForObjects",
		[]interface{}{keyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) EnableEventBridgeNotification() {
	_jsii_.InvokeVoid(
		b,
		"enableEventBridgeNotification",
		nil, // no parameters
	)
}

func (b *jsiiProxy_Bucket) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantDelete",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant {
	args := []interface{}{keyPrefix}
	for _, a := range allowedActions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPublicAccess",
		args,
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPut",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPutAcl",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantRead",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantReadWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailPutObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailWriteObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) S3UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"s3UrlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"transferAccelerationUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"urlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Bucket) VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"virtualHostedUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// Default bucket access control types.
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
//   		s3deploy.*cacheControl.maxAge(awscdk.Duration.hours(jsii.Number(1))),
//   	},
//   	accessControl: s3.bucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html
//
type BucketAccessControl string

const (
	// Owner gets FULL_CONTROL.
	//
	// No one else has access rights.
	BucketAccessControl_PRIVATE BucketAccessControl = "PRIVATE"
	// Owner gets FULL_CONTROL.
	//
	// The AllUsers group gets READ access.
	BucketAccessControl_PUBLIC_READ BucketAccessControl = "PUBLIC_READ"
	// Owner gets FULL_CONTROL.
	//
	// The AllUsers group gets READ and WRITE access.
	// Granting this on a bucket is generally not recommended.
	BucketAccessControl_PUBLIC_READ_WRITE BucketAccessControl = "PUBLIC_READ_WRITE"
	// Owner gets FULL_CONTROL.
	//
	// The AuthenticatedUsers group gets READ access.
	BucketAccessControl_AUTHENTICATED_READ BucketAccessControl = "AUTHENTICATED_READ"
	// The LogDelivery group gets WRITE and READ_ACP permissions on the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html
	//
	BucketAccessControl_LOG_DELIVERY_WRITE BucketAccessControl = "LOG_DELIVERY_WRITE"
	// Object owner gets FULL_CONTROL.
	//
	// Bucket owner gets READ access.
	// If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
	BucketAccessControl_BUCKET_OWNER_READ BucketAccessControl = "BUCKET_OWNER_READ"
	// Both the object owner and the bucket owner get FULL_CONTROL over the object.
	//
	// If you specify this canned ACL when creating a bucket, Amazon S3 ignores it.
	BucketAccessControl_BUCKET_OWNER_FULL_CONTROL BucketAccessControl = "BUCKET_OWNER_FULL_CONTROL"
	// Owner gets FULL_CONTROL.
	//
	// Amazon EC2 gets READ access to GET an Amazon Machine Image (AMI) bundle from Amazon S3.
	BucketAccessControl_AWS_EXEC_READ BucketAccessControl = "AWS_EXEC_READ"
)

// A reference to a bucket outside this stack.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myLambda function
//
//   bucket := s3.bucket.fromBucketAttributes(this, jsii.String("ImportedBucket"), &bucketAttributes{
//   	bucketArn: jsii.String("arn:aws:s3:::my-bucket"),
//   })
//
//   // now you can just call methods on the bucket
//   bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &notificationKeyFilter{
//   	prefix: jsii.String("home/myusername/*"),
//   })
//
type BucketAttributes struct {
	// The account this existing bucket belongs to.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The ARN of the bucket.
	//
	// At least one of bucketArn or bucketName must be
	// defined in order to initialize a bucket ref.
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The domain name of the bucket.
	BucketDomainName *string `field:"optional" json:"bucketDomainName" yaml:"bucketDomainName"`
	// The IPv6 DNS name of the specified bucket.
	BucketDualStackDomainName *string `field:"optional" json:"bucketDualStackDomainName" yaml:"bucketDualStackDomainName"`
	// The name of the bucket.
	//
	// If the underlying value of ARN is a string, the
	// name will be parsed from the ARN. Otherwise, the name is optional, but
	// some features that require the bucket name such as auto-creating a bucket
	// policy, won't work.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The regional domain name of the specified bucket.
	BucketRegionalDomainName *string `field:"optional" json:"bucketRegionalDomainName" yaml:"bucketRegionalDomainName"`
	// The format of the website URL of the bucket.
	//
	// This should be true for
	// regions launched since 2014.
	BucketWebsiteNewUrlFormat *bool `field:"optional" json:"bucketWebsiteNewUrlFormat" yaml:"bucketWebsiteNewUrlFormat"`
	// The website URL of the bucket (if static web hosting is enabled).
	BucketWebsiteUrl *string `field:"optional" json:"bucketWebsiteUrl" yaml:"bucketWebsiteUrl"`
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// If this bucket has been configured for static website hosting.
	IsWebsite *bool `field:"optional" json:"isWebsite" yaml:"isWebsite"`
	// The role to be used by the notifications handler.
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The region this existing bucket is in.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

// Represents an S3 Bucket.
//
// Buckets can be either defined within this stack:
//
//    new Bucket(this, 'MyBucket', { props });
//
// Or imported from an existing bucket:
//
//    Bucket.import(this, 'MyImportedBucket', { bucketArn: ... });
//
// You can also export a bucket and import it into another stack:
//
//    const ref = myBucket.export();
//    Bucket.import(this, 'MyImportedBucket', ref);
type BucketBase interface {
	awscdk.Resource
	IBucket
	// Indicates if a bucket resource policy should automatically created upon the first call to `addToResourcePolicy`.
	AutoCreatePolicy() *bool
	SetAutoCreatePolicy(val *bool)
	// The ARN of the bucket.
	BucketArn() *string
	// The IPv4 DNS name of the specified bucket.
	BucketDomainName() *string
	// The IPv6 DNS name of the specified bucket.
	BucketDualStackDomainName() *string
	// The name of the bucket.
	BucketName() *string
	// The regional domain name of the specified bucket.
	BucketRegionalDomainName() *string
	// The Domain name of the static website.
	BucketWebsiteDomainName() *string
	// The URL of the static website.
	BucketWebsiteUrl() *string
	// Whether to disallow public access.
	DisallowPublicAccess() *bool
	SetDisallowPublicAccess(val *bool)
	// Optional KMS encryption key associated with this bucket.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// If this bucket has been configured for static website hosting.
	IsWebsite() *bool
	// The tree node.
	Node() constructs.Node
	NotificationsHandlerRole() awsiam.IRole
	SetNotificationsHandlerRole(val awsiam.IRole)
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The resource policy associated with this bucket.
	//
	// If `autoCreatePolicy` is true, a `BucketPolicy` will be created upon the
	// first call to addToResourcePolicy(s).
	Policy() BucketPolicy
	SetPolicy(val BucketPolicy)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a bucket notification event destination.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   var myLambda function
	//
	//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
	//   bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &notificationKeyFilter{
	//   	prefix: jsii.String("home/myusername/*"),
	//   })
	//
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
	//
	AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Subscribes a destination to receive notifications when an object is created in the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_CREATED)`.
	AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Subscribes a destination to receive notifications when an object is removed from the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_REMOVED)`.
	AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Adds a statement to the resource policy for a principal (i.e. account/role/service) to perform actions on this bucket and/or its contents. Use `bucketArn` and `arnForObjects(keys)` to obtain ARNs for this bucket or objects.
	//
	// Note that the policy statement may or may not be added to the policy.
	// For example, when an `IBucket` is created from an existing bucket,
	// it's not possible to tell whether the bucket already has a policy
	// attached, let alone to re-use that policy to add more statements to it.
	// So it's safest to do nothing in these cases.
	//
	// Returns: metadata about the execution of this method. If the policy
	// was not added, the value of `statementAdded` will be `false`. You
	// should always check this value to make sure that the operation was
	// actually carried out. Otherwise, synthesis and deploy will terminate
	// silently, which may be confusing.
	AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Returns an ARN that represents all objects within the bucket that match the key pattern specified.
	//
	// To represent all keys, specify ``"*"``.
	//
	// If you need to specify a keyPattern with multiple components, concatenate them into a single string, e.g.:
	//
	// arnForObjects(`home/${team}/${user}/*`).
	ArnForObjects(keyPattern *string) *string
	EnableEventBridgeNotification()
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grants s3:DeleteObject* permission to an IAM principal for objects in this bucket.
	GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Allows unrestricted access to objects from this bucket.
	//
	// IMPORTANT: This permission allows anyone to perform actions on S3 objects
	// in this bucket, which is useful for when you configure your bucket as a
	// website and want everyone to be able to read objects in the bucket without
	// needing to authenticate.
	//
	// Without arguments, this method will grant read ("s3:GetObject") access to
	// all objects ("*") in the bucket.
	//
	// The method returns the `iam.Grant` object, which can then be modified
	// as needed. For example, you can add a condition that will restrict access only
	// to an IPv4 range like this:
	//
	//      const grant = bucket.grantPublicAccess();
	//      grant.resourceStatement!.addCondition(‘IpAddress’, { “aws:SourceIp”: “54.240.143.0/24” });
	//
	// Note that if this `IBucket` refers to an existing bucket, possibly not
	// managed by CloudFormation, this method will have no effect, since it's
	// impossible to modify the policy of an existing bucket.
	GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant
	// Grants s3:PutObject* and s3:Abort* permissions for this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant the given IAM identity permissions to modify the ACLs of objects in the given Bucket.
	//
	// If your application has the '@aws-cdk/aws-s3:grantWriteWithoutAcl' feature flag set,
	// calling {@link grantWrite} or {@link grantReadWrite} no longer grants permissions to modify the ACLs of the objects;
	// in this case, if you need to modify object ACLs, call this method explicitly.
	GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant
	// Grant read permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grants read/write permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to use the key for
	// encrypt/decrypt will also be granted.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant write permissions to this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Define a CloudWatch event that triggers when something happens to this repository.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object is uploaded to the specified paths (keys) in this bucket using the PutObject API call.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using `onCloudTrailWriteObject` may be preferable.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object at the specified paths (keys) in this bucket are written to.
	//
	// This includes
	// the events PutObject, CopyObject, and CompleteMultipartUpload.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using this method may be preferable to `onCloudTrailPutObject`.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// The S3 URL of an S3 object. For example:.
	//
	// - `s3://onlybucket`
	// - `s3://bucket/key`.
	//
	// Returns: an ObjectS3Url token.
	S3UrlForObject(key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
	// The https Transfer Acceleration URL of an S3 object.
	//
	// Specify `dualStack: true` at the options
	// for dual-stack endpoint (connect to the bucket over IPv6). For example:
	//
	// - `https://bucket.s3-accelerate.amazonaws.com`
	// - `https://bucket.s3-accelerate.amazonaws.com/key`
	//
	// Returns: an TransferAccelerationUrl token.
	TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string
	// The https URL of an S3 object. Specify `regional: false` at the options for non-regional URLs. For example:.
	//
	// - `https://s3.us-west-1.amazonaws.com/onlybucket`
	// - `https://s3.us-west-1.amazonaws.com/bucket/key`
	// - `https://s3.cn-north-1.amazonaws.com.cn/china-bucket/mykey`
	//
	// Returns: an ObjectS3Url token.
	UrlForObject(key *string) *string
	// The virtual hosted-style URL of an S3 object. Specify `regional: false` at the options for non-regional URL. For example:.
	//
	// - `https://only-bucket.s3.us-west-1.amazonaws.com`
	// - `https://bucket.s3.us-west-1.amazonaws.com/key`
	// - `https://bucket.s3.amazonaws.com/key`
	// - `https://china-bucket.s3.cn-north-1.amazonaws.com.cn/mykey`
	//
	// Returns: an ObjectS3Url token.
	VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string
}

// The jsii proxy struct for BucketBase
type jsiiProxy_BucketBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IBucket
}

func (j *jsiiProxy_BucketBase) AutoCreatePolicy() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"autoCreatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) BucketDualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) BucketRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) BucketWebsiteDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) BucketWebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) DisallowPublicAccess() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disallowPublicAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) IsWebsite() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isWebsite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) NotificationsHandlerRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"notificationsHandlerRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) Policy() BucketPolicy {
	var returns BucketPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBucketBase_Override(b BucketBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BucketBase",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_BucketBase) SetAutoCreatePolicy(val *bool) {
	_jsii_.Set(
		j,
		"autoCreatePolicy",
		val,
	)
}

func (j *jsiiProxy_BucketBase) SetDisallowPublicAccess(val *bool) {
	_jsii_.Set(
		j,
		"disallowPublicAccess",
		val,
	)
}

func (j *jsiiProxy_BucketBase) SetNotificationsHandlerRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"notificationsHandlerRole",
		val,
	)
}

func (j *jsiiProxy_BucketBase) SetPolicy(val BucketPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func BucketBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BucketBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{event, dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addEventNotification",
		args,
	)
}

func (b *jsiiProxy_BucketBase) AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addObjectCreatedNotification",
		args,
	)
}

func (b *jsiiProxy_BucketBase) AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"addObjectRemovedNotification",
		args,
	)
}

func (b *jsiiProxy_BucketBase) AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		b,
		"addToResourcePolicy",
		[]interface{}{permission},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BucketBase) ArnForObjects(keyPattern *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"arnForObjects",
		[]interface{}{keyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) EnableEventBridgeNotification() {
	_jsii_.InvokeVoid(
		b,
		"enableEventBridgeNotification",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BucketBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantDelete",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant {
	args := []interface{}{keyPrefix}
	for _, a := range allowedActions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPublicAccess",
		args,
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPut",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantPutAcl",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantRead",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantReadWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailPutObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		b,
		"onCloudTrailWriteObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) S3UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"s3UrlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"transferAccelerationUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"urlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketBase) VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"virtualHostedUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// What kind of server-side encryption to apply to this bucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyEncryptedBucket"), &bucketProps{
//   	encryption: s3.bucketEncryption_KMS,
//   })
//
//   // you can access the encryption key:
//   assert(bucket.encryptionKey instanceof kms.key)
//
type BucketEncryption string

const (
	// Objects in the bucket are not encrypted.
	BucketEncryption_UNENCRYPTED BucketEncryption = "UNENCRYPTED"
	// Server-side KMS encryption with a master key managed by KMS.
	BucketEncryption_KMS_MANAGED BucketEncryption = "KMS_MANAGED"
	// Server-side encryption with a master key managed by S3.
	BucketEncryption_S3_MANAGED BucketEncryption = "S3_MANAGED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	BucketEncryption_KMS BucketEncryption = "KMS"
)

// Specifies a metrics configuration for the CloudWatch request metrics from an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tagFilters interface{}
//
//   bucketMetrics := &bucketMetrics{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   	tagFilters: map[string]interface{}{
//   		"tagFiltersKey": tagFilters,
//   	},
//   }
//
type BucketMetrics struct {
	// The ID used to identify the metrics configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The prefix that an object must have to be included in the metrics results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies a list of tag filters to use as a metrics configuration filter.
	//
	// The metrics configuration includes only objects that meet the filter's criteria.
	TagFilters *map[string]interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

// Represents the properties of a notification destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var dependable iDependable
//
//   bucketNotificationDestinationConfig := &bucketNotificationDestinationConfig{
//   	arn: jsii.String("arn"),
//   	type: awscdk.Aws_s3.bucketNotificationDestinationType_LAMBDA,
//
//   	// the properties below are optional
//   	dependencies: []*iDependable{
//   		dependable,
//   	},
//   }
//
type BucketNotificationDestinationConfig struct {
	// The ARN of the destination (i.e. Lambda, SNS, SQS).
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The notification type.
	Type BucketNotificationDestinationType `field:"required" json:"type" yaml:"type"`
	// Any additional dependencies that should be resolved before the bucket notification can be configured (for example, the SNS Topic Policy resource).
	Dependencies *[]constructs.IDependable `field:"optional" json:"dependencies" yaml:"dependencies"`
}

// Supported types of notification destinations.
type BucketNotificationDestinationType string

const (
	BucketNotificationDestinationType_LAMBDA BucketNotificationDestinationType = "LAMBDA"
	BucketNotificationDestinationType_QUEUE BucketNotificationDestinationType = "QUEUE"
	BucketNotificationDestinationType_TOPIC BucketNotificationDestinationType = "TOPIC"
)

// The bucket policy for an Amazon S3 bucket.
//
// Policies define the operations that are allowed on this resource.
//
// You almost never need to define this construct directly.
//
// All AWS resources that support resource policies have a method called
// `addToResourcePolicy()`, which will automatically create a new resource
// policy if one doesn't exist yet, otherwise it will add to the existing
// policy.
//
// Prefer to use `addToResourcePolicy()` instead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   bucketPolicy := awscdk.Aws_s3.NewBucketPolicy(this, jsii.String("MyBucketPolicy"), &bucketPolicyProps{
//   	bucket: bucket,
//
//   	// the properties below are optional
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   })
//
type BucketPolicy interface {
	awscdk.Resource
	// A policy document containing permissions to add to the specified bucket.
	//
	// For more information, see Access Policy Language Overview in the Amazon
	// Simple Storage Service Developer Guide.
	Document() awsiam.PolicyDocument
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Sets the removal policy for the BucketPolicy.
	ApplyRemovalPolicy(removalPolicy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BucketPolicy
type jsiiProxy_BucketPolicy struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_BucketPolicy) Document() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BucketPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBucketPolicy(scope constructs.Construct, id *string, props *BucketPolicyProps) BucketPolicy {
	_init_.Initialize()

	j := jsiiProxy_BucketPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBucketPolicy_Override(b BucketPolicy, scope constructs.Construct, id *string, props *BucketPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		[]interface{}{scope, id, props},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func BucketPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BucketPolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.BucketPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicy) ApplyRemovalPolicy(removalPolicy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{removalPolicy},
	)
}

func (b *jsiiProxy_BucketPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   bucketPolicyProps := &bucketPolicyProps{
//   	bucket: bucket,
//
//   	// the properties below are optional
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   }
//
type BucketPolicyProps struct {
	// The Amazon S3 bucket that the policy applies to.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Policy to apply when the policy is removed from this stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   sourceBucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
//   	versioned: jsii.Boolean(true),
//   })
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewS3SourceAction(&s3SourceActionProps{
//   	actionName: jsii.String("S3Source"),
//   	bucket: sourceBucket,
//   	bucketKey: jsii.String("path/to/file.zip"),
//   	output: sourceOutput,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type BucketProps struct {
	// Specifies a canned ACL that grants predefined permissions to the bucket.
	AccessControl BucketAccessControl `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Whether all objects should be automatically deleted when the bucket is removed from the stack or when the stack is deleted.
	//
	// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
	//
	// **Warning** if you have deployed a bucket with `autoDeleteObjects: true`,
	// switching this to `false` in a CDK version *before* `1.126.0` will lead to
	// all objects in the bucket being deleted. Be sure to update your bucket resources
	// by deploying with CDK version `1.126.0` or later **before** switching this value to `false`.
	AutoDeleteObjects *bool `field:"optional" json:"autoDeleteObjects" yaml:"autoDeleteObjects"`
	// The block public access configuration of this bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
	//
	BlockPublicAccess BlockPublicAccess `field:"optional" json:"blockPublicAccess" yaml:"blockPublicAccess"`
	// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
	//
	// Only relevant, when Encryption is set to {@link BucketEncryption.KMS}
	BucketKeyEnabled *bool `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Physical name of this bucket.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The CORS configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-cors.html
	//
	Cors *[]*CorsRule `field:"optional" json:"cors" yaml:"cors"`
	// The kind of server-side encryption to apply to this bucket.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	Encryption BucketEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The 'encryption' property must be either not specified or set to "Kms".
	// An error will be emitted if encryption is set to "Unencrypted" or
	// "Managed".
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Enforces SSL for requests.
	//
	// S3.5 of the AWS Foundational Security Best Practices Regarding S3.
	// See: https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-ssl-requests-only.html
	//
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// Whether this bucket should send notifications to Amazon EventBridge or not.
	EventBridgeEnabled *bool `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
	// Inteligent Tiering Configurations.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/intelligent-tiering.html
	//
	IntelligentTieringConfigurations *[]*IntelligentTieringConfiguration `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// The inventory configuration of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
	//
	Inventories *[]*Inventory `field:"optional" json:"inventories" yaml:"inventories"`
	// Rules that define how Amazon S3 manages objects during their lifetime.
	LifecycleRules *[]*LifecycleRule `field:"optional" json:"lifecycleRules" yaml:"lifecycleRules"`
	// The metrics configuration of this bucket.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-metricsconfiguration.html
	//
	Metrics *[]*BucketMetrics `field:"optional" json:"metrics" yaml:"metrics"`
	// The role to be used by the notifications handler.
	NotificationsHandlerRole awsiam.IRole `field:"optional" json:"notificationsHandlerRole" yaml:"notificationsHandlerRole"`
	// The objectOwnership of the bucket.
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
	//
	ObjectOwnership ObjectOwnership `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
	// Grants public read access to all objects in the bucket.
	//
	// Similar to calling `bucket.grantPublicAccess()`
	PublicReadAccess *bool `field:"optional" json:"publicReadAccess" yaml:"publicReadAccess"`
	// Policy to apply when the bucket is removed from this stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Destination bucket for the server access logs.
	ServerAccessLogsBucket IBucket `field:"optional" json:"serverAccessLogsBucket" yaml:"serverAccessLogsBucket"`
	// Optional log file prefix to use for the bucket's access logs.
	//
	// If defined without "serverAccessLogsBucket", enables access logs to current bucket with this prefix.
	ServerAccessLogsPrefix *string `field:"optional" json:"serverAccessLogsPrefix" yaml:"serverAccessLogsPrefix"`
	// Whether this bucket should have transfer acceleration turned on or not.
	TransferAcceleration *bool `field:"optional" json:"transferAcceleration" yaml:"transferAcceleration"`
	// Whether this bucket should have versioning turned on or not.
	Versioned *bool `field:"optional" json:"versioned" yaml:"versioned"`
	// The name of the error document (e.g. "404.html") for the website. `websiteIndexDocument` must also be set if this is set.
	WebsiteErrorDocument *string `field:"optional" json:"websiteErrorDocument" yaml:"websiteErrorDocument"`
	// The name of the index document (e.g. "index.html") for the website. Enables static website hosting for this bucket.
	WebsiteIndexDocument *string `field:"optional" json:"websiteIndexDocument" yaml:"websiteIndexDocument"`
	// Specifies the redirect behavior of all requests to a website endpoint of a bucket.
	//
	// If you specify this property, you can't specify "websiteIndexDocument", "websiteErrorDocument" nor , "websiteRoutingRules".
	WebsiteRedirect *RedirectTarget `field:"optional" json:"websiteRedirect" yaml:"websiteRedirect"`
	// Rules that define when a redirect is applied and the redirect behavior.
	WebsiteRoutingRules *[]*RoutingRule `field:"optional" json:"websiteRoutingRules" yaml:"websiteRoutingRules"`
}

// A CloudFormation `AWS::S3::AccessPoint`.
//
// The AWS::S3::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//   var policyStatus interface{}
//
//   cfnAccessPoint := awscdk.Aws_s3.NewCfnAccessPoint(this, jsii.String("MyCfnAccessPoint"), &cfnAccessPointProps{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	policy: policy,
//   	policyStatus: policyStatus,
//   	publicAccessBlockConfiguration: &publicAccessBlockConfigurationProperty{
//   		blockPublicAcls: jsii.Boolean(false),
//   		blockPublicPolicy: jsii.Boolean(false),
//   		ignorePublicAcls: jsii.Boolean(false),
//   		restrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		vpcId: jsii.String("vpcId"),
//   	},
//   })
//
type CfnAccessPoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The alias for this access point.
	AttrAlias() *string
	// This property contains the details of the ARN for the access point.
	AttrArn() *string
	// The name of this access point.
	AttrName() *string
	// Indicates whether this access point allows access from the internet.
	//
	// If `VpcConfiguration` is specified for this access point, then `NetworkOrigin` is `VPC` , and the access point doesn't allow access from the internet. Otherwise, `NetworkOrigin` is `Internet` , and the access point allows access from the internet, subject to the access point and bucket access policies.
	//
	// *Allowed values* : `VPC` | `Internet`.
	AttrNetworkOrigin() *string
	// The name of the bucket associated with this access point.
	Bucket() *string
	SetBucket(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of this access point.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The access point policy associated with this access point.
	Policy() interface{}
	SetPolicy(val interface{})
	// The container element for a bucket's policy status.
	PolicyStatus() interface{}
	SetPolicyStatus(val interface{})
	// The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
	PublicAccessBlockConfiguration() interface{}
	SetPublicAccessBlockConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// The Virtual Private Cloud (VPC) configuration for this access point, if one exists.
	VpcConfiguration() interface{}
	SetVpcConfiguration(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnAccessPoint
type jsiiProxy_CfnAccessPoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAccessPoint) AttrAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) AttrNetworkOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrNetworkOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) Policy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) PolicyStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) PublicAccessBlockConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessBlockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessPoint) VpcConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::S3::AccessPoint`.
func NewCfnAccessPoint(scope constructs.Construct, id *string, props *CfnAccessPointProps) CfnAccessPoint {
	_init_.Initialize()

	j := jsiiProxy_CfnAccessPoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnAccessPoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::S3::AccessPoint`.
func NewCfnAccessPoint_Override(c CfnAccessPoint, scope constructs.Construct, id *string, props *CfnAccessPointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnAccessPoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAccessPoint) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_CfnAccessPoint) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAccessPoint) SetPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_CfnAccessPoint) SetPolicyStatus(val interface{}) {
	_jsii_.Set(
		j,
		"policyStatus",
		val,
	)
}

func (j *jsiiProxy_CfnAccessPoint) SetPublicAccessBlockConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"publicAccessBlockConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAccessPoint) SetVpcConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"vpcConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAccessPoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnAccessPoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAccessPoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnAccessPoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnAccessPoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnAccessPoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessPoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.CfnAccessPoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessPoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAccessPoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAccessPoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAccessPoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAccessPoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAccessPoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAccessPoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAccessPoint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPoint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAccessPoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAccessPoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAccessPoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket.
//
// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessBlockConfigurationProperty := &publicAccessBlockConfigurationProperty{
//   	blockPublicAcls: jsii.Boolean(false),
//   	blockPublicPolicy: jsii.Boolean(false),
//   	ignorePublicAcls: jsii.Boolean(false),
//   	restrictPublicBuckets: jsii.Boolean(false),
//   }
//
type CfnAccessPoint_PublicAccessBlockConfigurationProperty struct {
	// Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes the following behavior:
	//
	// - PUT Bucket ACL and PUT Object ACL calls fail if the specified ACL is public.
	// - PUT Object calls fail if the request includes a public ACL.
	// - PUT Bucket calls fail if the request includes a public ACL.
	//
	// Enabling this setting doesn't affect existing policies or ACLs.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Specifies whether Amazon S3 should block public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	//
	// Enabling this setting doesn't affect existing bucket policies.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.
	//
	// Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Specifies whether Amazon S3 should restrict public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` restricts access to this bucket to only AWS service principals and authorized users within this account if the bucket has a public policy.
	//
	// Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

// The Virtual Private Cloud (VPC) configuration for this access point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigurationProperty := &vpcConfigurationProperty{
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnAccessPoint_VpcConfigurationProperty struct {
	// If this field is specified, the access point will only allow connections from the specified VPC ID.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

// Properties for defining a `CfnAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//   var policyStatus interface{}
//
//   cfnAccessPointProps := &cfnAccessPointProps{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	policy: policy,
//   	policyStatus: policyStatus,
//   	publicAccessBlockConfiguration: &publicAccessBlockConfigurationProperty{
//   		blockPublicAcls: jsii.Boolean(false),
//   		blockPublicPolicy: jsii.Boolean(false),
//   		ignorePublicAcls: jsii.Boolean(false),
//   		restrictPublicBuckets: jsii.Boolean(false),
//   	},
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		vpcId: jsii.String("vpcId"),
//   	},
//   }
//
type CfnAccessPointProps struct {
	// The name of the bucket associated with this access point.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of this access point.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The access point policy associated with this access point.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The container element for a bucket's policy status.
	PolicyStatus interface{} `field:"optional" json:"policyStatus" yaml:"policyStatus"`
	// The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// The Virtual Private Cloud (VPC) configuration for this access point, if one exists.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

// A CloudFormation `AWS::S3::Bucket`.
//
// The `AWS::S3::Bucket` resource creates an Amazon S3 bucket in the same AWS Region where you create the AWS CloudFormation stack.
//
// To control how AWS CloudFormation handles the bucket when the stack is deleted, you can set a deletion policy for your bucket. You can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
//
// > You can only delete empty buckets. Deletion fails for buckets that have contents.
//
// Example:
//   var cfnTemplate cfnInclude
//
//   cfnBucket := cfnTemplate.getResource(jsii.String("Bucket")).(cfnBucket)
//
//   role := iam.NewRole(this, jsii.String("Role"), &roleProps{
//   	assumedBy: iam.NewAnyPrincipal(),
//   })
//   role.addToPolicy(iam.NewPolicyStatement(&policyStatementProps{
//   	actions: []*string{
//   		jsii.String("s3:*"),
//   	},
//   	resources: []*string{
//   		cfnBucket.attrArn,
//   	},
//   }))
//
type CfnBucket interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Configures the transfer acceleration state for an Amazon S3 bucket.
	//
	// For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide* .
	AccelerateConfiguration() interface{}
	SetAccelerateConfiguration(val interface{})
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	//
	// For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide* .
	//
	// Be aware that the syntax for this property differs from the information provided in the *Amazon S3 User Guide* . The AccessControl property is case-sensitive and must be one of the following values: Private, PublicRead, PublicReadWrite, AuthenticatedRead, LogDeliveryWrite, BucketOwnerRead, BucketOwnerFullControl, or AwsExecRead.
	AccessControl() *string
	SetAccessControl(val *string)
	// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations() interface{}
	SetAnalyticsConfigurations(val interface{})
	// Returns the Amazon Resource Name (ARN) of the specified bucket.
	//
	// Example: `arn:aws:s3:::DOC-EXAMPLE-BUCKET`.
	AttrArn() *string
	// Returns the IPv4 DNS name of the specified bucket.
	//
	// Example: `DOC-EXAMPLE-BUCKET.s3.amazonaws.com`
	AttrDomainName() *string
	// Returns the IPv6 DNS name of the specified bucket.
	//
	// Example: `DOC-EXAMPLE-BUCKET.s3.dualstack.us-east-2.amazonaws.com`
	//
	// For more information about dual-stack endpoints, see [Using Amazon S3 Dual-Stack Endpoints](https://docs.aws.amazon.com/AmazonS3/latest/dev/dual-stack-endpoints.html) .
	AttrDualStackDomainName() *string
	// Returns the regional domain name of the specified bucket.
	//
	// Example: `DOC-EXAMPLE-BUCKET.s3.us-east-2.amazonaws.com`
	AttrRegionalDomainName() *string
	// Returns the Amazon S3 website endpoint for the specified bucket.
	//
	// Example (IPv4): `http://DOC-EXAMPLE-BUCKET.s3-website.us-east-2.amazonaws.com`
	//
	// Example (IPv6): `http://DOC-EXAMPLE-BUCKET.s3.dualstack.us-east-2.amazonaws.com`
	AttrWebsiteUrl() *string
	// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS) bucket.
	//
	// For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide* .
	BucketEncryption() interface{}
	SetBucketEncryption(val interface{})
	// A name for the bucket.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html) . For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	BucketName() *string
	SetBucketName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Describes the cross-origin access configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide* .
	CorsConfiguration() interface{}
	SetCorsConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Defines how Amazon S3 handles Intelligent-Tiering storage.
	IntelligentTieringConfigurations() interface{}
	SetIntelligentTieringConfigurations(val interface{})
	// Specifies the inventory configuration for an Amazon S3 bucket.
	//
	// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference* .
	InventoryConfigurations() interface{}
	SetInventoryConfigurations(val interface{})
	// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide* .
	LifecycleConfiguration() interface{}
	SetLifecycleConfiguration(val interface{})
	// Settings that define where logs are stored.
	LoggingConfiguration() interface{}
	SetLoggingConfiguration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket.
	//
	// If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html) .
	MetricsConfigurations() interface{}
	SetMetricsConfigurations(val interface{})
	// The tree node.
	Node() constructs.Node
	// Configuration that defines how Amazon S3 handles bucket notifications.
	NotificationConfiguration() interface{}
	SetNotificationConfiguration(val interface{})
	// Places an Object Lock configuration on the specified bucket.
	//
	// The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) .
	//
	// > - The `DefaultRetention` settings require both a mode and a period.
	// > - The `DefaultRetention` period can be either `Days` or `Years` but you must select one. You cannot specify `Days` and `Years` at the same time.
	// > - You can only enable Object Lock for new buckets. If you want to turn on Object Lock for an existing bucket, contact AWS Support.
	ObjectLockConfiguration() interface{}
	SetObjectLockConfiguration(val interface{})
	// Indicates whether this bucket has an Object Lock configuration enabled.
	//
	// Enable `ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.
	ObjectLockEnabled() interface{}
	SetObjectLockEnabled(val interface{})
	// Configuration that defines how Amazon S3 handles Object Ownership rules.
	OwnershipControls() interface{}
	SetOwnershipControls(val interface{})
	// Configuration that defines how Amazon S3 handles public access.
	PublicAccessBlockConfiguration() interface{}
	SetPublicAccessBlockConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Configuration for replicating objects in an S3 bucket.
	//
	// To enable replication, you must also enable versioning by using the `VersioningConfiguration` property.
	//
	// Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
	ReplicationConfiguration() interface{}
	SetReplicationConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Enables multiple versions of all objects in this bucket.
	//
	// You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
	VersioningConfiguration() interface{}
	SetVersioningConfiguration(val interface{})
	// Information used to configure the bucket as a static website.
	//
	// For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) .
	WebsiteConfiguration() interface{}
	SetWebsiteConfiguration(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBucket
type jsiiProxy_CfnBucket struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBucket) AccelerateConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accelerateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AccessControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AnalyticsConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"analyticsConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrDualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) AttrWebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWebsiteUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) BucketEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CorsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) IntelligentTieringConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intelligentTieringConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) InventoryConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inventoryConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) LifecycleConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) LoggingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) MetricsConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) NotificationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) ObjectLockConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectLockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) ObjectLockEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectLockEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) OwnershipControls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ownershipControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) PublicAccessBlockConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessBlockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) ReplicationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) VersioningConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucket) WebsiteConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websiteConfiguration",
		&returns,
	)
	return returns
}


// Create a new `AWS::S3::Bucket`.
func NewCfnBucket(scope constructs.Construct, id *string, props *CfnBucketProps) CfnBucket {
	_init_.Initialize()

	j := jsiiProxy_CfnBucket{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnBucket",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::S3::Bucket`.
func NewCfnBucket_Override(c CfnBucket, scope constructs.Construct, id *string, props *CfnBucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnBucket",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBucket) SetAccelerateConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"accelerateConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetAccessControl(val *string) {
	_jsii_.Set(
		j,
		"accessControl",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetAnalyticsConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"analyticsConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetBucketEncryption(val interface{}) {
	_jsii_.Set(
		j,
		"bucketEncryption",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetCorsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"corsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetIntelligentTieringConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"intelligentTieringConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetInventoryConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"inventoryConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetLifecycleConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"lifecycleConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetLoggingConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"loggingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetMetricsConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"metricsConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetNotificationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"notificationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetObjectLockConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"objectLockConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetObjectLockEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"objectLockEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetOwnershipControls(val interface{}) {
	_jsii_.Set(
		j,
		"ownershipControls",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetPublicAccessBlockConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"publicAccessBlockConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetReplicationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"replicationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetVersioningConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"versioningConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket) SetWebsiteConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"websiteConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBucket_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnBucket",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBucket_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnBucket",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBucket_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.CfnBucket",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBucket) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBucket) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBucket) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBucket) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBucket) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBucket) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBucket) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBucket) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBucket) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBucket) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucket) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.
//
// For more information, see [Stopping Incomplete Multipart Uploads Using a Bucket Lifecycle Policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   abortIncompleteMultipartUploadProperty := &abortIncompleteMultipartUploadProperty{
//   	daysAfterInitiation: jsii.Number(123),
//   }
//
type CfnBucket_AbortIncompleteMultipartUploadProperty struct {
	// Specifies the number of days after which Amazon S3 stops an incomplete multipart upload.
	DaysAfterInitiation *float64 `field:"required" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

// Configures the transfer acceleration state for an Amazon S3 bucket.
//
// For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accelerateConfigurationProperty := &accelerateConfigurationProperty{
//   	accelerationStatus: jsii.String("accelerationStatus"),
//   }
//
type CfnBucket_AccelerateConfigurationProperty struct {
	// Specifies the transfer acceleration status of the bucket.
	AccelerationStatus *string `field:"required" json:"accelerationStatus" yaml:"accelerationStatus"`
}

// Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket.
//
// If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlTranslationProperty := &accessControlTranslationProperty{
//   	owner: jsii.String("owner"),
//   }
//
type CfnBucket_AccessControlTranslationProperty struct {
	// Specifies the replica ownership.
	//
	// For default and valid values, see [PUT bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) in the *Amazon S3 API Reference* .
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analyticsConfigurationProperty := &analyticsConfigurationProperty{
//   	id: jsii.String("id"),
//   	storageClassAnalysis: &storageClassAnalysisProperty{
//   		dataExport: &dataExportProperty{
//   			destination: &destinationProperty{
//   				bucketArn: jsii.String("bucketArn"),
//   				format: jsii.String("format"),
//
//   				// the properties below are optional
//   				bucketAccountId: jsii.String("bucketAccountId"),
//   				prefix: jsii.String("prefix"),
//   			},
//   			outputSchemaVersion: jsii.String("outputSchemaVersion"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_AnalyticsConfigurationProperty struct {
	// The ID that identifies the analytics configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Contains data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes.
	StorageClassAnalysis interface{} `field:"required" json:"storageClassAnalysis" yaml:"storageClassAnalysis"`
	// The prefix that an object must have to be included in the analytics results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The tags to use when evaluating an analytics filter.
	//
	// The analytics only includes objects that meet the filter's criteria. If no filter is specified, all of the contents of the bucket are included in the analysis.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS) bucket.
//
// For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketEncryptionProperty := &bucketEncryptionProperty{
//   	serverSideEncryptionConfiguration: []interface{}{
//   		&serverSideEncryptionRuleProperty{
//   			bucketKeyEnabled: jsii.Boolean(false),
//   			serverSideEncryptionByDefault: &serverSideEncryptionByDefaultProperty{
//   				sseAlgorithm: jsii.String("sseAlgorithm"),
//
//   				// the properties below are optional
//   				kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_BucketEncryptionProperty struct {
	// Specifies the default server-side-encryption configuration.
	ServerSideEncryptionConfiguration interface{} `field:"required" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

// Describes the cross-origin access configuration for objects in an Amazon S3 bucket.
//
// For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsConfigurationProperty := &corsConfigurationProperty{
//   	corsRules: []interface{}{
//   		&corsRuleProperty{
//   			allowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			allowedOrigins: []*string{
//   				jsii.String("allowedOrigins"),
//   			},
//
//   			// the properties below are optional
//   			allowedHeaders: []*string{
//   				jsii.String("allowedHeaders"),
//   			},
//   			exposedHeaders: []*string{
//   				jsii.String("exposedHeaders"),
//   			},
//   			id: jsii.String("id"),
//   			maxAge: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBucket_CorsConfigurationProperty struct {
	// A set of origins and methods (cross-origin access that you want to allow).
	//
	// You can add up to 100 rules to the configuration.
	CorsRules interface{} `field:"required" json:"corsRules" yaml:"corsRules"`
}

// Specifies a cross-origin access rule for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsRuleProperty := &corsRuleProperty{
//   	allowedMethods: []*string{
//   		jsii.String("allowedMethods"),
//   	},
//   	allowedOrigins: []*string{
//   		jsii.String("allowedOrigins"),
//   	},
//
//   	// the properties below are optional
//   	allowedHeaders: []*string{
//   		jsii.String("allowedHeaders"),
//   	},
//   	exposedHeaders: []*string{
//   		jsii.String("exposedHeaders"),
//   	},
//   	id: jsii.String("id"),
//   	maxAge: jsii.Number(123),
//   }
//
type CfnBucket_CorsRuleProperty struct {
	// An HTTP method that you allow the origin to run.
	//
	// *Allowed values* : `GET` | `PUT` | `HEAD` | `POST` | `DELETE`.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Headers that are specified in the `Access-Control-Request-Headers` header.
	//
	// These headers are allowed in a preflight OPTIONS request. In response to any preflight OPTIONS request, Amazon S3 returns any requested headers that are allowed.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript `XMLHttpRequest` object).
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// A unique identifier for this rule.
	//
	// The value must be no more than 255 characters.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataExportProperty := &dataExportProperty{
//   	destination: &destinationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		format: jsii.String("format"),
//
//   		// the properties below are optional
//   		bucketAccountId: jsii.String("bucketAccountId"),
//   		prefix: jsii.String("prefix"),
//   	},
//   	outputSchemaVersion: jsii.String("outputSchemaVersion"),
//   }
//
type CfnBucket_DataExportProperty struct {
	// The place to store the data for an analysis.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// The version of the output schema to use when exporting data.
	//
	// Must be `V_1` .
	OutputSchemaVersion *string `field:"required" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

// The container element for specifying the default Object Lock retention settings for new objects placed in the specified bucket.
//
// > - The `DefaultRetention` settings require both a mode and a period.
// > - The `DefaultRetention` period can be either `Days` or `Years` but you must select one. You cannot specify `Days` and `Years` at the same time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultRetentionProperty := &defaultRetentionProperty{
//   	days: jsii.Number(123),
//   	mode: jsii.String("mode"),
//   	years: jsii.Number(123),
//   }
//
type CfnBucket_DefaultRetentionProperty struct {
	// The number of days that you want to specify for the default retention period.
	//
	// If Object Lock is turned on, you must specify `Mode` and specify either `Days` or `Years` .
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket.
	//
	// If Object Lock is turned on, you must specify `Mode` and specify either `Days` or `Years` .
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The number of years that you want to specify for the default retention period.
	//
	// If Object Lock is turned on, you must specify `Mode` and specify either `Days` or `Years` .
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

// Specifies whether Amazon S3 replicates delete markers.
//
// If you specify a `Filter` in your replication configuration, you must also include a `DeleteMarkerReplication` element. If your `Filter` includes a `Tag` element, the `DeleteMarkerReplication` `Status` must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config) .
//
// For more information about delete marker replication, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html) .
//
// > If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see [Backward Compatibility](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deleteMarkerReplicationProperty := &deleteMarkerReplicationProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_DeleteMarkerReplicationProperty struct {
	// Indicates whether to replicate delete markers.
	//
	// Disabled by default.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

// Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationProperty := &destinationProperty{
//   	bucketArn: jsii.String("bucketArn"),
//   	format: jsii.String("format"),
//
//   	// the properties below are optional
//   	bucketAccountId: jsii.String("bucketAccountId"),
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnBucket_DestinationProperty struct {
	// The Amazon Resource Name (ARN) of the bucket to which data is exported.
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Specifies the file format used when exporting data to Amazon S3.
	//
	// *Allowed values* : `CSV` | `ORC` | `Parquet`.
	Format *string `field:"required" json:"format" yaml:"format"`
	// The account ID that owns the destination S3 bucket.
	//
	// If no account ID is provided, the owner is not validated before exporting data.
	//
	// > Although this value is optional, we strongly recommend that you set it to help prevent problems if the destination bucket ownership changes.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The prefix to use when exporting data.
	//
	// The prefix is prepended to all results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &encryptionConfigurationProperty{
//   	replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   }
//
type CfnBucket_EncryptionConfigurationProperty struct {
	// Specifies the ID (Key ARN or Alias ARN) of the customer managed AWS KMS key stored in AWS Key Management Service (KMS) for the destination bucket.
	//
	// Amazon S3 uses this key to encrypt replica objects. Amazon S3 only supports symmetric, customer managed KMS keys. For more information, see [Using symmetric and asymmetric keys](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *AWS Key Management Service Developer Guide* .
	ReplicaKmsKeyId *string `field:"required" json:"replicaKmsKeyId" yaml:"replicaKmsKeyId"`
}

// Amazon S3 can send events to Amazon EventBridge whenever certain events happen in your bucket, see [Using EventBridge](https://docs.aws.amazon.com/AmazonS3/latest/userguide/EventBridge.html) in the *Amazon S3 User Guide* .
//
// Unlike other destinations, delivery of events to EventBridge can be either enabled or disabled for a bucket. If enabled, all events will be sent to EventBridge and you can use EventBridge rules to route events to additional targets. For more information, see [What Is Amazon EventBridge](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-what-is.html) in the *Amazon EventBridge User Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBridgeConfigurationProperty := &eventBridgeConfigurationProperty{
//   	eventBridgeEnabled: jsii.Boolean(false),
//   }
//
type CfnBucket_EventBridgeConfigurationProperty struct {
	// Enables delivery of events to Amazon EventBridge.
	EventBridgeEnabled interface{} `field:"optional" json:"eventBridgeEnabled" yaml:"eventBridgeEnabled"`
}

// Specifies the Amazon S3 object key name to filter on and whether to filter on the suffix or prefix of the key name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterRuleProperty := &filterRuleProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnBucket_FilterRuleProperty struct {
	// The object key name prefix or suffix identifying one or more objects to which the filtering rule applies.
	//
	// The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported. For more information, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value that the filter searches for in object key names.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Specifies the S3 Intelligent-Tiering configuration for an Amazon S3 bucket.
//
// For information about the S3 Intelligent-Tiering storage class, see [Storage class for automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intelligentTieringConfigurationProperty := &intelligentTieringConfigurationProperty{
//   	id: jsii.String("id"),
//   	status: jsii.String("status"),
//   	tierings: []interface{}{
//   		&tieringProperty{
//   			accessTier: jsii.String("accessTier"),
//   			days: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_IntelligentTieringConfigurationProperty struct {
	// The ID used to identify the S3 Intelligent-Tiering configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Specifies the status of the configuration.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies a list of S3 Intelligent-Tiering storage class tiers in the configuration.
	//
	// At least one tier must be defined in the list. At most, you can specify two tiers in the list, one for each available AccessTier: `ARCHIVE_ACCESS` and `DEEP_ARCHIVE_ACCESS` .
	//
	// > You only need Intelligent Tiering Configuration enabled on a bucket if you want to automatically move objects stored in the Intelligent-Tiering storage class to Archive Access or Deep Archive Access tiers.
	Tierings interface{} `field:"required" json:"tierings" yaml:"tierings"`
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A container for a key-value pair.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

// Specifies the inventory configuration for an Amazon S3 bucket.
//
// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inventoryConfigurationProperty := &inventoryConfigurationProperty{
//   	destination: &destinationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		format: jsii.String("format"),
//
//   		// the properties below are optional
//   		bucketAccountId: jsii.String("bucketAccountId"),
//   		prefix: jsii.String("prefix"),
//   	},
//   	enabled: jsii.Boolean(false),
//   	id: jsii.String("id"),
//   	includedObjectVersions: jsii.String("includedObjectVersions"),
//   	scheduleFrequency: jsii.String("scheduleFrequency"),
//
//   	// the properties below are optional
//   	optionalFields: []*string{
//   		jsii.String("optionalFields"),
//   	},
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnBucket_InventoryConfigurationProperty struct {
	// Contains information about where to publish the inventory results.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the inventory is enabled or disabled.
	//
	// If set to `True` , an inventory list is generated. If set to `False` , no inventory list is generated.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The ID used to identify the inventory configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Object versions to include in the inventory list.
	//
	// If set to `All` , the list includes all the object versions, which adds the version-related fields `VersionId` , `IsLatest` , and `DeleteMarker` to the list. If set to `Current` , the list does not contain these version-related fields.
	IncludedObjectVersions *string `field:"required" json:"includedObjectVersions" yaml:"includedObjectVersions"`
	// Specifies the schedule for generating inventory results.
	//
	// *Allowed values* : `Daily` | `Weekly`.
	ScheduleFrequency *string `field:"required" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// Contains the optional fields that are included in the inventory results.
	//
	// *Valid values* : `Size | LastModifiedDate | StorageClass | ETag | IsMultipartUploaded | ReplicationStatus | EncryptionStatus | ObjectLockRetainUntilDate | ObjectLockMode | ObjectLockLegalHoldStatus | IntelligentTieringAccessTier | BucketKeyStatus`.
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// Specifies the inventory filter prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// Describes the AWS Lambda functions to invoke and the events for which to invoke them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaConfigurationProperty := &lambdaConfigurationProperty{
//   	event: jsii.String("event"),
//   	function: jsii.String("function"),
//
//   	// the properties below are optional
//   	filter: &notificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&filterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_LambdaConfigurationProperty struct {
	// The Amazon S3 bucket event for which to invoke the AWS Lambda function.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the AWS Lambda function that Amazon S3 invokes when the specified event type occurs.
	Function *string `field:"required" json:"function" yaml:"function"`
	// The filtering rules that determine which objects invoke the AWS Lambda function.
	//
	// For example, you can create a filter so that only image files with a `.jpg` extension invoke the function when they are added to the Amazon S3 bucket.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
//
// For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleConfigurationProperty := &lifecycleConfigurationProperty{
//   	rules: []interface{}{
//   		&ruleProperty{
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   				daysAfterInitiation: jsii.Number(123),
//   			},
//   			expirationDate: NewDate(),
//   			expirationInDays: jsii.Number(123),
//   			expiredObjectDeleteMarker: jsii.Boolean(false),
//   			id: jsii.String("id"),
//   			noncurrentVersionExpiration: &noncurrentVersionExpirationProperty{
//   				noncurrentDays: jsii.Number(123),
//
//   				// the properties below are optional
//   				newerNoncurrentVersions: jsii.Number(123),
//   			},
//   			noncurrentVersionExpirationInDays: jsii.Number(123),
//   			noncurrentVersionTransition: &noncurrentVersionTransitionProperty{
//   				storageClass: jsii.String("storageClass"),
//   				transitionInDays: jsii.Number(123),
//
//   				// the properties below are optional
//   				newerNoncurrentVersions: jsii.Number(123),
//   			},
//   			noncurrentVersionTransitions: []interface{}{
//   				&noncurrentVersionTransitionProperty{
//   					storageClass: jsii.String("storageClass"),
//   					transitionInDays: jsii.Number(123),
//
//   					// the properties below are optional
//   					newerNoncurrentVersions: jsii.Number(123),
//   				},
//   			},
//   			objectSizeGreaterThan: jsii.Number(123),
//   			objectSizeLessThan: jsii.Number(123),
//   			prefix: jsii.String("prefix"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			transition: &transitionProperty{
//   				storageClass: jsii.String("storageClass"),
//
//   				// the properties below are optional
//   				transitionDate: NewDate(),
//   				transitionInDays: jsii.Number(123),
//   			},
//   			transitions: []interface{}{
//   				&transitionProperty{
//   					storageClass: jsii.String("storageClass"),
//
//   					// the properties below are optional
//   					transitionDate: NewDate(),
//   					transitionInDays: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_LifecycleConfigurationProperty struct {
	// A lifecycle rule for individual objects in an Amazon S3 bucket.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

// Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket.
//
// For examples and more information, see [PUT Bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the *Amazon S3 API Reference* .
//
// > To successfully complete the `AWS::S3::Bucket LoggingConfiguration` request, you must have `s3:PutObject` and `s3:PutObjectAcl` in your IAM permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingConfigurationProperty := &loggingConfigurationProperty{
//   	destinationBucketName: jsii.String("destinationBucketName"),
//   	logFilePrefix: jsii.String("logFilePrefix"),
//   }
//
type CfnBucket_LoggingConfigurationProperty struct {
	// The name of the bucket where Amazon S3 should store server access log files.
	//
	// You can store log files in any bucket that you own. By default, logs are stored in the bucket where the `LoggingConfiguration` property is defined.
	DestinationBucketName *string `field:"optional" json:"destinationBucketName" yaml:"destinationBucketName"`
	// A prefix for all log object keys.
	//
	// If you store log files from multiple Amazon S3 buckets in a single bucket, you can use a prefix to distinguish which log files came from which bucket.
	LogFilePrefix *string `field:"optional" json:"logFilePrefix" yaml:"logFilePrefix"`
}

// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket.
//
// If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For examples, see [AWS::S3::Bucket](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples) . For more information, see [PUT Bucket metrics](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsConfigurationProperty := &metricsConfigurationProperty{
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	accessPointArn: jsii.String("accessPointArn"),
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_MetricsConfigurationProperty struct {
	// The ID used to identify the metrics configuration.
	//
	// This can be any value you choose that helps you identify your metrics configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The access point that was used while performing operations on the object.
	//
	// The metrics configuration only includes objects that meet the filter's criteria.
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// The prefix that an object must have to be included in the metrics results.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Specifies a list of tag filters to use as a metrics configuration filter.
	//
	// The metrics configuration includes only objects that meet the filter's criteria.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

// A container specifying replication metrics-related settings enabling replication metrics and events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsProperty := &metricsProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	eventThreshold: &replicationTimeValueProperty{
//   		minutes: jsii.Number(123),
//   	},
//   }
//
type CfnBucket_MetricsProperty struct {
	// Specifies whether the replication metrics are enabled.
	Status *string `field:"required" json:"status" yaml:"status"`
	// A container specifying the time threshold for emitting the `s3:Replication:OperationMissedThreshold` event.
	EventThreshold interface{} `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

// Specifies when noncurrent object versions expire.
//
// Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noncurrentVersionExpirationProperty := &noncurrentVersionExpirationProperty{
//   	noncurrentDays: jsii.Number(123),
//
//   	// the properties below are optional
//   	newerNoncurrentVersions: jsii.Number(123),
//   }
//
type CfnBucket_NoncurrentVersionExpirationProperty struct {
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	//
	// For information about the noncurrent days calculations, see [How Amazon S3 Calculates When an Object Became Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide* .
	NoncurrentDays *float64 `field:"required" json:"noncurrentDays" yaml:"noncurrentDays"`
	// Specifies how many noncurrent versions Amazon S3 will retain.
	//
	// If there are this many more recent noncurrent versions, Amazon S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide* .
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

// Container for the transition rule that describes when noncurrent objects transition to the `STANDARD_IA` , `ONEZONE_IA` , `INTELLIGENT_TIERING` , `GLACIER_IR` , `GLACIER` , or `DEEP_ARCHIVE` storage class.
//
// If your bucket is versioning-enabled (or versioning is suspended), you can set this action to request that Amazon S3 transition noncurrent object versions to the `STANDARD_IA` , `ONEZONE_IA` , `INTELLIGENT_TIERING` , `GLACIER_IR` , `GLACIER` , or `DEEP_ARCHIVE` storage class at a specific period in the object's lifetime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   noncurrentVersionTransitionProperty := &noncurrentVersionTransitionProperty{
//   	storageClass: jsii.String("storageClass"),
//   	transitionInDays: jsii.Number(123),
//
//   	// the properties below are optional
//   	newerNoncurrentVersions: jsii.Number(123),
//   }
//
type CfnBucket_NoncurrentVersionTransitionProperty struct {
	// The class of storage used to store the object.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	//
	// For information about the noncurrent days calculations, see [How Amazon S3 Calculates How Long an Object Has Been Noncurrent](https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#non-current-days-calculations) in the *Amazon S3 User Guide* .
	TransitionInDays *float64 `field:"required" json:"transitionInDays" yaml:"transitionInDays"`
	// Specifies how many noncurrent versions Amazon S3 will retain.
	//
	// If there are this many more recent noncurrent versions, Amazon S3 will take the associated action. For more information about noncurrent versions, see [Lifecycle configuration elements](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html) in the *Amazon S3 User Guide* .
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

// Describes the notification configuration for an Amazon S3 bucket.
//
// > If you create the target resource and related permissions in the same template, you might have a circular dependency.
// >
// > For example, you might use the `AWS::Lambda::Permission` resource to grant the bucket permission to invoke an AWS Lambda function. However, AWS CloudFormation can't create the bucket until the bucket has permission to invoke the function ( AWS CloudFormation checks whether the bucket can invoke the function). If you're using Refs to pass the bucket name, this leads to a circular dependency.
// >
// > To avoid this dependency, you can create all resources without specifying the notification configuration. Then, update the stack with a notification configuration.
// >
// > For more information on permissions, see [AWS::Lambda::Permission](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html) and [Granting Permissions to Publish Event Notification Messages to a Destination](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html#grant-destinations-permissions-to-s3) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationConfigurationProperty := &notificationConfigurationProperty{
//   	eventBridgeConfiguration: &eventBridgeConfigurationProperty{
//   		eventBridgeEnabled: jsii.Boolean(false),
//   	},
//   	lambdaConfigurations: []interface{}{
//   		&lambdaConfigurationProperty{
//   			event: jsii.String("event"),
//   			function: jsii.String("function"),
//
//   			// the properties below are optional
//   			filter: &notificationFilterProperty{
//   				s3Key: &s3KeyFilterProperty{
//   					rules: []interface{}{
//   						&filterRuleProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	queueConfigurations: []interface{}{
//   		&queueConfigurationProperty{
//   			event: jsii.String("event"),
//   			queue: jsii.String("queue"),
//
//   			// the properties below are optional
//   			filter: &notificationFilterProperty{
//   				s3Key: &s3KeyFilterProperty{
//   					rules: []interface{}{
//   						&filterRuleProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	topicConfigurations: []interface{}{
//   		&topicConfigurationProperty{
//   			event: jsii.String("event"),
//   			topic: jsii.String("topic"),
//
//   			// the properties below are optional
//   			filter: &notificationFilterProperty{
//   				s3Key: &s3KeyFilterProperty{
//   					rules: []interface{}{
//   						&filterRuleProperty{
//   							name: jsii.String("name"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_NotificationConfigurationProperty struct {
	// Enables delivery of events to Amazon EventBridge.
	EventBridgeConfiguration interface{} `field:"optional" json:"eventBridgeConfiguration" yaml:"eventBridgeConfiguration"`
	// Describes the AWS Lambda functions to invoke and the events for which to invoke them.
	LambdaConfigurations interface{} `field:"optional" json:"lambdaConfigurations" yaml:"lambdaConfigurations"`
	// The Amazon Simple Queue Service queues to publish messages to and the events for which to publish messages.
	QueueConfigurations interface{} `field:"optional" json:"queueConfigurations" yaml:"queueConfigurations"`
	// The topic to which notifications are sent and the events for which notifications are generated.
	TopicConfigurations interface{} `field:"optional" json:"topicConfigurations" yaml:"topicConfigurations"`
}

// Specifies object key name filtering rules.
//
// For information about key name filtering, see [Configuring Event Notifications](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationFilterProperty := &notificationFilterProperty{
//   	s3Key: &s3KeyFilterProperty{
//   		rules: []interface{}{
//   			&filterRuleProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_NotificationFilterProperty struct {
	// A container for object key name prefix and suffix filtering rules.
	S3Key interface{} `field:"required" json:"s3Key" yaml:"s3Key"`
}

// Places an Object Lock configuration on the specified bucket.
//
// The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectLockConfigurationProperty := &objectLockConfigurationProperty{
//   	objectLockEnabled: jsii.String("objectLockEnabled"),
//   	rule: &objectLockRuleProperty{
//   		defaultRetention: &defaultRetentionProperty{
//   			days: jsii.Number(123),
//   			mode: jsii.String("mode"),
//   			years: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBucket_ObjectLockConfigurationProperty struct {
	// Indicates whether this bucket has an Object Lock configuration enabled.
	//
	// Enable `ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.
	ObjectLockEnabled *string `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Specifies the Object Lock rule for the specified object.
	//
	// Enable the this rule when you apply `ObjectLockConfiguration` to a bucket. If Object Lock is turned on, bucket settings require both `Mode` and a period of either `Days` or `Years` . You cannot specify `Days` and `Years` at the same time. For more information, see [ObjectLockRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-objectlockrule.html) and [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html) .
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

// Specifies the Object Lock rule for the specified object.
//
// Enable the this rule when you apply `ObjectLockConfiguration` to a bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectLockRuleProperty := &objectLockRuleProperty{
//   	defaultRetention: &defaultRetentionProperty{
//   		days: jsii.Number(123),
//   		mode: jsii.String("mode"),
//   		years: jsii.Number(123),
//   	},
//   }
//
type CfnBucket_ObjectLockRuleProperty struct {
	// The default Object Lock retention mode and period that you want to apply to new objects placed in the specified bucket.
	//
	// If Object Lock is turned on, bucket settings require both `Mode` and a period of either `Days` or `Years` . You cannot specify `Days` and `Years` at the same time. For more information about allowable values for mode and period, see [DefaultRetention](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-defaultretention.html) .
	DefaultRetention interface{} `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
}

// Specifies the container element for Object Ownership rules.
//
// S3 Object Ownership is an Amazon S3 bucket-level setting that you can use to disable access control lists (ACLs) and take ownership of every object in your bucket, simplifying access management for data stored in Amazon S3. For more information, see [Controlling ownership of objects and disabling ACLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ownershipControlsProperty := &ownershipControlsProperty{
//   	rules: []interface{}{
//   		&ownershipControlsRuleProperty{
//   			objectOwnership: jsii.String("objectOwnership"),
//   		},
//   	},
//   }
//
type CfnBucket_OwnershipControlsProperty struct {
	// Specifies the container element for Object Ownership rules.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

// Specifies an Object Ownership rule.
//
// S3 Object Ownership is an Amazon S3 bucket-level setting that you can use to disable access control lists (ACLs) and take ownership of every object in your bucket, simplifying access management for data stored in Amazon S3. For more information, see [Controlling ownership of objects and disabling ACLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ownershipControlsRuleProperty := &ownershipControlsRuleProperty{
//   	objectOwnership: jsii.String("objectOwnership"),
//   }
//
type CfnBucket_OwnershipControlsRuleProperty struct {
	// Specifies an Object Ownership rule.
	//
	// *Allowed values* : `BucketOwnerEnforced` | `ObjectWriter` | `BucketOwnerPreferred`.
	ObjectOwnership *string `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
}

// The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket.
//
// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessBlockConfigurationProperty := &publicAccessBlockConfigurationProperty{
//   	blockPublicAcls: jsii.Boolean(false),
//   	blockPublicPolicy: jsii.Boolean(false),
//   	ignorePublicAcls: jsii.Boolean(false),
//   	restrictPublicBuckets: jsii.Boolean(false),
//   }
//
type CfnBucket_PublicAccessBlockConfigurationProperty struct {
	// Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes the following behavior:
	//
	// - PUT Bucket ACL and PUT Object ACL calls fail if the specified ACL is public.
	// - PUT Object calls fail if the request includes a public ACL.
	// - PUT Bucket calls fail if the request includes a public ACL.
	//
	// Enabling this setting doesn't affect existing policies or ACLs.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Specifies whether Amazon S3 should block public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	//
	// Enabling this setting doesn't affect existing bucket policies.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.
	//
	// Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Specifies whether Amazon S3 should restrict public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` restricts access to this bucket to only AWS service principals and authorized users within this account if the bucket has a public policy.
	//
	// Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

// Specifies the configuration for publishing messages to an Amazon Simple Queue Service (Amazon SQS) queue when Amazon S3 detects specified events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueConfigurationProperty := &queueConfigurationProperty{
//   	event: jsii.String("event"),
//   	queue: jsii.String("queue"),
//
//   	// the properties below are optional
//   	filter: &notificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&filterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_QueueConfigurationProperty struct {
	// The Amazon S3 bucket event about which you want to publish messages to Amazon SQS.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type.
	//
	// FIFO queues are not allowed when enabling an SQS queue as the event notification destination.
	Queue *string `field:"required" json:"queue" yaml:"queue"`
	// The filtering rules that determine which objects trigger notifications.
	//
	// For example, you can create a filter so that Amazon S3 sends notifications only when image files with a `.jpg` extension are added to the bucket. For more information, see [Configuring event notifications using object key name filtering](https://docs.aws.amazon.com/AmazonS3/latest/user-guide/notification-how-to-filtering.html) in the *Amazon S3 User Guide* .
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

// Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectAllRequestsToProperty := &redirectAllRequestsToProperty{
//   	hostName: jsii.String("hostName"),
//
//   	// the properties below are optional
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnBucket_RedirectAllRequestsToProperty struct {
	// Name of the host where requests are redirected.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests.
	//
	// The default is the protocol that is used in the original request.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

// Specifies how requests are redirected.
//
// In the event of an error, you can specify a different error code to return.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectRuleProperty := &redirectRuleProperty{
//   	hostName: jsii.String("hostName"),
//   	httpRedirectCode: jsii.String("httpRedirectCode"),
//   	protocol: jsii.String("protocol"),
//   	replaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   	replaceKeyWith: jsii.String("replaceKeyWith"),
//   }
//
type CfnBucket_RedirectRuleProperty struct {
	// The host name to use in the redirect request.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The HTTP redirect code to use on the response.
	//
	// Not required if one of the siblings is present.
	HttpRedirectCode *string `field:"optional" json:"httpRedirectCode" yaml:"httpRedirectCode"`
	// Protocol to use when redirecting requests.
	//
	// The default is the protocol that is used in the original request.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The object key prefix to use in the redirect request.
	//
	// For example, to redirect requests for all pages with prefix `docs/` (objects in the `docs/` folder) to `documents/` , you can set a condition block with `KeyPrefixEquals` set to `docs/` and in the Redirect set `ReplaceKeyPrefixWith` to `/documents` . Not required if one of the siblings is present. Can be present only if `ReplaceKeyWith` is not provided.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	ReplaceKeyPrefixWith *string `field:"optional" json:"replaceKeyPrefixWith" yaml:"replaceKeyPrefixWith"`
	// The specific object key to use in the redirect request.
	//
	// For example, redirect request to `error.html` . Not required if one of the siblings is present. Can be present only if `ReplaceKeyPrefixWith` is not provided.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	ReplaceKeyWith *string `field:"optional" json:"replaceKeyWith" yaml:"replaceKeyWith"`
}

// A filter that you can specify for selection for modifications on replicas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicaModificationsProperty := &replicaModificationsProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_ReplicaModificationsProperty struct {
	// Specifies whether Amazon S3 replicates modifications on replicas.
	//
	// *Allowed values* : `Enabled` | `Disabled`.
	Status *string `field:"required" json:"status" yaml:"status"`
}

// A container for replication rules.
//
// You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationConfigurationProperty := &replicationConfigurationProperty{
//   	role: jsii.String("role"),
//   	rules: []interface{}{
//   		&replicationRuleProperty{
//   			destination: &replicationDestinationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				accessControlTranslation: &accessControlTranslationProperty{
//   					owner: jsii.String("owner"),
//   				},
//   				account: jsii.String("account"),
//   				encryptionConfiguration: &encryptionConfigurationProperty{
//   					replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   				},
//   				metrics: &metricsProperty{
//   					status: jsii.String("status"),
//
//   					// the properties below are optional
//   					eventThreshold: &replicationTimeValueProperty{
//   						minutes: jsii.Number(123),
//   					},
//   				},
//   				replicationTime: &replicationTimeProperty{
//   					status: jsii.String("status"),
//   					time: &replicationTimeValueProperty{
//   						minutes: jsii.Number(123),
//   					},
//   				},
//   				storageClass: jsii.String("storageClass"),
//   			},
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			deleteMarkerReplication: &deleteMarkerReplicationProperty{
//   				status: jsii.String("status"),
//   			},
//   			filter: &replicationRuleFilterProperty{
//   				and: &replicationRuleAndOperatorProperty{
//   					prefix: jsii.String("prefix"),
//   					tagFilters: []interface{}{
//   						&tagFilterProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   				},
//   				prefix: jsii.String("prefix"),
//   				tagFilter: &tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			id: jsii.String("id"),
//   			prefix: jsii.String("prefix"),
//   			priority: jsii.Number(123),
//   			sourceSelectionCriteria: &sourceSelectionCriteriaProperty{
//   				replicaModifications: &replicaModificationsProperty{
//   					status: jsii.String("status"),
//   				},
//   				sseKmsEncryptedObjects: &sseKmsEncryptedObjectsProperty{
//   					status: jsii.String("status"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_ReplicationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating objects.
	//
	// For more information, see [How to Set Up Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-how-setup.html) in the *Amazon S3 User Guide* .
	Role *string `field:"required" json:"role" yaml:"role"`
	// A container for one or more replication rules.
	//
	// A replication configuration must have at least one rule and can contain a maximum of 1,000 rules.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

// A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationDestinationProperty := &replicationDestinationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	accessControlTranslation: &accessControlTranslationProperty{
//   		owner: jsii.String("owner"),
//   	},
//   	account: jsii.String("account"),
//   	encryptionConfiguration: &encryptionConfigurationProperty{
//   		replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   	},
//   	metrics: &metricsProperty{
//   		status: jsii.String("status"),
//
//   		// the properties below are optional
//   		eventThreshold: &replicationTimeValueProperty{
//   			minutes: jsii.Number(123),
//   		},
//   	},
//   	replicationTime: &replicationTimeProperty{
//   		status: jsii.String("status"),
//   		time: &replicationTimeValueProperty{
//   			minutes: jsii.Number(123),
//   		},
//   	},
//   	storageClass: jsii.String("storageClass"),
//   }
//
type CfnBucket_ReplicationDestinationProperty struct {
	// The Amazon Resource Name (ARN) of the bucket where you want Amazon S3 to store the results.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Specify this only in a cross-account scenario (where source and destination bucket owners are not the same), and you want to change replica ownership to the AWS account that owns the destination bucket.
	//
	// If this is not specified in the replication configuration, the replicas are owned by same AWS account that owns the source object.
	AccessControlTranslation interface{} `field:"optional" json:"accessControlTranslation" yaml:"accessControlTranslation"`
	// Destination bucket owner account ID.
	//
	// In a cross-account scenario, if you direct Amazon S3 to change replica ownership to the AWS account that owns the destination bucket by specifying the `AccessControlTranslation` property, this is the account ID of the destination bucket owner. For more information, see [Cross-Region Replication Additional Configuration: Change Replica Owner](https://docs.aws.amazon.com/AmazonS3/latest/dev/crr-change-owner.html) in the *Amazon S3 User Guide* .
	//
	// If you specify the `AccessControlTranslation` property, the `Account` property is required.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Specifies encryption-related information.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// A container specifying replication metrics-related settings enabling replication metrics and events.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
	// A container specifying S3 Replication Time Control (S3 RTC), including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated.
	//
	// Must be specified together with a `Metrics` block.
	ReplicationTime interface{} `field:"optional" json:"replicationTime" yaml:"replicationTime"`
	// The storage class to use when replicating objects, such as S3 Standard or reduced redundancy.
	//
	// By default, Amazon S3 uses the storage class of the source object to create the object replica.
	//
	// For valid values, see the `StorageClass` element of the [PUT Bucket replication](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTreplication.html) action in the *Amazon S3 API Reference* .
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

// A container for specifying rule filters.
//
// The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter.
//
// For example:
//
// - If you specify both a `Prefix` and a `TagFilter` , wrap these filters in an `And` tag.
// - If you specify a filter based on multiple tags, wrap the `TagFilter` elements in an `And` tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleAndOperatorProperty := &replicationRuleAndOperatorProperty{
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_ReplicationRuleAndOperatorProperty struct {
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// An array of tags containing key and value pairs.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

// A filter that identifies the subset of objects to which the replication rule applies.
//
// A `Filter` must specify exactly one `Prefix` , `TagFilter` , or an `And` child element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleFilterProperty := &replicationRuleFilterProperty{
//   	and: &replicationRuleAndOperatorProperty{
//   		prefix: jsii.String("prefix"),
//   		tagFilters: []interface{}{
//   			&tagFilterProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	prefix: jsii.String("prefix"),
//   	tagFilter: &tagFilterProperty{
//   		key: jsii.String("key"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnBucket_ReplicationRuleFilterProperty struct {
	// A container for specifying rule filters.
	//
	// The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter. For example:
	//
	// - If you specify both a `Prefix` and a `TagFilter` , wrap these filters in an `And` tag.
	// - If you specify a filter based on multiple tags, wrap the `TagFilter` elements in an `And` tag.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A container for specifying a tag key and value.
	//
	// The rule applies only to objects that have the tag in their tag set.
	TagFilter interface{} `field:"optional" json:"tagFilter" yaml:"tagFilter"`
}

// Specifies which Amazon S3 objects to replicate and where to store the replicas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleProperty := &replicationRuleProperty{
//   	destination: &replicationDestinationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		accessControlTranslation: &accessControlTranslationProperty{
//   			owner: jsii.String("owner"),
//   		},
//   		account: jsii.String("account"),
//   		encryptionConfiguration: &encryptionConfigurationProperty{
//   			replicaKmsKeyId: jsii.String("replicaKmsKeyId"),
//   		},
//   		metrics: &metricsProperty{
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			eventThreshold: &replicationTimeValueProperty{
//   				minutes: jsii.Number(123),
//   			},
//   		},
//   		replicationTime: &replicationTimeProperty{
//   			status: jsii.String("status"),
//   			time: &replicationTimeValueProperty{
//   				minutes: jsii.Number(123),
//   			},
//   		},
//   		storageClass: jsii.String("storageClass"),
//   	},
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	deleteMarkerReplication: &deleteMarkerReplicationProperty{
//   		status: jsii.String("status"),
//   	},
//   	filter: &replicationRuleFilterProperty{
//   		and: &replicationRuleAndOperatorProperty{
//   			prefix: jsii.String("prefix"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		prefix: jsii.String("prefix"),
//   		tagFilter: &tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	prefix: jsii.String("prefix"),
//   	priority: jsii.Number(123),
//   	sourceSelectionCriteria: &sourceSelectionCriteriaProperty{
//   		replicaModifications: &replicaModificationsProperty{
//   			status: jsii.String("status"),
//   		},
//   		sseKmsEncryptedObjects: &sseKmsEncryptedObjectsProperty{
//   			status: jsii.String("status"),
//   		},
//   	},
//   }
//
type CfnBucket_ReplicationRuleProperty struct {
	// A container for information about the replication destination and its configurations including enabling the S3 Replication Time Control (S3 RTC).
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Specifies whether the rule is enabled.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies whether Amazon S3 replicates delete markers.
	//
	// If you specify a `Filter` in your replication configuration, you must also include a `DeleteMarkerReplication` element. If your `Filter` includes a `Tag` element, the `DeleteMarkerReplication` `Status` must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config) .
	//
	// For more information about delete marker replication, see [Basic Rule Configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html) .
	//
	// > If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see [Backward Compatibility](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations) .
	DeleteMarkerReplication interface{} `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// A filter that identifies the subset of objects to which the replication rule applies.
	//
	// A `Filter` must specify exactly one `Prefix` , `TagFilter` , or an `And` child element. The use of the filter field indicates this is a V2 replication configuration. V1 does not have this field.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// A unique identifier for the rule.
	//
	// The maximum value is 255 characters. If you don't specify a value, AWS CloudFormation generates a random ID. When using a V2 replication configuration this property is capitalized as "ID".
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An object key name prefix that identifies the object or objects to which the rule applies.
	//
	// The maximum prefix length is 1,024 characters. To include all objects in a bucket, specify an empty string.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The priority indicates which rule has precedence whenever two or more replication rules conflict.
	//
	// Amazon S3 will attempt to replicate objects according to all replication rules. However, if there are two or more rules with the same destination bucket, then objects will be replicated according to the rule with the highest priority. The higher the number, the higher the priority.
	//
	// For more information, see [Replication](https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html) in the *Amazon S3 User Guide* .
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A container that describes additional filters for identifying the source objects that you want to replicate.
	//
	// You can choose to enable or disable the replication of these objects.
	SourceSelectionCriteria interface{} `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

// A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated.
//
// Must be specified together with a `Metrics` block.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationTimeProperty := &replicationTimeProperty{
//   	status: jsii.String("status"),
//   	time: &replicationTimeValueProperty{
//   		minutes: jsii.Number(123),
//   	},
//   }
//
type CfnBucket_ReplicationTimeProperty struct {
	// Specifies whether the replication time is enabled.
	Status *string `field:"required" json:"status" yaml:"status"`
	// A container specifying the time by which replication should be complete for all objects and operations on objects.
	Time interface{} `field:"required" json:"time" yaml:"time"`
}

// A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics `EventThreshold` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationTimeValueProperty := &replicationTimeValueProperty{
//   	minutes: jsii.Number(123),
//   }
//
type CfnBucket_ReplicationTimeValueProperty struct {
	// Contains an integer specifying time in minutes.
	//
	// Valid value: 15.
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
}

// A container for describing a condition that must be met for the specified redirect to apply.
//
// For example, 1. If request is for pages in the `/docs` folder, redirect to the `/documents` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingRuleConditionProperty := &routingRuleConditionProperty{
//   	httpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   	keyPrefixEquals: jsii.String("keyPrefixEquals"),
//   }
//
type CfnBucket_RoutingRuleConditionProperty struct {
	// The HTTP error code when the redirect is applied.
	//
	// In the event of an error, if the error code equals this value, then the specified redirect is applied.
	//
	// Required when parent element `Condition` is specified and sibling `KeyPrefixEquals` is not specified. If both are specified, then both must be true for the redirect to be applied.
	HttpErrorCodeReturnedEquals *string `field:"optional" json:"httpErrorCodeReturnedEquals" yaml:"httpErrorCodeReturnedEquals"`
	// The object key name prefix when the redirect is applied.
	//
	// For example, to redirect requests for `ExamplePage.html` , the key prefix will be `ExamplePage.html` . To redirect request for all pages with the prefix `docs/` , the key prefix will be `/docs` , which identifies all objects in the docs/ folder.
	//
	// Required when the parent element `Condition` is specified and sibling `HttpErrorCodeReturnedEquals` is not specified. If both conditions are specified, both must be true for the redirect to be applied.
	KeyPrefixEquals *string `field:"optional" json:"keyPrefixEquals" yaml:"keyPrefixEquals"`
}

// Specifies the redirect behavior and when a redirect is applied.
//
// For more information about routing rules, see [Configuring advanced conditional redirects](https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html#advanced-conditional-redirects) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingRuleProperty := &routingRuleProperty{
//   	redirectRule: &redirectRuleProperty{
//   		hostName: jsii.String("hostName"),
//   		httpRedirectCode: jsii.String("httpRedirectCode"),
//   		protocol: jsii.String("protocol"),
//   		replaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   		replaceKeyWith: jsii.String("replaceKeyWith"),
//   	},
//
//   	// the properties below are optional
//   	routingRuleCondition: &routingRuleConditionProperty{
//   		httpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   		keyPrefixEquals: jsii.String("keyPrefixEquals"),
//   	},
//   }
//
type CfnBucket_RoutingRuleProperty struct {
	// Container for redirect information.
	//
	// You can redirect requests to another host, to another page, or with another protocol. In the event of an error, you can specify a different error code to return.
	RedirectRule interface{} `field:"required" json:"redirectRule" yaml:"redirectRule"`
	// A container for describing a condition that must be met for the specified redirect to apply.
	//
	// For example, 1. If request is for pages in the `/docs` folder, redirect to the `/documents` folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.
	RoutingRuleCondition interface{} `field:"optional" json:"routingRuleCondition" yaml:"routingRuleCondition"`
}

// Specifies lifecycle rules for an Amazon S3 bucket.
//
// For more information, see [Put Bucket Lifecycle Configuration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlifecycle.html) in the *Amazon S3 API Reference* .
//
// You must specify at least one of the following properties: `AbortIncompleteMultipartUpload` , `ExpirationDate` , `ExpirationInDays` , `NoncurrentVersionExpirationInDays` , `NoncurrentVersionTransition` , `NoncurrentVersionTransitions` , `Transition` , or `Transitions` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &ruleProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   		daysAfterInitiation: jsii.Number(123),
//   	},
//   	expirationDate: NewDate(),
//   	expirationInDays: jsii.Number(123),
//   	expiredObjectDeleteMarker: jsii.Boolean(false),
//   	id: jsii.String("id"),
//   	noncurrentVersionExpiration: &noncurrentVersionExpirationProperty{
//   		noncurrentDays: jsii.Number(123),
//
//   		// the properties below are optional
//   		newerNoncurrentVersions: jsii.Number(123),
//   	},
//   	noncurrentVersionExpirationInDays: jsii.Number(123),
//   	noncurrentVersionTransition: &noncurrentVersionTransitionProperty{
//   		storageClass: jsii.String("storageClass"),
//   		transitionInDays: jsii.Number(123),
//
//   		// the properties below are optional
//   		newerNoncurrentVersions: jsii.Number(123),
//   	},
//   	noncurrentVersionTransitions: []interface{}{
//   		&noncurrentVersionTransitionProperty{
//   			storageClass: jsii.String("storageClass"),
//   			transitionInDays: jsii.Number(123),
//
//   			// the properties below are optional
//   			newerNoncurrentVersions: jsii.Number(123),
//   		},
//   	},
//   	objectSizeGreaterThan: jsii.Number(123),
//   	objectSizeLessThan: jsii.Number(123),
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transition: &transitionProperty{
//   		storageClass: jsii.String("storageClass"),
//
//   		// the properties below are optional
//   		transitionDate: NewDate(),
//   		transitionInDays: jsii.Number(123),
//   	},
//   	transitions: []interface{}{
//   		&transitionProperty{
//   			storageClass: jsii.String("storageClass"),
//
//   			// the properties below are optional
//   			transitionDate: NewDate(),
//   			transitionInDays: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBucket_RuleProperty struct {
	// If `Enabled` , the rule is currently being applied.
	//
	// If `Disabled` , the rule is not currently being applied.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3 bucket.
	AbortIncompleteMultipartUpload interface{} `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// Indicates when objects are deleted from Amazon S3 and Amazon S3 Glacier.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.
	ExpirationDate interface{} `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon S3 Glacier.
	//
	// If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Indicates whether Amazon S3 will remove a delete marker without any noncurrent versions.
	//
	// If set to true, the delete marker will be removed if there are no noncurrent versions. This cannot be specified with `ExpirationInDays` , `ExpirationDate` , or `TagFilters` .
	ExpiredObjectDeleteMarker interface{} `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
	// Unique identifier for the rule.
	//
	// The value can't be longer than 255 characters.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies when noncurrent object versions expire.
	//
	// Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.
	NoncurrentVersionExpiration interface{} `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// (Deprecated.) For buckets with versioning enabled (or suspended), specifies the time, in days, between when a new version of the object is uploaded to the bucket and when old versions of the object expire. When object versions expire, Amazon S3 permanently deletes them. If you specify a transition and expiration time, the expiration time must be later than the transition time.
	NoncurrentVersionExpirationInDays *float64 `field:"optional" json:"noncurrentVersionExpirationInDays" yaml:"noncurrentVersionExpirationInDays"`
	// (Deprecated.) For buckets with versioning enabled (or suspended), specifies when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the `NoncurrentVersionTransitions` property.
	NoncurrentVersionTransition interface{} `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// For buckets with versioning enabled (or suspended), one or more transition rules that specify when non-current objects transition to a specified storage class.
	//
	// If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the `NoncurrentVersionTransition` property.
	NoncurrentVersionTransitions interface{} `field:"optional" json:"noncurrentVersionTransitions" yaml:"noncurrentVersionTransitions"`
	// Specifies the minimum object size in bytes for this rule to apply to.
	//
	// For more information about size based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide* .
	ObjectSizeGreaterThan *float64 `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Specifies the maximum object size in bytes for this rule to apply to.
	//
	// For more information about sized based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide* .
	ObjectSizeLessThan *float64 `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Object key prefix that identifies one or more objects to which this rule applies.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Tags to use to identify a subset of objects to which the lifecycle rule applies.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// (Deprecated.) Specifies when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the `Transitions` property.
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
	// One or more transition rules that specify when an object transitions to a specified storage class.
	//
	// If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the `Transition` property.
	Transitions interface{} `field:"optional" json:"transitions" yaml:"transitions"`
}

// A container for object key name prefix and suffix filtering rules.
//
// > The same type of filter rule cannot be used more than once. For example, you cannot specify two prefix rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3KeyFilterProperty := &s3KeyFilterProperty{
//   	rules: []interface{}{
//   		&filterRuleProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_S3KeyFilterProperty struct {
	// A list of containers for the key-value pair that defines the criteria for the filter rule.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

// Describes the default server-side encryption to apply to new objects in the bucket.
//
// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied. If you don't specify a customer managed key at configuration, Amazon S3 automatically creates an AWS KMS key in your AWS account the first time that you add an object encrypted with SSE-KMS to a bucket. By default, Amazon S3 uses this KMS key for SSE-KMS. For more information, see [PUT Bucket encryption](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTencryption.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionByDefaultProperty := &serverSideEncryptionByDefaultProperty{
//   	sseAlgorithm: jsii.String("sseAlgorithm"),
//
//   	// the properties below are optional
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   }
//
type CfnBucket_ServerSideEncryptionByDefaultProperty struct {
	// Server-side encryption algorithm to use for the default encryption.
	SseAlgorithm *string `field:"required" json:"sseAlgorithm" yaml:"sseAlgorithm"`
	// KMS key ID to use for the default encryption. This parameter is allowed if SSEAlgorithm is aws:kms.
	//
	// You can specify the key ID or the Amazon Resource Name (ARN) of the CMK. However, if you are using encryption with cross-account operations, you must use a fully qualified CMK ARN. For more information, see [Using encryption for cross-account operations](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html#bucket-encryption-update-bucket-policy) .
	//
	// For example:
	//
	// - Key ID: `1234abcd-12ab-34cd-56ef-1234567890ab`
	// - Key ARN: `arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
	//
	// > Amazon S3 only supports symmetric KMS keys and not asymmetric KMS keys. For more information, see [Using Symmetric and Asymmetric Keys](https://docs.aws.amazon.com//kms/latest/developerguide/symmetric-asymmetric.html) in the *AWS Key Management Service Developer Guide* .
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

// Specifies the default server-side encryption configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serverSideEncryptionRuleProperty := &serverSideEncryptionRuleProperty{
//   	bucketKeyEnabled: jsii.Boolean(false),
//   	serverSideEncryptionByDefault: &serverSideEncryptionByDefaultProperty{
//   		sseAlgorithm: jsii.String("sseAlgorithm"),
//
//   		// the properties below are optional
//   		kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	},
//   }
//
type CfnBucket_ServerSideEncryptionRuleProperty struct {
	// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
	//
	// Existing objects are not affected. Setting the `BucketKeyEnabled` element to `true` causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.
	//
	// For more information, see [Amazon S3 Bucket Keys](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html) in the *Amazon S3 User Guide* .
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Specifies the default server-side encryption to apply to new objects in the bucket.
	//
	// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
	ServerSideEncryptionByDefault interface{} `field:"optional" json:"serverSideEncryptionByDefault" yaml:"serverSideEncryptionByDefault"`
}

// A container that describes additional filters for identifying the source objects that you want to replicate.
//
// You can choose to enable or disable the replication of these objects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceSelectionCriteriaProperty := &sourceSelectionCriteriaProperty{
//   	replicaModifications: &replicaModificationsProperty{
//   		status: jsii.String("status"),
//   	},
//   	sseKmsEncryptedObjects: &sseKmsEncryptedObjectsProperty{
//   		status: jsii.String("status"),
//   	},
//   }
//
type CfnBucket_SourceSelectionCriteriaProperty struct {
	// A filter that you can specify for selection for modifications on replicas.
	ReplicaModifications interface{} `field:"optional" json:"replicaModifications" yaml:"replicaModifications"`
	// A container for filter information for the selection of Amazon S3 objects encrypted with AWS KMS.
	SseKmsEncryptedObjects interface{} `field:"optional" json:"sseKmsEncryptedObjects" yaml:"sseKmsEncryptedObjects"`
}

// A container for filter information for the selection of S3 objects encrypted with AWS KMS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sseKmsEncryptedObjectsProperty := &sseKmsEncryptedObjectsProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_SseKmsEncryptedObjectsProperty struct {
	// Specifies whether Amazon S3 replicates objects created with server-side encryption using an AWS KMS key stored in AWS Key Management Service.
	Status *string `field:"required" json:"status" yaml:"status"`
}

// Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageClassAnalysisProperty := &storageClassAnalysisProperty{
//   	dataExport: &dataExportProperty{
//   		destination: &destinationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			format: jsii.String("format"),
//
//   			// the properties below are optional
//   			bucketAccountId: jsii.String("bucketAccountId"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		outputSchemaVersion: jsii.String("outputSchemaVersion"),
//   	},
//   }
//
type CfnBucket_StorageClassAnalysisProperty struct {
	// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
	DataExport interface{} `field:"optional" json:"dataExport" yaml:"dataExport"`
}

// Specifies tags to use to identify a subset of objects for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &tagFilterProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnBucket_TagFilterProperty struct {
	// The tag key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// The S3 Intelligent-Tiering storage class is designed to optimize storage costs by automatically moving data to the most cost-effective storage access tier, without additional operational overhead.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tieringProperty := &tieringProperty{
//   	accessTier: jsii.String("accessTier"),
//   	days: jsii.Number(123),
//   }
//
type CfnBucket_TieringProperty struct {
	// S3 Intelligent-Tiering access tier.
	//
	// See [Storage class for automatically optimizing frequently and infrequently accessed objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access) for a list of access tiers in the S3 Intelligent-Tiering storage class.
	AccessTier *string `field:"required" json:"accessTier" yaml:"accessTier"`
	// The number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier.
	//
	// The minimum number of days specified for Archive Access tier must be at least 90 days and Deep Archive Access tier must be at least 180 days. The maximum can be up to 2 years (730 days).
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

// A container for specifying the configuration for publication of messages to an Amazon Simple Notification Service (Amazon SNS) topic when Amazon S3 detects specified events.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicConfigurationProperty := &topicConfigurationProperty{
//   	event: jsii.String("event"),
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	filter: &notificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&filterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_TopicConfigurationProperty struct {
	// The Amazon S3 bucket event about which to send notifications.
	//
	// For more information, see [Supported Event Types](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the *Amazon S3 User Guide* .
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it detects events of the specified type.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The filtering rules that determine for which objects to send notifications.
	//
	// For example, you can create a filter so that Amazon S3 sends notifications only when image files with a `.jpg` extension are added to the bucket.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

// Specifies when an object transitions to a specified storage class.
//
// For more information about Amazon S3 lifecycle configuration rules, see [Transitioning Objects Using Amazon S3 Lifecycle](https://docs.aws.amazon.com/AmazonS3/latest/dev/lifecycle-transition-general-considerations.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitionProperty := &transitionProperty{
//   	storageClass: jsii.String("storageClass"),
//
//   	// the properties below are optional
//   	transitionDate: NewDate(),
//   	transitionInDays: jsii.Number(123),
//   }
//
type CfnBucket_TransitionProperty struct {
	// The storage class to which you want the object to transition.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates when objects are transitioned to the specified storage class.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	TransitionDate interface{} `field:"optional" json:"transitionDate" yaml:"transitionDate"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	//
	// The value must be a positive integer.
	TransitionInDays *float64 `field:"optional" json:"transitionInDays" yaml:"transitionInDays"`
}

// Describes the versioning state of an Amazon S3 bucket.
//
// For more information, see [PUT Bucket versioning](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTVersioningStatus.html) in the *Amazon S3 API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   versioningConfigurationProperty := &versioningConfigurationProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnBucket_VersioningConfigurationProperty struct {
	// The versioning state of the bucket.
	Status *string `field:"required" json:"status" yaml:"status"`
}

// Specifies website configuration parameters for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   websiteConfigurationProperty := &websiteConfigurationProperty{
//   	errorDocument: jsii.String("errorDocument"),
//   	indexDocument: jsii.String("indexDocument"),
//   	redirectAllRequestsTo: &redirectAllRequestsToProperty{
//   		hostName: jsii.String("hostName"),
//
//   		// the properties below are optional
//   		protocol: jsii.String("protocol"),
//   	},
//   	routingRules: []interface{}{
//   		&routingRuleProperty{
//   			redirectRule: &redirectRuleProperty{
//   				hostName: jsii.String("hostName"),
//   				httpRedirectCode: jsii.String("httpRedirectCode"),
//   				protocol: jsii.String("protocol"),
//   				replaceKeyPrefixWith: jsii.String("replaceKeyPrefixWith"),
//   				replaceKeyWith: jsii.String("replaceKeyWith"),
//   			},
//
//   			// the properties below are optional
//   			routingRuleCondition: &routingRuleConditionProperty{
//   				httpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   				keyPrefixEquals: jsii.String("keyPrefixEquals"),
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_WebsiteConfigurationProperty struct {
	// The name of the error document for the website.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// The name of the index document for the website.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// The redirect behavior for every request to this bucket's website endpoint.
	//
	// > If you specify this property, you can't specify any other property.
	RedirectAllRequestsTo interface{} `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Rules that define when a redirect is applied and the redirect behavior.
	RoutingRules interface{} `field:"optional" json:"routingRules" yaml:"routingRules"`
}

// A CloudFormation `AWS::S3::BucketPolicy`.
//
// Applies an Amazon S3 bucket policy to an Amazon S3 bucket. If you are using an identity other than the root user of the AWS account that owns the bucket, the calling identity must have the `PutBucketPolicy` permissions on the specified bucket and belong to the bucket owner's account in order to use this operation.
//
// If you don't have `PutBucketPolicy` permissions, Amazon S3 returns a `403 Access Denied` error. If you have the correct permissions, but you're not using an identity that belongs to the bucket owner's account, Amazon S3 returns a `405 Method Not Allowed` error.
//
// > As a security precaution, the root user of the AWS account that owns a bucket can always use this operation, even if the policy explicitly denies the root user the ability to perform this action.
//
// For more information, see [Bucket policy examples](https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html) .
//
// The following operations are related to `PutBucketPolicy` :
//
// - [CreateBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
// - [DeleteBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnBucketPolicy := awscdk.Aws_s3.NewCfnBucketPolicy(this, jsii.String("MyCfnBucketPolicy"), &cfnBucketPolicyProps{
//   	bucket: jsii.String("bucket"),
//   	policyDocument: policyDocument,
//   })
//
type CfnBucketPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The name of the Amazon S3 bucket to which the policy applies.
	Bucket() *string
	SetBucket(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// A policy document containing permissions to add to the specified bucket.
	//
	// In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language Overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the *Amazon S3 User Guide* .
	PolicyDocument() interface{}
	SetPolicyDocument(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnBucketPolicy
type jsiiProxy_CfnBucketPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnBucketPolicy) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) PolicyDocument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBucketPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::S3::BucketPolicy`.
func NewCfnBucketPolicy(scope constructs.Construct, id *string, props *CfnBucketPolicyProps) CfnBucketPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnBucketPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnBucketPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::S3::BucketPolicy`.
func NewCfnBucketPolicy_Override(c CfnBucketPolicy, scope constructs.Construct, id *string, props *CfnBucketPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnBucketPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnBucketPolicy) SetBucket(val *string) {
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_CfnBucketPolicy) SetPolicyDocument(val interface{}) {
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnBucketPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnBucketPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnBucketPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnBucketPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnBucketPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnBucketPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBucketPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.CfnBucketPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBucketPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBucketPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBucketPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnBucketPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnBucketPolicyProps := &cfnBucketPolicyProps{
//   	bucket: jsii.String("bucket"),
//   	policyDocument: policyDocument,
//   }
//
type CfnBucketPolicyProps struct {
	// The name of the Amazon S3 bucket to which the policy applies.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A policy document containing permissions to add to the specified bucket.
	//
	// In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM. For more information, see the AWS::IAM::Policy [PolicyDocument](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument) resource description in this guide and [Access Policy Language Overview](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the *Amazon S3 User Guide* .
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

// Properties for defining a `CfnBucket`.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   rawBucket := s3.NewCfnBucket(this, jsii.String("Bucket"), &cfnBucketProps{
//   })
//   // -or-
//   rawBucketAlt := myBucket.node.defaultChild.(cfnBucket)
//
//   // then
//   rawBucket.cfnOptions.condition = awscdk.NewCfnCondition(this, jsii.String("EnableBucket"), &cfnConditionProps{
//   })
//   rawBucket.cfnOptions.metadata = map[string]interface{}{
//   	"metadataKey": jsii.String("MetadataValue"),
//   }
//
type CfnBucketProps struct {
	// Configures the transfer acceleration state for an Amazon S3 bucket.
	//
	// For more information, see [Amazon S3 Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html) in the *Amazon S3 User Guide* .
	AccelerateConfiguration interface{} `field:"optional" json:"accelerateConfiguration" yaml:"accelerateConfiguration"`
	// A canned access control list (ACL) that grants predefined permissions to the bucket.
	//
	// For more information about canned ACLs, see [Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) in the *Amazon S3 User Guide* .
	//
	// Be aware that the syntax for this property differs from the information provided in the *Amazon S3 User Guide* . The AccessControl property is case-sensitive and must be one of the following values: Private, PublicRead, PublicReadWrite, AuthenticatedRead, LogDeliveryWrite, BucketOwnerRead, BucketOwnerFullControl, or AwsExecRead.
	AccessControl *string `field:"optional" json:"accessControl" yaml:"accessControl"`
	// Specifies the configuration and any analyses for the analytics filter of an Amazon S3 bucket.
	AnalyticsConfigurations interface{} `field:"optional" json:"analyticsConfigurations" yaml:"analyticsConfigurations"`
	// Specifies default encryption for a bucket using server-side encryption with Amazon S3-managed keys (SSE-S3) or AWS KMS-managed keys (SSE-KMS) bucket.
	//
	// For information about the Amazon S3 default encryption feature, see [Amazon S3 Default Encryption for S3 Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the *Amazon S3 User Guide* .
	BucketEncryption interface{} `field:"optional" json:"bucketEncryption" yaml:"bucketEncryption"`
	// A name for the bucket.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html) . For more information, see [Rules for naming Amazon S3 buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/BucketRestrictions.html#bucketnamingrules) in the *Amazon S3 User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Describes the cross-origin access configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide* .
	CorsConfiguration interface{} `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// Defines how Amazon S3 handles Intelligent-Tiering storage.
	IntelligentTieringConfigurations interface{} `field:"optional" json:"intelligentTieringConfigurations" yaml:"intelligentTieringConfigurations"`
	// Specifies the inventory configuration for an Amazon S3 bucket.
	//
	// For more information, see [GET Bucket inventory](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketGETInventoryConfig.html) in the *Amazon S3 API Reference* .
	InventoryConfigurations interface{} `field:"optional" json:"inventoryConfigurations" yaml:"inventoryConfigurations"`
	// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
	//
	// For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide* .
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Settings that define where logs are stored.
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket.
	//
	// If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For more information, see [PutBucketMetricsConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html) .
	MetricsConfigurations interface{} `field:"optional" json:"metricsConfigurations" yaml:"metricsConfigurations"`
	// Configuration that defines how Amazon S3 handles bucket notifications.
	NotificationConfiguration interface{} `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// Places an Object Lock configuration on the specified bucket.
	//
	// The rule specified in the Object Lock configuration will be applied by default to every new object placed in the specified bucket. For more information, see [Locking Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) .
	//
	// > - The `DefaultRetention` settings require both a mode and a period.
	// > - The `DefaultRetention` period can be either `Days` or `Years` but you must select one. You cannot specify `Days` and `Years` at the same time.
	// > - You can only enable Object Lock for new buckets. If you want to turn on Object Lock for an existing bucket, contact AWS Support.
	ObjectLockConfiguration interface{} `field:"optional" json:"objectLockConfiguration" yaml:"objectLockConfiguration"`
	// Indicates whether this bucket has an Object Lock configuration enabled.
	//
	// Enable `ObjectLockEnabled` when you apply `ObjectLockConfiguration` to a bucket.
	ObjectLockEnabled interface{} `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// Configuration that defines how Amazon S3 handles Object Ownership rules.
	OwnershipControls interface{} `field:"optional" json:"ownershipControls" yaml:"ownershipControls"`
	// Configuration that defines how Amazon S3 handles public access.
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
	// Configuration for replicating objects in an S3 bucket.
	//
	// To enable replication, you must also enable versioning by using the `VersioningConfiguration` property.
	//
	// Amazon S3 can store replicated objects in a single destination bucket or multiple destination buckets. The destination bucket or buckets must already exist.
	ReplicationConfiguration interface{} `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this S3 bucket.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Enables multiple versions of all objects in this bucket.
	//
	// You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.
	VersioningConfiguration interface{} `field:"optional" json:"versioningConfiguration" yaml:"versioningConfiguration"`
	// Information used to configure the bucket as a static website.
	//
	// For more information, see [Hosting Websites on Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) .
	WebsiteConfiguration interface{} `field:"optional" json:"websiteConfiguration" yaml:"websiteConfiguration"`
}

// A CloudFormation `AWS::S3::MultiRegionAccessPoint`.
//
// The `AWS::S3::MultiRegionAccessPoint` resource creates an Amazon S3 Multi-Region Access Point. To learn more about Multi-Region Access Points, see [Multi-Region Access Points in Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPoints.html) in the in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMultiRegionAccessPoint := awscdk.Aws_s3.NewCfnMultiRegionAccessPoint(this, jsii.String("MyCfnMultiRegionAccessPoint"), &cfnMultiRegionAccessPointProps{
//   	regions: []interface{}{
//   		&regionProperty{
//   			bucket: jsii.String("bucket"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	publicAccessBlockConfiguration: &publicAccessBlockConfigurationProperty{
//   		blockPublicAcls: jsii.Boolean(false),
//   		blockPublicPolicy: jsii.Boolean(false),
//   		ignorePublicAcls: jsii.Boolean(false),
//   		restrictPublicBuckets: jsii.Boolean(false),
//   	},
//   })
//
type CfnMultiRegionAccessPoint interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The alias for the Multi-Region Access Point.
	//
	// For more information about the distinction between the name and the alias of an Multi-Region Access Point, see [Managing Multi-Region Access Points](https://docs.aws.amazon.com/AmazonS3/latest/userguide/CreatingMultiRegionAccessPoints.html#multi-region-access-point-naming) in the *Amazon S3 User Guide* .
	AttrAlias() *string
	// The timestamp of when the Multi-Region Access Point is created.
	AttrCreatedAt() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the Multi-Region Access Point.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// The PublicAccessBlock configuration that you want to apply to this Multi-Region Access Point.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers an object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
	PublicAccessBlockConfiguration() interface{}
	SetPublicAccessBlockConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A collection of the Regions and buckets associated with the Multi-Region Access Point.
	Regions() interface{}
	SetRegions(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMultiRegionAccessPoint
type jsiiProxy_CfnMultiRegionAccessPoint struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) AttrAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) PublicAccessBlockConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicAccessBlockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) Regions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::S3::MultiRegionAccessPoint`.
func NewCfnMultiRegionAccessPoint(scope constructs.Construct, id *string, props *CfnMultiRegionAccessPointProps) CfnMultiRegionAccessPoint {
	_init_.Initialize()

	j := jsiiProxy_CfnMultiRegionAccessPoint{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::S3::MultiRegionAccessPoint`.
func NewCfnMultiRegionAccessPoint_Override(c CfnMultiRegionAccessPoint, scope constructs.Construct, id *string, props *CfnMultiRegionAccessPointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) SetPublicAccessBlockConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"publicAccessBlockConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnMultiRegionAccessPoint) SetRegions(val interface{}) {
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMultiRegionAccessPoint_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPoint",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMultiRegionAccessPoint_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPoint",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnMultiRegionAccessPoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMultiRegionAccessPoint_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPoint",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPoint) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The PublicAccessBlock configuration that you want to apply to this Amazon S3 Multi-Region Access Point.
//
// You can enable the configuration options in any combination. For more information about when Amazon S3 considers an object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessBlockConfigurationProperty := &publicAccessBlockConfigurationProperty{
//   	blockPublicAcls: jsii.Boolean(false),
//   	blockPublicPolicy: jsii.Boolean(false),
//   	ignorePublicAcls: jsii.Boolean(false),
//   	restrictPublicBuckets: jsii.Boolean(false),
//   }
//
type CfnMultiRegionAccessPoint_PublicAccessBlockConfigurationProperty struct {
	// Specifies whether Amazon S3 should block public access control lists (ACLs) for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes the following behavior:
	//
	// - PUT Bucket ACL and PUT Object ACL calls fail if the specified ACL is public.
	// - PUT Object calls fail if the request includes a public ACL.
	// - PUT Bucket calls fail if the request includes a public ACL.
	//
	// Enabling this setting doesn't affect existing policies or ACLs.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Specifies whether Amazon S3 should block public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access.
	//
	// Enabling this setting doesn't affect existing bucket policies.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this bucket.
	//
	// Setting this element to `TRUE` causes Amazon S3 to ignore all public ACLs on this bucket and objects in this bucket.
	//
	// Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Specifies whether Amazon S3 should restrict public bucket policies for this bucket.
	//
	// Setting this element to `TRUE` restricts access to this bucket to only AWS service principals and authorized users within this account if the bucket has a public policy.
	//
	// Enabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

// A bucket associated with a specific Region when creating Multi-Region Access Points.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionProperty := &regionProperty{
//   	bucket: jsii.String("bucket"),
//   }
//
type CfnMultiRegionAccessPoint_RegionProperty struct {
	// The name of the associated bucket for the Region.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
}

// A CloudFormation `AWS::S3::MultiRegionAccessPointPolicy`.
//
// Applies an Amazon S3 access policy to an Amazon S3 Multi-Region Access Point.
//
// It is not possible to delete an access policy for a Multi-Region Access Point from the CloudFormation template. When you attempt to delete the policy, CloudFormation updates the policy using `DeletionPolicy:Retain` and `UpdateReplacePolicy:Retain` . CloudFormation updates the policy to only allow access to the account that created the bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnMultiRegionAccessPointPolicy := awscdk.Aws_s3.NewCfnMultiRegionAccessPointPolicy(this, jsii.String("MyCfnMultiRegionAccessPointPolicy"), &cfnMultiRegionAccessPointPolicyProps{
//   	mrapName: jsii.String("mrapName"),
//   	policy: policy,
//   })
//
type CfnMultiRegionAccessPointPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The name of the Multi-Region Access Point.
	MrapName() *string
	SetMrapName(val *string)
	// The tree node.
	Node() constructs.Node
	// The access policy associated with the Multi-Region Access Point.
	Policy() interface{}
	SetPolicy(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnMultiRegionAccessPointPolicy
type jsiiProxy_CfnMultiRegionAccessPointPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) MrapName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mrapName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) Policy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::S3::MultiRegionAccessPointPolicy`.
func NewCfnMultiRegionAccessPointPolicy(scope constructs.Construct, id *string, props *CfnMultiRegionAccessPointPolicyProps) CfnMultiRegionAccessPointPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnMultiRegionAccessPointPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPointPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::S3::MultiRegionAccessPointPolicy`.
func NewCfnMultiRegionAccessPointPolicy_Override(c CfnMultiRegionAccessPointPolicy, scope constructs.Construct, id *string, props *CfnMultiRegionAccessPointPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPointPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) SetMrapName(val *string) {
	_jsii_.Set(
		j,
		"mrapName",
		val,
	)
}

func (j *jsiiProxy_CfnMultiRegionAccessPointPolicy) SetPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnMultiRegionAccessPointPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPointPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnMultiRegionAccessPointPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPointPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnMultiRegionAccessPointPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPointPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMultiRegionAccessPointPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.CfnMultiRegionAccessPointPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMultiRegionAccessPointPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnMultiRegionAccessPointPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   cfnMultiRegionAccessPointPolicyProps := &cfnMultiRegionAccessPointPolicyProps{
//   	mrapName: jsii.String("mrapName"),
//   	policy: policy,
//   }
//
type CfnMultiRegionAccessPointPolicyProps struct {
	// The name of the Multi-Region Access Point.
	MrapName *string `field:"required" json:"mrapName" yaml:"mrapName"`
	// The access policy associated with the Multi-Region Access Point.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
}

// Properties for defining a `CfnMultiRegionAccessPoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMultiRegionAccessPointProps := &cfnMultiRegionAccessPointProps{
//   	regions: []interface{}{
//   		&regionProperty{
//   			bucket: jsii.String("bucket"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	publicAccessBlockConfiguration: &publicAccessBlockConfigurationProperty{
//   		blockPublicAcls: jsii.Boolean(false),
//   		blockPublicPolicy: jsii.Boolean(false),
//   		ignorePublicAcls: jsii.Boolean(false),
//   		restrictPublicBuckets: jsii.Boolean(false),
//   	},
//   }
//
type CfnMultiRegionAccessPointProps struct {
	// A collection of the Regions and buckets associated with the Multi-Region Access Point.
	Regions interface{} `field:"required" json:"regions" yaml:"regions"`
	// The name of the Multi-Region Access Point.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The PublicAccessBlock configuration that you want to apply to this Multi-Region Access Point.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers an object public, see [The Meaning of "Public"](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the *Amazon S3 User Guide* .
	PublicAccessBlockConfiguration interface{} `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
}

// A CloudFormation `AWS::S3::StorageLens`.
//
// The AWS::S3::StorageLens resource creates an instance of an Amazon S3 Storage Lens configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryption interface{}
//
//   cfnStorageLens := awscdk.Aws_s3.NewCfnStorageLens(this, jsii.String("MyCfnStorageLens"), &cfnStorageLensProps{
//   	storageLensConfiguration: &storageLensConfigurationProperty{
//   		accountLevel: &accountLevelProperty{
//   			bucketLevel: &bucketLevelProperty{
//   				activityMetrics: &activityMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   				},
//   				prefixLevel: &prefixLevelProperty{
//   					storageMetrics: &prefixLevelStorageMetricsProperty{
//   						isEnabled: jsii.Boolean(false),
//   						selectionCriteria: &selectionCriteriaProperty{
//   							delimiter: jsii.String("delimiter"),
//   							maxDepth: jsii.Number(123),
//   							minStorageBytesPercentage: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			activityMetrics: &activityMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   		},
//   		id: jsii.String("id"),
//   		isEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		awsOrg: &awsOrgProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		dataExport: &dataExportProperty{
//   			cloudWatchMetrics: &cloudWatchMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   			s3BucketDestination: &s3BucketDestinationProperty{
//   				accountId: jsii.String("accountId"),
//   				arn: jsii.String("arn"),
//   				format: jsii.String("format"),
//   				outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   				// the properties below are optional
//   				encryption: encryption,
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   		exclude: &bucketsAndRegionsProperty{
//   			buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		include: &bucketsAndRegionsProperty{
//   			buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		storageLensArn: jsii.String("storageLensArn"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnStorageLens interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// This property contains the details of the ARN of the S3 Storage Lens configuration.
	//
	// This property is read-only.
	AttrStorageLensConfigurationStorageLensArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// This resource contains the details Amazon S3 Storage Lens configuration.
	StorageLensConfiguration() interface{}
	SetStorageLensConfiguration(val interface{})
	// A set of tags (key–value pairs) to associate with the Storage Lens configuration.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnStorageLens
type jsiiProxy_CfnStorageLens struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStorageLens) AttrStorageLensConfigurationStorageLensArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStorageLensConfigurationStorageLensArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) StorageLensConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageLensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStorageLens) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::S3::StorageLens`.
func NewCfnStorageLens(scope constructs.Construct, id *string, props *CfnStorageLensProps) CfnStorageLens {
	_init_.Initialize()

	j := jsiiProxy_CfnStorageLens{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnStorageLens",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::S3::StorageLens`.
func NewCfnStorageLens_Override(c CfnStorageLens, scope constructs.Construct, id *string, props *CfnStorageLensProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.CfnStorageLens",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStorageLens) SetStorageLensConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"storageLensConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStorageLens_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnStorageLens",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStorageLens_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnStorageLens",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnStorageLens_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.CfnStorageLens",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStorageLens_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.CfnStorageLens",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStorageLens) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStorageLens) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStorageLens) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStorageLens) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStorageLens) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStorageLens) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStorageLens) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStorageLens) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageLens) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageLens) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStorageLens) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStorageLens) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageLens) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageLens) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStorageLens) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// This resource contains the details of the account-level metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountLevelProperty := &accountLevelProperty{
//   	bucketLevel: &bucketLevelProperty{
//   		activityMetrics: &activityMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   		prefixLevel: &prefixLevelProperty{
//   			storageMetrics: &prefixLevelStorageMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   				selectionCriteria: &selectionCriteriaProperty{
//   					delimiter: jsii.String("delimiter"),
//   					maxDepth: jsii.Number(123),
//   					minStorageBytesPercentage: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	activityMetrics: &activityMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnStorageLens_AccountLevelProperty struct {
	// This property contains the details of the account-level bucket-level configurations for Amazon S3 Storage Lens.
	BucketLevel interface{} `field:"required" json:"bucketLevel" yaml:"bucketLevel"`
	// This property contains the details of the account-level activity metrics for Amazon S3 Storage Lens.
	ActivityMetrics interface{} `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
}

// This resource contains the details of the activity metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   activityMetricsProperty := &activityMetricsProperty{
//   	isEnabled: jsii.Boolean(false),
//   }
//
type CfnStorageLens_ActivityMetricsProperty struct {
	// A property that indicates whether the activity metrics is enabled.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

// This resource contains the details of the AWS Organization for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   awsOrgProperty := &awsOrgProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnStorageLens_AwsOrgProperty struct {
	// This resource contains the ARN of the AWS Organization.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

// A property for the bucket-level storage metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketLevelProperty := &bucketLevelProperty{
//   	activityMetrics: &activityMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   	prefixLevel: &prefixLevelProperty{
//   		storageMetrics: &prefixLevelStorageMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   			selectionCriteria: &selectionCriteriaProperty{
//   				delimiter: jsii.String("delimiter"),
//   				maxDepth: jsii.Number(123),
//   				minStorageBytesPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
type CfnStorageLens_BucketLevelProperty struct {
	// A property for the bucket-level activity metrics for Amazon S3 Storage Lens.
	ActivityMetrics interface{} `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// A property for the bucket-level prefix-level storage metrics for S3 Storage Lens.
	PrefixLevel interface{} `field:"optional" json:"prefixLevel" yaml:"prefixLevel"`
}

// This resource contains the details of the buckets and Regions for the Amazon S3 Storage Lens configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bucketsAndRegionsProperty := &bucketsAndRegionsProperty{
//   	buckets: []*string{
//   		jsii.String("buckets"),
//   	},
//   	regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
type CfnStorageLens_BucketsAndRegionsProperty struct {
	// This property contains the details of the buckets for the Amazon S3 Storage Lens configuration.
	//
	// This should be the bucket Amazon Resource Name(ARN). For valid values, see [Buckets ARN format here](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_Include.html#API_control_Include_Contents) in the *Amazon S3 API Reference* .
	Buckets *[]*string `field:"optional" json:"buckets" yaml:"buckets"`
	// This property contains the details of the Regions for the S3 Storage Lens configuration.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

// This resource enables the Amazon CloudWatch publishing option for S3 Storage Lens metrics.
//
// For more information, see [Monitor S3 Storage Lens metrics in CloudWatch](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_view_metrics_cloudwatch.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudWatchMetricsProperty := &cloudWatchMetricsProperty{
//   	isEnabled: jsii.Boolean(false),
//   }
//
type CfnStorageLens_CloudWatchMetricsProperty struct {
	// This property identifies whether the CloudWatch publishing option for S3 Storage Lens is enabled.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
}

// This resource contains the details of the Amazon S3 Storage Lens metrics export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryption interface{}
//
//   dataExportProperty := &dataExportProperty{
//   	cloudWatchMetrics: &cloudWatchMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   	},
//   	s3BucketDestination: &s3BucketDestinationProperty{
//   		accountId: jsii.String("accountId"),
//   		arn: jsii.String("arn"),
//   		format: jsii.String("format"),
//   		outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   		// the properties below are optional
//   		encryption: encryption,
//   		prefix: jsii.String("prefix"),
//   	},
//   }
//
type CfnStorageLens_DataExportProperty struct {
	// This property enables the Amazon CloudWatch publishing option for S3 Storage Lens metrics.
	CloudWatchMetrics interface{} `field:"optional" json:"cloudWatchMetrics" yaml:"cloudWatchMetrics"`
	// This property contains the details of the bucket where the S3 Storage Lens metrics export will be placed.
	S3BucketDestination interface{} `field:"optional" json:"s3BucketDestination" yaml:"s3BucketDestination"`
}

// This resource contains the details of the prefix-level of the Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixLevelProperty := &prefixLevelProperty{
//   	storageMetrics: &prefixLevelStorageMetricsProperty{
//   		isEnabled: jsii.Boolean(false),
//   		selectionCriteria: &selectionCriteriaProperty{
//   			delimiter: jsii.String("delimiter"),
//   			maxDepth: jsii.Number(123),
//   			minStorageBytesPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnStorageLens_PrefixLevelProperty struct {
	// A property for the prefix-level storage metrics for Amazon S3 Storage Lens.
	StorageMetrics interface{} `field:"required" json:"storageMetrics" yaml:"storageMetrics"`
}

// This resource contains the details of the prefix-level storage metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixLevelStorageMetricsProperty := &prefixLevelStorageMetricsProperty{
//   	isEnabled: jsii.Boolean(false),
//   	selectionCriteria: &selectionCriteriaProperty{
//   		delimiter: jsii.String("delimiter"),
//   		maxDepth: jsii.Number(123),
//   		minStorageBytesPercentage: jsii.Number(123),
//   	},
//   }
//
type CfnStorageLens_PrefixLevelStorageMetricsProperty struct {
	// This property identifies whether the details of the prefix-level storage metrics for S3 Storage Lens are enabled.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// This property identifies whether the details of the prefix-level storage metrics for S3 Storage Lens are enabled.
	SelectionCriteria interface{} `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

// This resource contains the details of the bucket where the Amazon S3 Storage Lens metrics export will be placed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryption interface{}
//
//   s3BucketDestinationProperty := &s3BucketDestinationProperty{
//   	accountId: jsii.String("accountId"),
//   	arn: jsii.String("arn"),
//   	format: jsii.String("format"),
//   	outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   	// the properties below are optional
//   	encryption: encryption,
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnStorageLens_S3BucketDestinationProperty struct {
	// This property contains the details of the AWS account ID of the S3 Storage Lens export bucket destination.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// This property contains the details of the ARN of the bucket destination of the S3 Storage Lens export.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// This property contains the details of the format of the S3 Storage Lens export bucket destination.
	Format *string `field:"required" json:"format" yaml:"format"`
	// This property contains the details of the output schema version of the S3 Storage Lens export bucket destination.
	OutputSchemaVersion *string `field:"required" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
	// This property contains the details of the encryption of the bucket destination of the Amazon S3 Storage Lens metrics export.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// This property contains the details of the prefix of the bucket destination of the S3 Storage Lens export .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// This resource contains the details of the Amazon S3 Storage Lens selection criteria.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectionCriteriaProperty := &selectionCriteriaProperty{
//   	delimiter: jsii.String("delimiter"),
//   	maxDepth: jsii.Number(123),
//   	minStorageBytesPercentage: jsii.Number(123),
//   }
//
type CfnStorageLens_SelectionCriteriaProperty struct {
	// This property contains the details of the S3 Storage Lens delimiter being used.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// This property contains the details of the max depth that S3 Storage Lens will collect metrics up to.
	MaxDepth *float64 `field:"optional" json:"maxDepth" yaml:"maxDepth"`
	// This property contains the details of the minimum storage bytes percentage threshold that S3 Storage Lens will collect metrics up to.
	MinStorageBytesPercentage *float64 `field:"optional" json:"minStorageBytesPercentage" yaml:"minStorageBytesPercentage"`
}

// This is the property of the Amazon S3 Storage Lens configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryption interface{}
//
//   storageLensConfigurationProperty := &storageLensConfigurationProperty{
//   	accountLevel: &accountLevelProperty{
//   		bucketLevel: &bucketLevelProperty{
//   			activityMetrics: &activityMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   			prefixLevel: &prefixLevelProperty{
//   				storageMetrics: &prefixLevelStorageMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   					selectionCriteria: &selectionCriteriaProperty{
//   						delimiter: jsii.String("delimiter"),
//   						maxDepth: jsii.Number(123),
//   						minStorageBytesPercentage: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		activityMetrics: &activityMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	isEnabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	awsOrg: &awsOrgProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	dataExport: &dataExportProperty{
//   		cloudWatchMetrics: &cloudWatchMetricsProperty{
//   			isEnabled: jsii.Boolean(false),
//   		},
//   		s3BucketDestination: &s3BucketDestinationProperty{
//   			accountId: jsii.String("accountId"),
//   			arn: jsii.String("arn"),
//   			format: jsii.String("format"),
//   			outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   			// the properties below are optional
//   			encryption: encryption,
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   	exclude: &bucketsAndRegionsProperty{
//   		buckets: []*string{
//   			jsii.String("buckets"),
//   		},
//   		regions: []*string{
//   			jsii.String("regions"),
//   		},
//   	},
//   	include: &bucketsAndRegionsProperty{
//   		buckets: []*string{
//   			jsii.String("buckets"),
//   		},
//   		regions: []*string{
//   			jsii.String("regions"),
//   		},
//   	},
//   	storageLensArn: jsii.String("storageLensArn"),
//   }
//
type CfnStorageLens_StorageLensConfigurationProperty struct {
	// This property contains the details of the account-level metrics for Amazon S3 Storage Lens configuration.
	AccountLevel interface{} `field:"required" json:"accountLevel" yaml:"accountLevel"`
	// This property contains the details of the ID of the S3 Storage Lens configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// This property contains the details of whether the Amazon S3 Storage Lens configuration is enabled.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// This property contains the details of the AWS Organization for the S3 Storage Lens configuration.
	AwsOrg interface{} `field:"optional" json:"awsOrg" yaml:"awsOrg"`
	// This property contains the details of this S3 Storage Lens configuration's metrics export.
	DataExport interface{} `field:"optional" json:"dataExport" yaml:"dataExport"`
	// This property contains the details of the bucket and or Regions excluded for Amazon S3 Storage Lens configuration.
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// This property contains the details of the bucket and or Regions included for Amazon S3 Storage Lens configuration.
	Include interface{} `field:"optional" json:"include" yaml:"include"`
	// This property contains the details of the ARN of the S3 Storage Lens configuration.
	//
	// This property is read-only.
	StorageLensArn *string `field:"optional" json:"storageLensArn" yaml:"storageLensArn"`
}

// Properties for defining a `CfnStorageLens`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var encryption interface{}
//
//   cfnStorageLensProps := &cfnStorageLensProps{
//   	storageLensConfiguration: &storageLensConfigurationProperty{
//   		accountLevel: &accountLevelProperty{
//   			bucketLevel: &bucketLevelProperty{
//   				activityMetrics: &activityMetricsProperty{
//   					isEnabled: jsii.Boolean(false),
//   				},
//   				prefixLevel: &prefixLevelProperty{
//   					storageMetrics: &prefixLevelStorageMetricsProperty{
//   						isEnabled: jsii.Boolean(false),
//   						selectionCriteria: &selectionCriteriaProperty{
//   							delimiter: jsii.String("delimiter"),
//   							maxDepth: jsii.Number(123),
//   							minStorageBytesPercentage: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			activityMetrics: &activityMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   		},
//   		id: jsii.String("id"),
//   		isEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		awsOrg: &awsOrgProperty{
//   			arn: jsii.String("arn"),
//   		},
//   		dataExport: &dataExportProperty{
//   			cloudWatchMetrics: &cloudWatchMetricsProperty{
//   				isEnabled: jsii.Boolean(false),
//   			},
//   			s3BucketDestination: &s3BucketDestinationProperty{
//   				accountId: jsii.String("accountId"),
//   				arn: jsii.String("arn"),
//   				format: jsii.String("format"),
//   				outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   				// the properties below are optional
//   				encryption: encryption,
//   				prefix: jsii.String("prefix"),
//   			},
//   		},
//   		exclude: &bucketsAndRegionsProperty{
//   			buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		include: &bucketsAndRegionsProperty{
//   			buckets: []*string{
//   				jsii.String("buckets"),
//   			},
//   			regions: []*string{
//   				jsii.String("regions"),
//   			},
//   		},
//   		storageLensArn: jsii.String("storageLensArn"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnStorageLensProps struct {
	// This resource contains the details Amazon S3 Storage Lens configuration.
	StorageLensConfiguration interface{} `field:"required" json:"storageLensConfiguration" yaml:"storageLensConfiguration"`
	// A set of tags (key–value pairs) to associate with the Storage Lens configuration.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

// Specifies a cross-origin access rule for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsRule := &corsRule{
//   	allowedMethods: []httpMethods{
//   		awscdk.Aws_s3.*httpMethods_GET,
//   	},
//   	allowedOrigins: []*string{
//   		jsii.String("allowedOrigins"),
//   	},
//
//   	// the properties below are optional
//   	allowedHeaders: []*string{
//   		jsii.String("allowedHeaders"),
//   	},
//   	exposedHeaders: []*string{
//   		jsii.String("exposedHeaders"),
//   	},
//   	id: jsii.String("id"),
//   	maxAge: jsii.Number(123),
//   }
//
type CorsRule struct {
	// An HTTP method that you allow the origin to execute.
	AllowedMethods *[]HttpMethods `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Headers that are specified in the Access-Control-Request-Headers header.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// One or more headers in the response that you want customers to be able to access from their applications.
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// A unique identifier for this rule.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

// Notification event types.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myQueue queue
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.addEventNotification(s3.eventType_OBJECT_REMOVED,
//   s3n.NewSqsDestination(myQueue), &notificationKeyFilter{
//   	prefix: jsii.String("foo/"),
//   	suffix: jsii.String(".jpg"),
//   })
//
type EventType string

const (
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED EventType = "OBJECT_CREATED"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_PUT EventType = "OBJECT_CREATED_PUT"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_POST EventType = "OBJECT_CREATED_POST"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_COPY EventType = "OBJECT_CREATED_COPY"
	// Amazon S3 APIs such as PUT, POST, and COPY can create an object.
	//
	// Using
	// these event types, you can enable notification when an object is created
	// using a specific API, or you can use the s3:ObjectCreated:* event type to
	// request notification regardless of the API that was used to create an
	// object.
	EventType_OBJECT_CREATED_COMPLETE_MULTIPART_UPLOAD EventType = "OBJECT_CREATED_COMPLETE_MULTIPART_UPLOAD"
	// By using the ObjectRemoved event types, you can enable notification when an object or a batch of objects is removed from a bucket.
	//
	// You can request notification when an object is deleted or a versioned
	// object is permanently deleted by using the s3:ObjectRemoved:Delete event
	// type. Or you can request notification when a delete marker is created for
	// a versioned object by using s3:ObjectRemoved:DeleteMarkerCreated. For
	// information about deleting versioned objects, see Deleting Object
	// Versions. You can also use a wildcard s3:ObjectRemoved:* to request
	// notification anytime an object is deleted.
	//
	// You will not receive event notifications from automatic deletes from
	// lifecycle policies or from failed operations.
	EventType_OBJECT_REMOVED EventType = "OBJECT_REMOVED"
	// By using the ObjectRemoved event types, you can enable notification when an object or a batch of objects is removed from a bucket.
	//
	// You can request notification when an object is deleted or a versioned
	// object is permanently deleted by using the s3:ObjectRemoved:Delete event
	// type. Or you can request notification when a delete marker is created for
	// a versioned object by using s3:ObjectRemoved:DeleteMarkerCreated. For
	// information about deleting versioned objects, see Deleting Object
	// Versions. You can also use a wildcard s3:ObjectRemoved:* to request
	// notification anytime an object is deleted.
	//
	// You will not receive event notifications from automatic deletes from
	// lifecycle policies or from failed operations.
	EventType_OBJECT_REMOVED_DELETE EventType = "OBJECT_REMOVED_DELETE"
	// By using the ObjectRemoved event types, you can enable notification when an object or a batch of objects is removed from a bucket.
	//
	// You can request notification when an object is deleted or a versioned
	// object is permanently deleted by using the s3:ObjectRemoved:Delete event
	// type. Or you can request notification when a delete marker is created for
	// a versioned object by using s3:ObjectRemoved:DeleteMarkerCreated. For
	// information about deleting versioned objects, see Deleting Object
	// Versions. You can also use a wildcard s3:ObjectRemoved:* to request
	// notification anytime an object is deleted.
	//
	// You will not receive event notifications from automatic deletes from
	// lifecycle policies or from failed operations.
	EventType_OBJECT_REMOVED_DELETE_MARKER_CREATED EventType = "OBJECT_REMOVED_DELETE_MARKER_CREATED"
	// Using restore object event types you can receive notifications for initiation and completion when restoring objects from the S3 Glacier storage class.
	//
	// You use s3:ObjectRestore:Post to request notification of object restoration
	// initiation.
	EventType_OBJECT_RESTORE_POST EventType = "OBJECT_RESTORE_POST"
	// Using restore object event types you can receive notifications for initiation and completion when restoring objects from the S3 Glacier storage class.
	//
	// You use s3:ObjectRestore:Completed to request notification of
	// restoration completion.
	EventType_OBJECT_RESTORE_COMPLETED EventType = "OBJECT_RESTORE_COMPLETED"
	// Using restore object event types you can receive notifications for initiation and completion when restoring objects from the S3 Glacier storage class.
	//
	// You use s3:ObjectRestore:Delete to request notification of
	// restoration completion.
	EventType_OBJECT_RESTORE_DELETE EventType = "OBJECT_RESTORE_DELETE"
	// You can use this event type to request Amazon S3 to send a notification message when Amazon S3 detects that an object of the RRS storage class is lost.
	EventType_REDUCED_REDUNDANCY_LOST_OBJECT EventType = "REDUCED_REDUNDANCY_LOST_OBJECT"
	// You receive this notification event when an object that was eligible for replication using Amazon S3 Replication Time Control failed to replicate.
	EventType_REPLICATION_OPERATION_FAILED_REPLICATION EventType = "REPLICATION_OPERATION_FAILED_REPLICATION"
	// You receive this notification event when an object that was eligible for replication using Amazon S3 Replication Time Control exceeded the 15-minute threshold for replication.
	EventType_REPLICATION_OPERATION_MISSED_THRESHOLD EventType = "REPLICATION_OPERATION_MISSED_THRESHOLD"
	// You receive this notification event for an object that was eligible for replication using the Amazon S3 Replication Time Control feature replicated after the 15-minute threshold.
	EventType_REPLICATION_OPERATION_REPLICATED_AFTER_THRESHOLD EventType = "REPLICATION_OPERATION_REPLICATED_AFTER_THRESHOLD"
	// You receive this notification event for an object that was eligible for replication using Amazon S3 Replication Time Control but is no longer tracked by replication metrics.
	EventType_REPLICATION_OPERATION_NOT_TRACKED EventType = "REPLICATION_OPERATION_NOT_TRACKED"
	// By using the LifecycleExpiration event types, you can receive a notification when Amazon S3 deletes an object based on your S3 Lifecycle configuration.
	EventType_LIFECYCLE_EXPIRATION EventType = "LIFECYCLE_EXPIRATION"
	// The s3:LifecycleExpiration:Delete event type notifies you when an object in an unversioned bucket is deleted.
	//
	// It also notifies you when an object version is permanently deleted by an
	// S3 Lifecycle configuration.
	EventType_LIFECYCLE_EXPIRATION_DELETE EventType = "LIFECYCLE_EXPIRATION_DELETE"
	// The s3:LifecycleExpiration:DeleteMarkerCreated event type notifies you when S3 Lifecycle creates a delete marker when a current version of an object in versioned bucket is deleted.
	EventType_LIFECYCLE_EXPIRATION_DELETE_MARKER_CREATED EventType = "LIFECYCLE_EXPIRATION_DELETE_MARKER_CREATED"
	// You receive this notification event when an object is transitioned to another Amazon S3 storage class by an S3 Lifecycle configuration.
	EventType_LIFECYCLE_TRANSITION EventType = "LIFECYCLE_TRANSITION"
	// You receive this notification event when an object within the S3 Intelligent-Tiering storage class moved to the Archive Access tier or Deep Archive Access tier.
	EventType_INTELLIGENT_TIERING EventType = "INTELLIGENT_TIERING"
	// By using the ObjectTagging event types, you can enable notification when an object tag is added or deleted from an object.
	EventType_OBJECT_TAGGING EventType = "OBJECT_TAGGING"
	// The s3:ObjectTagging:Put event type notifies you when a tag is PUT on an object or an existing tag is updated.
	EventType_OBJECT_TAGGING_PUT EventType = "OBJECT_TAGGING_PUT"
	// The s3:ObjectTagging:Delete event type notifies you when a tag is removed from an object.
	EventType_OBJECT_TAGGING_DELETE EventType = "OBJECT_TAGGING_DELETE"
	// You receive this notification event when an ACL is PUT on an object or when an existing ACL is changed.
	//
	// An event is not generated when a request results in no change to an
	// object’s ACL.
	EventType_OBJECT_ACL_PUT EventType = "OBJECT_ACL_PUT"
)

// All http request methods.
type HttpMethods string

const (
	// The GET method requests a representation of the specified resource.
	HttpMethods_GET HttpMethods = "GET"
	// The PUT method replaces all current representations of the target resource with the request payload.
	HttpMethods_PUT HttpMethods = "PUT"
	// The HEAD method asks for a response identical to that of a GET request, but without the response body.
	HttpMethods_HEAD HttpMethods = "HEAD"
	// The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
	HttpMethods_POST HttpMethods = "POST"
	// The DELETE method deletes the specified resource.
	HttpMethods_DELETE HttpMethods = "DELETE"
)

type IBucket interface {
	awscdk.IResource
	// Adds a bucket notification event destination.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   var myLambda function
	//
	//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
	//   bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &notificationKeyFilter{
	//   	prefix: jsii.String("home/myusername/*"),
	//   })
	//
	// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
	//
	AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Subscribes a destination to receive notifications when an object is created in the bucket.
	//
	// This is identical to calling
	// `onEvent(s3.EventType.OBJECT_CREATED)`.
	AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Subscribes a destination to receive notifications when an object is removed from the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_REMOVED)`.
	AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter)
	// Adds a statement to the resource policy for a principal (i.e. account/role/service) to perform actions on this bucket and/or its contents. Use `bucketArn` and `arnForObjects(keys)` to obtain ARNs for this bucket or objects.
	//
	// Note that the policy statement may or may not be added to the policy.
	// For example, when an `IBucket` is created from an existing bucket,
	// it's not possible to tell whether the bucket already has a policy
	// attached, let alone to re-use that policy to add more statements to it.
	// So it's safest to do nothing in these cases.
	//
	// Returns: metadata about the execution of this method. If the policy
	// was not added, the value of `statementAdded` will be `false`. You
	// should always check this value to make sure that the operation was
	// actually carried out. Otherwise, synthesis and deploy will terminate
	// silently, which may be confusing.
	AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Returns an ARN that represents all objects within the bucket that match the key pattern specified.
	//
	// To represent all keys, specify ``"*"``.
	ArnForObjects(keyPattern *string) *string
	// Grants s3:DeleteObject* permission to an IAM principal for objects in this bucket.
	GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Allows unrestricted access to objects from this bucket.
	//
	// IMPORTANT: This permission allows anyone to perform actions on S3 objects
	// in this bucket, which is useful for when you configure your bucket as a
	// website and want everyone to be able to read objects in the bucket without
	// needing to authenticate.
	//
	// Without arguments, this method will grant read ("s3:GetObject") access to
	// all objects ("*") in the bucket.
	//
	// The method returns the `iam.Grant` object, which can then be modified
	// as needed. For example, you can add a condition that will restrict access only
	// to an IPv4 range like this:
	//
	//      const grant = bucket.grantPublicAccess();
	//      grant.resourceStatement!.addCondition(‘IpAddress’, { “aws:SourceIp”: “54.240.143.0/24” });
	//
	// Returns: The `iam.PolicyStatement` object, which can be used to apply e.g. conditions.
	GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant
	// Grants s3:PutObject* and s3:Abort* permissions for this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant the given IAM identity permissions to modify the ACLs of objects in the given Bucket.
	//
	// If your application has the '@aws-cdk/aws-s3:grantWriteWithoutAcl' feature flag set,
	// calling {@link grantWrite} or {@link grantReadWrite} no longer grants permissions to modify the ACLs of the objects;
	// in this case, if you need to modify object ACLs, call this method explicitly.
	GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant
	// Grant read permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grants read/write permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to use the key for
	// encrypt/decrypt will also be granted.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Grant write permissions to this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant
	// Defines a CloudWatch event that triggers when something happens to this bucket.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object is uploaded to the specified paths (keys) in this bucket using the PutObject API call.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using `onCloudTrailWriteObject` may be preferable.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object at the specified paths (keys) in this bucket are written to.
	//
	// This includes
	// the events PutObject, CopyObject, and CompleteMultipartUpload.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using this method may be preferable to `onCloudTrailPutObject`.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule
	// The S3 URL of an S3 object.
	//
	// For example:
	// - `s3://onlybucket`
	// - `s3://bucket/key`.
	//
	// Returns: an ObjectS3Url token.
	S3UrlForObject(key *string) *string
	// The https Transfer Acceleration URL of an S3 object.
	//
	// Specify `dualStack: true` at the options
	// for dual-stack endpoint (connect to the bucket over IPv6). For example:
	//
	// - `https://bucket.s3-accelerate.amazonaws.com`
	// - `https://bucket.s3-accelerate.amazonaws.com/key`
	//
	// Returns: an TransferAccelerationUrl token.
	TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string
	// The https URL of an S3 object. For example:.
	//
	// - `https://s3.us-west-1.amazonaws.com/onlybucket`
	// - `https://s3.us-west-1.amazonaws.com/bucket/key`
	// - `https://s3.cn-north-1.amazonaws.com.cn/china-bucket/mykey`
	//
	// Returns: an ObjectS3Url token.
	UrlForObject(key *string) *string
	// The virtual hosted-style URL of an S3 object. Specify `regional: false` at the options for non-regional URL. For example:.
	//
	// - `https://only-bucket.s3.us-west-1.amazonaws.com`
	// - `https://bucket.s3.us-west-1.amazonaws.com/key`
	// - `https://bucket.s3.amazonaws.com/key`
	// - `https://china-bucket.s3.cn-north-1.amazonaws.com.cn/mykey`
	//
	// Returns: an ObjectS3Url token.
	VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string
	// The ARN of the bucket.
	BucketArn() *string
	// The IPv4 DNS name of the specified bucket.
	BucketDomainName() *string
	// The IPv6 DNS name of the specified bucket.
	BucketDualStackDomainName() *string
	// The name of the bucket.
	BucketName() *string
	// The regional domain name of the specified bucket.
	BucketRegionalDomainName() *string
	// The Domain name of the static website.
	BucketWebsiteDomainName() *string
	// The URL of the static website.
	BucketWebsiteUrl() *string
	// Optional KMS encryption key associated with this bucket.
	EncryptionKey() awskms.IKey
	// If this bucket has been configured for static website hosting.
	IsWebsite() *bool
	// The resource policy associated with this bucket.
	//
	// If `autoCreatePolicy` is true, a `BucketPolicy` will be created upon the
	// first call to addToResourcePolicy(s).
	Policy() BucketPolicy
	SetPolicy(p BucketPolicy)
}

// The jsii proxy for IBucket
type jsiiProxy_IBucket struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBucket) AddEventNotification(event EventType, dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{event, dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addEventNotification",
		args,
	)
}

func (i *jsiiProxy_IBucket) AddObjectCreatedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addObjectCreatedNotification",
		args,
	)
}

func (i *jsiiProxy_IBucket) AddObjectRemovedNotification(dest IBucketNotificationDestination, filters ...*NotificationKeyFilter) {
	args := []interface{}{dest}
	for _, a := range filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addObjectRemovedNotification",
		args,
	)
}

func (i *jsiiProxy_IBucket) AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{permission},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) ArnForObjects(keyPattern *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"arnForObjects",
		[]interface{}{keyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDelete",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantPublicAccess(keyPrefix *string, allowedActions ...*string) awsiam.Grant {
	args := []interface{}{keyPrefix}
	for _, a := range allowedActions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPublicAccess",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantPut(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPut",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantPutAcl(identity awsiam.IGrantable, objectsKeyPattern *string) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPutAcl",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantRead(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantReadWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) OnCloudTrailPutObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailPutObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) OnCloudTrailWriteObject(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onCloudTrailWriteObject",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) S3UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"s3UrlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) TransferAccelerationUrlForObject(key *string, options *TransferAccelerationUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"transferAccelerationUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) UrlForObject(key *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"urlForObject",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) VirtualHostedUrlForObject(key *string, options *VirtualHostedStyleUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"virtualHostedUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IBucket) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) BucketDualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) BucketRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) BucketWebsiteDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) BucketWebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) IsWebsite() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isWebsite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) Policy() BucketPolicy {
	var returns BucketPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBucket) SetPolicy(val BucketPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

// Implemented by constructs that can be used as bucket notification destinations.
type IBucketNotificationDestination interface {
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	Bind(scope constructs.Construct, bucket IBucket) *BucketNotificationDestinationConfig
}

// The jsii proxy for IBucketNotificationDestination
type jsiiProxy_IBucketNotificationDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IBucketNotificationDestination) Bind(scope constructs.Construct, bucket IBucket) *BucketNotificationDestinationConfig {
	var returns *BucketNotificationDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, bucket},
		&returns,
	)

	return returns
}

// The intelligent tiering configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   intelligentTieringConfiguration := &intelligentTieringConfiguration{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	archiveAccessTierTime: cdk.duration.minutes(jsii.Number(30)),
//   	deepArchiveAccessTierTime: cdk.*duration.minutes(jsii.Number(30)),
//   	prefix: jsii.String("prefix"),
//   	tags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type IntelligentTieringConfiguration struct {
	// Configuration name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 90 days to the Archive Access tier.
	ArchiveAccessTierTime awscdk.Duration `field:"optional" json:"archiveAccessTierTime" yaml:"archiveAccessTierTime"`
	// When enabled, Intelligent-Tiering will automatically move objects that haven’t been accessed for a minimum of 180 days to the Deep Archive Access tier.
	DeepArchiveAccessTierTime awscdk.Duration `field:"optional" json:"deepArchiveAccessTierTime" yaml:"deepArchiveAccessTierTime"`
	// Add a filter to limit the scope of this configuration to a single prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// You can limit the scope of this rule to the key value pairs added below.
	Tags *[]*Tag `field:"optional" json:"tags" yaml:"tags"`
}

// Specifies the inventory configuration of an S3 Bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   inventory := &inventory{
//   	destination: &inventoryDestination{
//   		bucket: bucket,
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	format: awscdk.Aws_s3.inventoryFormat_CSV,
//   	frequency: awscdk.*Aws_s3.inventoryFrequency_DAILY,
//   	includeObjectVersions: awscdk.*Aws_s3.inventoryObjectVersion_ALL,
//   	inventoryId: jsii.String("inventoryId"),
//   	objectsPrefix: jsii.String("objectsPrefix"),
//   	optionalFields: []*string{
//   		jsii.String("optionalFields"),
//   	},
//   }
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
//
type Inventory struct {
	// The destination of the inventory.
	Destination *InventoryDestination `field:"required" json:"destination" yaml:"destination"`
	// Whether the inventory is enabled or disabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The format of the inventory.
	Format InventoryFormat `field:"optional" json:"format" yaml:"format"`
	// Frequency at which the inventory should be generated.
	Frequency InventoryFrequency `field:"optional" json:"frequency" yaml:"frequency"`
	// If the inventory should contain all the object versions or only the current one.
	IncludeObjectVersions InventoryObjectVersion `field:"optional" json:"includeObjectVersions" yaml:"includeObjectVersions"`
	// The inventory configuration ID.
	InventoryId *string `field:"optional" json:"inventoryId" yaml:"inventoryId"`
	// The inventory will only include objects that meet the prefix filter criteria.
	ObjectsPrefix *string `field:"optional" json:"objectsPrefix" yaml:"objectsPrefix"`
	// A list of optional fields to be included in the inventory result.
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
}

// The destination of the inventory.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   inventoryBucket := s3.NewBucket(this, jsii.String("InventoryBucket"))
//
//   dataBucket := s3.NewBucket(this, jsii.String("DataBucket"), &bucketProps{
//   	inventories: []inventory{
//   		&inventory{
//   			frequency: s3.inventoryFrequency_DAILY,
//   			includeObjectVersions: s3.inventoryObjectVersion_CURRENT,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   			},
//   		},
//   		&inventory{
//   			frequency: s3.*inventoryFrequency_WEEKLY,
//   			includeObjectVersions: s3.*inventoryObjectVersion_ALL,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   				prefix: jsii.String("with-all-versions"),
//   			},
//   		},
//   	},
//   })
//
type InventoryDestination struct {
	// Bucket where all inventories will be saved in.
	Bucket IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The account ID that owns the destination S3 bucket.
	//
	// If no account ID is provided, the owner is not validated before exporting data.
	// It's recommended to set an account ID to prevent problems if the destination bucket ownership changes.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The prefix to be used when saving the inventory.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

// All supported inventory list formats.
type InventoryFormat string

const (
	// Generate the inventory list as CSV.
	InventoryFormat_CSV InventoryFormat = "CSV"
	// Generate the inventory list as Parquet.
	InventoryFormat_PARQUET InventoryFormat = "PARQUET"
	// Generate the inventory list as ORC.
	InventoryFormat_ORC InventoryFormat = "ORC"
)

// All supported inventory frequencies.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   inventoryBucket := s3.NewBucket(this, jsii.String("InventoryBucket"))
//
//   dataBucket := s3.NewBucket(this, jsii.String("DataBucket"), &bucketProps{
//   	inventories: []inventory{
//   		&inventory{
//   			frequency: s3.inventoryFrequency_DAILY,
//   			includeObjectVersions: s3.inventoryObjectVersion_CURRENT,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   			},
//   		},
//   		&inventory{
//   			frequency: s3.*inventoryFrequency_WEEKLY,
//   			includeObjectVersions: s3.*inventoryObjectVersion_ALL,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   				prefix: jsii.String("with-all-versions"),
//   			},
//   		},
//   	},
//   })
//
type InventoryFrequency string

const (
	// A report is generated every day.
	InventoryFrequency_DAILY InventoryFrequency = "DAILY"
	// A report is generated every Sunday (UTC timezone) after the initial report.
	InventoryFrequency_WEEKLY InventoryFrequency = "WEEKLY"
)

// Inventory version support.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   inventoryBucket := s3.NewBucket(this, jsii.String("InventoryBucket"))
//
//   dataBucket := s3.NewBucket(this, jsii.String("DataBucket"), &bucketProps{
//   	inventories: []inventory{
//   		&inventory{
//   			frequency: s3.inventoryFrequency_DAILY,
//   			includeObjectVersions: s3.inventoryObjectVersion_CURRENT,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   			},
//   		},
//   		&inventory{
//   			frequency: s3.*inventoryFrequency_WEEKLY,
//   			includeObjectVersions: s3.*inventoryObjectVersion_ALL,
//   			destination: &inventoryDestination{
//   				bucket: inventoryBucket,
//   				prefix: jsii.String("with-all-versions"),
//   			},
//   		},
//   	},
//   })
//
type InventoryObjectVersion string

const (
	// Includes all versions of each object in the report.
	InventoryObjectVersion_ALL InventoryObjectVersion = "ALL"
	// Includes only the current version of each object in the report.
	InventoryObjectVersion_CURRENT InventoryObjectVersion = "CURRENT"
)

// Declaration of a Life cycle rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var storageClass storageClass
//   var tagFilters interface{}
//
//   lifecycleRule := &lifecycleRule{
//   	abortIncompleteMultipartUploadAfter: cdk.duration.minutes(jsii.Number(30)),
//   	enabled: jsii.Boolean(false),
//   	expiration: cdk.*duration.minutes(jsii.Number(30)),
//   	expirationDate: NewDate(),
//   	expiredObjectDeleteMarker: jsii.Boolean(false),
//   	id: jsii.String("id"),
//   	noncurrentVersionExpiration: cdk.*duration.minutes(jsii.Number(30)),
//   	noncurrentVersionsToRetain: jsii.Number(123),
//   	noncurrentVersionTransitions: []noncurrentVersionTransition{
//   		&noncurrentVersionTransition{
//   			storageClass: storageClass,
//   			transitionAfter: cdk.*duration.minutes(jsii.Number(30)),
//
//   			// the properties below are optional
//   			noncurrentVersionsToRetain: jsii.Number(123),
//   		},
//   	},
//   	objectSizeGreaterThan: jsii.Number(123),
//   	prefix: jsii.String("prefix"),
//   	tagFilters: map[string]interface{}{
//   		"tagFiltersKey": tagFilters,
//   	},
//   	transitions: []transition{
//   		&transition{
//   			storageClass: storageClass,
//
//   			// the properties below are optional
//   			transitionAfter: cdk.*duration.minutes(jsii.Number(30)),
//   			transitionDate: NewDate(),
//   		},
//   	},
//   }
//
type LifecycleRule struct {
	// Specifies a lifecycle rule that aborts incomplete multipart uploads to an Amazon S3 bucket.
	//
	// The AbortIncompleteMultipartUpload property type creates a lifecycle
	// rule that aborts incomplete multipart uploads to an Amazon S3 bucket.
	// When Amazon S3 aborts a multipart upload, it deletes all parts
	// associated with the multipart upload.
	AbortIncompleteMultipartUploadAfter awscdk.Duration `field:"optional" json:"abortIncompleteMultipartUploadAfter" yaml:"abortIncompleteMultipartUploadAfter"`
	// Whether this rule is enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon Glacier.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	Expiration awscdk.Duration `field:"optional" json:"expiration" yaml:"expiration"`
	// Indicates when objects are deleted from Amazon S3 and Amazon Glacier.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	ExpirationDate *time.Time `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions.
	//
	// If set to true, the delete marker will be expired.
	ExpiredObjectDeleteMarker *bool `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
	// A unique identifier for this rule.
	//
	// The value cannot be more than 255 characters.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Time between when a new version of the object is uploaded to the bucket and when old versions of the object expire.
	//
	// For buckets with versioning enabled (or suspended), specifies the time,
	// in days, between when a new version of the object is uploaded to the
	// bucket and when old versions of the object expire. When object versions
	// expire, Amazon S3 permanently deletes them. If you specify a transition
	// and expiration time, the expiration time must be later than the
	// transition time.
	NoncurrentVersionExpiration awscdk.Duration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// Indicates a maximum number of noncurrent versions to retain.
	//
	// If there are this many more noncurrent versions,
	// Amazon S3 permanently deletes them.
	NoncurrentVersionsToRetain *float64 `field:"optional" json:"noncurrentVersionsToRetain" yaml:"noncurrentVersionsToRetain"`
	// One or more transition rules that specify when non-current objects transition to a specified storage class.
	//
	// Only for for buckets with versioning enabled (or suspended).
	//
	// If you specify a transition and expiration time, the expiration time
	// must be later than the transition time.
	NoncurrentVersionTransitions *[]*NoncurrentVersionTransition `field:"optional" json:"noncurrentVersionTransitions" yaml:"noncurrentVersionTransitions"`
	// Specifies the minimum object size in bytes for this rule to apply to.
	ObjectSizeGreaterThan *float64 `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Object key prefix that identifies one or more objects to which this rule applies.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The TagFilter property type specifies tags to use to identify a subset of objects for an Amazon S3 bucket.
	TagFilters *map[string]interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// One or more transition rules that specify when an object transitions to a specified storage class.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	Transitions *[]*Transition `field:"optional" json:"transitions" yaml:"transitions"`
}

// An interface that represents the location of a specific object in an S3 Bucket.
//
// Example:
//   startQueryExecutionJob := tasks.NewAthenaStartQueryExecution(this, jsii.String("Athena Start Query"), &athenaStartQueryExecutionProps{
//   	queryString: sfn.jsonPath.format(jsii.String("select contacts where year={};"), sfn.*jsonPath.stringAt(jsii.String("$.year"))),
//   	queryExecutionContext: &queryExecutionContext{
//   		databaseName: jsii.String("interactions"),
//   	},
//   	resultConfiguration: &resultConfiguration{
//   		encryptionConfiguration: &encryptionConfiguration{
//   			encryptionOption: tasks.encryptionOption_S3_MANAGED,
//   		},
//   		outputLocation: &location{
//   			bucketName: jsii.String("mybucket"),
//   			objectKey: jsii.String("myprefix"),
//   		},
//   	},
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   })
//
type Location struct {
	// The name of the S3 Bucket the object is in.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The path inside the Bucket where the object is located at.
	ObjectKey *string `field:"required" json:"objectKey" yaml:"objectKey"`
	// The S3 object version.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

// Describes when noncurrent versions transition to a specified storage class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var storageClass storageClass
//
//   noncurrentVersionTransition := &noncurrentVersionTransition{
//   	storageClass: storageClass,
//   	transitionAfter: cdk.duration.minutes(jsii.Number(30)),
//
//   	// the properties below are optional
//   	noncurrentVersionsToRetain: jsii.Number(123),
//   }
//
type NoncurrentVersionTransition struct {
	// The storage class to which you want the object to transition.
	StorageClass StorageClass `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	TransitionAfter awscdk.Duration `field:"required" json:"transitionAfter" yaml:"transitionAfter"`
	// Indicates the number of noncurrent version objects to be retained.
	//
	// Can be up to 100 noncurrent versions retained.
	NoncurrentVersionsToRetain *float64 `field:"optional" json:"noncurrentVersionsToRetain" yaml:"noncurrentVersionsToRetain"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myQueue queue
//
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.addEventNotification(s3.eventType_OBJECT_REMOVED,
//   s3n.NewSqsDestination(myQueue), &notificationKeyFilter{
//   	prefix: jsii.String("foo/"),
//   	suffix: jsii.String(".jpg"),
//   })
//
type NotificationKeyFilter struct {
	// S3 keys must have the specified prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// S3 keys must have the specified suffix.
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

// The ObjectOwnership of the bucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
//   	objectOwnership: s3.objectOwnership_OBJECT_WRITER,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
//
type ObjectOwnership string

const (
	// ACLs are disabled, and the bucket owner automatically owns and has full control over every object in the bucket.
	//
	// ACLs no longer affect permissions to data in the S3 bucket.
	// The bucket uses policies to define access control.
	ObjectOwnership_BUCKET_OWNER_ENFORCED ObjectOwnership = "BUCKET_OWNER_ENFORCED"
	// Objects uploaded to the bucket change ownership to the bucket owner .
	ObjectOwnership_BUCKET_OWNER_PREFERRED ObjectOwnership = "BUCKET_OWNER_PREFERRED"
	// The uploading account will own the object.
	ObjectOwnership_OBJECT_WRITER ObjectOwnership = "OBJECT_WRITER"
)

// Options for the onCloudTrailPutObject method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var detail interface{}
//   var ruleTarget iRuleTarget
//
//   onCloudTrailBucketEventOptions := &onCloudTrailBucketEventOptions{
//   	description: jsii.String("description"),
//   	eventPattern: &eventPattern{
//   		account: []*string{
//   			jsii.String("account"),
//   		},
//   		detail: map[string]interface{}{
//   			"detailKey": detail,
//   		},
//   		detailType: []*string{
//   			jsii.String("detailType"),
//   		},
//   		id: []*string{
//   			jsii.String("id"),
//   		},
//   		region: []*string{
//   			jsii.String("region"),
//   		},
//   		resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		source: []*string{
//   			jsii.String("source"),
//   		},
//   		time: []*string{
//   			jsii.String("time"),
//   		},
//   		version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	paths: []*string{
//   		jsii.String("paths"),
//   	},
//   	ruleName: jsii.String("ruleName"),
//   	target: ruleTarget,
//   }
//
type OnCloudTrailBucketEventOptions struct {
	// A description of the rule's purpose.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	EventPattern *awsevents.EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	Target awsevents.IRuleTarget `field:"optional" json:"target" yaml:"target"`
	// Only watch changes to these object paths.
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

// All http request methods.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
//   	websiteRoutingRules: []routingRule{
//   		&routingRule{
//   			hostName: jsii.String("www.example.com"),
//   			httpRedirectCode: jsii.String("302"),
//   			protocol: s3.redirectProtocol_HTTPS,
//   			replaceKey: s3.replaceKey.prefixWith(jsii.String("test/")),
//   			condition: &routingRuleCondition{
//   				httpErrorCodeReturnedEquals: jsii.String("200"),
//   				keyPrefixEquals: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   })
//
type RedirectProtocol string

const (
	RedirectProtocol_HTTP RedirectProtocol = "HTTP"
	RedirectProtocol_HTTPS RedirectProtocol = "HTTPS"
)

// Specifies a redirect behavior of all requests to a website endpoint of a bucket.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
//   	websiteRedirect: &redirectTarget{
//   		hostName: jsii.String("www.example.com"),
//   	},
//   })
//
type RedirectTarget struct {
	// Name of the host where requests are redirected.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Protocol to use when redirecting requests.
	Protocol RedirectProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
//   	websiteRoutingRules: []routingRule{
//   		&routingRule{
//   			hostName: jsii.String("www.example.com"),
//   			httpRedirectCode: jsii.String("302"),
//   			protocol: s3.redirectProtocol_HTTPS,
//   			replaceKey: s3.replaceKey.prefixWith(jsii.String("test/")),
//   			condition: &routingRuleCondition{
//   				httpErrorCodeReturnedEquals: jsii.String("200"),
//   				keyPrefixEquals: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   })
//
type ReplaceKey interface {
	PrefixWithKey() *string
	WithKey() *string
}

// The jsii proxy struct for ReplaceKey
type jsiiProxy_ReplaceKey struct {
	_ byte // padding
}

func (j *jsiiProxy_ReplaceKey) PrefixWithKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixWithKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplaceKey) WithKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"withKey",
		&returns,
	)
	return returns
}


// The object key prefix to use in the redirect request.
func ReplaceKey_PrefixWith(keyReplacement *string) ReplaceKey {
	_init_.Initialize()

	var returns ReplaceKey

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.ReplaceKey",
		"prefixWith",
		[]interface{}{keyReplacement},
		&returns,
	)

	return returns
}

// The specific object key to use in the redirect request.
func ReplaceKey_With(keyReplacement *string) ReplaceKey {
	_init_.Initialize()

	var returns ReplaceKey

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3.ReplaceKey",
		"with",
		[]interface{}{keyReplacement},
		&returns,
	)

	return returns
}

// Rule that define when a redirect is applied and the redirect behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var replaceKey replaceKey
//
//   routingRule := &routingRule{
//   	condition: &routingRuleCondition{
//   		httpErrorCodeReturnedEquals: jsii.String("httpErrorCodeReturnedEquals"),
//   		keyPrefixEquals: jsii.String("keyPrefixEquals"),
//   	},
//   	hostName: jsii.String("hostName"),
//   	httpRedirectCode: jsii.String("httpRedirectCode"),
//   	protocol: awscdk.Aws_s3.redirectProtocol_HTTP,
//   	replaceKey: replaceKey,
//   }
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html
//
type RoutingRule struct {
	// Specifies a condition that must be met for the specified redirect to apply.
	Condition *RoutingRuleCondition `field:"optional" json:"condition" yaml:"condition"`
	// The host name to use in the redirect request.
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The HTTP redirect code to use on the response.
	HttpRedirectCode *string `field:"optional" json:"httpRedirectCode" yaml:"httpRedirectCode"`
	// Protocol to use when redirecting requests.
	Protocol RedirectProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Specifies the object key prefix to use in the redirect request.
	ReplaceKey ReplaceKey `field:"optional" json:"replaceKey" yaml:"replaceKey"`
}

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
//   	websiteRoutingRules: []routingRule{
//   		&routingRule{
//   			hostName: jsii.String("www.example.com"),
//   			httpRedirectCode: jsii.String("302"),
//   			protocol: s3.redirectProtocol_HTTPS,
//   			replaceKey: s3.replaceKey.prefixWith(jsii.String("test/")),
//   			condition: &routingRuleCondition{
//   				httpErrorCodeReturnedEquals: jsii.String("200"),
//   				keyPrefixEquals: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   })
//
type RoutingRuleCondition struct {
	// The HTTP error code when the redirect is applied.
	//
	// In the event of an error, if the error code equals this value, then the specified redirect is applied.
	//
	// If both condition properties are specified, both must be true for the redirect to be applied.
	HttpErrorCodeReturnedEquals *string `field:"optional" json:"httpErrorCodeReturnedEquals" yaml:"httpErrorCodeReturnedEquals"`
	// The object key name prefix when the redirect is applied.
	//
	// If both condition properties are specified, both must be true for the redirect to be applied.
	KeyPrefixEquals *string `field:"optional" json:"keyPrefixEquals" yaml:"keyPrefixEquals"`
}

// Storage class to move an object to.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
//   	lifecycleRules: []lifecycleRule{
//   		&lifecycleRule{
//   			abortIncompleteMultipartUploadAfter: cdk.duration.minutes(jsii.Number(30)),
//   			enabled: jsii.Boolean(false),
//   			expiration: cdk.*duration.days(jsii.Number(30)),
//   			expirationDate: NewDate(),
//   			expiredObjectDeleteMarker: jsii.Boolean(false),
//   			id: jsii.String("id"),
//   			noncurrentVersionExpiration: cdk.*duration.days(jsii.Number(30)),
//
//   			// the properties below are optional
//   			noncurrentVersionsToRetain: jsii.Number(123),
//   			noncurrentVersionTransitions: []noncurrentVersionTransition{
//   				&noncurrentVersionTransition{
//   					storageClass: s3.storageClass_GLACIER(),
//   					transitionAfter: cdk.*duration.days(jsii.Number(30)),
//
//   					// the properties below are optional
//   					noncurrentVersionsToRetain: jsii.Number(123),
//   				},
//   			},
//   			objectSizeGreaterThan: jsii.Number(500),
//   			prefix: jsii.String("prefix"),
//   			transitions: []transition{
//   				&transition{
//   					storageClass: s3.*storageClass_GLACIER(),
//
//   					// the properties below are optional
//   					transitionAfter: cdk.*duration.days(jsii.Number(30)),
//   					transitionDate: NewDate(),
//   				},
//   			},
//   		},
//   	},
//   })
//
type StorageClass interface {
	Value() *string
	ToString() *string
}

// The jsii proxy struct for StorageClass
type jsiiProxy_StorageClass struct {
	_ byte // padding
}

func (j *jsiiProxy_StorageClass) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewStorageClass(value *string) StorageClass {
	_init_.Initialize()

	j := jsiiProxy_StorageClass{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.StorageClass",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewStorageClass_Override(s StorageClass, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3.StorageClass",
		[]interface{}{value},
		s,
	)
}

func StorageClass_DEEP_ARCHIVE() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"DEEP_ARCHIVE",
		&returns,
	)
	return returns
}

func StorageClass_GLACIER() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"GLACIER",
		&returns,
	)
	return returns
}

func StorageClass_GLACIER_INSTANT_RETRIEVAL() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"GLACIER_INSTANT_RETRIEVAL",
		&returns,
	)
	return returns
}

func StorageClass_INFREQUENT_ACCESS() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"INFREQUENT_ACCESS",
		&returns,
	)
	return returns
}

func StorageClass_INTELLIGENT_TIERING() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"INTELLIGENT_TIERING",
		&returns,
	)
	return returns
}

func StorageClass_ONE_ZONE_INFREQUENT_ACCESS() StorageClass {
	_init_.Initialize()
	var returns StorageClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_s3.StorageClass",
		"ONE_ZONE_INFREQUENT_ACCESS",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageClass) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tag := &tag{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type Tag struct {
	// key to e tagged.
	Key *string `field:"required" json:"key" yaml:"key"`
	// additional value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

// Options for creating a Transfer Acceleration URL.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transferAccelerationUrlOptions := &transferAccelerationUrlOptions{
//   	dualStack: jsii.Boolean(false),
//   }
//
type TransferAccelerationUrlOptions struct {
	// Dual-stack support to connect to the bucket over IPv6.
	DualStack *bool `field:"optional" json:"dualStack" yaml:"dualStack"`
}

// Describes when an object transitions to a specified storage class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var storageClass storageClass
//
//   transition := &transition{
//   	storageClass: storageClass,
//
//   	// the properties below are optional
//   	transitionAfter: cdk.duration.minutes(jsii.Number(30)),
//   	transitionDate: NewDate(),
//   }
//
type Transition struct {
	// The storage class to which you want the object to transition.
	StorageClass StorageClass `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	TransitionAfter awscdk.Duration `field:"optional" json:"transitionAfter" yaml:"transitionAfter"`
	// Indicates when objects are transitioned to the specified storage class.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	TransitionDate *time.Time `field:"optional" json:"transitionDate" yaml:"transitionDate"`
}

// Options for creating Virtual-Hosted style URL.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   bucket.urlForObject(jsii.String("objectname")) // Path-Style URL
//   bucket.virtualHostedUrlForObject(jsii.String("objectname")) // Virtual Hosted-Style URL
//   bucket.virtualHostedUrlForObject(jsii.String("objectname"), &virtualHostedStyleUrlOptions{
//   	regional: jsii.Boolean(false),
//   })
//
type VirtualHostedStyleUrlOptions struct {
	// Specifies the URL includes the region.
	Regional *bool `field:"optional" json:"regional" yaml:"regional"`
}

