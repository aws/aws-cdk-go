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

// The `AWS::ODB::CloudAutonomousVmCluster` resource creates an Autonomous VM cluster.
//
// An Autonomous VM cluster provides the infrastructure for running Autonomous Databases.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudAutonomousVmCluster := awscdk.Aws_odb.NewCfnCloudAutonomousVmCluster(this, jsii.String("MyCfnCloudAutonomousVmCluster"), &CfnCloudAutonomousVmClusterProps{
//   	AutonomousDataStorageSizeInTBs: jsii.Number(123),
//   	CloudExadataInfrastructureId: jsii.String("cloudExadataInfrastructureId"),
//   	CpuCoreCountPerNode: jsii.Number(123),
//   	DbServers: []*string{
//   		jsii.String("dbServers"),
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	IsMtlsEnabledVmCluster: jsii.Boolean(false),
//   	LicenseModel: jsii.String("licenseModel"),
//   	MaintenanceWindow: &MaintenanceWindowProperty{
//   		DaysOfWeek: []*string{
//   			jsii.String("daysOfWeek"),
//   		},
//   		HoursOfDay: []interface{}{
//   			jsii.Number(123),
//   		},
//   		LeadTimeInWeeks: jsii.Number(123),
//   		Months: []*string{
//   			jsii.String("months"),
//   		},
//   		Preference: jsii.String("preference"),
//   		WeeksOfMonth: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	MemoryPerOracleComputeUnitInGBs: jsii.Number(123),
//   	OdbNetworkId: jsii.String("odbNetworkId"),
//   	ScanListenerPortNonTls: jsii.Number(123),
//   	ScanListenerPortTls: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   	TotalContainerDatabases: jsii.Number(123),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html
//
type CfnCloudAutonomousVmCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsodb.ICloudAutonomousVmClusterRef
	awscdk.ITaggableV2
	// The percentage of data storage currently in use for Autonomous Databases in the Autonomous VM cluster.
	AttrAutonomousDataStoragePercentage() awscdk.IResolvable
	// The available data storage space for Autonomous Databases in the Autonomous VM cluster, in TB.
	AttrAvailableAutonomousDataStorageSizeInTBs() awscdk.IResolvable
	// The number of Autonomous CDBs that you can create with the currently available storage.
	AttrAvailableContainerDatabases() *float64
	// The number of CPU cores available for allocation to Autonomous Databases.
	AttrAvailableCpus() awscdk.IResolvable
	// The Amazon Resource Name (ARN) for the Autonomous VM cluster.
	AttrCloudAutonomousVmClusterArn() *string
	// The unique identifier of the Autonomous VM cluster.
	AttrCloudAutonomousVmClusterId() *string
	// The compute model of the Autonomous VM cluster: ECPU or OCPU.
	AttrComputeModel() *string
	// The total number of CPU cores in the Autonomous VM cluster.
	AttrCpuCoreCount() *float64
	// The percentage of total CPU cores currently in use in the Autonomous VM cluster.
	AttrCpuPercentage() awscdk.IResolvable
	// The total data storage allocated to the Autonomous VM cluster, in GB.
	AttrDataStorageSizeInGBs() awscdk.IResolvable
	// The total data storage allocated to the Autonomous VM cluster, in TB.
	AttrDataStorageSizeInTBs() awscdk.IResolvable
	// The local node storage allocated to the Autonomous VM cluster, in gigabytes (GB).
	AttrDbNodeStorageSizeInGBs() *float64
	// The domain name for the Autonomous VM cluster.
	AttrDomain() *string
	// The minimum value to which you can scale down the Exadata storage, in TB.
	AttrExadataStorageInTBsLowestScaledValue() awscdk.IResolvable
	// The hostname for the Autonomous VM cluster.
	AttrHostname() *string
	// The minimum value to which you can scale down the maximum number of Autonomous CDBs.
	AttrMaxAcdsLowestScaledValue() *float64
	// The total amount of memory allocated to the Autonomous VM cluster, in gigabytes (GB).
	AttrMemorySizeInGBs() *float64
	// The number of database server nodes in the Autonomous VM cluster.
	AttrNodeCount() *float64
	// The number of Autonomous CDBs that can't be provisioned because of resource constraints.
	AttrNonProvisionableAutonomousContainerDatabases() *float64
	// The Oracle Cloud Identifier (OCID) of the Autonomous VM cluster.
	AttrOcid() *string
	// The name of the OCI resource anchor associated with this Autonomous VM cluster.
	AttrOciResourceAnchorName() *string
	// The URL for accessing the OCI console page for this Autonomous VM cluster.
	AttrOciUrl() *string
	// The number of Autonomous CDBs that can be provisioned in the Autonomous VM cluster.
	AttrProvisionableAutonomousContainerDatabases() *float64
	// The number of Autonomous CDBs currently provisioned in the Autonomous VM cluster.
	AttrProvisionedAutonomousContainerDatabases() *float64
	// The number of CPU cores currently provisioned in the Autonomous VM cluster.
	AttrProvisionedCpus() awscdk.IResolvable
	// The number of CPU cores that can be reclaimed from terminated or scaled-down Autonomous Databases.
	AttrReclaimableCpus() awscdk.IResolvable
	// The number of CPU cores reserved for system operations and redundancy.
	AttrReservedCpus() awscdk.IResolvable
	// The shape of the Exadata infrastructure for the Autonomous VM cluster.
	AttrShape() *string
	// The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB.
	AutonomousDataStorageSizeInTBs() *float64
	SetAutonomousDataStorageSizeInTBs(val *float64)
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A reference to a CloudAutonomousVmCluster resource.
	CloudAutonomousVmClusterRef() *interfacesawsodb.CloudAutonomousVmClusterReference
	// The unique identifier of the Cloud Exadata Infrastructure containing this Autonomous VM cluster.
	CloudExadataInfrastructureId() *string
	SetCloudExadataInfrastructureId(val *string)
	// The number of CPU cores enabled per node in the Autonomous VM cluster.
	CpuCoreCountPerNode() *float64
	SetCpuCoreCountPerNode(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The list of database servers associated with the Autonomous VM cluster.
	DbServers() *[]*string
	SetDbServers(val *[]*string)
	// The user-provided description of the Autonomous VM cluster.
	Description() *string
	SetDescription(val *string)
	// The display name of the Autonomous VM cluster.
	DisplayName() *string
	SetDisplayName(val *string)
	Env() *interfaces.ResourceEnvironment
	// Specifies whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster.
	IsMtlsEnabledVmCluster() interface{}
	SetIsMtlsEnabledVmCluster(val interface{})
	// The Oracle license model that applies to the Autonomous VM cluster.
	LicenseModel() *string
	SetLicenseModel(val *string)
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
	// The scheduling details for the maintenance window.
	MaintenanceWindow() interface{}
	SetMaintenanceWindow(val interface{})
	// The amount of memory allocated per Oracle Compute Unit, in GB.
	MemoryPerOracleComputeUnitInGBs() *float64
	SetMemoryPerOracleComputeUnitInGBs(val *float64)
	// The tree node.
	Node() constructs.Node
	// The unique identifier of the ODB network associated with this Autonomous VM cluster.
	OdbNetworkId() *string
	SetOdbNetworkId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The SCAN listener port for non-TLS (TCP) protocol.
	ScanListenerPortNonTls() *float64
	SetScanListenerPortNonTls(val *float64)
	// The SCAN listener port for TLS (TCP) protocol.
	ScanListenerPortTls() *float64
	SetScanListenerPortTls(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags to assign to the Autonomous Vm Cluster.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The time zone of the Autonomous VM cluster.
	TimeZone() *string
	SetTimeZone(val *string)
	// The total number of Autonomous Container Databases that can be created with the allocated local storage.
	TotalContainerDatabases() *float64
	SetTotalContainerDatabases(val *float64)
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

// The jsii proxy struct for CfnCloudAutonomousVmCluster
type jsiiProxy_CfnCloudAutonomousVmCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsodbICloudAutonomousVmClusterRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrAutonomousDataStoragePercentage() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAutonomousDataStoragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrAvailableAutonomousDataStorageSizeInTBs() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAvailableAutonomousDataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrAvailableContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrAvailableContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrAvailableCpus() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAvailableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrCloudAutonomousVmClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloudAutonomousVmClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrCloudAutonomousVmClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloudAutonomousVmClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComputeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrCpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrCpuPercentage() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrCpuPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrDataStorageSizeInGBs() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrDataStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrDataStorageSizeInTBs() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrDataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrDbNodeStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrDbNodeStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrExadataStorageInTBsLowestScaledValue() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrExadataStorageInTBsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrMaxAcdsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrMaxAcdsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrMemorySizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrMemorySizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrNonProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrNonProvisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrOciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrOciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrProvisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrProvisionedAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrProvisionedAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrProvisionedCpus() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrProvisionedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrReclaimableCpus() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrReclaimableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrReservedCpus() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrReservedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AttrShape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrShape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) AutonomousDataStorageSizeInTBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CloudAutonomousVmClusterRef() *interfacesawsodb.CloudAutonomousVmClusterReference {
	var returns *interfacesawsodb.CloudAutonomousVmClusterReference
	_jsii_.Get(
		j,
		"cloudAutonomousVmClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CpuCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) IsMtlsEnabledVmCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMtlsEnabledVmCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) MaintenanceWindow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) MemoryPerOracleComputeUnitInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) ScanListenerPortNonTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortNonTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) ScanListenerPortTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) TotalContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ODB::CloudAutonomousVmCluster`.
func NewCfnCloudAutonomousVmCluster(scope constructs.Construct, id *string, props *CfnCloudAutonomousVmClusterProps) CfnCloudAutonomousVmCluster {
	_init_.Initialize()

	if err := validateNewCfnCloudAutonomousVmClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudAutonomousVmCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ODB::CloudAutonomousVmCluster`.
func NewCfnCloudAutonomousVmCluster_Override(c CfnCloudAutonomousVmCluster, scope constructs.Construct, id *string, props *CfnCloudAutonomousVmClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetAutonomousDataStorageSizeInTBs(val *float64) {
	_jsii_.Set(
		j,
		"autonomousDataStorageSizeInTBs",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetCloudExadataInfrastructureId(val *string) {
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetCpuCoreCountPerNode(val *float64) {
	_jsii_.Set(
		j,
		"cpuCoreCountPerNode",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetDbServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dbServers",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetIsMtlsEnabledVmCluster(val interface{}) {
	if err := j.validateSetIsMtlsEnabledVmClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMtlsEnabledVmCluster",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetLicenseModel(val *string) {
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetMaintenanceWindow(val interface{}) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetMemoryPerOracleComputeUnitInGBs(val *float64) {
	_jsii_.Set(
		j,
		"memoryPerOracleComputeUnitInGBs",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetOdbNetworkId(val *string) {
	_jsii_.Set(
		j,
		"odbNetworkId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetScanListenerPortNonTls(val *float64) {
	_jsii_.Set(
		j,
		"scanListenerPortNonTls",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetScanListenerPortTls(val *float64) {
	_jsii_.Set(
		j,
		"scanListenerPortTls",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_CfnCloudAutonomousVmCluster)SetTotalContainerDatabases(val *float64) {
	_jsii_.Set(
		j,
		"totalContainerDatabases",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCloudAutonomousVmCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudAutonomousVmCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCloudAutonomousVmCluster_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudAutonomousVmCluster_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster",
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
func CfnCloudAutonomousVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudAutonomousVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudAutonomousVmCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_odb.CfnCloudAutonomousVmCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudAutonomousVmCluster) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

