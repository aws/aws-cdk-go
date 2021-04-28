package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::ApiGatewayV2::Api`.
type CfnApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiKeySelectionExpression() *string
	SetApiKeySelectionExpression(val *string)
	AttrApiEndpoint() *string
	BasePath() *string
	SetBasePath(val *string)
	Body() interface{}
	SetBody(val interface{})
	BodyS3Location() interface{}
	SetBodyS3Location(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CorsConfiguration() interface{}
	SetCorsConfiguration(val interface{})
	CreationStack() *[]*string
	CredentialsArn() *string
	SetCredentialsArn(val *string)
	Description() *string
	SetDescription(val *string)
	DisableExecuteApiEndpoint() interface{}
	SetDisableExecuteApiEndpoint(val interface{})
	DisableSchemaValidation() interface{}
	SetDisableSchemaValidation(val interface{})
	FailOnWarnings() interface{}
	SetFailOnWarnings(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	ProtocolType() *string
	SetProtocolType(val *string)
	Ref() *string
	RouteKey() *string
	SetRouteKey(val *string)
	RouteSelectionExpression() *string
	SetRouteSelectionExpression(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Target() *string
	SetTarget(val *string)
	UpdatedProperites() *map[string]interface{}
	Version() *string
	SetVersion(val *string)
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

// The jsii proxy struct for CfnApi
type jsiiProxy_CfnApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApi) ApiKeySelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeySelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) AttrApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) BasePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Body() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) BodyS3Location() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyS3Location",
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

func (j *jsiiProxy_CfnApi) CorsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"corsConfiguration",
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

func (j *jsiiProxy_CfnApi) CredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArn",
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

func (j *jsiiProxy_CfnApi) DisableExecuteApiEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableExecuteApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) DisableSchemaValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableSchemaValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) FailOnWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnWarnings",
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

func (j *jsiiProxy_CfnApi) ProtocolType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolType",
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

func (j *jsiiProxy_CfnApi) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) RouteSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeSelectionExpression",
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

func (j *jsiiProxy_CfnApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApi) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
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

func (j *jsiiProxy_CfnApi) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::Api`.
func NewCfnApi(scope constructs.Construct, id *string, props *CfnApiProps) CfnApi {
	_init_.Initialize()

	j := jsiiProxy_CfnApi{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::Api`.
func NewCfnApi_Override(c CfnApi, scope constructs.Construct, id *string, props *CfnApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApi) SetApiKeySelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"apiKeySelectionExpression",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetBasePath(val *string) {
	_jsii_.Set(
		j,
		"basePath",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetBody(val interface{}) {
	_jsii_.Set(
		j,
		"body",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetBodyS3Location(val interface{}) {
	_jsii_.Set(
		j,
		"bodyS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCorsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"corsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetCredentialsArn(val *string) {
	_jsii_.Set(
		j,
		"credentialsArn",
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

func (j *jsiiProxy_CfnApi) SetDisableExecuteApiEndpoint(val interface{}) {
	_jsii_.Set(
		j,
		"disableExecuteApiEndpoint",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetDisableSchemaValidation(val interface{}) {
	_jsii_.Set(
		j,
		"disableSchemaValidation",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetFailOnWarnings(val interface{}) {
	_jsii_.Set(
		j,
		"failOnWarnings",
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

func (j *jsiiProxy_CfnApi) SetProtocolType(val *string) {
	_jsii_.Set(
		j,
		"protocolType",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetRouteKey(val *string) {
	_jsii_.Set(
		j,
		"routeKey",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetRouteSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"routeSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_CfnApi) SetVersion(val *string) {
	_jsii_.Set(
		j,
		"version",
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
func CfnApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi",
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
func CfnApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApi",
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
		"aws-cdk-lib.aws_apigatewayv2.CfnApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApi) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApi) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApi) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApi) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
// Experimental.
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnApi) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
// Experimental.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApi_BodyS3LocationProperty struct {
	// `CfnApi.BodyS3LocationProperty.Bucket`.
	Bucket *string `json:"bucket"`
	// `CfnApi.BodyS3LocationProperty.Etag`.
	Etag *string `json:"etag"`
	// `CfnApi.BodyS3LocationProperty.Key`.
	Key *string `json:"key"`
	// `CfnApi.BodyS3LocationProperty.Version`.
	Version *string `json:"version"`
}

type CfnApi_CorsProperty struct {
	// `CfnApi.CorsProperty.AllowCredentials`.
	AllowCredentials interface{} `json:"allowCredentials"`
	// `CfnApi.CorsProperty.AllowHeaders`.
	AllowHeaders *[]*string `json:"allowHeaders"`
	// `CfnApi.CorsProperty.AllowMethods`.
	AllowMethods *[]*string `json:"allowMethods"`
	// `CfnApi.CorsProperty.AllowOrigins`.
	AllowOrigins *[]*string `json:"allowOrigins"`
	// `CfnApi.CorsProperty.ExposeHeaders`.
	ExposeHeaders *[]*string `json:"exposeHeaders"`
	// `CfnApi.CorsProperty.MaxAge`.
	MaxAge *float64 `json:"maxAge"`
}

// A CloudFormation `AWS::ApiGatewayV2::ApiGatewayManagedOverrides`.
type CfnApiGatewayManagedOverrides interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Integration() interface{}
	SetIntegration(val interface{})
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Route() interface{}
	SetRoute(val interface{})
	Stack() awscdk.Stack
	Stage() interface{}
	SetStage(val interface{})
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

// The jsii proxy struct for CfnApiGatewayManagedOverrides
type jsiiProxy_CfnApiGatewayManagedOverrides struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) Integration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) Route() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) Stage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::ApiGatewayManagedOverrides`.
func NewCfnApiGatewayManagedOverrides(scope constructs.Construct, id *string, props *CfnApiGatewayManagedOverridesProps) CfnApiGatewayManagedOverrides {
	_init_.Initialize()

	j := jsiiProxy_CfnApiGatewayManagedOverrides{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::ApiGatewayManagedOverrides`.
func NewCfnApiGatewayManagedOverrides_Override(c CfnApiGatewayManagedOverrides, scope constructs.Construct, id *string, props *CfnApiGatewayManagedOverridesProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) SetIntegration(val interface{}) {
	_jsii_.Set(
		j,
		"integration",
		val,
	)
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) SetRoute(val interface{}) {
	_jsii_.Set(
		j,
		"route",
		val,
	)
}

func (j *jsiiProxy_CfnApiGatewayManagedOverrides) SetStage(val interface{}) {
	_jsii_.Set(
		j,
		"stage",
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
func CfnApiGatewayManagedOverrides_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApiGatewayManagedOverrides_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides",
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
func CfnApiGatewayManagedOverrides_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiGatewayManagedOverrides_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiGatewayManagedOverrides",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApiGatewayManagedOverrides) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApiGatewayManagedOverrides) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnApiGatewayManagedOverrides_AccessLogSettingsProperty struct {
	// `CfnApiGatewayManagedOverrides.AccessLogSettingsProperty.DestinationArn`.
	DestinationArn *string `json:"destinationArn"`
	// `CfnApiGatewayManagedOverrides.AccessLogSettingsProperty.Format`.
	Format *string `json:"format"`
}

type CfnApiGatewayManagedOverrides_IntegrationOverridesProperty struct {
	// `CfnApiGatewayManagedOverrides.IntegrationOverridesProperty.Description`.
	Description *string `json:"description"`
	// `CfnApiGatewayManagedOverrides.IntegrationOverridesProperty.IntegrationMethod`.
	IntegrationMethod *string `json:"integrationMethod"`
	// `CfnApiGatewayManagedOverrides.IntegrationOverridesProperty.PayloadFormatVersion`.
	PayloadFormatVersion *string `json:"payloadFormatVersion"`
	// `CfnApiGatewayManagedOverrides.IntegrationOverridesProperty.TimeoutInMillis`.
	TimeoutInMillis *float64 `json:"timeoutInMillis"`
}

type CfnApiGatewayManagedOverrides_RouteOverridesProperty struct {
	// `CfnApiGatewayManagedOverrides.RouteOverridesProperty.AuthorizationScopes`.
	AuthorizationScopes *[]*string `json:"authorizationScopes"`
	// `CfnApiGatewayManagedOverrides.RouteOverridesProperty.AuthorizationType`.
	AuthorizationType *string `json:"authorizationType"`
	// `CfnApiGatewayManagedOverrides.RouteOverridesProperty.AuthorizerId`.
	AuthorizerId *string `json:"authorizerId"`
	// `CfnApiGatewayManagedOverrides.RouteOverridesProperty.OperationName`.
	OperationName *string `json:"operationName"`
	// `CfnApiGatewayManagedOverrides.RouteOverridesProperty.Target`.
	Target *string `json:"target"`
}

type CfnApiGatewayManagedOverrides_RouteSettingsProperty struct {
	// `CfnApiGatewayManagedOverrides.RouteSettingsProperty.DataTraceEnabled`.
	DataTraceEnabled interface{} `json:"dataTraceEnabled"`
	// `CfnApiGatewayManagedOverrides.RouteSettingsProperty.DetailedMetricsEnabled`.
	DetailedMetricsEnabled interface{} `json:"detailedMetricsEnabled"`
	// `CfnApiGatewayManagedOverrides.RouteSettingsProperty.LoggingLevel`.
	LoggingLevel *string `json:"loggingLevel"`
	// `CfnApiGatewayManagedOverrides.RouteSettingsProperty.ThrottlingBurstLimit`.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit"`
	// `CfnApiGatewayManagedOverrides.RouteSettingsProperty.ThrottlingRateLimit`.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit"`
}

type CfnApiGatewayManagedOverrides_StageOverridesProperty struct {
	// `CfnApiGatewayManagedOverrides.StageOverridesProperty.AccessLogSettings`.
	AccessLogSettings interface{} `json:"accessLogSettings"`
	// `CfnApiGatewayManagedOverrides.StageOverridesProperty.AutoDeploy`.
	AutoDeploy interface{} `json:"autoDeploy"`
	// `CfnApiGatewayManagedOverrides.StageOverridesProperty.DefaultRouteSettings`.
	DefaultRouteSettings interface{} `json:"defaultRouteSettings"`
	// `CfnApiGatewayManagedOverrides.StageOverridesProperty.Description`.
	Description *string `json:"description"`
	// `CfnApiGatewayManagedOverrides.StageOverridesProperty.RouteSettings`.
	RouteSettings interface{} `json:"routeSettings"`
	// `CfnApiGatewayManagedOverrides.StageOverridesProperty.StageVariables`.
	StageVariables interface{} `json:"stageVariables"`
}

// Properties for defining a `AWS::ApiGatewayV2::ApiGatewayManagedOverrides`.
type CfnApiGatewayManagedOverridesProps struct {
	// `AWS::ApiGatewayV2::ApiGatewayManagedOverrides.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::ApiGatewayManagedOverrides.Integration`.
	Integration interface{} `json:"integration"`
	// `AWS::ApiGatewayV2::ApiGatewayManagedOverrides.Route`.
	Route interface{} `json:"route"`
	// `AWS::ApiGatewayV2::ApiGatewayManagedOverrides.Stage`.
	Stage interface{} `json:"stage"`
}

// A CloudFormation `AWS::ApiGatewayV2::ApiMapping`.
type CfnApiMapping interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	ApiMappingKey() *string
	SetApiMappingKey(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Stage() *string
	SetStage(val *string)
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

// The jsii proxy struct for CfnApiMapping
type jsiiProxy_CfnApiMapping struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiMapping) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) ApiMappingKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiMapping) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::ApiMapping`.
func NewCfnApiMapping(scope constructs.Construct, id *string, props *CfnApiMappingProps) CfnApiMapping {
	_init_.Initialize()

	j := jsiiProxy_CfnApiMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::ApiMapping`.
func NewCfnApiMapping_Override(c CfnApiMapping, scope constructs.Construct, id *string, props *CfnApiMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiMapping) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnApiMapping) SetApiMappingKey(val *string) {
	_jsii_.Set(
		j,
		"apiMappingKey",
		val,
	)
}

