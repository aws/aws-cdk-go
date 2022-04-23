package awssam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Serverless::Api`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//   var gatewayResponses interface{}
//   var methodSettings interface{}
//   var models interface{}
//   cfnApi := sam.NewCfnApi(this, jsii.String("MyCfnApi"), &cfnApiProps{
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &authProperty{
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	binaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	canarySetting: &canarySettingProperty{
//   		deploymentId: jsii.String("deploymentId"),
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	cors: jsii.String("cors"),
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	gatewayResponses: gatewayResponses,
//   	methodSettings: []interface{}{
//   		methodSettings,
//   	},
//   	minimumCompressionSize: jsii.Number(123),
//   	models: models,
//   	name: jsii.String("name"),
//   	openApiVersion: jsii.String("openApiVersion"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   })
//
type CfnApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::Serverless::Api.AccessLogSetting`.
	AccessLogSetting() interface{}
	SetAccessLogSetting(val interface{})
	// `AWS::Serverless::Api.Auth`.
	Auth() interface{}
	SetAuth(val interface{})
	// `AWS::Serverless::Api.BinaryMediaTypes`.
	BinaryMediaTypes() *[]*string
	SetBinaryMediaTypes(val *[]*string)
	// `AWS::Serverless::Api.CacheClusterEnabled`.
	CacheClusterEnabled() interface{}
	SetCacheClusterEnabled(val interface{})
	// `AWS::Serverless::Api.CacheClusterSize`.
	CacheClusterSize() *string
	SetCacheClusterSize(val *string)
	// `AWS::Serverless::Api.CanarySetting`.
	CanarySetting() interface{}
	SetCanarySetting(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Serverless::Api.Cors`.
	Cors() interface{}
	SetCors(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Serverless::Api.DefinitionBody`.
	DefinitionBody() interface{}
	SetDefinitionBody(val interface{})
	// `AWS::Serverless::Api.DefinitionUri`.
	DefinitionUri() interface{}
	SetDefinitionUri(val interface{})
	// `AWS::Serverless::Api.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::Api.EndpointConfiguration`.
	EndpointConfiguration() interface{}
	SetEndpointConfiguration(val interface{})
	// `AWS::Serverless::Api.GatewayResponses`.
	GatewayResponses() interface{}
	SetGatewayResponses(val interface{})
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
	// `AWS::Serverless::Api.MethodSettings`.
	MethodSettings() interface{}
	SetMethodSettings(val interface{})
	// `AWS::Serverless::Api.MinimumCompressionSize`.
	MinimumCompressionSize() *float64
	SetMinimumCompressionSize(val *float64)
	// `AWS::Serverless::Api.Models`.
	Models() interface{}
	SetModels(val interface{})
	// `AWS::Serverless::Api.Name`.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// `AWS::Serverless::Api.OpenApiVersion`.
	OpenApiVersion() *string
	SetOpenApiVersion(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Serverless::Api.StageName`.
	StageName() *string
	SetStageName(val *string)
	// `AWS::Serverless::Api.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::Api.TracingEnabled`.
	TracingEnabled() interface{}
	SetTracingEnabled(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// `AWS::Serverless::Api.Variables`.
	Variables() interface{}
	SetVariables(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnApi
type jsiiProxy_CfnApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApi) AccessLogSetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Auth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) BinaryMediaTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"binaryMediaTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CacheClusterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheClusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CacheClusterSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheClusterSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CanarySetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"canarySetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Cors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) DefinitionBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) DefinitionUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) EndpointConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) GatewayResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) MethodSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"methodSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) MinimumCompressionSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCompressionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Models() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) OpenApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) TracingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Variables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::Api`.
func NewCfnApi(scope constructs.Construct, id *string, props *CfnApiProps) CfnApi {
	_init_.Initialize()

	j := jsiiProxy_CfnApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::Api`.
func NewCfnApi_Override(c CfnApi, scope constructs.Construct, id *string, props *CfnApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApi) SetAccessLogSetting(val interface{}) {
	_jsii_.Set(
		j,
		"accessLogSetting",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetAuth(val interface{}) {
	_jsii_.Set(
		j,
		"auth",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetBinaryMediaTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"binaryMediaTypes",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCacheClusterEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"cacheClusterEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCacheClusterSize(val *string) {
	_jsii_.Set(
		j,
		"cacheClusterSize",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCanarySetting(val interface{}) {
	_jsii_.Set(
		j,
		"canarySetting",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCors(val interface{}) {
	_jsii_.Set(
		j,
		"cors",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDefinitionBody(val interface{}) {
	_jsii_.Set(
		j,
		"definitionBody",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDefinitionUri(val interface{}) {
	_jsii_.Set(
		j,
		"definitionUri",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetEndpointConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"endpointConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetGatewayResponses(val interface{}) {
	_jsii_.Set(
		j,
		"gatewayResponses",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetMethodSettings(val interface{}) {
	_jsii_.Set(
		j,
		"methodSettings",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetMinimumCompressionSize(val *float64) {
	_jsii_.Set(
		j,
		"minimumCompressionSize",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetModels(val interface{}) {
	_jsii_.Set(
		j,
		"models",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetOpenApiVersion(val *string) {
	_jsii_.Set(
		j,
		"openApiVersion",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetTracingEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"tracingEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetVariables(val interface{}) {
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApi",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnApi_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnApi",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApi) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApi) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApi) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApi) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApi) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   accessLogSettingProperty := &accessLogSettingProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
type CfnApi_AccessLogSettingProperty struct {
	// `CfnApi.AccessLogSettingProperty.DestinationArn`.
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
	// `CfnApi.AccessLogSettingProperty.Format`.
	Format *string `json:"format" yaml:"format"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var authorizers interface{}
//   authProperty := &authProperty{
//   	authorizers: authorizers,
//   	defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   }
//
type CfnApi_AuthProperty struct {
	// `CfnApi.AuthProperty.Authorizers`.
	Authorizers interface{} `json:"authorizers" yaml:"authorizers"`
	// `CfnApi.AuthProperty.DefaultAuthorizer`.
	DefaultAuthorizer *string `json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   canarySettingProperty := &canarySettingProperty{
//   	deploymentId: jsii.String("deploymentId"),
//   	percentTraffic: jsii.Number(123),
//   	stageVariableOverrides: map[string]*string{
//   		"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   	},
//   	useStageCache: jsii.Boolean(false),
//   }
//
type CfnApi_CanarySettingProperty struct {
	// `CfnApi.CanarySettingProperty.DeploymentId`.
	DeploymentId *string `json:"deploymentId" yaml:"deploymentId"`
	// `CfnApi.CanarySettingProperty.PercentTraffic`.
	PercentTraffic *float64 `json:"percentTraffic" yaml:"percentTraffic"`
	// `CfnApi.CanarySettingProperty.StageVariableOverrides`.
	StageVariableOverrides interface{} `json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// `CfnApi.CanarySettingProperty.UseStageCache`.
	UseStageCache interface{} `json:"useStageCache" yaml:"useStageCache"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   corsConfigurationProperty := &corsConfigurationProperty{
//   	allowOrigin: jsii.String("allowOrigin"),
//
//   	// the properties below are optional
//   	allowCredentials: jsii.Boolean(false),
//   	allowHeaders: jsii.String("allowHeaders"),
//   	allowMethods: jsii.String("allowMethods"),
//   	maxAge: jsii.String("maxAge"),
//   }
//
type CfnApi_CorsConfigurationProperty struct {
	// `CfnApi.CorsConfigurationProperty.AllowOrigin`.
	AllowOrigin *string `json:"allowOrigin" yaml:"allowOrigin"`
	// `CfnApi.CorsConfigurationProperty.AllowCredentials`.
	AllowCredentials interface{} `json:"allowCredentials" yaml:"allowCredentials"`
	// `CfnApi.CorsConfigurationProperty.AllowHeaders`.
	AllowHeaders *string `json:"allowHeaders" yaml:"allowHeaders"`
	// `CfnApi.CorsConfigurationProperty.AllowMethods`.
	AllowMethods *string `json:"allowMethods" yaml:"allowMethods"`
	// `CfnApi.CorsConfigurationProperty.MaxAge`.
	MaxAge *string `json:"maxAge" yaml:"maxAge"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   endpointConfigurationProperty := &endpointConfigurationProperty{
//   	type: jsii.String("type"),
//   	vpcEndpointIds: []*string{
//   		jsii.String("vpcEndpointIds"),
//   	},
//   }
//
type CfnApi_EndpointConfigurationProperty struct {
	// `CfnApi.EndpointConfigurationProperty.Type`.
	Type *string `json:"type" yaml:"type"`
	// `CfnApi.EndpointConfigurationProperty.VpcEndpointIds`.
	VpcEndpointIds *[]*string `json:"vpcEndpointIds" yaml:"vpcEndpointIds"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	version: jsii.Number(123),
//   }
//
type CfnApi_S3LocationProperty struct {
	// `CfnApi.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// `CfnApi.S3LocationProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnApi.S3LocationProperty.Version`.
	Version *float64 `json:"version" yaml:"version"`
}

// Properties for defining a `CfnApi`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//   var gatewayResponses interface{}
//   var methodSettings interface{}
//   var models interface{}
//   cfnApiProps := &cfnApiProps{
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &authProperty{
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	binaryMediaTypes: []*string{
//   		jsii.String("binaryMediaTypes"),
//   	},
//   	cacheClusterEnabled: jsii.Boolean(false),
//   	cacheClusterSize: jsii.String("cacheClusterSize"),
//   	canarySetting: &canarySettingProperty{
//   		deploymentId: jsii.String("deploymentId"),
//   		percentTraffic: jsii.Number(123),
//   		stageVariableOverrides: map[string]*string{
//   			"stageVariableOverridesKey": jsii.String("stageVariableOverrides"),
//   		},
//   		useStageCache: jsii.Boolean(false),
//   	},
//   	cors: jsii.String("cors"),
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	gatewayResponses: gatewayResponses,
//   	methodSettings: []interface{}{
//   		methodSettings,
//   	},
//   	minimumCompressionSize: jsii.Number(123),
//   	models: models,
//   	name: jsii.String("name"),
//   	openApiVersion: jsii.String("openApiVersion"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracingEnabled: jsii.Boolean(false),
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnApiProps struct {
	// `AWS::Serverless::Api.StageName`.
	StageName *string `json:"stageName" yaml:"stageName"`
	// `AWS::Serverless::Api.AccessLogSetting`.
	AccessLogSetting interface{} `json:"accessLogSetting" yaml:"accessLogSetting"`
	// `AWS::Serverless::Api.Auth`.
	Auth interface{} `json:"auth" yaml:"auth"`
	// `AWS::Serverless::Api.BinaryMediaTypes`.
	BinaryMediaTypes *[]*string `json:"binaryMediaTypes" yaml:"binaryMediaTypes"`
	// `AWS::Serverless::Api.CacheClusterEnabled`.
	CacheClusterEnabled interface{} `json:"cacheClusterEnabled" yaml:"cacheClusterEnabled"`
	// `AWS::Serverless::Api.CacheClusterSize`.
	CacheClusterSize *string `json:"cacheClusterSize" yaml:"cacheClusterSize"`
	// `AWS::Serverless::Api.CanarySetting`.
	CanarySetting interface{} `json:"canarySetting" yaml:"canarySetting"`
	// `AWS::Serverless::Api.Cors`.
	Cors interface{} `json:"cors" yaml:"cors"`
	// `AWS::Serverless::Api.DefinitionBody`.
	DefinitionBody interface{} `json:"definitionBody" yaml:"definitionBody"`
	// `AWS::Serverless::Api.DefinitionUri`.
	DefinitionUri interface{} `json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::Api.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::Serverless::Api.EndpointConfiguration`.
	EndpointConfiguration interface{} `json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `AWS::Serverless::Api.GatewayResponses`.
	GatewayResponses interface{} `json:"gatewayResponses" yaml:"gatewayResponses"`
	// `AWS::Serverless::Api.MethodSettings`.
	MethodSettings interface{} `json:"methodSettings" yaml:"methodSettings"`
	// `AWS::Serverless::Api.MinimumCompressionSize`.
	MinimumCompressionSize *float64 `json:"minimumCompressionSize" yaml:"minimumCompressionSize"`
	// `AWS::Serverless::Api.Models`.
	Models interface{} `json:"models" yaml:"models"`
	// `AWS::Serverless::Api.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::Serverless::Api.OpenApiVersion`.
	OpenApiVersion *string `json:"openApiVersion" yaml:"openApiVersion"`
	// `AWS::Serverless::Api.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// `AWS::Serverless::Api.TracingEnabled`.
	TracingEnabled interface{} `json:"tracingEnabled" yaml:"tracingEnabled"`
	// `AWS::Serverless::Api.Variables`.
	Variables interface{} `json:"variables" yaml:"variables"`
}

// A CloudFormation `AWS::Serverless::Application`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cfnApplication := sam.NewCfnApplication(this, jsii.String("MyCfnApplication"), &cfnApplicationProps{
//   	location: jsii.String("location"),
//
//   	// the properties below are optional
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeoutInMinutes: jsii.Number(123),
//   })
//
type CfnApplication interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Serverless::Application.Location`.
	Location() interface{}
	SetLocation(val interface{})
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
	// `AWS::Serverless::Application.NotificationArns`.
	NotificationArns() *[]*string
	SetNotificationArns(val *[]*string)
	// `AWS::Serverless::Application.Parameters`.
	Parameters() interface{}
	SetParameters(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Serverless::Application.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::Application.TimeoutInMinutes`.
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnApplication
type jsiiProxy_CfnApplication struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApplication) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Parameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplication) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::Application`.
func NewCfnApplication(scope constructs.Construct, id *string, props *CfnApplicationProps) CfnApplication {
	_init_.Initialize()

	j := jsiiProxy_CfnApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnApplication",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::Application`.
func NewCfnApplication_Override(c CfnApplication, scope constructs.Construct, id *string, props *CfnApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnApplication",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApplication) SetLocation(val interface{}) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetNotificationArns(val *[]*string) {
	_jsii_.Set(
		j,
		"notificationArns",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetParameters(val interface{}) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_CfnApplication) SetTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnApplication_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApplication",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnApplication_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApplication",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplication_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnApplication",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnApplication_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnApplication",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplication) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnApplication) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnApplication) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnApplication) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnApplication) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnApplication) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnApplication) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApplication) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnApplication) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   applicationLocationProperty := &applicationLocationProperty{
//   	applicationId: jsii.String("applicationId"),
//   	semanticVersion: jsii.String("semanticVersion"),
//   }
//
type CfnApplication_ApplicationLocationProperty struct {
	// `CfnApplication.ApplicationLocationProperty.ApplicationId`.
	ApplicationId *string `json:"applicationId" yaml:"applicationId"`
	// `CfnApplication.ApplicationLocationProperty.SemanticVersion`.
	SemanticVersion *string `json:"semanticVersion" yaml:"semanticVersion"`
}

// Properties for defining a `CfnApplication`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cfnApplicationProps := &cfnApplicationProps{
//   	location: jsii.String("location"),
//
//   	// the properties below are optional
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeoutInMinutes: jsii.Number(123),
//   }
//
type CfnApplicationProps struct {
	// `AWS::Serverless::Application.Location`.
	Location interface{} `json:"location" yaml:"location"`
	// `AWS::Serverless::Application.NotificationArns`.
	NotificationArns *[]*string `json:"notificationArns" yaml:"notificationArns"`
	// `AWS::Serverless::Application.Parameters`.
	Parameters interface{} `json:"parameters" yaml:"parameters"`
	// `AWS::Serverless::Application.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// `AWS::Serverless::Application.TimeoutInMinutes`.
	TimeoutInMinutes *float64 `json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

// A CloudFormation `AWS::Serverless::Function`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var assumeRolePolicyDocument interface{}
//   cfnFunction := sam.NewCfnFunction(this, jsii.String("MyCfnFunction"), &cfnFunctionProps{
//   	architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	assumeRolePolicyDocument: assumeRolePolicyDocument,
//   	autoPublishAlias: jsii.String("autoPublishAlias"),
//   	autoPublishCodeSha256: jsii.String("autoPublishCodeSha256"),
//   	codeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	codeUri: jsii.String("codeUri"),
//   	deadLetterQueue: &deadLetterQueueProperty{
//   		targetArn: jsii.String("targetArn"),
//   		type: jsii.String("type"),
//   	},
//   	deploymentPreference: &deploymentPreferenceProperty{
//   		enabled: jsii.Boolean(false),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		hooks: &hooksProperty{
//   			postTraffic: jsii.String("postTraffic"),
//   			preTraffic: jsii.String("preTraffic"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	environment: &functionEnvironmentProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	eventInvokeConfig: &eventInvokeConfigProperty{
//   		destinationConfig: &eventInvokeDestinationConfigProperty{
//   			onFailure: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   			onSuccess: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   		},
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &S3EventProperty{
//   				"variables": map[string]*string{
//   					"variablesKey": jsii.String("variables"),
//   				},
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	fileSystemConfigs: []interface{}{
//   		&fileSystemConfigProperty{
//   			arn: jsii.String("arn"),
//   			localMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	functionName: jsii.String("functionName"),
//   	handler: jsii.String("handler"),
//   	imageConfig: &imageConfigProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	imageUri: jsii.String("imageUri"),
//   	inlineCode: jsii.String("inlineCode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	memorySize: jsii.Number(123),
//   	packageType: jsii.String("packageType"),
//   	permissionsBoundary: jsii.String("permissionsBoundary"),
//   	policies: jsii.String("policies"),
//   	provisionedConcurrencyConfig: &provisionedConcurrencyConfigProperty{
//   		provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   	},
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	role: jsii.String("role"),
//   	runtime: jsii.String("runtime"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeout: jsii.Number(123),
//   	tracing: jsii.String("tracing"),
//   	versionDescription: jsii.String("versionDescription"),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   })
//
type CfnFunction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::Serverless::Function.Architectures`.
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	// `AWS::Serverless::Function.AssumeRolePolicyDocument`.
	AssumeRolePolicyDocument() interface{}
	SetAssumeRolePolicyDocument(val interface{})
	// `AWS::Serverless::Function.AutoPublishAlias`.
	AutoPublishAlias() *string
	SetAutoPublishAlias(val *string)
	// `AWS::Serverless::Function.AutoPublishCodeSha256`.
	AutoPublishCodeSha256() *string
	SetAutoPublishCodeSha256(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Serverless::Function.CodeSigningConfigArn`.
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	// `AWS::Serverless::Function.CodeUri`.
	CodeUri() interface{}
	SetCodeUri(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Serverless::Function.DeadLetterQueue`.
	DeadLetterQueue() interface{}
	SetDeadLetterQueue(val interface{})
	// `AWS::Serverless::Function.DeploymentPreference`.
	DeploymentPreference() interface{}
	SetDeploymentPreference(val interface{})
	// `AWS::Serverless::Function.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::Function.Environment`.
	Environment() interface{}
	SetEnvironment(val interface{})
	// `AWS::Serverless::Function.EventInvokeConfig`.
	EventInvokeConfig() interface{}
	SetEventInvokeConfig(val interface{})
	// `AWS::Serverless::Function.Events`.
	Events() interface{}
	SetEvents(val interface{})
	// `AWS::Serverless::Function.FileSystemConfigs`.
	FileSystemConfigs() interface{}
	SetFileSystemConfigs(val interface{})
	// `AWS::Serverless::Function.FunctionName`.
	FunctionName() *string
	SetFunctionName(val *string)
	// `AWS::Serverless::Function.Handler`.
	Handler() *string
	SetHandler(val *string)
	// `AWS::Serverless::Function.ImageConfig`.
	ImageConfig() interface{}
	SetImageConfig(val interface{})
	// `AWS::Serverless::Function.ImageUri`.
	ImageUri() *string
	SetImageUri(val *string)
	// `AWS::Serverless::Function.InlineCode`.
	InlineCode() *string
	SetInlineCode(val *string)
	// `AWS::Serverless::Function.KmsKeyArn`.
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	// `AWS::Serverless::Function.Layers`.
	Layers() *[]*string
	SetLayers(val *[]*string)
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
	// `AWS::Serverless::Function.MemorySize`.
	MemorySize() *float64
	SetMemorySize(val *float64)
	// The tree node.
	Node() constructs.Node
	// `AWS::Serverless::Function.PackageType`.
	PackageType() *string
	SetPackageType(val *string)
	// `AWS::Serverless::Function.PermissionsBoundary`.
	PermissionsBoundary() *string
	SetPermissionsBoundary(val *string)
	// `AWS::Serverless::Function.Policies`.
	Policies() interface{}
	SetPolicies(val interface{})
	// `AWS::Serverless::Function.ProvisionedConcurrencyConfig`.
	ProvisionedConcurrencyConfig() interface{}
	SetProvisionedConcurrencyConfig(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::Serverless::Function.ReservedConcurrentExecutions`.
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	// `AWS::Serverless::Function.Role`.
	Role() *string
	SetRole(val *string)
	// `AWS::Serverless::Function.Runtime`.
	Runtime() *string
	SetRuntime(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Serverless::Function.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::Function.Timeout`.
	Timeout() *float64
	SetTimeout(val *float64)
	// `AWS::Serverless::Function.Tracing`.
	Tracing() *string
	SetTracing(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// `AWS::Serverless::Function.VersionDescription`.
	VersionDescription() *string
	SetVersionDescription(val *string)
	// `AWS::Serverless::Function.VpcConfig`.
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnFunction
type jsiiProxy_CfnFunction struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AssumeRolePolicyDocument() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRolePolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AutoPublishAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoPublishAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) AutoPublishCodeSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoPublishCodeSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CodeUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) DeadLetterQueue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) DeploymentPreference() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deploymentPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) EventInvokeConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventInvokeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Events() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FileSystemConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ImageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) InlineCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) PermissionsBoundary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Policies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ProvisionedConcurrencyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedConcurrencyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) Tracing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::Function`.
func NewCfnFunction(scope constructs.Construct, id *string, props *CfnFunctionProps) CfnFunction {
	_init_.Initialize()

	j := jsiiProxy_CfnFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::Function`.
