package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define an EventBridge Api Destination.
//
// TODO: EXAMPLE
//
type ApiDestination interface {
	awscdk.Resource
	IApiDestination
	ApiDestinationArn() *string
	ApiDestinationName() *string
	Connection() IConnection
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for ApiDestination
type jsiiProxy_ApiDestination struct {
	internal.Type__awscdkResource
	jsiiProxy_IApiDestination
}

func (j *jsiiProxy_ApiDestination) ApiDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiDestination) ApiDestinationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDestinationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiDestination) Connection() IConnection {
	var returns IConnection
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiDestination) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiDestination) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewApiDestination(scope constructs.Construct, id *string, props *ApiDestinationProps) ApiDestination {
	_init_.Initialize()

	j := jsiiProxy_ApiDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.ApiDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApiDestination_Override(a ApiDestination, scope constructs.Construct, id *string, props *ApiDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.ApiDestination",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ApiDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.ApiDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ApiDestination_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.ApiDestination",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (a *jsiiProxy_ApiDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_ApiDestination) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (a *jsiiProxy_ApiDestination) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (a *jsiiProxy_ApiDestination) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_ApiDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The event API Destination properties.
//
// TODO: EXAMPLE
//
type ApiDestinationProps struct {
	// The ARN of the connection to use for the API destination.
	Connection IConnection `json:"connection" yaml:"connection"`
	// The URL to the HTTP invocation endpoint for the API destination..
	Endpoint *string `json:"endpoint" yaml:"endpoint"`
	// The name for the API destination.
	ApiDestinationName *string `json:"apiDestinationName" yaml:"apiDestinationName"`
	// A description for the API destination.
	Description *string `json:"description" yaml:"description"`
	// The method to use for the request to the HTTP invocation endpoint.
	HttpMethod HttpMethod `json:"httpMethod" yaml:"httpMethod"`
	// The maximum number of requests per second to send to the HTTP invocation endpoint.
	RateLimitPerSecond *float64 `json:"rateLimitPerSecond" yaml:"rateLimitPerSecond"`
}

// Define an EventBridge Archive.
//
// TODO: EXAMPLE
//
type Archive interface {
	awscdk.Resource
	ArchiveArn() *string
	ArchiveName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Archive
type jsiiProxy_Archive struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Archive) ArchiveArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Archive) ArchiveName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Archive) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Archive) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Archive) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Archive) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewArchive(scope constructs.Construct, id *string, props *ArchiveProps) Archive {
	_init_.Initialize()

	j := jsiiProxy_Archive{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Archive",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewArchive_Override(a Archive, scope constructs.Construct, id *string, props *ArchiveProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Archive",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Archive_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Archive",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Archive_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Archive",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (a *jsiiProxy_Archive) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_Archive) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (a *jsiiProxy_Archive) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (a *jsiiProxy_Archive) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_Archive) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The event archive properties.
//
// TODO: EXAMPLE
//
type ArchiveProps struct {
	// An event pattern to use to filter events sent to the archive.
	EventPattern *EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// The name of the archive.
	ArchiveName *string `json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	Description *string `json:"description" yaml:"description"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely.
	Retention awscdk.Duration `json:"retention" yaml:"retention"`
	// The event source associated with the archive.
	SourceEventBus IEventBus `json:"sourceEventBus" yaml:"sourceEventBus"`
}

// Authorization type for an API Destination Connection.
//
// TODO: EXAMPLE
//
type Authorization interface {
}

// The jsii proxy struct for Authorization
type jsiiProxy_Authorization struct {
	_ byte // padding
}

func NewAuthorization_Override(a Authorization) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Authorization",
		nil, // no parameters
		a,
	)
}

// Use API key authorization.
//
// API key authorization has two components: an API key name and an API key value.
// What these are depends on the target of your connection.
func Authorization_ApiKey(apiKeyName *string, apiKeyValue awscdk.SecretValue) Authorization {
	_init_.Initialize()

	var returns Authorization

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Authorization",
		"apiKey",
		[]interface{}{apiKeyName, apiKeyValue},
		&returns,
	)

	return returns
}

// Use username and password authorization.
func Authorization_Basic(username *string, password awscdk.SecretValue) Authorization {
	_init_.Initialize()

	var returns Authorization

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Authorization",
		"basic",
		[]interface{}{username, password},
		&returns,
	)

	return returns
}

// Use OAuth authorization.
func Authorization_Oauth(props *OAuthAuthorizationProps) Authorization {
	_init_.Initialize()

	var returns Authorization

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Authorization",
		"oauth",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The event archive base properties.
//
// TODO: EXAMPLE
//
type BaseArchiveProps struct {
	// An event pattern to use to filter events sent to the archive.
	EventPattern *EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// The name of the archive.
	ArchiveName *string `json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	Description *string `json:"description" yaml:"description"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely.
	Retention awscdk.Duration `json:"retention" yaml:"retention"`
}

// A CloudFormation `AWS::Events::ApiDestination`.
//
// Creates an API destination, which is an HTTP invocation endpoint configured as a target for events.
//
// When using ApiDesinations with OAuth authentication we recommend these best practices:
//
// - Create a secret in Secrets Manager with your OAuth credentials.
// - Reference that secret in your CloudFormation template for `AWS::Events::Connection` using CloudFormation dynamic reference syntax. For more information, see [Secrets Manager secrets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) .
//
// When the Connection resource is created the secret will be passed to EventBridge and stored in the customer account using “Service Linked Secrets,” effectively creating two secrets. This will minimize the cost because the original secret is only accessed when a CloudFormation template is created or updated, not every time an event is sent to the ApiDestination. The secret stored in the customer account by EventBridge is the one used for each event sent to the ApiDestination and AWS is responsible for the fees.
//
// > The secret stored in the customer account by EventBridge can’t be updated directly, only when a CloudFormation template is updated.
//
// For examples of CloudFormation templates that use secrets, see [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-connection.html#aws-resource-events-connection--examples) .
//
// TODO: EXAMPLE
//
type CfnApiDestination interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectionArn() *string
	SetConnectionArn(val *string)
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	HttpMethod() *string
	SetHttpMethod(val *string)
	InvocationEndpoint() *string
	SetInvocationEndpoint(val *string)
	InvocationRateLimitPerSecond() *float64
	SetInvocationRateLimitPerSecond(val *float64)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnApiDestination
type jsiiProxy_CfnApiDestination struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiDestination) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) InvocationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) InvocationRateLimitPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"invocationRateLimitPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiDestination) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Events::ApiDestination`.
func NewCfnApiDestination(scope constructs.Construct, id *string, props *CfnApiDestinationProps) CfnApiDestination {
	_init_.Initialize()

	j := jsiiProxy_CfnApiDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnApiDestination",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Events::ApiDestination`.
func NewCfnApiDestination_Override(c CfnApiDestination, scope constructs.Construct, id *string, props *CfnApiDestinationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnApiDestination",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiDestination) SetConnectionArn(val *string) {
	_jsii_.Set(
		j,
		"connectionArn",
		val,
	)
}

func (j *jsiiProxy_CfnApiDestination) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApiDestination) SetHttpMethod(val *string) {
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_CfnApiDestination) SetInvocationEndpoint(val *string) {
	_jsii_.Set(
		j,
		"invocationEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnApiDestination) SetInvocationRateLimitPerSecond(val *float64) {
	_jsii_.Set(
		j,
		"invocationRateLimitPerSecond",
		val,
	)
}

func (j *jsiiProxy_CfnApiDestination) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApiDestination_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnApiDestination",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApiDestination_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnApiDestination",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnApiDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnApiDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiDestination_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.CfnApiDestination",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnApiDestination) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApiDestination) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApiDestination) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnApiDestination) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnApiDestination) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnApiDestination) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnApiDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnApiDestination) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnApiDestination) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApiDestination) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnApiDestination) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApiDestination) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnApiDestination) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnApiDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApiDestination) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnApiDestination`.
//
// TODO: EXAMPLE
//
type CfnApiDestinationProps struct {
	// The ARN of the connection to use for the API destination.
	//
	// The destination endpoint must support the authorization type specified for the connection.
	ConnectionArn *string `json:"connectionArn" yaml:"connectionArn"`
	// The method to use for the request to the HTTP invocation endpoint.
	HttpMethod *string `json:"httpMethod" yaml:"httpMethod"`
	// The URL to the HTTP invocation endpoint for the API destination.
	InvocationEndpoint *string `json:"invocationEndpoint" yaml:"invocationEndpoint"`
	// A description for the API destination to create.
	Description *string `json:"description" yaml:"description"`
	// The maximum number of requests per second to send to the HTTP invocation endpoint.
	InvocationRateLimitPerSecond *float64 `json:"invocationRateLimitPerSecond" yaml:"invocationRateLimitPerSecond"`
	// The name for the API destination to create.
	Name *string `json:"name" yaml:"name"`
}

// A CloudFormation `AWS::Events::Archive`.
//
// Creates an archive of events with the specified settings. When you create an archive, incoming events might not immediately start being sent to the archive. Allow a short period of time for changes to take effect. If you do not specify a pattern to filter events sent to the archive, all events are sent to the archive except replayed events. Replayed events are not sent to an archive.
//
// TODO: EXAMPLE
//
type CfnArchive interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ArchiveName() *string
	SetArchiveName(val *string)
	AttrArchiveName() *string
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	EventPattern() interface{}
	SetEventPattern(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RetentionDays() *float64
	SetRetentionDays(val *float64)
	SourceArn() *string
	SetSourceArn(val *string)
	Stack() awscdk.Stack
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnArchive
type jsiiProxy_CfnArchive struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnArchive) ArchiveName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archiveName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) AttrArchiveName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArchiveName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) EventPattern() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) RetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnArchive) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Events::Archive`.
func NewCfnArchive(scope constructs.Construct, id *string, props *CfnArchiveProps) CfnArchive {
	_init_.Initialize()

	j := jsiiProxy_CfnArchive{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnArchive",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Events::Archive`.
func NewCfnArchive_Override(c CfnArchive, scope constructs.Construct, id *string, props *CfnArchiveProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnArchive",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnArchive) SetArchiveName(val *string) {
	_jsii_.Set(
		j,
		"archiveName",
		val,
	)
}

func (j *jsiiProxy_CfnArchive) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnArchive) SetEventPattern(val interface{}) {
	_jsii_.Set(
		j,
		"eventPattern",
		val,
	)
}

