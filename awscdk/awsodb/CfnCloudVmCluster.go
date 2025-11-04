package awsodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ODB::CloudVmCluster` resource creates a VM cluster on the specified Exadata infrastructure in the Oracle Database.
//
// A VM cluster provides the compute resources for Oracle Database workloads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudVmCluster := awscdk.Aws_odb.NewCfnCloudVmCluster(this, jsii.String("MyCfnCloudVmCluster"), &CfnCloudVmClusterProps{
//   	CloudExadataInfrastructureId: jsii.String("cloudExadataInfrastructureId"),
//   	ClusterName: jsii.String("clusterName"),
//   	CpuCoreCount: jsii.Number(123),
//   	DataCollectionOptions: &DataCollectionOptionsProperty{
//   		IsDiagnosticsEventsEnabled: jsii.Boolean(false),
//   		IsHealthMonitoringEnabled: jsii.Boolean(false),
//   		IsIncidentLogsEnabled: jsii.Boolean(false),
//   	},
//   	DataStorageSizeInTBs: jsii.Number(123),
//   	DbNodes: []interface{}{
//   		&DbNodeProperty{
//   			DbServerId: jsii.String("dbServerId"),
//
//   			// the properties below are optional
//   			BackupIpId: jsii.String("backupIpId"),
//   			BackupVnic2Id: jsii.String("backupVnic2Id"),
//   			CpuCoreCount: jsii.Number(123),
//   			DbNodeArn: jsii.String("dbNodeArn"),
//   			DbNodeId: jsii.String("dbNodeId"),
//   			DbNodeStorageSizeInGBs: jsii.Number(123),
//   			DbSystemId: jsii.String("dbSystemId"),
//   			HostIpId: jsii.String("hostIpId"),
//   			Hostname: jsii.String("hostname"),
//   			MemorySizeInGBs: jsii.Number(123),
//   			Ocid: jsii.String("ocid"),
//   			Status: jsii.String("status"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Vnic2Id: jsii.String("vnic2Id"),
//   			VnicId: jsii.String("vnicId"),
//   		},
//   	},
//   	DbNodeStorageSizeInGBs: jsii.Number(123),
//   	DbServers: []*string{
//   		jsii.String("dbServers"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	GiVersion: jsii.String("giVersion"),
//   	Hostname: jsii.String("hostname"),
//   	IsLocalBackupEnabled: jsii.Boolean(false),
//   	IsSparseDiskgroupEnabled: jsii.Boolean(false),
//   	LicenseModel: jsii.String("licenseModel"),
//   	MemorySizeInGBs: jsii.Number(123),
//   	OdbNetworkId: jsii.String("odbNetworkId"),
//   	ScanListenerPortTcp: jsii.Number(123),
//   	SshPublicKeys: []*string{
//   		jsii.String("sshPublicKeys"),
//   	},
//   	SystemVersion: jsii.String("systemVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudvmcluster.html
//
type CfnCloudVmCluster interface {
	awscdk.CfnResource
	ICloudVmClusterRef
	awscdk.IInspectable
	awscdk.ITaggableV2
	// The Amazon Resource Name (ARN) of the VM cluster.
	AttrCloudVmClusterArn() *string
	// The unique identifier of the VM cluster.
	AttrCloudVmClusterId() *string
	// The OCI model compute model used when you create or clone an instance: ECPU or OCPU.
	//
	// An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled.
	AttrComputeModel() *string
	// The type of redundancy configured for the VM cluster.
	//
	// `NORMAL` is 2-way redundancy. `HIGH` is 3-way redundancy.
	AttrDiskRedundancy() *string
	// The domain of the VM cluster.
	AttrDomain() *string
	// The port number configured for the listener on the VM cluster.
	AttrListenerPort() *float64
	// The number of nodes in the VM cluster.
	AttrNodeCount() *float64
	// The OCID of the VM cluster.
	AttrOcid() *string
	// The name of the OCI resource anchor for the VM cluster.
	AttrOciResourceAnchorName() *string
	// The HTTPS link to the VM cluster in OCI.
	AttrOciUrl() *string
	// The FQDN of the DNS record for the Single Client Access Name (SCAN) IP addresses that are associated with the VM cluster.
	AttrScanDnsName() *string
	// The OCID of the SCAN IP addresses that are associated with the VM cluster.
	AttrScanIpIds() *[]*string
	// The hardware model name of the Exadata infrastructure that's running the VM cluster.
	AttrShape() *string
	// The amount of local node storage, in gigabytes (GB), that's allocated to the VM cluster.
	AttrStorageSizeInGBs() *float64
	// The virtual IP (VIP) addresses that are associated with the VM cluster.
	//
	// Oracle's Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the VM cluster to enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.
	AttrVipIds() *[]*string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The unique identifier of the Exadata infrastructure that this VM cluster belongs to.
	CloudExadataInfrastructureId() *string
	SetCloudExadataInfrastructureId(val *string)
	// A reference to a CloudVmCluster resource.
	CloudVmClusterRef() *CloudVmClusterReference
	// The name of the Grid Infrastructure (GI) cluster.
	ClusterName() *string
	SetClusterName(val *string)
	// The number of CPU cores enabled on the VM cluster.
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The set of diagnostic collection options enabled for the VM cluster.
	DataCollectionOptions() interface{}
	SetDataCollectionOptions(val interface{})
	// The size of the data disk group, in terabytes (TB), that's allocated for the VM cluster.
	DataStorageSizeInTBs() *float64
	SetDataStorageSizeInTBs(val *float64)
	// The DB nodes that are implicitly created and managed as part of this VM Cluster.
	DbNodes() interface{}
	SetDbNodes(val interface{})
	// The amount of local node storage, in gigabytes (GB), that's allocated for the VM cluster.
	DbNodeStorageSizeInGBs() *float64
	SetDbNodeStorageSizeInGBs(val *float64)
	// The list of database servers for the VM cluster.
	DbServers() *[]*string
	SetDbServers(val *[]*string)
	// The user-friendly name for the VM cluster.
	DisplayName() *string
	SetDisplayName(val *string)
	Env() *awscdk.ResourceEnvironment
	// The software version of the Oracle Grid Infrastructure (GI) for the VM cluster.
	GiVersion() *string
	SetGiVersion(val *string)
	// The host name for the VM cluster.
	Hostname() *string
	SetHostname(val *string)
	// Specifies whether database backups to local Exadata storage are enabled for the VM cluster.
	IsLocalBackupEnabled() interface{}
	SetIsLocalBackupEnabled(val interface{})
	// Specifies whether the VM cluster is configured with a sparse disk group.
	IsSparseDiskgroupEnabled() interface{}
	SetIsSparseDiskgroupEnabled(val interface{})
	// The Oracle license model applied to the VM cluster.
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
	// The amount of memory, in gigabytes (GB), that's allocated for the VM cluster.
	MemorySizeInGBs() *float64
	SetMemorySizeInGBs(val *float64)
	// The tree node.
	Node() constructs.Node
	// The unique identifier of the ODB network for the VM cluster.
	OdbNetworkId() *string
	SetOdbNetworkId(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The port number for TCP connections to the single client access name (SCAN) listener.
	ScanListenerPortTcp() *float64
	SetScanListenerPortTcp(val *float64)
	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The operating system version of the image chosen for the VM cluster.
	SystemVersion() *string
	SetSystemVersion(val *string)
	// Tags to assign to the Vm Cluster.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The time zone of the VM cluster.
	TimeZone() *string
	SetTimeZone(val *string)
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

// The jsii proxy struct for CfnCloudVmCluster
type jsiiProxy_CfnCloudVmCluster struct {
	internal.Type__awscdkCfnResource
	jsiiProxy_ICloudVmClusterRef
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrCloudVmClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloudVmClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrCloudVmClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCloudVmClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComputeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrDiskRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDiskRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrListenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrOciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrOciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrOciUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrScanDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrScanDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrScanIpIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrScanIpIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrShape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrShape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) AttrVipIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrVipIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CloudVmClusterRef() *CloudVmClusterReference {
	var returns *CloudVmClusterReference
	_jsii_.Get(
		j,
		"cloudVmClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) DataCollectionOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) DataStorageSizeInTBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) DbNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dbNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) DbNodeStorageSizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) IsLocalBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLocalBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) IsSparseDiskgroupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSparseDiskgroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) MemorySizeInGBs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGBs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) SystemVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudVmCluster) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnCloudVmCluster(scope constructs.Construct, id *string, props *CfnCloudVmClusterProps) CfnCloudVmCluster {
	_init_.Initialize()

	if err := validateNewCfnCloudVmClusterParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudVmCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnCloudVmCluster_Override(c CfnCloudVmCluster, scope constructs.Construct, id *string, props *CfnCloudVmClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetCloudExadataInfrastructureId(val *string) {
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetClusterName(val *string) {
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetCpuCoreCount(val *float64) {
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetDataCollectionOptions(val interface{}) {
	if err := j.validateSetDataCollectionOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCollectionOptions",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetDataStorageSizeInTBs(val *float64) {
	_jsii_.Set(
		j,
		"dataStorageSizeInTBs",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetDbNodes(val interface{}) {
	if err := j.validateSetDbNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodes",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetDbNodeStorageSizeInGBs(val *float64) {
	_jsii_.Set(
		j,
		"dbNodeStorageSizeInGBs",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetDbServers(val *[]*string) {
	_jsii_.Set(
		j,
		"dbServers",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetDisplayName(val *string) {
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetGiVersion(val *string) {
	_jsii_.Set(
		j,
		"giVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetIsLocalBackupEnabled(val interface{}) {
	if err := j.validateSetIsLocalBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLocalBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetIsSparseDiskgroupEnabled(val interface{}) {
	if err := j.validateSetIsSparseDiskgroupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSparseDiskgroupEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetLicenseModel(val *string) {
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetMemorySizeInGBs(val *float64) {
	_jsii_.Set(
		j,
		"memorySizeInGBs",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetOdbNetworkId(val *string) {
	_jsii_.Set(
		j,
		"odbNetworkId",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetScanListenerPortTcp(val *float64) {
	_jsii_.Set(
		j,
		"scanListenerPortTcp",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetSshPublicKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetSystemVersion(val *string) {
	_jsii_.Set(
		j,
		"systemVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnCloudVmCluster)SetTimeZone(val *string) {
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCloudVmCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudVmCluster_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnCloudVmCluster_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudVmCluster_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster",
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
func CfnCloudVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudVmCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_odb.CfnCloudVmCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudVmCluster) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCloudVmCluster) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCloudVmCluster) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudVmCluster) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCloudVmCluster) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCloudVmCluster) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCloudVmCluster) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

