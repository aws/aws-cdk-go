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
// Creates a component. Components are software that run on Greengrass core devices. After you develop and test a component on your core device, you can use this operation to upload your component to AWS IoT Greengrass . Then, you can deploy the component to other core devices.
//
// You can use this operation to do the following:
//
// - *Create components from recipes*
//
// Create a component from a recipe, which is a file that defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform capability. For more information, see [AWS IoT Greengrass component recipe reference](https://docs.aws.amazon.com/greengrass/v2/developerguide/component-recipe-reference.html) in the *AWS IoT Greengrass V2 Developer Guide* .
//
// To create a component from a recipe, specify `inlineRecipe` when you call this operation.
// - *Create components from Lambda functions*
//
// Create a component from an AWS Lambda function that runs on AWS IoT Greengrass . This creates a recipe and artifacts from the Lambda function's deployment package. You can use this operation to migrate Lambda functions from AWS IoT Greengrass V1 to AWS IoT Greengrass V2 .
//
// This function only accepts Lambda functions that use the following runtimes:
//
// - Python 2.7 – `python2.7`
// - Python 3.7 – `python3.7`
// - Python 3.8 – `python3.8`
// - Java 8 – `java8`
// - Node.js 10 – `nodejs10.x`
// - Node.js 12 – `nodejs12.x`
//
// To create a component from a Lambda function, specify `lambdaFunction` when you call this operation.
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

// Contains information about a component dependency for a Lambda function component.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_ComponentDependencyRequirementProperty struct {
	// The type of this dependency. Choose from the following options:.
	//
	// - `SOFT` – The component doesn't restart if the dependency changes state.
	// - `HARD` – The component restarts if the dependency changes state.
	//
	// Default: `HARD`
	DependencyType *string `json:"dependencyType" yaml:"dependencyType"`
	// The component version requirement for the component dependency.
	//
	// AWS IoT Greengrass uses semantic version constraints. For more information, see [Semantic Versioning](https://docs.aws.amazon.com/https://semver.org/) .
	VersionRequirement *string `json:"versionRequirement" yaml:"versionRequirement"`
}

// Contains information about a platform that a component supports.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_ComponentPlatformProperty struct {
	// A dictionary of attributes for the platform.
	//
	// The  software defines the `os` and `platform` by default. You can specify additional platform attributes for a core device when you deploy the Greengrass nucleus component. For more information, see the [Greengrass nucleus component](https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-nucleus-component.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	Attributes interface{} `json:"attributes" yaml:"attributes"`
	// The friendly name of the platform. This name helps you identify the platform.
	//
	// If you omit this parameter, AWS IoT Greengrass creates a friendly name from the `os` and `architecture` of the platform.
	Name *string `json:"name" yaml:"name"`
}

// Contains information about a container in which AWS Lambda functions run on Greengrass core devices.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaContainerParamsProperty struct {
	// The list of system devices that the container can access.
	Devices interface{} `json:"devices" yaml:"devices"`
	// The memory size of the container, expressed in kilobytes.
	//
	// Default: `16384` (16 MB)
	MemorySizeInKb *float64 `json:"memorySizeInKb" yaml:"memorySizeInKb"`
	// Whether or not the container can read information from the device's `/sys` folder.
	//
	// Default: `false`
	MountRoSysfs interface{} `json:"mountRoSysfs" yaml:"mountRoSysfs"`
	// The list of volumes that the container can access.
	Volumes interface{} `json:"volumes" yaml:"volumes"`
}

// Contains information about a device that Linux processes in a container can access.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaDeviceMountProperty struct {
	// Whether or not to add the component's system user as an owner of the device.
	//
	// Default: `false`
	AddGroupOwner interface{} `json:"addGroupOwner" yaml:"addGroupOwner"`
	// The mount path for the device in the file system.
	Path *string `json:"path" yaml:"path"`
	// The permission to access the device: read/only ( `ro` ) or read/write ( `rw` ).
	//
	// Default: `ro`
	Permission *string `json:"permission" yaml:"permission"`
}

// Contains information about an event source for an AWS Lambda function.
//
// The event source defines the topics on which this Lambda function subscribes to receive messages that run the function.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaEventSourceProperty struct {
	// The topic to which to subscribe to receive event messages.
	Topic *string `json:"topic" yaml:"topic"`
	// The type of event source. Choose from the following options:.
	//
	// - `PUB_SUB` – Subscribe to local publish/subscribe messages. This event source type doesn't support MQTT wildcards ( `+` and `#` ) in the event source topic.
	// - `IOT_CORE` – Subscribe to AWS IoT Core MQTT messages. This event source type supports MQTT wildcards ( `+` and `#` ) in the event source topic.
	Type *string `json:"type" yaml:"type"`
}