func (j *jsiiProxy_CfnArchive) SetRetentionDays(val *float64) {
	_jsii_.Set(
		j,
		"retentionDays",
		val,
	)
}

func (j *jsiiProxy_CfnArchive) SetSourceArn(val *string) {
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnArchive_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnArchive",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnArchive_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnArchive",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnArchive_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnArchive",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnArchive_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.CfnArchive",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnArchive) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnArchive) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnArchive) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnArchive) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnArchive) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnArchive) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnArchive) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnArchive) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnArchive) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnArchive) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnArchive) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnArchive) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnArchive) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnArchive) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnArchive) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnArchive`.
//
// TODO: EXAMPLE
//
type CfnArchiveProps struct {
	// The ARN of the event bus that sends events to the archive.
	SourceArn *string `json:"sourceArn" yaml:"sourceArn"`
	// The name for the archive to create.
	ArchiveName *string `json:"archiveName" yaml:"archiveName"`
	// A description for the archive.
	Description *string `json:"description" yaml:"description"`
	// An event pattern to use to filter events sent to the archive.
	EventPattern interface{} `json:"eventPattern" yaml:"eventPattern"`
	// The number of days to retain events for.
	//
	// Default value is 0. If set to 0, events are retained indefinitely
	RetentionDays *float64 `json:"retentionDays" yaml:"retentionDays"`
}

// A CloudFormation `AWS::Events::Connection`.
//
// Creates a connection. A connection defines the authorization type and credentials to use for authorization with an API destination HTTP endpoint.
//
// TODO: EXAMPLE
//
type CfnConnection interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrSecretArn() *string
	AuthorizationType() *string
	SetAuthorizationType(val *string)
	AuthParameters() interface{}
	SetAuthParameters(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnConnection
type jsiiProxy_CfnConnection struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnConnection) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AttrSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) AuthParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnConnection) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Events::Connection`.
func NewCfnConnection(scope constructs.Construct, id *string, props *CfnConnectionProps) CfnConnection {
	_init_.Initialize()

	j := jsiiProxy_CfnConnection{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnConnection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Events::Connection`.
func NewCfnConnection_Override(c CfnConnection, scope constructs.Construct, id *string, props *CfnConnectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnConnection",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnConnection) SetAuthorizationType(val *string) {
	_jsii_.Set(
		j,
		"authorizationType",
		val,
	)
}

func (j *jsiiProxy_CfnConnection) SetAuthParameters(val interface{}) {
	_jsii_.Set(
		j,
		"authParameters",
		val,
	)
}

func (j *jsiiProxy_CfnConnection) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnConnection) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnConnection_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnConnection",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnConnection_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnConnection",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnConnection_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.CfnConnection",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnConnection) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnConnection) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnConnection) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnConnection) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnConnection) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnConnection) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnConnection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnConnection) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnConnection) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnConnection) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnConnection) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnConnection) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnConnection) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnConnection) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnConnection_ApiKeyAuthParametersProperty struct {
	// `CfnConnection.ApiKeyAuthParametersProperty.ApiKeyName`.
	ApiKeyName *string `json:"apiKeyName" yaml:"apiKeyName"`
	// `CfnConnection.ApiKeyAuthParametersProperty.ApiKeyValue`.
	ApiKeyValue *string `json:"apiKeyValue" yaml:"apiKeyValue"`
}

// TODO: EXAMPLE
//
type CfnConnection_AuthParametersProperty struct {
	// `CfnConnection.AuthParametersProperty.ApiKeyAuthParameters`.
	ApiKeyAuthParameters interface{} `json:"apiKeyAuthParameters" yaml:"apiKeyAuthParameters"`
	// `CfnConnection.AuthParametersProperty.BasicAuthParameters`.
	BasicAuthParameters interface{} `json:"basicAuthParameters" yaml:"basicAuthParameters"`
	// `CfnConnection.AuthParametersProperty.InvocationHttpParameters`.
	InvocationHttpParameters interface{} `json:"invocationHttpParameters" yaml:"invocationHttpParameters"`
	// `CfnConnection.AuthParametersProperty.OAuthParameters`.
	OAuthParameters interface{} `json:"oAuthParameters" yaml:"oAuthParameters"`
}

// TODO: EXAMPLE
//
type CfnConnection_BasicAuthParametersProperty struct {
	// `CfnConnection.BasicAuthParametersProperty.Password`.
	Password *string `json:"password" yaml:"password"`
	// `CfnConnection.BasicAuthParametersProperty.Username`.
	Username *string `json:"username" yaml:"username"`
}

// TODO: EXAMPLE
//
type CfnConnection_ClientParametersProperty struct {
	// `CfnConnection.ClientParametersProperty.ClientID`.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// `CfnConnection.ClientParametersProperty.ClientSecret`.
	ClientSecret *string `json:"clientSecret" yaml:"clientSecret"`
}

// Contains additional parameters for the connection.
//
// TODO: EXAMPLE
//
type CfnConnection_ConnectionHttpParametersProperty struct {
	// Contains additional body string parameters for the connection.
	BodyParameters interface{} `json:"bodyParameters" yaml:"bodyParameters"`
	// Contains additional header parameters for the connection.
	HeaderParameters interface{} `json:"headerParameters" yaml:"headerParameters"`
	// Contains additional query string parameters for the connection.
	QueryStringParameters interface{} `json:"queryStringParameters" yaml:"queryStringParameters"`
}

// TODO: EXAMPLE
//
type CfnConnection_OAuthParametersProperty struct {
	// `CfnConnection.OAuthParametersProperty.AuthorizationEndpoint`.
	AuthorizationEndpoint *string `json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// `CfnConnection.OAuthParametersProperty.ClientParameters`.
	ClientParameters interface{} `json:"clientParameters" yaml:"clientParameters"`
	// `CfnConnection.OAuthParametersProperty.HttpMethod`.
	HttpMethod *string `json:"httpMethod" yaml:"httpMethod"`
	// `CfnConnection.OAuthParametersProperty.OAuthHttpParameters`.
	OAuthHttpParameters interface{} `json:"oAuthHttpParameters" yaml:"oAuthHttpParameters"`
}

// TODO: EXAMPLE
//
type CfnConnection_ParameterProperty struct {
	// `CfnConnection.ParameterProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnConnection.ParameterProperty.Value`.
	Value *string `json:"value" yaml:"value"`
	// `CfnConnection.ParameterProperty.IsValueSecret`.
	IsValueSecret interface{} `json:"isValueSecret" yaml:"isValueSecret"`
}

// Properties for defining a `CfnConnection`.
//
// TODO: EXAMPLE
//
type CfnConnectionProps struct {
	// The type of authorization to use for the connection.
	AuthorizationType *string `json:"authorizationType" yaml:"authorizationType"`
	// A `CreateConnectionAuthRequestParameters` object that contains the authorization parameters to use to authorize with the endpoint.
	AuthParameters interface{} `json:"authParameters" yaml:"authParameters"`
	// A description for the connection to create.
	Description *string `json:"description" yaml:"description"`
	// The name for the connection to create.
	Name *string `json:"name" yaml:"name"`
}

// A CloudFormation `AWS::Events::EventBus`.
//
// Creates a new event bus within your account. This can be a custom event bus which you can use to receive events from your custom applications and services, or it can be a partner event bus which can be matched to a partner event source.
//
// TODO: EXAMPLE
//
type CfnEventBus interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrName() *string
	AttrPolicy() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EventSourceName() *string
	SetEventSourceName(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() *[]*CfnEventBus_TagEntryProperty
	SetTags(val *[]*CfnEventBus_TagEntryProperty)
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnEventBus
type jsiiProxy_CfnEventBus struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventBus) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) AttrPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) EventSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) Tags() *[]*CfnEventBus_TagEntryProperty {
	var returns *[]*CfnEventBus_TagEntryProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBus) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Events::EventBus`.
func NewCfnEventBus(scope constructs.Construct, id *string, props *CfnEventBusProps) CfnEventBus {
	_init_.Initialize()

	j := jsiiProxy_CfnEventBus{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnEventBus",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Events::EventBus`.
func NewCfnEventBus_Override(c CfnEventBus, scope constructs.Construct, id *string, props *CfnEventBusProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnEventBus",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventBus) SetEventSourceName(val *string) {
	_jsii_.Set(
		j,
		"eventSourceName",
		val,
	)
}

func (j *jsiiProxy_CfnEventBus) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnEventBus) SetTags(val *[]*CfnEventBus_TagEntryProperty) {
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
func CfnEventBus_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnEventBus",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventBus_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnEventBus",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnEventBus_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnEventBus",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventBus_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.CfnEventBus",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEventBus) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEventBus) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEventBus) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnEventBus) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEventBus) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEventBus) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventBus) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEventBus) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEventBus) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEventBus) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEventBus) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventBus) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEventBus) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnEventBus) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventBus) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// TODO: EXAMPLE
//
type CfnEventBus_TagEntryProperty struct {
	// `CfnEventBus.TagEntryProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnEventBus.TagEntryProperty.Value`.
	Value *string `json:"value" yaml:"value"`
}

// A CloudFormation `AWS::Events::EventBusPolicy`.
//
// Running `PutPermission` permits the specified AWS account or AWS organization to put events to the specified *event bus* . Amazon EventBridge (CloudWatch Events) rules in your account are triggered by these events arriving to an event bus in your account.
//
// For another account to send events to your account, that external account must have an EventBridge rule with your account's event bus as a target.
//
// To enable multiple AWS accounts to put events to your event bus, run `PutPermission` once for each of these accounts. Or, if all the accounts are members of the same AWS organization, you can run `PutPermission` once specifying `Principal` as "*" and specifying the AWS organization ID in `Condition` , to grant permissions to all accounts in that organization.
//
// If you grant permissions using an organization, then accounts in that organization must specify a `RoleArn` with proper permissions when they use `PutTarget` to add your account's event bus as a target. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .
//
// The permission policy on the event bus cannot exceed 10 KB in size.
//
// TODO: EXAMPLE
//
type CfnEventBusPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Action() *string
	SetAction(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Condition() interface{}
	SetCondition(val interface{})
	CreationStack() *[]*string
	EventBusName() *string
	SetEventBusName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Principal() *string
	SetPrincipal(val *string)
	Ref() *string
	Stack() awscdk.Stack
	Statement() interface{}
	SetStatement(val interface{})
	StatementId() *string
	SetStatementId(val *string)
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnEventBusPolicy
type jsiiProxy_CfnEventBusPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventBusPolicy) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) Condition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) Statement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) StatementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Events::EventBusPolicy`.
func NewCfnEventBusPolicy(scope constructs.Construct, id *string, props *CfnEventBusPolicyProps) CfnEventBusPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnEventBusPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnEventBusPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Events::EventBusPolicy`.
func NewCfnEventBusPolicy_Override(c CfnEventBusPolicy, scope constructs.Construct, id *string, props *CfnEventBusPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnEventBusPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventBusPolicy) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CfnEventBusPolicy) SetCondition(val interface{}) {
	_jsii_.Set(
		j,
		"condition",
		val,
	)
}

