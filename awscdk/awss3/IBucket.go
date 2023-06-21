package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
)

type IBucket interface {
	awscdk.IResource
	// Adds a bucket notification event destination.
	//
	// Example:
	//   var myLambda function
	//
	//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
	//   bucket.AddEventNotification(s3.EventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &NotificationKeyFilter{
	//   	Prefix: jsii.String("home/myusername/*"),
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
	// Enables event bridge notification, causing all events below to be sent to EventBridge:.
	//
	// - Object Deleted (DeleteObject)
	// - Object Deleted (Lifecycle expiration)
	// - Object Restore Initiated
	// - Object Restore Completed
	// - Object Restore Expired
	// - Object Storage Class Changed
	// - Object Access Tier Changed
	// - Object ACL Updated
	// - Object Tags Added
	// - Object Tags Deleted.
	EnableEventBridgeNotification()
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
	//     const grant = bucket.grantPublicAccess();
	//     grant.resourceStatement!.addCondition(‘IpAddress’, { “aws:SourceIp”: “54.240.143.0/24” });
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
	// calling `grantWrite` or `grantReadWrite` no longer grants permissions to modify the ACLs of the objects;
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
	// use the `grantPutAcl` method.
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
	// use the `grantPutAcl` method.
	GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}, allowedActionPatterns *[]*string) awsiam.Grant
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
	if err := i.validateAddEventNotificationParameters(event, dest, &filters); err != nil {
		panic(err)
	}
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
	if err := i.validateAddObjectCreatedNotificationParameters(dest, &filters); err != nil {
		panic(err)
	}
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
	if err := i.validateAddObjectRemovedNotificationParameters(dest, &filters); err != nil {
		panic(err)
	}
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
	if err := i.validateAddToResourcePolicyParameters(permission); err != nil {
		panic(err)
	}
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
	if err := i.validateArnForObjectsParameters(keyPattern); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"arnForObjects",
		[]interface{}{keyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) EnableEventBridgeNotification() {
	_jsii_.InvokeVoid(
		i,
		"enableEventBridgeNotification",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IBucket) GrantDelete(identity awsiam.IGrantable, objectsKeyPattern interface{}) awsiam.Grant {
	if err := i.validateGrantDeleteParameters(identity); err != nil {
		panic(err)
	}
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
	if err := i.validateGrantPutParameters(identity); err != nil {
		panic(err)
	}
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
	if err := i.validateGrantPutAclParameters(identity); err != nil {
		panic(err)
	}
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
	if err := i.validateGrantReadParameters(identity); err != nil {
		panic(err)
	}
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
	if err := i.validateGrantReadWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadWrite",
		[]interface{}{identity, objectsKeyPattern},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) GrantWrite(identity awsiam.IGrantable, objectsKeyPattern interface{}, allowedActionPatterns *[]*string) awsiam.Grant {
	if err := i.validateGrantWriteParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{identity, objectsKeyPattern, allowedActionPatterns},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBucket) OnCloudTrailEvent(id *string, options *OnCloudTrailBucketEventOptions) awsevents.Rule {
	if err := i.validateOnCloudTrailEventParameters(id, options); err != nil {
		panic(err)
	}
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
	if err := i.validateOnCloudTrailPutObjectParameters(id, options); err != nil {
		panic(err)
	}
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
	if err := i.validateOnCloudTrailWriteObjectParameters(id, options); err != nil {
		panic(err)
	}
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
	if err := i.validateTransferAccelerationUrlForObjectParameters(options); err != nil {
		panic(err)
	}
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
	if err := i.validateVirtualHostedUrlForObjectParameters(options); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_IBucket)SetPolicy(val BucketPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

