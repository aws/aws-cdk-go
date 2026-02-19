package awsworkspacesinstances

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsworkspacesinstances/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesinstances"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::WorkspacesInstances::WorkspaceInstance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceInstance := awscdk.Aws_workspacesinstances.NewCfnWorkspaceInstance(this, jsii.String("MyCfnWorkspaceInstance"), &CfnWorkspaceInstanceProps{
//   	ManagedInstance: &ManagedInstanceProperty{
//   		ImageId: jsii.String("imageId"),
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		BlockDeviceMappings: []interface{}{
//   			&BlockDeviceMappingProperty{
//   				DeviceName: jsii.String("deviceName"),
//   				Ebs: &EbsBlockDeviceProperty{
//   					Encrypted: jsii.Boolean(false),
//   					Iops: jsii.Number(123),
//   					KmsKeyId: jsii.String("kmsKeyId"),
//   					Throughput: jsii.Number(123),
//   					VolumeSize: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//   				},
//   				NoDevice: jsii.String("noDevice"),
//   				VirtualName: jsii.String("virtualName"),
//   			},
//   		},
//   		CapacityReservationSpecification: &CapacityReservationSpecificationProperty{
//   			CapacityReservationPreference: jsii.String("capacityReservationPreference"),
//   			CapacityReservationTarget: &CapacityReservationTargetProperty{
//   				CapacityReservationId: jsii.String("capacityReservationId"),
//   				CapacityReservationResourceGroupArn: jsii.String("capacityReservationResourceGroupArn"),
//   			},
//   		},
//   		CpuOptions: &CpuOptionsRequestProperty{
//   			CoreCount: jsii.Number(123),
//   			ThreadsPerCore: jsii.Number(123),
//   		},
//   		CreditSpecification: &CreditSpecificationRequestProperty{
//   			CpuCredits: jsii.String("cpuCredits"),
//   		},
//   		DisableApiStop: jsii.Boolean(false),
//   		EbsOptimized: jsii.Boolean(false),
//   		EnablePrimaryIpv6: jsii.Boolean(false),
//   		EnclaveOptions: &EnclaveOptionsRequestProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		HibernationOptions: &HibernationOptionsRequestProperty{
//   			Configured: jsii.Boolean(false),
//   		},
//   		IamInstanceProfile: &IamInstanceProfileSpecificationProperty{
//   			Arn: jsii.String("arn"),
//   			Name: jsii.String("name"),
//   		},
//   		InstanceMarketOptions: &InstanceMarketOptionsRequestProperty{
//   			MarketType: jsii.String("marketType"),
//   			SpotOptions: &SpotMarketOptionsProperty{
//   				InstanceInterruptionBehavior: jsii.String("instanceInterruptionBehavior"),
//   				MaxPrice: jsii.String("maxPrice"),
//   				SpotInstanceType: jsii.String("spotInstanceType"),
//   				ValidUntilUtc: jsii.String("validUntilUtc"),
//   			},
//   		},
//   		Ipv6AddressCount: jsii.Number(123),
//   		KeyName: jsii.String("keyName"),
//   		LicenseSpecifications: []interface{}{
//   			&LicenseConfigurationRequestProperty{
//   				LicenseConfigurationArn: jsii.String("licenseConfigurationArn"),
//   			},
//   		},
//   		MaintenanceOptions: &InstanceMaintenanceOptionsRequestProperty{
//   			AutoRecovery: jsii.String("autoRecovery"),
//   		},
//   		MetadataOptions: &InstanceMetadataOptionsRequestProperty{
//   			HttpEndpoint: jsii.String("httpEndpoint"),
//   			HttpProtocolIpv6: jsii.String("httpProtocolIpv6"),
//   			HttpPutResponseHopLimit: jsii.Number(123),
//   			HttpTokens: jsii.String("httpTokens"),
//   			InstanceMetadataTags: jsii.String("instanceMetadataTags"),
//   		},
//   		Monitoring: &RunInstancesMonitoringEnabledProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   		NetworkInterfaces: []interface{}{
//   			&InstanceNetworkInterfaceSpecificationProperty{
//   				Description: jsii.String("description"),
//   				DeviceIndex: jsii.Number(123),
//   				Groups: []*string{
//   					jsii.String("groups"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		NetworkPerformanceOptions: &InstanceNetworkPerformanceOptionsRequestProperty{
//   			BandwidthWeighting: jsii.String("bandwidthWeighting"),
//   		},
//   		Placement: &PlacementProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			GroupId: jsii.String("groupId"),
//   			GroupName: jsii.String("groupName"),
//   			PartitionNumber: jsii.Number(123),
//   			Tenancy: jsii.String("tenancy"),
//   		},
//   		PrivateDnsNameOptions: &PrivateDnsNameOptionsRequestProperty{
//   			EnableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   			EnableResourceNameDnsARecord: jsii.Boolean(false),
//   			HostnameType: jsii.String("hostnameType"),
//   		},
//   		SubnetId: jsii.String("subnetId"),
//   		TagSpecifications: []interface{}{
//   			&TagSpecificationProperty{
//   				ResourceType: jsii.String("resourceType"),
//   				Tags: []CfnTag{
//   					&CfnTag{
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   		UserData: jsii.String("userData"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspacesinstances-workspaceinstance.html
//
type CfnWorkspaceInstance interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsworkspacesinstances.IWorkspaceInstanceRef
	awscdk.ITaggableV2
	AttrEc2ManagedInstance() awscdk.IResolvable
	AttrEc2ManagedInstanceInstanceId() *string
	// The current state of the workspace instance.
	AttrProvisionState() *string
	// Unique identifier for the workspace instance.
	AttrWorkspaceInstanceId() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	ManagedInstance() interface{}
	SetManagedInstance(val interface{})
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
	// A reference to a WorkspaceInstance resource.
	WorkspaceInstanceRef() *interfacesawsworkspacesinstances.WorkspaceInstanceReference
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnWorkspaceInstance
type jsiiProxy_CfnWorkspaceInstance struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsworkspacesinstancesIWorkspaceInstanceRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnWorkspaceInstance) AttrEc2ManagedInstance() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEc2ManagedInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) AttrEc2ManagedInstanceInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEc2ManagedInstanceInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) AttrProvisionState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrProvisionState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) AttrWorkspaceInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWorkspaceInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) ManagedInstance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspaceInstance) WorkspaceInstanceRef() *interfacesawsworkspacesinstances.WorkspaceInstanceReference {
	var returns *interfacesawsworkspacesinstances.WorkspaceInstanceReference
	_jsii_.Get(
		j,
		"workspaceInstanceRef",
		&returns,
	)
	return returns
}