func NewCfnFunction_Override(c CfnFunction, scope constructs.Construct, id *string, props *CfnFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnFunction",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunction) SetArchitectures(val *[]*string) {
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetAssumeRolePolicyDocument(val interface{}) {
	_jsii_.Set(
		j,
		"assumeRolePolicyDocument",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetAutoPublishAlias(val *string) {
	_jsii_.Set(
		j,
		"autoPublishAlias",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetAutoPublishCodeSha256(val *string) {
	_jsii_.Set(
		j,
		"autoPublishCodeSha256",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetCodeSigningConfigArn(val *string) {
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetCodeUri(val interface{}) {
	_jsii_.Set(
		j,
		"codeUri",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetDeadLetterQueue(val interface{}) {
	_jsii_.Set(
		j,
		"deadLetterQueue",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetDeploymentPreference(val interface{}) {
	_jsii_.Set(
		j,
		"deploymentPreference",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetEnvironment(val interface{}) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetEventInvokeConfig(val interface{}) {
	_jsii_.Set(
		j,
		"eventInvokeConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetEvents(val interface{}) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetFileSystemConfigs(val interface{}) {
	_jsii_.Set(
		j,
		"fileSystemConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetHandler(val *string) {
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetImageConfig(val interface{}) {
	_jsii_.Set(
		j,
		"imageConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetImageUri(val *string) {
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetInlineCode(val *string) {
	_jsii_.Set(
		j,
		"inlineCode",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetKmsKeyArn(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetLayers(val *[]*string) {
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetMemorySize(val *float64) {
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetPackageType(val *string) {
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetPermissionsBoundary(val *string) {
	_jsii_.Set(
		j,
		"permissionsBoundary",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetProvisionedConcurrencyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"provisionedConcurrencyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetReservedConcurrentExecutions(val *float64) {
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetRuntime(val *string) {
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetTracing(val *string) {
	_jsii_.Set(
		j,
		"tracing",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetVersionDescription(val *string) {
	_jsii_.Set(
		j,
		"versionDescription",
		val,
	)
}

func (j *jsiiProxy_CfnFunction) SetVpcConfig(val interface{}) {
	_jsii_.Set(
		j,
		"vpcConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFunction_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnFunction",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFunction_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnFunction",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunction_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnFunction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnFunction_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnFunction",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFunction) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFunction) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFunction) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFunction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFunction) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFunction) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFunction) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFunction) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFunction) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   alexaSkillEventProperty := &alexaSkillEventProperty{
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnFunction_AlexaSkillEventProperty struct {
	// `CfnFunction.AlexaSkillEventProperty.Variables`.
	Variables interface{} `json:"variables" yaml:"variables"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var customStatements interface{}
//   apiEventProperty := &apiEventProperty{
//   	method: jsii.String("method"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	auth: &authProperty{
//   		apiKeyRequired: jsii.Boolean(false),
//   		authorizationScopes: []*string{
//   			jsii.String("authorizationScopes"),
//   		},
//   		authorizer: jsii.String("authorizer"),
//   		resourcePolicy: &authResourcePolicyProperty{
//   			awsAccountBlacklist: []*string{
//   				jsii.String("awsAccountBlacklist"),
//   			},
//   			awsAccountWhitelist: []*string{
//   				jsii.String("awsAccountWhitelist"),
//   			},
//   			customStatements: []interface{}{
//   				customStatements,
//   			},
//   			intrinsicVpcBlacklist: []*string{
//   				jsii.String("intrinsicVpcBlacklist"),
//   			},
//   			intrinsicVpceBlacklist: []*string{
//   				jsii.String("intrinsicVpceBlacklist"),
//   			},
//   			intrinsicVpceWhitelist: []*string{
//   				jsii.String("intrinsicVpceWhitelist"),
//   			},
//   			intrinsicVpcWhitelist: []*string{
//   				jsii.String("intrinsicVpcWhitelist"),
//   			},
//   			ipRangeBlacklist: []*string{
//   				jsii.String("ipRangeBlacklist"),
//   			},
//   			ipRangeWhitelist: []*string{
//   				jsii.String("ipRangeWhitelist"),
//   			},
//   			sourceVpcBlacklist: []*string{
//   				jsii.String("sourceVpcBlacklist"),
//   			},
//   			sourceVpcWhitelist: []*string{
//   				jsii.String("sourceVpcWhitelist"),
//   			},
//   		},
//   	},
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnFunction_ApiEventProperty struct {
	// `CfnFunction.ApiEventProperty.Method`.
	Method *string `json:"method" yaml:"method"`
	// `CfnFunction.ApiEventProperty.Path`.
	Path *string `json:"path" yaml:"path"`
	// `CfnFunction.ApiEventProperty.Auth`.
	Auth interface{} `json:"auth" yaml:"auth"`
	// `CfnFunction.ApiEventProperty.RestApiId`.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var customStatements interface{}
//   authProperty := &authProperty{
//   	apiKeyRequired: jsii.Boolean(false),
//   	authorizationScopes: []*string{
//   		jsii.String("authorizationScopes"),
//   	},
//   	authorizer: jsii.String("authorizer"),
//   	resourcePolicy: &authResourcePolicyProperty{
//   		awsAccountBlacklist: []*string{
//   			jsii.String("awsAccountBlacklist"),
//   		},
//   		awsAccountWhitelist: []*string{
//   			jsii.String("awsAccountWhitelist"),
//   		},
//   		customStatements: []interface{}{
//   			customStatements,
//   		},
//   		intrinsicVpcBlacklist: []*string{
//   			jsii.String("intrinsicVpcBlacklist"),
//   		},
//   		intrinsicVpceBlacklist: []*string{
//   			jsii.String("intrinsicVpceBlacklist"),
//   		},
//   		intrinsicVpceWhitelist: []*string{
//   			jsii.String("intrinsicVpceWhitelist"),
//   		},
//   		intrinsicVpcWhitelist: []*string{
//   			jsii.String("intrinsicVpcWhitelist"),
//   		},
//   		ipRangeBlacklist: []*string{
//   			jsii.String("ipRangeBlacklist"),
//   		},
//   		ipRangeWhitelist: []*string{
//   			jsii.String("ipRangeWhitelist"),
//   		},
//   		sourceVpcBlacklist: []*string{
//   			jsii.String("sourceVpcBlacklist"),
//   		},
//   		sourceVpcWhitelist: []*string{
//   			jsii.String("sourceVpcWhitelist"),
//   		},
//   	},
//   }
//
type CfnFunction_AuthProperty struct {
	// `CfnFunction.AuthProperty.ApiKeyRequired`.
	ApiKeyRequired interface{} `json:"apiKeyRequired" yaml:"apiKeyRequired"`
	// `CfnFunction.AuthProperty.AuthorizationScopes`.
	AuthorizationScopes *[]*string `json:"authorizationScopes" yaml:"authorizationScopes"`
	// `CfnFunction.AuthProperty.Authorizer`.
	Authorizer *string `json:"authorizer" yaml:"authorizer"`
	// `CfnFunction.AuthProperty.ResourcePolicy`.
	ResourcePolicy interface{} `json:"resourcePolicy" yaml:"resourcePolicy"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var customStatements interface{}
//   authResourcePolicyProperty := &authResourcePolicyProperty{
//   	awsAccountBlacklist: []*string{
//   		jsii.String("awsAccountBlacklist"),
//   	},
//   	awsAccountWhitelist: []*string{
//   		jsii.String("awsAccountWhitelist"),
//   	},
//   	customStatements: []interface{}{
//   		customStatements,
//   	},
//   	intrinsicVpcBlacklist: []*string{
//   		jsii.String("intrinsicVpcBlacklist"),
//   	},
//   	intrinsicVpceBlacklist: []*string{
//   		jsii.String("intrinsicVpceBlacklist"),
//   	},
//   	intrinsicVpceWhitelist: []*string{
//   		jsii.String("intrinsicVpceWhitelist"),
//   	},
//   	intrinsicVpcWhitelist: []*string{
//   		jsii.String("intrinsicVpcWhitelist"),
//   	},
//   	ipRangeBlacklist: []*string{
//   		jsii.String("ipRangeBlacklist"),
//   	},
//   	ipRangeWhitelist: []*string{
//   		jsii.String("ipRangeWhitelist"),
//   	},
//   	sourceVpcBlacklist: []*string{
//   		jsii.String("sourceVpcBlacklist"),
//   	},
//   	sourceVpcWhitelist: []*string{
//   		jsii.String("sourceVpcWhitelist"),
//   	},
//   }
//
type CfnFunction_AuthResourcePolicyProperty struct {
	// `CfnFunction.AuthResourcePolicyProperty.AwsAccountBlacklist`.
	AwsAccountBlacklist *[]*string `json:"awsAccountBlacklist" yaml:"awsAccountBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.AwsAccountWhitelist`.
	AwsAccountWhitelist *[]*string `json:"awsAccountWhitelist" yaml:"awsAccountWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.CustomStatements`.
	CustomStatements interface{} `json:"customStatements" yaml:"customStatements"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpcBlacklist`.
	IntrinsicVpcBlacklist *[]*string `json:"intrinsicVpcBlacklist" yaml:"intrinsicVpcBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpceBlacklist`.
	IntrinsicVpceBlacklist *[]*string `json:"intrinsicVpceBlacklist" yaml:"intrinsicVpceBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpceWhitelist`.
	IntrinsicVpceWhitelist *[]*string `json:"intrinsicVpceWhitelist" yaml:"intrinsicVpceWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.IntrinsicVpcWhitelist`.
	IntrinsicVpcWhitelist *[]*string `json:"intrinsicVpcWhitelist" yaml:"intrinsicVpcWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.IpRangeBlacklist`.
	IpRangeBlacklist *[]*string `json:"ipRangeBlacklist" yaml:"ipRangeBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.IpRangeWhitelist`.
	IpRangeWhitelist *[]*string `json:"ipRangeWhitelist" yaml:"ipRangeWhitelist"`
	// `CfnFunction.AuthResourcePolicyProperty.SourceVpcBlacklist`.
	SourceVpcBlacklist *[]*string `json:"sourceVpcBlacklist" yaml:"sourceVpcBlacklist"`
	// `CfnFunction.AuthResourcePolicyProperty.SourceVpcWhitelist`.
	SourceVpcWhitelist *[]*string `json:"sourceVpcWhitelist" yaml:"sourceVpcWhitelist"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   bucketSAMPTProperty := &bucketSAMPTProperty{
//   	bucketName: jsii.String("bucketName"),
//   }
//
type CfnFunction_BucketSAMPTProperty struct {
	// `CfnFunction.BucketSAMPTProperty.BucketName`.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var pattern interface{}
//   cloudWatchEventEventProperty := &cloudWatchEventEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnFunction_CloudWatchEventEventProperty struct {
	// `CfnFunction.CloudWatchEventEventProperty.Pattern`.
	Pattern interface{} `json:"pattern" yaml:"pattern"`
	// `CfnFunction.CloudWatchEventEventProperty.Input`.
	Input *string `json:"input" yaml:"input"`
	// `CfnFunction.CloudWatchEventEventProperty.InputPath`.
	InputPath *string `json:"inputPath" yaml:"inputPath"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cloudWatchLogsEventProperty := &cloudWatchLogsEventProperty{
//   	filterPattern: jsii.String("filterPattern"),
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnFunction_CloudWatchLogsEventProperty struct {
	// `CfnFunction.CloudWatchLogsEventProperty.FilterPattern`.
	FilterPattern *string `json:"filterPattern" yaml:"filterPattern"`
	// `CfnFunction.CloudWatchLogsEventProperty.LogGroupName`.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   collectionSAMPTProperty := &collectionSAMPTProperty{
//   	collectionId: jsii.String("collectionId"),
//   }
//
type CfnFunction_CollectionSAMPTProperty struct {
	// `CfnFunction.CollectionSAMPTProperty.CollectionId`.
	CollectionId *string `json:"collectionId" yaml:"collectionId"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   deadLetterQueueProperty := &deadLetterQueueProperty{
//   	targetArn: jsii.String("targetArn"),
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_DeadLetterQueueProperty struct {
	// `CfnFunction.DeadLetterQueueProperty.TargetArn`.
	TargetArn *string `json:"targetArn" yaml:"targetArn"`
	// `CfnFunction.DeadLetterQueueProperty.Type`.
	Type *string `json:"type" yaml:"type"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   deploymentPreferenceProperty := &deploymentPreferenceProperty{
//   	enabled: jsii.Boolean(false),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	alarms: []*string{
//   		jsii.String("alarms"),
//   	},
//   	hooks: &hooksProperty{
//   		postTraffic: jsii.String("postTraffic"),
//   		preTraffic: jsii.String("preTraffic"),
//   	},
//   }
//
type CfnFunction_DeploymentPreferenceProperty struct {
	// `CfnFunction.DeploymentPreferenceProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// `CfnFunction.DeploymentPreferenceProperty.Type`.
	Type *string `json:"type" yaml:"type"`
	// `CfnFunction.DeploymentPreferenceProperty.Alarms`.
	Alarms *[]*string `json:"alarms" yaml:"alarms"`
	// `CfnFunction.DeploymentPreferenceProperty.Hooks`.
	Hooks interface{} `json:"hooks" yaml:"hooks"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   destinationConfigProperty := &destinationConfigProperty{
//   	onFailure: &destinationProperty{
//   		destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnFunction_DestinationConfigProperty struct {
	// `CfnFunction.DestinationConfigProperty.OnFailure`.
	OnFailure interface{} `json:"onFailure" yaml:"onFailure"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   destinationProperty := &destinationProperty{
//   	destination: jsii.String("destination"),
//
//   	// the properties below are optional
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_DestinationProperty struct {
	// `CfnFunction.DestinationProperty.Destination`.
	Destination *string `json:"destination" yaml:"destination"`
	// `CfnFunction.DestinationProperty.Type`.
	Type *string `json:"type" yaml:"type"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   domainSAMPTProperty := &domainSAMPTProperty{
//   	domainName: jsii.String("domainName"),
//   }
//
type CfnFunction_DomainSAMPTProperty struct {
	// `CfnFunction.DomainSAMPTProperty.DomainName`.
	DomainName *string `json:"domainName" yaml:"domainName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   dynamoDBEventProperty := &dynamoDBEventProperty{
//   	startingPosition: jsii.String("startingPosition"),
//   	stream: jsii.String("stream"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	bisectBatchOnFunctionError: jsii.Boolean(false),
//   	destinationConfig: &destinationConfigProperty{
//   		onFailure: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   	},
//   	enabled: jsii.Boolean(false),
//   	maximumBatchingWindowInSeconds: jsii.Number(123),
//   	maximumRecordAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   	parallelizationFactor: jsii.Number(123),
//   }
//
type CfnFunction_DynamoDBEventProperty struct {
	// `CfnFunction.DynamoDBEventProperty.StartingPosition`.
	StartingPosition *string `json:"startingPosition" yaml:"startingPosition"`
	// `CfnFunction.DynamoDBEventProperty.Stream`.
	Stream *string `json:"stream" yaml:"stream"`
	// `CfnFunction.DynamoDBEventProperty.BatchSize`.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.DynamoDBEventProperty.BisectBatchOnFunctionError`.
	BisectBatchOnFunctionError interface{} `json:"bisectBatchOnFunctionError" yaml:"bisectBatchOnFunctionError"`
	// `CfnFunction.DynamoDBEventProperty.DestinationConfig`.
	DestinationConfig interface{} `json:"destinationConfig" yaml:"destinationConfig"`
	// `CfnFunction.DynamoDBEventProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// `CfnFunction.DynamoDBEventProperty.MaximumBatchingWindowInSeconds`.
	MaximumBatchingWindowInSeconds *float64 `json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// `CfnFunction.DynamoDBEventProperty.MaximumRecordAgeInSeconds`.
	MaximumRecordAgeInSeconds *float64 `json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// `CfnFunction.DynamoDBEventProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// `CfnFunction.DynamoDBEventProperty.ParallelizationFactor`.
	ParallelizationFactor *float64 `json:"parallelizationFactor" yaml:"parallelizationFactor"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   emptySAMPTProperty := &emptySAMPTProperty{
//   }
//
type CfnFunction_EmptySAMPTProperty struct {
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var pattern interface{}
//   eventBridgeRuleEventProperty := &eventBridgeRuleEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	eventBusName: jsii.String("eventBusName"),
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnFunction_EventBridgeRuleEventProperty struct {
	// `CfnFunction.EventBridgeRuleEventProperty.Pattern`.
	Pattern interface{} `json:"pattern" yaml:"pattern"`
	// `CfnFunction.EventBridgeRuleEventProperty.EventBusName`.
	EventBusName *string `json:"eventBusName" yaml:"eventBusName"`
	// `CfnFunction.EventBridgeRuleEventProperty.Input`.
	Input *string `json:"input" yaml:"input"`
	// `CfnFunction.EventBridgeRuleEventProperty.InputPath`.
	InputPath *string `json:"inputPath" yaml:"inputPath"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   eventInvokeConfigProperty := &eventInvokeConfigProperty{
//   	destinationConfig: &eventInvokeDestinationConfigProperty{
//   		onFailure: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   		onSuccess: &destinationProperty{
//   			destination: jsii.String("destination"),
//
//   			// the properties below are optional
//   			type: jsii.String("type"),
//   		},
//   	},
//   	maximumEventAgeInSeconds: jsii.Number(123),
//   	maximumRetryAttempts: jsii.Number(123),
//   }
//
type CfnFunction_EventInvokeConfigProperty struct {
	// `CfnFunction.EventInvokeConfigProperty.DestinationConfig`.
	DestinationConfig interface{} `json:"destinationConfig" yaml:"destinationConfig"`
	// `CfnFunction.EventInvokeConfigProperty.MaximumEventAgeInSeconds`.
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds" yaml:"maximumEventAgeInSeconds"`
	// `CfnFunction.EventInvokeConfigProperty.MaximumRetryAttempts`.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   eventInvokeDestinationConfigProperty := &eventInvokeDestinationConfigProperty{
//   	onFailure: &destinationProperty{
//   		destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   	},
//   	onSuccess: &destinationProperty{
//   		destination: jsii.String("destination"),
//
//   		// the properties below are optional
//   		type: jsii.String("type"),
//   	},
//   }
//
type CfnFunction_EventInvokeDestinationConfigProperty struct {
	// `CfnFunction.EventInvokeDestinationConfigProperty.OnFailure`.
	OnFailure interface{} `json:"onFailure" yaml:"onFailure"`
	// `CfnFunction.EventInvokeDestinationConfigProperty.OnSuccess`.
	OnSuccess interface{} `json:"onSuccess" yaml:"onSuccess"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   eventSourceProperty := &eventSourceProperty{
//   	properties: &s3EventProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnFunction_EventSourceProperty struct {
	// `CfnFunction.EventSourceProperty.Properties`.
	Properties interface{} `json:"properties" yaml:"properties"`
	// `CfnFunction.EventSourceProperty.Type`.
	Type *string `json:"type" yaml:"type"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   fileSystemConfigProperty := &fileSystemConfigProperty{
//   	arn: jsii.String("arn"),
//   	localMountPath: jsii.String("localMountPath"),
//   }
//
type CfnFunction_FileSystemConfigProperty struct {
	// `CfnFunction.FileSystemConfigProperty.Arn`.
	Arn *string `json:"arn" yaml:"arn"`
	// `CfnFunction.FileSystemConfigProperty.LocalMountPath`.
	LocalMountPath *string `json:"localMountPath" yaml:"localMountPath"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   functionEnvironmentProperty := &functionEnvironmentProperty{
//   	variables: map[string]*string{
//   		"variablesKey": jsii.String("variables"),
//   	},
//   }
//
type CfnFunction_FunctionEnvironmentProperty struct {
	// `CfnFunction.FunctionEnvironmentProperty.Variables`.
	Variables interface{} `json:"variables" yaml:"variables"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   functionSAMPTProperty := &functionSAMPTProperty{
//   	functionName: jsii.String("functionName"),
//   }
//
type CfnFunction_FunctionSAMPTProperty struct {
	// `CfnFunction.FunctionSAMPTProperty.FunctionName`.
	FunctionName *string `json:"functionName" yaml:"functionName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   hooksProperty := &hooksProperty{
//   	postTraffic: jsii.String("postTraffic"),
//   	preTraffic: jsii.String("preTraffic"),
//   }
//
type CfnFunction_HooksProperty struct {
	// `CfnFunction.HooksProperty.PostTraffic`.
	PostTraffic *string `json:"postTraffic" yaml:"postTraffic"`
	// `CfnFunction.HooksProperty.PreTraffic`.
	PreTraffic *string `json:"preTraffic" yaml:"preTraffic"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var statement interface{}
//   iAMPolicyDocumentProperty := map[string]interface{}{
//   	"statement": statement,
//   }
//
type CfnFunction_IAMPolicyDocumentProperty struct {
	// `CfnFunction.IAMPolicyDocumentProperty.Statement`.
	Statement interface{} `json:"statement" yaml:"statement"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   identitySAMPTProperty := &identitySAMPTProperty{
//   	identityName: jsii.String("identityName"),
//   }
//
type CfnFunction_IdentitySAMPTProperty struct {
	// `CfnFunction.IdentitySAMPTProperty.IdentityName`.
	IdentityName *string `json:"identityName" yaml:"identityName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   imageConfigProperty := &imageConfigProperty{
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type CfnFunction_ImageConfigProperty struct {
	// `CfnFunction.ImageConfigProperty.Command`.
	Command *[]*string `json:"command" yaml:"command"`
	// `CfnFunction.ImageConfigProperty.EntryPoint`.
	EntryPoint *[]*string `json:"entryPoint" yaml:"entryPoint"`
	// `CfnFunction.ImageConfigProperty.WorkingDirectory`.
	WorkingDirectory *string `json:"workingDirectory" yaml:"workingDirectory"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   ioTRuleEventProperty := &ioTRuleEventProperty{
//   	sql: jsii.String("sql"),
//
//   	// the properties below are optional
//   	awsIotSqlVersion: jsii.String("awsIotSqlVersion"),
//   }
//
type CfnFunction_IoTRuleEventProperty struct {
	// `CfnFunction.IoTRuleEventProperty.Sql`.
	Sql *string `json:"sql" yaml:"sql"`
	// `CfnFunction.IoTRuleEventProperty.AwsIotSqlVersion`.
	AwsIotSqlVersion *string `json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   keySAMPTProperty := &keySAMPTProperty{
//   	keyId: jsii.String("keyId"),
//   }
//
type CfnFunction_KeySAMPTProperty struct {
	// `CfnFunction.KeySAMPTProperty.KeyId`.
	KeyId *string `json:"keyId" yaml:"keyId"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   kinesisEventProperty := &kinesisEventProperty{
//   	startingPosition: jsii.String("startingPosition"),
//   	stream: jsii.String("stream"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnFunction_KinesisEventProperty struct {
	// `CfnFunction.KinesisEventProperty.StartingPosition`.
	StartingPosition *string `json:"startingPosition" yaml:"startingPosition"`
	// `CfnFunction.KinesisEventProperty.Stream`.
	Stream *string `json:"stream" yaml:"stream"`
	// `CfnFunction.KinesisEventProperty.BatchSize`.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.KinesisEventProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   logGroupSAMPTProperty := &logGroupSAMPTProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnFunction_LogGroupSAMPTProperty struct {
	// `CfnFunction.LogGroupSAMPTProperty.LogGroupName`.
	LogGroupName *string `json:"logGroupName" yaml:"logGroupName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   provisionedConcurrencyConfigProperty := &provisionedConcurrencyConfigProperty{
//   	provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   }
//
type CfnFunction_ProvisionedConcurrencyConfigProperty struct {
	// `CfnFunction.ProvisionedConcurrencyConfigProperty.ProvisionedConcurrentExecutions`.
	ProvisionedConcurrentExecutions *string `json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   queueSAMPTProperty := &queueSAMPTProperty{
//   	queueName: jsii.String("queueName"),
//   }
//
type CfnFunction_QueueSAMPTProperty struct {
	// `CfnFunction.QueueSAMPTProperty.QueueName`.
	QueueName *string `json:"queueName" yaml:"queueName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3EventProperty := &s3EventProperty{
//   	bucket: jsii.String("bucket"),
//   	events: jsii.String("events"),
//
//   	// the properties below are optional
//   	filter: &s3NotificationFilterProperty{
//   		s3Key: &s3KeyFilterProperty{
//   			rules: []interface{}{
//   				&s3KeyFilterRuleProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnFunction_S3EventProperty struct {
	// `CfnFunction.S3EventProperty.Bucket`.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// `CfnFunction.S3EventProperty.Events`.
	Events interface{} `json:"events" yaml:"events"`
	// `CfnFunction.S3EventProperty.Filter`.
	Filter interface{} `json:"filter" yaml:"filter"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3KeyFilterProperty := &s3KeyFilterProperty{
//   	rules: []interface{}{
//   		&s3KeyFilterRuleProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFunction_S3KeyFilterProperty struct {
	// `CfnFunction.S3KeyFilterProperty.Rules`.
	Rules interface{} `json:"rules" yaml:"rules"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3KeyFilterRuleProperty := &s3KeyFilterRuleProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnFunction_S3KeyFilterRuleProperty struct {
	// `CfnFunction.S3KeyFilterRuleProperty.Name`.
	Name *string `json:"name" yaml:"name"`
	// `CfnFunction.S3KeyFilterRuleProperty.Value`.
	Value *string `json:"value" yaml:"value"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnFunction_S3LocationProperty struct {
	// `CfnFunction.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// `CfnFunction.S3LocationProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnFunction.S3LocationProperty.Version`.
	Version *float64 `json:"version" yaml:"version"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3NotificationFilterProperty := &s3NotificationFilterProperty{
//   	s3Key: &s3KeyFilterProperty{
//   		rules: []interface{}{
//   			&s3KeyFilterRuleProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnFunction_S3NotificationFilterProperty struct {
	// `CfnFunction.S3NotificationFilterProperty.S3Key`.
	S3Key interface{} `json:"s3Key" yaml:"s3Key"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   sAMPolicyTemplateProperty := &sAMPolicyTemplateProperty{
//   	amiDescribePolicy: &emptySAMPTProperty{
//   	},
//   	cloudFormationDescribeStacksPolicy: &emptySAMPTProperty{
//   	},
//   	cloudWatchPutMetricPolicy: &emptySAMPTProperty{
//   	},
//   	dynamoDbCrudPolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbReadPolicy: &tableSAMPTProperty{
//   		tableName: jsii.String("tableName"),
//   	},
//   	dynamoDbStreamReadPolicy: &tableStreamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   		tableName: jsii.String("tableName"),
//   	},
//   	ec2DescribePolicy: &emptySAMPTProperty{
//   	},
//   	elasticsearchHttpPostPolicy: &domainSAMPTProperty{
//   		domainName: jsii.String("domainName"),
//   	},
//   	filterLogEventsPolicy: &logGroupSAMPTProperty{
//   		logGroupName: jsii.String("logGroupName"),
//   	},
//   	kinesisCrudPolicy: &streamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   	},
//   	kinesisStreamReadPolicy: &streamSAMPTProperty{
//   		streamName: jsii.String("streamName"),
//   	},
//   	kmsDecryptPolicy: &keySAMPTProperty{
//   		keyId: jsii.String("keyId"),
//   	},
//   	lambdaInvokePolicy: &functionSAMPTProperty{
//   		functionName: jsii.String("functionName"),
//   	},
//   	rekognitionDetectOnlyPolicy: &emptySAMPTProperty{
//   	},
//   	rekognitionLabelsPolicy: &emptySAMPTProperty{
//   	},
//   	rekognitionNoDataAccessPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	rekognitionReadPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	rekognitionWriteOnlyAccessPolicy: &collectionSAMPTProperty{
//   		collectionId: jsii.String("collectionId"),
//   	},
//   	s3CrudPolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	s3ReadPolicy: &bucketSAMPTProperty{
//   		bucketName: jsii.String("bucketName"),
//   	},
//   	sesBulkTemplatedCrudPolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	sesCrudPolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	sesEmailTemplateCrudPolicy: &emptySAMPTProperty{
//   	},
//   	sesSendBouncePolicy: &identitySAMPTProperty{
//   		identityName: jsii.String("identityName"),
//   	},
//   	snsCrudPolicy: &topicSAMPTProperty{
//   		topicName: jsii.String("topicName"),
//   	},
//   	snsPublishMessagePolicy: &topicSAMPTProperty{
//   		topicName: jsii.String("topicName"),
//   	},
//   	sqsPollerPolicy: &queueSAMPTProperty{
//   		queueName: jsii.String("queueName"),
//   	},
//   	sqsSendMessagePolicy: &queueSAMPTProperty{
//   		queueName: jsii.String("queueName"),
//   	},
//   	stepFunctionsExecutionPolicy: &stateMachineSAMPTProperty{
//   		stateMachineName: jsii.String("stateMachineName"),
//   	},
//   	vpcAccessPolicy: &emptySAMPTProperty{
//   	},
//   }
//
type CfnFunction_SAMPolicyTemplateProperty struct {
	// `CfnFunction.SAMPolicyTemplateProperty.AMIDescribePolicy`.
	AmiDescribePolicy interface{} `json:"amiDescribePolicy" yaml:"amiDescribePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.CloudFormationDescribeStacksPolicy`.
	CloudFormationDescribeStacksPolicy interface{} `json:"cloudFormationDescribeStacksPolicy" yaml:"cloudFormationDescribeStacksPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.CloudWatchPutMetricPolicy`.
	CloudWatchPutMetricPolicy interface{} `json:"cloudWatchPutMetricPolicy" yaml:"cloudWatchPutMetricPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBCrudPolicy`.
	DynamoDbCrudPolicy interface{} `json:"dynamoDbCrudPolicy" yaml:"dynamoDbCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBReadPolicy`.
	DynamoDbReadPolicy interface{} `json:"dynamoDbReadPolicy" yaml:"dynamoDbReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.DynamoDBStreamReadPolicy`.
	DynamoDbStreamReadPolicy interface{} `json:"dynamoDbStreamReadPolicy" yaml:"dynamoDbStreamReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.EC2DescribePolicy`.
	Ec2DescribePolicy interface{} `json:"ec2DescribePolicy" yaml:"ec2DescribePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.ElasticsearchHttpPostPolicy`.
	ElasticsearchHttpPostPolicy interface{} `json:"elasticsearchHttpPostPolicy" yaml:"elasticsearchHttpPostPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.FilterLogEventsPolicy`.
	FilterLogEventsPolicy interface{} `json:"filterLogEventsPolicy" yaml:"filterLogEventsPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KinesisCrudPolicy`.
	KinesisCrudPolicy interface{} `json:"kinesisCrudPolicy" yaml:"kinesisCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KinesisStreamReadPolicy`.
	KinesisStreamReadPolicy interface{} `json:"kinesisStreamReadPolicy" yaml:"kinesisStreamReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.KMSDecryptPolicy`.
	KmsDecryptPolicy interface{} `json:"kmsDecryptPolicy" yaml:"kmsDecryptPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.LambdaInvokePolicy`.
	LambdaInvokePolicy interface{} `json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionDetectOnlyPolicy`.
	RekognitionDetectOnlyPolicy interface{} `json:"rekognitionDetectOnlyPolicy" yaml:"rekognitionDetectOnlyPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionLabelsPolicy`.
	RekognitionLabelsPolicy interface{} `json:"rekognitionLabelsPolicy" yaml:"rekognitionLabelsPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionNoDataAccessPolicy`.
	RekognitionNoDataAccessPolicy interface{} `json:"rekognitionNoDataAccessPolicy" yaml:"rekognitionNoDataAccessPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionReadPolicy`.
	RekognitionReadPolicy interface{} `json:"rekognitionReadPolicy" yaml:"rekognitionReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.RekognitionWriteOnlyAccessPolicy`.
	RekognitionWriteOnlyAccessPolicy interface{} `json:"rekognitionWriteOnlyAccessPolicy" yaml:"rekognitionWriteOnlyAccessPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3CrudPolicy`.
	S3CrudPolicy interface{} `json:"s3CrudPolicy" yaml:"s3CrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.S3ReadPolicy`.
	S3ReadPolicy interface{} `json:"s3ReadPolicy" yaml:"s3ReadPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESBulkTemplatedCrudPolicy`.
	SesBulkTemplatedCrudPolicy interface{} `json:"sesBulkTemplatedCrudPolicy" yaml:"sesBulkTemplatedCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESCrudPolicy`.
	SesCrudPolicy interface{} `json:"sesCrudPolicy" yaml:"sesCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESEmailTemplateCrudPolicy`.
	SesEmailTemplateCrudPolicy interface{} `json:"sesEmailTemplateCrudPolicy" yaml:"sesEmailTemplateCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SESSendBouncePolicy`.
	SesSendBouncePolicy interface{} `json:"sesSendBouncePolicy" yaml:"sesSendBouncePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SNSCrudPolicy`.
	SnsCrudPolicy interface{} `json:"snsCrudPolicy" yaml:"snsCrudPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SNSPublishMessagePolicy`.
	SnsPublishMessagePolicy interface{} `json:"snsPublishMessagePolicy" yaml:"snsPublishMessagePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SQSPollerPolicy`.
	SqsPollerPolicy interface{} `json:"sqsPollerPolicy" yaml:"sqsPollerPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.SQSSendMessagePolicy`.
	SqsSendMessagePolicy interface{} `json:"sqsSendMessagePolicy" yaml:"sqsSendMessagePolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.StepFunctionsExecutionPolicy`.
	StepFunctionsExecutionPolicy interface{} `json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
	// `CfnFunction.SAMPolicyTemplateProperty.VPCAccessPolicy`.
	VpcAccessPolicy interface{} `json:"vpcAccessPolicy" yaml:"vpcAccessPolicy"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   sNSEventProperty := &sNSEventProperty{
//   	topic: jsii.String("topic"),
//   }
//
type CfnFunction_SNSEventProperty struct {
	// `CfnFunction.SNSEventProperty.Topic`.
	Topic *string `json:"topic" yaml:"topic"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   sQSEventProperty := &sQSEventProperty{
//   	queue: jsii.String("queue"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnFunction_SQSEventProperty struct {
	// `CfnFunction.SQSEventProperty.Queue`.
	Queue *string `json:"queue" yaml:"queue"`
	// `CfnFunction.SQSEventProperty.BatchSize`.
	BatchSize *float64 `json:"batchSize" yaml:"batchSize"`
	// `CfnFunction.SQSEventProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   scheduleEventProperty := &scheduleEventProperty{
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	input: jsii.String("input"),
//   }
//
type CfnFunction_ScheduleEventProperty struct {
	// `CfnFunction.ScheduleEventProperty.Schedule`.
	Schedule *string `json:"schedule" yaml:"schedule"`
	// `CfnFunction.ScheduleEventProperty.Input`.
	Input *string `json:"input" yaml:"input"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   stateMachineSAMPTProperty := &stateMachineSAMPTProperty{
//   	stateMachineName: jsii.String("stateMachineName"),
//   }
//
type CfnFunction_StateMachineSAMPTProperty struct {
	// `CfnFunction.StateMachineSAMPTProperty.StateMachineName`.
	StateMachineName *string `json:"stateMachineName" yaml:"stateMachineName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   streamSAMPTProperty := &streamSAMPTProperty{
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnFunction_StreamSAMPTProperty struct {
	// `CfnFunction.StreamSAMPTProperty.StreamName`.
	StreamName *string `json:"streamName" yaml:"streamName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   tableSAMPTProperty := &tableSAMPTProperty{
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnFunction_TableSAMPTProperty struct {
	// `CfnFunction.TableSAMPTProperty.TableName`.
	TableName *string `json:"tableName" yaml:"tableName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   tableStreamSAMPTProperty := &tableStreamSAMPTProperty{
//   	streamName: jsii.String("streamName"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnFunction_TableStreamSAMPTProperty struct {
	// `CfnFunction.TableStreamSAMPTProperty.StreamName`.
	StreamName *string `json:"streamName" yaml:"streamName"`
	// `CfnFunction.TableStreamSAMPTProperty.TableName`.
	TableName *string `json:"tableName" yaml:"tableName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   topicSAMPTProperty := &topicSAMPTProperty{
//   	topicName: jsii.String("topicName"),
//   }
//
type CfnFunction_TopicSAMPTProperty struct {
	// `CfnFunction.TopicSAMPTProperty.TopicName`.
	TopicName *string `json:"topicName" yaml:"topicName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnFunction_VpcConfigProperty struct {
	// `CfnFunction.VpcConfigProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnFunction.VpcConfigProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

// Properties for defining a `CfnFunction`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var assumeRolePolicyDocument interface{}
//   cfnFunctionProps := &cfnFunctionProps{
//   	architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	assumeRolePolicyDocument: assumeRolePolicyDocument,
//   	autoPublishAlias: jsii.String("autoPublishAlias"),
//   	autoPublishCodeSha256: jsii.String("autoPublishCodeSha256"),
//   	codeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	codeUri: jsii.String("codeUri"),
//   	deadLetterQueue: &deadLetterQueueProperty{
//   		targetArn: jsii.String("targetArn"),
//   		type: jsii.String("type"),
//   	},
//   	deploymentPreference: &deploymentPreferenceProperty{
//   		enabled: jsii.Boolean(false),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		hooks: &hooksProperty{
//   			postTraffic: jsii.String("postTraffic"),
//   			preTraffic: jsii.String("preTraffic"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	environment: &functionEnvironmentProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	eventInvokeConfig: &eventInvokeConfigProperty{
//   		destinationConfig: &eventInvokeDestinationConfigProperty{
//   			onFailure: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   			onSuccess: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   		},
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &S3EventProperty{
//   				"variables": map[string]*string{
//   					"variablesKey": jsii.String("variables"),
//   				},
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	fileSystemConfigs: []interface{}{
//   		&fileSystemConfigProperty{
//   			arn: jsii.String("arn"),
//   			localMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	functionName: jsii.String("functionName"),
//   	handler: jsii.String("handler"),
//   	imageConfig: &imageConfigProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	imageUri: jsii.String("imageUri"),
//   	inlineCode: jsii.String("inlineCode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	memorySize: jsii.Number(123),
//   	packageType: jsii.String("packageType"),
//   	permissionsBoundary: jsii.String("permissionsBoundary"),
//   	policies: jsii.String("policies"),
//   	provisionedConcurrencyConfig: &provisionedConcurrencyConfigProperty{
//   		provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   	},
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	role: jsii.String("role"),
//   	runtime: jsii.String("runtime"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeout: jsii.Number(123),
//   	tracing: jsii.String("tracing"),
//   	versionDescription: jsii.String("versionDescription"),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnFunctionProps struct {
	// `AWS::Serverless::Function.Architectures`.
	Architectures *[]*string `json:"architectures" yaml:"architectures"`
	// `AWS::Serverless::Function.AssumeRolePolicyDocument`.
	AssumeRolePolicyDocument interface{} `json:"assumeRolePolicyDocument" yaml:"assumeRolePolicyDocument"`
	// `AWS::Serverless::Function.AutoPublishAlias`.
	AutoPublishAlias *string `json:"autoPublishAlias" yaml:"autoPublishAlias"`
	// `AWS::Serverless::Function.AutoPublishCodeSha256`.
	AutoPublishCodeSha256 *string `json:"autoPublishCodeSha256" yaml:"autoPublishCodeSha256"`
	// `AWS::Serverless::Function.CodeSigningConfigArn`.
	CodeSigningConfigArn *string `json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// `AWS::Serverless::Function.CodeUri`.
	CodeUri interface{} `json:"codeUri" yaml:"codeUri"`
	// `AWS::Serverless::Function.DeadLetterQueue`.
	DeadLetterQueue interface{} `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// `AWS::Serverless::Function.DeploymentPreference`.
	DeploymentPreference interface{} `json:"deploymentPreference" yaml:"deploymentPreference"`
	// `AWS::Serverless::Function.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::Serverless::Function.Environment`.
	Environment interface{} `json:"environment" yaml:"environment"`
	// `AWS::Serverless::Function.EventInvokeConfig`.
	EventInvokeConfig interface{} `json:"eventInvokeConfig" yaml:"eventInvokeConfig"`
	// `AWS::Serverless::Function.Events`.
	Events interface{} `json:"events" yaml:"events"`
	// `AWS::Serverless::Function.FileSystemConfigs`.
	FileSystemConfigs interface{} `json:"fileSystemConfigs" yaml:"fileSystemConfigs"`
	// `AWS::Serverless::Function.FunctionName`.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// `AWS::Serverless::Function.Handler`.
	Handler *string `json:"handler" yaml:"handler"`
	// `AWS::Serverless::Function.ImageConfig`.
	ImageConfig interface{} `json:"imageConfig" yaml:"imageConfig"`
	// `AWS::Serverless::Function.ImageUri`.
	ImageUri *string `json:"imageUri" yaml:"imageUri"`
	// `AWS::Serverless::Function.InlineCode`.
	InlineCode *string `json:"inlineCode" yaml:"inlineCode"`
	// `AWS::Serverless::Function.KmsKeyArn`.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// `AWS::Serverless::Function.Layers`.
	Layers *[]*string `json:"layers" yaml:"layers"`
	// `AWS::Serverless::Function.MemorySize`.
	MemorySize *float64 `json:"memorySize" yaml:"memorySize"`
	// `AWS::Serverless::Function.PackageType`.
	PackageType *string `json:"packageType" yaml:"packageType"`
	// `AWS::Serverless::Function.PermissionsBoundary`.
	PermissionsBoundary *string `json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// `AWS::Serverless::Function.Policies`.
	Policies interface{} `json:"policies" yaml:"policies"`
	// `AWS::Serverless::Function.ProvisionedConcurrencyConfig`.
	ProvisionedConcurrencyConfig interface{} `json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// `AWS::Serverless::Function.ReservedConcurrentExecutions`.
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// `AWS::Serverless::Function.Role`.
	Role *string `json:"role" yaml:"role"`
	// `AWS::Serverless::Function.Runtime`.
	Runtime *string `json:"runtime" yaml:"runtime"`
	// `AWS::Serverless::Function.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// `AWS::Serverless::Function.Timeout`.
	Timeout *float64 `json:"timeout" yaml:"timeout"`
	// `AWS::Serverless::Function.Tracing`.
	Tracing *string `json:"tracing" yaml:"tracing"`
	// `AWS::Serverless::Function.VersionDescription`.
	VersionDescription *string `json:"versionDescription" yaml:"versionDescription"`
	// `AWS::Serverless::Function.VpcConfig`.
	VpcConfig interface{} `json:"vpcConfig" yaml:"vpcConfig"`
}

// A CloudFormation `AWS::Serverless::HttpApi`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//   cfnHttpApi := sam.NewCfnHttpApi(this, jsii.String("MyCfnHttpApi"), &cfnHttpApiProps{
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &httpApiAuthProperty{
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	corsConfiguration: jsii.Boolean(false),
//   	defaultRouteSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	domain: &httpApiDomainConfigurationProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   		endpointConfiguration: jsii.String("endpointConfiguration"),
//   		mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   			truststoreUri: jsii.String("truststoreUri"),
//   			truststoreVersion: jsii.Boolean(false),
//   		},
//   		route53: &route53ConfigurationProperty{
//   			distributedDomainName: jsii.String("distributedDomainName"),
//   			evaluateTargetHealth: jsii.Boolean(false),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			ipV6: jsii.Boolean(false),
//   		},
//   		securityPolicy: jsii.String("securityPolicy"),
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	routeSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	stageName: jsii.String("stageName"),
//   	stageVariables: map[string]*string{
//   		"stageVariablesKey": jsii.String("stageVariables"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnHttpApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::Serverless::HttpApi.AccessLogSetting`.
	AccessLogSetting() interface{}
	SetAccessLogSetting(val interface{})
	// `AWS::Serverless::HttpApi.Auth`.
	Auth() interface{}
	SetAuth(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Serverless::HttpApi.CorsConfiguration`.
	CorsConfiguration() interface{}
	SetCorsConfiguration(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Serverless::HttpApi.DefaultRouteSettings`.
	DefaultRouteSettings() interface{}
	SetDefaultRouteSettings(val interface{})
	// `AWS::Serverless::HttpApi.DefinitionBody`.
	DefinitionBody() interface{}
	SetDefinitionBody(val interface{})
	// `AWS::Serverless::HttpApi.DefinitionUri`.
	DefinitionUri() interface{}
	SetDefinitionUri(val interface{})
	// `AWS::Serverless::HttpApi.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::HttpApi.DisableExecuteApiEndpoint`.
	DisableExecuteApiEndpoint() interface{}
	SetDisableExecuteApiEndpoint(val interface{})
	// `AWS::Serverless::HttpApi.Domain`.
	Domain() interface{}
	SetDomain(val interface{})
	// `AWS::Serverless::HttpApi.FailOnWarnings`.
	FailOnWarnings() interface{}
	SetFailOnWarnings(val interface{})
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
	// `AWS::Serverless::HttpApi.RouteSettings`.
	RouteSettings() interface{}
	SetRouteSettings(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Serverless::HttpApi.StageName`.
	StageName() *string
	SetStageName(val *string)
	// `AWS::Serverless::HttpApi.StageVariables`.
	StageVariables() interface{}
	SetStageVariables(val interface{})
	// `AWS::Serverless::HttpApi.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnHttpApi
type jsiiProxy_CfnHttpApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnHttpApi) AccessLogSetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Auth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CorsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DefaultRouteSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRouteSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DefinitionBody() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DefinitionUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) DisableExecuteApiEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Domain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) FailOnWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) RouteSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) StageVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHttpApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::HttpApi`.
func NewCfnHttpApi(scope constructs.Construct, id *string, props *CfnHttpApiProps) CfnHttpApi {
	_init_.Initialize()

	j := jsiiProxy_CfnHttpApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnHttpApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::HttpApi`.
func NewCfnHttpApi_Override(c CfnHttpApi, scope constructs.Construct, id *string, props *CfnHttpApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnHttpApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetAccessLogSetting(val interface{}) {
	_jsii_.Set(
		j,
		"accessLogSetting",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetAuth(val interface{}) {
	_jsii_.Set(
		j,
		"auth",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetCorsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"corsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDefaultRouteSettings(val interface{}) {
	_jsii_.Set(
		j,
		"defaultRouteSettings",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDefinitionBody(val interface{}) {
	_jsii_.Set(
		j,
		"definitionBody",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDefinitionUri(val interface{}) {
	_jsii_.Set(
		j,
		"definitionUri",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDisableExecuteApiEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"disableExecuteApiEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetDomain(val interface{}) {
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetFailOnWarnings(val interface{}) {
	_jsii_.Set(
		j,
		"failOnWarnings",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetRouteSettings(val interface{}) {
	_jsii_.Set(
		j,
		"routeSettings",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_CfnHttpApi) SetStageVariables(val interface{}) {
	_jsii_.Set(
		j,
		"stageVariables",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHttpApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnHttpApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnHttpApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnHttpApi",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnHttpApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnHttpApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHttpApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnHttpApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnHttpApi_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnHttpApi",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHttpApi) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnHttpApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnHttpApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnHttpApi) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnHttpApi) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHttpApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHttpApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   accessLogSettingProperty := &accessLogSettingProperty{
//   	destinationArn: jsii.String("destinationArn"),
//   	format: jsii.String("format"),
//   }
//
type CfnHttpApi_AccessLogSettingProperty struct {
	// `CfnHttpApi.AccessLogSettingProperty.DestinationArn`.
	DestinationArn *string `json:"destinationArn" yaml:"destinationArn"`
	// `CfnHttpApi.AccessLogSettingProperty.Format`.
	Format *string `json:"format" yaml:"format"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   corsConfigurationObjectProperty := &corsConfigurationObjectProperty{
//   	allowCredentials: jsii.Boolean(false),
//   	allowHeaders: jsii.String("allowHeaders"),
//   	allowMethods: jsii.String("allowMethods"),
//   	allowOrigin: jsii.String("allowOrigin"),
//   	exposeHeaders: []*string{
//   		jsii.String("exposeHeaders"),
//   	},
//   	maxAge: jsii.String("maxAge"),
//   }
//
type CfnHttpApi_CorsConfigurationObjectProperty struct {
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowCredentials`.
	AllowCredentials interface{} `json:"allowCredentials" yaml:"allowCredentials"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowHeaders`.
	AllowHeaders *string `json:"allowHeaders" yaml:"allowHeaders"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowMethods`.
	AllowMethods *string `json:"allowMethods" yaml:"allowMethods"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.AllowOrigin`.
	AllowOrigin *string `json:"allowOrigin" yaml:"allowOrigin"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.ExposeHeaders`.
	ExposeHeaders *[]*string `json:"exposeHeaders" yaml:"exposeHeaders"`
	// `CfnHttpApi.CorsConfigurationObjectProperty.MaxAge`.
	MaxAge *string `json:"maxAge" yaml:"maxAge"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var authorizers interface{}
//   httpApiAuthProperty := &httpApiAuthProperty{
//   	authorizers: authorizers,
//   	defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   }
//
type CfnHttpApi_HttpApiAuthProperty struct {
	// `CfnHttpApi.HttpApiAuthProperty.Authorizers`.
	Authorizers interface{} `json:"authorizers" yaml:"authorizers"`
	// `CfnHttpApi.HttpApiAuthProperty.DefaultAuthorizer`.
	DefaultAuthorizer *string `json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   httpApiDomainConfigurationProperty := &httpApiDomainConfigurationProperty{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	basePath: jsii.String("basePath"),
//   	endpointConfiguration: jsii.String("endpointConfiguration"),
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.Boolean(false),
//   	},
//   	route53: &route53ConfigurationProperty{
//   		distributedDomainName: jsii.String("distributedDomainName"),
//   		evaluateTargetHealth: jsii.Boolean(false),
//   		hostedZoneId: jsii.String("hostedZoneId"),
//   		hostedZoneName: jsii.String("hostedZoneName"),
//   		ipV6: jsii.Boolean(false),
//   	},
//   	securityPolicy: jsii.String("securityPolicy"),
//   }
//
type CfnHttpApi_HttpApiDomainConfigurationProperty struct {
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.CertificateArn`.
	CertificateArn *string `json:"certificateArn" yaml:"certificateArn"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.DomainName`.
	DomainName *string `json:"domainName" yaml:"domainName"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.BasePath`.
	BasePath *string `json:"basePath" yaml:"basePath"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.EndpointConfiguration`.
	EndpointConfiguration *string `json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.MutualTlsAuthentication`.
	MutualTlsAuthentication interface{} `json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.Route53`.
	Route53 interface{} `json:"route53" yaml:"route53"`
	// `CfnHttpApi.HttpApiDomainConfigurationProperty.SecurityPolicy`.
	SecurityPolicy *string `json:"securityPolicy" yaml:"securityPolicy"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   mutualTlsAuthenticationProperty := &mutualTlsAuthenticationProperty{
//   	truststoreUri: jsii.String("truststoreUri"),
//   	truststoreVersion: jsii.Boolean(false),
//   }
//
type CfnHttpApi_MutualTlsAuthenticationProperty struct {
	// `CfnHttpApi.MutualTlsAuthenticationProperty.TruststoreUri`.
	TruststoreUri *string `json:"truststoreUri" yaml:"truststoreUri"`
	// `CfnHttpApi.MutualTlsAuthenticationProperty.TruststoreVersion`.
	TruststoreVersion interface{} `json:"truststoreVersion" yaml:"truststoreVersion"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   route53ConfigurationProperty := &route53ConfigurationProperty{
//   	distributedDomainName: jsii.String("distributedDomainName"),
//   	evaluateTargetHealth: jsii.Boolean(false),
//   	hostedZoneId: jsii.String("hostedZoneId"),
//   	hostedZoneName: jsii.String("hostedZoneName"),
//   	ipV6: jsii.Boolean(false),
//   }
//
type CfnHttpApi_Route53ConfigurationProperty struct {
	// `CfnHttpApi.Route53ConfigurationProperty.DistributedDomainName`.
	DistributedDomainName *string `json:"distributedDomainName" yaml:"distributedDomainName"`
	// `CfnHttpApi.Route53ConfigurationProperty.EvaluateTargetHealth`.
	EvaluateTargetHealth interface{} `json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// `CfnHttpApi.Route53ConfigurationProperty.HostedZoneId`.
	HostedZoneId *string `json:"hostedZoneId" yaml:"hostedZoneId"`
	// `CfnHttpApi.Route53ConfigurationProperty.HostedZoneName`.
	HostedZoneName *string `json:"hostedZoneName" yaml:"hostedZoneName"`
	// `CfnHttpApi.Route53ConfigurationProperty.IpV6`.
	IpV6 interface{} `json:"ipV6" yaml:"ipV6"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   routeSettingsProperty := &routeSettingsProperty{
//   	dataTraceEnabled: jsii.Boolean(false),
//   	detailedMetricsEnabled: jsii.Boolean(false),
//   	loggingLevel: jsii.String("loggingLevel"),
//   	throttlingBurstLimit: jsii.Number(123),
//   	throttlingRateLimit: jsii.Number(123),
//   }
//
type CfnHttpApi_RouteSettingsProperty struct {
	// `CfnHttpApi.RouteSettingsProperty.DataTraceEnabled`.
	DataTraceEnabled interface{} `json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// `CfnHttpApi.RouteSettingsProperty.DetailedMetricsEnabled`.
	DetailedMetricsEnabled interface{} `json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// `CfnHttpApi.RouteSettingsProperty.LoggingLevel`.
	LoggingLevel *string `json:"loggingLevel" yaml:"loggingLevel"`
	// `CfnHttpApi.RouteSettingsProperty.ThrottlingBurstLimit`.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// `CfnHttpApi.RouteSettingsProperty.ThrottlingRateLimit`.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	version: jsii.Number(123),
//   }
//
type CfnHttpApi_S3LocationProperty struct {
	// `CfnHttpApi.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// `CfnHttpApi.S3LocationProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnHttpApi.S3LocationProperty.Version`.
	Version *float64 `json:"version" yaml:"version"`
}

// Properties for defining a `CfnHttpApi`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var authorizers interface{}
//   var definitionBody interface{}
//   cfnHttpApiProps := &cfnHttpApiProps{
//   	accessLogSetting: &accessLogSettingProperty{
//   		destinationArn: jsii.String("destinationArn"),
//   		format: jsii.String("format"),
//   	},
//   	auth: &httpApiAuthProperty{
//   		authorizers: authorizers,
//   		defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   	},
//   	corsConfiguration: jsii.Boolean(false),
//   	defaultRouteSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	definitionBody: definitionBody,
//   	definitionUri: jsii.String("definitionUri"),
//   	description: jsii.String("description"),
//   	disableExecuteApiEndpoint: jsii.Boolean(false),
//   	domain: &httpApiDomainConfigurationProperty{
//   		certificateArn: jsii.String("certificateArn"),
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		basePath: jsii.String("basePath"),
//   		endpointConfiguration: jsii.String("endpointConfiguration"),
//   		mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   			truststoreUri: jsii.String("truststoreUri"),
//   			truststoreVersion: jsii.Boolean(false),
//   		},
//   		route53: &route53ConfigurationProperty{
//   			distributedDomainName: jsii.String("distributedDomainName"),
//   			evaluateTargetHealth: jsii.Boolean(false),
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			hostedZoneName: jsii.String("hostedZoneName"),
//   			ipV6: jsii.Boolean(false),
//   		},
//   		securityPolicy: jsii.String("securityPolicy"),
//   	},
//   	failOnWarnings: jsii.Boolean(false),
//   	routeSettings: &routeSettingsProperty{
//   		dataTraceEnabled: jsii.Boolean(false),
//   		detailedMetricsEnabled: jsii.Boolean(false),
//   		loggingLevel: jsii.String("loggingLevel"),
//   		throttlingBurstLimit: jsii.Number(123),
//   		throttlingRateLimit: jsii.Number(123),
//   	},
//   	stageName: jsii.String("stageName"),
//   	stageVariables: map[string]*string{
//   		"stageVariablesKey": jsii.String("stageVariables"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnHttpApiProps struct {
	// `AWS::Serverless::HttpApi.AccessLogSetting`.
	AccessLogSetting interface{} `json:"accessLogSetting" yaml:"accessLogSetting"`
	// `AWS::Serverless::HttpApi.Auth`.
	Auth interface{} `json:"auth" yaml:"auth"`
	// `AWS::Serverless::HttpApi.CorsConfiguration`.
	CorsConfiguration interface{} `json:"corsConfiguration" yaml:"corsConfiguration"`
	// `AWS::Serverless::HttpApi.DefaultRouteSettings`.
	DefaultRouteSettings interface{} `json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// `AWS::Serverless::HttpApi.DefinitionBody`.
	DefinitionBody interface{} `json:"definitionBody" yaml:"definitionBody"`
	// `AWS::Serverless::HttpApi.DefinitionUri`.
	DefinitionUri interface{} `json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::HttpApi.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::Serverless::HttpApi.DisableExecuteApiEndpoint`.
	DisableExecuteApiEndpoint interface{} `json:"disableExecuteApiEndpoint" yaml:"disableExecuteApiEndpoint"`
	// `AWS::Serverless::HttpApi.Domain`.
	Domain interface{} `json:"domain" yaml:"domain"`
	// `AWS::Serverless::HttpApi.FailOnWarnings`.
	FailOnWarnings interface{} `json:"failOnWarnings" yaml:"failOnWarnings"`
	// `AWS::Serverless::HttpApi.RouteSettings`.
	RouteSettings interface{} `json:"routeSettings" yaml:"routeSettings"`
	// `AWS::Serverless::HttpApi.StageName`.
	StageName *string `json:"stageName" yaml:"stageName"`
	// `AWS::Serverless::HttpApi.StageVariables`.
	StageVariables interface{} `json:"stageVariables" yaml:"stageVariables"`
	// `AWS::Serverless::HttpApi.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Serverless::LayerVersion`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cfnLayerVersion := sam.NewCfnLayerVersion(this, jsii.String("MyCfnLayerVersion"), &cfnLayerVersionProps{
//   	compatibleRuntimes: []*string{
//   		jsii.String("compatibleRuntimes"),
//   	},
//   	contentUri: jsii.String("contentUri"),
//   	description: jsii.String("description"),
//   	layerName: jsii.String("layerName"),
//   	licenseInfo: jsii.String("licenseInfo"),
//   	retentionPolicy: jsii.String("retentionPolicy"),
//   })
//
type CfnLayerVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// `AWS::Serverless::LayerVersion.CompatibleRuntimes`.
	CompatibleRuntimes() *[]*string
	SetCompatibleRuntimes(val *[]*string)
	// `AWS::Serverless::LayerVersion.ContentUri`.
	ContentUri() interface{}
	SetContentUri(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Serverless::LayerVersion.Description`.
	Description() *string
	SetDescription(val *string)
	// `AWS::Serverless::LayerVersion.LayerName`.
	LayerName() *string
	SetLayerName(val *string)
	// `AWS::Serverless::LayerVersion.LicenseInfo`.
	LicenseInfo() *string
	SetLicenseInfo(val *string)
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
	// `AWS::Serverless::LayerVersion.RetentionPolicy`.
	RetentionPolicy() *string
	SetRetentionPolicy(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnLayerVersion
type jsiiProxy_CfnLayerVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLayerVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CompatibleRuntimes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) ContentUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) LayerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) LicenseInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) RetentionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::LayerVersion`.
func NewCfnLayerVersion(scope constructs.Construct, id *string, props *CfnLayerVersionProps) CfnLayerVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnLayerVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnLayerVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::LayerVersion`.
func NewCfnLayerVersion_Override(c CfnLayerVersion, scope constructs.Construct, id *string, props *CfnLayerVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnLayerVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetCompatibleRuntimes(val *[]*string) {
	_jsii_.Set(
		j,
		"compatibleRuntimes",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetContentUri(val interface{}) {
	_jsii_.Set(
		j,
		"contentUri",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetLayerName(val *string) {
	_jsii_.Set(
		j,
		"layerName",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetLicenseInfo(val *string) {
	_jsii_.Set(
		j,
		"licenseInfo",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetRetentionPolicy(val *string) {
	_jsii_.Set(
		j,
		"retentionPolicy",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLayerVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnLayerVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLayerVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnLayerVersion",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnLayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnLayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayerVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnLayerVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnLayerVersion_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnLayerVersion",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLayerVersion) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnLayerVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnLayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnLayerVersion) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnLayerVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLayerVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnLayerVersion_S3LocationProperty struct {
	// `CfnLayerVersion.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// `CfnLayerVersion.S3LocationProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnLayerVersion.S3LocationProperty.Version`.
	Version *float64 `json:"version" yaml:"version"`
}

// Properties for defining a `CfnLayerVersion`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cfnLayerVersionProps := &cfnLayerVersionProps{
//   	compatibleRuntimes: []*string{
//   		jsii.String("compatibleRuntimes"),
//   	},
//   	contentUri: jsii.String("contentUri"),
//   	description: jsii.String("description"),
//   	layerName: jsii.String("layerName"),
//   	licenseInfo: jsii.String("licenseInfo"),
//   	retentionPolicy: jsii.String("retentionPolicy"),
//   }
//
type CfnLayerVersionProps struct {
	// `AWS::Serverless::LayerVersion.CompatibleRuntimes`.
	CompatibleRuntimes *[]*string `json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// `AWS::Serverless::LayerVersion.ContentUri`.
	ContentUri interface{} `json:"contentUri" yaml:"contentUri"`
	// `AWS::Serverless::LayerVersion.Description`.
	Description *string `json:"description" yaml:"description"`
	// `AWS::Serverless::LayerVersion.LayerName`.
	LayerName *string `json:"layerName" yaml:"layerName"`
	// `AWS::Serverless::LayerVersion.LicenseInfo`.
	LicenseInfo *string `json:"licenseInfo" yaml:"licenseInfo"`
	// `AWS::Serverless::LayerVersion.RetentionPolicy`.
	RetentionPolicy *string `json:"retentionPolicy" yaml:"retentionPolicy"`
}

// A CloudFormation `AWS::Serverless::SimpleTable`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cfnSimpleTable := sam.NewCfnSimpleTable(this, jsii.String("MyCfnSimpleTable"), &cfnSimpleTableProps{
//   	primaryKey: &primaryKeyProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		writeCapacityUnits: jsii.Number(123),
//
//   		// the properties below are optional
//   		readCapacityUnits: jsii.Number(123),
//   	},
//   	sseSpecification: &sSESpecificationProperty{
//   		sseEnabled: jsii.Boolean(false),
//   	},
//   	tableName: jsii.String("tableName"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnSimpleTable interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
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
	// `AWS::Serverless::SimpleTable.PrimaryKey`.
	PrimaryKey() interface{}
	SetPrimaryKey(val interface{})
	// `AWS::Serverless::SimpleTable.ProvisionedThroughput`.
	ProvisionedThroughput() interface{}
	SetProvisionedThroughput(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::Serverless::SimpleTable.SSESpecification`.
	SseSpecification() interface{}
	SetSseSpecification(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Serverless::SimpleTable.TableName`.
	TableName() *string
	SetTableName(val *string)
	// `AWS::Serverless::SimpleTable.Tags`.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnSimpleTable
type jsiiProxy_CfnSimpleTable struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSimpleTable) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) PrimaryKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) ProvisionedThroughput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) SseSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSimpleTable) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::SimpleTable`.
func NewCfnSimpleTable(scope constructs.Construct, id *string, props *CfnSimpleTableProps) CfnSimpleTable {
	_init_.Initialize()

	j := jsiiProxy_CfnSimpleTable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnSimpleTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::SimpleTable`.
func NewCfnSimpleTable_Override(c CfnSimpleTable, scope constructs.Construct, id *string, props *CfnSimpleTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnSimpleTable",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetPrimaryKey(val interface{}) {
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetProvisionedThroughput(val interface{}) {
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetSseSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"sseSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnSimpleTable) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSimpleTable_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnSimpleTable",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSimpleTable_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnSimpleTable",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnSimpleTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnSimpleTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSimpleTable_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnSimpleTable",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnSimpleTable_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnSimpleTable",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSimpleTable) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnSimpleTable) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnSimpleTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnSimpleTable) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnSimpleTable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSimpleTable) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSimpleTable) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   primaryKeyProperty := &primaryKeyProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnSimpleTable_PrimaryKeyProperty struct {
	// `CfnSimpleTable.PrimaryKeyProperty.Type`.
	Type *string `json:"type" yaml:"type"`
	// `CfnSimpleTable.PrimaryKeyProperty.Name`.
	Name *string `json:"name" yaml:"name"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   provisionedThroughputProperty := &provisionedThroughputProperty{
//   	writeCapacityUnits: jsii.Number(123),
//
//   	// the properties below are optional
//   	readCapacityUnits: jsii.Number(123),
//   }
//
type CfnSimpleTable_ProvisionedThroughputProperty struct {
	// `CfnSimpleTable.ProvisionedThroughputProperty.WriteCapacityUnits`.
	WriteCapacityUnits *float64 `json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
	// `CfnSimpleTable.ProvisionedThroughputProperty.ReadCapacityUnits`.
	ReadCapacityUnits *float64 `json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   sSESpecificationProperty := &sSESpecificationProperty{
//   	sseEnabled: jsii.Boolean(false),
//   }
//
type CfnSimpleTable_SSESpecificationProperty struct {
	// `CfnSimpleTable.SSESpecificationProperty.SSEEnabled`.
	SseEnabled interface{} `json:"sseEnabled" yaml:"sseEnabled"`
}

// Properties for defining a `CfnSimpleTable`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cfnSimpleTableProps := &cfnSimpleTableProps{
//   	primaryKey: &primaryKeyProperty{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		writeCapacityUnits: jsii.Number(123),
//
//   		// the properties below are optional
//   		readCapacityUnits: jsii.Number(123),
//   	},
//   	sseSpecification: &sSESpecificationProperty{
//   		sseEnabled: jsii.Boolean(false),
//   	},
//   	tableName: jsii.String("tableName"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSimpleTableProps struct {
	// `AWS::Serverless::SimpleTable.PrimaryKey`.
	PrimaryKey interface{} `json:"primaryKey" yaml:"primaryKey"`
	// `AWS::Serverless::SimpleTable.ProvisionedThroughput`.
	ProvisionedThroughput interface{} `json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// `AWS::Serverless::SimpleTable.SSESpecification`.
	SseSpecification interface{} `json:"sseSpecification" yaml:"sseSpecification"`
	// `AWS::Serverless::SimpleTable.TableName`.
	TableName *string `json:"tableName" yaml:"tableName"`
	// `AWS::Serverless::SimpleTable.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Serverless::StateMachine`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var definition interface{}
//   cfnStateMachine := sam.NewCfnStateMachine(this, jsii.String("MyCfnStateMachine"), &cfnStateMachineProps{
//   	definition: definition,
//   	definitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	definitionUri: jsii.String("definitionUri"),
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &CloudWatchEventEventProperty{
//   				"method": jsii.String("method"),
//   				"path": jsii.String("path"),
//
//   				// the properties below are optional
//   				"restApiId": jsii.String("restApiId"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	logging: &loggingConfigurationProperty{
//   		destinations: []interface{}{
//   			&logDestinationProperty{
//   				cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   					logGroupArn: jsii.String("logGroupArn"),
//   				},
//   			},
//   		},
//   		includeExecutionData: jsii.Boolean(false),
//   		level: jsii.String("level"),
//   	},
//   	name: jsii.String("name"),
//   	permissionsBoundaries: jsii.String("permissionsBoundaries"),
//   	policies: jsii.String("policies"),
//   	role: jsii.String("role"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracing: &tracingConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	type: jsii.String("type"),
//   })
//
type CfnStateMachine interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// `AWS::Serverless::StateMachine.Definition`.
	Definition() interface{}
	SetDefinition(val interface{})
	// `AWS::Serverless::StateMachine.DefinitionSubstitutions`.
	DefinitionSubstitutions() interface{}
	SetDefinitionSubstitutions(val interface{})
	// `AWS::Serverless::StateMachine.DefinitionUri`.
	DefinitionUri() interface{}
	SetDefinitionUri(val interface{})
	// `AWS::Serverless::StateMachine.Events`.
	Events() interface{}
	SetEvents(val interface{})
	// `AWS::Serverless::StateMachine.Logging`.
	Logging() interface{}
	SetLogging(val interface{})
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
	// `AWS::Serverless::StateMachine.Name`.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// `AWS::Serverless::StateMachine.PermissionsBoundaries`.
	PermissionsBoundaries() *string
	SetPermissionsBoundaries(val *string)
	// `AWS::Serverless::StateMachine.Policies`.
	Policies() interface{}
	SetPolicies(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// `AWS::Serverless::StateMachine.Role`.
	Role() *string
	SetRole(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// `AWS::Serverless::StateMachine.Tags`.
	Tags() awscdk.TagManager
	// `AWS::Serverless::StateMachine.Tracing`.
	Tracing() interface{}
	SetTracing(val interface{})
	// `AWS::Serverless::StateMachine.Type`.
	Type() *string
	SetType(val *string)
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnStateMachine
type jsiiProxy_CfnStateMachine struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStateMachine) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Definition() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) DefinitionSubstitutions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionSubstitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) DefinitionUri() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"definitionUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Events() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Logging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) PermissionsBoundaries() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsBoundaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Policies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Tracing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachine) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Serverless::StateMachine`.
func NewCfnStateMachine(scope constructs.Construct, id *string, props *CfnStateMachineProps) CfnStateMachine {
	_init_.Initialize()

	j := jsiiProxy_CfnStateMachine{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnStateMachine",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Serverless::StateMachine`.
func NewCfnStateMachine_Override(c CfnStateMachine, scope constructs.Construct, id *string, props *CfnStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sam.CfnStateMachine",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetDefinition(val interface{}) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetDefinitionSubstitutions(val interface{}) {
	_jsii_.Set(
		j,
		"definitionSubstitutions",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetDefinitionUri(val interface{}) {
	_jsii_.Set(
		j,
		"definitionUri",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetEvents(val interface{}) {
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetLogging(val interface{}) {
	_jsii_.Set(
		j,
		"logging",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetPermissionsBoundaries(val *string) {
	_jsii_.Set(
		j,
		"permissionsBoundaries",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"policies",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetTracing(val interface{}) {
	_jsii_.Set(
		j,
		"tracing",
		val,
	)
}

func (j *jsiiProxy_CfnStateMachine) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnStateMachine_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnStateMachine",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnStateMachine_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnStateMachine",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func CfnStateMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_sam.CfnStateMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStateMachine_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnStateMachine",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func CfnStateMachine_REQUIRED_TRANSFORM() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_sam.CfnStateMachine",
		"REQUIRED_TRANSFORM",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStateMachine) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnStateMachine) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnStateMachine) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnStateMachine) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnStateMachine) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStateMachine) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachine) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   apiEventProperty := &apiEventProperty{
//   	method: jsii.String("method"),
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnStateMachine_ApiEventProperty struct {
	// `CfnStateMachine.ApiEventProperty.Method`.
	Method *string `json:"method" yaml:"method"`
	// `CfnStateMachine.ApiEventProperty.Path`.
	Path *string `json:"path" yaml:"path"`
	// `CfnStateMachine.ApiEventProperty.RestApiId`.
	RestApiId *string `json:"restApiId" yaml:"restApiId"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var pattern interface{}
//   cloudWatchEventEventProperty := &cloudWatchEventEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	eventBusName: jsii.String("eventBusName"),
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnStateMachine_CloudWatchEventEventProperty struct {
	// `CfnStateMachine.CloudWatchEventEventProperty.Pattern`.
	Pattern interface{} `json:"pattern" yaml:"pattern"`
	// `CfnStateMachine.CloudWatchEventEventProperty.EventBusName`.
	EventBusName *string `json:"eventBusName" yaml:"eventBusName"`
	// `CfnStateMachine.CloudWatchEventEventProperty.Input`.
	Input *string `json:"input" yaml:"input"`
	// `CfnStateMachine.CloudWatchEventEventProperty.InputPath`.
	InputPath *string `json:"inputPath" yaml:"inputPath"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   cloudWatchLogsLogGroupProperty := &cloudWatchLogsLogGroupProperty{
//   	logGroupArn: jsii.String("logGroupArn"),
//   }
//
type CfnStateMachine_CloudWatchLogsLogGroupProperty struct {
	// `CfnStateMachine.CloudWatchLogsLogGroupProperty.LogGroupArn`.
	LogGroupArn *string `json:"logGroupArn" yaml:"logGroupArn"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var pattern interface{}
//   eventBridgeRuleEventProperty := &eventBridgeRuleEventProperty{
//   	pattern: pattern,
//
//   	// the properties below are optional
//   	eventBusName: jsii.String("eventBusName"),
//   	input: jsii.String("input"),
//   	inputPath: jsii.String("inputPath"),
//   }
//
type CfnStateMachine_EventBridgeRuleEventProperty struct {
	// `CfnStateMachine.EventBridgeRuleEventProperty.Pattern`.
	Pattern interface{} `json:"pattern" yaml:"pattern"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.EventBusName`.
	EventBusName *string `json:"eventBusName" yaml:"eventBusName"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.Input`.
	Input *string `json:"input" yaml:"input"`
	// `CfnStateMachine.EventBridgeRuleEventProperty.InputPath`.
	InputPath *string `json:"inputPath" yaml:"inputPath"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   eventSourceProperty := &eventSourceProperty{
//   	properties: &cloudWatchEventEventProperty{
//   		method: jsii.String("method"),
//   		path: jsii.String("path"),
//
//   		// the properties below are optional
//   		restApiId: jsii.String("restApiId"),
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnStateMachine_EventSourceProperty struct {
	// `CfnStateMachine.EventSourceProperty.Properties`.
	Properties interface{} `json:"properties" yaml:"properties"`
	// `CfnStateMachine.EventSourceProperty.Type`.
	Type *string `json:"type" yaml:"type"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   functionSAMPTProperty := &functionSAMPTProperty{
//   	functionName: jsii.String("functionName"),
//   }
//
type CfnStateMachine_FunctionSAMPTProperty struct {
	// `CfnStateMachine.FunctionSAMPTProperty.FunctionName`.
	FunctionName *string `json:"functionName" yaml:"functionName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var statement interface{}
//   iAMPolicyDocumentProperty := map[string]interface{}{
//   	"statement": statement,
//   }
//
type CfnStateMachine_IAMPolicyDocumentProperty struct {
	// `CfnStateMachine.IAMPolicyDocumentProperty.Statement`.
	Statement interface{} `json:"statement" yaml:"statement"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   logDestinationProperty := &logDestinationProperty{
//   	cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   		logGroupArn: jsii.String("logGroupArn"),
//   	},
//   }
//
type CfnStateMachine_LogDestinationProperty struct {
	// `CfnStateMachine.LogDestinationProperty.CloudWatchLogsLogGroup`.
	CloudWatchLogsLogGroup interface{} `json:"cloudWatchLogsLogGroup" yaml:"cloudWatchLogsLogGroup"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   loggingConfigurationProperty := &loggingConfigurationProperty{
//   	destinations: []interface{}{
//   		&logDestinationProperty{
//   			cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   				logGroupArn: jsii.String("logGroupArn"),
//   			},
//   		},
//   	},
//   	includeExecutionData: jsii.Boolean(false),
//   	level: jsii.String("level"),
//   }
//
type CfnStateMachine_LoggingConfigurationProperty struct {
	// `CfnStateMachine.LoggingConfigurationProperty.Destinations`.
	Destinations interface{} `json:"destinations" yaml:"destinations"`
	// `CfnStateMachine.LoggingConfigurationProperty.IncludeExecutionData`.
	IncludeExecutionData interface{} `json:"includeExecutionData" yaml:"includeExecutionData"`
	// `CfnStateMachine.LoggingConfigurationProperty.Level`.
	Level *string `json:"level" yaml:"level"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//
//   	// the properties below are optional
//   	version: jsii.Number(123),
//   }
//
type CfnStateMachine_S3LocationProperty struct {
	// `CfnStateMachine.S3LocationProperty.Bucket`.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// `CfnStateMachine.S3LocationProperty.Key`.
	Key *string `json:"key" yaml:"key"`
	// `CfnStateMachine.S3LocationProperty.Version`.
	Version *float64 `json:"version" yaml:"version"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   sAMPolicyTemplateProperty := &sAMPolicyTemplateProperty{
//   	lambdaInvokePolicy: &functionSAMPTProperty{
//   		functionName: jsii.String("functionName"),
//   	},
//   	stepFunctionsExecutionPolicy: &stateMachineSAMPTProperty{
//   		stateMachineName: jsii.String("stateMachineName"),
//   	},
//   }
//
type CfnStateMachine_SAMPolicyTemplateProperty struct {
	// `CfnStateMachine.SAMPolicyTemplateProperty.LambdaInvokePolicy`.
	LambdaInvokePolicy interface{} `json:"lambdaInvokePolicy" yaml:"lambdaInvokePolicy"`
	// `CfnStateMachine.SAMPolicyTemplateProperty.StepFunctionsExecutionPolicy`.
	StepFunctionsExecutionPolicy interface{} `json:"stepFunctionsExecutionPolicy" yaml:"stepFunctionsExecutionPolicy"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   scheduleEventProperty := &scheduleEventProperty{
//   	schedule: jsii.String("schedule"),
//
//   	// the properties below are optional
//   	input: jsii.String("input"),
//   }
//
type CfnStateMachine_ScheduleEventProperty struct {
	// `CfnStateMachine.ScheduleEventProperty.Schedule`.
	Schedule *string `json:"schedule" yaml:"schedule"`
	// `CfnStateMachine.ScheduleEventProperty.Input`.
	Input *string `json:"input" yaml:"input"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   stateMachineSAMPTProperty := &stateMachineSAMPTProperty{
//   	stateMachineName: jsii.String("stateMachineName"),
//   }
//
type CfnStateMachine_StateMachineSAMPTProperty struct {
	// `CfnStateMachine.StateMachineSAMPTProperty.StateMachineName`.
	StateMachineName *string `json:"stateMachineName" yaml:"stateMachineName"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//   tracingConfigurationProperty := &tracingConfigurationProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnStateMachine_TracingConfigurationProperty struct {
	// `CfnStateMachine.TracingConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Properties for defining a `CfnStateMachine`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sam "github.com/aws/aws-cdk-go/awscdk/aws_sam"
//
//   var definition interface{}
//   cfnStateMachineProps := &cfnStateMachineProps{
//   	definition: definition,
//   	definitionSubstitutions: map[string]*string{
//   		"definitionSubstitutionsKey": jsii.String("definitionSubstitutions"),
//   	},
//   	definitionUri: jsii.String("definitionUri"),
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &CloudWatchEventEventProperty{
//   				"method": jsii.String("method"),
//   				"path": jsii.String("path"),
//
//   				// the properties below are optional
//   				"restApiId": jsii.String("restApiId"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	logging: &loggingConfigurationProperty{
//   		destinations: []interface{}{
//   			&logDestinationProperty{
//   				cloudWatchLogsLogGroup: &cloudWatchLogsLogGroupProperty{
//   					logGroupArn: jsii.String("logGroupArn"),
//   				},
//   			},
//   		},
//   		includeExecutionData: jsii.Boolean(false),
//   		level: jsii.String("level"),
//   	},
//   	name: jsii.String("name"),
//   	permissionsBoundaries: jsii.String("permissionsBoundaries"),
//   	policies: jsii.String("policies"),
//   	role: jsii.String("role"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	tracing: &tracingConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnStateMachineProps struct {
	// `AWS::Serverless::StateMachine.Definition`.
	Definition interface{} `json:"definition" yaml:"definition"`
	// `AWS::Serverless::StateMachine.DefinitionSubstitutions`.
	DefinitionSubstitutions interface{} `json:"definitionSubstitutions" yaml:"definitionSubstitutions"`
	// `AWS::Serverless::StateMachine.DefinitionUri`.
	DefinitionUri interface{} `json:"definitionUri" yaml:"definitionUri"`
	// `AWS::Serverless::StateMachine.Events`.
	Events interface{} `json:"events" yaml:"events"`
	// `AWS::Serverless::StateMachine.Logging`.
	Logging interface{} `json:"logging" yaml:"logging"`
	// `AWS::Serverless::StateMachine.Name`.
	Name *string `json:"name" yaml:"name"`
	// `AWS::Serverless::StateMachine.PermissionsBoundaries`.
	PermissionsBoundaries *string `json:"permissionsBoundaries" yaml:"permissionsBoundaries"`
	// `AWS::Serverless::StateMachine.Policies`.
	Policies interface{} `json:"policies" yaml:"policies"`
	// `AWS::Serverless::StateMachine.Role`.
	Role *string `json:"role" yaml:"role"`
	// `AWS::Serverless::StateMachine.Tags`.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// `AWS::Serverless::StateMachine.Tracing`.
	Tracing interface{} `json:"tracing" yaml:"tracing"`
	// `AWS::Serverless::StateMachine.Type`.
	Type *string `json:"type" yaml:"type"`
}

