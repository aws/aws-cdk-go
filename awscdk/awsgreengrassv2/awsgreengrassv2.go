package awsgreengrassv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsgreengrassv2/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::GreengrassV2::ComponentVersion`.
//
// TODO: EXAMPLE
//
type CfnComponentVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrComponentName() *string
	AttrComponentVersion() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	InlineRecipe() *string
	SetInlineRecipe(val *string)
	LambdaFunction() interface{}
	SetLambdaFunction(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnComponentVersion
type jsiiProxy_CfnComponentVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponentVersion) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) AttrComponentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComponentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) AttrComponentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrComponentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) InlineRecipe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineRecipe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) LambdaFunction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::GreengrassV2::ComponentVersion`.
func NewCfnComponentVersion(scope awscdk.Construct, id *string, props *CfnComponentVersionProps) CfnComponentVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnComponentVersion{}

	_jsii_.Create(
		"monocdk.aws_greengrassv2.CfnComponentVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::GreengrassV2::ComponentVersion`.
func NewCfnComponentVersion_Override(c CfnComponentVersion, scope awscdk.Construct, id *string, props *CfnComponentVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_greengrassv2.CfnComponentVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponentVersion) SetInlineRecipe(val *string) {
	_jsii_.Set(
		j,
		"inlineRecipe",
		val,
	)
}

func (j *jsiiProxy_CfnComponentVersion) SetLambdaFunction(val interface{}) {
	_jsii_.Set(
		j,
		"lambdaFunction",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnComponentVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_greengrassv2.CfnComponentVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnComponentVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_greengrassv2.CfnComponentVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnComponentVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_greengrassv2.CfnComponentVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponentVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_greengrassv2.CfnComponentVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnComponentVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComponentVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnComponentVersion) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnComponentVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnComponentVersion_ComponentDependencyRequirementProperty struct {
	// `CfnComponentVersion.ComponentDependencyRequirementProperty.DependencyType`.
	DependencyType *string `json:"dependencyType"`
	// `CfnComponentVersion.ComponentDependencyRequirementProperty.VersionRequirement`.
	VersionRequirement *string `json:"versionRequirement"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_ComponentPlatformProperty struct {
	// `CfnComponentVersion.ComponentPlatformProperty.Attributes`.
	Attributes interface{} `json:"attributes"`
	// `CfnComponentVersion.ComponentPlatformProperty.Name`.
	Name *string `json:"name"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaContainerParamsProperty struct {
	// `CfnComponentVersion.LambdaContainerParamsProperty.Devices`.
	Devices interface{} `json:"devices"`
	// `CfnComponentVersion.LambdaContainerParamsProperty.MemorySizeInKB`.
	MemorySizeInKb *float64 `json:"memorySizeInKb"`
	// `CfnComponentVersion.LambdaContainerParamsProperty.MountROSysfs`.
	MountRoSysfs interface{} `json:"mountRoSysfs"`
	// `CfnComponentVersion.LambdaContainerParamsProperty.Volumes`.
	Volumes interface{} `json:"volumes"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaDeviceMountProperty struct {
	// `CfnComponentVersion.LambdaDeviceMountProperty.AddGroupOwner`.
	AddGroupOwner interface{} `json:"addGroupOwner"`
	// `CfnComponentVersion.LambdaDeviceMountProperty.Path`.
	Path *string `json:"path"`
	// `CfnComponentVersion.LambdaDeviceMountProperty.Permission`.
	Permission *string `json:"permission"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaEventSourceProperty struct {
	// `CfnComponentVersion.LambdaEventSourceProperty.Topic`.
	Topic *string `json:"topic"`
	// `CfnComponentVersion.LambdaEventSourceProperty.Type`.
	Type *string `json:"type"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaExecutionParametersProperty struct {
	// `CfnComponentVersion.LambdaExecutionParametersProperty.EnvironmentVariables`.
	EnvironmentVariables interface{} `json:"environmentVariables"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.EventSources`.
	EventSources interface{} `json:"eventSources"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.ExecArgs`.
	ExecArgs *[]*string `json:"execArgs"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.InputPayloadEncodingType`.
	InputPayloadEncodingType *string `json:"inputPayloadEncodingType"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.LinuxProcessParams`.
	LinuxProcessParams interface{} `json:"linuxProcessParams"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.MaxIdleTimeInSeconds`.
	MaxIdleTimeInSeconds *float64 `json:"maxIdleTimeInSeconds"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.MaxInstancesCount`.
	MaxInstancesCount *float64 `json:"maxInstancesCount"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.MaxQueueSize`.
	MaxQueueSize *float64 `json:"maxQueueSize"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.Pinned`.
	Pinned interface{} `json:"pinned"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.StatusTimeoutInSeconds`.
	StatusTimeoutInSeconds *float64 `json:"statusTimeoutInSeconds"`
	// `CfnComponentVersion.LambdaExecutionParametersProperty.TimeoutInSeconds`.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaFunctionRecipeSourceProperty struct {
	// `CfnComponentVersion.LambdaFunctionRecipeSourceProperty.ComponentDependencies`.
	ComponentDependencies interface{} `json:"componentDependencies"`
	// `CfnComponentVersion.LambdaFunctionRecipeSourceProperty.ComponentLambdaParameters`.
	ComponentLambdaParameters interface{} `json:"componentLambdaParameters"`
	// `CfnComponentVersion.LambdaFunctionRecipeSourceProperty.ComponentName`.
	ComponentName *string `json:"componentName"`
	// `CfnComponentVersion.LambdaFunctionRecipeSourceProperty.ComponentPlatforms`.
	ComponentPlatforms interface{} `json:"componentPlatforms"`
	// `CfnComponentVersion.LambdaFunctionRecipeSourceProperty.ComponentVersion`.
	ComponentVersion *string `json:"componentVersion"`
	// `CfnComponentVersion.LambdaFunctionRecipeSourceProperty.LambdaArn`.
	LambdaArn *string `json:"lambdaArn"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaLinuxProcessParamsProperty struct {
	// `CfnComponentVersion.LambdaLinuxProcessParamsProperty.ContainerParams`.
	ContainerParams interface{} `json:"containerParams"`
	// `CfnComponentVersion.LambdaLinuxProcessParamsProperty.IsolationMode`.
	IsolationMode *string `json:"isolationMode"`
}

// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaVolumeMountProperty struct {
	// `CfnComponentVersion.LambdaVolumeMountProperty.AddGroupOwner`.
	AddGroupOwner interface{} `json:"addGroupOwner"`
	// `CfnComponentVersion.LambdaVolumeMountProperty.DestinationPath`.
	DestinationPath *string `json:"destinationPath"`
	// `CfnComponentVersion.LambdaVolumeMountProperty.Permission`.
	Permission *string `json:"permission"`
	// `CfnComponentVersion.LambdaVolumeMountProperty.SourcePath`.
	SourcePath *string `json:"sourcePath"`
}

// Properties for defining a `AWS::GreengrassV2::ComponentVersion`.
//
// TODO: EXAMPLE
//
type CfnComponentVersionProps struct {
	// `AWS::GreengrassV2::ComponentVersion.InlineRecipe`.
	InlineRecipe *string `json:"inlineRecipe"`
	// `AWS::GreengrassV2::ComponentVersion.LambdaFunction`.
	LambdaFunction interface{} `json:"lambdaFunction"`
	// `AWS::GreengrassV2::ComponentVersion.Tags`.
	Tags *map[string]*string `json:"tags"`
}