func (j *jsiiProxy_CfnApiMapping) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnApiMapping) SetStage(val *string) {
	_jsii_.Set(
		j,
		"stage",
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
func CfnApiMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApiMapping_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMapping",
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
func CfnApiMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiMapping_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnApiMapping",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApiMapping) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApiMapping) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApiMapping) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnApiMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApiMapping) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnApiMapping) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApiMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnApiMapping) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnApiMapping) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApiMapping) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnApiMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnApiMapping) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnApiMapping) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnApiMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnApiMapping) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::ApiGatewayV2::ApiMapping`.
type CfnApiMappingProps struct {
	// `AWS::ApiGatewayV2::ApiMapping.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::ApiMapping.DomainName`.
	DomainName *string `json:"domainName"`
	// `AWS::ApiGatewayV2::ApiMapping.Stage`.
	Stage *string `json:"stage"`
	// `AWS::ApiGatewayV2::ApiMapping.ApiMappingKey`.
	ApiMappingKey *string `json:"apiMappingKey"`
}

// Properties for defining a `AWS::ApiGatewayV2::Api`.
type CfnApiProps struct {
	// `AWS::ApiGatewayV2::Api.ApiKeySelectionExpression`.
	ApiKeySelectionExpression *string `json:"apiKeySelectionExpression"`
	// `AWS::ApiGatewayV2::Api.BasePath`.
	BasePath *string `json:"basePath"`
	// `AWS::ApiGatewayV2::Api.Body`.
	Body interface{} `json:"body"`
	// `AWS::ApiGatewayV2::Api.BodyS3Location`.
	BodyS3Location interface{} `json:"bodyS3Location"`
	// `AWS::ApiGatewayV2::Api.CorsConfiguration`.
	CorsConfiguration interface{} `json:"corsConfiguration"`
	// `AWS::ApiGatewayV2::Api.CredentialsArn`.
	CredentialsArn *string `json:"credentialsArn"`
	// `AWS::ApiGatewayV2::Api.Description`.
	Description *string `json:"description"`
	// `AWS::ApiGatewayV2::Api.DisableExecuteApiEndpoint`.
	DisableExecuteApiEndpoint interface{} `json:"disableExecuteApiEndpoint"`
	// `AWS::ApiGatewayV2::Api.DisableSchemaValidation`.
	DisableSchemaValidation interface{} `json:"disableSchemaValidation"`
	// `AWS::ApiGatewayV2::Api.FailOnWarnings`.
	FailOnWarnings interface{} `json:"failOnWarnings"`
	// `AWS::ApiGatewayV2::Api.Name`.
	Name *string `json:"name"`
	// `AWS::ApiGatewayV2::Api.ProtocolType`.
	ProtocolType *string `json:"protocolType"`
	// `AWS::ApiGatewayV2::Api.RouteKey`.
	RouteKey *string `json:"routeKey"`
	// `AWS::ApiGatewayV2::Api.RouteSelectionExpression`.
	RouteSelectionExpression *string `json:"routeSelectionExpression"`
	// `AWS::ApiGatewayV2::Api.Tags`.
	Tags interface{} `json:"tags"`
	// `AWS::ApiGatewayV2::Api.Target`.
	Target *string `json:"target"`
	// `AWS::ApiGatewayV2::Api.Version`.
	Version *string `json:"version"`
}

// A CloudFormation `AWS::ApiGatewayV2::Authorizer`.
type CfnAuthorizer interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	AuthorizerCredentialsArn() *string
	SetAuthorizerCredentialsArn(val *string)
	AuthorizerPayloadFormatVersion() *string
	SetAuthorizerPayloadFormatVersion(val *string)
	AuthorizerResultTtlInSeconds() *float64
	SetAuthorizerResultTtlInSeconds(val *float64)
	AuthorizerType() *string
	SetAuthorizerType(val *string)
	AuthorizerUri() *string
	SetAuthorizerUri(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EnableSimpleResponses() interface{}
	SetEnableSimpleResponses(val interface{})
	IdentitySource() *[]*string
	SetIdentitySource(val *[]*string)
	IdentityValidationExpression() *string
	SetIdentityValidationExpression(val *string)
	JwtConfiguration() interface{}
	SetJwtConfiguration(val interface{})
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

// The jsii proxy struct for CfnAuthorizer
type jsiiProxy_CfnAuthorizer struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAuthorizer) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerCredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerCredentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerPayloadFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerPayloadFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerResultTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authorizerResultTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) AuthorizerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) EnableSimpleResponses() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSimpleResponses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) IdentitySource() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) IdentityValidationExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityValidationExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) JwtConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAuthorizer) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::Authorizer`.
func NewCfnAuthorizer(scope constructs.Construct, id *string, props *CfnAuthorizerProps) CfnAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_CfnAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::Authorizer`.
func NewCfnAuthorizer_Override(c CfnAuthorizer, scope constructs.Construct, id *string, props *CfnAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerCredentialsArn(val *string) {
	_jsii_.Set(
		j,
		"authorizerCredentialsArn",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerPayloadFormatVersion(val *string) {
	_jsii_.Set(
		j,
		"authorizerPayloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerResultTtlInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"authorizerResultTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerType(val *string) {
	_jsii_.Set(
		j,
		"authorizerType",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetAuthorizerUri(val *string) {
	_jsii_.Set(
		j,
		"authorizerUri",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetEnableSimpleResponses(val interface{}) {
	_jsii_.Set(
		j,
		"enableSimpleResponses",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetIdentitySource(val *[]*string) {
	_jsii_.Set(
		j,
		"identitySource",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetIdentityValidationExpression(val *string) {
	_jsii_.Set(
		j,
		"identityValidationExpression",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetJwtConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"jwtConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAuthorizer) SetName(val *string) {
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
// Experimental.
func CfnAuthorizer_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnAuthorizer_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer",
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
func CfnAuthorizer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAuthorizer_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnAuthorizer",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnAuthorizer) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAuthorizer) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAuthorizer) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAuthorizer) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnAuthorizer) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnAuthorizer) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnAuthorizer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAuthorizer) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAuthorizer) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAuthorizer) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnAuthorizer) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAuthorizer) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAuthorizer) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnAuthorizer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnAuthorizer) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnAuthorizer_JWTConfigurationProperty struct {
	// `CfnAuthorizer.JWTConfigurationProperty.Audience`.
	Audience *[]*string `json:"audience"`
	// `CfnAuthorizer.JWTConfigurationProperty.Issuer`.
	Issuer *string `json:"issuer"`
}

// Properties for defining a `AWS::ApiGatewayV2::Authorizer`.
type CfnAuthorizerProps struct {
	// `AWS::ApiGatewayV2::Authorizer.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerType`.
	AuthorizerType *string `json:"authorizerType"`
	// `AWS::ApiGatewayV2::Authorizer.IdentitySource`.
	IdentitySource *[]*string `json:"identitySource"`
	// `AWS::ApiGatewayV2::Authorizer.Name`.
	Name *string `json:"name"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerCredentialsArn`.
	AuthorizerCredentialsArn *string `json:"authorizerCredentialsArn"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerPayloadFormatVersion`.
	AuthorizerPayloadFormatVersion *string `json:"authorizerPayloadFormatVersion"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerResultTtlInSeconds`.
	AuthorizerResultTtlInSeconds *float64 `json:"authorizerResultTtlInSeconds"`
	// `AWS::ApiGatewayV2::Authorizer.AuthorizerUri`.
	AuthorizerUri *string `json:"authorizerUri"`
	// `AWS::ApiGatewayV2::Authorizer.EnableSimpleResponses`.
	EnableSimpleResponses interface{} `json:"enableSimpleResponses"`
	// `AWS::ApiGatewayV2::Authorizer.IdentityValidationExpression`.
	IdentityValidationExpression *string `json:"identityValidationExpression"`
	// `AWS::ApiGatewayV2::Authorizer.JwtConfiguration`.
	JwtConfiguration interface{} `json:"jwtConfiguration"`
}

// A CloudFormation `AWS::ApiGatewayV2::Deployment`.
type CfnDeployment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	StageName() *string
	SetStageName(val *string)
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

// The jsii proxy struct for CfnDeployment
type jsiiProxy_CfnDeployment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeployment) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeployment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::Deployment`.
func NewCfnDeployment(scope constructs.Construct, id *string, props *CfnDeploymentProps) CfnDeployment {
	_init_.Initialize()

	j := jsiiProxy_CfnDeployment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::Deployment`.
func NewCfnDeployment_Override(c CfnDeployment, scope constructs.Construct, id *string, props *CfnDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeployment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeployment) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnDeployment) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDeployment) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
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
func CfnDeployment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeployment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDeployment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeployment",
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
func CfnDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeployment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnDeployment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDeployment) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDeployment) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDeployment) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDeployment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDeployment) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDeployment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDeployment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDeployment) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDeployment) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDeployment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDeployment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeployment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDeployment) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnDeployment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::ApiGatewayV2::Deployment`.
type CfnDeploymentProps struct {
	// `AWS::ApiGatewayV2::Deployment.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::Deployment.Description`.
	Description *string `json:"description"`
	// `AWS::ApiGatewayV2::Deployment.StageName`.
	StageName *string `json:"stageName"`
}

// A CloudFormation `AWS::ApiGatewayV2::DomainName`.
type CfnDomainName interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrRegionalDomainName() *string
	AttrRegionalHostedZoneId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameConfigurations() interface{}
	SetDomainNameConfigurations(val interface{})
	LogicalId() *string
	MutualTlsAuthentication() interface{}
	SetMutualTlsAuthentication(val interface{})
	Node() constructs.Node
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDomainName
type jsiiProxy_CfnDomainName struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDomainName) AttrRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) AttrRegionalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrRegionalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) DomainNameConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"domainNameConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) MutualTlsAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutualTlsAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDomainName) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::DomainName`.
func NewCfnDomainName(scope constructs.Construct, id *string, props *CfnDomainNameProps) CfnDomainName {
	_init_.Initialize()

	j := jsiiProxy_CfnDomainName{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::DomainName`.
func NewCfnDomainName_Override(c CfnDomainName, scope constructs.Construct, id *string, props *CfnDomainNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDomainName) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetDomainNameConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"domainNameConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnDomainName) SetMutualTlsAuthentication(val interface{}) {
	_jsii_.Set(
		j,
		"mutualTlsAuthentication",
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
func CfnDomainName_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDomainName_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName",
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
func CfnDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDomainName_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnDomainName",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDomainName) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDomainName) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDomainName) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDomainName) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDomainName) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDomainName) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDomainName) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDomainName) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDomainName) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDomainName) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnDomainName) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDomainName) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDomainName) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnDomainName) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDomainName_DomainNameConfigurationProperty struct {
	// `CfnDomainName.DomainNameConfigurationProperty.CertificateArn`.
	CertificateArn *string `json:"certificateArn"`
	// `CfnDomainName.DomainNameConfigurationProperty.CertificateName`.
	CertificateName *string `json:"certificateName"`
	// `CfnDomainName.DomainNameConfigurationProperty.EndpointType`.
	EndpointType *string `json:"endpointType"`
	// `CfnDomainName.DomainNameConfigurationProperty.SecurityPolicy`.
	SecurityPolicy *string `json:"securityPolicy"`
}

type CfnDomainName_MutualTlsAuthenticationProperty struct {
	// `CfnDomainName.MutualTlsAuthenticationProperty.TruststoreUri`.
	TruststoreUri *string `json:"truststoreUri"`
	// `CfnDomainName.MutualTlsAuthenticationProperty.TruststoreVersion`.
	TruststoreVersion *string `json:"truststoreVersion"`
}

// Properties for defining a `AWS::ApiGatewayV2::DomainName`.
type CfnDomainNameProps struct {
	// `AWS::ApiGatewayV2::DomainName.DomainName`.
	DomainName *string `json:"domainName"`
	// `AWS::ApiGatewayV2::DomainName.DomainNameConfigurations`.
	DomainNameConfigurations interface{} `json:"domainNameConfigurations"`
	// `AWS::ApiGatewayV2::DomainName.MutualTlsAuthentication`.
	MutualTlsAuthentication interface{} `json:"mutualTlsAuthentication"`
	// `AWS::ApiGatewayV2::DomainName.Tags`.
	Tags interface{} `json:"tags"`
}

// A CloudFormation `AWS::ApiGatewayV2::Integration`.
type CfnIntegration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionType() *string
	SetConnectionType(val *string)
	ContentHandlingStrategy() *string
	SetContentHandlingStrategy(val *string)
	CreationStack() *[]*string
	CredentialsArn() *string
	SetCredentialsArn(val *string)
	Description() *string
	SetDescription(val *string)
	IntegrationMethod() *string
	SetIntegrationMethod(val *string)
	IntegrationSubtype() *string
	SetIntegrationSubtype(val *string)
	IntegrationType() *string
	SetIntegrationType(val *string)
	IntegrationUri() *string
	SetIntegrationUri(val *string)
	LogicalId() *string
	Node() constructs.Node
	PassthroughBehavior() *string
	SetPassthroughBehavior(val *string)
	PayloadFormatVersion() *string
	SetPayloadFormatVersion(val *string)
	Ref() *string
	RequestParameters() interface{}
	SetRequestParameters(val interface{})
	RequestTemplates() interface{}
	SetRequestTemplates(val interface{})
	ResponseParameters() interface{}
	SetResponseParameters(val interface{})
	Stack() awscdk.Stack
	TemplateSelectionExpression() *string
	SetTemplateSelectionExpression(val *string)
	TimeoutInMillis() *float64
	SetTimeoutInMillis(val *float64)
	TlsConfig() interface{}
	SetTlsConfig(val interface{})
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

// The jsii proxy struct for CfnIntegration
type jsiiProxy_CfnIntegration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIntegration) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) ContentHandlingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) CredentialsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) IntegrationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) IntegrationSubtype() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationSubtype",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) IntegrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) IntegrationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) PassthroughBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) PayloadFormatVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadFormatVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) RequestParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) RequestTemplates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) ResponseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) TemplateSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) TimeoutInMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) TlsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::Integration`.
func NewCfnIntegration(scope constructs.Construct, id *string, props *CfnIntegrationProps) CfnIntegration {
	_init_.Initialize()

	j := jsiiProxy_CfnIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::Integration`.
func NewCfnIntegration_Override(c CfnIntegration, scope constructs.Construct, id *string, props *CfnIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIntegration) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetConnectionId(val *string) {
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetConnectionType(val *string) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetContentHandlingStrategy(val *string) {
	_jsii_.Set(
		j,
		"contentHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetCredentialsArn(val *string) {
	_jsii_.Set(
		j,
		"credentialsArn",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetIntegrationMethod(val *string) {
	_jsii_.Set(
		j,
		"integrationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetIntegrationSubtype(val *string) {
	_jsii_.Set(
		j,
		"integrationSubtype",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetIntegrationType(val *string) {
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetIntegrationUri(val *string) {
	_jsii_.Set(
		j,
		"integrationUri",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetPassthroughBehavior(val *string) {
	_jsii_.Set(
		j,
		"passthroughBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetPayloadFormatVersion(val *string) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetRequestParameters(val interface{}) {
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetRequestTemplates(val interface{}) {
	_jsii_.Set(
		j,
		"requestTemplates",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetResponseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetTemplateSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"templateSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetTimeoutInMillis(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMillis",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration) SetTlsConfig(val interface{}) {
	_jsii_.Set(
		j,
		"tlsConfig",
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
func CfnIntegration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIntegration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
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
func CfnIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnIntegration) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnIntegration) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnIntegration) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnIntegration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnIntegration) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnIntegration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnIntegration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnIntegration) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnIntegration) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnIntegration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnIntegration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIntegration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnIntegration) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnIntegration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnIntegration_ResponseParameterListProperty struct {
	// `CfnIntegration.ResponseParameterListProperty.ResponseParameters`.
	ResponseParameters interface{} `json:"responseParameters"`
}

type CfnIntegration_ResponseParameterProperty struct {
	// `CfnIntegration.ResponseParameterProperty.Destination`.
	Destination *string `json:"destination"`
	// `CfnIntegration.ResponseParameterProperty.Source`.
	Source *string `json:"source"`
}

type CfnIntegration_TlsConfigProperty struct {
	// `CfnIntegration.TlsConfigProperty.ServerNameToVerify`.
	ServerNameToVerify *string `json:"serverNameToVerify"`
}

// Properties for defining a `AWS::ApiGatewayV2::Integration`.
type CfnIntegrationProps struct {
	// `AWS::ApiGatewayV2::Integration.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::Integration.IntegrationType`.
	IntegrationType *string `json:"integrationType"`
	// `AWS::ApiGatewayV2::Integration.ConnectionId`.
	ConnectionId *string `json:"connectionId"`
	// `AWS::ApiGatewayV2::Integration.ConnectionType`.
	ConnectionType *string `json:"connectionType"`
	// `AWS::ApiGatewayV2::Integration.ContentHandlingStrategy`.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy"`
	// `AWS::ApiGatewayV2::Integration.CredentialsArn`.
	CredentialsArn *string `json:"credentialsArn"`
	// `AWS::ApiGatewayV2::Integration.Description`.
	Description *string `json:"description"`
	// `AWS::ApiGatewayV2::Integration.IntegrationMethod`.
	IntegrationMethod *string `json:"integrationMethod"`
	// `AWS::ApiGatewayV2::Integration.IntegrationSubtype`.
	IntegrationSubtype *string `json:"integrationSubtype"`
	// `AWS::ApiGatewayV2::Integration.IntegrationUri`.
	IntegrationUri *string `json:"integrationUri"`
	// `AWS::ApiGatewayV2::Integration.PassthroughBehavior`.
	PassthroughBehavior *string `json:"passthroughBehavior"`
	// `AWS::ApiGatewayV2::Integration.PayloadFormatVersion`.
	PayloadFormatVersion *string `json:"payloadFormatVersion"`
	// `AWS::ApiGatewayV2::Integration.RequestParameters`.
	RequestParameters interface{} `json:"requestParameters"`
	// `AWS::ApiGatewayV2::Integration.RequestTemplates`.
	RequestTemplates interface{} `json:"requestTemplates"`
	// `AWS::ApiGatewayV2::Integration.ResponseParameters`.
	ResponseParameters interface{} `json:"responseParameters"`
	// `AWS::ApiGatewayV2::Integration.TemplateSelectionExpression`.
	TemplateSelectionExpression *string `json:"templateSelectionExpression"`
	// `AWS::ApiGatewayV2::Integration.TimeoutInMillis`.
	TimeoutInMillis *float64 `json:"timeoutInMillis"`
	// `AWS::ApiGatewayV2::Integration.TlsConfig`.
	TlsConfig interface{} `json:"tlsConfig"`
}

// A CloudFormation `AWS::ApiGatewayV2::IntegrationResponse`.
type CfnIntegrationResponse interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ContentHandlingStrategy() *string
	SetContentHandlingStrategy(val *string)
	CreationStack() *[]*string
	IntegrationId() *string
	SetIntegrationId(val *string)
	IntegrationResponseKey() *string
	SetIntegrationResponseKey(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ResponseParameters() interface{}
	SetResponseParameters(val interface{})
	ResponseTemplates() interface{}
	SetResponseTemplates(val interface{})
	Stack() awscdk.Stack
	TemplateSelectionExpression() *string
	SetTemplateSelectionExpression(val *string)
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

// The jsii proxy struct for CfnIntegrationResponse
type jsiiProxy_CfnIntegrationResponse struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnIntegrationResponse) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) ContentHandlingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) IntegrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) IntegrationResponseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationResponseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) ResponseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) ResponseTemplates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) TemplateSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationResponse) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::IntegrationResponse`.
func NewCfnIntegrationResponse(scope constructs.Construct, id *string, props *CfnIntegrationResponseProps) CfnIntegrationResponse {
	_init_.Initialize()

	j := jsiiProxy_CfnIntegrationResponse{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponse",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::IntegrationResponse`.
