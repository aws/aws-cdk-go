package awskinesisanalytics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisanalyticsv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::KinesisAnalyticsV2::Snapshot, a point-in-time snapshot of a Managed Service for Apache Flink (Kinesis Analytics V2) application's state.
//
// Snapshots are account-scoped and immutable: every input property except Tags is create-only, so changing any of them forces replacement, while changing Tags updates the snapshot in place. Specify exactly one source: ApplicationName to snapshot a running application (CreateSnapshot), or ImportSourceConfiguration to import a snapshot from Amazon S3 (ImportSnapshot). CloudFormation deletes the snapshot via DeleteSnapshot; to retain a snapshot beyond the stack, set DeletionPolicy: Retain on the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotV2 := awscdk.Aws_kinesisanalytics.NewCfnSnapshotV2(this, jsii.String("MyCfnSnapshotV2"), &CfnSnapshotV2Props{
//   	SnapshotName: jsii.String("snapshotName"),
//
//   	// the properties below are optional
//   	ApplicationName: jsii.String("applicationName"),
//   	ImportSourceConfiguration: &ImportSourceConfigurationProperty{
//   		ImportRoleArn: jsii.String("importRoleArn"),
//   		SourceS3PrefixLocation: &S3PrefixLocationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   		SourceType: jsii.String("sourceType"),
//
//   		// the properties below are optional
//   		SourceEncryptionConfiguration: &SnapshotEncryptionConfigurationDescriptionProperty{
//   			KeyType: jsii.String("keyType"),
//
//   			// the properties below are optional
//   			KeyId: jsii.String("keyId"),
//   		},
//   	},
//   	SnapshotDescription: jsii.String("snapshotDescription"),
//   	SnapshotEncryptionConfigurationDescription: &SnapshotEncryptionConfigurationDescriptionProperty{
//   		KeyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		KeyId: jsii.String("keyId"),
//   	},
//   	SnapshotFormat: jsii.String("snapshotFormat"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-snapshot.html
//
// Deprecated: use `aws-kinesisanalyticsv2` instead.
type CfnSnapshotV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawskinesisanalyticsv2.ISnapshotRef
	awscdk.ITaggableV2
	// The name of the Managed Service for Apache Flink application to snapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplicationName() *string
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SetApplicationName(val *string)
	// The runtime environment of the source application when the snapshot was taken with ApplicationName (CreateSnapshot), for example FLINK-1_20.
	//
	// Populated by DescribeSnapshot. Not returned for snapshots created with ImportSourceConfiguration (ImportSnapshot), which carry no runtime version.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrRuntimeEnvironment() *string
	// The Amazon Resource Name (ARN) of the snapshot.
	//
	// Account-scoped: arn:<partition>:kinesisanalytics:<region>:<account>:snapshot/<SnapshotName>. Assigned by the service at create time and used as the resource's primary identifier.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotArn() *string
	// The time at which the service created the snapshot, in ISO-8601 format.
	//
	// Populated by DescribeSnapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotCreationTimestamp() *string
	// Indicates whether the snapshot was initiated by the customer or by the service.
	//
	// Populated by DescribeSnapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotInitiator() *string
	// The operation that initiated the snapshot.
	//
	// Populated by DescribeSnapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotInitiatorOperation() *string
	// The size of the snapshot in bytes.
	//
	// Populated by DescribeSnapshot once the snapshot reaches READY.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotSizeInBytes() *float64
	// Details about the source (application and/or S3 location) the snapshot was taken from.
	//
	// Populated by DescribeSnapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotSourceDescription() awscdk.IResolvable
	// The ARN of the S3 bucket.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotSourceDescriptionS3PrefixLocationBucketArn() *string
	// The S3 key prefix.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotSourceDescriptionS3PrefixLocationKeyPrefix() *string
	// The time at which the source application was created.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationCreationTimestamp() *string
	// The name of the source application.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationName() *string
	// The version of the source application captured in the snapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationVersionId() *float64
	// The lifecycle status of the snapshot as reported by the service (CREATING, READY, DELETING, FAILED).
	//
	// Populated by DescribeSnapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AttrSnapshotStatus() *string
	// Tag Manager which manages the tags for this resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CfnOptions() awscdk.ICfnResourceOptions
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CfnProperties() *map[string]interface{}
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CreationStack() *[]*string
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	Env() *interfaces.ResourceEnvironment
	// The Amazon S3 location to import a snapshot from, and the role the service assumes to read it.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ImportSourceConfiguration() interface{}
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SetImportSourceConfiguration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	LogicalId() *string
	// The tree node.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	Ref() *string
	// Optional customer-supplied description of the snapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SnapshotDescription() *string
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SetSnapshotDescription(val *string)
	// A KMS encryption configuration for snapshot data: the key type, and the key when it is a customer managed key.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SnapshotEncryptionConfigurationDescription() interface{}
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SetSnapshotEncryptionConfigurationDescription(val interface{})
	// The format of the snapshot (CANONICAL).
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SnapshotFormat() *string
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SetSnapshotFormat(val *string)
	// The customer-chosen name of the snapshot, unique within the account.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SnapshotName() *string
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SetSnapshotName(val *string)
	// A reference to a Snapshot resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	SnapshotRef() *interfacesawskinesisanalyticsv2.SnapshotReference
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	Stack() awscdk.Stack
	// Tags to apply to the snapshot.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	Tags() *[]*awscdk.CfnTag
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
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
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
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
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
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
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
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
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	CfnPropertyName(cdkPropertyName *string) *string
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
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
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	RemoveResourceDependency(target awscdk.CfnResource)
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ToString() *string
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	ValidateProperties(_properties interface{})
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	// Deprecated: use `aws-kinesisanalyticsv2` instead.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnSnapshotV2
type jsiiProxy_CfnSnapshotV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawskinesisanalyticsv2ISnapshotRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnSnapshotV2) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrRuntimeEnvironment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRuntimeEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotCreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotCreationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotInitiator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotInitiator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotInitiatorOperation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotInitiatorOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrSnapshotSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotSourceDescription() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrSnapshotSourceDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotSourceDescriptionS3PrefixLocationBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotSourceDescriptionS3PrefixLocationBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotSourceDescriptionS3PrefixLocationKeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotSourceDescriptionS3PrefixLocationKeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationCreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationCreationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationVersionId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrSnapshotSourceDescriptionSourceApplicationDescriptionApplicationVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) AttrSnapshotStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSnapshotStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) ImportSourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) SnapshotDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) SnapshotEncryptionConfigurationDescription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotEncryptionConfigurationDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) SnapshotFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) SnapshotRef() *interfacesawskinesisanalyticsv2.SnapshotReference {
	var returns *interfacesawskinesisanalyticsv2.SnapshotReference
	_jsii_.Get(
		j,
		"snapshotRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSnapshotV2) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisAnalyticsV2::Snapshot`.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
func NewCfnSnapshotV2(scope constructs.Construct, id *string, props *CfnSnapshotV2Props) CfnSnapshotV2 {
	_init_.Initialize()

	if err := validateNewCfnSnapshotV2Parameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSnapshotV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisAnalyticsV2::Snapshot`.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
func NewCfnSnapshotV2_Override(c CfnSnapshotV2, scope constructs.Construct, id *string, props *CfnSnapshotV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSnapshotV2)SetApplicationName(val *string) {
	_jsii_.Set(
		j,
		"applicationName",
		val,
	)
}

func (j *jsiiProxy_CfnSnapshotV2)SetImportSourceConfiguration(val interface{}) {
	if err := j.validateSetImportSourceConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnSnapshotV2)SetSnapshotDescription(val *string) {
	_jsii_.Set(
		j,
		"snapshotDescription",
		val,
	)
}

func (j *jsiiProxy_CfnSnapshotV2)SetSnapshotEncryptionConfigurationDescription(val interface{}) {
	if err := j.validateSetSnapshotEncryptionConfigurationDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotEncryptionConfigurationDescription",
		val,
	)
}

func (j *jsiiProxy_CfnSnapshotV2)SetSnapshotFormat(val *string) {
	_jsii_.Set(
		j,
		"snapshotFormat",
		val,
	)
}

func (j *jsiiProxy_CfnSnapshotV2)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_CfnSnapshotV2)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Deprecated: use `aws-kinesisanalyticsv2` instead.
func CfnSnapshotV2_ArnForSnapshot(resource interfacesawskinesisanalyticsv2.ISnapshotRef) *string {
	_init_.Initialize()

	if err := validateCfnSnapshotV2_ArnForSnapshotParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		"arnForSnapshot",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
func CfnSnapshotV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSnapshotV2_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
func CfnSnapshotV2_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSnapshotV2_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnSnapshotV2.
// Deprecated: use `aws-kinesisanalyticsv2` instead.
func CfnSnapshotV2_IsCfnSnapshotV2(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSnapshotV2_IsCfnSnapshotV2Parameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		"isCfnSnapshotV2",
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
// Deprecated: use `aws-kinesisanalyticsv2` instead.
func CfnSnapshotV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSnapshotV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSnapshotV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisanalytics.CfnSnapshotV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSnapshotV2) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) CfnPropertyName(cdkPropertyName *string) *string {
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

func (c *jsiiProxy_CfnSnapshotV2) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnSnapshotV2) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnSnapshotV2) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshotV2) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnSnapshotV2) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshotV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSnapshotV2) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnSnapshotV2) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