func (j *jsiiProxy_CfnEventBusPolicy) SetEventBusName(val *string) {
	_jsii_.Set(
		j,
		"eventBusName",
		val,
	)
}

func (j *jsiiProxy_CfnEventBusPolicy) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_CfnEventBusPolicy) SetStatement(val interface{}) {
	_jsii_.Set(
		j,
		"statement",
		val,
	)
}

func (j *jsiiProxy_CfnEventBusPolicy) SetStatementId(val *string) {
	_jsii_.Set(
		j,
		"statementId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventBusPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnEventBusPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventBusPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnEventBusPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnEventBusPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnEventBusPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventBusPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.CfnEventBusPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEventBusPolicy) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEventBusPolicy) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEventBusPolicy) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnEventBusPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEventBusPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEventBusPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventBusPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEventBusPolicy) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEventBusPolicy) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEventBusPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEventBusPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventBusPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEventBusPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnEventBusPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventBusPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A JSON string which you can use to limit the event bus permissions you are granting to only accounts that fulfill the condition.
//
// Currently, the only supported condition is membership in a certain AWS organization. The string must contain `Type` , `Key` , and `Value` fields. The `Value` field specifies the ID of the AWS organization. Following is an example value for `Condition` :
//
// `'{"Type" : "StringEquals", "Key": "aws:PrincipalOrgID", "Value": "o-1234567890"}'`
//
// TODO: EXAMPLE
//
type CfnEventBusPolicy_ConditionProperty struct {
	// Specifies the key for the condition.
	//
	// Currently the only supported key is `aws:PrincipalOrgID` .
	Key *string `json:"key" yaml:"key"`
	// Specifies the type of condition.
	//
	// Currently the only supported value is `StringEquals` .
	Type *string `json:"type" yaml:"type"`
	// Specifies the value for the key.
	//
	// Currently, this must be the ID of the organization.
	Value *string `json:"value" yaml:"value"`
}

// Properties for defining a `CfnEventBusPolicy`.
//
// TODO: EXAMPLE
//
type CfnEventBusPolicyProps struct {
	// An identifier string for the external account that you are granting permissions to.
	//
	// If you later want to revoke the permission for this external account, specify this `StatementId` when you run [RemovePermission](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_RemovePermission.html) .
	//
	// > Each `StatementId` must be unique.
	StatementId *string `json:"statementId" yaml:"statementId"`
	// The action that you are enabling the other account to perform.
	Action *string `json:"action" yaml:"action"`
	// This parameter enables you to limit the permission to accounts that fulfill a certain condition, such as being a member of a certain AWS organization.
	//
	// For more information about AWS Organizations, see [What Is AWS Organizations](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html) in the *AWS Organizations User Guide* .
	//
	// If you specify `Condition` with an AWS organization ID, and specify "*" as the value for `Principal` , you grant permission to all the accounts in the named organization.
	//
	// The `Condition` is a JSON string which must contain `Type` , `Key` , and `Value` fields.
	Condition interface{} `json:"condition" yaml:"condition"`
	// The name of the event bus associated with the rule.
	//
	// If you omit this, the default event bus is used.
	EventBusName *string `json:"eventBusName" yaml:"eventBusName"`
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus.
	//
	// Specify "*" to permit any account to put events to your default event bus.
	//
	// If you specify "*" without specifying `Condition` , avoid creating rules that may match undesirable events. To create more secure rules, make sure that the event pattern for each rule contains an `account` field with a specific account ID from which to receive events. Rules with an account field do not match any events sent from other accounts.
	Principal *string `json:"principal" yaml:"principal"`
	// A JSON string that describes the permission policy statement.
	//
	// You can include a `Policy` parameter in the request instead of using the `StatementId` , `Action` , `Principal` , or `Condition` parameters.
	Statement interface{} `json:"statement" yaml:"statement"`
}

// Properties for defining a `CfnEventBus`.
//
// TODO: EXAMPLE
//
type CfnEventBusProps struct {
	// The name of the new event bus.
	//
	// Event bus names cannot contain the / character. You can't use the name `default` for a custom event bus, as this name is already used for your account's default event bus.
	//
	// If this is a partner event bus, the name must exactly match the name of the partner event source that this event bus is matched to.
	Name *string `json:"name" yaml:"name"`
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	EventSourceName *string `json:"eventSourceName" yaml:"eventSourceName"`
	// `AWS::Events::EventBus.Tags`.
	Tags *[]*CfnEventBus_TagEntryProperty `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Events::Rule`.
//
// Creates or updates the specified rule. Rules are enabled by default, or based on value of the state. You can disable a rule using [DisableRule](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DisableRule.html) .
//
// A single rule watches for events from a single event bus. Events generated by AWS services go to your account's default event bus. Events generated by SaaS partner services or applications go to the matching partner event bus. If you have custom applications or services, you can specify whether their events go to your default event bus or a custom event bus that you have created. For more information, see [CreateEventBus](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateEventBus.html) .
//
// If you are updating an existing rule, the rule is replaced with what you specify in this `PutRule` command. If you omit arguments in `PutRule` , the old values for those arguments are not kept. Instead, they are replaced with null values.
//
// When you create or update a rule, incoming events might not immediately start matching to new or updated rules. Allow a short period of time for changes to take effect.
//
// A rule must contain at least an EventPattern or ScheduleExpression. Rules with EventPatterns are triggered when a matching event is observed. Rules with ScheduleExpressions self-trigger based on the given schedule. A rule can have both an EventPattern and a ScheduleExpression, in which case the rule triggers on matching events as well as on a schedule.
//
// Most services in AWS treat : or / as the same character in Amazon Resource Names (ARNs). However, EventBridge uses an exact match in event patterns and rules. Be sure to use the correct ARN characters when creating event patterns so that they match the ARN syntax in the event you want to match.
//
// In EventBridge, it is possible to create rules that lead to infinite loops, where a rule is fired repeatedly. For example, a rule might detect that ACLs have changed on an S3 bucket, and trigger software to change them to the desired state. If the rule is not written carefully, the subsequent change to the ACLs fires the rule again, creating an infinite loop.
//
// To prevent this, write the rules so that the triggered actions do not re-fire the same rule. For example, your rule could fire only if ACLs are found to be in a bad state, instead of after any change.
//
// An infinite loop can quickly cause higher than expected charges. We recommend that you use budgeting, which alerts you when charges exceed your specified limit. For more information, see [Managing Your Costs with Budgets](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/budgets-managing-costs.html) .
//
// TODO: EXAMPLE
//
type CfnRule interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	EventBusName() *string
	SetEventBusName(val *string)
	EventPattern() interface{}
	SetEventPattern(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	RoleArn() *string
	SetRoleArn(val *string)
	ScheduleExpression() *string
	SetScheduleExpression(val *string)
	Stack() awscdk.Stack
	State() *string
	SetState(val *string)
	Targets() interface{}
	SetTargets(val interface{})
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnRule
type jsiiProxy_CfnRule struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRule) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) EventPattern() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) ScheduleExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) Targets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRule) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Events::Rule`.
func NewCfnRule(scope constructs.Construct, id *string, props *CfnRuleProps) CfnRule {
	_init_.Initialize()

	j := jsiiProxy_CfnRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnRule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Events::Rule`.