func NewCfnIntegrationResponse_Override(c CfnIntegrationResponse, scope constructs.Construct, id *string, props *CfnIntegrationResponseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponse",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIntegrationResponse) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponse) SetContentHandlingStrategy(val *string) {
	_jsii_.Set(
		j,
		"contentHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponse) SetIntegrationId(val *string) {
	_jsii_.Set(
		j,
		"integrationId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponse) SetIntegrationResponseKey(val *string) {
	_jsii_.Set(
		j,
		"integrationResponseKey",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponse) SetResponseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponse) SetResponseTemplates(val interface{}) {
	_jsii_.Set(
		j,
		"responseTemplates",
		val,
	)
}

func (j *jsiiProxy_CfnIntegrationResponse) SetTemplateSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"templateSelectionExpression",
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
func CfnIntegrationResponse_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponse",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnIntegrationResponse_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponse",
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
func CfnIntegrationResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegrationResponse_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegrationResponse",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnIntegrationResponse) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnIntegrationResponse) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnIntegrationResponse) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnIntegrationResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnIntegrationResponse) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnIntegrationResponse) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnIntegrationResponse) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnIntegrationResponse) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnIntegrationResponse) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnIntegrationResponse) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnIntegrationResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIntegrationResponse) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnIntegrationResponse) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnIntegrationResponse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnIntegrationResponse) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::ApiGatewayV2::IntegrationResponse`.
type CfnIntegrationResponseProps struct {
	// `AWS::ApiGatewayV2::IntegrationResponse.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::IntegrationResponse.IntegrationId`.
	IntegrationId *string `json:"integrationId"`
	// `AWS::ApiGatewayV2::IntegrationResponse.IntegrationResponseKey`.
	IntegrationResponseKey *string `json:"integrationResponseKey"`
	// `AWS::ApiGatewayV2::IntegrationResponse.ContentHandlingStrategy`.
	ContentHandlingStrategy *string `json:"contentHandlingStrategy"`
	// `AWS::ApiGatewayV2::IntegrationResponse.ResponseParameters`.
	ResponseParameters interface{} `json:"responseParameters"`
	// `AWS::ApiGatewayV2::IntegrationResponse.ResponseTemplates`.
	ResponseTemplates interface{} `json:"responseTemplates"`
	// `AWS::ApiGatewayV2::IntegrationResponse.TemplateSelectionExpression`.
	TemplateSelectionExpression *string `json:"templateSelectionExpression"`
}

// A CloudFormation `AWS::ApiGatewayV2::Model`.
type CfnModel interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ContentType() *string
	SetContentType(val *string)
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Schema() interface{}
	SetSchema(val interface{})
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

// The jsii proxy struct for CfnModel
type jsiiProxy_CfnModel struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnModel) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Schema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnModel) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::Model`.
func NewCfnModel(scope constructs.Construct, id *string, props *CfnModelProps) CfnModel {
	_init_.Initialize()

	j := jsiiProxy_CfnModel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::Model`.
