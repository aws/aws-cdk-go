package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A new alias to a particular version of a Lambda function.
//
// TODO: EXAMPLE
//
type Alias interface {
	QualifiedFunctionBase
	IAlias
	AliasName() *string
	Architecture() Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	Lambda() IFunction
	LatestVersion() IVersion
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Qualifier() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	Version() IVersion
	AddAutoScaling(options *AutoScalingOptions) IScalableFunctionAttribute
	AddEventSource(source IEventSource)
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	AddPermission(id *string, permission *Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for Alias
type jsiiProxy_Alias struct {
	jsiiProxy_QualifiedFunctionBase
	jsiiProxy_IAlias
}

func (j *jsiiProxy_Alias) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Lambda() IFunction {
	var returns IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Alias) Version() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewAlias(scope constructs.Construct, id *string, props *AliasProps) Alias {
	_init_.Initialize()

	j := jsiiProxy_Alias{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Alias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAlias_Override(a Alias, scope constructs.Construct, id *string, props *AliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Alias",
		[]interface{}{scope, id, props},
		a,
	)
}

func Alias_FromAliasAttributes(scope constructs.Construct, id *string, attrs *AliasAttributes) IAlias {
	_init_.Initialize()

	var returns IAlias

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Alias",
		"fromAliasAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Alias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Alias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Alias_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Alias",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Configure provisioned concurrency autoscaling on a function alias.
//
// Returns a scalable attribute that can call
// `scaleOnUtilization()` and `scaleOnSchedule()`.
func (a *jsiiProxy_Alias) AddAutoScaling(options *AutoScalingOptions) IScalableFunctionAttribute {
	var returns IScalableFunctionAttribute

	_jsii_.Invoke(
		a,
		"addAutoScaling",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
//
// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
//
// The following example adds an SQS Queue as an event source:
// ```
// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
// myFunction.addEventSource(new SqsEventSource(myQueue));
// ```
func (a *jsiiProxy_Alias) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		a,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (a *jsiiProxy_Alias) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		a,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
func (a *jsiiProxy_Alias) AddPermission(id *string, permission *Permission) {
	_jsii_.InvokeVoid(
		a,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (a *jsiiProxy_Alias) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		a,
		"addToRolePolicy",
		[]interface{}{statement},
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
func (a *jsiiProxy_Alias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (a *jsiiProxy_Alias) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		a,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (a *jsiiProxy_Alias) GeneratePhysicalName() *string {
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
func (a *jsiiProxy_Alias) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (a *jsiiProxy_Alias) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (a *jsiiProxy_Alias) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
func (a *jsiiProxy_Alias) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
func (a *jsiiProxy_Alias) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (a *jsiiProxy_Alias) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
func (a *jsiiProxy_Alias) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
func (a *jsiiProxy_Alias) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (a *jsiiProxy_Alias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type AliasAttributes struct {
	AliasName *string `json:"aliasName"`
	AliasVersion IVersion `json:"aliasVersion"`
}

// Options for `lambda.Alias`.
//
// TODO: EXAMPLE
//
type AliasOptions struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// Additional versions with individual weights this alias points to.
	//
	// Individual additional version weights specified here should add up to
	// (less than) one. All remaining weight is routed to the default
	// version.
	//
	// For example, the config is
	//
	//     version: "1"
	//     additionalVersions: [{ version: "2", weight: 0.05 }]
	//
	// Then 5% of traffic will be routed to function version 2, while
	// the remaining 95% of traffic will be routed to function version 1.
	AdditionalVersions *[]*VersionWeight `json:"additionalVersions"`
	// Description for the alias.
	Description *string `json:"description"`
	// Specifies a provisioned concurrency configuration for a function's alias.
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions"`
}

// Properties for a new Lambda alias.
//
// TODO: EXAMPLE
//
type AliasProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// Additional versions with individual weights this alias points to.
	//
	// Individual additional version weights specified here should add up to
	// (less than) one. All remaining weight is routed to the default
	// version.
	//
	// For example, the config is
	//
	//     version: "1"
	//     additionalVersions: [{ version: "2", weight: 0.05 }]
	//
	// Then 5% of traffic will be routed to function version 2, while
	// the remaining 95% of traffic will be routed to function version 1.
	AdditionalVersions *[]*VersionWeight `json:"additionalVersions"`
	// Description for the alias.
	Description *string `json:"description"`
	// Specifies a provisioned concurrency configuration for a function's alias.
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions"`
	// Name of this alias.
	AliasName *string `json:"aliasName"`
	// Function version this alias refers to.
	//
	// Use lambda.addVersion() to obtain a new lambda version to refer to.
	Version IVersion `json:"version"`
}

// Architectures supported by AWS Lambda.
//
// TODO: EXAMPLE
//
type Architecture interface {
	DockerPlatform() *string
	Name() *string
}

// The jsii proxy struct for Architecture
type jsiiProxy_Architecture struct {
	_ byte // padding
}

func (j *jsiiProxy_Architecture) DockerPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Architecture) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Used to specify a custom architecture name.
//
// Use this if the architecture name is not yet supported by the CDK.
func Architecture_Custom(name *string, dockerPlatform *string) Architecture {
	_init_.Initialize()

	var returns Architecture

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Architecture",
		"custom",
		[]interface{}{name, dockerPlatform},
		&returns,
	)

	return returns
}

func Architecture_ARM_64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Architecture",
		"ARM_64",
		&returns,
	)
	return returns
}

func Architecture_X86_64() Architecture {
	_init_.Initialize()
	var returns Architecture
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Architecture",
		"X86_64",
		&returns,
	)
	return returns
}

// Lambda code from a local directory.
//
// TODO: EXAMPLE
//
type AssetCode interface {
	Code
	IsInline() *bool
	Path() *string
	Bind(scope constructs.Construct) *CodeConfig
	BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions)
}

// The jsii proxy struct for AssetCode
type jsiiProxy_AssetCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_AssetCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetCode) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}


func NewAssetCode(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	j := jsiiProxy_AssetCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.AssetCode",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetCode_Override(a AssetCode, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.AssetCode",
		[]interface{}{path, options},
		a,
	)
}

// Loads the function code from a local disk path.
func AssetCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func AssetCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func AssetCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`
func AssetCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func AssetCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func AssetCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func AssetCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
func (a *jsiiProxy_AssetCode) Bind(scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Called after the CFN function resource has been created to allow the code class to bind to it.
//
// Specifically it's required to allow assets to add
// metadata for tooling like SAM CLI to be able to find their origins.
func (a *jsiiProxy_AssetCode) BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		a,
		"bindToResource",
		[]interface{}{resource, options},
	)
}

// Represents an ECR image that will be constructed from the specified asset and can be bound as Lambda code.
//
// TODO: EXAMPLE
//
type AssetImageCode interface {
	Code
	IsInline() *bool
	Bind(scope constructs.Construct) *CodeConfig
	BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions)
}

// The jsii proxy struct for AssetImageCode
type jsiiProxy_AssetImageCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_AssetImageCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewAssetImageCode(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	j := jsiiProxy_AssetImageCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		[]interface{}{directory, props},
		&j,
	)

	return &j
}

func NewAssetImageCode_Override(a AssetImageCode, directory *string, props *AssetImageCodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		[]interface{}{directory, props},
		a,
	)
}

// Loads the function code from a local disk path.
func AssetImageCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func AssetImageCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func AssetImageCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`
func AssetImageCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func AssetImageCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func AssetImageCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func AssetImageCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
func (a *jsiiProxy_AssetImageCode) Bind(scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Called after the CFN function resource has been created to allow the code class to bind to it.
//
// Specifically it's required to allow assets to add
// metadata for tooling like SAM CLI to be able to find their origins.
func (a *jsiiProxy_AssetImageCode) BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		a,
		"bindToResource",
		[]interface{}{resource, options},
	)
}

// Properties to initialize a new AssetImage.
//
// TODO: EXAMPLE
//
type AssetImageCodeProps struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `json:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *awsecrassets.DockerImageAssetInvalidationOptions `json:"invalidation"`
	// Docker target to build to.
	Target *string `json:"target"`
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	Cmd *[]*string `json:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	Entrypoint *[]*string `json:"entrypoint"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	WorkingDirectory *string `json:"workingDirectory"`
}

// Properties for enabling Lambda autoscaling.
//
// TODO: EXAMPLE
//
type AutoScalingOptions struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `json:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `json:"minCapacity"`
}

// A CloudFormation `AWS::Lambda::Alias`.
//
// The `AWS::Lambda::Alias` resource creates an [alias](https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html) for a Lambda function version. Use aliases to provide clients with a function identifier that you can update to invoke a different version.
//
// You can also map an alias to split invocation requests between two versions. Use the `RoutingConfig` parameter to specify a second version and the percentage of invocation requests that it receives.
//
// TODO: EXAMPLE
//
type CfnAlias interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionVersion() *string
	SetFunctionVersion(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	ProvisionedConcurrencyConfig() interface{}
	SetProvisionedConcurrencyConfig(val interface{})
	Ref() *string
	RoutingConfig() interface{}
	SetRoutingConfig(val interface{})
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

// The jsii proxy struct for CfnAlias
type jsiiProxy_CfnAlias struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnAlias) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) FunctionVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) ProvisionedConcurrencyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedConcurrencyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) RoutingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAlias) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::Alias`.
func NewCfnAlias(scope constructs.Construct, id *string, props *CfnAliasProps) CfnAlias {
	_init_.Initialize()

	j := jsiiProxy_CfnAlias{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::Alias`.
func NewCfnAlias_Override(c CfnAlias, scope constructs.Construct, id *string, props *CfnAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnAlias",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAlias) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnAlias) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnAlias) SetFunctionVersion(val *string) {
	_jsii_.Set(
		j,
		"functionVersion",
		val,
	)
}

func (j *jsiiProxy_CfnAlias) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnAlias) SetProvisionedConcurrencyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"provisionedConcurrencyConfig",
		val,
	)
}

func (j *jsiiProxy_CfnAlias) SetRoutingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"routingConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAlias_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnAlias",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnAlias_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnAlias",
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
func CfnAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAlias_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnAlias",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnAlias) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnAlias) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnAlias) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnAlias) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnAlias) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnAlias) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnAlias) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnAlias) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnAlias) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnAlias) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAlias) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnAlias) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAlias) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The [traffic-shifting](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) configuration of a Lambda function alias.
//
// TODO: EXAMPLE
//
type CfnAlias_AliasRoutingConfigurationProperty struct {
	// The second version, and the percentage of traffic that's routed to it.
	AdditionalVersionWeights interface{} `json:"additionalVersionWeights"`
}

// A provisioned concurrency configuration for a function's alias.
//
// TODO: EXAMPLE
//
type CfnAlias_ProvisionedConcurrencyConfigurationProperty struct {
	// The amount of provisioned concurrency to allocate for the alias.
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions"`
}

// The [traffic-shifting](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) configuration of a Lambda function alias.
//
// TODO: EXAMPLE
//
type CfnAlias_VersionWeightProperty struct {
	// The qualifier of the second version.
	FunctionVersion *string `json:"functionVersion"`
	// The percentage of traffic that the alias routes to the second version.
	FunctionWeight *float64 `json:"functionWeight"`
}

// Properties for defining a `CfnAlias`.
//
// TODO: EXAMPLE
//
type CfnAliasProps struct {
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName *string `json:"functionName"`
	// The function version that the alias invokes.
	FunctionVersion *string `json:"functionVersion"`
	// The name of the alias.
	Name *string `json:"name"`
	// A description of the alias.
	Description *string `json:"description"`
	// Specifies a [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html) configuration for a function's alias.
	ProvisionedConcurrencyConfig interface{} `json:"provisionedConcurrencyConfig"`
	// The [routing configuration](https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html) of the alias.
	RoutingConfig interface{} `json:"routingConfig"`
}

// A CloudFormation `AWS::Lambda::CodeSigningConfig`.
//
// Details about a [Code signing configuration](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html) .
//
// TODO: EXAMPLE
//
type CfnCodeSigningConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AllowedPublishers() interface{}
	SetAllowedPublishers(val interface{})
	AttrCodeSigningConfigArn() *string
	AttrCodeSigningConfigId() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CodeSigningPolicies() interface{}
	SetCodeSigningPolicies(val interface{})
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LogicalId() *string
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

// The jsii proxy struct for CfnCodeSigningConfig
type jsiiProxy_CfnCodeSigningConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCodeSigningConfig) AllowedPublishers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedPublishers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) AttrCodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCodeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) AttrCodeSigningConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCodeSigningConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) CodeSigningPolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeSigningPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCodeSigningConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::CodeSigningConfig`.
func NewCfnCodeSigningConfig(scope constructs.Construct, id *string, props *CfnCodeSigningConfigProps) CfnCodeSigningConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnCodeSigningConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnCodeSigningConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::CodeSigningConfig`.
func NewCfnCodeSigningConfig_Override(c CfnCodeSigningConfig, scope constructs.Construct, id *string, props *CfnCodeSigningConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnCodeSigningConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCodeSigningConfig) SetAllowedPublishers(val interface{}) {
	_jsii_.Set(
		j,
		"allowedPublishers",
		val,
	)
}

func (j *jsiiProxy_CfnCodeSigningConfig) SetCodeSigningPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"codeSigningPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnCodeSigningConfig) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCodeSigningConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnCodeSigningConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCodeSigningConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnCodeSigningConfig",
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
func CfnCodeSigningConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnCodeSigningConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCodeSigningConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnCodeSigningConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnCodeSigningConfig) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCodeSigningConfig) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCodeSigningConfig) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnCodeSigningConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnCodeSigningConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnCodeSigningConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnCodeSigningConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCodeSigningConfig) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCodeSigningConfig) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCodeSigningConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnCodeSigningConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCodeSigningConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCodeSigningConfig) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnCodeSigningConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCodeSigningConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// List of signing profiles that can sign a code package.
//
// TODO: EXAMPLE
//
type CfnCodeSigningConfig_AllowedPublishersProperty struct {
	// The Amazon Resource Name (ARN) for each of the signing profiles.
	//
	// A signing profile defines a trusted user who can sign a code package.
	SigningProfileVersionArns *[]*string `json:"signingProfileVersionArns"`
}

// Code signing configuration [policies](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html#config-codesigning-policies) specify the validation failure action for signature mismatch or expiry.
//
// TODO: EXAMPLE
//
type CfnCodeSigningConfig_CodeSigningPoliciesProperty struct {
	// Code signing configuration policy for deployment validation failure.
	//
	// If you set the policy to `Enforce` , Lambda blocks the deployment request if signature validation checks fail. If you set the policy to `Warn` , Lambda allows the deployment and creates a CloudWatch log.
	//
	// Default value: `Warn`
	UntrustedArtifactOnDeployment *string `json:"untrustedArtifactOnDeployment"`
}

// Properties for defining a `CfnCodeSigningConfig`.
//
// TODO: EXAMPLE
//
type CfnCodeSigningConfigProps struct {
	// List of allowed publishers.
	AllowedPublishers interface{} `json:"allowedPublishers"`
	// The code signing policy controls the validation failure action for signature mismatch or expiry.
	CodeSigningPolicies interface{} `json:"codeSigningPolicies"`
	// Code signing configuration description.
	Description *string `json:"description"`
}

// A CloudFormation `AWS::Lambda::EventInvokeConfig`.
//
// The `AWS::Lambda::EventInvokeConfig` resource configures options for [asynchronous invocation](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html) on a version or an alias.
//
// By default, Lambda retries an asynchronous invocation twice if the function returns an error. It retains events in a queue for up to six hours. When an event fails all processing attempts or stays in the asynchronous invocation queue for too long, Lambda discards it.
//
// TODO: EXAMPLE
//
type CfnEventInvokeConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DestinationConfig() interface{}
	SetDestinationConfig(val interface{})
	FunctionName() *string
	SetFunctionName(val *string)
	LogicalId() *string
	MaximumEventAgeInSeconds() *float64
	SetMaximumEventAgeInSeconds(val *float64)
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	Node() constructs.Node
	Qualifier() *string
	SetQualifier(val *string)
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

