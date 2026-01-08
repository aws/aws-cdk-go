package awsgroundstation

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgroundstation/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgroundstation"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a `DataflowEndpoint` group containing the specified list of Ground Station Agent based endpoints.
//
// The `name` field in each endpoint is used in your mission profile `DataflowEndpointConfig` to specify which endpoints to use during a contact.
//
// When a contact uses multiple `DataflowEndpointConfig` objects, each `Config` must match a `DataflowEndpoint` in the same group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataflowEndpointGroupV2 := awscdk.Aws_groundstation.NewCfnDataflowEndpointGroupV2(this, jsii.String("MyCfnDataflowEndpointGroupV2"), &CfnDataflowEndpointGroupV2Props{
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	Endpoints: []interface{}{
//   		&CreateEndpointDetailsProperty{
//   			DownlinkAwsGroundStationAgentEndpoint: &DownlinkAwsGroundStationAgentEndpointProperty{
//   				DataflowDetails: &DownlinkDataflowDetailsProperty{
//   					AgentConnectionDetails: &DownlinkConnectionDetailsProperty{
//   						AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   							SocketAddress: &RangedSocketAddressProperty{
//   								Name: jsii.String("name"),
//   								PortRange: &IntegerRangeProperty{
//   									Maximum: jsii.Number(123),
//   									Minimum: jsii.Number(123),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   						EgressAddressAndPort: &ConnectionDetailsProperty{
//   							SocketAddress: &SocketAddressProperty{
//   								Name: jsii.String("name"),
//   								Port: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   			UplinkAwsGroundStationAgentEndpoint: &UplinkAwsGroundStationAgentEndpointProperty{
//   				DataflowDetails: &UplinkDataflowDetailsProperty{
//   					AgentConnectionDetails: &UplinkConnectionDetailsProperty{
//   						AgentIpAndPortAddress: &RangedConnectionDetailsProperty{
//   							SocketAddress: &RangedSocketAddressProperty{
//   								Name: jsii.String("name"),
//   								PortRange: &IntegerRangeProperty{
//   									Maximum: jsii.Number(123),
//   									Minimum: jsii.Number(123),
//   								},
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   						IngressAddressAndPort: &ConnectionDetailsProperty{
//   							SocketAddress: &SocketAddressProperty{
//   								Name: jsii.String("name"),
//   								Port: jsii.Number(123),
//   							},
//
//   							// the properties below are optional
//   							Mtu: jsii.Number(123),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-dataflowendpointgroupv2.html
//
type CfnDataflowEndpointGroupV2 interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsgroundstation.IDataflowEndpointGroupV2Ref
	awscdk.ITaggableV2
	AttrArn() *string
	// Information about the endpoint details.
	AttrEndpointDetails() awscdk.IResolvable
	AttrId() *string
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state.
	ContactPostPassDurationSeconds() *float64
	SetContactPostPassDurationSeconds(val *float64)
	// Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state.
	ContactPrePassDurationSeconds() *float64
	SetContactPrePassDurationSeconds(val *float64)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A reference to a DataflowEndpointGroupV2 resource.
	DataflowEndpointGroupV2Ref() *interfacesawsgroundstation.DataflowEndpointGroupV2Reference
	// List of endpoints for the dataflow endpoint group.
	Endpoints() interface{}
	SetEndpoints(val interface{})
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
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tags assigned to a resource.
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

// The jsii proxy struct for CfnDataflowEndpointGroupV2
type jsiiProxy_CfnDataflowEndpointGroupV2 struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsgroundstationIDataflowEndpointGroupV2Ref
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) AttrEndpointDetails() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrEndpointDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) ContactPostPassDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPostPassDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) ContactPrePassDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"contactPrePassDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) DataflowEndpointGroupV2Ref() *interfacesawsgroundstation.DataflowEndpointGroupV2Reference {
	var returns *interfacesawsgroundstation.DataflowEndpointGroupV2Reference
	_jsii_.Get(
		j,
		"dataflowEndpointGroupV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) Endpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::GroundStation::DataflowEndpointGroupV2`.
func NewCfnDataflowEndpointGroupV2(scope constructs.Construct, id *string, props *CfnDataflowEndpointGroupV2Props) CfnDataflowEndpointGroupV2 {
	_init_.Initialize()

	if err := validateNewCfnDataflowEndpointGroupV2Parameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDataflowEndpointGroupV2{}

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GroundStation::DataflowEndpointGroupV2`.
func NewCfnDataflowEndpointGroupV2_Override(c CfnDataflowEndpointGroupV2, scope constructs.Construct, id *string, props *CfnDataflowEndpointGroupV2Props) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2)SetContactPostPassDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"contactPostPassDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2)SetContactPrePassDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"contactPrePassDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2)SetEndpoints(val interface{}) {
	if err := j.validateSetEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_CfnDataflowEndpointGroupV2)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnDataflowEndpointGroupV2_ArnForDataflowEndpointGroupV2(resource interfacesawsgroundstation.IDataflowEndpointGroupV2Ref) *string {
	_init_.Initialize()

	if err := validateCfnDataflowEndpointGroupV2_ArnForDataflowEndpointGroupV2Parameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
		"arnForDataflowEndpointGroupV2",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnDataflowEndpointGroupV2.
func CfnDataflowEndpointGroupV2_IsCfnDataflowEndpointGroupV2(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataflowEndpointGroupV2_IsCfnDataflowEndpointGroupV2Parameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
		"isCfnDataflowEndpointGroupV2",
		[]interface{}{x},
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
func CfnDataflowEndpointGroupV2_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataflowEndpointGroupV2_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnDataflowEndpointGroupV2_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataflowEndpointGroupV2_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
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
func CfnDataflowEndpointGroupV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDataflowEndpointGroupV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataflowEndpointGroupV2_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_groundstation.CfnDataflowEndpointGroupV2",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDataflowEndpointGroupV2) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