func NewCfnRule_Override(c CfnRule, scope constructs.Construct, id *string, props *CfnRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.CfnRule",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRule) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetEventBusName(val *string) {
	_jsii_.Set(
		j,
		"eventBusName",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetEventPattern(val interface{}) {
	_jsii_.Set(
		j,
		"eventPattern",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetScheduleExpression(val *string) {
	_jsii_.Set(
		j,
		"scheduleExpression",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetState(val *string) {
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_CfnRule) SetTargets(val interface{}) {
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRule_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnRule",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRule_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnRule",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.CfnRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRule_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.CfnRule",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnRule) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRule) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRule) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnRule) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnRule) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnRule) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRule) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRule) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRule) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnRule) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRule) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRule) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRule) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// This structure specifies the VPC subnets and security groups for the task, and whether a public IP address is to be used.
//
// This structure is relevant only for ECS tasks that use the `awsvpc` network mode.
//
// TODO: EXAMPLE
//
type CfnRule_AwsVpcConfigurationProperty struct {
	// Specifies the subnets associated with the task.
	//
	// These subnets must all be in the same VPC. You can specify as many as 16 subnets.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// Specifies whether the task's elastic network interface receives a public IP address.
	//
	// You can specify `ENABLED` only when `LaunchType` in `EcsParameters` is set to `FARGATE` .
	AssignPublicIp *string `json:"assignPublicIp" yaml:"assignPublicIp"`
	// Specifies the security groups associated with the task.
	//
	// These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
	SecurityGroups *[]*string `json:"securityGroups" yaml:"securityGroups"`
}

// The array properties for the submitted job, such as the size of the array.
//
// The array size can be between 2 and 10,000. If you specify array properties for a job, it becomes an array job. This parameter is used only if the target is an AWS Batch job.
//
// TODO: EXAMPLE
//
type CfnRule_BatchArrayPropertiesProperty struct {
	// The size of the array, if this is an array batch job.
	//
	// Valid values are integers between 2 and 10,000.
	Size *float64 `json:"size" yaml:"size"`
}

// The custom parameters to be used when the target is an AWS Batch job.
//
// TODO: EXAMPLE
//
type CfnRule_BatchParametersProperty struct {
	// The ARN or name of the job definition to use if the event target is an AWS Batch job.
	//
	// This job definition must already exist.
	JobDefinition *string `json:"jobDefinition" yaml:"jobDefinition"`
	// The name to use for this execution of the job, if the target is an AWS Batch job.
	JobName *string `json:"jobName" yaml:"jobName"`
	// The array properties for the submitted job, such as the size of the array.
	//
	// The array size can be between 2 and 10,000. If you specify array properties for a job, it becomes an array job. This parameter is used only if the target is an AWS Batch job.
	ArrayProperties interface{} `json:"arrayProperties" yaml:"arrayProperties"`
	// The retry strategy to use for failed jobs, if the target is an AWS Batch job.
	//
	// The retry strategy is the number of times to retry the failed job execution. Valid values are 1–10. When you specify a retry strategy here, it overrides the retry strategy defined in the job definition.
	RetryStrategy interface{} `json:"retryStrategy" yaml:"retryStrategy"`
}

// The retry strategy to use for failed jobs, if the target is an AWS Batch job.
//
// If you specify a retry strategy here, it overrides the retry strategy defined in the job definition.
//
// TODO: EXAMPLE
//
type CfnRule_BatchRetryStrategyProperty struct {
	// The number of times to attempt to retry, if the job fails.
	//
	// Valid values are 1–10.
	Attempts *float64 `json:"attempts" yaml:"attempts"`
}

// The details of a capacity provider strategy.
//
// To learn more, see [CapacityProviderStrategyItem](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CapacityProviderStrategyItem.html) in the Amazon ECS API Reference.
//
// TODO: EXAMPLE
//
type CfnRule_CapacityProviderStrategyItemProperty struct {
	// The short name of the capacity provider.
	CapacityProvider *string `json:"capacityProvider" yaml:"capacityProvider"`
	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one capacity provider in a capacity provider strategy can have a base defined. If no value is specified, the default value of 0 is used.
	Base *float64 `json:"base" yaml:"base"`
	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The weight value is taken into consideration after the base value, if defined, is satisfied.
	Weight *float64 `json:"weight" yaml:"weight"`
}

// A `DeadLetterConfig` object that contains information about a dead-letter queue configuration.
//
// TODO: EXAMPLE
//
type CfnRule_DeadLetterConfigProperty struct {
	// The ARN of the SQS queue specified as the target for the dead-letter queue.
	Arn *string `json:"arn" yaml:"arn"`
}

// The custom parameters to be used when the target is an Amazon ECS task.
//
// TODO: EXAMPLE
//
type CfnRule_EcsParametersProperty struct {
	// The ARN of the task definition to use if the event target is an Amazon ECS task.
	TaskDefinitionArn *string `json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The capacity provider strategy to use for the task.
	//
	// If a `capacityProviderStrategy` is specified, the `launchType` parameter must be omitted. If no `capacityProviderStrategy` or launchType is specified, the `defaultCapacityProviderStrategy` for the cluster is used.
	CapacityProviderStrategy interface{} `json:"capacityProviderStrategy" yaml:"capacityProviderStrategy"`
	// Specifies whether to enable Amazon ECS managed tags for the task.
	//
	// For more information, see [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the Amazon Elastic Container Service Developer Guide.
	EnableEcsManagedTags interface{} `json:"enableEcsManagedTags" yaml:"enableEcsManagedTags"`
	// Whether or not to enable the execute command functionality for the containers in this task.
	//
	// If true, this enables execute command functionality on all containers in the task.
	EnableExecuteCommand interface{} `json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// Specifies an ECS task group for the task.
	//
	// The maximum length is 255 characters.
	Group *string `json:"group" yaml:"group"`
	// Specifies the launch type on which your task is running.
	//
	// The launch type that you specify here must match one of the launch type (compatibilities) of the target task. The `FARGATE` value is supported only in the Regions where AWS Fargate with Amazon ECS is supported. For more information, see [AWS Fargate on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/AWS-Fargate.html) in the *Amazon Elastic Container Service Developer Guide* .
	LaunchType *string `json:"launchType" yaml:"launchType"`
	// Use this structure if the Amazon ECS task uses the `awsvpc` network mode.
	//
	// This structure specifies the VPC subnets and security groups associated with the task, and whether a public IP address is to be used. This structure is required if `LaunchType` is `FARGATE` because the `awsvpc` mode is required for Fargate tasks.
	//
	// If you specify `NetworkConfiguration` when the target ECS task does not use the `awsvpc` network mode, the task fails.
	NetworkConfiguration interface{} `json:"networkConfiguration" yaml:"networkConfiguration"`
	// An array of placement constraint objects to use for the task.
	//
	// You can specify up to 10 constraints per task (including constraints in the task definition and those specified at runtime).
	PlacementConstraints interface{} `json:"placementConstraints" yaml:"placementConstraints"`
	// The placement strategy objects to use for the task.
	//
	// You can specify a maximum of five strategy rules per task.
	PlacementStrategies interface{} `json:"placementStrategies" yaml:"placementStrategies"`
	// Specifies the platform version for the task.
	//
	// Specify only the numeric portion of the platform version, such as `1.1.0` .
	//
	// This structure is used only if `LaunchType` is `FARGATE` . For more information about valid platform versions, see [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide* .
	PlatformVersion *string `json:"platformVersion" yaml:"platformVersion"`
	// Specifies whether to propagate the tags from the task definition to the task.
	//
	// If no value is specified, the tags are not propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the TagResource API action.
	PropagateTags *string `json:"propagateTags" yaml:"propagateTags"`
	// The reference ID to use for the task.
	ReferenceId *string `json:"referenceId" yaml:"referenceId"`
	// The metadata that you apply to the task to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define. To learn more, see [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html#ECS-RunTask-request-tags) in the Amazon ECS API Reference.
	TagList interface{} `json:"tagList" yaml:"tagList"`
	// The number of tasks to create based on `TaskDefinition` .
	//
	// The default is 1.
	TaskCount *float64 `json:"taskCount" yaml:"taskCount"`
}

// These are custom parameter to be used when the target is an API Gateway REST APIs or EventBridge ApiDestinations.
//
// In the latter case, these are merged with any InvocationParameters specified on the Connection, with any values from the Connection taking precedence.
//
// TODO: EXAMPLE
//
type CfnRule_HttpParametersProperty struct {
	// The headers that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
	HeaderParameters interface{} `json:"headerParameters" yaml:"headerParameters"`
	// The path parameter values to be used to populate API Gateway REST API or EventBridge ApiDestination path wildcards ("*").
	PathParameterValues *[]*string `json:"pathParameterValues" yaml:"pathParameterValues"`
	// The query string keys/values that need to be sent as part of request invoking the API Gateway REST API or EventBridge ApiDestination.
	QueryStringParameters interface{} `json:"queryStringParameters" yaml:"queryStringParameters"`
}

// Contains the parameters needed for you to provide custom input to a target based on one or more pieces of data extracted from the event.
//
// TODO: EXAMPLE
//
type CfnRule_InputTransformerProperty struct {
	// Input template where you specify placeholders that will be filled with the values of the keys from `InputPathsMap` to customize the data sent to the target.
	//
	// Enclose each `InputPathsMaps` value in brackets: < *value* > The InputTemplate must be valid JSON.
	//
	// If `InputTemplate` is a JSON object (surrounded by curly braces), the following restrictions apply:
	//
	// - The placeholder cannot be used as an object key.
	//
	// The following example shows the syntax for using `InputPathsMap` and `InputTemplate` .
	//
	// `"InputTransformer":`
	//
	// `{`
	//
	// `"InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},`
	//
	// `"InputTemplate": "<instance> is in state <status>"`
	//
	// `}`
	//
	// To have the `InputTemplate` include quote marks within a JSON string, escape each quote marks with a slash, as in the following example:
	//
	// `"InputTransformer":`
	//
	// `{`
	//
	// `"InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},`
	//
	// `"InputTemplate": "<instance> is in state \"<status>\""`
	//
	// `}`
	//
	// The `InputTemplate` can also be valid JSON with varibles in quotes or out, as in the following example:
	//
	// `"InputTransformer":`
	//
	// `{`
	//
	// `"InputPathsMap": {"instance": "$.detail.instance","status": "$.detail.status"},`
	//
	// `"InputTemplate": '{"myInstance": <instance>,"myStatus": "<instance> is in state \"<status>\""}'`
	//
	// `}`
	InputTemplate *string `json:"inputTemplate" yaml:"inputTemplate"`
	// Map of JSON paths to be extracted from the event.
	//
	// You can then insert these in the template in `InputTemplate` to produce the output you want to be sent to the target.
	//
	// `InputPathsMap` is an array key-value pairs, where each value is a valid JSON path. You can have as many as 100 key-value pairs. You must use JSON dot notation, not bracket notation.
	//
	// The keys cannot start with " AWS ."
	InputPathsMap interface{} `json:"inputPathsMap" yaml:"inputPathsMap"`
}

// This object enables you to specify a JSON path to extract from the event and use as the partition key for the Amazon Kinesis data stream, so that you can control the shard to which the event goes.
//
// If you do not include this parameter, the default is to use the `eventId` as the partition key.
//
// TODO: EXAMPLE
//
type CfnRule_KinesisParametersProperty struct {
	// The JSON path to be extracted from the event and used as the partition key.
	//
	// For more information, see [Amazon Kinesis Streams Key Concepts](https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html#partition-key) in the *Amazon Kinesis Streams Developer Guide* .
	PartitionKeyPath *string `json:"partitionKeyPath" yaml:"partitionKeyPath"`
}

// This structure specifies the network configuration for an ECS task.
//
// TODO: EXAMPLE
//
type CfnRule_NetworkConfigurationProperty struct {
	// Use this structure to specify the VPC subnets and security groups for the task, and whether a public IP address is to be used.
	//
	// This structure is relevant only for ECS tasks that use the `awsvpc` network mode.
	AwsVpcConfiguration interface{} `json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

// An object representing a constraint on task placement.
//
// To learn more, see [Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html) in the Amazon Elastic Container Service Developer Guide.
//
// TODO: EXAMPLE
//
type CfnRule_PlacementConstraintProperty struct {
	// A cluster query language expression to apply to the constraint.
	//
	// You cannot specify an expression if the constraint type is `distinctInstance` . To learn more, see [Cluster Query Language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the Amazon Elastic Container Service Developer Guide.
	Expression *string `json:"expression" yaml:"expression"`
	// The type of constraint.
	//
	// Use distinctInstance to ensure that each task in a particular group is running on a different container instance. Use memberOf to restrict the selection to a group of valid candidates.
	Type *string `json:"type" yaml:"type"`
}

// The task placement strategy for a task or service.
//
// To learn more, see [Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html) in the Amazon Elastic Container Service Service Developer Guide.
//
// TODO: EXAMPLE
//
type CfnRule_PlacementStrategyProperty struct {
	// The field to apply the placement strategy against.
	//
	// For the spread placement strategy, valid values are instanceId (or host, which has the same effect), or any platform or custom attribute that is applied to a container instance, such as attribute:ecs.availability-zone. For the binpack placement strategy, valid values are cpu and memory. For the random placement strategy, this field is not used.
	Field *string `json:"field" yaml:"field"`
	// The type of placement strategy.
	//
	// The random placement strategy randomly places tasks on available candidates. The spread placement strategy spreads placement across available candidates evenly based on the field parameter. The binpack strategy places tasks on available candidates that have the least available amount of the resource that is specified with the field parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory (but still enough to run the task).
	Type *string `json:"type" yaml:"type"`
}

// These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the Amazon Redshift Data API ExecuteStatement based on EventBridge events.
//
// TODO: EXAMPLE
//
type CfnRule_RedshiftDataParametersProperty struct {
	// The name of the database.
	//
	// Required when authenticating using temporary credentials.
	Database *string `json:"database" yaml:"database"`
	// The SQL statement text to run.
	Sql *string `json:"sql" yaml:"sql"`
	// The database user name.
	//
	// Required when authenticating using temporary credentials.
	DbUser *string `json:"dbUser" yaml:"dbUser"`
	// The name or ARN of the secret that enables access to the database.
	//
	// Required when authenticating using AWS Secrets Manager.
	SecretManagerArn *string `json:"secretManagerArn" yaml:"secretManagerArn"`
	// The name of the SQL statement.
	//
	// You can name the SQL statement when you create it to identify the query.
	StatementName *string `json:"statementName" yaml:"statementName"`
	// Indicates whether to send an event back to EventBridge after the SQL statement runs.
	WithEvent interface{} `json:"withEvent" yaml:"withEvent"`
}

// A `RetryPolicy` object that includes information about the retry policy settings.
//
// TODO: EXAMPLE
//
type CfnRule_RetryPolicyProperty struct {
	// The maximum amount of time, in seconds, to continue to make retry attempts.
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// The maximum number of retry attempts to make before the request fails.
	//
	// Retry attempts continue until either the maximum number of attempts is made or until the duration of the `MaximumEventAgeInSeconds` is met.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

// This parameter contains the criteria (either InstanceIds or a tag) used to specify which EC2 instances are to be sent the command.
//
// TODO: EXAMPLE
//
type CfnRule_RunCommandParametersProperty struct {
	// Currently, we support including only one RunCommandTarget block, which specifies either an array of InstanceIds or a tag.
	RunCommandTargets interface{} `json:"runCommandTargets" yaml:"runCommandTargets"`
}

// Information about the EC2 instances that are to be sent the command, specified as key-value pairs.
//
// Each `RunCommandTarget` block can include only one key, but this key may specify multiple values.
//
// TODO: EXAMPLE
//
type CfnRule_RunCommandTargetProperty struct {
	// Can be either `tag:` *tag-key* or `InstanceIds` .
	Key *string `json:"key" yaml:"key"`
	// If `Key` is `tag:` *tag-key* , `Values` is a list of tag values.
	//
	// If `Key` is `InstanceIds` , `Values` is a list of Amazon EC2 instance IDs.
	Values *[]*string `json:"values" yaml:"values"`
}

// Name/Value pair of a parameter to start execution of a SageMaker Model Building Pipeline.
//
// TODO: EXAMPLE
//
type CfnRule_SageMakerPipelineParameterProperty struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	Name *string `json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	Value *string `json:"value" yaml:"value"`
}

// These are custom parameters to use when the target is a SageMaker Model Building Pipeline that starts based on EventBridge events.
//
// TODO: EXAMPLE
//
type CfnRule_SageMakerPipelineParametersProperty struct {
	// List of Parameter names and values for SageMaker Model Building Pipeline execution.
	PipelineParameterList interface{} `json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

// This structure includes the custom parameter to be used when the target is an SQS FIFO queue.
//
// TODO: EXAMPLE
//
type CfnRule_SqsParametersProperty struct {
	// The FIFO message group ID to use as the target.
	MessageGroupId *string `json:"messageGroupId" yaml:"messageGroupId"`
}

// A key-value pair associated with an ECS Target of an EventBridge rule.
//
// The tag will be propagated to ECS by EventBridge when starting an ECS task based on a matched event.
//
// > Currently, tags are only available when using ECS with EventBridge .
//
// TODO: EXAMPLE
//
type CfnRule_TagProperty struct {
	// A string you can use to assign a value.
	//
	// The combination of tag keys and values can help you organize and categorize your resources.
	Key *string `json:"key" yaml:"key"`
	// The value for the specified tag key.
	Value *string `json:"value" yaml:"value"`
}

// Targets are the resources to be invoked when a rule is triggered.
//
// For a complete list of services and resources that can be set as a target, see [PutTargets](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutTargets.html) .
//
// If you are setting the event bus of another account as the target, and that account granted permission to your account through an organization instead of directly by the account ID, then you must specify a `RoleArn` with proper permissions in the `Target` structure. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .
//
// TODO: EXAMPLE
//
type CfnRule_TargetProperty struct {
	// The Amazon Resource Name (ARN) of the target.
	Arn *string `json:"arn" yaml:"arn"`
	// The ID of the target within the specified rule.
	//
	// Use this ID to reference the target when updating the rule. We recommend using a memorable and unique string.
	Id *string `json:"id" yaml:"id"`
	// If the event target is an AWS Batch job, this contains the job definition, job name, and other parameters.
	//
	// For more information, see [Jobs](https://docs.aws.amazon.com/batch/latest/userguide/jobs.html) in the *AWS Batch User Guide* .
	BatchParameters interface{} `json:"batchParameters" yaml:"batchParameters"`
	// The `DeadLetterConfig` that defines the target queue to send dead-letter queue events to.
	DeadLetterConfig interface{} `json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Contains the Amazon ECS task definition and task count to be used, if the event target is an Amazon ECS task.
	//
	// For more information about Amazon ECS tasks, see [Task Definitions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html) in the *Amazon EC2 Container Service Developer Guide* .
	EcsParameters interface{} `json:"ecsParameters" yaml:"ecsParameters"`
	// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge ApiDestination.
	//
	// If you specify an API Gateway REST API or EventBridge ApiDestination as a target, you can use this parameter to specify headers, path parameters, and query string keys/values as part of your target invoking request. If you're using ApiDestinations, the corresponding Connection can also have these values configured. In case of any conflicting keys, values from the Connection take precedence.
	HttpParameters interface{} `json:"httpParameters" yaml:"httpParameters"`
	// Valid JSON text passed to the target.
	//
	// In this case, nothing from the event itself is passed to the target. For more information, see [The JavaScript Object Notation (JSON) Data Interchange Format](https://docs.aws.amazon.com/http://www.rfc-editor.org/rfc/rfc7159.txt) .
	Input *string `json:"input" yaml:"input"`
	// The value of the JSONPath that is used for extracting part of the matched event when passing it to the target.
	//
	// You must use JSON dot notation, not bracket notation. For more information about JSON paths, see [JSONPath](https://docs.aws.amazon.com/http://goessner.net/articles/JsonPath/) .
	InputPath *string `json:"inputPath" yaml:"inputPath"`
	// Settings to enable you to provide custom input to a target based on certain event data.
	//
	// You can extract one or more key-value pairs from the event and then use that data to send customized input to the target.
	InputTransformer interface{} `json:"inputTransformer" yaml:"inputTransformer"`
	// The custom parameter you can use to control the shard assignment, when the target is a Kinesis data stream.
	//
	// If you do not include this parameter, the default is to use the `eventId` as the partition key.
	KinesisParameters interface{} `json:"kinesisParameters" yaml:"kinesisParameters"`
	// Contains the Amazon Redshift Data API parameters to use when the target is a Amazon Redshift cluster.
	//
	// If you specify a Amazon Redshift Cluster as a Target, you can use this to specify parameters to invoke the Amazon Redshift Data API ExecuteStatement based on EventBridge events.
	RedshiftDataParameters interface{} `json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// The `RetryPolicy` object that contains the retry policy configuration to use for the dead-letter queue.
	RetryPolicy interface{} `json:"retryPolicy" yaml:"retryPolicy"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered.
	//
	// If one rule triggers multiple targets, you can use a different IAM role for each target.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// Parameters used when you are using the rule to invoke Amazon EC2 Run Command.
	RunCommandParameters interface{} `json:"runCommandParameters" yaml:"runCommandParameters"`
	// Contains the SageMaker Model Building Pipeline parameters to start execution of a SageMaker Model Building Pipeline.
	//
	// If you specify a SageMaker Model Building Pipeline as a target, you can use this to specify parameters to start a pipeline execution based on EventBridge events.
	SageMakerPipelineParameters interface{} `json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// Contains the message group ID to use when the target is a FIFO queue.
	//
	// If you specify an SQS FIFO queue as a target, the queue must have content-based deduplication enabled.
	SqsParameters interface{} `json:"sqsParameters" yaml:"sqsParameters"`
}

// Properties for defining a `CfnRule`.
//
// TODO: EXAMPLE
//
type CfnRuleProps struct {
	// The description of the rule.
	Description *string `json:"description" yaml:"description"`
	// The name or ARN of the event bus associated with the rule.
	//
	// If you omit this, the default event bus is used.
	EventBusName *string `json:"eventBusName" yaml:"eventBusName"`
	// The event pattern of the rule.
	//
	// For more information, see [Events and Event Patterns](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html) in the *Amazon EventBridge User Guide* .
	EventPattern interface{} `json:"eventPattern" yaml:"eventPattern"`
	// The name of the rule.
	Name *string `json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role that is used for target invocation.
	//
	// If you're setting an event bus in another account as the target and that account granted permission to your account through an organization instead of directly by the account ID, you must specify a `RoleArn` with proper permissions in the `Target` structure, instead of here in this parameter.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The scheduling expression.
	//
	// For example, "cron(0 20 * * ? *)", "rate(5 minutes)". For more information, see [Creating an Amazon EventBridge rule that runs on a schedule](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-create-rule-schedule.html) .
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
	// The state of the rule.
	State *string `json:"state" yaml:"state"`
	// Adds the specified targets to the specified rule, or updates the targets if they are already associated with the rule.
	//
	// Targets are the resources that are invoked when a rule is triggered.
	//
	// > Each rule can have up to five (5) targets associated with it at one time.
	//
	// You can configure the following as targets for Events:
	//
	// - [API destination](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-destinations.html)
	// - [API Gateway](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-gateway-target.html)
	// - Batch job queue
	// - CloudWatch group
	// - CodeBuild project
	// - CodePipeline
	// - EC2 `CreateSnapshot` API call
	// - EC2 Image Builder
	// - EC2 `RebootInstances` API call
	// - EC2 `StopInstances` API call
	// - EC2 `TerminateInstances` API call
	// - ECS task
	// - [Event bus in a different account or Region](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-cross-account.html)
	// - [Event bus in the same account and Region](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-bus-to-bus.html)
	// - Firehose delivery stream
	// - Glue workflow
	// - [Incident Manager response plan](https://docs.aws.amazon.com//incident-manager/latest/userguide/incident-creation.html#incident-tracking-auto-eventbridge)
	// - Inspector assessment template
	// - Kinesis stream
	// - Lambda function
	// - Redshift cluster
	// - SageMaker Pipeline
	// - SNS topic
	// - SQS queue
	// - Step Functions state machine
	// - Systems Manager Automation
	// - Systems Manager OpsItem
	// - Systems Manager Run Command
	//
	// Creating rules with built-in targets is supported only in the AWS Management Console . The built-in targets are `EC2 CreateSnapshot API call` , `EC2 RebootInstances API call` , `EC2 StopInstances API call` , and `EC2 TerminateInstances API call` .
	//
	// For some target types, `PutTargets` provides target-specific parameters. If the target is a Kinesis data stream, you can optionally specify which shard the event goes to by using the `KinesisParameters` argument. To invoke a command on multiple EC2 instances with one rule, you can use the `RunCommandParameters` field.
	//
	// To be able to make API calls against the resources that you own, Amazon EventBridge needs the appropriate permissions. For AWS Lambda and Amazon SNS resources, EventBridge relies on resource-based policies. For EC2 instances, Kinesis Data Streams, AWS Step Functions state machines and API Gateway REST APIs, EventBridge relies on IAM roles that you specify in the `RoleARN` argument in `PutTargets` . For more information, see [Authentication and Access Control](https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html) in the *Amazon EventBridge User Guide* .
	//
	// If another AWS account is in the same region and has granted you permission (using `PutPermission` ), you can send events to that account. Set that account's event bus as a target of the rules in your account. To send the matched events to the other account, specify that account's event bus as the `Arn` value when you run `PutTargets` . If your account sends events to another account, your account is charged for each sent event. Each event sent to another account is charged as a custom event. The account receiving the event is not charged. For more information, see [Amazon EventBridge Pricing](https://docs.aws.amazon.com/eventbridge/pricing/) .
	//
	// > `Input` , `InputPath` , and `InputTransformer` are not available with `PutTarget` if the target is an event bus of a different AWS account.
	//
	// If you are setting the event bus of another account as the target, and that account granted permission to your account through an organization instead of directly by the account ID, then you must specify a `RoleArn` with proper permissions in the `Target` structure. For more information, see [Sending and Receiving Events Between AWS Accounts](https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html) in the *Amazon EventBridge User Guide* .
	//
	// For more information about enabling cross-account events, see [PutPermission](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html) .
	//
	// *Input* , *InputPath* , and *InputTransformer* are mutually exclusive and optional parameters of a target. When a rule is triggered due to a matched event:
	//
	// - If none of the following arguments are specified for a target, then the entire event is passed to the target in JSON format (unless the target is Amazon EC2 Run Command or Amazon ECS task, in which case nothing from the event is passed to the target).
	// - If *Input* is specified in the form of valid JSON, then the matched event is overridden with this constant.
	// - If *InputPath* is specified in the form of JSONPath (for example, `$.detail` ), then only the part of the event specified in the path is passed to the target (for example, only the detail part of the event is passed).
	// - If *InputTransformer* is specified, then one or more specified JSONPaths are extracted from the event and used as values in a template that you specify as the input to the target.
	//
	// When you specify `InputPath` or `InputTransformer` , you must use JSON dot notation, not bracket notation.
	//
	// When you add targets to a rule and the associated rule triggers soon after, new or updated targets might not be immediately invoked. Allow a short period of time for changes to take effect.
	//
	// This action can partially fail if too many requests are made at the same time. If that happens, `FailedEntryCount` is non-zero in the response and each entry in `FailedEntries` provides the ID of the failed target and the error code.
	Targets interface{} `json:"targets" yaml:"targets"`
}

// Define an EventBridge Connection.
//
// TODO: EXAMPLE
//
type Connection interface {
	awscdk.Resource
	IConnection
	ConnectionArn() *string
	ConnectionName() *string
	ConnectionSecretArn() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Connection
type jsiiProxy_Connection struct {
	internal.Type__awscdkResource
	jsiiProxy_IConnection
}

func (j *jsiiProxy_Connection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) ConnectionSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Connection) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewConnection(scope constructs.Construct, id *string, props *ConnectionProps) Connection {
	_init_.Initialize()

	j := jsiiProxy_Connection{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Connection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewConnection_Override(c Connection, scope constructs.Construct, id *string, props *ConnectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Connection",
		[]interface{}{scope, id, props},
		c,
	)
}

// Import an existing connection resource.
func Connection_FromConnectionAttributes(scope constructs.Construct, id *string, attrs *ConnectionAttributes) IConnection {
	_init_.Initialize()

	var returns IConnection

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Connection",
		"fromConnectionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing connection resource.
func Connection_FromEventBusArn(scope constructs.Construct, id *string, connectionArn *string, connectionSecretArn *string) IConnection {
	_init_.Initialize()

	var returns IConnection

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Connection",
		"fromEventBusArn",
		[]interface{}{scope, id, connectionArn, connectionSecretArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Connection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Connection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Connection_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Connection",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_Connection) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_Connection) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (c *jsiiProxy_Connection) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (c *jsiiProxy_Connection) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (c *jsiiProxy_Connection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface with properties necessary to import a reusable Connection.
//
// TODO: EXAMPLE
//
type ConnectionAttributes struct {
	// The ARN of the connection created.
	ConnectionArn *string `json:"connectionArn" yaml:"connectionArn"`
	// The Name for the connection.
	ConnectionName *string `json:"connectionName" yaml:"connectionName"`
	// The ARN for the secret created for the connection.
	ConnectionSecretArn *string `json:"connectionSecretArn" yaml:"connectionSecretArn"`
}

// An API Destination Connection.
//
// A connection defines the authorization type and credentials to use for authorization with an API destination HTTP endpoint.
//
// TODO: EXAMPLE
//
type ConnectionProps struct {
	// The authorization type for the connection.
	Authorization Authorization `json:"authorization" yaml:"authorization"`
	// Additional string parameters to add to the invocation bodies.
	BodyParameters *map[string]HttpParameter `json:"bodyParameters" yaml:"bodyParameters"`
	// The name of the connection.
	ConnectionName *string `json:"connectionName" yaml:"connectionName"`
	// The name of the connection.
	Description *string `json:"description" yaml:"description"`
	// Additional string parameters to add to the invocation headers.
	HeaderParameters *map[string]HttpParameter `json:"headerParameters" yaml:"headerParameters"`
	// Additional string parameters to add to the invocation query strings.
	QueryStringParameters *map[string]HttpParameter `json:"queryStringParameters" yaml:"queryStringParameters"`
}

// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html#cron-expressions
//
type CronOptions struct {
	// The day of the month to run this rule at.
	Day *string `json:"day" yaml:"day"`
	// The hour to run this rule at.
	Hour *string `json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	Minute *string `json:"minute" yaml:"minute"`
	// The month to run this rule at.
	Month *string `json:"month" yaml:"month"`
	// The day of the week to run this rule at.
	WeekDay *string `json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	Year *string `json:"year" yaml:"year"`
}

// Define an EventBridge EventBus.
//
// TODO: EXAMPLE
//
type EventBus interface {
	awscdk.Resource
	IEventBus
	Env() *awscdk.ResourceEnvironment
	EventBusArn() *string
	EventBusName() *string
	EventBusPolicy() *string
	EventSourceName() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	Archive(id *string, props *BaseArchiveProps) Archive
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantPutEventsTo(grantee awsiam.IGrantable) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for EventBus
type jsiiProxy_EventBus struct {
	internal.Type__awscdkResource
	jsiiProxy_IEventBus
}

func (j *jsiiProxy_EventBus) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBus) EventBusArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBus) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBus) EventBusPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBus) EventSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBus) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBus) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBus) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewEventBus(scope constructs.Construct, id *string, props *EventBusProps) EventBus {
	_init_.Initialize()

	j := jsiiProxy_EventBus{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.EventBus",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventBus_Override(e EventBus, scope constructs.Construct, id *string, props *EventBusProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.EventBus",
		[]interface{}{scope, id, props},
		e,
	)
}

// Import an existing event bus resource.
func EventBus_FromEventBusArn(scope constructs.Construct, id *string, eventBusArn *string) IEventBus {
	_init_.Initialize()

	var returns IEventBus

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventBus",
		"fromEventBusArn",
		[]interface{}{scope, id, eventBusArn},
		&returns,
	)

	return returns
}

// Import an existing event bus resource.
func EventBus_FromEventBusAttributes(scope constructs.Construct, id *string, attrs *EventBusAttributes) IEventBus {
	_init_.Initialize()

	var returns IEventBus

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventBus",
		"fromEventBusAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing event bus resource.
func EventBus_FromEventBusName(scope constructs.Construct, id *string, eventBusName *string) IEventBus {
	_init_.Initialize()

	var returns IEventBus

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventBus",
		"fromEventBusName",
		[]interface{}{scope, id, eventBusName},
		&returns,
	)

	return returns
}

// Permits an IAM Principal to send custom events to EventBridge so that they can be matched to rules.
func EventBus_GrantAllPutEvents(grantee awsiam.IGrantable) awsiam.Grant {
	_init_.Initialize()

	var returns awsiam.Grant

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventBus",
		"grantAllPutEvents",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventBus_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventBus",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func EventBus_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventBus",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (e *jsiiProxy_EventBus) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Create an EventBridge archive to send events to.
//
// When you create an archive, incoming events might not immediately start being sent to the archive.
// Allow a short period of time for changes to take effect.
func (e *jsiiProxy_EventBus) Archive(id *string, props *BaseArchiveProps) Archive {
	var returns Archive

	_jsii_.Invoke(
		e,
		"archive",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBus) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (e *jsiiProxy_EventBus) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (e *jsiiProxy_EventBus) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grants an IAM Principal to send custom events to the eventBus so that they can be matched to rules.
func (e *jsiiProxy_EventBus) GrantPutEventsTo(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		e,
		"grantPutEventsTo",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventBus) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface with properties necessary to import a reusable EventBus.
//
// TODO: EXAMPLE
//
type EventBusAttributes struct {
	// The ARN of this event bus resource.
	EventBusArn *string `json:"eventBusArn" yaml:"eventBusArn"`
	// The physical ID of this event bus resource.
	EventBusName *string `json:"eventBusName" yaml:"eventBusName"`
	// The JSON policy of this event bus resource.
	EventBusPolicy *string `json:"eventBusPolicy" yaml:"eventBusPolicy"`
	// The partner event source to associate with this event bus resource.
	EventSourceName *string `json:"eventSourceName" yaml:"eventSourceName"`
}

// Properties to define an event bus.
//
// TODO: EXAMPLE
//
type EventBusProps struct {
	// The name of the event bus you are creating Note: If 'eventSourceName' is passed in, you cannot set this.
	EventBusName *string `json:"eventBusName" yaml:"eventBusName"`
	// The partner event source to associate with this event bus resource Note: If 'eventBusName' is passed in, you cannot set this.
	EventSourceName *string `json:"eventSourceName" yaml:"eventSourceName"`
}

// Represents a field in the event pattern.
type EventField interface {
	awscdk.IResolvable
	CreationStack() *[]*string
	DisplayHint() *string
	Path() *string
	Resolve(_ctx awscdk.IResolveContext) interface{}
	ToJSON() *string
	ToString() *string
}

// The jsii proxy struct for EventField
type jsiiProxy_EventField struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_EventField) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventField) DisplayHint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayHint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventField) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


// Extract a custom JSON path from the event.
func EventField_FromPath(path *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.EventField",
		"fromPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func EventField_Account() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"account",
		&returns,
	)
	return returns
}

func EventField_DetailType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"detailType",
		&returns,
	)
	return returns
}