// The jsii proxy struct for CfnEventInvokeConfig
type jsiiProxy_CfnEventInvokeConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventInvokeConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) DestinationConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) MaximumEventAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumEventAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventInvokeConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::EventInvokeConfig`.
func NewCfnEventInvokeConfig(scope constructs.Construct, id *string, props *CfnEventInvokeConfigProps) CfnEventInvokeConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnEventInvokeConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnEventInvokeConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::EventInvokeConfig`.
func NewCfnEventInvokeConfig_Override(c CfnEventInvokeConfig, scope constructs.Construct, id *string, props *CfnEventInvokeConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnEventInvokeConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventInvokeConfig) SetDestinationConfig(val interface{}) {
	_jsii_.Set(
		j,
		"destinationConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEventInvokeConfig) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnEventInvokeConfig) SetMaximumEventAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumEventAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnEventInvokeConfig) SetMaximumRetryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_CfnEventInvokeConfig) SetQualifier(val *string) {
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventInvokeConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventInvokeConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventInvokeConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventInvokeConfig",
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
func CfnEventInvokeConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventInvokeConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventInvokeConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnEventInvokeConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEventInvokeConfig) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEventInvokeConfig) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEventInvokeConfig) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventInvokeConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEventInvokeConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEventInvokeConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventInvokeConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEventInvokeConfig) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEventInvokeConfig) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEventInvokeConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEventInvokeConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventInvokeConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEventInvokeConfig) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnEventInvokeConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventInvokeConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A configuration object that specifies the destination of an event after Lambda processes it.
//
// TODO: EXAMPLE
//
type CfnEventInvokeConfig_DestinationConfigProperty struct {
	// The destination configuration for failed invocations.
	OnFailure interface{} `json:"onFailure"`
	// The destination configuration for successful invocations.
	OnSuccess interface{} `json:"onSuccess"`
}

// A destination for events that failed processing.
//
// TODO: EXAMPLE
//
type CfnEventInvokeConfig_OnFailureProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `json:"destination"`
}

// A destination for events that were processed successfully.
//
// TODO: EXAMPLE
//
type CfnEventInvokeConfig_OnSuccessProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `json:"destination"`
}

// Properties for defining a `CfnEventInvokeConfig`.
//
// TODO: EXAMPLE
//
type CfnEventInvokeConfigProps struct {
	// The name of the Lambda function.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `64`
	//
	// *Pattern* : `([a-zA-Z0-9-_]+)`
	FunctionName *string `json:"functionName"`
	// The identifier of a version or alias.
	//
	// - *Version* - A version number.
	// - *Alias* - An alias name.
	// - *Latest* - To specify the unpublished version, use `$LATEST` .
	Qualifier *string `json:"qualifier"`
	// A destination for events after they have been sent to a function for processing.
	//
	// **Destinations** - *Function* - The Amazon Resource Name (ARN) of a Lambda function.
	// - *Queue* - The ARN of an SQS queue.
	// - *Topic* - The ARN of an SNS topic.
	// - *Event Bus* - The ARN of an Amazon EventBridge event bus.
	DestinationConfig interface{} `json:"destinationConfig"`
	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *float64 `json:"maximumEventAgeInSeconds"`
	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts"`
}

// A CloudFormation `AWS::Lambda::EventSourceMapping`.
//
// The `AWS::Lambda::EventSourceMapping` resource creates a mapping between an event source and an AWS Lambda function. Lambda reads items from the event source and triggers the function.
//
// For details about each event source type, see the following topics. In particular, each of the topics describes the required and optional parameters for the specific event source.
//
// - [Configuring a Dynamo DB stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-dynamodb-eventsourcemapping)
// - [Configuring a Kinesis stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-eventsourcemapping)
// - [Configuring an SQS queue as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-eventsource)
// - [Configuring an MQ broker as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-eventsourcemapping)
// - [Configuring MSK as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html)
// - [Configuring Self-Managed Apache Kafka as an event source](https://docs.aws.amazon.com/lambda/latest/dg/kafka-smaa.html)
//
// TODO: EXAMPLE
//
type CfnEventSourceMapping interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrId() *string
	BatchSize() *float64
	SetBatchSize(val *float64)
	BisectBatchOnFunctionError() interface{}
	SetBisectBatchOnFunctionError(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DestinationConfig() interface{}
	SetDestinationConfig(val interface{})
	Enabled() interface{}
	SetEnabled(val interface{})
	EventSourceArn() *string
	SetEventSourceArn(val *string)
	FilterCriteria() interface{}
	SetFilterCriteria(val interface{})
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionResponseTypes() *[]*string
	SetFunctionResponseTypes(val *[]*string)
	LogicalId() *string
	MaximumBatchingWindowInSeconds() *float64
	SetMaximumBatchingWindowInSeconds(val *float64)
	MaximumRecordAgeInSeconds() *float64
	SetMaximumRecordAgeInSeconds(val *float64)
	MaximumRetryAttempts() *float64
	SetMaximumRetryAttempts(val *float64)
	Node() constructs.Node
	ParallelizationFactor() *float64
	SetParallelizationFactor(val *float64)
	Queues() *[]*string
	SetQueues(val *[]*string)
	Ref() *string
	SelfManagedEventSource() interface{}
	SetSelfManagedEventSource(val interface{})
	SourceAccessConfigurations() interface{}
	SetSourceAccessConfigurations(val interface{})
	Stack() awscdk.Stack
	StartingPosition() *string
	SetStartingPosition(val *string)
	StartingPositionTimestamp() *float64
	SetStartingPositionTimestamp(val *float64)
	Topics() *[]*string
	SetTopics(val *[]*string)
	TumblingWindowInSeconds() *float64
	SetTumblingWindowInSeconds(val *float64)
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

// The jsii proxy struct for CfnEventSourceMapping
type jsiiProxy_CfnEventSourceMapping struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEventSourceMapping) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) BisectBatchOnFunctionError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) DestinationConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) EventSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) FilterCriteria() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) FunctionResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Queues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) SelfManagedEventSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedEventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) SourceAccessConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAccessConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) StartingPositionTimestamp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startingPositionTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) TumblingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventSourceMapping) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMapping(scope constructs.Construct, id *string, props *CfnEventSourceMappingProps) CfnEventSourceMapping {
	_init_.Initialize()

	j := jsiiProxy_CfnEventSourceMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::EventSourceMapping`.
func NewCfnEventSourceMapping_Override(c CfnEventSourceMapping, scope constructs.Construct, id *string, props *CfnEventSourceMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetBatchSize(val *float64) {
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetBisectBatchOnFunctionError(val interface{}) {
	_jsii_.Set(
		j,
		"bisectBatchOnFunctionError",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetDestinationConfig(val interface{}) {
	_jsii_.Set(
		j,
		"destinationConfig",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetEventSourceArn(val *string) {
	_jsii_.Set(
		j,
		"eventSourceArn",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetFilterCriteria(val interface{}) {
	_jsii_.Set(
		j,
		"filterCriteria",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetFunctionResponseTypes(val *[]*string) {
	_jsii_.Set(
		j,
		"functionResponseTypes",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetMaximumBatchingWindowInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetMaximumRecordAgeInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetMaximumRetryAttempts(val *float64) {
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetParallelizationFactor(val *float64) {
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetQueues(val *[]*string) {
	_jsii_.Set(
		j,
		"queues",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetSelfManagedEventSource(val interface{}) {
	_jsii_.Set(
		j,
		"selfManagedEventSource",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetSourceAccessConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"sourceAccessConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetStartingPosition(val *string) {
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetStartingPositionTimestamp(val *float64) {
	_jsii_.Set(
		j,
		"startingPositionTimestamp",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetTopics(val *[]*string) {
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_CfnEventSourceMapping) SetTumblingWindowInSeconds(val *float64) {
	_jsii_.Set(
		j,
		"tumblingWindowInSeconds",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnEventSourceMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnEventSourceMapping_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
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
func CfnEventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventSourceMapping_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnEventSourceMapping",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnEventSourceMapping) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEventSourceMapping) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEventSourceMapping) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventSourceMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnEventSourceMapping) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnEventSourceMapping) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnEventSourceMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEventSourceMapping) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEventSourceMapping) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEventSourceMapping) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnEventSourceMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEventSourceMapping) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEventSourceMapping) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnEventSourceMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEventSourceMapping) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A configuration object that specifies the destination of an event after Lambda processes it.
//
// TODO: EXAMPLE
//
type CfnEventSourceMapping_DestinationConfigProperty struct {
	// The destination configuration for failed invocations.
	OnFailure interface{} `json:"onFailure"`
}

// The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]` .
//
// TODO: EXAMPLE
//
type CfnEventSourceMapping_EndpointsProperty struct {
	// The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]` .
	KafkaBootstrapServers *[]*string `json:"kafkaBootstrapServers"`
}

// A destination for events that failed processing.
//
// TODO: EXAMPLE
//
type CfnEventSourceMapping_OnFailureProperty struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `json:"destination"`
}

// The self-managed Apache Kafka cluster for your event source.
//
// TODO: EXAMPLE
//
type CfnEventSourceMapping_SelfManagedEventSourceProperty struct {
	// The list of bootstrap servers for your Kafka brokers in the following format: `"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]` .
	Endpoints interface{} `json:"endpoints"`
}

// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
//
// TODO: EXAMPLE
//
type CfnEventSourceMapping_SourceAccessConfigurationProperty struct {
	// The type of authentication protocol, VPC components, or virtual host for your event source. For example: `"Type":"SASL_SCRAM_512_AUTH"` .
	//
	// - `BASIC_AUTH` - (Amazon MQ) The AWS Secrets Manager secret that stores your broker credentials.
	// - `BASIC_AUTH` - (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key used for SASL/PLAIN authentication of your Apache Kafka brokers.
	// - `VPC_SUBNET` - The subnets associated with your VPC. Lambda connects to these subnets to fetch data from your self-managed Apache Kafka cluster.
	// - `VPC_SECURITY_GROUP` - The VPC security group used to manage access to your self-managed Apache Kafka brokers.
	// - `SASL_SCRAM_256_AUTH` - The Secrets Manager ARN of your secret key used for SASL SCRAM-256 authentication of your self-managed Apache Kafka brokers.
	// - `SASL_SCRAM_512_AUTH` - The Secrets Manager ARN of your secret key used for SASL SCRAM-512 authentication of your self-managed Apache Kafka brokers.
	// - `VIRTUAL_HOST` - (Amazon MQ) The name of the virtual host in your RabbitMQ broker. Lambda uses this RabbitMQ host as the event source. This property cannot be specified in an UpdateEventSourceMapping API call.
	// - `CLIENT_CERTIFICATE_TLS_AUTH` - (Amazon MSK, Self-managed Apache Kafka) The Secrets Manager ARN of your secret key containing the certificate chain (X.509 PEM), private key (PKCS#8 PEM), and private key password (optional) used for mutual TLS authentication of your MSK/Apache Kafka brokers.
	// - `SERVER_ROOT_CA_CERTIFICATE` - (Self-managed Apache Kafka) The Secrets Manager ARN of your secret key containing the root CA certificate (X.509 PEM) used for TLS encryption of your Apache Kafka brokers.
	Type *string `json:"type"`
	// The value for your chosen configuration in `Type` .
	//
	// For example: `"URI": "arn:aws:secretsmanager:us-east-1:01234567890:secret:MyBrokerSecretName"` .
	Uri *string `json:"uri"`
}

// Properties for defining a `CfnEventSourceMapping`.
//
// TODO: EXAMPLE
//
type CfnEventSourceMappingProps struct {
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Version or Alias ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.
	FunctionName *string `json:"functionName"`
	// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function.
	//
	// Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
	//
	// - *Amazon Kinesis* - Default 100. Max 10,000.
	// - *Amazon DynamoDB Streams* - Default 100. Max 1,000.
	// - *Amazon Simple Queue Service* - Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
	// - *Amazon Managed Streaming for Apache Kafka* - Default 100. Max 10,000.
	// - *Self-Managed Apache Kafka* - Default 100. Max 10,000.
	// - *Amazon MQ (ActiveMQ and RabbitMQ)* - Default 100. Max 10,000.
	BatchSize *float64 `json:"batchSize"`
	// (Streams only) If the function returns an error, split the batch in two and retry.
	//
	// The default value is false.
	BisectBatchOnFunctionError interface{} `json:"bisectBatchOnFunctionError"`
	// (Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	DestinationConfig interface{} `json:"destinationConfig"`
	// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.
	//
	// Default: True
	Enabled interface{} `json:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// - *Amazon Kinesis* - The ARN of the data stream or a stream consumer.
	// - *Amazon DynamoDB Streams* - The ARN of the stream.
	// - *Amazon Simple Queue Service* - The ARN of the queue.
	// - *Amazon Managed Streaming for Apache Kafka* - The ARN of the cluster.
	EventSourceArn *string `json:"eventSourceArn"`
	// (Streams and Amazon SQS) An object that defines the filter criteria that determine whether Lambda should process an event.
	//
	// For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) .
	FilterCriteria interface{} `json:"filterCriteria"`
	// (Streams and SQS) A list of current response type enums applied to the event source mapping.
	//
	// Valid Values: `ReportBatchItemFailures`
	FunctionResponseTypes *[]*string `json:"functionResponseTypes"`
	// (Streams and Amazon SQS standard queues) The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
	//
	// Default: 0
	//
	// Related setting: When you set `BatchSize` to a value greater than 10, you must set `MaximumBatchingWindowInSeconds` to at least 1.
	MaximumBatchingWindowInSeconds *float64 `json:"maximumBatchingWindowInSeconds"`
	// (Streams only) Discard records older than the specified age.
	//
	// The default value is -1,
	// which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
	MaximumRecordAgeInSeconds *float64 `json:"maximumRecordAgeInSeconds"`
	// (Streams only) Discard records after the specified number of retries.
	//
	// The default value is -1,
	// which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts"`
	// (Streams only) The number of batches to process concurrently from each shard.
	//
	// The default value is 1.
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues *[]*string `json:"queues"`
	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource interface{} `json:"selfManagedEventSource"`
	// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
	SourceAccessConfigurations interface{} `json:"sourceAccessConfigurations"`
	// The position in a stream from which to start reading.
	//
	// Required for Amazon Kinesis, Amazon DynamoDB, and Amazon MSK Streams sources.
	//
	// - *LATEST* - Read only new records.
	// - *TRIM_HORIZON* - Process all available records.
	StartingPosition *string `json:"startingPosition"`
	// With `StartingPosition` set to `AT_TIMESTAMP` , the time from which to start reading, in Unix time seconds.
	StartingPositionTimestamp *float64 `json:"startingPositionTimestamp"`
	// The name of the Kafka topic.
	Topics *[]*string `json:"topics"`
	// (Streams only) The duration in seconds of a processing window.
	//
	// The range is between 1 second up to 900 seconds.
	TumblingWindowInSeconds *float64 `json:"tumblingWindowInSeconds"`
}