// Create a new `AWS::WorkspacesInstances::WorkspaceInstance`.
func NewCfnWorkspaceInstance(scope constructs.Construct, id *string, props *CfnWorkspaceInstanceProps) CfnWorkspaceInstance {
	_init_.Initialize()

	if err := validateNewCfnWorkspaceInstanceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspaceInstance{}

	_jsii_.Create(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WorkspacesInstances::WorkspaceInstance`.
func NewCfnWorkspaceInstance_Override(c CfnWorkspaceInstance, scope constructs.Construct, id *string, props *CfnWorkspaceInstanceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnWorkspaceInstance)SetManagedInstance(val interface{}) {
	if err := j.validateSetManagedInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedInstance",
		val,
	)
}

func (j *jsiiProxy_CfnWorkspaceInstance)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnWorkspaceInstance_ArnForWorkspaceInstance(resource interfacesawsworkspacesinstances.IWorkspaceInstanceRef) *string {
	_init_.Initialize()

	if err := validateCfnWorkspaceInstance_ArnForWorkspaceInstanceParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		"arnForWorkspaceInstance",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Creates a new IWorkspaceInstanceRef from a workspaceInstanceId.
func CfnWorkspaceInstance_FromWorkspaceInstanceId(scope constructs.Construct, id *string, workspaceInstanceId *string) interfacesawsworkspacesinstances.IWorkspaceInstanceRef {
	_init_.Initialize()

	if err := validateCfnWorkspaceInstance_FromWorkspaceInstanceIdParameters(scope, id, workspaceInstanceId); err != nil {
		panic(err)
	}
	var returns interfacesawsworkspacesinstances.IWorkspaceInstanceRef

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		"fromWorkspaceInstanceId",
		[]interface{}{scope, id, workspaceInstanceId},
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
func CfnWorkspaceInstance_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspaceInstance_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnWorkspaceInstance_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspaceInstance_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnWorkspaceInstance.
func CfnWorkspaceInstance_IsCfnWorkspaceInstance(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspaceInstance_IsCfnWorkspaceInstanceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		"isCfnWorkspaceInstance",
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
func CfnWorkspaceInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspaceInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspaceInstance_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_workspacesinstances.CfnWorkspaceInstance",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspaceInstance) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnWorkspaceInstance) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnWorkspaceInstance) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspaceInstance) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspaceInstance) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnWorkspaceInstance) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspaceInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnWorkspaceInstance) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnWorkspaceInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