func NewCfnModel_Override(c CfnModel, scope constructs.Construct, id *string, props *CfnModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnModel",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnModel) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnModel) SetSchema(val interface{}) {
	_jsii_.Set(
		j,
		"schema",
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
func CfnModel_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnModel",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnModel_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnModel",
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
func CfnModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnModel_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnModel",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnModel) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnModel) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnModel) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnModel) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnModel) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnModel) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnModel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnModel) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnModel) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnModel) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnModel) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnModel) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnModel) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnModel) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::ApiGatewayV2::Model`.
type CfnModelProps struct {
	// `AWS::ApiGatewayV2::Model.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::Model.Name`.
	Name *string `json:"name"`
	// `AWS::ApiGatewayV2::Model.Schema`.
	Schema interface{} `json:"schema"`
	// `AWS::ApiGatewayV2::Model.ContentType`.
	ContentType *string `json:"contentType"`
	// `AWS::ApiGatewayV2::Model.Description`.
	Description *string `json:"description"`
}

// A CloudFormation `AWS::ApiGatewayV2::Route`.
type CfnRoute interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	ApiKeyRequired() interface{}
	SetApiKeyRequired(val interface{})
	AuthorizationScopes() *[]*string
	SetAuthorizationScopes(val *[]*string)
	AuthorizationType() *string
	SetAuthorizationType(val *string)
	AuthorizerId() *string
	SetAuthorizerId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	ModelSelectionExpression() *string
	SetModelSelectionExpression(val *string)
	Node() constructs.Node
	OperationName() *string
	SetOperationName(val *string)
	Ref() *string
	RequestModels() interface{}
	SetRequestModels(val interface{})
	RequestParameters() interface{}
	SetRequestParameters(val interface{})
	RouteKey() *string
	SetRouteKey(val *string)
	RouteResponseSelectionExpression() *string
	SetRouteResponseSelectionExpression(val *string)
	Stack() awscdk.Stack
	Target() *string
	SetTarget(val *string)
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

// The jsii proxy struct for CfnRoute
type jsiiProxy_CfnRoute struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRoute) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) ApiKeyRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AuthorizationScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizationScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AuthorizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) AuthorizerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) ModelSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) RequestModels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) RequestParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) RouteResponseSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeResponseSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRoute) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::Route`.
func NewCfnRoute(scope constructs.Construct, id *string, props *CfnRouteProps) CfnRoute {
	_init_.Initialize()

	j := jsiiProxy_CfnRoute{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::Route`.
