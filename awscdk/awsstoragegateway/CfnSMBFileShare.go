package awsstoragegateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstoragegateway/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstoragegateway"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::StorageGateway::SMBFileShare.
//
// An SMB file share exposes an Amazon S3 bucket through a Storage Gateway file gateway, so clients read and write objects as files over SMB. The share's variant is fixed by its resource type rather than by a property: the service constructs the two variants with different APIs, and a property carrying the choice would belong to neither API's input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSMBFileShare := awscdk.Aws_storagegateway.NewCfnSMBFileShare(this, jsii.String("MyCfnSMBFileShare"), &CfnSMBFileShareProps{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	LocationArn: jsii.String("locationArn"),
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	AccessBasedEnumeration: jsii.Boolean(false),
//   	AdminUserList: []*string{
//   		jsii.String("adminUserList"),
//   	},
//   	AuditDestinationArn: jsii.String("auditDestinationArn"),
//   	Authentication: jsii.String("authentication"),
//   	BucketRegion: jsii.String("bucketRegion"),
//   	CacheAttributes: &CacheAttributesProperty{
//   		CacheStaleTimeoutInSeconds: jsii.Number(123),
//   	},
//   	CaseSensitivity: jsii.String("caseSensitivity"),
//   	DefaultStorageClass: jsii.String("defaultStorageClass"),
//   	EncryptionType: jsii.String("encryptionType"),
//   	FileShareName: jsii.String("fileShareName"),
//   	GuessMimeTypeEnabled: jsii.Boolean(false),
//   	InvalidUserList: []*string{
//   		jsii.String("invalidUserList"),
//   	},
//   	KmsEncrypted: jsii.Boolean(false),
//   	KmsKey: jsii.String("kmsKey"),
//   	NotificationPolicy: jsii.String("notificationPolicy"),
//   	ObjectAcl: jsii.String("objectAcl"),
//   	OplocksEnabled: jsii.Boolean(false),
//   	ReadOnly: jsii.Boolean(false),
//   	RequesterPays: jsii.Boolean(false),
//   	SmbaclEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ValidUserList: []*string{
//   		jsii.String("validUserList"),
//   	},
//   	VpcEndpointDnsName: jsii.String("vpcEndpointDnsName"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-storagegateway-smbfileshare.html
//
type CfnSMBFileShare interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsstoragegateway.ISMBFileShareRef
	awscdk.ITaggableV2
	// Whether clients see only the files and folders they may access.
	AccessBasedEnumeration() interface{}
	SetAccessBasedEnumeration(val interface{})
	// Users and groups granted administrative privileges on the SMB share.
	AdminUserList() *[]*string
	SetAdminUserList(val *[]*string)
	// The ARN of the file share.
	AttrFileShareArn() *string
	// The ID of the file share.
	AttrFileShareId() *string
	// The status of the file share.
	//
	// The service reports CREATING, UPDATING, AVAILABLE, DELETING and FORCE_DELETING; it is a length-bounded string in the API model rather than an enumeration.
	AttrFileShareStatus() *string
	// The export path the share is mounted or mapped at.
	AttrPath() *string
	// The ARN of the CloudWatch log group that receives the share's audit logs.
	AuditDestinationArn() *string
	SetAuditDestinationArn(val *string)
	// The authentication the SMB share uses.
	Authentication() *string
	SetAuthentication(val *string)
	// The Region of the S3 bucket, required when the bucket is reached through a VPC endpoint.
	BucketRegion() *string
	SetBucketRegion(val *string)
	// Refresh cache information for the file share.
	CacheAttributes() interface{}
	SetCacheAttributes(val interface{})
	// Whether the SMB share is case sensitive or defers to the client.
	CaseSensitivity() *string
	SetCaseSensitivity(val *string)
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The S3 storage class for new objects.
	DefaultStorageClass() *string
	SetDefaultStorageClass(val *string)
	// The server-side encryption the gateway applies to objects it writes.
	EncryptionType() *string
	SetEncryptionType(val *string)
	Env() *interfaces.ResourceEnvironment
	// The name of the file share, which defaults to the bucket name when omitted.
	FileShareName() *string
	SetFileShareName(val *string)
	// The ARN of the S3 File Gateway that hosts the share.
	GatewayArn() *string
	SetGatewayArn(val *string)
	// Whether the gateway infers the MIME type of uploaded objects from their extension.
	GuessMimeTypeEnabled() interface{}
	SetGuessMimeTypeEnabled(val interface{})
	// Users and groups denied access to the SMB share.
	InvalidUserList() *[]*string
	SetInvalidUserList(val *[]*string)
	// Deprecated by the service in favour of EncryptionType;.
	KmsEncrypted() interface{}
	SetKmsEncrypted(val interface{})
	// The KMS key used for encryption, when EncryptionType selects KMS.
	KmsKey() *string
	SetKmsKey(val *string)
	// The ARN of the S3 bucket, or bucket prefix, the share exposes.
	LocationArn() *string
	SetLocationArn(val *string)
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
	// The upload-notification settling period, as a JSON document.
	NotificationPolicy() *string
	SetNotificationPolicy(val *string)
	// The canned S3 ACL applied to objects the gateway writes.
	ObjectAcl() *string
	SetObjectAcl(val *string)
	// Whether the SMB share grants opportunistic locks.
	OplocksEnabled() interface{}
	SetOplocksEnabled(val interface{})
	// Whether clients may write to the share.
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Whether the bucket is a Requester Pays bucket.
	RequesterPays() interface{}
	SetRequesterPays(val interface{})
	// The ARN of the IAM role the gateway assumes when it accesses the S3 location.
	Role() *string
	SetRole(val *string)
	// Whether the SMB share uses Windows ACLs rather than POSIX permissions.
	SmbaclEnabled() interface{}
	SetSmbaclEnabled(val interface{})
	// A reference to a SMBFileShare resource.
	SmbFileShareRef() *interfacesawsstoragegateway.SMBFileShareReference
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags assigned to the file share.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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
	// Users and groups permitted to access the SMB share.
	ValidUserList() *[]*string
	SetValidUserList(val *[]*string)
	// The DNS name of the S3 VPC endpoint the gateway uses to reach the bucket.
	VpcEndpointDnsName() *string
	SetVpcEndpointDnsName(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
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
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
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
	CfnPropertyName(cdkPropertyName *string) *string
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
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
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnSMBFileShare
type jsiiProxy_CfnSMBFileShare struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsstoragegatewayISMBFileShareRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnSMBFileShare) AccessBasedEnumeration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBasedEnumeration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) AdminUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) AttrFileShareArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFileShareArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) AttrFileShareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFileShareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) AttrFileShareStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFileShareStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) AttrPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) Authentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CacheAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CaseSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) GuessMimeTypeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) InvalidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) OplocksEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oplocksEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) SmbaclEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbaclEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) SmbFileShareRef() *interfacesawsstoragegateway.SMBFileShareReference {
	var returns *interfacesawsstoragegateway.SMBFileShareReference
	_jsii_.Get(
		j,
		"smbFileShareRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) ValidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSMBFileShare) VpcEndpointDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsName",
		&returns,
	)
	return returns
}


