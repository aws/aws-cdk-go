package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApiGatewayV2::Integration` resource creates an integration for an API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var requestParameters interface{}
//   var requestTemplates interface{}
//   var responseParameters interface{}
//
//   cfnIntegration := awscdk.Aws_apigatewayv2.NewCfnIntegration(this, jsii.String("MyCfnIntegration"), &CfnIntegrationProps{
//   	ApiId: jsii.String("apiId"),
//   	IntegrationType: jsii.String("integrationType"),
//
//   	// the properties below are optional
//   	ConnectionId: jsii.String("connectionId"),
//   	ConnectionType: jsii.String("connectionType"),
//   	ContentHandlingStrategy: jsii.String("contentHandlingStrategy"),
//   	CredentialsArn: jsii.String("credentialsArn"),
//   	Description: jsii.String("description"),
//   	IntegrationMethod: jsii.String("integrationMethod"),
//   	IntegrationSubtype: jsii.String("integrationSubtype"),
//   	IntegrationUri: jsii.String("integrationUri"),
//   	PassthroughBehavior: jsii.String("passthroughBehavior"),
//   	PayloadFormatVersion: jsii.String("payloadFormatVersion"),
//   	RequestParameters: requestParameters,
//   	RequestTemplates: requestTemplates,
//   	ResponseParameters: responseParameters,
//   	TemplateSelectionExpression: jsii.String("templateSelectionExpression"),
//   	TimeoutInMillis: jsii.Number(123),
//   	TlsConfig: &TlsConfigProperty{
//   		ServerNameToVerify: jsii.String("serverNameToVerify"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html
//
type CfnIntegration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The API identifier.
	ApiId() *string
	SetApiId(val *string)
	// The identifier.
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The ID of the VPC link for a private integration.
	ConnectionId() *string
	SetConnectionId(val *string)
	// The type of the network connection to the integration endpoint.
	ConnectionType() *string
	SetConnectionType(val *string)
	// Supported only for WebSocket APIs.
	ContentHandlingStrategy() *string
	SetContentHandlingStrategy(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies the credentials required for the integration, if any.
	CredentialsArn() *string
	SetCredentialsArn(val *string)
	// The description of the integration.
	Description() *string
	SetDescription(val *string)
	// Specifies the integration's HTTP method type.
	IntegrationMethod() *string
	SetIntegrationMethod(val *string)
	// Supported only for HTTP API `AWS_PROXY` integrations.
	IntegrationSubtype() *string
	SetIntegrationSubtype(val *string)
	// The integration type of an integration.
	//
	// One of the following:.
	IntegrationType() *string
	SetIntegrationType(val *string)
	// For a Lambda integration, specify the URI of a Lambda function.
	IntegrationUri() *string
	SetIntegrationUri(val *string)
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
	// Specifies the pass-through behavior for incoming requests based on the `Content-Type` header in the request, and the available mapping templates specified as the `requestTemplates` property on the `Integration` resource.
	PassthroughBehavior() *string
	SetPassthroughBehavior(val *string)
	// Specifies the format of the payload sent to an integration.
	PayloadFormatVersion() *string
	SetPayloadFormatVersion(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.
	RequestParameters() interface{}
	SetRequestParameters(val interface{})
	// Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
	RequestTemplates() interface{}
	SetRequestTemplates(val interface{})
	// Supported only for HTTP APIs.
	ResponseParameters() interface{}
	SetResponseParameters(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The template selection expression for the integration.
	TemplateSelectionExpression() *string
	SetTemplateSelectionExpression(val *string)
	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	TimeoutInMillis() *float64
	SetTimeoutInMillis(val *float64)
	// The TLS configuration for a private integration.
	TlsConfig() interface{}
	SetTlsConfig(val interface{})
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

func (j *jsiiProxy_CfnIntegration) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
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

func (j *jsiiProxy_CfnIntegration) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnIntegration(scope constructs.Construct, id *string, props *CfnIntegrationProps) CfnIntegration {
	_init_.Initialize()

	if err := validateNewCfnIntegrationParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnIntegration_Override(c CfnIntegration, scope constructs.Construct, id *string, props *CfnIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnIntegration)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetConnectionId(val *string) {
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetConnectionType(val *string) {
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetContentHandlingStrategy(val *string) {
	_jsii_.Set(
		j,
		"contentHandlingStrategy",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetCredentialsArn(val *string) {
	_jsii_.Set(
		j,
		"credentialsArn",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetIntegrationMethod(val *string) {
	_jsii_.Set(
		j,
		"integrationMethod",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetIntegrationSubtype(val *string) {
	_jsii_.Set(
		j,
		"integrationSubtype",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetIntegrationType(val *string) {
	if err := j.validateSetIntegrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationType",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetIntegrationUri(val *string) {
	_jsii_.Set(
		j,
		"integrationUri",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetPassthroughBehavior(val *string) {
	_jsii_.Set(
		j,
		"passthroughBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetPayloadFormatVersion(val *string) {
	_jsii_.Set(
		j,
		"payloadFormatVersion",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetRequestParameters(val interface{}) {
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetRequestTemplates(val interface{}) {
	_jsii_.Set(
		j,
		"requestTemplates",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetResponseParameters(val interface{}) {
	_jsii_.Set(
		j,
		"responseParameters",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetTemplateSelectionExpression(val *string) {
	_jsii_.Set(
		j,
		"templateSelectionExpression",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetTimeoutInMillis(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMillis",
		val,
	)
}

func (j *jsiiProxy_CfnIntegration)SetTlsConfig(val interface{}) {
	if err := j.validateSetTlsConfigParameters(val); err != nil {
		panic(err)
	}
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
func CfnIntegration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegration_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnIntegration_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegration_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigatewayv2.CfnIntegration",
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
func CfnIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnIntegration) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnIntegration) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIntegration) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIntegration) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnIntegration) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnIntegration) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnIntegration) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnIntegration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnIntegration) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnIntegration) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnIntegration) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnIntegration) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegration) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnIntegration) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnIntegration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnIntegration) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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

func (c *jsiiProxy_CfnIntegration) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