func EventField_EventId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"eventId",
		&returns,
	)
	return returns
}

func EventField_Region() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"region",
		&returns,
	)
	return returns
}

func EventField_Source() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"source",
		&returns,
	)
	return returns
}

func EventField_Time() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_events.EventField",
		"time",
		&returns,
	)
	return returns
}

// Produce the Token's value at resolution time.
func (e *jsiiProxy_EventField) Resolve(_ctx awscdk.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_ctx},
		&returns,
	)

	return returns
}

// Convert the path to the field in the event pattern to JSON.
func (e *jsiiProxy_EventField) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Return a string representation of this resolvable object.
//
// Returns a reversible string representation.
func (e *jsiiProxy_EventField) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Events in Amazon CloudWatch Events are represented as JSON objects. For more information about JSON objects, see RFC 7159.
//
// **Important**: this class can only be used with a `Rule` class. In particular,
// do not use it with `CfnRule` class: your pattern will not be rendered
// correctly. In a `CfnRule` class, write the pattern as you normally would when
// directly writing CloudFormation.
//
// Rules use event patterns to select events and route them to targets. A
// pattern either matches an event or it doesn't. Event patterns are represented
// as JSON objects with a structure that is similar to that of events.
//
// It is important to remember the following about event pattern matching:
//
// - For a pattern to match an event, the event must contain all the field names
//    listed in the pattern. The field names must appear in the event with the
//    same nesting structure.
//
// - Other fields of the event not mentioned in the pattern are ignored;
//    effectively, there is a ``"*": "*"`` wildcard for fields not mentioned.
//
// - The matching is exact (character-by-character), without case-folding or any
//    other string normalization.
//
// - The values being matched follow JSON rules: Strings enclosed in quotes,
//    numbers, and the unquoted keywords true, false, and null.
//
// - Number matching is at the string representation level. For example, 300,
//    300.0, and 3.0e2 are not considered equal.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/CloudWatchEventsandEventPatterns.html
//
type EventPattern struct {
	// The 12-digit number identifying an AWS account.
	Account *[]*string `json:"account" yaml:"account"`
	// A JSON object, whose content is at the discretion of the service originating the event.
	Detail *map[string]interface{} `json:"detail" yaml:"detail"`
	// Identifies, in combination with the source field, the fields and values that appear in the detail field.
	//
	// Represents the "detail-type" event field.
	DetailType *[]*string `json:"detailType" yaml:"detailType"`
	// A unique value is generated for every event.
	//
	// This can be helpful in
	// tracing events as they move through rules to targets, and are processed.
	Id *[]*string `json:"id" yaml:"id"`
	// Identifies the AWS region where the event originated.
	Region *[]*string `json:"region" yaml:"region"`
	// This JSON array contains ARNs that identify resources that are involved in the event.
	//
	// Inclusion of these ARNs is at the discretion of the
	// service.
	//
	// For example, Amazon EC2 instance state-changes include Amazon EC2
	// instance ARNs, Auto Scaling events include ARNs for both instances and
	// Auto Scaling groups, but API calls with AWS CloudTrail do not include
	// resource ARNs.
	Resources *[]*string `json:"resources" yaml:"resources"`
	// Identifies the service that sourced the event.
	//
	// All events sourced from
	// within AWS begin with "aws." Customer-generated events can have any value
	// here, as long as it doesn't begin with "aws." We recommend the use of
	// Java package-name style reverse domain-name strings.
	//
	// To find the correct value for source for an AWS service, see the table in
	// AWS Service Namespaces. For example, the source value for Amazon
	// CloudFront is aws.cloudfront.
	// See: http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces
	//
	Source *[]*string `json:"source" yaml:"source"`
	// The event timestamp, which can be specified by the service originating the event.
	//
	// If the event spans a time interval, the service might choose
	// to report the start time, so this value can be noticeably before the time
	// the event is actually received.
	Time *[]*string `json:"time" yaml:"time"`
	// By default, this is set to 0 (zero) in all events.
	Version *[]*string `json:"version" yaml:"version"`
}