func NewCfnRoute_Override(c CfnRoute, scope constructs.Construct, id *string, props *CfnRouteProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRoute) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetApiKeyRequired(val interface{}) {
	_jsii_.Set(
		j,
		"apiKeyRequired",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetAuthorizationScopes(val *[]*string) {
	_jsii_.Set(
		j,
		"authorizationScopes",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetAuthorizationType(val *string) {
	_jsii_.Set(
		j,
		"authorizationType",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetAuthorizerId(val *string) {
	_jsii_.Set(
		j,
		"authorizerId",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetModelSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"modelSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetOperationName(val *string) {
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetRequestModels(val interface{}) {
	_jsii_.Set(
		j,
		"requestModels",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetRequestParameters(val interface{}) {
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetRouteKey(val *string) {
	_jsii_.Set(
		j,
		"routeKey",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetRouteResponseSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"routeResponseSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_CfnRoute) SetTarget(val *string) {
	_jsii_.Set(
		j,
		"target",
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
func CfnRoute_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRoute_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute",
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
func CfnRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRoute_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnRoute",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnRoute) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRoute) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRoute) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRoute) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnRoute) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnRoute) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnRoute) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRoute) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRoute) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRoute) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnRoute) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRoute) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRoute) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnRoute) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnRoute_ParameterConstraintsProperty struct {
	// `CfnRoute.ParameterConstraintsProperty.Required`.
	Required interface{} `json:"required"`
}

// Properties for defining a `AWS::ApiGatewayV2::Route`.
type CfnRouteProps struct {
	// `AWS::ApiGatewayV2::Route.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::Route.RouteKey`.
	RouteKey *string `json:"routeKey"`
	// `AWS::ApiGatewayV2::Route.ApiKeyRequired`.
	ApiKeyRequired interface{} `json:"apiKeyRequired"`
	// `AWS::ApiGatewayV2::Route.AuthorizationScopes`.
	AuthorizationScopes *[]*string `json:"authorizationScopes"`
	// `AWS::ApiGatewayV2::Route.AuthorizationType`.
	AuthorizationType *string `json:"authorizationType"`
	// `AWS::ApiGatewayV2::Route.AuthorizerId`.
	AuthorizerId *string `json:"authorizerId"`
	// `AWS::ApiGatewayV2::Route.ModelSelectionExpression`.
	ModelSelectionExpression *string `json:"modelSelectionExpression"`
	// `AWS::ApiGatewayV2::Route.OperationName`.
	OperationName *string `json:"operationName"`
	// `AWS::ApiGatewayV2::Route.RequestModels`.
	RequestModels interface{} `json:"requestModels"`
	// `AWS::ApiGatewayV2::Route.RequestParameters`.
	RequestParameters interface{} `json:"requestParameters"`
	// `AWS::ApiGatewayV2::Route.RouteResponseSelectionExpression`.
	RouteResponseSelectionExpression *string `json:"routeResponseSelectionExpression"`
	// `AWS::ApiGatewayV2::Route.Target`.
	Target *string `json:"target"`
}

// A CloudFormation `AWS::ApiGatewayV2::RouteResponse`.
type CfnRouteResponse interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	ModelSelectionExpression() *string
	SetModelSelectionExpression(val *string)
	Node() constructs.Node
	Ref() *string
	ResponseModels() interface{}
	SetResponseModels(val interface{})
	ResponseParameters() interface{}
	SetResponseParameters(val interface{})
	RouteId() *string
	SetRouteId(val *string)
	RouteResponseKey() *string
	SetRouteResponseKey(val *string)
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

// The jsii proxy struct for CfnRouteResponse
type jsiiProxy_CfnRouteResponse struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRouteResponse) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) ModelSelectionExpression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelSelectionExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) ResponseModels() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) ResponseParameters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) RouteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) RouteResponseKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeResponseKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRouteResponse) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::RouteResponse`.
func NewCfnRouteResponse(scope constructs.Construct, id *string, props *CfnRouteResponseProps) CfnRouteResponse {
	_init_.Initialize()

	j := jsiiProxy_CfnRouteResponse{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::RouteResponse`.
func NewCfnRouteResponse_Override(c CfnRouteResponse, scope constructs.Construct, id *string, props *CfnRouteResponseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRouteResponse) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnRouteResponse) SetModelSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"modelSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_CfnRouteResponse) SetResponseModels(val interface{}) {
	_jsii_.Set(
		j,
		"responseModels",
		val,
	)
}

func (j *jsiiProxy_CfnRouteResponse) SetResponseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnRouteResponse) SetRouteId(val *string) {
	_jsii_.Set(
		j,
		"routeId",
		val,
	)
}

func (j *jsiiProxy_CfnRouteResponse) SetRouteResponseKey(val *string) {
	_jsii_.Set(
		j,
		"routeResponseKey",
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
func CfnRouteResponse_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnRouteResponse_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse",
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
func CfnRouteResponse_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRouteResponse_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnRouteResponse",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnRouteResponse) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnRouteResponse) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnRouteResponse) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnRouteResponse) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnRouteResponse) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnRouteResponse) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnRouteResponse) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnRouteResponse) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnRouteResponse) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnRouteResponse) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnRouteResponse) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRouteResponse) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnRouteResponse) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnRouteResponse) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnRouteResponse) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnRouteResponse_ParameterConstraintsProperty struct {
	// `CfnRouteResponse.ParameterConstraintsProperty.Required`.
	Required interface{} `json:"required"`
}

// Properties for defining a `AWS::ApiGatewayV2::RouteResponse`.
type CfnRouteResponseProps struct {
	// `AWS::ApiGatewayV2::RouteResponse.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::RouteResponse.RouteId`.
	RouteId *string `json:"routeId"`
	// `AWS::ApiGatewayV2::RouteResponse.RouteResponseKey`.
	RouteResponseKey *string `json:"routeResponseKey"`
	// `AWS::ApiGatewayV2::RouteResponse.ModelSelectionExpression`.
	ModelSelectionExpression *string `json:"modelSelectionExpression"`
	// `AWS::ApiGatewayV2::RouteResponse.ResponseModels`.
	ResponseModels interface{} `json:"responseModels"`
	// `AWS::ApiGatewayV2::RouteResponse.ResponseParameters`.
	ResponseParameters interface{} `json:"responseParameters"`
}

// A CloudFormation `AWS::ApiGatewayV2::Stage`.
type CfnStage interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AccessLogSettings() interface{}
	SetAccessLogSettings(val interface{})
	AccessPolicyId() *string
	SetAccessPolicyId(val *string)
	ApiId() *string
	SetApiId(val *string)
	AutoDeploy() interface{}
	SetAutoDeploy(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClientCertificateId() *string
	SetClientCertificateId(val *string)
	CreationStack() *[]*string
	DefaultRouteSettings() interface{}
	SetDefaultRouteSettings(val interface{})
	DeploymentId() *string
	SetDeploymentId(val *string)
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	RouteSettings() interface{}
	SetRouteSettings(val interface{})
	Stack() awscdk.Stack
	StageName() *string
	SetStageName(val *string)
	StageVariables() interface{}
	SetStageVariables(val interface{})
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnStage
type jsiiProxy_CfnStage struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStage) AccessLogSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessLogSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) AccessPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) AutoDeploy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) ClientCertificateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) DefaultRouteSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultRouteSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) DeploymentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) RouteSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) StageVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stageVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStage) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::Stage`.
func NewCfnStage(scope constructs.Construct, id *string, props *CfnStageProps) CfnStage {
	_init_.Initialize()

	j := jsiiProxy_CfnStage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::Stage`.