// Create a new `AWS::StorageGateway::SMBFileShare`.
func NewCfnSMBFileShare(scope constructs.Construct, id *string, props *CfnSMBFileShareProps) CfnSMBFileShare {
	_init_.Initialize()

	if err := validateNewCfnSMBFileShareParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSMBFileShare{}

	_jsii_.Create(
		"aws-cdk-lib.aws_storagegateway.CfnSMBFileShare",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::StorageGateway::SMBFileShare`.
func NewCfnSMBFileShare_Override(c CfnSMBFileShare, scope constructs.Construct, id *string, props *CfnSMBFileShareProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_storagegateway.CfnSMBFileShare",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetAccessBasedEnumeration(val interface{}) {
	if err := j.validateSetAccessBasedEnumerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessBasedEnumeration",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetAdminUserList(val *[]*string) {
	_jsii_.Set(
		j,
		"adminUserList",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetAuditDestinationArn(val *string) {
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetAuthentication(val *string) {
	_jsii_.Set(
		j,
		"authentication",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetBucketRegion(val *string) {
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetCacheAttributes(val interface{}) {
	if err := j.validateSetCacheAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetCaseSensitivity(val *string) {
	_jsii_.Set(
		j,
		"caseSensitivity",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetDefaultStorageClass(val *string) {
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetEncryptionType(val *string) {
	_jsii_.Set(
		j,
		"encryptionType",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetFileShareName(val *string) {
	_jsii_.Set(
		j,
		"fileShareName",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetGatewayArn(val *string) {
	if err := j.validateSetGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetGuessMimeTypeEnabled(val interface{}) {
	if err := j.validateSetGuessMimeTypeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guessMimeTypeEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetInvalidUserList(val *[]*string) {
	_jsii_.Set(
		j,
		"invalidUserList",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetKmsEncrypted(val interface{}) {
	if err := j.validateSetKmsEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetLocationArn(val *string) {
	if err := j.validateSetLocationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetNotificationPolicy(val *string) {
	_jsii_.Set(
		j,
		"notificationPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetObjectAcl(val *string) {
	_jsii_.Set(
		j,
		"objectAcl",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetOplocksEnabled(val interface{}) {
	if err := j.validateSetOplocksEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplocksEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetRequesterPays(val interface{}) {
	if err := j.validateSetRequesterPaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetSmbaclEnabled(val interface{}) {
	if err := j.validateSetSmbaclEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbaclEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetValidUserList(val *[]*string) {
	_jsii_.Set(
		j,
		"validUserList",
		val,
	)
}

func (j *jsiiProxy_CfnSMBFileShare)SetVpcEndpointDnsName(val *string) {
	_jsii_.Set(
		j,
		"vpcEndpointDnsName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSMBFileShare_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSMBFileShare_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_storagegateway.CfnSMBFileShare",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnSMBFileShare_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSMBFileShare_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_storagegateway.CfnSMBFileShare",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnSMBFileShare.
func CfnSMBFileShare_IsCfnSMBFileShare(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSMBFileShare_IsCfnSMBFileShareParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_storagegateway.CfnSMBFileShare",
		"isCfnSMBFileShare",
		[]interface{}{x},
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
func CfnSMBFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSMBFileShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_storagegateway.CfnSMBFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSMBFileShare_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_storagegateway.CfnSMBFileShare",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSMBFileShare) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMBFileShare) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMBFileShare) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnSMBFileShare) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMBFileShare) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnSMBFileShare) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMBFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSMBFileShare) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnSMBFileShare) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