// Supported HTTP operations.
type HttpMethod string

const (
	HttpMethod_POST HttpMethod = "POST"
	HttpMethod_GET HttpMethod = "GET"
	HttpMethod_HEAD HttpMethod = "HEAD"
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	HttpMethod_PUT HttpMethod = "PUT"
	HttpMethod_PATCH HttpMethod = "PATCH"
	HttpMethod_DELETE HttpMethod = "DELETE"
)

// An additional HTTP parameter to send along with the OAuth request.
//
// TODO: EXAMPLE
//
type HttpParameter interface {
}

// The jsii proxy struct for HttpParameter
type jsiiProxy_HttpParameter struct {
	_ byte // padding
}

func NewHttpParameter_Override(h HttpParameter) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.HttpParameter",
		nil, // no parameters
		h,
	)
}

// Make an OAuthParameter from a secret.
func HttpParameter_FromSecret(value awscdk.SecretValue) HttpParameter {
	_init_.Initialize()

	var returns HttpParameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.HttpParameter",
		"fromSecret",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Make an OAuthParameter from a string value.
//
// The value is not treated as a secret.
func HttpParameter_FromString(value *string) HttpParameter {
	_init_.Initialize()

	var returns HttpParameter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.HttpParameter",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Interface for API Destinations.
type IApiDestination interface {
	awscdk.IResource
	// The ARN of the Api Destination created.
	ApiDestinationArn() *string
	// The Name of the Api Destination created.
	ApiDestinationName() *string
}

// The jsii proxy for IApiDestination
type jsiiProxy_IApiDestination struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApiDestination) ApiDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiDestination) ApiDestinationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDestinationName",
		&returns,
	)
	return returns
}