func NewCfnStage_Override(c CfnStage, scope constructs.Construct, id *string, props *CfnStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStage) SetAccessLogSettings(val interface{}) {
	_jsii_.Set(
		j,
		"accessLogSettings",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetAccessPolicyId(val *string) {
	_jsii_.Set(
		j,
		"accessPolicyId",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetAutoDeploy(val interface{}) {
	_jsii_.Set(
		j,
		"autoDeploy",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetClientCertificateId(val *string) {
	_jsii_.Set(
		j,
		"clientCertificateId",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetDefaultRouteSettings(val interface{}) {
	_jsii_.Set(
		j,
		"defaultRouteSettings",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetDeploymentId(val *string) {
	_jsii_.Set(
		j,
		"deploymentId",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetRouteSettings(val interface{}) {
	_jsii_.Set(
		j,
		"routeSettings",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetStageName(val *string) {
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

func (j *jsiiProxy_CfnStage) SetStageVariables(val interface{}) {
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
// Experimental.
func CfnStage_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStage_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage",
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
func CfnStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStage_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnStage",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnStage) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnStage) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnStage) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnStage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnStage) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnStage) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnStage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnStage) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnStage) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnStage) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnStage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStage) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnStage) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnStage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnStage) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnStage_AccessLogSettingsProperty struct {
	// `CfnStage.AccessLogSettingsProperty.DestinationArn`.
	DestinationArn *string `json:"destinationArn"`
	// `CfnStage.AccessLogSettingsProperty.Format`.
	Format *string `json:"format"`
}

type CfnStage_RouteSettingsProperty struct {
	// `CfnStage.RouteSettingsProperty.DataTraceEnabled`.
	DataTraceEnabled interface{} `json:"dataTraceEnabled"`
	// `CfnStage.RouteSettingsProperty.DetailedMetricsEnabled`.
	DetailedMetricsEnabled interface{} `json:"detailedMetricsEnabled"`
	// `CfnStage.RouteSettingsProperty.LoggingLevel`.
	LoggingLevel *string `json:"loggingLevel"`
	// `CfnStage.RouteSettingsProperty.ThrottlingBurstLimit`.
	ThrottlingBurstLimit *float64 `json:"throttlingBurstLimit"`
	// `CfnStage.RouteSettingsProperty.ThrottlingRateLimit`.
	ThrottlingRateLimit *float64 `json:"throttlingRateLimit"`
}

// Properties for defining a `AWS::ApiGatewayV2::Stage`.
type CfnStageProps struct {
	// `AWS::ApiGatewayV2::Stage.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::ApiGatewayV2::Stage.StageName`.
	StageName *string `json:"stageName"`
	// `AWS::ApiGatewayV2::Stage.AccessLogSettings`.
	AccessLogSettings interface{} `json:"accessLogSettings"`
	// `AWS::ApiGatewayV2::Stage.AccessPolicyId`.
	AccessPolicyId *string `json:"accessPolicyId"`
	// `AWS::ApiGatewayV2::Stage.AutoDeploy`.
	AutoDeploy interface{} `json:"autoDeploy"`
	// `AWS::ApiGatewayV2::Stage.ClientCertificateId`.
	ClientCertificateId *string `json:"clientCertificateId"`
	// `AWS::ApiGatewayV2::Stage.DefaultRouteSettings`.
	DefaultRouteSettings interface{} `json:"defaultRouteSettings"`
	// `AWS::ApiGatewayV2::Stage.DeploymentId`.
	DeploymentId *string `json:"deploymentId"`
	// `AWS::ApiGatewayV2::Stage.Description`.
	Description *string `json:"description"`
	// `AWS::ApiGatewayV2::Stage.RouteSettings`.
	RouteSettings interface{} `json:"routeSettings"`
	// `AWS::ApiGatewayV2::Stage.StageVariables`.
	StageVariables interface{} `json:"stageVariables"`
	// `AWS::ApiGatewayV2::Stage.Tags`.
	Tags interface{} `json:"tags"`
}

// A CloudFormation `AWS::ApiGatewayV2::VpcLink`.
type CfnVpcLink interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	Stack() awscdk.Stack
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnVpcLink
type jsiiProxy_CfnVpcLink struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVpcLink) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVpcLink) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApiGatewayV2::VpcLink`.
func NewCfnVpcLink(scope constructs.Construct, id *string, props *CfnVpcLinkProps) CfnVpcLink {
	_init_.Initialize()

	j := jsiiProxy_CfnVpcLink{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLink",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApiGatewayV2::VpcLink`.
func NewCfnVpcLink_Override(c CfnVpcLink, scope constructs.Construct, id *string, props *CfnVpcLinkProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLink",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVpcLink) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnVpcLink) SetSecurityGroupIds(val *[]*string) {
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_CfnVpcLink) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
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
func CfnVpcLink_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLink",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnVpcLink_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLink",
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
func CfnVpcLink_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLink",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVpcLink_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.CfnVpcLink",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnVpcLink) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnVpcLink) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnVpcLink) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnVpcLink) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnVpcLink) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnVpcLink) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnVpcLink) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnVpcLink) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnVpcLink) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnVpcLink) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnVpcLink) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVpcLink) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnVpcLink) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnVpcLink) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnVpcLink) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::ApiGatewayV2::VpcLink`.
type CfnVpcLinkProps struct {
	// `AWS::ApiGatewayV2::VpcLink.Name`.
	Name *string `json:"name"`
	// `AWS::ApiGatewayV2::VpcLink.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `AWS::ApiGatewayV2::VpcLink.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `AWS::ApiGatewayV2::VpcLink.Tags`.
	Tags interface{} `json:"tags"`
}