// A CloudFormation `AWS::Lambda::Function`.
//
// The `AWS::Lambda::Function` resource creates a Lambda function. To create a function, you need a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and an [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html) . The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.
//
// You set the package type to `Image` if the deployment package is a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) . For a container image, the code property must include the URI of a container image in the Amazon ECR registry. You do not need to specify the handler and runtime properties.
//
// You set the package type to `Zip` if the deployment package is a [.zip file archive](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip) . For a .zip file archive, the code property specifies the location of the .zip file. You must also specify the handler and runtime properties. For a Python example, see [Deploy Python Lambda functions with .zip file archives](https://docs.aws.amazon.com/lambda/latest/dg/python-package.html) .
//
// You can use [code signing](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html) if your deployment package is a .zip file archive. To enable code signing for this function, specify the ARN of a code-signing configuration. When a user attempts to deploy a code package with `UpdateFunctionCode` , Lambda checks that the code package has a valid signature from a trusted publisher. The code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.
//
// Note that you configure [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html) on a `AWS::Lambda::Version` or a `AWS::Lambda::Alias` .
//
// For a complete introduction to Lambda functions, see [What is Lambda?](https://docs.aws.amazon.com/lambda/latest/dg/lambda-welcome.html) in the *Lambda developer guide.*
//
// TODO: EXAMPLE
//
type CfnFunction interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Code() interface{}
	SetCode(val interface{})
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	CreationStack() *[]*string
	DeadLetterConfig() interface{}
	SetDeadLetterConfig(val interface{})
	Description() *string
	SetDescription(val *string)
	Environment() interface{}
	SetEnvironment(val interface{})
	FileSystemConfigs() interface{}
	SetFileSystemConfigs(val interface{})
	FunctionName() *string
	SetFunctionName(val *string)
	Handler() *string
	SetHandler(val *string)
	ImageConfig() interface{}
	SetImageConfig(val interface{})
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	Layers() *[]*string
	SetLayers(val *[]*string)
	LogicalId() *string
	MemorySize() *float64
	SetMemorySize(val *float64)
	Node() constructs.Node
	PackageType() *string
	SetPackageType(val *string)
	Ref() *string
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	Role() *string
	SetRole(val *string)
	Runtime() *string
	SetRuntime(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	Timeout() *float64
	SetTimeout(val *float64)
	TracingConfig() interface{}
	SetTracingConfig(val interface{})
	UpdatedProperites() *map[string]interface{}
	VpcConfig() interface{}
	SetVpcConfig(val interface{})
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

func (j *jsiiProxy_CfnFunction) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
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

func (j *jsiiProxy_CfnFunction) Code() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"code",
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

func (j *jsiiProxy_CfnFunction) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFunction) DeadLetterConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterConfig",
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

func (j *jsiiProxy_CfnFunction) TracingConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingConfig",
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

func (j *jsiiProxy_CfnFunction) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::Function`.
func NewCfnFunction(scope constructs.Construct, id *string, props *CfnFunctionProps) CfnFunction {
	_init_.Initialize()

	j := jsiiProxy_CfnFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::Function`.
func NewCfnFunction_Override(c CfnFunction, scope constructs.Construct, id *string, props *CfnFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnFunction",
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

func (j *jsiiProxy_CfnFunction) SetCode(val interface{}) {
	_jsii_.Set(
		j,
		"code",
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

func (j *jsiiProxy_CfnFunction) SetDeadLetterConfig(val interface{}) {
	_jsii_.Set(
		j,
		"deadLetterConfig",
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

func (j *jsiiProxy_CfnFunction) SetTracingConfig(val interface{}) {
	_jsii_.Set(
		j,
		"tracingConfig",
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
		"aws-cdk-lib.aws_lambda.CfnFunction",
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
		"aws-cdk-lib.aws_lambda.CfnFunction",
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
func CfnFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnFunction",
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
		"aws-cdk-lib.aws_lambda.CfnFunction",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnFunction) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnFunction) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnFunction) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnFunction) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnFunction) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnFunction) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnFunction) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

// The [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) for a Lambda function. To deploy a function defined as a container image, you specify the location of a container image in the Amazon ECR registry. For a .zip file deployment package, you can specify the location of an object in Amazon S3. For Node.js and Python functions, you can specify the function code inline in the template.
//
// Changes to a deployment package in Amazon S3 are not detected automatically during stack updates. To update the function code, change the object key or version in the template.
//
// TODO: EXAMPLE
//
type CfnFunction_CodeProperty struct {
	// URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon ECR registry.
	ImageUri *string `json:"imageUri"`
	// An Amazon S3 bucket in the same AWS Region as your function.
	//
	// The bucket can be in a different AWS account.
	S3Bucket *string `json:"s3Bucket"`
	// The Amazon S3 key of the deployment package.
	S3Key *string `json:"s3Key"`
	// For versioned objects, the version of the deployment package object to use.
	S3ObjectVersion *string `json:"s3ObjectVersion"`
	// (Node.js and Python) The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named `index` and zips it to create a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) . For the `Handler` property, the first part of the handler identifier must be `index` . For example, `index.handler` .
	//
	// Your source code can contain up to 4096 characters. For JSON, you must escape quotes and special characters such as newline ( `\n` ) with a backslash.
	//
	// If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module ( [cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html) ) that simplifies sending responses. See [Using AWS Lambda with AWS CloudFormation](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudformation.html) for details.
	ZipFile *string `json:"zipFile"`
}

// The [dead-letter queue](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq) for failed asynchronous invocations.
//
// TODO: EXAMPLE
//
type CfnFunction_DeadLetterConfigProperty struct {
	// The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
	TargetArn *string `json:"targetArn"`
}

// A function's environment variable settings.
//
// You can use environment variables to adjust your function's behavior without updating code. An environment variable is a pair of strings that are stored in a function's version-specific configuration.
//
// TODO: EXAMPLE
//
type CfnFunction_EnvironmentProperty struct {
	// Environment variable key-value pairs.
	//
	// For more information, see [Using Lambda environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html) .
	Variables interface{} `json:"variables"`
}

// Details about the connection between a Lambda function and an [Amazon EFS file system](https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html) .
//
// TODO: EXAMPLE
//
type CfnFunction_FileSystemConfigProperty struct {
	// The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.
	Arn *string `json:"arn"`
	// The path where the function can access the file system, starting with `/mnt/` .
	LocalMountPath *string `json:"localMountPath"`
}

// Configuration values that override the container image Dockerfile settings.
//
// See [Container settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) .
//
// TODO: EXAMPLE
//
type CfnFunction_ImageConfigProperty struct {
	// Specifies parameters that you want to pass in with ENTRYPOINT.
	Command *[]*string `json:"command"`
	// Specifies the entry point to their application, which is typically the location of the runtime executable.
	EntryPoint *[]*string `json:"entryPoint"`
	// Specifies the working directory.
	WorkingDirectory *string `json:"workingDirectory"`
}

// The function's [AWS X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) tracing configuration. To sample and record incoming requests, set `Mode` to `Active` .
//
// TODO: EXAMPLE
//
type CfnFunction_TracingConfigProperty struct {
	// The tracing mode.
	Mode *string `json:"mode"`
}

// The VPC security groups and subnets that are attached to a Lambda function.
//
// When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC. For more information, see [VPC Settings](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html) .
//
// > When you delete a function, AWS CloudFormation monitors the state of its network interfaces and waits for Lambda to delete them before proceeding. If the VPC is defined in the same stack, the network interfaces need to be deleted by Lambda before AWS CloudFormation can delete the VPC's resources.
// >
// > To monitor network interfaces, AWS CloudFormation needs the `ec2:DescribeNetworkInterfaces` permission. It obtains this from the user or role that modifies the stack. If you don't provide this permission, AWS CloudFormation does not wait for network interfaces to be deleted.
//
// TODO: EXAMPLE
//
type CfnFunction_VpcConfigProperty struct {
	// A list of VPC security groups IDs.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// A list of VPC subnet IDs.
	SubnetIds *[]*string `json:"subnetIds"`
}

// Properties for defining a `CfnFunction`.
//
// TODO: EXAMPLE
//
type CfnFunctionProps struct {
	// The code for the function.
	Code interface{} `json:"code"`
	// The Amazon Resource Name (ARN) of the function's execution role.
	Role *string `json:"role"`
	// The instruction set architecture that the function supports.
	//
	// Enter a string array with one of the valid values (arm64 or x86_64). The default value is `x86_64` .
	Architectures *[]*string `json:"architectures"`
	// To enable code signing for this function, specify the ARN of a code-signing configuration.
	//
	// A code-signing configuration
	// includes a set of signing profiles, which define the trusted publishers for this function.
	CodeSigningConfigArn *string `json:"codeSigningConfigArn"`
	// A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.
	//
	// For more information, see [Dead Letter Queues](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq) .
	DeadLetterConfig interface{} `json:"deadLetterConfig"`
	// A description of the function.
	Description *string `json:"description"`
	// Environment variables that are accessible from function code during execution.
	Environment interface{} `json:"environment"`
	// Connection settings for an Amazon EFS file system.
	//
	// To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) resource, you must also specify a `DependsOn` attribute to ensure that the mount target is created or updated before the function.
	//
	// For more information about using the `DependsOn` attribute, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	FileSystemConfigs interface{} `json:"fileSystemConfigs"`
	// The name of the Lambda function, up to 64 characters in length.
	//
	// If you don't specify a name, AWS CloudFormation generates one.
	//
	// If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	FunctionName *string `json:"functionName"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see [Programming Model](https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html) .
	Handler *string `json:"handler"`
	// Configuration values that override the container image Dockerfile settings.
	//
	// See [Container settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) .
	ImageConfig interface{} `json:"imageConfig"`
	// The ARN of the AWS Key Management Service ( AWS KMS ) key that's used to encrypt your function's environment variables.
	//
	// If it's not provided, AWS Lambda uses a default service key.
	KmsKeyArn *string `json:"kmsKeyArn"`
	// A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.
	Layers *[]*string `json:"layers"`
	// The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-memory.html) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.
	MemorySize *float64 `json:"memorySize"`
	// The type of deployment package.
	//
	// Set to `Image` for container image and set `Zip` for .zip file archive.
	PackageType *string `json:"packageType"`
	// The number of simultaneous executions to reserve for the function.
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions"`
	// The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) . Runtime is required if the deployment package is a .zip file archive.
	Runtime *string `json:"runtime"`
	// A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// The amount of time (in seconds) that Lambda allows a function to run before stopping it.
	//
	// The default is 3 seconds. The maximum allowed value is 900 seconds. For additional information, see [Lambda execution environment](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html) .
	Timeout *float64 `json:"timeout"`
	// Set `Mode` to `Active` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) .
	TracingConfig interface{} `json:"tracingConfig"`
	// For network connectivity to AWS resources in a [VPC](https://docs.aws.amazon.com/lambda/latest/dg/configuration-network.html) , specify a list of security groups and subnets in the VPC.
	VpcConfig interface{} `json:"vpcConfig"`
}

// A CloudFormation `AWS::Lambda::LayerVersion`.
//
// The `AWS::Lambda::LayerVersion` resource creates a [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) from a ZIP archive.
//
// TODO: EXAMPLE
//
type CfnLayerVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CompatibleArchitectures() *[]*string
	SetCompatibleArchitectures(val *[]*string)
	CompatibleRuntimes() *[]*string
	SetCompatibleRuntimes(val *[]*string)
	Content() interface{}
	SetContent(val interface{})
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	LayerName() *string
	SetLayerName(val *string)
	LicenseInfo() *string
	SetLicenseInfo(val *string)
	LogicalId() *string
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

func (j *jsiiProxy_CfnLayerVersion) CompatibleArchitectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"compatibleArchitectures",
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

func (j *jsiiProxy_CfnLayerVersion) Content() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"content",
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


// Create a new `AWS::Lambda::LayerVersion`.
func NewCfnLayerVersion(scope constructs.Construct, id *string, props *CfnLayerVersionProps) CfnLayerVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnLayerVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnLayerVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::LayerVersion`.
func NewCfnLayerVersion_Override(c CfnLayerVersion, scope constructs.Construct, id *string, props *CfnLayerVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnLayerVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetCompatibleArchitectures(val *[]*string) {
	_jsii_.Set(
		j,
		"compatibleArchitectures",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetCompatibleRuntimes(val *[]*string) {
	_jsii_.Set(
		j,
		"compatibleRuntimes",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersion) SetContent(val interface{}) {
	_jsii_.Set(
		j,
		"content",
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
		"aws-cdk-lib.aws_lambda.CfnLayerVersion",
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
		"aws-cdk-lib.aws_lambda.CfnLayerVersion",
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
func CfnLayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnLayerVersion",
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
		"aws-cdk-lib.aws_lambda.CfnLayerVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLayerVersion) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLayerVersion) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLayerVersion) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLayerVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLayerVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLayerVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
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

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnLayerVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
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

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
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

// A ZIP archive that contains the contents of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) .
//
// TODO: EXAMPLE
//
type CfnLayerVersion_ContentProperty struct {
	// The Amazon S3 bucket of the layer archive.
	S3Bucket *string `json:"s3Bucket"`
	// The Amazon S3 key of the layer archive.
	S3Key *string `json:"s3Key"`
	// For versioned objects, the version of the layer archive object to use.
	S3ObjectVersion *string `json:"s3ObjectVersion"`
}

// A CloudFormation `AWS::Lambda::LayerVersionPermission`.
//
// The `AWS::Lambda::LayerVersionPermission` resource adds permissions to the resource-based policy of a version of an [Lambda layer](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) . Use this action to grant layer usage permission to other accounts. You can grant permission to a single account, all AWS accounts, or all accounts in an organization.
//
// > Since the release of the [UpdateReplacePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html) both `UpdateReplacePolicy` and `DeletionPolicy` are required to protect your Resources/LayerPermissions from deletion.
//
// TODO: EXAMPLE
//
type CfnLayerVersionPermission interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Action() *string
	SetAction(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LayerVersionArn() *string
	SetLayerVersionArn(val *string)
	LogicalId() *string
	Node() constructs.Node
	OrganizationId() *string
	SetOrganizationId(val *string)
	Principal() *string
	SetPrincipal(val *string)
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

// The jsii proxy struct for CfnLayerVersionPermission
type jsiiProxy_CfnLayerVersionPermission struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnLayerVersionPermission) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) OrganizationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLayerVersionPermission) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::LayerVersionPermission`.
func NewCfnLayerVersionPermission(scope constructs.Construct, id *string, props *CfnLayerVersionPermissionProps) CfnLayerVersionPermission {
	_init_.Initialize()

	j := jsiiProxy_CfnLayerVersionPermission{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnLayerVersionPermission",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::LayerVersionPermission`.
func NewCfnLayerVersionPermission_Override(c CfnLayerVersionPermission, scope constructs.Construct, id *string, props *CfnLayerVersionPermissionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnLayerVersionPermission",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnLayerVersionPermission) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersionPermission) SetLayerVersionArn(val *string) {
	_jsii_.Set(
		j,
		"layerVersionArn",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersionPermission) SetOrganizationId(val *string) {
	_jsii_.Set(
		j,
		"organizationId",
		val,
	)
}

func (j *jsiiProxy_CfnLayerVersionPermission) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnLayerVersionPermission_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnLayerVersionPermission",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnLayerVersionPermission_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnLayerVersionPermission",
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
func CfnLayerVersionPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnLayerVersionPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLayerVersionPermission_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnLayerVersionPermission",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnLayerVersionPermission) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnLayerVersionPermission) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnLayerVersionPermission) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnLayerVersionPermission) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnLayerVersionPermission) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnLayerVersionPermission) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnLayerVersionPermission) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnLayerVersionPermission) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnLayerVersionPermission) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnLayerVersionPermission) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnLayerVersionPermission) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnLayerVersionPermission) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnLayerVersionPermission) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnLayerVersionPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnLayerVersionPermission) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnLayerVersionPermission`.
//
// TODO: EXAMPLE
//
type CfnLayerVersionPermissionProps struct {
	// The API action that grants access to the layer.
	//
	// For example, `lambda:GetLayerVersion` .
	Action *string `json:"action"`
	// The name or Amazon Resource Name (ARN) of the layer.
	LayerVersionArn *string `json:"layerVersionArn"`
	// An account ID, or `*` to grant layer usage permission to all accounts in an organization, or all AWS accounts (if `organizationId` is not specified).
	//
	// For the last case, make sure that you really do want all AWS accounts to have usage permission to this layer.
	Principal *string `json:"principal"`
	// With the principal set to `*` , grant permission to all accounts in the specified organization.
	OrganizationId *string `json:"organizationId"`
}

// Properties for defining a `CfnLayerVersion`.
//
// TODO: EXAMPLE
//
type CfnLayerVersionProps struct {
	// The function layer archive.
	Content interface{} `json:"content"`
	// A list of compatible [instruction set architectures](https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html) .
	CompatibleArchitectures *[]*string `json:"compatibleArchitectures"`
	// A list of compatible [function runtimes](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) . Used for filtering with [ListLayers](https://docs.aws.amazon.com/lambda/latest/dg/API_ListLayers.html) and [ListLayerVersions](https://docs.aws.amazon.com/lambda/latest/dg/API_ListLayerVersions.html) .
	CompatibleRuntimes *[]*string `json:"compatibleRuntimes"`
	// The description of the version.
	Description *string `json:"description"`
	// The name or Amazon Resource Name (ARN) of the layer.
	LayerName *string `json:"layerName"`
	// The layer's software license. It can be any of the following:.
	//
	// - An [SPDX license identifier](https://docs.aws.amazon.com/https://spdx.org/licenses/) . For example, `MIT` .
	// - The URL of a license hosted on the internet. For example, `https://opensource.org/licenses/MIT` .
	// - The full text of the license.
	LicenseInfo *string `json:"licenseInfo"`
}

// Lambda code defined using 2 CloudFormation parameters.
//
// Useful when you don't have access to the code of your Lambda from your CDK code, so you can't use Assets,
// and you want to deploy the Lambda in a CodePipeline, using CloudFormation Actions -
// you can fill the parameters using the {@link #assign} method.
//
// TODO: EXAMPLE
//
type CfnParametersCode interface {
	Code
	BucketNameParam() *string
	IsInline() *bool
	ObjectKeyParam() *string
	Assign(location *awss3.Location) *map[string]interface{}
	Bind(scope constructs.Construct) *CodeConfig
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for CfnParametersCode
type jsiiProxy_CfnParametersCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_CfnParametersCode) BucketNameParam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParametersCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnParametersCode) ObjectKeyParam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectKeyParam",
		&returns,
	)
	return returns
}


