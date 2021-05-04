package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappsync/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v3"
)

// The options to add a field to an Intermediate Type.
// Experimental.
type AddFieldOptions struct {
	// The resolvable field to add.
	//
	// This option must be configured for Object, Interface,
	// Input and Union Types.
	// Experimental.
	Field IField `json:"field"`
	// The name of the field.
	//
	// This option must be configured for Object, Interface,
	// Input and Enum Types.
	// Experimental.
	FieldName *string `json:"fieldName"`
}

// Configuration for API Key authorization in AppSync.
// Experimental.
type ApiKeyConfig struct {
	// Description of API key.
	// Experimental.
	Description *string `json:"description"`
	// The time from creation time after which the API key expires.
	//
	// It must be a minimum of 1 day and a maximum of 365 days from date of creation.
	// Rounded down to the nearest hour.
	// Experimental.
	Expires awscdk.Expiration `json:"expires"`
	// Unique name of the API Key.
	// Experimental.
	Name *string `json:"name"`
}

// AppSync Functions are local functions that perform certain operations onto a backend data source.
//
// Developers can compose operations (Functions)
// and execute them in sequence with Pipeline Resolvers.
// Experimental.
type AppsyncFunction interface {
	awscdk.Resource
	IAppsyncFunction
	DataSource() BaseDataSource
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionId() *string
	FunctionName() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for AppsyncFunction
type jsiiProxy_AppsyncFunction struct {
	internal.Type__awscdkResource
	jsiiProxy_IAppsyncFunction
}

func (j *jsiiProxy_AppsyncFunction) DataSource() BaseDataSource {
	var returns BaseDataSource
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppsyncFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAppsyncFunction(scope constructs.Construct, id *string, props *AppsyncFunctionProps) AppsyncFunction {
	_init_.Initialize()

	j := jsiiProxy_AppsyncFunction{}

	_jsii_.Create(
		"monocdk.aws_appsync.AppsyncFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAppsyncFunction_Override(a AppsyncFunction, scope constructs.Construct, id *string, props *AppsyncFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.AppsyncFunction",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import Appsync Function from arn.
// Experimental.
func AppsyncFunction_FromAppsyncFunctionAttributes(scope constructs.Construct, id *string, attrs *AppsyncFunctionAttributes) IAppsyncFunction {
	_init_.Initialize()

	var returns IAppsyncFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.AppsyncFunction",
		"fromAppsyncFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AppsyncFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.AppsyncFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AppsyncFunction_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.AppsyncFunction",
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (a *jsiiProxy_AppsyncFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_AppsyncFunction) GeneratePhysicalName() *string {
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
// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
// Experimental.
func (a *jsiiProxy_AppsyncFunction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
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
func (a *jsiiProxy_AppsyncFunction) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AppsyncFunction) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
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
func (a *jsiiProxy_AppsyncFunction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (a *jsiiProxy_AppsyncFunction) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AppsyncFunction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AppsyncFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_AppsyncFunction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The attributes for imported AppSync Functions.
// Experimental.
type AppsyncFunctionAttributes struct {
	// the ARN of the AppSync function.
	// Experimental.
	FunctionArn *string `json:"functionArn"`
}

// the CDK properties for AppSync Functions.
// Experimental.
type AppsyncFunctionProps struct {
	// the name of the AppSync Function.
	// Experimental.
	Name *string `json:"name"`
	// the description for this AppSync Function.
	// Experimental.
	Description *string `json:"description"`
	// the request mapping template for the AppSync Function.
	// Experimental.
	RequestMappingTemplate MappingTemplate `json:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `json:"responseMappingTemplate"`
	// the GraphQL Api linked to this AppSync Function.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the data source linked to this AppSync Function.
	// Experimental.
	DataSource BaseDataSource `json:"dataSource"`
}

// Utility class representing the assigment of a value to an attribute.
// Experimental.
type Assign interface {
	PutInMap(map_ *string) *string
	RenderAsAssignment() *string
}

// The jsii proxy struct for Assign
type jsiiProxy_Assign struct {
	_ byte // padding
}

// Experimental.
func NewAssign(attr *string, arg *string) Assign {
	_init_.Initialize()

	j := jsiiProxy_Assign{}

	_jsii_.Create(
		"monocdk.aws_appsync.Assign",
		[]interface{}{attr, arg},
		&j,
	)

	return &j
}

// Experimental.
func NewAssign_Override(a Assign, attr *string, arg *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.Assign",
		[]interface{}{attr, arg},
		a,
	)
}

// Renders the assignment as a map element.
// Experimental.
func (a *jsiiProxy_Assign) PutInMap(map_ *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"putInMap",
		[]interface{}{map_},
		&returns,
	)

	return returns
}

// Renders the assignment as a VTL string.
// Experimental.
func (a *jsiiProxy_Assign) RenderAsAssignment() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"renderAsAssignment",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Specifies the attribute value assignments.
// Experimental.
type AttributeValues interface {
	Attribute(attr *string) AttributeValuesStep
	RenderTemplate() *string
	RenderVariables() *string
}

// The jsii proxy struct for AttributeValues
type jsiiProxy_AttributeValues struct {
	_ byte // padding
}

// Experimental.
func NewAttributeValues(container *string, assignments *[]Assign) AttributeValues {
	_init_.Initialize()

	j := jsiiProxy_AttributeValues{}

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValues",
		[]interface{}{container, assignments},
		&j,
	)

	return &j
}

// Experimental.
func NewAttributeValues_Override(a AttributeValues, container *string, assignments *[]Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValues",
		[]interface{}{container, assignments},
		a,
	)
}

// Allows assigning a value to the specified attribute.
// Experimental.
func (a *jsiiProxy_AttributeValues) Attribute(attr *string) AttributeValuesStep {
	var returns AttributeValuesStep

	_jsii_.Invoke(
		a,
		"attribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

// Renders the attribute value assingments to a VTL string.
// Experimental.
func (a *jsiiProxy_AttributeValues) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Renders the variables required for `renderTemplate`.
// Experimental.
func (a *jsiiProxy_AttributeValues) RenderVariables() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"renderVariables",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Utility class to allow assigning a value to an attribute.
// Experimental.
type AttributeValuesStep interface {
	Is(val *string) AttributeValues
}

// The jsii proxy struct for AttributeValuesStep
type jsiiProxy_AttributeValuesStep struct {
	_ byte // padding
}

// Experimental.
func NewAttributeValuesStep(attr *string, container *string, assignments *[]Assign) AttributeValuesStep {
	_init_.Initialize()

	j := jsiiProxy_AttributeValuesStep{}

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValuesStep",
		[]interface{}{attr, container, assignments},
		&j,
	)

	return &j
}

// Experimental.
func NewAttributeValuesStep_Override(a AttributeValuesStep, attr *string, container *string, assignments *[]Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.AttributeValuesStep",
		[]interface{}{attr, container, assignments},
		a,
	)
}

// Assign the value to the current attribute.
// Experimental.
func (a *jsiiProxy_AttributeValuesStep) Is(val *string) AttributeValues {
	var returns AttributeValues

	_jsii_.Invoke(
		a,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}

// Configuration of the API authorization modes.
// Experimental.
type AuthorizationConfig struct {
	// Additional authorization modes.
	// Experimental.
	AdditionalAuthorizationModes *[]*AuthorizationMode `json:"additionalAuthorizationModes"`
	// Optional authorization configuration.
	// Experimental.
	DefaultAuthorization *AuthorizationMode `json:"defaultAuthorization"`
}

// Interface to specify default or additional authorization(s).
// Experimental.
type AuthorizationMode struct {
	// One of possible four values AppSync supports.
	// See: https://docs.aws.amazon.com/appsync/latest/devguide/security.html
	//
	// Experimental.
	AuthorizationType AuthorizationType `json:"authorizationType"`
	// If authorizationType is `AuthorizationType.API_KEY`, this option can be configured.
	// Experimental.
	ApiKeyConfig *ApiKeyConfig `json:"apiKeyConfig"`
	// If authorizationType is `AuthorizationType.OIDC`, this option is required.
	// Experimental.
	OpenIdConnectConfig *OpenIdConnectConfig `json:"openIdConnectConfig"`
	// If authorizationType is `AuthorizationType.USER_POOL`, this option is required.
	// Experimental.
	UserPoolConfig *UserPoolConfig `json:"userPoolConfig"`
}

// enum with all possible values for AppSync authorization type.
// Experimental.
type AuthorizationType string

const (
	AuthorizationType_API_KEY AuthorizationType = "API_KEY"
	AuthorizationType_IAM AuthorizationType = "IAM"
	AuthorizationType_USER_POOL AuthorizationType = "USER_POOL"
	AuthorizationType_OIDC AuthorizationType = "OIDC"
)

// The authorization config in case the HTTP endpoint requires authorization.
// Experimental.
type AwsIamConfig struct {
	// The signing region for AWS IAM authorization.
	// Experimental.
	SigningRegion *string `json:"signingRegion"`
	// The signing service name for AWS IAM authorization.
	// Experimental.
	SigningServiceName *string `json:"signingServiceName"`
}

// Abstract AppSync datasource implementation.
//
// Do not use directly but use subclasses for resource backed datasources
// Experimental.
type BackedDataSource interface {
	BaseDataSource
	awsiam.IGrantable
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	Ds() CfnDataSource
	GrantPrincipal() awsiam.IPrincipal
	Name() *string
	Node() awscdk.ConstructNode
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	CreateResolver(props *BaseResolverProps) Resolver
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for BackedDataSource
type jsiiProxy_BackedDataSource struct {
	jsiiProxy_BaseDataSource
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_BackedDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackedDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewBackedDataSource_Override(b BackedDataSource, scope constructs.Construct, id *string, props *BackedDataSourceProps, extended *ExtendedDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.BackedDataSource",
		[]interface{}{scope, id, props, extended},
		b,
	)
}

func (j *jsiiProxy_BackedDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_BackedDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func BackedDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.BackedDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// creates a new appsync function for this datasource and API using the given properties.
// Experimental.
func (b *jsiiProxy_BackedDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		b,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (b *jsiiProxy_BackedDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		b,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackedDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackedDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
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
func (b *jsiiProxy_BackedDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BackedDataSource) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BackedDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (b *jsiiProxy_BackedDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BackedDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// properties for an AppSync datasource backed by a resource.
// Experimental.
type BackedDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `json:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Experimental.
	ServiceRole awsiam.IRole `json:"serviceRole"`
}

// the base properties for AppSync Functions.
// Experimental.
type BaseAppsyncFunctionProps struct {
	// the name of the AppSync Function.
	// Experimental.
	Name *string `json:"name"`
	// the description for this AppSync Function.
	// Experimental.
	Description *string `json:"description"`
	// the request mapping template for the AppSync Function.
	// Experimental.
	RequestMappingTemplate MappingTemplate `json:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `json:"responseMappingTemplate"`
}

// Abstract AppSync datasource implementation.
//
// Do not use directly but use subclasses for concrete datasources
// Experimental.
type BaseDataSource interface {
	awscdk.Construct
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	Ds() CfnDataSource
	Name() *string
	Node() awscdk.ConstructNode
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	CreateResolver(props *BaseResolverProps) Resolver
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for BaseDataSource
type jsiiProxy_BaseDataSource struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_BaseDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseDataSource_Override(b BaseDataSource, scope constructs.Construct, id *string, props *BackedDataSourceProps, extended *ExtendedDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.BaseDataSource",
		[]interface{}{scope, id, props, extended},
		b,
	)
}

func (j *jsiiProxy_BaseDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_BaseDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func BaseDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.BaseDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// creates a new appsync function for this datasource and API using the given properties.
// Experimental.
func (b *jsiiProxy_BaseDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		b,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (b *jsiiProxy_BaseDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		b,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BaseDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BaseDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
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
func (b *jsiiProxy_BaseDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (b *jsiiProxy_BaseDataSource) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BaseDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (b *jsiiProxy_BaseDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BaseDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Base properties for an AppSync datasource.
// Experimental.
type BaseDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `json:"name"`
}

// Basic properties for an AppSync resolver.
// Experimental.
type BaseResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	// Experimental.
	FieldName *string `json:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	// Experimental.
	TypeName *string `json:"typeName"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `json:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `json:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `json:"responseMappingTemplate"`
}

// Base options for GraphQL Types.
// Experimental.
type BaseTypeOptions struct {
	// property determining if this attribute is a list i.e. if true, attribute would be [Type].
	// Experimental.
	IsList *bool `json:"isList"`
	// property determining if this attribute is non-nullable i.e. if true, attribute would be Type!
	// Experimental.
	IsRequired *bool `json:"isRequired"`
	// property determining if this attribute is a non-nullable list i.e. if true, attribute would be [ Type ]! or if isRequired true, attribe would be [ Type! ]!
	// Experimental.
	IsRequiredList *bool `json:"isRequiredList"`
}

// A CloudFormation `AWS::AppSync::ApiCache`.
type CfnApiCache interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiCachingBehavior() *string
	SetApiCachingBehavior(val *string)
	ApiId() *string
	SetApiId(val *string)
	AtRestEncryptionEnabled() interface{}
	SetAtRestEncryptionEnabled(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	Stack() awscdk.Stack
	TransitEncryptionEnabled() interface{}
	SetTransitEncryptionEnabled(val interface{})
	Ttl() *float64
	SetTtl(val *float64)
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnApiCache
type jsiiProxy_CfnApiCache struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiCache) ApiCachingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiCachingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) AtRestEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"atRestEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiCache) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::ApiCache`.
func NewCfnApiCache(scope awscdk.Construct, id *string, props *CfnApiCacheProps) CfnApiCache {
	_init_.Initialize()

	j := jsiiProxy_CfnApiCache{}

	_jsii_.Create(
		"monocdk.aws_appsync.CfnApiCache",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::ApiCache`.
func NewCfnApiCache_Override(c CfnApiCache, scope awscdk.Construct, id *string, props *CfnApiCacheProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.CfnApiCache",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiCache) SetApiCachingBehavior(val *string) {
	_jsii_.Set(
		j,
		"apiCachingBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetAtRestEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"atRestEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetTransitEncryptionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetTtl(val *float64) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_CfnApiCache) SetType(val *string) {
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
// Experimental.
func CfnApiCache_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnApiCache",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApiCache_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnApiCache",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApiCache_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnApiCache",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiCache_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appsync.CfnApiCache",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApiCache) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApiCache) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApiCache) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnApiCache) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApiCache) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnApiCache) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApiCache) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnApiCache) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnApiCache) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApiCache) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnApiCache) OnPrepare() {
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
func (c *jsiiProxy_CfnApiCache) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnApiCache) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnApiCache) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnApiCache) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApiCache) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnApiCache) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnApiCache) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnApiCache) ToString() *string {
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
func (c *jsiiProxy_CfnApiCache) Validate() *[]*string {
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
func (c *jsiiProxy_CfnApiCache) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::AppSync::ApiCache`.
type CfnApiCacheProps struct {
	// `AWS::AppSync::ApiCache.ApiCachingBehavior`.
	ApiCachingBehavior *string `json:"apiCachingBehavior"`
	// `AWS::AppSync::ApiCache.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::AppSync::ApiCache.Ttl`.
	Ttl *float64 `json:"ttl"`
	// `AWS::AppSync::ApiCache.Type`.
	Type *string `json:"type"`
	// `AWS::AppSync::ApiCache.AtRestEncryptionEnabled`.
	AtRestEncryptionEnabled interface{} `json:"atRestEncryptionEnabled"`
	// `AWS::AppSync::ApiCache.TransitEncryptionEnabled`.
	TransitEncryptionEnabled interface{} `json:"transitEncryptionEnabled"`
}

// A CloudFormation `AWS::AppSync::ApiKey`.
type CfnApiKey interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	ApiKeyId() *string
	SetApiKeyId(val *string)
	AttrApiKey() *string
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	Expires() *float64
	SetExpires(val *float64)
	LogicalId() *string
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for CfnApiKey
type jsiiProxy_CfnApiKey struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnApiKey) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) ApiKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) AttrApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Expires() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApiKey) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::ApiKey`.
func NewCfnApiKey(scope awscdk.Construct, id *string, props *CfnApiKeyProps) CfnApiKey {
	_init_.Initialize()

	j := jsiiProxy_CfnApiKey{}

	_jsii_.Create(
		"monocdk.aws_appsync.CfnApiKey",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::ApiKey`.
func NewCfnApiKey_Override(c CfnApiKey, scope awscdk.Construct, id *string, props *CfnApiKeyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.CfnApiKey",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnApiKey) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetApiKeyId(val *string) {
	_jsii_.Set(
		j,
		"apiKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnApiKey) SetExpires(val *float64) {
	_jsii_.Set(
		j,
		"expires",
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
func CfnApiKey_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnApiKey",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnApiKey_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnApiKey",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnApiKey_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnApiKey",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApiKey_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appsync.CfnApiKey",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnApiKey) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnApiKey) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnApiKey) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnApiKey) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnApiKey) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnApiKey) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnApiKey) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnApiKey) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnApiKey) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnApiKey) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnApiKey) OnPrepare() {
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
func (c *jsiiProxy_CfnApiKey) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnApiKey) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnApiKey) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnApiKey) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnApiKey) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnApiKey) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnApiKey) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnApiKey) ToString() *string {
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
func (c *jsiiProxy_CfnApiKey) Validate() *[]*string {
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
func (c *jsiiProxy_CfnApiKey) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::AppSync::ApiKey`.
type CfnApiKeyProps struct {
	// `AWS::AppSync::ApiKey.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::AppSync::ApiKey.ApiKeyId`.
	ApiKeyId *string `json:"apiKeyId"`
	// `AWS::AppSync::ApiKey.Description`.
	Description *string `json:"description"`
	// `AWS::AppSync::ApiKey.Expires`.
	Expires *float64 `json:"expires"`
}

// A CloudFormation `AWS::AppSync::DataSource`.
type CfnDataSource interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	AttrDataSourceArn() *string
	AttrName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DynamoDbConfig() interface{}
	SetDynamoDbConfig(val interface{})
	ElasticsearchConfig() interface{}
	SetElasticsearchConfig(val interface{})
	HttpConfig() interface{}
	SetHttpConfig(val interface{})
	LambdaConfig() interface{}
	SetLambdaConfig(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RelationalDatabaseConfig() interface{}
	SetRelationalDatabaseConfig(val interface{})
	ServiceRoleArn() *string
	SetServiceRoleArn(val *string)
	Stack() awscdk.Stack
	Type() *string
	SetType(val *string)
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

// The jsii proxy struct for CfnDataSource
type jsiiProxy_CfnDataSource struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDataSource) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrDataSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDataSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) DynamoDbConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ElasticsearchConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) HttpConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LambdaConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) RelationalDatabaseConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relationalDatabaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDataSource) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::DataSource`.
func NewCfnDataSource(scope awscdk.Construct, id *string, props *CfnDataSourceProps) CfnDataSource {
	_init_.Initialize()

	j := jsiiProxy_CfnDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.CfnDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::DataSource`.
func NewCfnDataSource_Override(c CfnDataSource, scope awscdk.Construct, id *string, props *CfnDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.CfnDataSource",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDataSource) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetDynamoDbConfig(val interface{}) {
	_jsii_.Set(
		j,
		"dynamoDbConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetElasticsearchConfig(val interface{}) {
	_jsii_.Set(
		j,
		"elasticsearchConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetHttpConfig(val interface{}) {
	_jsii_.Set(
		j,
		"httpConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetLambdaConfig(val interface{}) {
	_jsii_.Set(
		j,
		"lambdaConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetRelationalDatabaseConfig(val interface{}) {
	_jsii_.Set(
		j,
		"relationalDatabaseConfig",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetServiceRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnDataSource) SetType(val *string) {
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
// Experimental.
func CfnDataSource_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnDataSource",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnDataSource_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnDataSource",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDataSource_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appsync.CfnDataSource",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnDataSource) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDataSource) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDataSource) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDataSource) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnDataSource) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnDataSource) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnDataSource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDataSource) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDataSource) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDataSource) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnDataSource) OnPrepare() {
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
func (c *jsiiProxy_CfnDataSource) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDataSource) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnDataSource) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnDataSource) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnDataSource) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDataSource) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDataSource) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnDataSource) ToString() *string {
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
func (c *jsiiProxy_CfnDataSource) Validate() *[]*string {
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
func (c *jsiiProxy_CfnDataSource) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnDataSource_AuthorizationConfigProperty struct {
	// `CfnDataSource.AuthorizationConfigProperty.AuthorizationType`.
	AuthorizationType *string `json:"authorizationType"`
	// `CfnDataSource.AuthorizationConfigProperty.AwsIamConfig`.
	AwsIamConfig interface{} `json:"awsIamConfig"`
}

type CfnDataSource_AwsIamConfigProperty struct {
	// `CfnDataSource.AwsIamConfigProperty.SigningRegion`.
	SigningRegion *string `json:"signingRegion"`
	// `CfnDataSource.AwsIamConfigProperty.SigningServiceName`.
	SigningServiceName *string `json:"signingServiceName"`
}

type CfnDataSource_DeltaSyncConfigProperty struct {
	// `CfnDataSource.DeltaSyncConfigProperty.BaseTableTTL`.
	BaseTableTtl *string `json:"baseTableTtl"`
	// `CfnDataSource.DeltaSyncConfigProperty.DeltaSyncTableName`.
	DeltaSyncTableName *string `json:"deltaSyncTableName"`
	// `CfnDataSource.DeltaSyncConfigProperty.DeltaSyncTableTTL`.
	DeltaSyncTableTtl *string `json:"deltaSyncTableTtl"`
}

type CfnDataSource_DynamoDBConfigProperty struct {
	// `CfnDataSource.DynamoDBConfigProperty.AwsRegion`.
	AwsRegion *string `json:"awsRegion"`
	// `CfnDataSource.DynamoDBConfigProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnDataSource.DynamoDBConfigProperty.DeltaSyncConfig`.
	DeltaSyncConfig interface{} `json:"deltaSyncConfig"`
	// `CfnDataSource.DynamoDBConfigProperty.UseCallerCredentials`.
	UseCallerCredentials interface{} `json:"useCallerCredentials"`
	// `CfnDataSource.DynamoDBConfigProperty.Versioned`.
	Versioned interface{} `json:"versioned"`
}

type CfnDataSource_ElasticsearchConfigProperty struct {
	// `CfnDataSource.ElasticsearchConfigProperty.AwsRegion`.
	AwsRegion *string `json:"awsRegion"`
	// `CfnDataSource.ElasticsearchConfigProperty.Endpoint`.
	Endpoint *string `json:"endpoint"`
}

type CfnDataSource_HttpConfigProperty struct {
	// `CfnDataSource.HttpConfigProperty.Endpoint`.
	Endpoint *string `json:"endpoint"`
	// `CfnDataSource.HttpConfigProperty.AuthorizationConfig`.
	AuthorizationConfig interface{} `json:"authorizationConfig"`
}

type CfnDataSource_LambdaConfigProperty struct {
	// `CfnDataSource.LambdaConfigProperty.LambdaFunctionArn`.
	LambdaFunctionArn *string `json:"lambdaFunctionArn"`
}

type CfnDataSource_RdsHttpEndpointConfigProperty struct {
	// `CfnDataSource.RdsHttpEndpointConfigProperty.AwsRegion`.
	AwsRegion *string `json:"awsRegion"`
	// `CfnDataSource.RdsHttpEndpointConfigProperty.AwsSecretStoreArn`.
	AwsSecretStoreArn *string `json:"awsSecretStoreArn"`
	// `CfnDataSource.RdsHttpEndpointConfigProperty.DbClusterIdentifier`.
	DbClusterIdentifier *string `json:"dbClusterIdentifier"`
	// `CfnDataSource.RdsHttpEndpointConfigProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnDataSource.RdsHttpEndpointConfigProperty.Schema`.
	Schema *string `json:"schema"`
}

type CfnDataSource_RelationalDatabaseConfigProperty struct {
	// `CfnDataSource.RelationalDatabaseConfigProperty.RelationalDatabaseSourceType`.
	RelationalDatabaseSourceType *string `json:"relationalDatabaseSourceType"`
	// `CfnDataSource.RelationalDatabaseConfigProperty.RdsHttpEndpointConfig`.
	RdsHttpEndpointConfig interface{} `json:"rdsHttpEndpointConfig"`
}

// Properties for defining a `AWS::AppSync::DataSource`.
type CfnDataSourceProps struct {
	// `AWS::AppSync::DataSource.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::AppSync::DataSource.Name`.
	Name *string `json:"name"`
	// `AWS::AppSync::DataSource.Type`.
	Type *string `json:"type"`
	// `AWS::AppSync::DataSource.Description`.
	Description *string `json:"description"`
	// `AWS::AppSync::DataSource.DynamoDBConfig`.
	DynamoDbConfig interface{} `json:"dynamoDbConfig"`
	// `AWS::AppSync::DataSource.ElasticsearchConfig`.
	ElasticsearchConfig interface{} `json:"elasticsearchConfig"`
	// `AWS::AppSync::DataSource.HttpConfig`.
	HttpConfig interface{} `json:"httpConfig"`
	// `AWS::AppSync::DataSource.LambdaConfig`.
	LambdaConfig interface{} `json:"lambdaConfig"`
	// `AWS::AppSync::DataSource.RelationalDatabaseConfig`.
	RelationalDatabaseConfig interface{} `json:"relationalDatabaseConfig"`
	// `AWS::AppSync::DataSource.ServiceRoleArn`.
	ServiceRoleArn *string `json:"serviceRoleArn"`
}

// A CloudFormation `AWS::AppSync::FunctionConfiguration`.
type CfnFunctionConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	AttrDataSourceName() *string
	AttrFunctionArn() *string
	AttrFunctionId() *string
	AttrName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataSourceName() *string
	SetDataSourceName(val *string)
	Description() *string
	SetDescription(val *string)
	FunctionVersion() *string
	SetFunctionVersion(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RequestMappingTemplate() *string
	SetRequestMappingTemplate(val *string)
	RequestMappingTemplateS3Location() *string
	SetRequestMappingTemplateS3Location(val *string)
	ResponseMappingTemplate() *string
	SetResponseMappingTemplate(val *string)
	ResponseMappingTemplateS3Location() *string
	SetResponseMappingTemplateS3Location(val *string)
	Stack() awscdk.Stack
	SyncConfig() interface{}
	SetSyncConfig(val interface{})
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

// The jsii proxy struct for CfnFunctionConfiguration
type jsiiProxy_CfnFunctionConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFunctionConfiguration) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrDataSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDataSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrFunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFunctionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) DataSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) FunctionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) RequestMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) RequestMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) ResponseMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) ResponseMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) SyncConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunctionConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::FunctionConfiguration`.
func NewCfnFunctionConfiguration(scope awscdk.Construct, id *string, props *CfnFunctionConfigurationProps) CfnFunctionConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnFunctionConfiguration{}

	_jsii_.Create(
		"monocdk.aws_appsync.CfnFunctionConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::FunctionConfiguration`.
func NewCfnFunctionConfiguration_Override(c CfnFunctionConfiguration, scope awscdk.Construct, id *string, props *CfnFunctionConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.CfnFunctionConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetDataSourceName(val *string) {
	_jsii_.Set(
		j,
		"dataSourceName",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetFunctionVersion(val *string) {
	_jsii_.Set(
		j,
		"functionVersion",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetRequestMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetRequestMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetResponseMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetResponseMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnFunctionConfiguration) SetSyncConfig(val interface{}) {
	_jsii_.Set(
		j,
		"syncConfig",
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
func CfnFunctionConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnFunctionConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnFunctionConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnFunctionConfiguration",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnFunctionConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnFunctionConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFunctionConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appsync.CfnFunctionConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnFunctionConfiguration) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnFunctionConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnFunctionConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnFunctionConfiguration) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnFunctionConfiguration) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) OnPrepare() {
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
func (c *jsiiProxy_CfnFunctionConfiguration) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnFunctionConfiguration) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnFunctionConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnFunctionConfiguration) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnFunctionConfiguration) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnFunctionConfiguration) ToString() *string {
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
func (c *jsiiProxy_CfnFunctionConfiguration) Validate() *[]*string {
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
func (c *jsiiProxy_CfnFunctionConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnFunctionConfiguration_LambdaConflictHandlerConfigProperty struct {
	// `CfnFunctionConfiguration.LambdaConflictHandlerConfigProperty.LambdaConflictHandlerArn`.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn"`
}

type CfnFunctionConfiguration_SyncConfigProperty struct {
	// `CfnFunctionConfiguration.SyncConfigProperty.ConflictDetection`.
	ConflictDetection *string `json:"conflictDetection"`
	// `CfnFunctionConfiguration.SyncConfigProperty.ConflictHandler`.
	ConflictHandler *string `json:"conflictHandler"`
	// `CfnFunctionConfiguration.SyncConfigProperty.LambdaConflictHandlerConfig`.
	LambdaConflictHandlerConfig interface{} `json:"lambdaConflictHandlerConfig"`
}

// Properties for defining a `AWS::AppSync::FunctionConfiguration`.
type CfnFunctionConfigurationProps struct {
	// `AWS::AppSync::FunctionConfiguration.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::AppSync::FunctionConfiguration.DataSourceName`.
	DataSourceName *string `json:"dataSourceName"`
	// `AWS::AppSync::FunctionConfiguration.FunctionVersion`.
	FunctionVersion *string `json:"functionVersion"`
	// `AWS::AppSync::FunctionConfiguration.Name`.
	Name *string `json:"name"`
	// `AWS::AppSync::FunctionConfiguration.Description`.
	Description *string `json:"description"`
	// `AWS::AppSync::FunctionConfiguration.RequestMappingTemplate`.
	RequestMappingTemplate *string `json:"requestMappingTemplate"`
	// `AWS::AppSync::FunctionConfiguration.RequestMappingTemplateS3Location`.
	RequestMappingTemplateS3Location *string `json:"requestMappingTemplateS3Location"`
	// `AWS::AppSync::FunctionConfiguration.ResponseMappingTemplate`.
	ResponseMappingTemplate *string `json:"responseMappingTemplate"`
	// `AWS::AppSync::FunctionConfiguration.ResponseMappingTemplateS3Location`.
	ResponseMappingTemplateS3Location *string `json:"responseMappingTemplateS3Location"`
	// `AWS::AppSync::FunctionConfiguration.SyncConfig`.
	SyncConfig interface{} `json:"syncConfig"`
}

// A CloudFormation `AWS::AppSync::GraphQLApi`.
type CfnGraphQLApi interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AdditionalAuthenticationProviders() interface{}
	SetAdditionalAuthenticationProviders(val interface{})
	AttrApiId() *string
	AttrArn() *string
	AttrGraphQlUrl() *string
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogConfig() interface{}
	SetLogConfig(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	OpenIdConnectConfig() interface{}
	SetOpenIdConnectConfig(val interface{})
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UserPoolConfig() interface{}
	SetUserPoolConfig(val interface{})
	XrayEnabled() interface{}
	SetXrayEnabled(val interface{})
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

// The jsii proxy struct for CfnGraphQLApi
type jsiiProxy_CfnGraphQLApi struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGraphQLApi) AdditionalAuthenticationProviders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AttrGraphQlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrGraphQlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LogConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) OpenIdConnectConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openIdConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) UserPoolConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLApi) XrayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::GraphQLApi`.
func NewCfnGraphQLApi(scope awscdk.Construct, id *string, props *CfnGraphQLApiProps) CfnGraphQLApi {
	_init_.Initialize()

	j := jsiiProxy_CfnGraphQLApi{}

	_jsii_.Create(
		"monocdk.aws_appsync.CfnGraphQLApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::GraphQLApi`.
func NewCfnGraphQLApi_Override(c CfnGraphQLApi, scope awscdk.Construct, id *string, props *CfnGraphQLApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.CfnGraphQLApi",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetAdditionalAuthenticationProviders(val interface{}) {
	_jsii_.Set(
		j,
		"additionalAuthenticationProviders",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetAuthenticationType(val *string) {
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetLogConfig(val interface{}) {
	_jsii_.Set(
		j,
		"logConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetOpenIdConnectConfig(val interface{}) {
	_jsii_.Set(
		j,
		"openIdConnectConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetUserPoolConfig(val interface{}) {
	_jsii_.Set(
		j,
		"userPoolConfig",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLApi) SetXrayEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"xrayEnabled",
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
func CfnGraphQLApi_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnGraphQLApi",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGraphQLApi_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnGraphQLApi",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGraphQLApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnGraphQLApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGraphQLApi_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appsync.CfnGraphQLApi",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnGraphQLApi) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnGraphQLApi) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnGraphQLApi) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnGraphQLApi) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnGraphQLApi) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnGraphQLApi) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnGraphQLApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnGraphQLApi) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnGraphQLApi) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnGraphQLApi) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnGraphQLApi) OnPrepare() {
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
func (c *jsiiProxy_CfnGraphQLApi) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnGraphQLApi) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnGraphQLApi) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnGraphQLApi) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGraphQLApi) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnGraphQLApi) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnGraphQLApi) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnGraphQLApi) ToString() *string {
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
func (c *jsiiProxy_CfnGraphQLApi) Validate() *[]*string {
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
func (c *jsiiProxy_CfnGraphQLApi) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnGraphQLApi_AdditionalAuthenticationProviderProperty struct {
	// `CfnGraphQLApi.AdditionalAuthenticationProviderProperty.AuthenticationType`.
	AuthenticationType *string `json:"authenticationType"`
	// `CfnGraphQLApi.AdditionalAuthenticationProviderProperty.OpenIDConnectConfig`.
	OpenIdConnectConfig interface{} `json:"openIdConnectConfig"`
	// `CfnGraphQLApi.AdditionalAuthenticationProviderProperty.UserPoolConfig`.
	UserPoolConfig interface{} `json:"userPoolConfig"`
}

type CfnGraphQLApi_CognitoUserPoolConfigProperty struct {
	// `CfnGraphQLApi.CognitoUserPoolConfigProperty.AppIdClientRegex`.
	AppIdClientRegex *string `json:"appIdClientRegex"`
	// `CfnGraphQLApi.CognitoUserPoolConfigProperty.AwsRegion`.
	AwsRegion *string `json:"awsRegion"`
	// `CfnGraphQLApi.CognitoUserPoolConfigProperty.UserPoolId`.
	UserPoolId *string `json:"userPoolId"`
}

type CfnGraphQLApi_LogConfigProperty struct {
	// `CfnGraphQLApi.LogConfigProperty.CloudWatchLogsRoleArn`.
	CloudWatchLogsRoleArn *string `json:"cloudWatchLogsRoleArn"`
	// `CfnGraphQLApi.LogConfigProperty.ExcludeVerboseContent`.
	ExcludeVerboseContent interface{} `json:"excludeVerboseContent"`
	// `CfnGraphQLApi.LogConfigProperty.FieldLogLevel`.
	FieldLogLevel *string `json:"fieldLogLevel"`
}

type CfnGraphQLApi_OpenIDConnectConfigProperty struct {
	// `CfnGraphQLApi.OpenIDConnectConfigProperty.AuthTTL`.
	AuthTtl *float64 `json:"authTtl"`
	// `CfnGraphQLApi.OpenIDConnectConfigProperty.ClientId`.
	ClientId *string `json:"clientId"`
	// `CfnGraphQLApi.OpenIDConnectConfigProperty.IatTTL`.
	IatTtl *float64 `json:"iatTtl"`
	// `CfnGraphQLApi.OpenIDConnectConfigProperty.Issuer`.
	Issuer *string `json:"issuer"`
}

type CfnGraphQLApi_UserPoolConfigProperty struct {
	// `CfnGraphQLApi.UserPoolConfigProperty.AppIdClientRegex`.
	AppIdClientRegex *string `json:"appIdClientRegex"`
	// `CfnGraphQLApi.UserPoolConfigProperty.AwsRegion`.
	AwsRegion *string `json:"awsRegion"`
	// `CfnGraphQLApi.UserPoolConfigProperty.DefaultAction`.
	DefaultAction *string `json:"defaultAction"`
	// `CfnGraphQLApi.UserPoolConfigProperty.UserPoolId`.
	UserPoolId *string `json:"userPoolId"`
}

// Properties for defining a `AWS::AppSync::GraphQLApi`.
type CfnGraphQLApiProps struct {
	// `AWS::AppSync::GraphQLApi.AuthenticationType`.
	AuthenticationType *string `json:"authenticationType"`
	// `AWS::AppSync::GraphQLApi.Name`.
	Name *string `json:"name"`
	// `AWS::AppSync::GraphQLApi.AdditionalAuthenticationProviders`.
	AdditionalAuthenticationProviders interface{} `json:"additionalAuthenticationProviders"`
	// `AWS::AppSync::GraphQLApi.LogConfig`.
	LogConfig interface{} `json:"logConfig"`
	// `AWS::AppSync::GraphQLApi.OpenIDConnectConfig`.
	OpenIdConnectConfig interface{} `json:"openIdConnectConfig"`
	// `AWS::AppSync::GraphQLApi.Tags`.
	Tags interface{} `json:"tags"`
	// `AWS::AppSync::GraphQLApi.UserPoolConfig`.
	UserPoolConfig interface{} `json:"userPoolConfig"`
	// `AWS::AppSync::GraphQLApi.XrayEnabled`.
	XrayEnabled interface{} `json:"xrayEnabled"`
}

// A CloudFormation `AWS::AppSync::GraphQLSchema`.
type CfnGraphQLSchema interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Definition() *string
	SetDefinition(val *string)
	DefinitionS3Location() *string
	SetDefinitionS3Location(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
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

// The jsii proxy struct for CfnGraphQLSchema
type jsiiProxy_CfnGraphQLSchema struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGraphQLSchema) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) DefinitionS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphQLSchema) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::GraphQLSchema`.
func NewCfnGraphQLSchema(scope awscdk.Construct, id *string, props *CfnGraphQLSchemaProps) CfnGraphQLSchema {
	_init_.Initialize()

	j := jsiiProxy_CfnGraphQLSchema{}

	_jsii_.Create(
		"monocdk.aws_appsync.CfnGraphQLSchema",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::GraphQLSchema`.
func NewCfnGraphQLSchema_Override(c CfnGraphQLSchema, scope awscdk.Construct, id *string, props *CfnGraphQLSchemaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.CfnGraphQLSchema",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGraphQLSchema) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLSchema) SetDefinition(val *string) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

func (j *jsiiProxy_CfnGraphQLSchema) SetDefinitionS3Location(val *string) {
	_jsii_.Set(
		j,
		"definitionS3Location",
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
func CfnGraphQLSchema_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnGraphQLSchema",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGraphQLSchema_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnGraphQLSchema",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGraphQLSchema_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnGraphQLSchema",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGraphQLSchema_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appsync.CfnGraphQLSchema",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnGraphQLSchema) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnGraphQLSchema) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnGraphQLSchema) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnGraphQLSchema) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnGraphQLSchema) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnGraphQLSchema) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnGraphQLSchema) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnGraphQLSchema) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnGraphQLSchema) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnGraphQLSchema) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnGraphQLSchema) OnPrepare() {
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
func (c *jsiiProxy_CfnGraphQLSchema) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnGraphQLSchema) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnGraphQLSchema) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnGraphQLSchema) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGraphQLSchema) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnGraphQLSchema) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnGraphQLSchema) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnGraphQLSchema) ToString() *string {
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
func (c *jsiiProxy_CfnGraphQLSchema) Validate() *[]*string {
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
func (c *jsiiProxy_CfnGraphQLSchema) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::AppSync::GraphQLSchema`.
type CfnGraphQLSchemaProps struct {
	// `AWS::AppSync::GraphQLSchema.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::AppSync::GraphQLSchema.Definition`.
	Definition *string `json:"definition"`
	// `AWS::AppSync::GraphQLSchema.DefinitionS3Location`.
	DefinitionS3Location *string `json:"definitionS3Location"`
}

// A CloudFormation `AWS::AppSync::Resolver`.
type CfnResolver interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApiId() *string
	SetApiId(val *string)
	AttrFieldName() *string
	AttrResolverArn() *string
	AttrTypeName() *string
	CachingConfig() interface{}
	SetCachingConfig(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DataSourceName() *string
	SetDataSourceName(val *string)
	FieldName() *string
	SetFieldName(val *string)
	Kind() *string
	SetKind(val *string)
	LogicalId() *string
	Node() awscdk.ConstructNode
	PipelineConfig() interface{}
	SetPipelineConfig(val interface{})
	Ref() *string
	RequestMappingTemplate() *string
	SetRequestMappingTemplate(val *string)
	RequestMappingTemplateS3Location() *string
	SetRequestMappingTemplateS3Location(val *string)
	ResponseMappingTemplate() *string
	SetResponseMappingTemplate(val *string)
	ResponseMappingTemplateS3Location() *string
	SetResponseMappingTemplateS3Location(val *string)
	Stack() awscdk.Stack
	SyncConfig() interface{}
	SetSyncConfig(val interface{})
	TypeName() *string
	SetTypeName(val *string)
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

// The jsii proxy struct for CfnResolver
type jsiiProxy_CfnResolver struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnResolver) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) AttrFieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) AttrResolverArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrResolverArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) AttrTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CachingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) DataSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) PipelineConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pipelineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) RequestMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) RequestMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) ResponseMappingTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) ResponseMappingTemplateS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseMappingTemplateS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) SyncConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResolver) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppSync::Resolver`.
func NewCfnResolver(scope awscdk.Construct, id *string, props *CfnResolverProps) CfnResolver {
	_init_.Initialize()

	j := jsiiProxy_CfnResolver{}

	_jsii_.Create(
		"monocdk.aws_appsync.CfnResolver",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppSync::Resolver`.
func NewCfnResolver_Override(c CfnResolver, scope awscdk.Construct, id *string, props *CfnResolverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.CfnResolver",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnResolver) SetApiId(val *string) {
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetCachingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"cachingConfig",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetDataSourceName(val *string) {
	_jsii_.Set(
		j,
		"dataSourceName",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetFieldName(val *string) {
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetKind(val *string) {
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetPipelineConfig(val interface{}) {
	_jsii_.Set(
		j,
		"pipelineConfig",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetRequestMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetRequestMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"requestMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetResponseMappingTemplate(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetResponseMappingTemplateS3Location(val *string) {
	_jsii_.Set(
		j,
		"responseMappingTemplateS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetSyncConfig(val interface{}) {
	_jsii_.Set(
		j,
		"syncConfig",
		val,
	)
}

func (j *jsiiProxy_CfnResolver) SetTypeName(val *string) {
	_jsii_.Set(
		j,
		"typeName",
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
func CfnResolver_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnResolver",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnResolver_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnResolver",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnResolver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.CfnResolver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResolver_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_appsync.CfnResolver",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnResolver) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnResolver) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnResolver) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnResolver) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnResolver) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnResolver) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnResolver) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnResolver) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnResolver) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnResolver) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnResolver) OnPrepare() {
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
func (c *jsiiProxy_CfnResolver) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnResolver) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnResolver) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnResolver) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnResolver) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnResolver) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnResolver) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnResolver) ToString() *string {
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
func (c *jsiiProxy_CfnResolver) Validate() *[]*string {
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
func (c *jsiiProxy_CfnResolver) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnResolver_CachingConfigProperty struct {
	// `CfnResolver.CachingConfigProperty.CachingKeys`.
	CachingKeys *[]*string `json:"cachingKeys"`
	// `CfnResolver.CachingConfigProperty.Ttl`.
	Ttl *float64 `json:"ttl"`
}

type CfnResolver_LambdaConflictHandlerConfigProperty struct {
	// `CfnResolver.LambdaConflictHandlerConfigProperty.LambdaConflictHandlerArn`.
	LambdaConflictHandlerArn *string `json:"lambdaConflictHandlerArn"`
}

type CfnResolver_PipelineConfigProperty struct {
	// `CfnResolver.PipelineConfigProperty.Functions`.
	Functions *[]*string `json:"functions"`
}

type CfnResolver_SyncConfigProperty struct {
	// `CfnResolver.SyncConfigProperty.ConflictDetection`.
	ConflictDetection *string `json:"conflictDetection"`
	// `CfnResolver.SyncConfigProperty.ConflictHandler`.
	ConflictHandler *string `json:"conflictHandler"`
	// `CfnResolver.SyncConfigProperty.LambdaConflictHandlerConfig`.
	LambdaConflictHandlerConfig interface{} `json:"lambdaConflictHandlerConfig"`
}

// Properties for defining a `AWS::AppSync::Resolver`.
type CfnResolverProps struct {
	// `AWS::AppSync::Resolver.ApiId`.
	ApiId *string `json:"apiId"`
	// `AWS::AppSync::Resolver.FieldName`.
	FieldName *string `json:"fieldName"`
	// `AWS::AppSync::Resolver.TypeName`.
	TypeName *string `json:"typeName"`
	// `AWS::AppSync::Resolver.CachingConfig`.
	CachingConfig interface{} `json:"cachingConfig"`
	// `AWS::AppSync::Resolver.DataSourceName`.
	DataSourceName *string `json:"dataSourceName"`
	// `AWS::AppSync::Resolver.Kind`.
	Kind *string `json:"kind"`
	// `AWS::AppSync::Resolver.PipelineConfig`.
	PipelineConfig interface{} `json:"pipelineConfig"`
	// `AWS::AppSync::Resolver.RequestMappingTemplate`.
	RequestMappingTemplate *string `json:"requestMappingTemplate"`
	// `AWS::AppSync::Resolver.RequestMappingTemplateS3Location`.
	RequestMappingTemplateS3Location *string `json:"requestMappingTemplateS3Location"`
	// `AWS::AppSync::Resolver.ResponseMappingTemplate`.
	ResponseMappingTemplate *string `json:"responseMappingTemplate"`
	// `AWS::AppSync::Resolver.ResponseMappingTemplateS3Location`.
	ResponseMappingTemplateS3Location *string `json:"responseMappingTemplateS3Location"`
	// `AWS::AppSync::Resolver.SyncConfig`.
	SyncConfig interface{} `json:"syncConfig"`
}

// Optional configuration for data sources.
// Experimental.
type DataSourceOptions struct {
	// The description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source, overrides the id given by cdk.
	// Experimental.
	Name *string `json:"name"`
}

// Directives for types.
//
// i.e. @aws_iam or @aws_subscribe
// Experimental.
type Directive interface {
	Mode() AuthorizationType
	Modes() *[]AuthorizationType
	SetModes(val *[]AuthorizationType)
	MutationFields() *[]*string
	ToString() *string
}

// The jsii proxy struct for Directive
type jsiiProxy_Directive struct {
	_ byte // padding
}

func (j *jsiiProxy_Directive) Mode() AuthorizationType {
	var returns AuthorizationType
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Directive) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Directive) MutationFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mutationFields",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_Directive) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

// Add the @aws_api_key directive.
// Experimental.
func Directive_ApiKey() Directive {
	_init_.Initialize()

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"apiKey",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add the @aws_auth or @aws_cognito_user_pools directive.
// Experimental.
func Directive_Cognito(groups ...*string) Directive {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range groups {
		args = append(args, a)
	}

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"cognito",
		args,
		&returns,
	)

	return returns
}

// Add a custom directive.
// Experimental.
func Directive_Custom(statement *string) Directive {
	_init_.Initialize()

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"custom",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

// Add the @aws_iam directive.
// Experimental.
func Directive_Iam() Directive {
	_init_.Initialize()

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"iam",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add the @aws_oidc directive.
// Experimental.
func Directive_Oidc() Directive {
	_init_.Initialize()

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"oidc",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Add the @aws_subscribe directive.
//
// Only use for top level Subscription type.
// Experimental.
func Directive_Subscribe(mutations ...*string) Directive {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range mutations {
		args = append(args, a)
	}

	var returns Directive

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Directive",
		"subscribe",
		args,
		&returns,
	)

	return returns
}

// Generate the directive statement.
// Experimental.
func (d *jsiiProxy_Directive) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An AppSync datasource backed by a DynamoDB table.
// Experimental.
type DynamoDbDataSource interface {
	BackedDataSource
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	Ds() CfnDataSource
	GrantPrincipal() awsiam.IPrincipal
	Name() *string
	Node() awscdk.ConstructNode
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	CreateResolver(props *BaseResolverProps) Resolver
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for DynamoDbDataSource
type jsiiProxy_DynamoDbDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_DynamoDbDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDbDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDbDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDbDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDbDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDbDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoDbDataSource(scope constructs.Construct, id *string, props *DynamoDbDataSourceProps) DynamoDbDataSource {
	_init_.Initialize()

	j := jsiiProxy_DynamoDbDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.DynamoDbDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoDbDataSource_Override(d DynamoDbDataSource, scope constructs.Construct, id *string, props *DynamoDbDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.DynamoDbDataSource",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DynamoDbDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_DynamoDbDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func DynamoDbDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.DynamoDbDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// creates a new appsync function for this datasource and API using the given properties.
// Experimental.
func (d *jsiiProxy_DynamoDbDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		d,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (d *jsiiProxy_DynamoDbDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		d,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
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
func (d *jsiiProxy_DynamoDbDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoDbDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
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
func (d *jsiiProxy_DynamoDbDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (d *jsiiProxy_DynamoDbDataSource) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoDbDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DynamoDbDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DynamoDbDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an AppSync DynamoDB datasource.
// Experimental.
type DynamoDbDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `json:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Experimental.
	ServiceRole awsiam.IRole `json:"serviceRole"`
	// The DynamoDB table backing this data source.
	// Experimental.
	Table awsdynamodb.ITable `json:"table"`
	// Specify whether this DS is read only or has read and write permissions to the DynamoDB table.
	// Experimental.
	ReadOnlyAccess *bool `json:"readOnlyAccess"`
	// use credentials of caller to access DynamoDB.
	// Experimental.
	UseCallerCredentials *bool `json:"useCallerCredentials"`
}

// Enum Types are abstract types that includes a set of fields that represent the strings this type can create.
// Experimental.
type EnumType interface {
	IIntermediateType
	Definition() *map[string]IField
	Modes() *[]AuthorizationType
	SetModes(val *[]AuthorizationType)
	Name() *string
	AddField(options *AddFieldOptions)
	Attribute(options *BaseTypeOptions) GraphqlType
	ToString() *string
}

// The jsii proxy struct for EnumType
type jsiiProxy_EnumType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_EnumType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnumType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnumType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewEnumType(name *string, options *EnumTypeOptions) EnumType {
	_init_.Initialize()

	j := jsiiProxy_EnumType{}

	_jsii_.Create(
		"monocdk.aws_appsync.EnumType",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewEnumType_Override(e EnumType, name *string, options *EnumTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.EnumType",
		[]interface{}{name, options},
		e,
	)
}

func (j *jsiiProxy_EnumType) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

// Add a field to this Enum Type.
//
// To add a field to this Enum Type, you must only configure
// addField with the fieldName options.
// Experimental.
func (e *jsiiProxy_EnumType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		e,
		"addField",
		[]interface{}{options},
	)
}

// Create an GraphQL Type representing this Enum Type.
// Experimental.
func (e *jsiiProxy_EnumType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		e,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the string of this enum type.
// Experimental.
func (e *jsiiProxy_EnumType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for configuring an Enum Type.
// Experimental.
type EnumTypeOptions struct {
	// the attributes of this type.
	// Experimental.
	Definition *[]*string `json:"definition"`
}

// props used by implementations of BaseDataSource to provide configuration.
//
// Should not be used directly.
// Experimental.
type ExtendedDataSourceProps struct {
	// the type of the AppSync datasource.
	// Experimental.
	Type *string `json:"type"`
	// configuration for DynamoDB Datasource.
	// Experimental.
	DynamoDbConfig interface{} `json:"dynamoDbConfig"`
	// configuration for Elasticsearch Datasource.
	// Experimental.
	ElasticsearchConfig interface{} `json:"elasticsearchConfig"`
	// configuration for HTTP Datasource.
	// Experimental.
	HttpConfig interface{} `json:"httpConfig"`
	// configuration for Lambda Datasource.
	// Experimental.
	LambdaConfig interface{} `json:"lambdaConfig"`
	// configuration for RDS Datasource.
	// Experimental.
	RelationalDatabaseConfig interface{} `json:"relationalDatabaseConfig"`
}

// Additional property for an AppSync resolver for data source reference.
// Experimental.
type ExtendedResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	// Experimental.
	FieldName *string `json:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	// Experimental.
	TypeName *string `json:"typeName"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `json:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `json:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `json:"responseMappingTemplate"`
	// The data source this resolver is using.
	// Experimental.
	DataSource BaseDataSource `json:"dataSource"`
}

// Fields build upon Graphql Types and provide typing and arguments.
// Experimental.
type Field interface {
	GraphqlType
	IField
	FieldOptions() *ResolvableFieldOptions
	IntermediateType() IIntermediateType
	IsList() *bool
	IsRequired() *bool
	IsRequiredList() *bool
	Type() Type
	ArgsToString() *string
	DirectivesToString(modes *[]AuthorizationType) *string
	ToString() *string
}

// The jsii proxy struct for Field
type jsiiProxy_Field struct {
	jsiiProxy_GraphqlType
	jsiiProxy_IField
}

func (j *jsiiProxy_Field) FieldOptions() *ResolvableFieldOptions {
	var returns *ResolvableFieldOptions
	_jsii_.Get(
		j,
		"fieldOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Field) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewField(options *FieldOptions) Field {
	_init_.Initialize()

	j := jsiiProxy_Field{}

	_jsii_.Create(
		"monocdk.aws_appsync.Field",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewField_Override(f Field, options *FieldOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.Field",
		[]interface{}{options},
		f,
	)
}

// `AWSDate` scalar type represents a valid extended `ISO 8601 Date` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DD`. It accepts time zone offsets.
// Experimental.
func Field_AwsDate(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsDate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSDateTime` scalar type represents a valid extended `ISO 8601 DateTime` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DDThh:mm:ss.sssZ`. It accepts time zone offsets.
// Experimental.
func Field_AwsDateTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsDateTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSEmail` scalar type represents an email address string (i.e.`username@example.com`).
// Experimental.
func Field_AwsEmail(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsEmail",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSIPAddress` scalar type respresents a valid `IPv4` of `IPv6` address string.
// Experimental.
func Field_AwsIpAddress(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsIpAddress",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSJson` scalar type represents a JSON string.
// Experimental.
func Field_AwsJson(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSPhone` scalar type represents a valid phone number. Phone numbers maybe be whitespace delimited or hyphenated.
//
// The number can specify a country code at the beginning, but is not required for US phone numbers.
// Experimental.
func Field_AwsPhone(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsPhone",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTime` scalar type represents a valid extended `ISO 8601 Time` string.
//
// In other words, accepts date strings in the form of `hh:mm:ss.sss`. It accepts time zone offsets.
// Experimental.
func Field_AwsTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTimestamp` scalar type represents the number of seconds since `1970-01-01T00:00Z`.
//
// Timestamps are serialized and deserialized as numbers.
// Experimental.
func Field_AwsTimestamp(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsTimestamp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSURL` scalar type represetns a valid URL string.
//
// URLs wihtout schemes or contain double slashes are considered invalid.
// Experimental.
func Field_AwsUrl(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"awsUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Boolean` scalar type is a boolean value: true or false.
// Experimental.
func Field_Boolean(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"boolean",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Float` scalar type is a signed double-precision fractional value.
// Experimental.
func Field_Float(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"float",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `ID` scalar type is a unique identifier. `ID` type is serialized similar to `String`.
//
// Often used as a key for a cache and not intended to be human-readable.
// Experimental.
func Field_Id(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"id",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Int` scalar type is a signed non-fractional numerical value.
// Experimental.
func Field_Int(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"int",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// an intermediate type to be added as an attribute (i.e. an interface or an object type).
// Experimental.
func Field_Intermediate(options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"intermediate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `String` scalar type is a free-form human-readable text.
// Experimental.
func Field_String(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Field",
		"string",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the args string of this resolvable field.
// Experimental.
func (f *jsiiProxy_Field) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the directives for this field.
// Experimental.
func (f *jsiiProxy_Field) DirectivesToString(modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"directivesToString",
		[]interface{}{modes},
		&returns,
	)

	return returns
}

// Generate the string for this attribute.
// Experimental.
func (f *jsiiProxy_Field) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// log-level for fields in AppSync.
// Experimental.
type FieldLogLevel string

const (
	FieldLogLevel_NONE FieldLogLevel = "NONE"
	FieldLogLevel_ERROR FieldLogLevel = "ERROR"
	FieldLogLevel_ALL FieldLogLevel = "ALL"
)

// Properties for configuring a field.
// Experimental.
type FieldOptions struct {
	// The return type for this field.
	// Experimental.
	ReturnType GraphqlType `json:"returnType"`
	// The arguments for this field.
	//
	// i.e. type Example (first: String second: String) {}
	// - where 'first' and 'second' are key values for args
	// and 'String' is the GraphqlType
	// Experimental.
	Args *map[string]GraphqlType `json:"args"`
	// the directives for this field.
	// Experimental.
	Directives *[]Directive `json:"directives"`
}

// An AppSync GraphQL API.
// Experimental.
type GraphqlApi interface {
	GraphqlApiBase
	ApiId() *string
	ApiKey() *string
	Arn() *string
	Env() *awscdk.ResourceEnvironment
	GraphqlUrl() *string
	Modes() *[]AuthorizationType
	Name() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Schema() Schema
	Stack() awscdk.Stack
	AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource
	AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource
	AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource
	AddMutation(fieldName *string, field ResolvableField) ObjectType
	AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource
	AddQuery(fieldName *string, field ResolvableField) ObjectType
	AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource
	AddSchemaDependency(construct awscdk.CfnResource) *bool
	AddSubscription(fieldName *string, field ResolvableField) ObjectType
	AddToSchema(addition *string, delimiter *string)
	AddType(type_ IIntermediateType) IIntermediateType
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	CreateResolver(props *ExtendedResolverProps) Resolver
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, resources IamResource, actions ...*string) awsiam.Grant
	GrantMutation(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant
	GrantQuery(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant
	GrantSubscription(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for GraphqlApi
type jsiiProxy_GraphqlApi struct {
	jsiiProxy_GraphqlApiBase
}

func (j *jsiiProxy_GraphqlApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) GraphqlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphqlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Schema() Schema {
	var returns Schema
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewGraphqlApi(scope constructs.Construct, id *string, props *GraphqlApiProps) GraphqlApi {
	_init_.Initialize()

	j := jsiiProxy_GraphqlApi{}

	_jsii_.Create(
		"monocdk.aws_appsync.GraphqlApi",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGraphqlApi_Override(g GraphqlApi, scope constructs.Construct, id *string, props *GraphqlApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.GraphqlApi",
		[]interface{}{scope, id, props},
		g,
	)
}

// Import a GraphQL API through this function.
// Experimental.
func GraphqlApi_FromGraphqlApiAttributes(scope constructs.Construct, id *string, attrs *GraphqlApiAttributes) IGraphqlApi {
	_init_.Initialize()

	var returns IGraphqlApi

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlApi",
		"fromGraphqlApiAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func GraphqlApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GraphqlApi_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlApi",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// add a new DynamoDB data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource {
	var returns DynamoDbDataSource

	_jsii_.Invoke(
		g,
		"addDynamoDbDataSource",
		[]interface{}{id, table, options},
		&returns,
	)

	return returns
}

// add a new http data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource {
	var returns HttpDataSource

	_jsii_.Invoke(
		g,
		"addHttpDataSource",
		[]interface{}{id, endpoint, options},
		&returns,
	)

	return returns
}

// add a new Lambda data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource {
	var returns LambdaDataSource

	_jsii_.Invoke(
		g,
		"addLambdaDataSource",
		[]interface{}{id, lambdaFunction, options},
		&returns,
	)

	return returns
}

// Add a mutation field to the schema's Mutation. CDK will create an Object Type called 'Mutation'. For example,.
//
// type Mutation {
//    fieldName: Field.returnType
// }
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddMutation(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		g,
		"addMutation",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

// add a new dummy data source to this API.
//
// Useful for pipeline resolvers
// and for backend changes that don't require a data source.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource {
	var returns NoneDataSource

	_jsii_.Invoke(
		g,
		"addNoneDataSource",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Add a query field to the schema's Query. CDK will create an Object Type called 'Query'. For example,.
//
// type Query {
//    fieldName: Field.returnType
// }
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddQuery(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		g,
		"addQuery",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

// add a new Rds data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource {
	var returns RdsDataSource

	_jsii_.Invoke(
		g,
		"addRdsDataSource",
		[]interface{}{id, serverlessCluster, secretStore, databaseName, options},
		&returns,
	)

	return returns
}

// Add schema dependency to a given construct.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddSchemaDependency(construct awscdk.CfnResource) *bool {
	var returns *bool

	_jsii_.Invoke(
		g,
		"addSchemaDependency",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a subscription field to the schema's Subscription. CDK will create an Object Type called 'Subscription'. For example,.
//
// type Subscription {
//    fieldName: Field.returnType
// }
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddSubscription(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		g,
		"addSubscription",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

// Escape hatch to append to Schema as desired.
//
// Will always result
// in a newline.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddToSchema(addition *string, delimiter *string) {
	_jsii_.InvokeVoid(
		g,
		"addToSchema",
		[]interface{}{addition, delimiter},
	)
}

// Add type to the schema.
// Experimental.
func (g *jsiiProxy_GraphqlApi) AddType(type_ IIntermediateType) IIntermediateType {
	var returns IIntermediateType

	_jsii_.Invoke(
		g,
		"addType",
		[]interface{}{type_},
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (g *jsiiProxy_GraphqlApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (g *jsiiProxy_GraphqlApi) CreateResolver(props *ExtendedResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		g,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GraphqlApi) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
// Experimental.
func (g *jsiiProxy_GraphqlApi) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
// Experimental.
func (g *jsiiProxy_GraphqlApi) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Adds an IAM policy statement associated with this GraphQLApi to an IAM principal's policy.
// Experimental.
func (g *jsiiProxy_GraphqlApi) Grant(grantee awsiam.IGrantable, resources IamResource, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee, resources}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grant",
		args,
		&returns,
	)

	return returns
}

// Adds an IAM policy statement for Mutation access to this GraphQLApi to an IAM principal's policy.
// Experimental.
func (g *jsiiProxy_GraphqlApi) GrantMutation(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantMutation",
		args,
		&returns,
	)

	return returns
}

// Adds an IAM policy statement for Query access to this GraphQLApi to an IAM principal's policy.
// Experimental.
func (g *jsiiProxy_GraphqlApi) GrantQuery(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantQuery",
		args,
		&returns,
	)

	return returns
}

// Adds an IAM policy statement for Subscription access to this GraphQLApi to an IAM principal's policy.
// Experimental.
func (g *jsiiProxy_GraphqlApi) GrantSubscription(grantee awsiam.IGrantable, fields ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantSubscription",
		args,
		&returns,
	)

	return returns
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
func (g *jsiiProxy_GraphqlApi) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GraphqlApi) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
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
func (g *jsiiProxy_GraphqlApi) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (g *jsiiProxy_GraphqlApi) Prepare() {
	_jsii_.InvokeVoid(
		g,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GraphqlApi) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GraphqlApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
func (g *jsiiProxy_GraphqlApi) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Attributes for GraphQL imports.
// Experimental.
type GraphqlApiAttributes struct {
	// an unique AWS AppSync GraphQL API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	// Experimental.
	GraphqlApiId *string `json:"graphqlApiId"`
	// the arn for the GraphQL Api.
	// Experimental.
	GraphqlApiArn *string `json:"graphqlApiArn"`
}

// Base Class for GraphQL API.
// Experimental.
type GraphqlApiBase interface {
	awscdk.Resource
	IGraphqlApi
	ApiId() *string
	Arn() *string
	Env() *awscdk.ResourceEnvironment
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource
	AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource
	AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource
	AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource
	AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource
	AddSchemaDependency(construct awscdk.CfnResource) *bool
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	CreateResolver(props *ExtendedResolverProps) Resolver
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for GraphqlApiBase
type jsiiProxy_GraphqlApiBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IGraphqlApi
}

func (j *jsiiProxy_GraphqlApiBase) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlApiBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewGraphqlApiBase_Override(g GraphqlApiBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.GraphqlApiBase",
		[]interface{}{scope, id, props},
		g,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func GraphqlApiBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlApiBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GraphqlApiBase_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlApiBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// add a new DynamoDB data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource {
	var returns DynamoDbDataSource

	_jsii_.Invoke(
		g,
		"addDynamoDbDataSource",
		[]interface{}{id, table, options},
		&returns,
	)

	return returns
}

// add a new http data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource {
	var returns HttpDataSource

	_jsii_.Invoke(
		g,
		"addHttpDataSource",
		[]interface{}{id, endpoint, options},
		&returns,
	)

	return returns
}

// add a new Lambda data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource {
	var returns LambdaDataSource

	_jsii_.Invoke(
		g,
		"addLambdaDataSource",
		[]interface{}{id, lambdaFunction, options},
		&returns,
	)

	return returns
}

// add a new dummy data source to this API.
//
// Useful for pipeline resolvers
// and for backend changes that don't require a data source.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource {
	var returns NoneDataSource

	_jsii_.Invoke(
		g,
		"addNoneDataSource",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// add a new Rds data source to this API.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource {
	var returns RdsDataSource

	_jsii_.Invoke(
		g,
		"addRdsDataSource",
		[]interface{}{id, serverlessCluster, secretStore, databaseName, options},
		&returns,
	)

	return returns
}

// Add schema dependency if not imported.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) AddSchemaDependency(construct awscdk.CfnResource) *bool {
	var returns *bool

	_jsii_.Invoke(
		g,
		"addSchemaDependency",
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) CreateResolver(props *ExtendedResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		g,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Experimental.
func (g *jsiiProxy_GraphqlApiBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
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
func (g *jsiiProxy_GraphqlApiBase) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
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
func (g *jsiiProxy_GraphqlApiBase) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (g *jsiiProxy_GraphqlApiBase) Prepare() {
	_jsii_.InvokeVoid(
		g,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GraphqlApiBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
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
func (g *jsiiProxy_GraphqlApiBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an AppSync GraphQL API.
// Experimental.
type GraphqlApiProps struct {
	// the name of the GraphQL API.
	// Experimental.
	Name *string `json:"name"`
	// Optional authorization configuration.
	// Experimental.
	AuthorizationConfig *AuthorizationConfig `json:"authorizationConfig"`
	// Logging configuration for this api.
	// Experimental.
	LogConfig *LogConfig `json:"logConfig"`
	// GraphQL schema definition. Specify how you want to define your schema.
	//
	// Schema.fromFile(filePath: string) allows schema definition through schema.graphql file
	// Experimental.
	Schema Schema `json:"schema"`
	// A flag indicating whether or not X-Ray tracing is enabled for the GraphQL API.
	// Experimental.
	XrayEnabled *bool `json:"xrayEnabled"`
}

// The GraphQL Types in AppSync's GraphQL.
//
// GraphQL Types are the
// building blocks for object types, queries, mutations, etc. They are
// types like String, Int, Id or even Object Types you create.
//
// i.e. `String`, `String!`, `[String]`, `[String!]`, `[String]!`
//
// GraphQL Types are used to define the entirety of schema.
// Experimental.
type GraphqlType interface {
	IField
	IntermediateType() IIntermediateType
	IsList() *bool
	IsRequired() *bool
	IsRequiredList() *bool
	Type() Type
	ArgsToString() *string
	DirectivesToString(_modes *[]AuthorizationType) *string
	ToString() *string
}

// The jsii proxy struct for GraphqlType
type jsiiProxy_GraphqlType struct {
	jsiiProxy_IField
}

func (j *jsiiProxy_GraphqlType) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GraphqlType) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewGraphqlType(type_ Type, options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	j := jsiiProxy_GraphqlType{}

	_jsii_.Create(
		"monocdk.aws_appsync.GraphqlType",
		[]interface{}{type_, options},
		&j,
	)

	return &j
}

// Experimental.
func NewGraphqlType_Override(g GraphqlType, type_ Type, options *GraphqlTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.GraphqlType",
		[]interface{}{type_, options},
		g,
	)
}

// `AWSDate` scalar type represents a valid extended `ISO 8601 Date` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DD`. It accepts time zone offsets.
// Experimental.
func GraphqlType_AwsDate(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsDate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSDateTime` scalar type represents a valid extended `ISO 8601 DateTime` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DDThh:mm:ss.sssZ`. It accepts time zone offsets.
// Experimental.
func GraphqlType_AwsDateTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsDateTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSEmail` scalar type represents an email address string (i.e.`username@example.com`).
// Experimental.
func GraphqlType_AwsEmail(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsEmail",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSIPAddress` scalar type respresents a valid `IPv4` of `IPv6` address string.
// Experimental.
func GraphqlType_AwsIpAddress(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsIpAddress",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSJson` scalar type represents a JSON string.
// Experimental.
func GraphqlType_AwsJson(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSPhone` scalar type represents a valid phone number. Phone numbers maybe be whitespace delimited or hyphenated.
//
// The number can specify a country code at the beginning, but is not required for US phone numbers.
// Experimental.
func GraphqlType_AwsPhone(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsPhone",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTime` scalar type represents a valid extended `ISO 8601 Time` string.
//
// In other words, accepts date strings in the form of `hh:mm:ss.sss`. It accepts time zone offsets.
// Experimental.
func GraphqlType_AwsTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTimestamp` scalar type represents the number of seconds since `1970-01-01T00:00Z`.
//
// Timestamps are serialized and deserialized as numbers.
// Experimental.
func GraphqlType_AwsTimestamp(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsTimestamp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSURL` scalar type represetns a valid URL string.
//
// URLs wihtout schemes or contain double slashes are considered invalid.
// Experimental.
func GraphqlType_AwsUrl(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"awsUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Boolean` scalar type is a boolean value: true or false.
// Experimental.
func GraphqlType_Boolean(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"boolean",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Float` scalar type is a signed double-precision fractional value.
// Experimental.
func GraphqlType_Float(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"float",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `ID` scalar type is a unique identifier. `ID` type is serialized similar to `String`.
//
// Often used as a key for a cache and not intended to be human-readable.
// Experimental.
func GraphqlType_Id(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"id",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Int` scalar type is a signed non-fractional numerical value.
// Experimental.
func GraphqlType_Int(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"int",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// an intermediate type to be added as an attribute (i.e. an interface or an object type).
// Experimental.
func GraphqlType_Intermediate(options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"intermediate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `String` scalar type is a free-form human-readable text.
// Experimental.
func GraphqlType_String(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.GraphqlType",
		"string",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the arguments for this field.
// Experimental.
func (g *jsiiProxy_GraphqlType) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the directives for this field.
// Experimental.
func (g *jsiiProxy_GraphqlType) DirectivesToString(_modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"directivesToString",
		[]interface{}{_modes},
		&returns,
	)

	return returns
}

// Generate the string for this attribute.
// Experimental.
func (g *jsiiProxy_GraphqlType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for GraphQL Types.
// Experimental.
type GraphqlTypeOptions struct {
	// property determining if this attribute is a list i.e. if true, attribute would be [Type].
	// Experimental.
	IsList *bool `json:"isList"`
	// property determining if this attribute is non-nullable i.e. if true, attribute would be Type!
	// Experimental.
	IsRequired *bool `json:"isRequired"`
	// property determining if this attribute is a non-nullable list i.e. if true, attribute would be [ Type ]! or if isRequired true, attribe would be [ Type! ]!
	// Experimental.
	IsRequiredList *bool `json:"isRequiredList"`
	// the intermediate type linked to this attribute.
	// Experimental.
	IntermediateType IIntermediateType `json:"intermediateType"`
}

// An AppSync datasource backed by a http endpoint.
// Experimental.
type HttpDataSource interface {
	BackedDataSource
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	Ds() CfnDataSource
	GrantPrincipal() awsiam.IPrincipal
	Name() *string
	Node() awscdk.ConstructNode
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	CreateResolver(props *BaseResolverProps) Resolver
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for HttpDataSource
type jsiiProxy_HttpDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_HttpDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpDataSource(scope constructs.Construct, id *string, props *HttpDataSourceProps) HttpDataSource {
	_init_.Initialize()

	j := jsiiProxy_HttpDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.HttpDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpDataSource_Override(h HttpDataSource, scope constructs.Construct, id *string, props *HttpDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.HttpDataSource",
		[]interface{}{scope, id, props},
		h,
	)
}

func (j *jsiiProxy_HttpDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_HttpDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func HttpDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.HttpDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// creates a new appsync function for this datasource and API using the given properties.
// Experimental.
func (h *jsiiProxy_HttpDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		h,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (h *jsiiProxy_HttpDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		h,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
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
func (h *jsiiProxy_HttpDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		h,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (h *jsiiProxy_HttpDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
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
func (h *jsiiProxy_HttpDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (h *jsiiProxy_HttpDataSource) Prepare() {
	_jsii_.InvokeVoid(
		h,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (h *jsiiProxy_HttpDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		h,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (h *jsiiProxy_HttpDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
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
func (h *jsiiProxy_HttpDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Optional configuration for Http data sources.
// Experimental.
type HttpDataSourceOptions struct {
	// The description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source, overrides the id given by cdk.
	// Experimental.
	Name *string `json:"name"`
	// The authorization config in case the HTTP endpoint requires authorization.
	// Experimental.
	AuthorizationConfig *AwsIamConfig `json:"authorizationConfig"`
}

// Properties for an AppSync http datasource.
// Experimental.
type HttpDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `json:"name"`
	// The http endpoint.
	// Experimental.
	Endpoint *string `json:"endpoint"`
	// The authorization config in case the HTTP endpoint requires authorization.
	// Experimental.
	AuthorizationConfig *AwsIamConfig `json:"authorizationConfig"`
}

// Interface for AppSync Functions.
// Experimental.
type IAppsyncFunction interface {
	awscdk.IResource
	// the ARN of the AppSync function.
	// Experimental.
	FunctionArn() *string
	// the name of this AppSync Function.
	// Experimental.
	FunctionId() *string
}

// The jsii proxy for IAppsyncFunction
type jsiiProxy_IAppsyncFunction struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAppsyncFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppsyncFunction) FunctionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionId",
		&returns,
	)
	return returns
}

// A Graphql Field.
// Experimental.
type IField interface {
	// Generate the arguments for this field.
	// Experimental.
	ArgsToString() *string
	// Generate the directives for this field.
	// Experimental.
	DirectivesToString(modes *[]AuthorizationType) *string
	// Generate the string for this attribute.
	// Experimental.
	ToString() *string
	// The options to make this field resolvable.
	// Experimental.
	FieldOptions() *ResolvableFieldOptions
	// the intermediate type linked to this attribute (i.e. an interface or an object).
	// Experimental.
	IntermediateType() IIntermediateType
	// property determining if this attribute is a list i.e. if true, attribute would be `[Type]`.
	// Experimental.
	IsList() *bool
	// property determining if this attribute is non-nullable i.e. if true, attribute would be `Type!` and this attribute must always have a value.
	// Experimental.
	IsRequired() *bool
	// property determining if this attribute is a non-nullable list i.e. if true, attribute would be `[ Type ]!` and this attribute's list must always have a value.
	// Experimental.
	IsRequiredList() *bool
	// the type of attribute.
	// Experimental.
	Type() Type
}

// The jsii proxy for IField
type jsiiProxy_IField struct {
	_ byte // padding
}

func (i *jsiiProxy_IField) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IField) DirectivesToString(modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"directivesToString",
		[]interface{}{modes},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IField) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IField) FieldOptions() *ResolvableFieldOptions {
	var returns *ResolvableFieldOptions
	_jsii_.Get(
		j,
		"fieldOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IField) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// Interface for GraphQL.
// Experimental.
type IGraphqlApi interface {
	awscdk.IResource
	// add a new DynamoDB data source to this API.
	// Experimental.
	AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource
	// add a new http data source to this API.
	// Experimental.
	AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource
	// add a new Lambda data source to this API.
	// Experimental.
	AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource
	// add a new dummy data source to this API.
	//
	// Useful for pipeline resolvers
	// and for backend changes that don't require a data source.
	// Experimental.
	AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource
	// add a new Rds data source to this API.
	// Experimental.
	AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource
	// Add schema dependency if not imported.
	// Experimental.
	AddSchemaDependency(construct awscdk.CfnResource) *bool
	// creates a new resolver for this datasource and API using the given properties.
	// Experimental.
	CreateResolver(props *ExtendedResolverProps) Resolver
	// an unique AWS AppSync GraphQL API identifier i.e. 'lxz775lwdrgcndgz3nurvac7oa'.
	// Experimental.
	ApiId() *string
	// the ARN of the API.
	// Experimental.
	Arn() *string
}

// The jsii proxy for IGraphqlApi
type jsiiProxy_IGraphqlApi struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGraphqlApi) AddDynamoDbDataSource(id *string, table awsdynamodb.ITable, options *DataSourceOptions) DynamoDbDataSource {
	var returns DynamoDbDataSource

	_jsii_.Invoke(
		i,
		"addDynamoDbDataSource",
		[]interface{}{id, table, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddHttpDataSource(id *string, endpoint *string, options *HttpDataSourceOptions) HttpDataSource {
	var returns HttpDataSource

	_jsii_.Invoke(
		i,
		"addHttpDataSource",
		[]interface{}{id, endpoint, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddLambdaDataSource(id *string, lambdaFunction awslambda.IFunction, options *DataSourceOptions) LambdaDataSource {
	var returns LambdaDataSource

	_jsii_.Invoke(
		i,
		"addLambdaDataSource",
		[]interface{}{id, lambdaFunction, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddNoneDataSource(id *string, options *DataSourceOptions) NoneDataSource {
	var returns NoneDataSource

	_jsii_.Invoke(
		i,
		"addNoneDataSource",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddRdsDataSource(id *string, serverlessCluster awsrds.IServerlessCluster, secretStore awssecretsmanager.ISecret, databaseName *string, options *DataSourceOptions) RdsDataSource {
	var returns RdsDataSource

	_jsii_.Invoke(
		i,
		"addRdsDataSource",
		[]interface{}{id, serverlessCluster, secretStore, databaseName, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) AddSchemaDependency(construct awscdk.CfnResource) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"addSchemaDependency",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGraphqlApi) CreateResolver(props *ExtendedResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		i,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IGraphqlApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

// Intermediate Types are types that includes a certain set of fields that define the entirety of your schema.
// Experimental.
type IIntermediateType interface {
	// Add a field to this Intermediate Type.
	// Experimental.
	AddField(options *AddFieldOptions)
	// Create an GraphQL Type representing this Intermediate Type.
	// Experimental.
	Attribute(options *BaseTypeOptions) GraphqlType
	// Generate the string of this object type.
	// Experimental.
	ToString() *string
	// the attributes of this type.
	// Experimental.
	Definition() *map[string]IField
	// the directives for this object type.
	// Experimental.
	Directives() *[]Directive
	// The Interface Types this Intermediate Type implements.
	// Experimental.
	InterfaceTypes() *[]InterfaceType
	// the intermediate type linked to this attribute (i.e. an interface or an object).
	// Experimental.
	IntermediateType() IIntermediateType
	// the name of this type.
	// Experimental.
	Name() *string
	// The resolvers linked to this data source.
	// Experimental.
	Resolvers() *[]Resolver
	// The resolvers linked to this data source.
	// Experimental.
	SetResolvers(r *[]Resolver)
}

// The jsii proxy for IIntermediateType
type jsiiProxy_IIntermediateType struct {
	_ byte // padding
}

func (i *jsiiProxy_IIntermediateType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		i,
		"addField",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IIntermediateType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		i,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIntermediateType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIntermediateType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) Directives() *[]Directive {
	var returns *[]Directive
	_jsii_.Get(
		j,
		"directives",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) InterfaceTypes() *[]InterfaceType {
	var returns *[]InterfaceType
	_jsii_.Get(
		j,
		"interfaceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) Resolvers() *[]Resolver {
	var returns *[]Resolver
	_jsii_.Get(
		j,
		"resolvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIntermediateType) SetResolvers(val *[]Resolver) {
	_jsii_.Set(
		j,
		"resolvers",
		val,
	)
}

// A class used to generate resource arns for AppSync.
// Experimental.
type IamResource interface {
	ResourceArns(api GraphqlApi) *[]*string
}

// The jsii proxy struct for IamResource
type jsiiProxy_IamResource struct {
	_ byte // padding
}

// Generate the resource names that accepts all types: `*`.
// Experimental.
func IamResource_All() IamResource {
	_init_.Initialize()

	var returns IamResource

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.IamResource",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the resource names given custom arns.
// Experimental.
func IamResource_Custom(arns ...*string) IamResource {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range arns {
		args = append(args, a)
	}

	var returns IamResource

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.IamResource",
		"custom",
		args,
		&returns,
	)

	return returns
}

// Generate the resource names given a type and fields.
// Experimental.
func IamResource_OfType(type_ *string, fields ...*string) IamResource {
	_init_.Initialize()

	args := []interface{}{type_}
	for _, a := range fields {
		args = append(args, a)
	}

	var returns IamResource

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.IamResource",
		"ofType",
		args,
		&returns,
	)

	return returns
}

// Return the Resource ARN.
// Experimental.
func (i *jsiiProxy_IamResource) ResourceArns(api GraphqlApi) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"resourceArns",
		[]interface{}{api},
		&returns,
	)

	return returns
}

// Input Types are abstract types that define complex objects.
//
// They are used in arguments to represent
// Experimental.
type InputType interface {
	IIntermediateType
	Definition() *map[string]IField
	Modes() *[]AuthorizationType
	SetModes(val *[]AuthorizationType)
	Name() *string
	AddField(options *AddFieldOptions)
	Attribute(options *BaseTypeOptions) GraphqlType
	ToString() *string
}

// The jsii proxy struct for InputType
type jsiiProxy_InputType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_InputType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewInputType(name *string, props *IntermediateTypeOptions) InputType {
	_init_.Initialize()

	j := jsiiProxy_InputType{}

	_jsii_.Create(
		"monocdk.aws_appsync.InputType",
		[]interface{}{name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewInputType_Override(i InputType, name *string, props *IntermediateTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.InputType",
		[]interface{}{name, props},
		i,
	)
}

func (j *jsiiProxy_InputType) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

// Add a field to this Input Type.
//
// Input Types must have both fieldName and field options.
// Experimental.
func (i *jsiiProxy_InputType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		i,
		"addField",
		[]interface{}{options},
	)
}

// Create a GraphQL Type representing this Input Type.
// Experimental.
func (i *jsiiProxy_InputType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		i,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the string of this input type.
// Experimental.
func (i *jsiiProxy_InputType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Interface Types are abstract types that includes a certain set of fields that other types must include if they implement the interface.
// Experimental.
type InterfaceType interface {
	IIntermediateType
	Definition() *map[string]IField
	Directives() *[]Directive
	Modes() *[]AuthorizationType
	SetModes(val *[]AuthorizationType)
	Name() *string
	AddField(options *AddFieldOptions)
	Attribute(options *BaseTypeOptions) GraphqlType
	ToString() *string
}

// The jsii proxy struct for InterfaceType
type jsiiProxy_InterfaceType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_InterfaceType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceType) Directives() *[]Directive {
	var returns *[]Directive
	_jsii_.Get(
		j,
		"directives",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InterfaceType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewInterfaceType(name *string, props *IntermediateTypeOptions) InterfaceType {
	_init_.Initialize()

	j := jsiiProxy_InterfaceType{}

	_jsii_.Create(
		"monocdk.aws_appsync.InterfaceType",
		[]interface{}{name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewInterfaceType_Override(i InterfaceType, name *string, props *IntermediateTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.InterfaceType",
		[]interface{}{name, props},
		i,
	)
}

func (j *jsiiProxy_InterfaceType) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

// Add a field to this Interface Type.
//
// Interface Types must have both fieldName and field options.
// Experimental.
func (i *jsiiProxy_InterfaceType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		i,
		"addField",
		[]interface{}{options},
	)
}

// Create a GraphQL Type representing this Intermediate Type.
// Experimental.
func (i *jsiiProxy_InterfaceType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		i,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the string of this object type.
// Experimental.
func (i *jsiiProxy_InterfaceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for configuring an Intermediate Type.
// Experimental.
type IntermediateTypeOptions struct {
	// the attributes of this type.
	// Experimental.
	Definition *map[string]IField `json:"definition"`
	// the directives for this object type.
	// Experimental.
	Directives *[]Directive `json:"directives"`
}

// Factory class for DynamoDB key conditions.
// Experimental.
type KeyCondition interface {
	And(keyCond KeyCondition) KeyCondition
	RenderTemplate() *string
}

// The jsii proxy struct for KeyCondition
type jsiiProxy_KeyCondition struct {
	_ byte // padding
}

// Condition (k, arg).
//
// True if the key attribute k begins with the Query argument.
// Experimental.
func KeyCondition_BeginsWith(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.KeyCondition",
		"beginsWith",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k BETWEEN arg1 AND arg2, true if k >= arg1 and k <= arg2.
// Experimental.
func KeyCondition_Between(keyName *string, arg1 *string, arg2 *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.KeyCondition",
		"between",
		[]interface{}{keyName, arg1, arg2},
		&returns,
	)

	return returns
}

// Condition k = arg, true if the key attribute k is equal to the Query argument.
// Experimental.
func KeyCondition_Eq(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.KeyCondition",
		"eq",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k >= arg, true if the key attribute k is greater or equal to the Query argument.
// Experimental.
func KeyCondition_Ge(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.KeyCondition",
		"ge",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k > arg, true if the key attribute k is greater than the the Query argument.
// Experimental.
func KeyCondition_Gt(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.KeyCondition",
		"gt",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k <= arg, true if the key attribute k is less than or equal to the Query argument.
// Experimental.
func KeyCondition_Le(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.KeyCondition",
		"le",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Condition k < arg, true if the key attribute k is less than the Query argument.
// Experimental.
func KeyCondition_Lt(keyName *string, arg *string) KeyCondition {
	_init_.Initialize()

	var returns KeyCondition

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.KeyCondition",
		"lt",
		[]interface{}{keyName, arg},
		&returns,
	)

	return returns
}

// Conjunction between two conditions.
// Experimental.
func (k *jsiiProxy_KeyCondition) And(keyCond KeyCondition) KeyCondition {
	var returns KeyCondition

	_jsii_.Invoke(
		k,
		"and",
		[]interface{}{keyCond},
		&returns,
	)

	return returns
}

// Renders the key condition to a VTL string.
// Experimental.
func (k *jsiiProxy_KeyCondition) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An AppSync datasource backed by a Lambda function.
// Experimental.
type LambdaDataSource interface {
	BackedDataSource
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	Ds() CfnDataSource
	GrantPrincipal() awsiam.IPrincipal
	Name() *string
	Node() awscdk.ConstructNode
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	CreateResolver(props *BaseResolverProps) Resolver
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for LambdaDataSource
type jsiiProxy_LambdaDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_LambdaDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaDataSource(scope constructs.Construct, id *string, props *LambdaDataSourceProps) LambdaDataSource {
	_init_.Initialize()

	j := jsiiProxy_LambdaDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.LambdaDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaDataSource_Override(l LambdaDataSource, scope constructs.Construct, id *string, props *LambdaDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.LambdaDataSource",
		[]interface{}{scope, id, props},
		l,
	)
}

func (j *jsiiProxy_LambdaDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_LambdaDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func LambdaDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.LambdaDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// creates a new appsync function for this datasource and API using the given properties.
// Experimental.
func (l *jsiiProxy_LambdaDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		l,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (l *jsiiProxy_LambdaDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		l,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
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
func (l *jsiiProxy_LambdaDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
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
func (l *jsiiProxy_LambdaDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (l *jsiiProxy_LambdaDataSource) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (l *jsiiProxy_LambdaDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
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
func (l *jsiiProxy_LambdaDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an AppSync Lambda datasource.
// Experimental.
type LambdaDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `json:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Experimental.
	ServiceRole awsiam.IRole `json:"serviceRole"`
	// The Lambda function to call to interact with this data source.
	// Experimental.
	LambdaFunction awslambda.IFunction `json:"lambdaFunction"`
}

// Logging configuration for AppSync.
// Experimental.
type LogConfig struct {
	// exclude verbose content.
	// Experimental.
	ExcludeVerboseContent interface{} `json:"excludeVerboseContent"`
	// log level for fields.
	// Experimental.
	FieldLogLevel FieldLogLevel `json:"fieldLogLevel"`
	// The role for CloudWatch Logs.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// MappingTemplates for AppSync resolvers.
// Experimental.
type MappingTemplate interface {
	RenderTemplate() *string
}

// The jsii proxy struct for MappingTemplate
type jsiiProxy_MappingTemplate struct {
	_ byte // padding
}

// Experimental.
func NewMappingTemplate_Override(m MappingTemplate) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.MappingTemplate",
		nil, // no parameters
		m,
	)
}

// Mapping template to delete a single item from a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbDeleteItem(keyName *string, idArg *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"dynamoDbDeleteItem",
		[]interface{}{keyName, idArg},
		&returns,
	)

	return returns
}

// Mapping template to get a single item from a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbGetItem(keyName *string, idArg *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"dynamoDbGetItem",
		[]interface{}{keyName, idArg},
		&returns,
	)

	return returns
}

// Mapping template to save a single item to a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbPutItem(key PrimaryKey, values AttributeValues) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"dynamoDbPutItem",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Mapping template to query a set of items from a DynamoDB table.
// Experimental.
func MappingTemplate_DynamoDbQuery(cond KeyCondition, indexName *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"dynamoDbQuery",
		[]interface{}{cond, indexName},
		&returns,
	)

	return returns
}

// Mapping template for a single result item from DynamoDB.
// Experimental.
func MappingTemplate_DynamoDbResultItem() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"dynamoDbResultItem",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Mapping template for a result list from DynamoDB.
// Experimental.
func MappingTemplate_DynamoDbResultList() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"dynamoDbResultList",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Mapping template to scan a DynamoDB table to fetch all entries.
// Experimental.
func MappingTemplate_DynamoDbScanTable() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"dynamoDbScanTable",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create a mapping template from the given file.
// Experimental.
func MappingTemplate_FromFile(fileName *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"fromFile",
		[]interface{}{fileName},
		&returns,
	)

	return returns
}

// Create a mapping template from the given string.
// Experimental.
func MappingTemplate_FromString(template *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"fromString",
		[]interface{}{template},
		&returns,
	)

	return returns
}

// Mapping template to invoke a Lambda function.
// Experimental.
func MappingTemplate_LambdaRequest(payload *string) MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"lambdaRequest",
		[]interface{}{payload},
		&returns,
	)

	return returns
}

// Mapping template to return the Lambda result to the caller.
// Experimental.
func MappingTemplate_LambdaResult() MappingTemplate {
	_init_.Initialize()

	var returns MappingTemplate

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.MappingTemplate",
		"lambdaResult",
		nil, // no parameters
		&returns,
	)

	return returns
}

// this is called to render the mapping template to a VTL string.
// Experimental.
func (m *jsiiProxy_MappingTemplate) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An AppSync dummy datasource.
// Experimental.
type NoneDataSource interface {
	BaseDataSource
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	Ds() CfnDataSource
	Name() *string
	Node() awscdk.ConstructNode
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	CreateResolver(props *BaseResolverProps) Resolver
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for NoneDataSource
type jsiiProxy_NoneDataSource struct {
	jsiiProxy_BaseDataSource
}

func (j *jsiiProxy_NoneDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NoneDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewNoneDataSource(scope constructs.Construct, id *string, props *NoneDataSourceProps) NoneDataSource {
	_init_.Initialize()

	j := jsiiProxy_NoneDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.NoneDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNoneDataSource_Override(n NoneDataSource, scope constructs.Construct, id *string, props *NoneDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.NoneDataSource",
		[]interface{}{scope, id, props},
		n,
	)
}

func (j *jsiiProxy_NoneDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_NoneDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func NoneDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.NoneDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// creates a new appsync function for this datasource and API using the given properties.
// Experimental.
func (n *jsiiProxy_NoneDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		n,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (n *jsiiProxy_NoneDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		n,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
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
func (n *jsiiProxy_NoneDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		n,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (n *jsiiProxy_NoneDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NoneDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (n *jsiiProxy_NoneDataSource) Prepare() {
	_jsii_.InvokeVoid(
		n,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (n *jsiiProxy_NoneDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NoneDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NoneDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an AppSync dummy datasource.
// Experimental.
type NoneDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `json:"name"`
}

// Object Types are types declared by you.
// Experimental.
type ObjectType interface {
	InterfaceType
	IIntermediateType
	Definition() *map[string]IField
	Directives() *[]Directive
	InterfaceTypes() *[]InterfaceType
	Modes() *[]AuthorizationType
	SetModes(val *[]AuthorizationType)
	Name() *string
	Resolvers() *[]Resolver
	SetResolvers(val *[]Resolver)
	AddField(options *AddFieldOptions)
	Attribute(options *BaseTypeOptions) GraphqlType
	GenerateResolver(api IGraphqlApi, fieldName *string, options *ResolvableFieldOptions) Resolver
	ToString() *string
}

// The jsii proxy struct for ObjectType
type jsiiProxy_ObjectType struct {
	jsiiProxy_InterfaceType
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_ObjectType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Directives() *[]Directive {
	var returns *[]Directive
	_jsii_.Get(
		j,
		"directives",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) InterfaceTypes() *[]InterfaceType {
	var returns *[]InterfaceType
	_jsii_.Get(
		j,
		"interfaceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObjectType) Resolvers() *[]Resolver {
	var returns *[]Resolver
	_jsii_.Get(
		j,
		"resolvers",
		&returns,
	)
	return returns
}


// Experimental.
func NewObjectType(name *string, props *ObjectTypeOptions) ObjectType {
	_init_.Initialize()

	j := jsiiProxy_ObjectType{}

	_jsii_.Create(
		"monocdk.aws_appsync.ObjectType",
		[]interface{}{name, props},
		&j,
	)

	return &j
}

// Experimental.
func NewObjectType_Override(o ObjectType, name *string, props *ObjectTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.ObjectType",
		[]interface{}{name, props},
		o,
	)
}

func (j *jsiiProxy_ObjectType) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

func (j *jsiiProxy_ObjectType) SetResolvers(val *[]Resolver) {
	_jsii_.Set(
		j,
		"resolvers",
		val,
	)
}

// Add a field to this Object Type.
//
// Object Types must have both fieldName and field options.
// Experimental.
func (o *jsiiProxy_ObjectType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		o,
		"addField",
		[]interface{}{options},
	)
}

// Create a GraphQL Type representing this Intermediate Type.
// Experimental.
func (o *jsiiProxy_ObjectType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		o,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the resolvers linked to this Object Type.
// Experimental.
func (o *jsiiProxy_ObjectType) GenerateResolver(api IGraphqlApi, fieldName *string, options *ResolvableFieldOptions) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		o,
		"generateResolver",
		[]interface{}{api, fieldName, options},
		&returns,
	)

	return returns
}

// Generate the string of this object type.
// Experimental.
func (o *jsiiProxy_ObjectType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for configuring an Object Type.
// Experimental.
type ObjectTypeOptions struct {
	// the attributes of this type.
	// Experimental.
	Definition *map[string]IField `json:"definition"`
	// the directives for this object type.
	// Experimental.
	Directives *[]Directive `json:"directives"`
	// The Interface Types this Object Type implements.
	// Experimental.
	InterfaceTypes *[]InterfaceType `json:"interfaceTypes"`
}

// Configuration for OpenID Connect authorization in AppSync.
// Experimental.
type OpenIdConnectConfig struct {
	// The issuer for the OIDC configuration.
	//
	// The issuer returned by discovery must exactly match the value of `iss` in the OIDC token.
	// Experimental.
	OidcProvider *string `json:"oidcProvider"`
	// The client identifier of the Relying party at the OpenID identity provider.
	//
	// A regular expression can be specified so AppSync can validate against multiple client identifiers at a time.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	ClientId *string `json:"clientId"`
	// The number of milliseconds an OIDC token is valid after being authenticated by OIDC provider.
	//
	// `auth_time` claim in OIDC token is required for this validation to work.
	// Experimental.
	TokenExpiryFromAuth *float64 `json:"tokenExpiryFromAuth"`
	// The number of milliseconds an OIDC token is valid after being issued to a user.
	//
	// This validation uses `iat` claim of OIDC token.
	// Experimental.
	TokenExpiryFromIssue *float64 `json:"tokenExpiryFromIssue"`
}

// Specifies the assignment to the partition key.
//
// It can be
// enhanced with the assignment of the sort key.
// Experimental.
type PartitionKey interface {
	PrimaryKey
	Pkey() Assign
	RenderTemplate() *string
	Sort(key *string) SortKeyStep
}

// The jsii proxy struct for PartitionKey
type jsiiProxy_PartitionKey struct {
	jsiiProxy_PrimaryKey
}

func (j *jsiiProxy_PartitionKey) Pkey() Assign {
	var returns Assign
	_jsii_.Get(
		j,
		"pkey",
		&returns,
	)
	return returns
}


// Experimental.
func NewPartitionKey(pkey Assign) PartitionKey {
	_init_.Initialize()

	j := jsiiProxy_PartitionKey{}

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKey",
		[]interface{}{pkey},
		&j,
	)

	return &j
}

// Experimental.
func NewPartitionKey_Override(p PartitionKey, pkey Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKey",
		[]interface{}{pkey},
		p,
	)
}

// Allows assigning a value to the partition key.
// Experimental.
func PartitionKey_Partition(key *string) PartitionKeyStep {
	_init_.Initialize()

	var returns PartitionKeyStep

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.PartitionKey",
		"partition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Renders the key assignment to a VTL string.
// Experimental.
func (p *jsiiProxy_PartitionKey) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows assigning a value to the sort key.
// Experimental.
func (p *jsiiProxy_PartitionKey) Sort(key *string) SortKeyStep {
	var returns SortKeyStep

	_jsii_.Invoke(
		p,
		"sort",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Utility class to allow assigning a value or an auto-generated id to a partition key.
// Experimental.
type PartitionKeyStep interface {
	Auto() PartitionKey
	Is(val *string) PartitionKey
}

// The jsii proxy struct for PartitionKeyStep
type jsiiProxy_PartitionKeyStep struct {
	_ byte // padding
}

// Experimental.
func NewPartitionKeyStep(key *string) PartitionKeyStep {
	_init_.Initialize()

	j := jsiiProxy_PartitionKeyStep{}

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKeyStep",
		[]interface{}{key},
		&j,
	)

	return &j
}

// Experimental.
func NewPartitionKeyStep_Override(p PartitionKeyStep, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.PartitionKeyStep",
		[]interface{}{key},
		p,
	)
}

// Assign an auto-generated value to the partition key.
// Experimental.
func (p *jsiiProxy_PartitionKeyStep) Auto() PartitionKey {
	var returns PartitionKey

	_jsii_.Invoke(
		p,
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Assign an auto-generated value to the partition key.
// Experimental.
func (p *jsiiProxy_PartitionKeyStep) Is(val *string) PartitionKey {
	var returns PartitionKey

	_jsii_.Invoke(
		p,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}

// Specifies the assignment to the primary key.
//
// It either
// contains the full primary key or only the partition key.
// Experimental.
type PrimaryKey interface {
	Pkey() Assign
	RenderTemplate() *string
}

// The jsii proxy struct for PrimaryKey
type jsiiProxy_PrimaryKey struct {
	_ byte // padding
}

func (j *jsiiProxy_PrimaryKey) Pkey() Assign {
	var returns Assign
	_jsii_.Get(
		j,
		"pkey",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrimaryKey(pkey Assign, skey Assign) PrimaryKey {
	_init_.Initialize()

	j := jsiiProxy_PrimaryKey{}

	_jsii_.Create(
		"monocdk.aws_appsync.PrimaryKey",
		[]interface{}{pkey, skey},
		&j,
	)

	return &j
}

// Experimental.
func NewPrimaryKey_Override(p PrimaryKey, pkey Assign, skey Assign) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.PrimaryKey",
		[]interface{}{pkey, skey},
		p,
	)
}

// Allows assigning a value to the partition key.
// Experimental.
func PrimaryKey_Partition(key *string) PartitionKeyStep {
	_init_.Initialize()

	var returns PartitionKeyStep

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.PrimaryKey",
		"partition",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Renders the key assignment to a VTL string.
// Experimental.
func (p *jsiiProxy_PrimaryKey) RenderTemplate() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"renderTemplate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An AppSync datasource backed by RDS.
// Experimental.
type RdsDataSource interface {
	BackedDataSource
	Api() IGraphqlApi
	SetApi(val IGraphqlApi)
	Ds() CfnDataSource
	GrantPrincipal() awsiam.IPrincipal
	Name() *string
	Node() awscdk.ConstructNode
	ServiceRole() awsiam.IRole
	SetServiceRole(val awsiam.IRole)
	CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction
	CreateResolver(props *BaseResolverProps) Resolver
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for RdsDataSource
type jsiiProxy_RdsDataSource struct {
	jsiiProxy_BackedDataSource
}

func (j *jsiiProxy_RdsDataSource) Api() IGraphqlApi {
	var returns IGraphqlApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) Ds() CfnDataSource {
	var returns CfnDataSource
	_jsii_.Get(
		j,
		"ds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDataSource) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewRdsDataSource(scope constructs.Construct, id *string, props *RdsDataSourceProps) RdsDataSource {
	_init_.Initialize()

	j := jsiiProxy_RdsDataSource{}

	_jsii_.Create(
		"monocdk.aws_appsync.RdsDataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRdsDataSource_Override(r RdsDataSource, scope constructs.Construct, id *string, props *RdsDataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.RdsDataSource",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RdsDataSource) SetApi(val IGraphqlApi) {
	_jsii_.Set(
		j,
		"api",
		val,
	)
}

func (j *jsiiProxy_RdsDataSource) SetServiceRole(val awsiam.IRole) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func RdsDataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.RdsDataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// creates a new appsync function for this datasource and API using the given properties.
// Experimental.
func (r *jsiiProxy_RdsDataSource) CreateFunction(props *BaseAppsyncFunctionProps) AppsyncFunction {
	var returns AppsyncFunction

	_jsii_.Invoke(
		r,
		"createFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// creates a new resolver for this datasource and API using the given properties.
// Experimental.
func (r *jsiiProxy_RdsDataSource) CreateResolver(props *BaseResolverProps) Resolver {
	var returns Resolver

	_jsii_.Invoke(
		r,
		"createResolver",
		[]interface{}{props},
		&returns,
	)

	return returns
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
func (r *jsiiProxy_RdsDataSource) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_RdsDataSource) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_RdsDataSource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (r *jsiiProxy_RdsDataSource) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_RdsDataSource) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (r *jsiiProxy_RdsDataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_RdsDataSource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for an AppSync RDS datasource.
// Experimental.
type RdsDataSourceProps struct {
	// The API to attach this data source to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
	// the description of the data source.
	// Experimental.
	Description *string `json:"description"`
	// The name of the data source.
	// Experimental.
	Name *string `json:"name"`
	// The IAM service role to be assumed by AppSync to interact with the data source.
	// Experimental.
	ServiceRole awsiam.IRole `json:"serviceRole"`
	// The secret containing the credentials for the database.
	// Experimental.
	SecretStore awssecretsmanager.ISecret `json:"secretStore"`
	// The serverless cluster to call to interact with this data source.
	// Experimental.
	ServerlessCluster awsrds.IServerlessCluster `json:"serverlessCluster"`
	// The name of the database to use within the cluster.
	// Experimental.
	DatabaseName *string `json:"databaseName"`
}

// Resolvable Fields build upon Graphql Types and provide fields that can resolve into operations on a data source.
// Experimental.
type ResolvableField interface {
	Field
	IField
	FieldOptions() *ResolvableFieldOptions
	IntermediateType() IIntermediateType
	IsList() *bool
	IsRequired() *bool
	IsRequiredList() *bool
	Type() Type
	ArgsToString() *string
	DirectivesToString(modes *[]AuthorizationType) *string
	ToString() *string
}

// The jsii proxy struct for ResolvableField
type jsiiProxy_ResolvableField struct {
	jsiiProxy_Field
	jsiiProxy_IField
}

func (j *jsiiProxy_ResolvableField) FieldOptions() *ResolvableFieldOptions {
	var returns *ResolvableFieldOptions
	_jsii_.Get(
		j,
		"fieldOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IntermediateType() IIntermediateType {
	var returns IIntermediateType
	_jsii_.Get(
		j,
		"intermediateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IsList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IsRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) IsRequiredList() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isRequiredList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResolvableField) Type() Type {
	var returns Type
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewResolvableField(options *ResolvableFieldOptions) ResolvableField {
	_init_.Initialize()

	j := jsiiProxy_ResolvableField{}

	_jsii_.Create(
		"monocdk.aws_appsync.ResolvableField",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewResolvableField_Override(r ResolvableField, options *ResolvableFieldOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.ResolvableField",
		[]interface{}{options},
		r,
	)
}

// `AWSDate` scalar type represents a valid extended `ISO 8601 Date` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DD`. It accepts time zone offsets.
// Experimental.
func ResolvableField_AwsDate(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsDate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSDateTime` scalar type represents a valid extended `ISO 8601 DateTime` string.
//
// In other words, accepts date strings in the form of `YYYY-MM-DDThh:mm:ss.sssZ`. It accepts time zone offsets.
// Experimental.
func ResolvableField_AwsDateTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsDateTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSEmail` scalar type represents an email address string (i.e.`username@example.com`).
// Experimental.
func ResolvableField_AwsEmail(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsEmail",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSIPAddress` scalar type respresents a valid `IPv4` of `IPv6` address string.
// Experimental.
func ResolvableField_AwsIpAddress(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsIpAddress",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSJson` scalar type represents a JSON string.
// Experimental.
func ResolvableField_AwsJson(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsJson",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSPhone` scalar type represents a valid phone number. Phone numbers maybe be whitespace delimited or hyphenated.
//
// The number can specify a country code at the beginning, but is not required for US phone numbers.
// Experimental.
func ResolvableField_AwsPhone(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsPhone",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTime` scalar type represents a valid extended `ISO 8601 Time` string.
//
// In other words, accepts date strings in the form of `hh:mm:ss.sss`. It accepts time zone offsets.
// Experimental.
func ResolvableField_AwsTime(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSTimestamp` scalar type represents the number of seconds since `1970-01-01T00:00Z`.
//
// Timestamps are serialized and deserialized as numbers.
// Experimental.
func ResolvableField_AwsTimestamp(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsTimestamp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `AWSURL` scalar type represetns a valid URL string.
//
// URLs wihtout schemes or contain double slashes are considered invalid.
// Experimental.
func ResolvableField_AwsUrl(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"awsUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Boolean` scalar type is a boolean value: true or false.
// Experimental.
func ResolvableField_Boolean(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"boolean",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Float` scalar type is a signed double-precision fractional value.
// Experimental.
func ResolvableField_Float(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"float",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `ID` scalar type is a unique identifier. `ID` type is serialized similar to `String`.
//
// Often used as a key for a cache and not intended to be human-readable.
// Experimental.
func ResolvableField_Id(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"id",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `Int` scalar type is a signed non-fractional numerical value.
// Experimental.
func ResolvableField_Int(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"int",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// an intermediate type to be added as an attribute (i.e. an interface or an object type).
// Experimental.
func ResolvableField_Intermediate(options *GraphqlTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"intermediate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// `String` scalar type is a free-form human-readable text.
// Experimental.
func ResolvableField_String(options *BaseTypeOptions) GraphqlType {
	_init_.Initialize()

	var returns GraphqlType

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.ResolvableField",
		"string",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the args string of this resolvable field.
// Experimental.
func (r *jsiiProxy_ResolvableField) ArgsToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"argsToString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Generate the directives for this field.
// Experimental.
func (r *jsiiProxy_ResolvableField) DirectivesToString(modes *[]AuthorizationType) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"directivesToString",
		[]interface{}{modes},
		&returns,
	)

	return returns
}

// Generate the string for this attribute.
// Experimental.
func (r *jsiiProxy_ResolvableField) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for configuring a resolvable field.
// Experimental.
type ResolvableFieldOptions struct {
	// The return type for this field.
	// Experimental.
	ReturnType GraphqlType `json:"returnType"`
	// The arguments for this field.
	//
	// i.e. type Example (first: String second: String) {}
	// - where 'first' and 'second' are key values for args
	// and 'String' is the GraphqlType
	// Experimental.
	Args *map[string]GraphqlType `json:"args"`
	// the directives for this field.
	// Experimental.
	Directives *[]Directive `json:"directives"`
	// The data source creating linked to this resolvable field.
	// Experimental.
	DataSource BaseDataSource `json:"dataSource"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `json:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `json:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `json:"responseMappingTemplate"`
}

// An AppSync resolver.
// Experimental.
type Resolver interface {
	awscdk.Construct
	Arn() *string
	Node() awscdk.ConstructNode
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Resolver
type jsiiProxy_Resolver struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_Resolver) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resolver) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewResolver(scope constructs.Construct, id *string, props *ResolverProps) Resolver {
	_init_.Initialize()

	j := jsiiProxy_Resolver{}

	_jsii_.Create(
		"monocdk.aws_appsync.Resolver",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewResolver_Override(r Resolver, scope constructs.Construct, id *string, props *ResolverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.Resolver",
		[]interface{}{scope, id, props},
		r,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Resolver_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Resolver",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
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
func (r *jsiiProxy_Resolver) OnPrepare() {
	_jsii_.InvokeVoid(
		r,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_Resolver) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
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
func (r *jsiiProxy_Resolver) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
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
func (r *jsiiProxy_Resolver) Prepare() {
	_jsii_.InvokeVoid(
		r,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (r *jsiiProxy_Resolver) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		r,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (r *jsiiProxy_Resolver) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_Resolver) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Additional property for an AppSync resolver for GraphQL API reference.
// Experimental.
type ResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	// Experimental.
	FieldName *string `json:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	// Experimental.
	TypeName *string `json:"typeName"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `json:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `json:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `json:"responseMappingTemplate"`
	// The data source this resolver is using.
	// Experimental.
	DataSource BaseDataSource `json:"dataSource"`
	// The API this resolver is attached to.
	// Experimental.
	Api IGraphqlApi `json:"api"`
}

// The Schema for a GraphQL Api.
//
// If no options are configured, schema will be generated
// code-first.
// Experimental.
type Schema interface {
	Definition() *string
	SetDefinition(val *string)
	AddMutation(fieldName *string, field ResolvableField) ObjectType
	AddQuery(fieldName *string, field ResolvableField) ObjectType
	AddSubscription(fieldName *string, field ResolvableField) ObjectType
	AddToSchema(addition *string, delimiter *string)
	AddType(type_ IIntermediateType) IIntermediateType
	Bind(api GraphqlApi) CfnGraphQLSchema
}

// The jsii proxy struct for Schema
type jsiiProxy_Schema struct {
	_ byte // padding
}

func (j *jsiiProxy_Schema) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}


// Experimental.
func NewSchema(options *SchemaOptions) Schema {
	_init_.Initialize()

	j := jsiiProxy_Schema{}

	_jsii_.Create(
		"monocdk.aws_appsync.Schema",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewSchema_Override(s Schema, options *SchemaOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.Schema",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_Schema) SetDefinition(val *string) {
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

// Generate a Schema from file.
//
// Returns: `SchemaAsset` with immutable schema defintion
// Experimental.
func Schema_FromAsset(filePath *string) Schema {
	_init_.Initialize()

	var returns Schema

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Schema",
		"fromAsset",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Add a mutation field to the schema's Mutation. CDK will create an Object Type called 'Mutation'. For example,.
//
// type Mutation {
//    fieldName: Field.returnType
// }
// Experimental.
func (s *jsiiProxy_Schema) AddMutation(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addMutation",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

// Add a query field to the schema's Query. CDK will create an Object Type called 'Query'. For example,.
//
// type Query {
//    fieldName: Field.returnType
// }
// Experimental.
func (s *jsiiProxy_Schema) AddQuery(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addQuery",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

// Add a subscription field to the schema's Subscription. CDK will create an Object Type called 'Subscription'. For example,.
//
// type Subscription {
//    fieldName: Field.returnType
// }
// Experimental.
func (s *jsiiProxy_Schema) AddSubscription(fieldName *string, field ResolvableField) ObjectType {
	var returns ObjectType

	_jsii_.Invoke(
		s,
		"addSubscription",
		[]interface{}{fieldName, field},
		&returns,
	)

	return returns
}

// Escape hatch to add to Schema as desired.
//
// Will always result
// in a newline.
// Experimental.
func (s *jsiiProxy_Schema) AddToSchema(addition *string, delimiter *string) {
	_jsii_.InvokeVoid(
		s,
		"addToSchema",
		[]interface{}{addition, delimiter},
	)
}

// Add type to the schema.
// Experimental.
func (s *jsiiProxy_Schema) AddType(type_ IIntermediateType) IIntermediateType {
	var returns IIntermediateType

	_jsii_.Invoke(
		s,
		"addType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Called when the GraphQL Api is initialized to allow this object to bind to the stack.
// Experimental.
func (s *jsiiProxy_Schema) Bind(api GraphqlApi) CfnGraphQLSchema {
	var returns CfnGraphQLSchema

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{api},
		&returns,
	)

	return returns
}

// The options for configuring a schema.
//
// If no options are specified, then the schema will
// be generated code-first.
// Experimental.
type SchemaOptions struct {
	// The file path for the schema.
	//
	// When this option is
	// configured, then the schema will be generated from an
	// existing file from disk.
	// Experimental.
	FilePath *string `json:"filePath"`
}

// Utility class to allow assigning a value or an auto-generated id to a sort key.
// Experimental.
type SortKeyStep interface {
	Auto() PrimaryKey
	Is(val *string) PrimaryKey
}

// The jsii proxy struct for SortKeyStep
type jsiiProxy_SortKeyStep struct {
	_ byte // padding
}

// Experimental.
func NewSortKeyStep(pkey Assign, skey *string) SortKeyStep {
	_init_.Initialize()

	j := jsiiProxy_SortKeyStep{}

	_jsii_.Create(
		"monocdk.aws_appsync.SortKeyStep",
		[]interface{}{pkey, skey},
		&j,
	)

	return &j
}

// Experimental.
func NewSortKeyStep_Override(s SortKeyStep, pkey Assign, skey *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.SortKeyStep",
		[]interface{}{pkey, skey},
		s,
	)
}

// Assign an auto-generated value to the sort key.
// Experimental.
func (s *jsiiProxy_SortKeyStep) Auto() PrimaryKey {
	var returns PrimaryKey

	_jsii_.Invoke(
		s,
		"auto",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Assign an auto-generated value to the sort key.
// Experimental.
func (s *jsiiProxy_SortKeyStep) Is(val *string) PrimaryKey {
	var returns PrimaryKey

	_jsii_.Invoke(
		s,
		"is",
		[]interface{}{val},
		&returns,
	)

	return returns
}

// Enum containing the Types that can be used to define ObjectTypes.
// Experimental.
type Type string

const (
	Type_ID Type = "ID"
	Type_STRING Type = "STRING"
	Type_INT Type = "INT"
	Type_FLOAT Type = "FLOAT"
	Type_BOOLEAN Type = "BOOLEAN"
	Type_AWS_DATE Type = "AWS_DATE"
	Type_AWS_TIME Type = "AWS_TIME"
	Type_AWS_DATE_TIME Type = "AWS_DATE_TIME"
	Type_AWS_TIMESTAMP Type = "AWS_TIMESTAMP"
	Type_AWS_EMAIL Type = "AWS_EMAIL"
	Type_AWS_JSON Type = "AWS_JSON"
	Type_AWS_URL Type = "AWS_URL"
	Type_AWS_PHONE Type = "AWS_PHONE"
	Type_AWS_IP_ADDRESS Type = "AWS_IP_ADDRESS"
	Type_INTERMEDIATE Type = "INTERMEDIATE"
)

// Union Types are abstract types that are similar to Interface Types, but they cannot to specify any common fields between types.
//
// Note that fields of a union type need to be object types. In other words,
// you can't create a union type out of interfaces, other unions, or inputs.
// Experimental.
type UnionType interface {
	IIntermediateType
	Definition() *map[string]IField
	Modes() *[]AuthorizationType
	SetModes(val *[]AuthorizationType)
	Name() *string
	AddField(options *AddFieldOptions)
	Attribute(options *BaseTypeOptions) GraphqlType
	ToString() *string
}

// The jsii proxy struct for UnionType
type jsiiProxy_UnionType struct {
	jsiiProxy_IIntermediateType
}

func (j *jsiiProxy_UnionType) Definition() *map[string]IField {
	var returns *map[string]IField
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UnionType) Modes() *[]AuthorizationType {
	var returns *[]AuthorizationType
	_jsii_.Get(
		j,
		"modes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UnionType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewUnionType(name *string, options *UnionTypeOptions) UnionType {
	_init_.Initialize()

	j := jsiiProxy_UnionType{}

	_jsii_.Create(
		"monocdk.aws_appsync.UnionType",
		[]interface{}{name, options},
		&j,
	)

	return &j
}

// Experimental.
func NewUnionType_Override(u UnionType, name *string, options *UnionTypeOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.UnionType",
		[]interface{}{name, options},
		u,
	)
}

func (j *jsiiProxy_UnionType) SetModes(val *[]AuthorizationType) {
	_jsii_.Set(
		j,
		"modes",
		val,
	)
}

// Add a field to this Union Type.
//
// Input Types must have field options and the IField must be an Object Type.
// Experimental.
func (u *jsiiProxy_UnionType) AddField(options *AddFieldOptions) {
	_jsii_.InvokeVoid(
		u,
		"addField",
		[]interface{}{options},
	)
}

// Create a GraphQL Type representing this Union Type.
// Experimental.
func (u *jsiiProxy_UnionType) Attribute(options *BaseTypeOptions) GraphqlType {
	var returns GraphqlType

	_jsii_.Invoke(
		u,
		"attribute",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Generate the string of this Union type.
// Experimental.
func (u *jsiiProxy_UnionType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for configuring an Union Type.
// Experimental.
type UnionTypeOptions struct {
	// the object types for this union type.
	// Experimental.
	Definition *[]IIntermediateType `json:"definition"`
}

// Configuration for Cognito user-pools in AppSync.
// Experimental.
type UserPoolConfig struct {
	// The Cognito user pool to use as identity source.
	// Experimental.
	UserPool awscognito.IUserPool `json:"userPool"`
	// the optional app id regex.
	// Experimental.
	AppIdClientRegex *string `json:"appIdClientRegex"`
	// Default auth action.
	// Experimental.
	DefaultAction UserPoolDefaultAction `json:"defaultAction"`
}

// enum with all possible values for Cognito user-pool default actions.
// Experimental.
type UserPoolDefaultAction string

const (
	UserPoolDefaultAction_ALLOW UserPoolDefaultAction = "ALLOW"
	UserPoolDefaultAction_DENY UserPoolDefaultAction = "DENY"
)

// Factory class for attribute value assignments.
// Experimental.
type Values interface {
}

// The jsii proxy struct for Values
type jsiiProxy_Values struct {
	_ byte // padding
}

// Experimental.
func NewValues() Values {
	_init_.Initialize()

	j := jsiiProxy_Values{}

	_jsii_.Create(
		"monocdk.aws_appsync.Values",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewValues_Override(v Values) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appsync.Values",
		nil, // no parameters
		v,
	)
}

// Allows assigning a value to the specified attribute.
// Experimental.
func Values_Attribute(attr *string) AttributeValuesStep {
	_init_.Initialize()

	var returns AttributeValuesStep

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Values",
		"attribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

// Treats the specified object as a map of assignments, where the property names represent attribute names.
//
// It’s opinionated about how it represents
// some of the nested objects: e.g., it will use lists (“L”) rather than sets
// (“SS”, “NS”, “BS”). By default it projects the argument container ("$ctx.args").
// Experimental.
func Values_Projecting(arg *string) AttributeValues {
	_init_.Initialize()

	var returns AttributeValues

	_jsii_.StaticInvoke(
		"monocdk.aws_appsync.Values",
		"projecting",
		[]interface{}{arg},
		&returns,
	)

	return returns
}

