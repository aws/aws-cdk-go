package awsodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsodb"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ODB::OdbNetwork` resource creates an ODB network.
//
// An ODB network provides the networking foundation for Oracle Database resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOdbNetwork := awscdk.Aws_odb.NewCfnOdbNetwork(this, jsii.String("MyCfnOdbNetwork"), &CfnOdbNetworkProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	BackupSubnetCidr: jsii.String("backupSubnetCidr"),
//   	ClientSubnetCidr: jsii.String("clientSubnetCidr"),
//   	CustomDomainName: jsii.String("customDomainName"),
//   	DefaultDnsPrefix: jsii.String("defaultDnsPrefix"),
//   	DeleteAssociatedResources: jsii.Boolean(false),
//   	DisplayName: jsii.String("displayName"),
//   	S3Access: jsii.String("s3Access"),
//   	S3PolicyDocument: jsii.String("s3PolicyDocument"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ZeroEtlAccess: jsii.String("zeroEtlAccess"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html
//
type CfnOdbNetwork interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsodb.IOdbNetworkRef
	awscdk.ITaggableV2
	// The managed services configuration for the ODB network.
	AttrManagedServices() awscdk.IResolvable
	// The unique identifier of the OCI network anchor for the ODB network.
	AttrOciNetworkAnchorId() *string
	// The name of the OCI resource anchor that's associated with the ODB network.
	AttrOciResourceAnchorName() *string
	// The URL for the VCN that's associated with the ODB network.
	AttrOciVcnUrl() *string
	// The Amazon Resource Name (ARN) of the ODB network.
	AttrOdbNetworkArn() *string
	// The unique identifier of the ODB network.
	AttrOdbNetworkId() *string
	// The Availability Zone (AZ) where the ODB network is located.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// The AZ ID of the AZ where the ODB network is located.
	AvailabilityZoneId() *string
	SetAvailabilityZoneId(val *string)
	// The CIDR range of the backup subnet in the ODB network.
	BackupSubnetCidr() *string
	SetBackupSubnetCidr(val *string)
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The CIDR range of the client subnet in the ODB network.
	ClientSubnetCidr() *string
	SetClientSubnetCidr(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The domain name for the resources in the ODB network.
	CustomDomainName() *string
	SetCustomDomainName(val *string)
	// The DNS prefix to the default DNS domain name.
	DefaultDnsPrefix() *string
	SetDefaultDnsPrefix(val *string)
	// Specifies whether to delete associated OCI networking resources along with the ODB network.
	DeleteAssociatedResources() interface{}
	SetDeleteAssociatedResources(val interface{})
	// The user-friendly name of the ODB network.
	DisplayName() *string
	SetDisplayName(val *string)
	Env() *interfaces.ResourceEnvironment
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
	// A reference to a OdbNetwork resource.
	OdbNetworkRef() *interfacesawsodb.OdbNetworkReference
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The configuration for Amazon S3 access from the ODB network.
	S3Access() *string
	SetS3Access(val *string)
	// Specifies the endpoint policy for Amazon S3 access from the ODB network.
	S3PolicyDocument() *string
	SetS3PolicyDocument(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags to assign to the Odb Network.
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
	// The configuration for Zero-ETL access from the ODB network.
	ZeroEtlAccess() *string
	SetZeroEtlAccess(val *string)
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
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
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
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
}

// The jsii proxy struct for CfnOdbNetwork
type jsiiProxy_CfnOdbNetwork struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsodbIOdbNetworkRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnOdbNetwork) AttrManagedServices() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrManagedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) AttrOciNetworkAnchorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciNetworkAnchorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) AttrOciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) AttrOciVcnUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciVcnUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) AttrOdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOdbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) AttrOdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOdbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) BackupSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) ClientSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) CustomDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) DefaultDnsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDnsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) DeleteAssociatedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAssociatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) OdbNetworkRef() *interfacesawsodb.OdbNetworkReference {
	var returns *interfacesawsodb.OdbNetworkReference
	_jsii_.Get(
		j,
		"odbNetworkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) S3Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) S3PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOdbNetwork) ZeroEtlAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroEtlAccess",
		&returns,
	)
	return returns
}


// Create a new `AWS::ODB::OdbNetwork`.
func NewCfnOdbNetwork(scope constructs.Construct, id *string, props *CfnOdbNetworkProps) CfnOdbNetwork {
	_init_.Initialize()

	if err := validateNewCfnOdbNetworkParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOdbNetwork{}

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ODB::OdbNetwork`.
func NewCfnOdbNetwork_Override(c CfnOdbNetwork, scope constructs.Construct, id *string, props *CfnOdbNetworkProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetAvailabilityZoneId(val *string) {
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetBackupSubnetCidr(val *string) {
	_jsii_.Set(
		j,
		"backupSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetClientSubnetCidr(val *string) {
	_jsii_.Set(
		j,
		"clientSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetCustomDomainName(val *string) {
	_jsii_.Set(
		j,
		"customDomainName",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetDefaultDnsPrefix(val *string) {
	_jsii_.Set(
		j,
		"defaultDnsPrefix",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetDeleteAssociatedResources(val interface{}) {
	if err := j.validateSetDeleteAssociatedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAssociatedResources",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetS3Access(val *string) {
	_jsii_.Set(
		j,
		"s3Access",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetS3PolicyDocument(val *string) {
	_jsii_.Set(
		j,
		"s3PolicyDocument",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnOdbNetwork)SetZeroEtlAccess(val *string) {
	_jsii_.Set(
		j,
		"zeroEtlAccess",
		val,
	)
}

func CfnOdbNetwork_ArnForOdbNetwork(resource interfacesawsodb.IOdbNetworkRef) *string {
	_init_.Initialize()

	if err := validateCfnOdbNetwork_ArnForOdbNetworkParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		"arnForOdbNetwork",
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
func CfnOdbNetwork_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOdbNetwork_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnOdbNetwork.
func CfnOdbNetwork_IsCfnOdbNetwork(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOdbNetwork_IsCfnOdbNetworkParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		"isCfnOdbNetwork",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnOdbNetwork_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOdbNetwork_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		"isCfnResource",
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
func CfnOdbNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOdbNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOdbNetwork_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_odb.CfnOdbNetwork",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOdbNetwork) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnOdbNetwork) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnOdbNetwork) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOdbNetwork) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOdbNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnOdbNetwork) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnOdbNetwork) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOdbNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOdbNetwork) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

