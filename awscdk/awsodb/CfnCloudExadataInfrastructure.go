package awsodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ODB::CloudExadataInfrastructure` resource creates an Exadata infrastructure.
//
// An Exadata infrastructure provides the underlying compute and storage resources for Oracle Database workloads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudExadataInfrastructure := awscdk.Aws_odb.NewCfnCloudExadataInfrastructure(this, jsii.String("MyCfnCloudExadataInfrastructure"), &CfnCloudExadataInfrastructureProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	ComputeCount: jsii.Number(123),
//   	CustomerContactsToSendToOci: []interface{}{
//   		&CustomerContactProperty{
//   			Email: jsii.String("email"),
//   		},
//   	},
//   	DatabaseServerType: jsii.String("databaseServerType"),
//   	DisplayName: jsii.String("displayName"),
//   	Shape: jsii.String("shape"),
//   	StorageCount: jsii.Number(123),
//   	StorageServerType: jsii.String("storageServerType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudexadatainfrastructure.html
//
type CfnCloudExadataInfrastructure interface {
	awscdk.CfnResource
	ICloudExadataInfrastructureRef
	awscdk.IInspectable
	awscdk.ITaggableV2
	// The number of storage servers requested for the Exadata infrastructure.
	AttrActivatedStorageCount() *float64
	// The number of storage servers requested for the Exadata infrastructure.
	AttrAdditionalStorageCount() *float64
	// The amount of available storage, in gigabytes (GB), for the Exadata infrastructure.
	AttrAvailableStorageSizeInGBs() *float64
	// The Amazon Resource Name (ARN) for the Exadata infrastructure.
	AttrCloudExadataInfrastructureArn() *string
	// The unique identifier for the Exadata infrastructure.
	AttrCloudExadataInfrastructureId() *string
	// The OCI model compute model used when you create or clone an instance: ECPU or OCPU.
	//
	// An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled.
	AttrComputeModel() *string
	// The total number of CPU cores that are allocated to the Exadata infrastructure.
	AttrCpuCount() *float64
	// The size of the Exadata infrastructure's data disk group, in terabytes (TB).
	AttrDataStorageSizeInTBs() awscdk.IResolvable
	// The size of the Exadata infrastructure's local node storage, in gigabytes (GB).
	AttrDbNodeStorageSizeInGBs() *float64
	// The list of database server identifiers for the Exadata infrastructure.
	AttrDbServerIds() *[]*string
	// The software version of the database servers (dom0) in the Exadata infrastructure.
	AttrDbServerVersion() *string
	// The total number of CPU cores available on the Exadata infrastructure.
	AttrMaxCpuCount() *float64
	// The total amount of data disk group storage, in terabytes (TB), that's available on the Exadata infrastructure.
	AttrMaxDataStorageInTBs() awscdk.IResolvable
	// The total amount of local node storage, in gigabytes (GB), that's available on the Exadata infrastructure.
	AttrMaxDbNodeStorageSizeInGBs() *float64
	// The total amount of memory, in gigabytes (GB), that's available on the Exadata infrastructure.
	AttrMaxMemoryInGBs() *float64
	// The amount of memory, in gigabytes (GB), that's allocated on the Exadata infrastructure.
	AttrMemorySizeInGBs() *float64
	// The OCID of the Exadata infrastructure.
	AttrOcid() *string
	// The name of the OCI resource anchor for the Exadata infrastructure.
	AttrOciResourceAnchorName() *string
	// The HTTPS link to the Exadata infrastructure in OCI.
	AttrOciUrl() *string
	// The software version of the storage servers on the Exadata infrastructure.
	AttrStorageServerVersion() *string
	// The total amount of storage, in gigabytes (GB), on the the Exadata infrastructure.
	AttrTotalStorageSizeInGBs() *float64
	// The name of the Availability Zone (AZ) where the Exadata infrastructure is located.
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	// The AZ ID of the AZ where the Exadata infrastructure is located.
	AvailabilityZoneId() *string
	SetAvailabilityZoneId(val *string)
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A reference to a CloudExadataInfrastructure resource.
	CloudExadataInfrastructureRef() *CloudExadataInfrastructureReference
	// The number of database servers for the Exadata infrastructure.
	ComputeCount() *float64
	SetComputeCount(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.
	CustomerContactsToSendToOci() interface{}
	SetCustomerContactsToSendToOci(val interface{})
	// The database server model type of the Exadata infrastructure.
	DatabaseServerType() *string
	SetDatabaseServerType(val *string)
	// The user-friendly name for the Exadata infrastructure.
	DisplayName() *string
	SetDisplayName(val *string)
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
	// The model name of the Exadata infrastructure.
	Shape() *string
	SetShape(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The number of storage servers that are activated for the Exadata infrastructure.
	StorageCount() *float64
	SetStorageCount(val *float64)
	// The storage server model type of the Exadata infrastructure.
	StorageServerType() *string
	SetStorageServerType(val *string)
	// Tags to assign to the Exadata Infrastructure.
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

// The jsii proxy struct for CfnCloudExadataInfrastructure
type jsiiProxy_CfnCloudExadataInfrastructure struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_ICloudExadataInfrastructureRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrActivatedStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrActivatedStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrAdditionalStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAdditionalStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrAvailableStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAvailableStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrCloudExadataInfrastructureArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloudExadataInfrastructureArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrCloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComputeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrDataStorageSizeInTBs() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrDataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrDbNodeStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrDbNodeStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrDbServerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDbServerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrDbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrMaxCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrMaxCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrMaxDataStorageInTBs() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrMaxDataStorageInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrMaxDbNodeStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrMaxDbNodeStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrMaxMemoryInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrMaxMemoryInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrMemorySizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrMemorySizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrOciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrOciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrStorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStorageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AttrTotalStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrTotalStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) CloudExadataInfrastructureRef() *CloudExadataInfrastructureReference {
	var returns *CloudExadataInfrastructureReference
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) CustomerContactsToSendToOci() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsToSendToOci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) DatabaseServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseServerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) StorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) StorageServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnCloudExadataInfrastructure(scope constructs.Construct, id *string, props *CfnCloudExadataInfrastructureProps) CfnCloudExadataInfrastructure {
	_init_.Initialize()

	if err := validateNewCfnCloudExadataInfrastructureParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudExadataInfrastructure{}

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnCloudExadataInfrastructure_Override(c CfnCloudExadataInfrastructure, scope constructs.Construct, id *string, props *CfnCloudExadataInfrastructureProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetAvailabilityZone(val *string) {
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetAvailabilityZoneId(val *string) {
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetComputeCount(val *float64) {
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetCustomerContactsToSendToOci(val interface{}) {
	if err := j.validateSetCustomerContactsToSendToOciParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerContactsToSendToOci",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetDatabaseServerType(val *string) {
	_jsii_.Set(
		j,
		"databaseServerType",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetShape(val *string) {
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetStorageCount(val *float64) {
	_jsii_.Set(
		j,
		"storageCount",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetStorageServerType(val *string) {
	_jsii_.Set(
		j,
		"storageServerType",
		val,
	)
}

func (j *jsiiProxy_CfnCloudExadataInfrastructure)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCloudExadataInfrastructure_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudExadataInfrastructure_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCloudExadataInfrastructure_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudExadataInfrastructure_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure",
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
func CfnCloudExadataInfrastructure_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudExadataInfrastructure_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudExadataInfrastructure_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_odb.CfnCloudExadataInfrastructure",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCloudExadataInfrastructure) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCloudExadataInfrastructure) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCloudExadataInfrastructure) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudExadataInfrastructure) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