// Interface for EventBus Connections.
type IConnection interface {
	awscdk.IResource
	// The ARN of the connection created.
	ConnectionArn() *string
	// The Name for the connection.
	ConnectionName() *string
	// The ARN for the secret created for the connection.
	ConnectionSecretArn() *string
}

// The jsii proxy for IConnection
type jsiiProxy_IConnection struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IConnection) ConnectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) ConnectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnection) ConnectionSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionSecretArn",
		&returns,
	)
	return returns
}

// Interface which all EventBus based classes MUST implement.
type IEventBus interface {
	awscdk.IResource
	// Create an EventBridge archive to send events to.
	//
	// When you create an archive, incoming events might not immediately start being sent to the archive.
	// Allow a short period of time for changes to take effect.
	Archive(id *string, props *BaseArchiveProps) Archive
	// Grants an IAM Principal to send custom events to the eventBus so that they can be matched to rules.
	GrantPutEventsTo(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of this event bus resource.
	EventBusArn() *string
	// The physical ID of this event bus resource.
	EventBusName() *string
	// The JSON policy of this event bus resource.
	EventBusPolicy() *string
	// The partner event source to associate with this event bus resource.
	EventSourceName() *string
}

// The jsii proxy for IEventBus
type jsiiProxy_IEventBus struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IEventBus) Archive(id *string, props *BaseArchiveProps) Archive {
	var returns Archive

	_jsii_.Invoke(
		i,
		"archive",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEventBus) GrantPutEventsTo(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantPutEventsTo",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IEventBus) EventBusArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBus) EventBusName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBus) EventBusPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBusPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBus) EventSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceName",
		&returns,
	)
	return returns
}

// Represents an EventBridge Rule.
type IRule interface {
	awscdk.IResource
	// The value of the event rule Amazon Resource Name (ARN), such as arn:aws:events:us-east-2:123456789012:rule/example.
	RuleArn() *string
	// The name event rule.
	RuleName() *string
}

// The jsii proxy for IRule
type jsiiProxy_IRule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRule) RuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

// An abstract target for EventRules.
type IRuleTarget interface {
	// Returns the rule target specification.
	//
	// NOTE: Do not use the various `inputXxx` options. They can be set in a call to `addTarget`.
	Bind(rule IRule, id *string) *RuleTargetConfig
}