// Contains parameters for a Lambda function that runs on AWS IoT Greengrass .
//
// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaExecutionParametersProperty struct {
	// The map of environment variables that are available to the Lambda function when it runs.
	EnvironmentVariables interface{} `json:"environmentVariables" yaml:"environmentVariables"`
	// The list of event sources to which to subscribe to receive work messages.
	//
	// The Lambda function runs when it receives a message from an event source. You can subscribe this function to local publish/subscribe messages and AWS IoT Core MQTT messages.
	EventSources interface{} `json:"eventSources" yaml:"eventSources"`
	// The list of arguments to pass to the Lambda function when it runs.
	ExecArgs *[]*string `json:"execArgs" yaml:"execArgs"`
	// The encoding type that the Lambda function supports.
	//
	// Default: `json`
	InputPayloadEncodingType *string `json:"inputPayloadEncodingType" yaml:"inputPayloadEncodingType"`
	// The parameters for the Linux process that contains the Lambda function.
	LinuxProcessParams interface{} `json:"linuxProcessParams" yaml:"linuxProcessParams"`
	// The maximum amount of time in seconds that a non-pinned Lambda function can idle before the  software stops its process.
	MaxIdleTimeInSeconds *float64 `json:"maxIdleTimeInSeconds" yaml:"maxIdleTimeInSeconds"`
	// The maximum number of instances that a non-pinned Lambda function can run at the same time.
	MaxInstancesCount *float64 `json:"maxInstancesCount" yaml:"maxInstancesCount"`
	// The maximum size of the message queue for the Lambda function component.
	//
	// The Greengrass core device stores messages in a FIFO (first-in-first-out) queue until it can run the Lambda function to consume each message.
	MaxQueueSize *float64 `json:"maxQueueSize" yaml:"maxQueueSize"`
	// Whether or not the Lambda function is pinned, or long-lived.
	//
	// - A pinned Lambda function starts when the  starts and keeps running in its own container.
	// - A non-pinned Lambda function starts only when it receives a work item and exists after it idles for `maxIdleTimeInSeconds` . If the function has multiple work items, the  software creates multiple instances of the function.
	//
	// Default: `true`
	Pinned interface{} `json:"pinned" yaml:"pinned"`
	// The interval in seconds at which a pinned (also known as long-lived) Lambda function component sends status updates to the Lambda manager component.
	StatusTimeoutInSeconds *float64 `json:"statusTimeoutInSeconds" yaml:"statusTimeoutInSeconds"`
	// The maximum amount of time in seconds that the Lambda function can process a work item.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

// Contains information about an AWS Lambda function to import to create a component.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaFunctionRecipeSourceProperty struct {
	// The component versions on which this Lambda function component depends.
	ComponentDependencies interface{} `json:"componentDependencies" yaml:"componentDependencies"`
	// The system and runtime parameters for the Lambda function as it runs on the Greengrass core device.
	ComponentLambdaParameters interface{} `json:"componentLambdaParameters" yaml:"componentLambdaParameters"`
	// The name of the component.
	//
	// Defaults to the name of the Lambda function.
	ComponentName *string `json:"componentName" yaml:"componentName"`
	// The platforms that the component version supports.
	ComponentPlatforms interface{} `json:"componentPlatforms" yaml:"componentPlatforms"`
	// The version of the component.
	//
	// Defaults to the version of the Lambda function as a semantic version. For example, if your function version is `3` , the component version becomes `3.0.0` .
	ComponentVersion *string `json:"componentVersion" yaml:"componentVersion"`
	// The ARN of the Lambda function.
	//
	// The ARN must include the version of the function to import. You can't use version aliases like `$LATEST` .
	LambdaArn *string `json:"lambdaArn" yaml:"lambdaArn"`
}

// Contains parameters for a Linux process that contains an AWS Lambda function.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaLinuxProcessParamsProperty struct {
	// The parameters for the container in which the Lambda function runs.
	ContainerParams interface{} `json:"containerParams" yaml:"containerParams"`
	// The isolation mode for the process that contains the Lambda function.
	//
	// The process can run in an isolated runtime environment inside the AWS IoT Greengrass container, or as a regular process outside any container.
	//
	// Default: `GreengrassContainer`
	IsolationMode *string `json:"isolationMode" yaml:"isolationMode"`
}

// Contains information about a volume that Linux processes in a container can access.
//
// When you define a volume, the  software mounts the source files to the destination inside the container.
//
// TODO: EXAMPLE
//
type CfnComponentVersion_LambdaVolumeMountProperty struct {
	// Whether or not to add the AWS IoT Greengrass user group as an owner of the volume.
	//
	// Default: `false`
	AddGroupOwner interface{} `json:"addGroupOwner" yaml:"addGroupOwner"`
	// The path to the logical volume in the file system.
	DestinationPath *string `json:"destinationPath" yaml:"destinationPath"`
	// The permission to access the volume: read/only ( `ro` ) or read/write ( `rw` ).
	//
	// Default: `ro`
	Permission *string `json:"permission" yaml:"permission"`
	// The path to the physical volume in the file system.
	SourcePath *string `json:"sourcePath" yaml:"sourcePath"`
}

// Properties for defining a `CfnComponentVersion`.
//
// TODO: EXAMPLE
//
type CfnComponentVersionProps struct {
	// The recipe to use to create the component.
	//
	// The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.
	//
	// You must specify either `InlineRecipe` or `LambdaFunction` .
	InlineRecipe *string `json:"inlineRecipe" yaml:"inlineRecipe"`
	// The parameters to create a component from a Lambda function.
	//
	// You must specify either `InlineRecipe` or `LambdaFunction` .
	LambdaFunction interface{} `json:"lambdaFunction" yaml:"lambdaFunction"`
	// Application-specific metadata to attach to the component version.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

