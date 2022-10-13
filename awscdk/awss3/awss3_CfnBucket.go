package awss3

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

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
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
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

func (j *jsiiProxy_CfnBucket) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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

	if err := validateNewCfnBucketParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnBucket)SetAccelerateConfiguration(val interface{}) {
	if err := j.validateSetAccelerateConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accelerateConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetAccessControl(val *string) {
	_jsii_.Set(
		j,
		"accessControl",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetAnalyticsConfigurations(val interface{}) {
	if err := j.validateSetAnalyticsConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analyticsConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetBucketEncryption(val interface{}) {
	if err := j.validateSetBucketEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketEncryption",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetBucketName(val *string) {
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetCorsConfiguration(val interface{}) {
	if err := j.validateSetCorsConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"corsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetIntelligentTieringConfigurations(val interface{}) {
	if err := j.validateSetIntelligentTieringConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intelligentTieringConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetInventoryConfigurations(val interface{}) {
	if err := j.validateSetInventoryConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inventoryConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetLifecycleConfiguration(val interface{}) {
	if err := j.validateSetLifecycleConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetLoggingConfiguration(val interface{}) {
	if err := j.validateSetLoggingConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetMetricsConfigurations(val interface{}) {
	if err := j.validateSetMetricsConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetNotificationConfiguration(val interface{}) {
	if err := j.validateSetNotificationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetObjectLockConfiguration(val interface{}) {
	if err := j.validateSetObjectLockConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetObjectLockEnabled(val interface{}) {
	if err := j.validateSetObjectLockEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetOwnershipControls(val interface{}) {
	if err := j.validateSetOwnershipControlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownershipControls",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetPublicAccessBlockConfiguration(val interface{}) {
	if err := j.validateSetPublicAccessBlockConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicAccessBlockConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetReplicationConfiguration(val interface{}) {
	if err := j.validateSetReplicationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetVersioningConfiguration(val interface{}) {
	if err := j.validateSetVersioningConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versioningConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnBucket)SetWebsiteConfiguration(val interface{}) {
	if err := j.validateSetWebsiteConfigurationParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateCfnBucket_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnBucket_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnBucket) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnBucket) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnBucket) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnBucket) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnBucket) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnBucket) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnBucket) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
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
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
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
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnBucket) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnBucket) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