// The jsii proxy for IRuleTarget
type jsiiProxy_IRuleTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IRuleTarget) Bind(rule IRule, id *string) *RuleTargetConfig {
	var returns *RuleTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

// Properties for `Authorization.oauth()`.
//
// TODO: EXAMPLE
//
type OAuthAuthorizationProps struct {
	// The URL to the authorization endpoint.
	AuthorizationEndpoint *string `json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The client ID to use for OAuth authorization for the connection.
	ClientId *string `json:"clientId" yaml:"clientId"`
	// The client secret associated with the client ID to use for OAuth authorization for the connection.
	ClientSecret awscdk.SecretValue `json:"clientSecret" yaml:"clientSecret"`
	// The method to use for the authorization request.
	//
	// (Can only choose POST, GET or PUT).
	HttpMethod HttpMethod `json:"httpMethod" yaml:"httpMethod"`
	// Additional string parameters to add to the OAuth request body.
	BodyParameters *map[string]HttpParameter `json:"bodyParameters" yaml:"bodyParameters"`
	// Additional string parameters to add to the OAuth request header.
	HeaderParameters *map[string]HttpParameter `json:"headerParameters" yaml:"headerParameters"`
	// Additional string parameters to add to the OAuth request query string.
	QueryStringParameters *map[string]HttpParameter `json:"queryStringParameters" yaml:"queryStringParameters"`
}

// Standard set of options for `onXxx` event handlers on construct.
//
// TODO: EXAMPLE
//
type OnEventOptions struct {
	// A description of the rule's purpose.
	Description *string `json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	EventPattern *EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	Target IRuleTarget `json:"target" yaml:"target"`
}

// Defines an EventBridge Rule in this stack.
//
// TODO: EXAMPLE
//
type Rule interface {
	awscdk.Resource
	IRule
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	RuleArn() *string
	RuleName() *string
	Stack() awscdk.Stack
	AddEventPattern(eventPattern *EventPattern)
	AddTarget(target IRuleTarget)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	ValidateRule() *[]*string
}

// The jsii proxy struct for Rule
type jsiiProxy_Rule struct {
	internal.Type__awscdkResource
	jsiiProxy_IRule
}

func (j *jsiiProxy_Rule) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) RuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewRule(scope constructs.Construct, id *string, props *RuleProps) Rule {
	_init_.Initialize()

	j := jsiiProxy_Rule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Rule",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewRule_Override(r Rule, scope constructs.Construct, id *string, props *RuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Rule",
		[]interface{}{scope, id, props},
		r,
	)
}

// Import an existing EventBridge Rule provided an ARN.
func Rule_FromEventRuleArn(scope constructs.Construct, id *string, eventRuleArn *string) IRule {
	_init_.Initialize()

	var returns IRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Rule",
		"fromEventRuleArn",
		[]interface{}{scope, id, eventRuleArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Rule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Rule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Rule_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Rule",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds an event pattern filter to this rule.
//
// If a pattern was already specified,
// these values are merged into the existing pattern.
//
// For example, if the rule already contains the pattern:
//
//     {
//       "resources": [ "r1" ],
//       "detail": {
//         "hello": [ 1 ]
//       }
//     }
//
// And `addEventPattern` is called with the pattern:
//
//     {
//       "resources": [ "r2" ],
//       "detail": {
//         "foo": [ "bar" ]
//       }
//     }
//
// The resulting event pattern will be:
//
//     {
//       "resources": [ "r1", "r2" ],
//       "detail": {
//         "hello": [ 1 ],
//         "foo": [ "bar" ]
//       }
//     }
func (r *jsiiProxy_Rule) AddEventPattern(eventPattern *EventPattern) {
	_jsii_.InvokeVoid(
		r,
		"addEventPattern",
		[]interface{}{eventPattern},
	)
}

// Adds a target to the rule. The abstract class RuleTarget can be extended to define new targets.
//
// No-op if target is undefined.
func (r *jsiiProxy_Rule) AddTarget(target IRuleTarget) {
	_jsii_.InvokeVoid(
		r,
		"addTarget",
		[]interface{}{target},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (r *jsiiProxy_Rule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_Rule) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
func (r *jsiiProxy_Rule) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
func (r *jsiiProxy_Rule) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_Rule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Rule) ValidateRule() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validateRule",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for defining an EventBridge Rule.
//
// TODO: EXAMPLE
//
type RuleProps struct {
	// A description of the rule's purpose.
	Description *string `json:"description" yaml:"description"`
	// Indicates whether the rule is enabled.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The event bus to associate with this rule.
	EventBus IEventBus `json:"eventBus" yaml:"eventBus"`
	// Describes which events EventBridge routes to the specified target.
	//
	// These routed events are matched events. For more information, see Events
	// and Event Patterns in the Amazon EventBridge User Guide.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	// You must specify this property (either via props or via
	// `addEventPattern`), the `scheduleExpression` property, or both. The
	// method `addEventPattern` can be used to add filter values to the event
	// pattern.
	//
	EventPattern *EventPattern `json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `json:"ruleName" yaml:"ruleName"`
	// The schedule or rate (frequency) that determines when EventBridge runs the rule.
	//
	// For more information, see Schedule Expression Syntax for
	// Rules in the Amazon EventBridge User Guide.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/scheduled-events.html
	//
	// You must specify this property, the `eventPattern` property, or both.
	//
	Schedule Schedule `json:"schedule" yaml:"schedule"`
	// Targets to invoke when this rule matches an event.
	//
	// Input will be the full matched event. If you wish to specify custom
	// target input, use `addTarget(target[, inputOptions])`.
	Targets *[]IRuleTarget `json:"targets" yaml:"targets"`
}

// Properties for an event rule target.
//
// TODO: EXAMPLE
//
type RuleTargetConfig struct {
	// The Amazon Resource Name (ARN) of the target.
	Arn *string `json:"arn" yaml:"arn"`
	// Parameters used when the rule invokes Amazon AWS Batch Job/Queue.
	BatchParameters *CfnRule_BatchParametersProperty `json:"batchParameters" yaml:"batchParameters"`
	// Contains information about a dead-letter queue configuration.
	DeadLetterConfig *CfnRule_DeadLetterConfigProperty `json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The Amazon ECS task definition and task count to use, if the event target is an Amazon ECS task.
	EcsParameters *CfnRule_EcsParametersProperty `json:"ecsParameters" yaml:"ecsParameters"`
	// Contains the HTTP parameters to use when the target is a API Gateway REST endpoint or EventBridge API destination.
	HttpParameters *CfnRule_HttpParametersProperty `json:"httpParameters" yaml:"httpParameters"`
	// What input to send to the event target.
	Input RuleTargetInput `json:"input" yaml:"input"`
	// Settings that control shard assignment, when the target is a Kinesis stream.
	//
	// If you don't include this parameter, eventId is used as the
	// partition key.
	KinesisParameters *CfnRule_KinesisParametersProperty `json:"kinesisParameters" yaml:"kinesisParameters"`
	// A RetryPolicy object that includes information about the retry policy settings.
	RetryPolicy *CfnRule_RetryPolicyProperty `json:"retryPolicy" yaml:"retryPolicy"`
	// Role to use to invoke this event target.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Parameters used when the rule invokes Amazon EC2 Systems Manager Run Command.
	RunCommandParameters *CfnRule_RunCommandParametersProperty `json:"runCommandParameters" yaml:"runCommandParameters"`
	// Parameters used when the FIFO sqs queue is used an event target by the rule.
	SqsParameters *CfnRule_SqsParametersProperty `json:"sqsParameters" yaml:"sqsParameters"`
	// The resource that is backing this target.
	//
	// This is the resource that will actually have some action performed on it when used as a target
	// (for example, start a build for a CodeBuild project).
	// We need it to determine whether the rule belongs to a different account than the target -
	// if so, we generate a more complex setup,
	// including an additional stack containing the EventBusPolicy.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html
	//
	TargetResource constructs.IConstruct `json:"targetResource" yaml:"targetResource"`
}

// The input to send to the event target.
//
// TODO: EXAMPLE
//
type RuleTargetInput interface {
	Bind(rule IRule) *RuleTargetInputProperties
}

// The jsii proxy struct for RuleTargetInput
type jsiiProxy_RuleTargetInput struct {
	_ byte // padding
}

func NewRuleTargetInput_Override(r RuleTargetInput) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		nil, // no parameters
		r,
	)
}

// Take the event target input from a path in the event JSON.
func RuleTargetInput_FromEventPath(path *string) RuleTargetInput {
	_init_.Initialize()

	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromEventPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Pass text to the event target, splitting on newlines.
//
// This is only useful when passing to a target that does not
// take a single argument.
//
// May contain strings returned by `EventField.from()` to substitute in parts
// of the matched event.
func RuleTargetInput_FromMultilineText(text *string) RuleTargetInput {
	_init_.Initialize()

	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromMultilineText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

// Pass a JSON object to the event target.
//
// May contain strings returned by `EventField.from()` to substitute in parts of the
// matched event.
func RuleTargetInput_FromObject(obj interface{}) RuleTargetInput {
	_init_.Initialize()

	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromObject",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

// Pass text to the event target.
//
// May contain strings returned by `EventField.from()` to substitute in parts of the
// matched event.
//
// The Rule Target input value will be a single string: the string you pass
// here.  Do not use this method to pass a complex value like a JSON object to
// a Rule Target.  Use `RuleTargetInput.fromObject()` instead.
func RuleTargetInput_FromText(text *string) RuleTargetInput {
	_init_.Initialize()

	var returns RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.RuleTargetInput",
		"fromText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

// Return the input properties for this input object.
func (r *jsiiProxy_RuleTargetInput) Bind(rule IRule) *RuleTargetInputProperties {
	var returns *RuleTargetInputProperties

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// The input properties for an event target.
//
// TODO: EXAMPLE
//
type RuleTargetInputProperties struct {
	// Literal input to the target service (must be valid JSON).
	Input *string `json:"input" yaml:"input"`
	// JsonPath to take input from the input event.
	InputPath *string `json:"inputPath" yaml:"inputPath"`
	// Paths map to extract values from event and insert into `inputTemplate`.
	InputPathsMap *map[string]*string `json:"inputPathsMap" yaml:"inputPathsMap"`
	// Input template to insert paths map into.
	InputTemplate *string `json:"inputTemplate" yaml:"inputTemplate"`
}

// Schedule for scheduled event rules.
//
// TODO: EXAMPLE
//
type Schedule interface {
	ExpressionString() *string
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Schedule",
		nil, // no parameters
		s,
	)
}

// Create a schedule from a set of cron fields.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
func Schedule_Rate(duration awscdk.Duration) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Schedule",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