func NewCfnParametersCode(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	j := jsiiProxy_CfnParametersCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCfnParametersCode_Override(c CfnParametersCode, props *CfnParametersCodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		[]interface{}{props},
		c,
	)
}

// Loads the function code from a local disk path.
func CfnParametersCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func CfnParametersCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func CfnParametersCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`
func CfnParametersCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func CfnParametersCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func CfnParametersCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func CfnParametersCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnParametersCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Create a parameters map from this instance's CloudFormation parameters.
//
// It returns a map with 2 keys that correspond to the names of the parameters defined in this Lambda code,
// and as values it contains the appropriate expressions pointing at the provided S3 location
// (most likely, obtained from a CodePipeline Artifact by calling the `artifact.s3Location` method).
// The result should be provided to the CloudFormation Action
// that is deploying the Stack that the Lambda with this code is part of,
// in the `parameterOverrides` property.
func (c *jsiiProxy_CfnParametersCode) Assign(location *awss3.Location) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"assign",
		[]interface{}{location},
		&returns,
	)

	return returns
}

// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
func (c *jsiiProxy_CfnParametersCode) Bind(scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Called after the CFN function resource has been created to allow the code class to bind to it.
//
// Specifically it's required to allow assets to add
// metadata for tooling like SAM CLI to be able to find their origins.
func (c *jsiiProxy_CfnParametersCode) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		c,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

// Construction properties for {@link CfnParametersCode}.
//
// TODO: EXAMPLE
//
type CfnParametersCodeProps struct {
	// The CloudFormation parameter that represents the name of the S3 Bucket where the Lambda code will be located in.
	//
	// Must be of type 'String'.
	BucketNameParam awscdk.CfnParameter `json:"bucketNameParam"`
	// The CloudFormation parameter that represents the path inside the S3 Bucket where the Lambda code will be located at.
	//
	// Must be of type 'String'.
	ObjectKeyParam awscdk.CfnParameter `json:"objectKeyParam"`
}

// A CloudFormation `AWS::Lambda::Permission`.
//
// The `AWS::Lambda::Permission` resource grants an AWS service or another account permission to use a function. You can apply the policy at the function level, or specify a qualifier to restrict access to a single version or alias. If you use a qualifier, the invoker must use the full Amazon Resource Name (ARN) of that version or alias to invoke the function.
//
// To grant permission to another account, specify the account ID as the `Principal` . For AWS services, the principal is a domain-style identifier defined by the service, like `s3.amazonaws.com` or `sns.amazonaws.com` . For AWS services, you can also specify the ARN of the associated resource as the `SourceArn` . If you grant permission to a service principal without specifying the source, other accounts could potentially configure resources in their account to invoke your Lambda function.
//
// This resource adds a statement to a resource-based permission policy for the function. For more information about function policies, see [Lambda Function Policies](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) .
//
// TODO: EXAMPLE
//
type CfnPermission interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Action() *string
	SetAction(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	EventSourceToken() *string
	SetEventSourceToken(val *string)
	FunctionName() *string
	SetFunctionName(val *string)
	LogicalId() *string
	Node() constructs.Node
	Principal() *string
	SetPrincipal(val *string)
	Ref() *string
	SourceAccount() *string
	SetSourceAccount(val *string)
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

// The jsii proxy struct for CfnPermission
type jsiiProxy_CfnPermission struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnPermission) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) EventSourceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) SourceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermission) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::Permission`.
func NewCfnPermission(scope constructs.Construct, id *string, props *CfnPermissionProps) CfnPermission {
	_init_.Initialize()

	j := jsiiProxy_CfnPermission{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnPermission",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::Permission`.
func NewCfnPermission_Override(c CfnPermission, scope constructs.Construct, id *string, props *CfnPermissionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnPermission",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnPermission) SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetEventSourceToken(val *string) {
	_jsii_.Set(
		j,
		"eventSourceToken",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetPrincipal(val *string) {
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetSourceAccount(val *string) {
	_jsii_.Set(
		j,
		"sourceAccount",
		val,
	)
}

func (j *jsiiProxy_CfnPermission) SetSourceArn(val *string) {
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
func CfnPermission_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnPermission",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnPermission_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnPermission",
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
func CfnPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermission_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnPermission",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnPermission) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnPermission) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnPermission) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnPermission) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnPermission) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnPermission) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnPermission) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnPermission) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnPermission) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnPermission) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnPermission) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnPermission) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnPermission) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermission) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnPermission`.
//
// TODO: EXAMPLE
//
type CfnPermissionProps struct {
	// The action that the principal can use on the function.
	//
	// For example, `lambda:InvokeFunction` or `lambda:GetFunction` .
	Action *string `json:"action"`
	// The name of the Lambda function, version, or alias.
	//
	// **Name formats** - *Function name* - `my-function` (name-only), `my-function:v1` (with alias).
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:my-function` .
	// - *Partial ARN* - `123456789012:function:my-function` .
	//
	// You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName *string `json:"functionName"`
	// The AWS service or account that invokes the function.
	//
	// If you specify a service, use `SourceArn` or `SourceAccount` to limit who can invoke the function through that service.
	Principal *string `json:"principal"`
	// For Alexa Smart Home functions, a token that must be supplied by the invoker.
	EventSourceToken *string `json:"eventSourceToken"`
	// For Amazon S3, the ID of the account that owns the resource.
	//
	// Use this together with `SourceArn` to ensure that the resource is owned by the specified account. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.
	SourceAccount *string `json:"sourceAccount"`
	// For AWS services, the ARN of the AWS resource that invokes the function.
	//
	// For example, an Amazon S3 bucket or Amazon SNS topic.
	//
	// Note that Lambda configures the comparison using the `StringLike` operator.
	SourceArn *string `json:"sourceArn"`
}

// A CloudFormation `AWS::Lambda::Version`.
//
// The `AWS::Lambda::Version` resource creates a [version](https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) from the current code and configuration of a function. Use versions to create a snapshot of your function code and configuration that doesn't change.
//
// TODO: EXAMPLE
//
type CfnVersion interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrVersion() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CodeSha256() *string
	SetCodeSha256(val *string)
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	FunctionName() *string
	SetFunctionName(val *string)
	LogicalId() *string
	Node() constructs.Node
	ProvisionedConcurrencyConfig() interface{}
	SetProvisionedConcurrencyConfig(val interface{})
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

// The jsii proxy struct for CfnVersion
type jsiiProxy_CfnVersion struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnVersion) AttrVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) CodeSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) ProvisionedConcurrencyConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedConcurrencyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVersion) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Lambda::Version`.
func NewCfnVersion(scope constructs.Construct, id *string, props *CfnVersionProps) CfnVersion {
	_init_.Initialize()

	j := jsiiProxy_CfnVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Lambda::Version`.
func NewCfnVersion_Override(c CfnVersion, scope constructs.Construct, id *string, props *CfnVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CfnVersion",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnVersion) SetCodeSha256(val *string) {
	_jsii_.Set(
		j,
		"codeSha256",
		val,
	)
}

func (j *jsiiProxy_CfnVersion) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnVersion) SetFunctionName(val *string) {
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_CfnVersion) SetProvisionedConcurrencyConfig(val interface{}) {
	_jsii_.Set(
		j,
		"provisionedConcurrencyConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnVersion_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnVersion",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnVersion_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnVersion",
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
func CfnVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CfnVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVersion_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.CfnVersion",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnVersion) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnVersion) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnVersion) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnVersion) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnVersion) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnVersion) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnVersion) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnVersion) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnVersion) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnVersion) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnVersion) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnVersion) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnVersion) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A [provisioned concurrency](https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html) configuration for a function's version.
//
// TODO: EXAMPLE
//
type CfnVersion_ProvisionedConcurrencyConfigurationProperty struct {
	// The amount of provisioned concurrency to allocate for the version.
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions"`
}

// Properties for defining a `CfnVersion`.
//
// TODO: EXAMPLE
//
type CfnVersionProps struct {
	// The name of the Lambda function.
	//
	// **Name formats** - *Function name* - `MyFunction` .
	// - *Function ARN* - `arn:aws:lambda:us-west-2:123456789012:function:MyFunction` .
	// - *Partial ARN* - `123456789012:function:MyFunction` .
	//
	// The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.
	FunctionName *string `json:"functionName"`
	// Only publish a version if the hash value matches the value that's specified.
	//
	// Use this option to avoid publishing a version if the function code has changed since you last updated it. Updates are not supported for this property.
	CodeSha256 *string `json:"codeSha256"`
	// A description for the version to override the description in the function configuration.
	//
	// Updates are not supported for this property.
	Description *string `json:"description"`
	// Specifies a provisioned concurrency configuration for a function's version.
	//
	// Updates are not supported for this property.
	ProvisionedConcurrencyConfig interface{} `json:"provisionedConcurrencyConfig"`
}

// Represents the Lambda Handler Code.
//
// TODO: EXAMPLE
//
type Code interface {
	Bind(scope constructs.Construct) *CodeConfig
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Code",
		nil, // no parameters
		c,
	)
}

// Loads the function code from a local disk path.
func Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func Code_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func Code_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`
func Code_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func Code_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func Code_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
func (c *jsiiProxy_Code) Bind(scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Called after the CFN function resource has been created to allow the code class to bind to it.
//
// Specifically it's required to allow assets to add
// metadata for tooling like SAM CLI to be able to find their origins.
func (c *jsiiProxy_Code) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		c,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

// Result of binding `Code` into a `Function`.
//
// TODO: EXAMPLE
//
type CodeConfig struct {
	// Docker image configuration (mutually exclusive with `s3Location` and `inlineCode`).
	Image *CodeImageConfig `json:"image"`
	// Inline code (mutually exclusive with `s3Location` and `image`).
	InlineCode *string `json:"inlineCode"`
	// The location of the code in S3 (mutually exclusive with `inlineCode` and `image`).
	S3Location *awss3.Location `json:"s3Location"`
}

// Result of the bind when an ECR image is used.
//
// TODO: EXAMPLE
//
type CodeImageConfig struct {
	// URI to the Docker image.
	ImageUri *string `json:"imageUri"`
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	Cmd *[]*string `json:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	Entrypoint *[]*string `json:"entrypoint"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	WorkingDirectory *string `json:"workingDirectory"`
}

// Defines a Code Signing Config.
//
// TODO: EXAMPLE
//
type CodeSigningConfig interface {
	awscdk.Resource
	ICodeSigningConfig
	CodeSigningConfigArn() *string
	CodeSigningConfigId() *string
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

// The jsii proxy struct for CodeSigningConfig
type jsiiProxy_CodeSigningConfig struct {
	internal.Type__awscdkResource
	jsiiProxy_ICodeSigningConfig
}

func (j *jsiiProxy_CodeSigningConfig) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeSigningConfig) CodeSigningConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeSigningConfig) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeSigningConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeSigningConfig) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeSigningConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCodeSigningConfig(scope constructs.Construct, id *string, props *CodeSigningConfigProps) CodeSigningConfig {
	_init_.Initialize()

	j := jsiiProxy_CodeSigningConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CodeSigningConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCodeSigningConfig_Override(c CodeSigningConfig, scope constructs.Construct, id *string, props *CodeSigningConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.CodeSigningConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

// Creates a Signing Profile construct that represents an external Signing Profile.
func CodeSigningConfig_FromCodeSigningConfigArn(scope constructs.Construct, id *string, codeSigningConfigArn *string) ICodeSigningConfig {
	_init_.Initialize()

	var returns ICodeSigningConfig

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CodeSigningConfig",
		"fromCodeSigningConfigArn",
		[]interface{}{scope, id, codeSigningConfigArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CodeSigningConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CodeSigningConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CodeSigningConfig_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.CodeSigningConfig",
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
func (c *jsiiProxy_CodeSigningConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CodeSigningConfig) GeneratePhysicalName() *string {
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
func (c *jsiiProxy_CodeSigningConfig) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (c *jsiiProxy_CodeSigningConfig) GetResourceNameAttribute(nameAttr *string) *string {
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
func (c *jsiiProxy_CodeSigningConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a Code Signing Config object.
//
// TODO: EXAMPLE
//
type CodeSigningConfigProps struct {
	// List of signing profiles that defines a trusted user who can sign a code package.
	SigningProfiles *[]awssigner.ISigningProfile `json:"signingProfiles"`
	// Code signing configuration description.
	Description *string `json:"description"`
	// Code signing configuration policy for deployment validation failure.
	//
	// If you set the policy to Enforce, Lambda blocks the deployment request
	// if signature validation checks fail.
	// If you set the policy to Warn, Lambda allows the deployment and
	// creates a CloudWatch log.
	UntrustedArtifactOnDeployment UntrustedArtifactOnDeployment `json:"untrustedArtifactOnDeployment"`
}

// A destination configuration.
//
// TODO: EXAMPLE
//
type DestinationConfig struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `json:"destination"`
}

// Options when binding a destination to a function.
//
// TODO: EXAMPLE
//
type DestinationOptions struct {
	// The destination type.
	Type DestinationType `json:"type"`
}

// The type of destination.
type DestinationType string

const (
	DestinationType_FAILURE DestinationType = "FAILURE"
	DestinationType_SUCCESS DestinationType = "SUCCESS"
)

// A destination configuration.
//
// TODO: EXAMPLE
//
type DlqDestinationConfig struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string `json:"destination"`
}

// Options when creating an asset from a Docker build.
//
// TODO: EXAMPLE
//
type DockerBuildAssetOptions struct {
	// Build args.
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Name of the Dockerfile, must relative to the docker build path.
	File *string `json:"file"`
	// Set platform if server is multi-platform capable. _Requires Docker Engine API v1.38+_.
	//
	// Example value: `linux/amd64`
	Platform *string `json:"platform"`
	// The path in the Docker image where the asset is located after the build operation.
	ImagePath *string `json:"imagePath"`
	// The path on the local filesystem where the asset will be copied using `docker cp`.
	OutputPath *string `json:"outputPath"`
}

// Code property for the DockerImageFunction construct.
//
// TODO: EXAMPLE
//
type DockerImageCode interface {
}

// The jsii proxy struct for DockerImageCode
type jsiiProxy_DockerImageCode struct {
	_ byte // padding
}

func NewDockerImageCode_Override(d DockerImageCode) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.DockerImageCode",
		nil, // no parameters
		d,
	)
}

// Use an existing ECR image as the Lambda code.
func DockerImageCode_FromEcr(repository awsecr.IRepository, props *EcrImageCodeProps) DockerImageCode {
	_init_.Initialize()

	var returns DockerImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageCode",
		"fromEcr",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func DockerImageCode_FromImageAsset(directory *string, props *AssetImageCodeProps) DockerImageCode {
	_init_.Initialize()

	var returns DockerImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageCode",
		"fromImageAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Create a lambda function where the handler is a docker image.
//
// TODO: EXAMPLE
//
type DockerImageFunction interface {
	Function
	Architecture() Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	CurrentVersion() Version
	DeadLetterQueue() awssqs.IQueue
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	LatestVersion() IVersion
	LogGroup() awslogs.ILogGroup
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Runtime() Runtime
	Stack() awscdk.Stack
	Timeout() awscdk.Duration
	AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function
	AddEventSource(source IEventSource)
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	AddLayers(layers ...ILayerVersion)
	AddPermission(id *string, permission *Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for DockerImageFunction
type jsiiProxy_DockerImageFunction struct {
	jsiiProxy_Function
}

func (j *jsiiProxy_DockerImageFunction) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) CurrentVersion() Version {
	var returns Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Runtime() Runtime {
	var returns Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageFunction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewDockerImageFunction(scope constructs.Construct, id *string, props *DockerImageFunctionProps) DockerImageFunction {
	_init_.Initialize()

	j := jsiiProxy_DockerImageFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDockerImageFunction_Override(d DockerImageFunction, scope constructs.Construct, id *string, props *DockerImageFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		[]interface{}{scope, id, props},
		d,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
func DockerImageFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
func DockerImageFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) IFunction {
	_init_.Initialize()

	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
func DockerImageFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *FunctionAttributes) IFunction {
	_init_.Initialize()

	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DockerImageFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DockerImageFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
func DockerImageFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
func DockerImageFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
func DockerImageFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
func DockerImageFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
func DockerImageFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
func DockerImageFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
func DockerImageFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.DockerImageFunction",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Adds an environment variable to this Lambda function.
//
// If this is a ref to a Lambda function, this operation results in a no-op.
func (d *jsiiProxy_DockerImageFunction) AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function {
	var returns Function

	_jsii_.Invoke(
		d,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
//
// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
//
// The following example adds an SQS Queue as an event source:
// ```
// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
// myFunction.addEventSource(new SqsEventSource(myQueue));
// ```
func (d *jsiiProxy_DockerImageFunction) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		d,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (d *jsiiProxy_DockerImageFunction) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		d,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds one or more Lambda Layers to this Lambda function.
func (d *jsiiProxy_DockerImageFunction) AddLayers(layers ...ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"addLayers",
		args,
	)
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
func (d *jsiiProxy_DockerImageFunction) AddPermission(id *string, permission *Permission) {
	_jsii_.InvokeVoid(
		d,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (d *jsiiProxy_DockerImageFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		d,
		"addToRolePolicy",
		[]interface{}{statement},
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
func (d *jsiiProxy_DockerImageFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (d *jsiiProxy_DockerImageFunction) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		d,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (d *jsiiProxy_DockerImageFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DockerImageFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		d,
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
func (d *jsiiProxy_DockerImageFunction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (d *jsiiProxy_DockerImageFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
func (d *jsiiProxy_DockerImageFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
func (d *jsiiProxy_DockerImageFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (d *jsiiProxy_DockerImageFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
func (d *jsiiProxy_DockerImageFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
func (d *jsiiProxy_DockerImageFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DockerImageFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties to configure a new DockerImageFunction construct.
//
// TODO: EXAMPLE
//
type DockerImageFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	AllowAllOutbound *bool `json:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	AllowPublicSubnet *bool `json:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	Architecture Architecture `json:"architecture"`
	// Code signing config associated with this function.
	CodeSigningConfig ICodeSigningConfig `json:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	CurrentVersionOptions *VersionOptions `json:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `json:"deadLetterQueueEnabled"`
	// A description of the function.
	Description *string `json:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	Environment *map[string]*string `json:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	EnvironmentEncryption awskms.IKey `json:"environmentEncryption"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	Events *[]IEventSource `json:"events"`
	// The filesystem configuration for the lambda function.
	Filesystem FileSystem `json:"filesystem"`
	// A name for the function.
	FunctionName *string `json:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	InitialPolicy *[]awsiam.PolicyStatement `json:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	InsightsVersion LambdaInsightsVersion `json:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	Layers *[]ILayerVersion `json:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	LogRetentionRetryOptions *LogRetentionRetryOptions `json:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	LogRetentionRole awsiam.IRole `json:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	MemorySize *float64 `json:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	Profiling *bool `json:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `json:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions"`
	// Lambda execution role.
	//
	// This is the role that will be assumed by the function upon execution.
	// It controls the permissions that the function will have. The Role must
	// be assumable by the 'lambda.amazonaws.com' service principal.
	//
	// The default Role automatically has permissions granted for Lambda execution. If you
	// provide a Role, you must add the relevant AWS managed policies yourself.
	//
	// The relevant managed policies are "service-role/AWSLambdaBasicExecutionRole" and
	// "service-role/AWSLambdaVPCAccessExecutionRole".
	Role awsiam.IRole `json:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	Timeout awscdk.Duration `json:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	Tracing Tracing `json:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	Code DockerImageCode `json:"code"`
}

// Represents a Docker image in ECR that can be bound as Lambda Code.
//
// TODO: EXAMPLE
//
type EcrImageCode interface {
	Code
	IsInline() *bool
	Bind(_arg constructs.Construct) *CodeConfig
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for EcrImageCode
type jsiiProxy_EcrImageCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_EcrImageCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewEcrImageCode(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	j := jsiiProxy_EcrImageCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		[]interface{}{repository, props},
		&j,
	)

	return &j
}

func NewEcrImageCode_Override(e EcrImageCode, repository awsecr.IRepository, props *EcrImageCodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		[]interface{}{repository, props},
		e,
	)
}

// Loads the function code from a local disk path.
func EcrImageCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func EcrImageCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func EcrImageCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`
func EcrImageCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func EcrImageCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func EcrImageCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func EcrImageCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EcrImageCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
func (e *jsiiProxy_EcrImageCode) Bind(_arg constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// Called after the CFN function resource has been created to allow the code class to bind to it.
//
// Specifically it's required to allow assets to add
// metadata for tooling like SAM CLI to be able to find their origins.
func (e *jsiiProxy_EcrImageCode) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		e,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

// Properties to initialize a new EcrImageCode.
//
// TODO: EXAMPLE
//
type EcrImageCodeProps struct {
	// Specify or override the CMD on the specified Docker image or Dockerfile.
	//
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#cmd
	//
	Cmd *[]*string `json:"cmd"`
	// Specify or override the ENTRYPOINT on the specified Docker image or Dockerfile.
	//
	// An ENTRYPOINT allows you to configure a container that will run as an executable.
	// This needs to be in the 'exec form', viz., `[ 'executable', 'param1', 'param2' ]`.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	Entrypoint *[]*string `json:"entrypoint"`
	// The image tag to use when pulling the image from ECR.
	Tag *string `json:"tag"`
	// Specify or override the WORKDIR on the specified Docker image or Dockerfile.
	//
	// A WORKDIR allows you to configure the working directory the container will use.
	// See: https://docs.docker.com/engine/reference/builder/#workdir
	//
	WorkingDirectory *string `json:"workingDirectory"`
}

// Environment variables options.
//
// TODO: EXAMPLE
//
type EnvironmentOptions struct {
	// When used in Lambda@Edge via edgeArn() API, these environment variables will be removed.
	//
	// If not set, an error will be thrown.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-requirements-limits.html#lambda-requirements-lambda-function-configuration
	//
	RemoveInEdge *bool `json:"removeInEdge"`
}

// Configure options for asynchronous invocation on a version or an alias.
//
// By default, Lambda retries an asynchronous invocation twice if the function
// returns an error. It retains events in a queue for up to six hours. When an
// event fails all processing attempts or stays in the asynchronous invocation
// queue for too long, Lambda discards it.
//
// TODO: EXAMPLE
//
type EventInvokeConfig interface {
	awscdk.Resource
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

// The jsii proxy struct for EventInvokeConfig
type jsiiProxy_EventInvokeConfig struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_EventInvokeConfig) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventInvokeConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventInvokeConfig) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventInvokeConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewEventInvokeConfig(scope constructs.Construct, id *string, props *EventInvokeConfigProps) EventInvokeConfig {
	_init_.Initialize()

	j := jsiiProxy_EventInvokeConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.EventInvokeConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventInvokeConfig_Override(e EventInvokeConfig, scope constructs.Construct, id *string, props *EventInvokeConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.EventInvokeConfig",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventInvokeConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EventInvokeConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func EventInvokeConfig_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EventInvokeConfig",
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
func (e *jsiiProxy_EventInvokeConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_EventInvokeConfig) GeneratePhysicalName() *string {
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
func (e *jsiiProxy_EventInvokeConfig) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (e *jsiiProxy_EventInvokeConfig) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventInvokeConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to add an EventInvokeConfig to a function.
//
// TODO: EXAMPLE
//
type EventInvokeConfigOptions struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
}

// Properties for an EventInvokeConfig.
//
// TODO: EXAMPLE
//
type EventInvokeConfigProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// The Lambda function.
	Function IFunction `json:"function"`
	// The qualifier.
	Qualifier *string `json:"qualifier"`
}

// Defines a Lambda EventSourceMapping resource.
//
// Usually, you won't need to define the mapping yourself. This will usually be done by
// event sources. For example, to add an SQS event source to a function:
//
//     import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
//     lambda.addEventSource(new SqsEventSource(sqs));
//
// The `SqsEventSource` class will automatically create the mapping, and will also
// modify the Lambda's execution role so it can consume messages from the queue.
//
// TODO: EXAMPLE
//
type EventSourceMapping interface {
	awscdk.Resource
	IEventSourceMapping
	Env() *awscdk.ResourceEnvironment
	EventSourceMappingId() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for EventSourceMapping
type jsiiProxy_EventSourceMapping struct {
	internal.Type__awscdkResource
	jsiiProxy_IEventSourceMapping
}

func (j *jsiiProxy_EventSourceMapping) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventSourceMapping) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventSourceMapping) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventSourceMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewEventSourceMapping(scope constructs.Construct, id *string, props *EventSourceMappingProps) EventSourceMapping {
	_init_.Initialize()

	j := jsiiProxy_EventSourceMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.EventSourceMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEventSourceMapping_Override(e EventSourceMapping, scope constructs.Construct, id *string, props *EventSourceMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.EventSourceMapping",
		[]interface{}{scope, id, props},
		e,
	)
}

// Import an event source into this stack from its event source id.
func EventSourceMapping_FromEventSourceMappingId(scope constructs.Construct, id *string, eventSourceMappingId *string) IEventSourceMapping {
	_init_.Initialize()

	var returns IEventSourceMapping

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EventSourceMapping",
		"fromEventSourceMappingId",
		[]interface{}{scope, id, eventSourceMappingId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EventSourceMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func EventSourceMapping_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.EventSourceMapping",
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
func (e *jsiiProxy_EventSourceMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_EventSourceMapping) GeneratePhysicalName() *string {
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
func (e *jsiiProxy_EventSourceMapping) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (e *jsiiProxy_EventSourceMapping) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EventSourceMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type EventSourceMappingOptions struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range: Minimum value of 1. Maximum value of 10000.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// Set to false to disable the event source upon creation.
	Enabled *bool `json:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// Any record added to
	// this stream can invoke the Lambda function.
	EventSourceArn *string `json:"eventSourceArn"`
	// A list of host and port pairs that are the addresses of the Kafka brokers in a self managed "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.example.com:9096`.
	KafkaBootstrapServers *[]*string `json:"kafkaBootstrapServers"`
	// The name of the Kafka topic.
	KafkaTopic *string `json:"kafkaTopic"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	OnFailure IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Set to `undefined` if you want lambda to keep retrying infinitely or until
	// the record expires.
	//
	// Valid Range:
	// * Minimum value of 0
	// * Maximum value of 10000
	RetryAttempts *float64 `json:"retryAttempts"`
	// Specific settings like the authentication protocol or the VPC components to secure access to your event source.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html
	//
	SourceAccessConfigurations *[]*SourceAccessConfiguration `json:"sourceAccessConfigurations"`
	// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
	// See: https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType
	//
	StartingPosition StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-windows
	//
	// Valid Range: 0 - 15 minutes
	//
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
}

// Properties for declaring a new event source mapping.
//
// TODO: EXAMPLE
//
type EventSourceMappingProps struct {
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range: Minimum value of 1. Maximum value of 10000.
	BatchSize *float64 `json:"batchSize"`
	// If the function returns an error, split the batch in two and retry.
	BisectBatchOnError *bool `json:"bisectBatchOnError"`
	// Set to false to disable the event source upon creation.
	Enabled *bool `json:"enabled"`
	// The Amazon Resource Name (ARN) of the event source.
	//
	// Any record added to
	// this stream can invoke the Lambda function.
	EventSourceArn *string `json:"eventSourceArn"`
	// A list of host and port pairs that are the addresses of the Kafka brokers in a self managed "bootstrap" Kafka cluster that a Kafka client connects to initially to bootstrap itself.
	//
	// They are in the format `abc.example.com:9096`.
	KafkaBootstrapServers *[]*string `json:"kafkaBootstrapServers"`
	// The name of the Kafka topic.
	KafkaTopic *string `json:"kafkaTopic"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `json:"maxBatchingWindow"`
	// The maximum age of a record that Lambda sends to a function for processing.
	//
	// Valid Range:
	// * Minimum value of 60 seconds
	// * Maximum value of 7 days
	MaxRecordAge awscdk.Duration `json:"maxRecordAge"`
	// An Amazon SQS queue or Amazon SNS topic destination for discarded records.
	OnFailure IEventSourceDlq `json:"onFailure"`
	// The number of batches to process from each shard concurrently.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of 10
	ParallelizationFactor *float64 `json:"parallelizationFactor"`
	// Allow functions to return partially successful responses for a batch of records.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting
	//
	ReportBatchItemFailures *bool `json:"reportBatchItemFailures"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Set to `undefined` if you want lambda to keep retrying infinitely or until
	// the record expires.
	//
	// Valid Range:
	// * Minimum value of 0
	// * Maximum value of 10000
	RetryAttempts *float64 `json:"retryAttempts"`
	// Specific settings like the authentication protocol or the VPC components to secure access to your event source.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-sourceaccessconfiguration.html
	//
	SourceAccessConfigurations *[]*SourceAccessConfiguration `json:"sourceAccessConfigurations"`
	// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
	// See: https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType
	//
	StartingPosition StartingPosition `json:"startingPosition"`
	// The size of the tumbling windows to group records sent to DynamoDB or Kinesis.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-windows
	//
	// Valid Range: 0 - 15 minutes
	//
	TumblingWindow awscdk.Duration `json:"tumblingWindow"`
	// The target AWS Lambda function.
	Target IFunction `json:"target"`
}

// Represents the filesystem for the Lambda function.
//
// TODO: EXAMPLE
//
type FileSystem interface {
	Config() *FileSystemConfig
}

// The jsii proxy struct for FileSystem
type jsiiProxy_FileSystem struct {
	_ byte // padding
}

func (j *jsiiProxy_FileSystem) Config() *FileSystemConfig {
	var returns *FileSystemConfig
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewFileSystem(config *FileSystemConfig) FileSystem {
	_init_.Initialize()

	j := jsiiProxy_FileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FileSystem",
		[]interface{}{config},
		&j,
	)

	return &j
}

func NewFileSystem_Override(f FileSystem, config *FileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FileSystem",
		[]interface{}{config},
		f,
	)
}

// mount the filesystem from Amazon EFS.
func FileSystem_FromEfsAccessPoint(ap awsefs.IAccessPoint, mountPath *string) FileSystem {
	_init_.Initialize()

	var returns FileSystem

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FileSystem",
		"fromEfsAccessPoint",
		[]interface{}{ap, mountPath},
		&returns,
	)

	return returns
}

// FileSystem configurations for the Lambda function.
//
// TODO: EXAMPLE
//
type FileSystemConfig struct {
	// ARN of the access point.
	Arn *string `json:"arn"`
	// mount path in the lambda runtime environment.
	LocalMountPath *string `json:"localMountPath"`
	// connections object used to allow ingress traffic from lambda function.
	Connections awsec2.Connections `json:"connections"`
	// array of IDependable that lambda function depends on.
	Dependency *[]constructs.IDependable `json:"dependency"`
	// additional IAM policies required for the lambda function.
	Policies *[]awsiam.PolicyStatement `json:"policies"`
}

// Deploys a file from inside the construct library as a function.
//
// The supplied file is subject to the 4096 bytes limit of being embedded in a
// CloudFormation template.
//
// The construct includes an associated role with the lambda.
//
// This construct does not yet reproduce all features from the underlying resource
// library.
//
// TODO: EXAMPLE
//
type Function interface {
	FunctionBase
	Architecture() Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	CurrentVersion() Version
	DeadLetterQueue() awssqs.IQueue
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	LatestVersion() IVersion
	LogGroup() awslogs.ILogGroup
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Runtime() Runtime
	Stack() awscdk.Stack
	Timeout() awscdk.Duration
	AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function
	AddEventSource(source IEventSource)
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	AddLayers(layers ...ILayerVersion)
	AddPermission(id *string, permission *Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for Function
type jsiiProxy_Function struct {
	jsiiProxy_FunctionBase
}

func (j *jsiiProxy_Function) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) CurrentVersion() Version {
	var returns Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Runtime() Runtime {
	var returns Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewFunction(scope constructs.Construct, id *string, props *FunctionProps) Function {
	_init_.Initialize()

	j := jsiiProxy_Function{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Function",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewFunction_Override(f Function, scope constructs.Construct, id *string, props *FunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Function",
		[]interface{}{scope, id, props},
		f,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
func Function_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_lambda.Function",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
func Function_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) IFunction {
	_init_.Initialize()

	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
func Function_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *FunctionAttributes) IFunction {
	_init_.Initialize()

	var returns IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Function_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Function_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
func Function_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
func Function_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
func Function_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
func Function_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
func Function_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
func Function_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
func Function_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Function",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Adds an environment variable to this Lambda function.
//
// If this is a ref to a Lambda function, this operation results in a no-op.
func (f *jsiiProxy_Function) AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function {
	var returns Function

	_jsii_.Invoke(
		f,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
//
// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
//
// The following example adds an SQS Queue as an event source:
// ```
// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
// myFunction.addEventSource(new SqsEventSource(myQueue));
// ```
func (f *jsiiProxy_Function) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		f,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (f *jsiiProxy_Function) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		f,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds one or more Lambda Layers to this Lambda function.
func (f *jsiiProxy_Function) AddLayers(layers ...ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		f,
		"addLayers",
		args,
	)
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
func (f *jsiiProxy_Function) AddPermission(id *string, permission *Permission) {
	_jsii_.InvokeVoid(
		f,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (f *jsiiProxy_Function) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToRolePolicy",
		[]interface{}{statement},
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
func (f *jsiiProxy_Function) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (f *jsiiProxy_Function) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		f,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (f *jsiiProxy_Function) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
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
func (f *jsiiProxy_Function) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
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
func (f *jsiiProxy_Function) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (f *jsiiProxy_Function) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
func (f *jsiiProxy_Function) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
func (f *jsiiProxy_Function) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (f *jsiiProxy_Function) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
func (f *jsiiProxy_Function) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
func (f *jsiiProxy_Function) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_Function) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Represents a Lambda function defined outside of this stack.
//
// TODO: EXAMPLE
//
type FunctionAttributes struct {
	// The ARN of the Lambda function.
	//
	// Format: arn:<partition>:lambda:<region>:<account-id>:function:<function-name>
	FunctionArn *string `json:"functionArn"`
	// The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64).
	Architecture Architecture `json:"architecture"`
	// The IAM execution role associated with this function.
	//
	// If the role is not specified, any role-related operations will no-op.
	Role awsiam.IRole `json:"role"`
	// Setting this property informs the CDK that the imported function is in the same environment as the stack.
	//
	// This affects certain behaviours such as, whether this function's permission can be modified.
	// When not configured, the CDK attempts to auto-determine this. For environment agnostic stacks, i.e., stacks
	// where the account is not specified with the `env` property, this is determined to be false.
	//
	// Set this to property *ONLY IF* the imported function is in the same account as the stack
	// it's imported in.
	SameEnvironment *bool `json:"sameEnvironment"`
	// The security group of this Lambda, if in a VPC.
	//
	// This needs to be given in order to support allowing connections
	// to this Lambda.
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
}

type FunctionBase interface {
	awscdk.Resource
	awsec2.IClientVpnConnectionHandler
	IFunction
	Architecture() Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	LatestVersion() IVersion
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddEventSource(source IEventSource)
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	AddPermission(id *string, permission *Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for FunctionBase
type jsiiProxy_FunctionBase struct {
	internal.Type__awscdkResource
	internal.Type__awsec2IClientVpnConnectionHandler
	jsiiProxy_IFunction
}

func (j *jsiiProxy_FunctionBase) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewFunctionBase_Override(f FunctionBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.FunctionBase",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func FunctionBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FunctionBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func FunctionBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.FunctionBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
//
// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
//
// The following example adds an SQS Queue as an event source:
// ```
// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
// myFunction.addEventSource(new SqsEventSource(myQueue));
// ```
func (f *jsiiProxy_FunctionBase) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		f,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (f *jsiiProxy_FunctionBase) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		f,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
func (f *jsiiProxy_FunctionBase) AddPermission(id *string, permission *Permission) {
	_jsii_.InvokeVoid(
		f,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (f *jsiiProxy_FunctionBase) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		f,
		"addToRolePolicy",
		[]interface{}{statement},
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
func (f *jsiiProxy_FunctionBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (f *jsiiProxy_FunctionBase) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		f,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (f *jsiiProxy_FunctionBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
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
func (f *jsiiProxy_FunctionBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		f,
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
func (f *jsiiProxy_FunctionBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (f *jsiiProxy_FunctionBase) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
func (f *jsiiProxy_FunctionBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
func (f *jsiiProxy_FunctionBase) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (f *jsiiProxy_FunctionBase) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
func (f *jsiiProxy_FunctionBase) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
func (f *jsiiProxy_FunctionBase) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		f,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (f *jsiiProxy_FunctionBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Non runtime options.
//
// TODO: EXAMPLE
//
type FunctionOptions struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	AllowAllOutbound *bool `json:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	AllowPublicSubnet *bool `json:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	Architecture Architecture `json:"architecture"`
	// Code signing config associated with this function.
	CodeSigningConfig ICodeSigningConfig `json:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	CurrentVersionOptions *VersionOptions `json:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `json:"deadLetterQueueEnabled"`
	// A description of the function.
	Description *string `json:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	Environment *map[string]*string `json:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	EnvironmentEncryption awskms.IKey `json:"environmentEncryption"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	Events *[]IEventSource `json:"events"`
	// The filesystem configuration for the lambda function.
	Filesystem FileSystem `json:"filesystem"`
	// A name for the function.
	FunctionName *string `json:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	InitialPolicy *[]awsiam.PolicyStatement `json:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	InsightsVersion LambdaInsightsVersion `json:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	Layers *[]ILayerVersion `json:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	LogRetentionRetryOptions *LogRetentionRetryOptions `json:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	LogRetentionRole awsiam.IRole `json:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	MemorySize *float64 `json:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	Profiling *bool `json:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `json:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions"`
	// Lambda execution role.
	//
	// This is the role that will be assumed by the function upon execution.
	// It controls the permissions that the function will have. The Role must
	// be assumable by the 'lambda.amazonaws.com' service principal.
	//
	// The default Role automatically has permissions granted for Lambda execution. If you
	// provide a Role, you must add the relevant AWS managed policies yourself.
	//
	// The relevant managed policies are "service-role/AWSLambdaBasicExecutionRole" and
	// "service-role/AWSLambdaVPCAccessExecutionRole".
	Role awsiam.IRole `json:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	Timeout awscdk.Duration `json:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	Tracing Tracing `json:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
}

// TODO: EXAMPLE
//
type FunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	AllowAllOutbound *bool `json:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	AllowPublicSubnet *bool `json:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	Architecture Architecture `json:"architecture"`
	// Code signing config associated with this function.
	CodeSigningConfig ICodeSigningConfig `json:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	CurrentVersionOptions *VersionOptions `json:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `json:"deadLetterQueueEnabled"`
	// A description of the function.
	Description *string `json:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	Environment *map[string]*string `json:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	EnvironmentEncryption awskms.IKey `json:"environmentEncryption"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	Events *[]IEventSource `json:"events"`
	// The filesystem configuration for the lambda function.
	Filesystem FileSystem `json:"filesystem"`
	// A name for the function.
	FunctionName *string `json:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	InitialPolicy *[]awsiam.PolicyStatement `json:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	InsightsVersion LambdaInsightsVersion `json:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	Layers *[]ILayerVersion `json:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	LogRetentionRetryOptions *LogRetentionRetryOptions `json:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	LogRetentionRole awsiam.IRole `json:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	MemorySize *float64 `json:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	Profiling *bool `json:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `json:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions"`
	// Lambda execution role.
	//
	// This is the role that will be assumed by the function upon execution.
	// It controls the permissions that the function will have. The Role must
	// be assumable by the 'lambda.amazonaws.com' service principal.
	//
	// The default Role automatically has permissions granted for Lambda execution. If you
	// provide a Role, you must add the relevant AWS managed policies yourself.
	//
	// The relevant managed policies are "service-role/AWSLambdaBasicExecutionRole" and
	// "service-role/AWSLambdaVPCAccessExecutionRole".
	Role awsiam.IRole `json:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	Timeout awscdk.Duration `json:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	Tracing Tracing `json:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	Code Code `json:"code"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// The format includes the file name. It can also include
	// namespaces and other qualifiers, depending on the runtime.
	// For more information, see https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-features.html#gettingstarted-features-programmingmodel.
	//
	// Use `Handler.FROM_IMAGE` when defining a function from a Docker image.
	//
	// NOTE: If you specify your source code as inline text by specifying the
	// ZipFile property within the Code property, specify index.function_name as
	// the handler.
	Handler *string `json:"handler"`
	// The runtime environment for the Lambda function that you are uploading.
	//
	// For valid values, see the Runtime property in the AWS Lambda Developer
	// Guide.
	//
	// Use `Runtime.FROM_IMAGE` when when defining a function from a Docker image.
	Runtime Runtime `json:"runtime"`
}

// Lambda function handler.
type Handler interface {
}

// The jsii proxy struct for Handler
type jsiiProxy_Handler struct {
	_ byte // padding
}

func Handler_FROM_IMAGE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Handler",
		"FROM_IMAGE",
		&returns,
	)
	return returns
}

type IAlias interface {
	IFunction
	// Name of this alias.
	AliasName() *string
	// The underlying Lambda function version.
	Version() IVersion
}

// The jsii proxy for IAlias
type jsiiProxy_IAlias struct {
	jsiiProxy_IFunction
}

func (j *jsiiProxy_IAlias) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlias) Version() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// A Code Signing Config.
type ICodeSigningConfig interface {
	awscdk.IResource
	// The ARN of Code Signing Config.
	CodeSigningConfigArn() *string
	// The id of Code Signing Config.
	CodeSigningConfigId() *string
}

// The jsii proxy for ICodeSigningConfig
type jsiiProxy_ICodeSigningConfig struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ICodeSigningConfig) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICodeSigningConfig) CodeSigningConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigId",
		&returns,
	)
	return returns
}

// A Lambda destination.
type IDestination interface {
	// Binds this destination to the Lambda function.
	Bind(scope constructs.Construct, fn IFunction, options *DestinationOptions) *DestinationConfig
}

// The jsii proxy for IDestination
type jsiiProxy_IDestination struct {
	_ byte // padding
}

func (i *jsiiProxy_IDestination) Bind(scope constructs.Construct, fn IFunction, options *DestinationOptions) *DestinationConfig {
	var returns *DestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, fn, options},
		&returns,
	)

	return returns
}

// An abstract class which represents an AWS Lambda event source.
type IEventSource interface {
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target IFunction)
}

// The jsii proxy for IEventSource
type jsiiProxy_IEventSource struct {
	_ byte // padding
}

func (i *jsiiProxy_IEventSource) Bind(target IFunction) {
	_jsii_.InvokeVoid(
		i,
		"bind",
		[]interface{}{target},
	)
}

// A DLQ for an event source.
type IEventSourceDlq interface {
	// Returns the DLQ destination config of the DLQ.
	Bind(target IEventSourceMapping, targetHandler IFunction) *DlqDestinationConfig
}

// The jsii proxy for IEventSourceDlq
type jsiiProxy_IEventSourceDlq struct {
	_ byte // padding
}

func (i *jsiiProxy_IEventSourceDlq) Bind(target IEventSourceMapping, targetHandler IFunction) *DlqDestinationConfig {
	var returns *DlqDestinationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{target, targetHandler},
		&returns,
	)

	return returns
}

// Represents an event source mapping for a lambda function.
// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html
//
type IEventSourceMapping interface {
	awscdk.IResource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
}

// The jsii proxy for IEventSourceMapping
type jsiiProxy_IEventSourceMapping struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEventSourceMapping) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

type IFunction interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// Adds an event source to this function.
	//
	// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
	//
	// The following example adds an SQS Queue as an event source:
	// ```
	// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
	// myFunction.addEventSource(new SqsEventSource(myQueue));
	// ```
	AddEventSource(source IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	// Adds a permission to the Lambda resource policy.
	// See: Permission for details.
	//
	AddPermission(id *string, permission *Permission)
	// Adds a statement to the IAM role assumed by the instance.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Configures options for asynchronous invocation.
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	// Grant the given identity permissions to invoke this Lambda.
	GrantInvoke(identity awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this Lambda Return the given named metric for this Function.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the Duration of this Lambda How long execution of this Lambda takes.
	//
	// Average over 5 minutes
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How many invocations of this Lambda fail.
	//
	// Sum over 5 minutes
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of invocations of this Lambda How often this Lambda is invoked.
	//
	// Sum over 5 minutes
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of throttled invocations of this Lambda How often this Lambda is throttled.
	//
	// Sum over 5 minutes
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The system architectures compatible with this lambda function.
	Architecture() Architecture
	// The ARN of the function.
	FunctionArn() *string
	// The name of the function.
	FunctionName() *string
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	IsBoundToVpc() *bool
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	LatestVersion() IVersion
	// The construct node where permissions are attached.
	PermissionsNode() constructs.Node
	// The IAM role associated with this function.
	Role() awsiam.IRole
}

// The jsii proxy for IFunction
type jsiiProxy_IFunction struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFunction) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		i,
		"addEventSource",
		[]interface{}{source},
	)
}

func (i *jsiiProxy_IFunction) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		i,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) AddPermission(id *string, permission *Permission) {
	_jsii_.InvokeVoid(
		i,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (i *jsiiProxy_IFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_IFunction) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		i,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_IFunction) GrantInvoke(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IFunction) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

type ILayerVersion interface {
	awscdk.IResource
	// Add permission for this layer version to specific entities.
	//
	// Usage within
	// the same account where the layer is defined is always allowed and does not
	// require calling this method. Note that the principal that creates the
	// Lambda function using the layer (for example, a CloudFormation changeset
	// execution role) also needs to have the ``lambda:GetLayerVersion``
	// permission on the layer version.
	AddPermission(id *string, permission *LayerVersionPermission)
	// The runtimes compatible with this Layer.
	CompatibleRuntimes() *[]Runtime
	// The ARN of the Lambda Layer version that this Layer defines.
	LayerVersionArn() *string
}

// The jsii proxy for ILayerVersion
type jsiiProxy_ILayerVersion struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ILayerVersion) AddPermission(id *string, permission *LayerVersionPermission) {
	_jsii_.InvokeVoid(
		i,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (j *jsiiProxy_ILayerVersion) CompatibleRuntimes() *[]Runtime {
	var returns *[]Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

// Interface for scalable attributes.
type IScalableFunctionAttribute interface {
	constructs.IConstruct
	// Scale out or in based on schedule.
	ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in to keep utilization at a given level.
	//
	// The utilization is tracked by the
	// LambdaProvisionedConcurrencyUtilization metric, emitted by lambda. See:
	// https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics.html#monitoring-metrics-concurrency
	ScaleOnUtilization(options *UtilizationScalingOptions)
}

// The jsii proxy for IScalableFunctionAttribute
type jsiiProxy_IScalableFunctionAttribute struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IScalableFunctionAttribute) ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		i,
		"scaleOnSchedule",
		[]interface{}{id, actions},
	)
}

func (i *jsiiProxy_IScalableFunctionAttribute) ScaleOnUtilization(options *UtilizationScalingOptions) {
	_jsii_.InvokeVoid(
		i,
		"scaleOnUtilization",
		[]interface{}{options},
	)
}

type IVersion interface {
	IFunction
	// Defines an alias for this version.
	AddAlias(aliasName *string, options *AliasOptions) Alias
	// The ARN of the version for Lambda@Edge.
	EdgeArn() *string
	// The underlying AWS Lambda function.
	Lambda() IFunction
	// The most recently deployed version of this function.
	Version() *string
}

// The jsii proxy for IVersion
type jsiiProxy_IVersion struct {
	jsiiProxy_IFunction
}

func (i *jsiiProxy_IVersion) AddAlias(aliasName *string, options *AliasOptions) Alias {
	var returns Alias

	_jsii_.Invoke(
		i,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVersion) EdgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Lambda() IFunction {
	var returns IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

// Lambda code from an inline string (limited to 4KiB).
//
// TODO: EXAMPLE
//
type InlineCode interface {
	Code
	IsInline() *bool
	Bind(_scope constructs.Construct) *CodeConfig
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for InlineCode
type jsiiProxy_InlineCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_InlineCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewInlineCode(code *string) InlineCode {
	_init_.Initialize()

	j := jsiiProxy_InlineCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.InlineCode",
		[]interface{}{code},
		&j,
	)

	return &j
}

func NewInlineCode_Override(i InlineCode, code *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.InlineCode",
		[]interface{}{code},
		i,
	)
}

// Loads the function code from a local disk path.
func InlineCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func InlineCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func InlineCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`
func InlineCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func InlineCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func InlineCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func InlineCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InlineCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
func (i *jsiiProxy_InlineCode) Bind(_scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Called after the CFN function resource has been created to allow the code class to bind to it.
//
// Specifically it's required to allow assets to add
// metadata for tooling like SAM CLI to be able to find their origins.
func (i *jsiiProxy_InlineCode) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		i,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

// Version of CloudWatch Lambda Insights.
//
// TODO: EXAMPLE
//
type LambdaInsightsVersion interface {
	LayerVersionArn() *string
}

// The jsii proxy struct for LambdaInsightsVersion
type jsiiProxy_LambdaInsightsVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_LambdaInsightsVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}


func NewLambdaInsightsVersion_Override(l LambdaInsightsVersion) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		nil, // no parameters
		l,
	)
}

// Use the insights extension associated with the provided ARN.
//
// Make sure the ARN is associated
// with same region as your function
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-extension-versions.html
//
func LambdaInsightsVersion_FromInsightVersionArn(arn *string) LambdaInsightsVersion {
	_init_.Initialize()

	var returns LambdaInsightsVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"fromInsightVersionArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func LambdaInsightsVersion_VERSION_1_0_119_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_119_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_54_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_54_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_86_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_86_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_89_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_89_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_98_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_98_0",
		&returns,
	)
	return returns
}

// TODO: EXAMPLE
//
type LambdaRuntimeProps struct {
	// The Docker image name to be used for bundling in this runtime.
	BundlingDockerImage *string `json:"bundlingDockerImage"`
	// Whether this runtime is integrated with and supported for profiling using Amazon CodeGuru Profiler.
	SupportsCodeGuruProfiling *bool `json:"supportsCodeGuruProfiling"`
	// Whether the ``ZipFile`` (aka inline code) property can be used with this runtime.
	SupportsInlineCode *bool `json:"supportsInlineCode"`
}

// Defines a new Lambda Layer version.
//
// TODO: EXAMPLE
//
type LayerVersion interface {
	awscdk.Resource
	ILayerVersion
	CompatibleRuntimes() *[]Runtime
	Env() *awscdk.ResourceEnvironment
	LayerVersionArn() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	AddPermission(id *string, permission *LayerVersionPermission)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for LayerVersion
type jsiiProxy_LayerVersion struct {
	internal.Type__awscdkResource
	jsiiProxy_ILayerVersion
}

func (j *jsiiProxy_LayerVersion) CompatibleRuntimes() *[]Runtime {
	var returns *[]Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LayerVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewLayerVersion(scope constructs.Construct, id *string, props *LayerVersionProps) LayerVersion {
	_init_.Initialize()

	j := jsiiProxy_LayerVersion{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLayerVersion_Override(l LayerVersion, scope constructs.Construct, id *string, props *LayerVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		[]interface{}{scope, id, props},
		l,
	)
}

// Imports a layer version by ARN.
//
// Assumes it is compatible with all Lambda runtimes.
func LayerVersion_FromLayerVersionArn(scope constructs.Construct, id *string, layerVersionArn *string) ILayerVersion {
	_init_.Initialize()

	var returns ILayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"fromLayerVersionArn",
		[]interface{}{scope, id, layerVersionArn},
		&returns,
	)

	return returns
}

// Imports a Layer that has been defined externally.
func LayerVersion_FromLayerVersionAttributes(scope constructs.Construct, id *string, attrs *LayerVersionAttributes) ILayerVersion {
	_init_.Initialize()

	var returns ILayerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"fromLayerVersionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func LayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func LayerVersion_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LayerVersion",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add permission for this layer version to specific entities.
//
// Usage within
// the same account where the layer is defined is always allowed and does not
// require calling this method. Note that the principal that creates the
// Lambda function using the layer (for example, a CloudFormation changeset
// execution role) also needs to have the ``lambda:GetLayerVersion``
// permission on the layer version.
func (l *jsiiProxy_LayerVersion) AddPermission(id *string, permission *LayerVersionPermission) {
	_jsii_.InvokeVoid(
		l,
		"addPermission",
		[]interface{}{id, permission},
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
func (l *jsiiProxy_LayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		l,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (l *jsiiProxy_LayerVersion) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		l,
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
func (l *jsiiProxy_LayerVersion) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		l,
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
func (l *jsiiProxy_LayerVersion) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (l *jsiiProxy_LayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties necessary to import a LayerVersion.
//
// TODO: EXAMPLE
//
type LayerVersionAttributes struct {
	// The ARN of the LayerVersion.
	LayerVersionArn *string `json:"layerVersionArn"`
	// The list of compatible runtimes with this Layer.
	CompatibleRuntimes *[]Runtime `json:"compatibleRuntimes"`
}

// Non runtime options.
//
// TODO: EXAMPLE
//
type LayerVersionOptions struct {
	// The description the this Lambda Layer.
	Description *string `json:"description"`
	// The name of the layer.
	LayerVersionName *string `json:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	License *string `json:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
}

// Identification of an account (or organization) that is allowed to access a Lambda Layer Version.
//
// TODO: EXAMPLE
//
type LayerVersionPermission struct {
	// The AWS Account id of the account that is authorized to use a Lambda Layer Version.
	//
	// The wild-card ``'*'`` can be
	// used to grant access to "any" account (or any account in an organization when ``organizationId`` is specified).
	AccountId *string `json:"accountId"`
	// The ID of the AWS Organization to which the grant is restricted.
	//
	// Can only be specified if ``accountId`` is ``'*'``
	OrganizationId *string `json:"organizationId"`
}

// TODO: EXAMPLE
//
type LayerVersionProps struct {
	// The description the this Lambda Layer.
	Description *string `json:"description"`
	// The name of the layer.
	LayerVersionName *string `json:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	License *string `json:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// The content of this Layer.
	//
	// Using `Code.fromInline` is not supported.
	Code Code `json:"code"`
	// The system architectures compatible with this layer.
	CompatibleArchitectures *[]Architecture `json:"compatibleArchitectures"`
	// The runtimes compatible with this Layer.
	CompatibleRuntimes *[]Runtime `json:"compatibleRuntimes"`
}

// Retry options for all AWS API calls.
//
// TODO: EXAMPLE
//
type LogRetentionRetryOptions struct {
	// The base duration to use in the exponential backoff for operation retries.
	Base awscdk.Duration `json:"base"`
	// The maximum amount of retries.
	MaxRetries *float64 `json:"maxRetries"`
}

// Represents a permission statement that can be added to a Lambda function's resource policy via the `addPermissions()` method.
//
// TODO: EXAMPLE
//
type Permission struct {
	// The entity for which you are granting permission to invoke the Lambda function.
	//
	// This entity can be any valid AWS service principal, such as
	// s3.amazonaws.com or sns.amazonaws.com, or, if you are granting
	// cross-account permission, an AWS account ID. For example, you might want
	// to allow a custom application in another AWS account to push events to
	// Lambda by invoking your function.
	//
	// The principal can be either an AccountPrincipal or a ServicePrincipal.
	Principal awsiam.IPrincipal `json:"principal"`
	// The Lambda actions that you want to allow in this statement.
	//
	// For example,
	// you can specify lambda:CreateFunction to specify a certain action, or use
	// a wildcard (``lambda:*``) to grant permission to all Lambda actions. For a
	// list of actions, see Actions and Condition Context Keys for AWS Lambda in
	// the IAM User Guide.
	Action *string `json:"action"`
	// A unique token that must be supplied by the principal invoking the function.
	EventSourceToken *string `json:"eventSourceToken"`
	// The scope to which the permission constructs be attached.
	//
	// The default is
	// the Lambda function construct itself, but this would need to be different
	// in cases such as cross-stack references where the Permissions would need
	// to sit closer to the consumer of this permission (i.e., the caller).
	Scope constructs.Construct `json:"scope"`
	// The AWS account ID (without hyphens) of the source owner.
	//
	// For example, if
	// you specify an S3 bucket in the SourceArn property, this value is the
	// bucket owner's account ID. You can use this property to ensure that all
	// source principals are owned by a specific account.
	SourceAccount *string `json:"sourceAccount"`
	// The ARN of a resource that is invoking your function.
	//
	// When granting
	// Amazon Simple Storage Service (Amazon S3) permission to invoke your
	// function, specify this property with the bucket ARN as its value. This
	// ensures that events generated only from the specified bucket, not just
	// any bucket from any AWS account that creates a mapping to your function,
	// can invoke the function.
	SourceArn *string `json:"sourceArn"`
}

type QualifiedFunctionBase interface {
	FunctionBase
	Architecture() Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	Lambda() IFunction
	LatestVersion() IVersion
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Qualifier() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddEventSource(source IEventSource)
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	AddPermission(id *string, permission *Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for QualifiedFunctionBase
type jsiiProxy_QualifiedFunctionBase struct {
	jsiiProxy_FunctionBase
}

func (j *jsiiProxy_QualifiedFunctionBase) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Lambda() IFunction {
	var returns IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QualifiedFunctionBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewQualifiedFunctionBase_Override(q QualifiedFunctionBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.QualifiedFunctionBase",
		[]interface{}{scope, id, props},
		q,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func QualifiedFunctionBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.QualifiedFunctionBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func QualifiedFunctionBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.QualifiedFunctionBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
//
// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
//
// The following example adds an SQS Queue as an event source:
// ```
// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
// myFunction.addEventSource(new SqsEventSource(myQueue));
// ```
func (q *jsiiProxy_QualifiedFunctionBase) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		q,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (q *jsiiProxy_QualifiedFunctionBase) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		q,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
func (q *jsiiProxy_QualifiedFunctionBase) AddPermission(id *string, permission *Permission) {
	_jsii_.InvokeVoid(
		q,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (q *jsiiProxy_QualifiedFunctionBase) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		q,
		"addToRolePolicy",
		[]interface{}{statement},
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
func (q *jsiiProxy_QualifiedFunctionBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		q,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (q *jsiiProxy_QualifiedFunctionBase) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		q,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (q *jsiiProxy_QualifiedFunctionBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		q,
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
func (q *jsiiProxy_QualifiedFunctionBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		q,
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
func (q *jsiiProxy_QualifiedFunctionBase) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (q *jsiiProxy_QualifiedFunctionBase) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		q,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
func (q *jsiiProxy_QualifiedFunctionBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
func (q *jsiiProxy_QualifiedFunctionBase) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (q *jsiiProxy_QualifiedFunctionBase) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
func (q *jsiiProxy_QualifiedFunctionBase) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
func (q *jsiiProxy_QualifiedFunctionBase) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		q,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (q *jsiiProxy_QualifiedFunctionBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type ResourceBindOptions struct {
	// The name of the CloudFormation property to annotate with asset metadata.
	// See: https://github.com/aws/aws-cdk/issues/1432
	//
	ResourceProperty *string `json:"resourceProperty"`
}

// Lambda function runtime environment.
//
// If you need to use a runtime name that doesn't exist as a static member, you
// can instantiate a `Runtime` object, e.g: `new Runtime('nodejs99.99')`.
//
// TODO: EXAMPLE
//
type Runtime interface {
	BundlingImage() awscdk.DockerImage
	Family() RuntimeFamily
	Name() *string
	SupportsCodeGuruProfiling() *bool
	SupportsInlineCode() *bool
	RuntimeEquals(other Runtime) *bool
	ToString() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) BundlingImage() awscdk.DockerImage {
	var returns awscdk.DockerImage
	_jsii_.Get(
		j,
		"bundlingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) Family() RuntimeFamily {
	var returns RuntimeFamily
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) SupportsCodeGuruProfiling() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsCodeGuruProfiling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) SupportsInlineCode() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsInlineCode",
		&returns,
	)
	return returns
}


func NewRuntime(name *string, family RuntimeFamily, props *LambdaRuntimeProps) Runtime {
	_init_.Initialize()

	j := jsiiProxy_Runtime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Runtime",
		[]interface{}{name, family, props},
		&j,
	)

	return &j
}

func NewRuntime_Override(r Runtime, name *string, family RuntimeFamily, props *LambdaRuntimeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Runtime",
		[]interface{}{name, family, props},
		r,
	)
}

func Runtime_ALL() *[]Runtime {
	_init_.Initialize()
	var returns *[]Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"ALL",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_1",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_2",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_2_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_2_1",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_3_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_3_1",
		&returns,
	)
	return returns
}

func Runtime_FROM_IMAGE() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"FROM_IMAGE",
		&returns,
	)
	return returns
}

func Runtime_GO_1_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"GO_1_X",
		&returns,
	)
	return returns
}

func Runtime_JAVA_11() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_11",
		&returns,
	)
	return returns
}

func Runtime_JAVA_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_8",
		&returns,
	)
	return returns
}

func Runtime_JAVA_8_CORRETTO() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_8_CORRETTO",
		&returns,
	)
	return returns
}

func Runtime_NODEJS() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_10_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_10_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_12_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_12_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_14_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_14_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_4_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_4_3",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_6_10() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_6_10",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_8_10() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_8_10",
		&returns,
	)
	return returns
}

func Runtime_PROVIDED() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PROVIDED",
		&returns,
	)
	return returns
}

func Runtime_PROVIDED_AL2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PROVIDED_AL2",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_2_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_2_7",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_6",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_7",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_8",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_9() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_9",
		&returns,
	)
	return returns
}

func Runtime_RUBY_2_5() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"RUBY_2_5",
		&returns,
	)
	return returns
}

func Runtime_RUBY_2_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"RUBY_2_7",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Runtime) RuntimeEquals(other Runtime) *bool {
	var returns *bool

	_jsii_.Invoke(
		r,
		"runtimeEquals",
		[]interface{}{other},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Runtime) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type RuntimeFamily string

const (
	RuntimeFamily_NODEJS RuntimeFamily = "NODEJS"
	RuntimeFamily_JAVA RuntimeFamily = "JAVA"
	RuntimeFamily_PYTHON RuntimeFamily = "PYTHON"
	RuntimeFamily_DOTNET_CORE RuntimeFamily = "DOTNET_CORE"
	RuntimeFamily_GO RuntimeFamily = "GO"
	RuntimeFamily_RUBY RuntimeFamily = "RUBY"
	RuntimeFamily_OTHER RuntimeFamily = "OTHER"
)

// Lambda code from an S3 archive.
//
// TODO: EXAMPLE
//
type S3Code interface {
	Code
	IsInline() *bool
	Bind(_scope constructs.Construct) *CodeConfig
	BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions)
}

// The jsii proxy struct for S3Code
type jsiiProxy_S3Code struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_S3Code) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewS3Code(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	j := jsiiProxy_S3Code{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.S3Code",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

func NewS3Code_Override(s S3Code, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.S3Code",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Loads the function code from a local disk path.
func S3Code_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3Code",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func S3Code_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3Code",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func S3Code_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3Code",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`
func S3Code_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3Code",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func S3Code_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3Code",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func S3Code_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3Code",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func S3Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.S3Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
func (s *jsiiProxy_S3Code) Bind(_scope constructs.Construct) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// Called after the CFN function resource has been created to allow the code class to bind to it.
//
// Specifically it's required to allow assets to add
// metadata for tooling like SAM CLI to be able to find their origins.
func (s *jsiiProxy_S3Code) BindToResource(_resource awscdk.CfnResource, _options *ResourceBindOptions) {
	_jsii_.InvokeVoid(
		s,
		"bindToResource",
		[]interface{}{_resource, _options},
	)
}

// A Lambda that will only ever be added to a stack once.
//
// This construct is a way to guarantee that the lambda function will be guaranteed to be part of the stack,
// once and only once, irrespective of how many times the construct is declared to be part of the stack.
// This is guaranteed as long as the `uuid` property and the optional `lambdaPurpose` property stay the same
// whenever they're declared into the stack.
//
// TODO: EXAMPLE
//
type SingletonFunction interface {
	FunctionBase
	Architecture() Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	CurrentVersion() Version
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	LatestVersion() IVersion
	LogGroup() awslogs.ILogGroup
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Runtime() Runtime
	Stack() awscdk.Stack
	AddDependency(up ...constructs.IDependable)
	AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function
	AddEventSource(source IEventSource)
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	AddLayers(layers ...ILayerVersion)
	AddPermission(name *string, permission *Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	DependOn(down constructs.IConstruct)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for SingletonFunction
type jsiiProxy_SingletonFunction struct {
	jsiiProxy_FunctionBase
}

func (j *jsiiProxy_SingletonFunction) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) CurrentVersion() Version {
	var returns Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Runtime() Runtime {
	var returns Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingletonFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewSingletonFunction(scope constructs.Construct, id *string, props *SingletonFunctionProps) SingletonFunction {
	_init_.Initialize()

	j := jsiiProxy_SingletonFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSingletonFunction_Override(s SingletonFunction, scope constructs.Construct, id *string, props *SingletonFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func SingletonFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SingletonFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.SingletonFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Using node.addDependency() does not work on this method as the underlying lambda function is modeled as a singleton across the stack. Use this method instead to declare dependencies.
func (s *jsiiProxy_SingletonFunction) AddDependency(up ...constructs.IDependable) {
	args := []interface{}{}
	for _, a := range up {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addDependency",
		args,
	)
}

// Adds an environment variable to this Lambda function.
//
// If this is a ref to a Lambda function, this operation results in a no-op.
func (s *jsiiProxy_SingletonFunction) AddEnvironment(key *string, value *string, options *EnvironmentOptions) Function {
	var returns Function

	_jsii_.Invoke(
		s,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
//
// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
//
// The following example adds an SQS Queue as an event source:
// ```
// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
// myFunction.addEventSource(new SqsEventSource(myQueue));
// ```
func (s *jsiiProxy_SingletonFunction) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		s,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (s *jsiiProxy_SingletonFunction) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		s,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds one or more Lambda Layers to this Lambda function.
func (s *jsiiProxy_SingletonFunction) AddLayers(layers ...ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addLayers",
		args,
	)
}

// Adds a permission to the Lambda resource policy.
func (s *jsiiProxy_SingletonFunction) AddPermission(name *string, permission *Permission) {
	_jsii_.InvokeVoid(
		s,
		"addPermission",
		[]interface{}{name, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (s *jsiiProxy_SingletonFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
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
func (s *jsiiProxy_SingletonFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (s *jsiiProxy_SingletonFunction) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		s,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

// The SingletonFunction construct cannot be added as a dependency of another construct using node.addDependency(). Use this method instead to declare this as a dependency of another construct.
func (s *jsiiProxy_SingletonFunction) DependOn(down constructs.IConstruct) {
	_jsii_.InvokeVoid(
		s,
		"dependOn",
		[]interface{}{down},
	)
}

func (s *jsiiProxy_SingletonFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_SingletonFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
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
func (s *jsiiProxy_SingletonFunction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (s *jsiiProxy_SingletonFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
func (s *jsiiProxy_SingletonFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
func (s *jsiiProxy_SingletonFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (s *jsiiProxy_SingletonFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
func (s *jsiiProxy_SingletonFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
func (s *jsiiProxy_SingletonFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (s *jsiiProxy_SingletonFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a newly created singleton Lambda.
//
// TODO: EXAMPLE
//
type SingletonFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	AllowAllOutbound *bool `json:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	AllowPublicSubnet *bool `json:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	Architecture Architecture `json:"architecture"`
	// Code signing config associated with this function.
	CodeSigningConfig ICodeSigningConfig `json:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	CurrentVersionOptions *VersionOptions `json:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `json:"deadLetterQueueEnabled"`
	// A description of the function.
	Description *string `json:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	Environment *map[string]*string `json:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	EnvironmentEncryption awskms.IKey `json:"environmentEncryption"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	Events *[]IEventSource `json:"events"`
	// The filesystem configuration for the lambda function.
	Filesystem FileSystem `json:"filesystem"`
	// A name for the function.
	FunctionName *string `json:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	InitialPolicy *[]awsiam.PolicyStatement `json:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	InsightsVersion LambdaInsightsVersion `json:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	Layers *[]ILayerVersion `json:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	LogRetentionRetryOptions *LogRetentionRetryOptions `json:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	LogRetentionRole awsiam.IRole `json:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	MemorySize *float64 `json:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	Profiling *bool `json:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `json:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions"`
	// Lambda execution role.
	//
	// This is the role that will be assumed by the function upon execution.
	// It controls the permissions that the function will have. The Role must
	// be assumable by the 'lambda.amazonaws.com' service principal.
	//
	// The default Role automatically has permissions granted for Lambda execution. If you
	// provide a Role, you must add the relevant AWS managed policies yourself.
	//
	// The relevant managed policies are "service-role/AWSLambdaBasicExecutionRole" and
	// "service-role/AWSLambdaVPCAccessExecutionRole".
	Role awsiam.IRole `json:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	Timeout awscdk.Duration `json:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	Tracing Tracing `json:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	Code Code `json:"code"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// The format includes the file name. It can also include
	// namespaces and other qualifiers, depending on the runtime.
	// For more information, see https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-features.html#gettingstarted-features-programmingmodel.
	//
	// Use `Handler.FROM_IMAGE` when defining a function from a Docker image.
	//
	// NOTE: If you specify your source code as inline text by specifying the
	// ZipFile property within the Code property, specify index.function_name as
	// the handler.
	Handler *string `json:"handler"`
	// The runtime environment for the Lambda function that you are uploading.
	//
	// For valid values, see the Runtime property in the AWS Lambda Developer
	// Guide.
	//
	// Use `Runtime.FROM_IMAGE` when when defining a function from a Docker image.
	Runtime Runtime `json:"runtime"`
	// A unique identifier to identify this lambda.
	//
	// The identifier should be unique across all custom resource providers.
	// We recommend generating a UUID per provider.
	Uuid *string `json:"uuid"`
	// A descriptive name for the purpose of this Lambda.
	//
	// If the Lambda does not have a physical name, this string will be
	// reflected its generated name. The combination of lambdaPurpose
	// and uuid must be unique.
	LambdaPurpose *string `json:"lambdaPurpose"`
}

// Specific settings like the authentication protocol or the VPC components to secure access to your event source.
//
// TODO: EXAMPLE
//
type SourceAccessConfiguration struct {
	// The type of authentication protocol or the VPC components for your event source.
	//
	// For example: "SASL_SCRAM_512_AUTH".
	Type SourceAccessConfigurationType `json:"type"`
	// The value for your chosen configuration in type.
	//
	// For example: "URI": "arn:aws:secretsmanager:us-east-1:01234567890:secret:MyBrokerSecretName".
	// The exact string depends on the type.
	// See: SourceAccessConfigurationType
	//
	Uri *string `json:"uri"`
}

// The type of authentication protocol or the VPC components for your event source's SourceAccessConfiguration.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/lambda/latest/dg/API_SourceAccessConfiguration.html#SSS-Type-SourceAccessConfiguration-Type
//
type SourceAccessConfigurationType interface {
	Type() *string
}

// The jsii proxy struct for SourceAccessConfigurationType
type jsiiProxy_SourceAccessConfigurationType struct {
	_ byte // padding
}

func (j *jsiiProxy_SourceAccessConfigurationType) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// A custom source access configuration property.
func SourceAccessConfigurationType_Of(name *string) SourceAccessConfigurationType {
	_init_.Initialize()

	var returns SourceAccessConfigurationType

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func SourceAccessConfigurationType_BASIC_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"BASIC_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_CLIENT_CERTIFICATE_TLS_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"CLIENT_CERTIFICATE_TLS_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_SASL_SCRAM_256_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"SASL_SCRAM_256_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_SASL_SCRAM_512_AUTH() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"SASL_SCRAM_512_AUTH",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_VPC_SECURITY_GROUP() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"VPC_SECURITY_GROUP",
		&returns,
	)
	return returns
}

func SourceAccessConfigurationType_VPC_SUBNET() SourceAccessConfigurationType {
	_init_.Initialize()
	var returns SourceAccessConfigurationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.SourceAccessConfigurationType",
		"VPC_SUBNET",
		&returns,
	)
	return returns
}

// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
//
// TODO: EXAMPLE
//
type StartingPosition string

const (
	StartingPosition_TRIM_HORIZON StartingPosition = "TRIM_HORIZON"
	StartingPosition_LATEST StartingPosition = "LATEST"
)

// X-Ray Tracing Modes (https://docs.aws.amazon.com/lambda/latest/dg/API_TracingConfig.html).
//
// TODO: EXAMPLE
//
type Tracing string

const (
	Tracing_ACTIVE Tracing = "ACTIVE"
	Tracing_PASS_THROUGH Tracing = "PASS_THROUGH"
	Tracing_DISABLED Tracing = "DISABLED"
)

// Code signing configuration policy for deployment validation failure.
type UntrustedArtifactOnDeployment string

const (
	UntrustedArtifactOnDeployment_ENFORCE UntrustedArtifactOnDeployment = "ENFORCE"
	UntrustedArtifactOnDeployment_WARN UntrustedArtifactOnDeployment = "WARN"
)

// Options for enabling Lambda utilization tracking.
//
// TODO: EXAMPLE
//
type UtilizationScalingOptions struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// A name for the scaling policy.
	PolicyName *string `json:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown"`
	// Utilization target for the attribute.
	//
	// For example, .5 indicates that 50 percent of allocated provisioned concurrency is in use.
	UtilizationTarget *float64 `json:"utilizationTarget"`
}

// A single newly-deployed version of a Lambda function.
//
// This object exists to--at deploy time--query the "then-current" version of
// the Lambda function that it refers to. This Version object can then be
// used in `Alias` to refer to a particular deployment of a Lambda.
//
// This means that for every new update you deploy to your Lambda (using the
// CDK and Aliases), you must always create a new Version object. In
// particular, it must have a different name, so that a new resource is
// created.
//
// If you want to ensure that you're associating the right version with
// the right deployment, specify the `codeSha256` property while
// creating the `Version.
//
// TODO: EXAMPLE
//
type Version interface {
	QualifiedFunctionBase
	IVersion
	Architecture() Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	EdgeArn() *string
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	Lambda() IFunction
	LatestVersion() IVersion
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Qualifier() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	Version() *string
	AddAlias(aliasName *string, options *AliasOptions) Alias
	AddEventSource(source IEventSource)
	AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping
	AddPermission(id *string, permission *Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *EventInvokeConfigOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for Version
type jsiiProxy_Version struct {
	jsiiProxy_QualifiedFunctionBase
	jsiiProxy_IVersion
}

func (j *jsiiProxy_Version) Architecture() Architecture {
	var returns Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) EdgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Lambda() IFunction {
	var returns IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) LatestVersion() IVersion {
	var returns IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Version) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewVersion(scope constructs.Construct, id *string, props *VersionProps) Version {
	_init_.Initialize()

	j := jsiiProxy_Version{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Version",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewVersion_Override(v Version, scope constructs.Construct, id *string, props *VersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Version",
		[]interface{}{scope, id, props},
		v,
	)
}

// Construct a Version object from a Version ARN.
func Version_FromVersionArn(scope constructs.Construct, id *string, versionArn *string) IVersion {
	_init_.Initialize()

	var returns IVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Version",
		"fromVersionArn",
		[]interface{}{scope, id, versionArn},
		&returns,
	)

	return returns
}

func Version_FromVersionAttributes(scope constructs.Construct, id *string, attrs *VersionAttributes) IVersion {
	_init_.Initialize()

	var returns IVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Version",
		"fromVersionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Version_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Version",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Version_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.Version",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Defines an alias for this version.
func (v *jsiiProxy_Version) AddAlias(aliasName *string, options *AliasOptions) Alias {
	var returns Alias

	_jsii_.Invoke(
		v,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
//
// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
//
// The following example adds an SQS Queue as an event source:
// ```
// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
// myFunction.addEventSource(new SqsEventSource(myQueue));
// ```
func (v *jsiiProxy_Version) AddEventSource(source IEventSource) {
	_jsii_.InvokeVoid(
		v,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (v *jsiiProxy_Version) AddEventSourceMapping(id *string, options *EventSourceMappingOptions) EventSourceMapping {
	var returns EventSourceMapping

	_jsii_.Invoke(
		v,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
func (v *jsiiProxy_Version) AddPermission(id *string, permission *Permission) {
	_jsii_.InvokeVoid(
		v,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (v *jsiiProxy_Version) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		v,
		"addToRolePolicy",
		[]interface{}{statement},
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
func (v *jsiiProxy_Version) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		v,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (v *jsiiProxy_Version) ConfigureAsyncInvoke(options *EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		v,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (v *jsiiProxy_Version) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		v,
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
func (v *jsiiProxy_Version) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		v,
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
func (v *jsiiProxy_Version) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (v *jsiiProxy_Version) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		v,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
func (v *jsiiProxy_Version) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
func (v *jsiiProxy_Version) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (v *jsiiProxy_Version) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
func (v *jsiiProxy_Version) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
func (v *jsiiProxy_Version) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		v,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (v *jsiiProxy_Version) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type VersionAttributes struct {
	// The lambda function.
	Lambda IFunction `json:"lambda"`
	// The version.
	Version *string `json:"version"`
}

// Options for `lambda.Version`.
//
// TODO: EXAMPLE
//
type VersionOptions struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// SHA256 of the version of the Lambda source code.
	//
	// Specify to validate that you're deploying the right version.
	CodeSha256 *string `json:"codeSha256"`
	// Description of the version.
	Description *string `json:"description"`
	// Specifies a provisioned concurrency configuration for a function's version.
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions"`
	// Whether to retain old versions of this function when a new version is created.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
}

// Properties for a new Lambda version.
//
// TODO: EXAMPLE
//
type VersionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `json:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts"`
	// SHA256 of the version of the Lambda source code.
	//
	// Specify to validate that you're deploying the right version.
	CodeSha256 *string `json:"codeSha256"`
	// Description of the version.
	Description *string `json:"description"`
	// Specifies a provisioned concurrency configuration for a function's version.
	ProvisionedConcurrentExecutions *float64 `json:"provisionedConcurrentExecutions"`
	// Whether to retain old versions of this function when a new version is created.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Function to get the value of.
	Lambda IFunction `json:"lambda"`
}

// A version/weight pair for routing traffic to Lambda functions.
//
// TODO: EXAMPLE
//
type VersionWeight struct {
	// The version to route traffic to.
	Version IVersion `json:"version"`
	// How much weight to assign to this version (0..1).
	Weight *float64 `json:"weight"`
}

