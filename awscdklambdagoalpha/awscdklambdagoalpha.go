// The CDK Construct Library for AWS Lambda in Golang
package awscdklambdagoalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Bundling options.
//
// TODO: EXAMPLE
//
// Experimental.
type BundlingOptions struct {
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Experimental.
	AssetHash *string `json:"assetHash"`
	// Determines how the asset hash is calculated. Assets will get rebuilt and uploaded only if their hash has changed.
	//
	// If the asset hash is set to `OUTPUT` (default), the hash is calculated
	// after bundling. This means that any change in the output will cause
	// the asset to be invalidated and uploaded. Bear in mind that the
	// go binary that is output can be different depending on the environment
	// that it was compiled in. If you want to control when the output is changed
	// it is recommended that you use immutable build images such as
	// `public.ecr.aws/bitnami/golang:1.16.3-debian-10-r16`.
	//
	// If the asset hash is set to `SOURCE`, then only changes to the source
	// directory will cause the asset to rebuild. If your go project has multiple
	// Lambda functions this means that an update to any one function could cause
	// all the functions to be rebuilt and uploaded.
	// Experimental.
	AssetHashType awscdk.AssetHashType `json:"assetHashType"`
	// Build arguments to pass when building the bundling image.
	// Experimental.
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Whether or not to enable cgo during go build.
	//
	// This will set the CGO_ENABLED environment variable
	// Experimental.
	CgoEnabled *bool `json:"cgoEnabled"`
	// Command hooks.
	// Experimental.
	CommandHooks ICommandHooks `json:"commandHooks"`
	// A custom bundling Docker image.
	// Experimental.
	DockerImage awscdk.DockerImage `json:"dockerImage"`
	// Environment variables defined when go runs.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// Force bundling in a Docker container even if local bundling is possible.
	// Experimental.
	ForcedDockerBundling *bool `json:"forcedDockerBundling"`
	// List of additional flags to use while building.
	//
	// For example:
	// ['ldflags "-s -w"']
	// Experimental.
	GoBuildFlags *[]*string `json:"goBuildFlags"`
}

// A Golang Lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type GoFunction interface {
	awslambda.Function
	Architecture() awslambda.Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	CurrentVersion() awslambda.Version
	DeadLetterQueue() awssqs.IQueue
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	LatestVersion() awslambda.IVersion
	LogGroup() awslogs.ILogGroup
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Runtime() awslambda.Runtime
	Stack() awscdk.Stack
	AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function
	AddEventSource(source awslambda.IEventSource)
	AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping
	AddLayers(layers ...awslambda.ILayerVersion)
	AddPermission(id *string, permission *awslambda.Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions)
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

// The jsii proxy struct for GoFunction
type jsiiProxy_GoFunction struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_GoFunction) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) CurrentVersion() awslambda.Version {
	var returns awslambda.Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Runtime() awslambda.Runtime {
	var returns awslambda.Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewGoFunction(scope constructs.Construct, id *string, props *GoFunctionProps) GoFunction {
	_init_.Initialize()

	j := jsiiProxy_GoFunction{}

	_jsii_.Create(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGoFunction_Override(g GoFunction, scope constructs.Construct, id *string, props *GoFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		[]interface{}{scope, id, props},
		g,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
// Experimental.
func GoFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
// Experimental.
func GoFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
// Experimental.
func GoFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
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
func GoFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GoFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
// Experimental.
func GoFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
// Experimental.
func GoFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
// Experimental.
func GoFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
// Experimental.
func GoFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
// Experimental.
func GoFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
// Experimental.
func GoFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
// Experimental.
func GoFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-lambda-go-alpha.GoFunction",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Adds an environment variable to this Lambda function.
//
// If this is a ref to a Lambda function, this operation results in a no-op.
// Experimental.
func (g *jsiiProxy_GoFunction) AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function {
	var returns awslambda.Function

	_jsii_.Invoke(
		g,
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
// Experimental.
func (g *jsiiProxy_GoFunction) AddEventSource(source awslambda.IEventSource) {
	_jsii_.InvokeVoid(
		g,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
// Experimental.
func (g *jsiiProxy_GoFunction) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		g,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds one or more Lambda Layers to this Lambda function.
// Experimental.
func (g *jsiiProxy_GoFunction) AddLayers(layers ...awslambda.ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		g,
		"addLayers",
		args,
	)
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
// Experimental.
func (g *jsiiProxy_GoFunction) AddPermission(id *string, permission *awslambda.Permission) {
	_jsii_.InvokeVoid(
		g,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
// Experimental.
func (g *jsiiProxy_GoFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		g,
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
// Experimental.
func (g *jsiiProxy_GoFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
// Experimental.
func (g *jsiiProxy_GoFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		g,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

// Experimental.
func (g *jsiiProxy_GoFunction) GeneratePhysicalName() *string {
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
func (g *jsiiProxy_GoFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (g *jsiiProxy_GoFunction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
// Experimental.
func (g *jsiiProxy_GoFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
// Experimental.
func (g *jsiiProxy_GoFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// How long execution of this Lambda takes.
//
// Average over 5 minutes
// Experimental.
func (g *jsiiProxy_GoFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
// Experimental.
func (g *jsiiProxy_GoFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is invoked.
//
// Sum over 5 minutes
// Experimental.
func (g *jsiiProxy_GoFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How often this Lambda is throttled.
//
// Sum over 5 minutes
// Experimental.
func (g *jsiiProxy_GoFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GoFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a GolangFunction.
//
// TODO: EXAMPLE
//
// Experimental.
type GoFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	// Experimental.
	MaxEventAge awscdk.Duration `json:"maxEventAge"`
	// The destination for failed invocations.
	// Experimental.
	OnFailure awslambda.IDestination `json:"onFailure"`
	// The destination for successful invocations.
	// Experimental.
	OnSuccess awslambda.IDestination `json:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	// Experimental.
	RetryAttempts *float64 `json:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	// Experimental.
	AllowAllOutbound *bool `json:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	// Experimental.
	AllowPublicSubnet *bool `json:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	// Experimental.
	Architecture awslambda.Architecture `json:"architecture"`
	// Code signing config associated with this function.
	// Experimental.
	CodeSigningConfig awslambda.ICodeSigningConfig `json:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	// Experimental.
	CurrentVersionOptions *awslambda.VersionOptions `json:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	// Experimental.
	DeadLetterQueueEnabled *bool `json:"deadLetterQueueEnabled"`
	// A description of the function.
	// Experimental.
	Description *string `json:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	// Experimental.
	EnvironmentEncryption awskms.IKey `json:"environmentEncryption"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	// Experimental.
	Events *[]awslambda.IEventSource `json:"events"`
	// The filesystem configuration for the lambda function.
	// Experimental.
	Filesystem awslambda.FileSystem `json:"filesystem"`
	// A name for the function.
	// Experimental.
	FunctionName *string `json:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	// Experimental.
	InitialPolicy *[]awsiam.PolicyStatement `json:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	// Experimental.
	InsightsVersion awslambda.LambdaInsightsVersion `json:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	// Experimental.
	Layers *[]awslambda.ILayerVersion `json:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	// Experimental.
	LogRetention awslogs.RetentionDays `json:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	// Experimental.
	LogRetentionRetryOptions *awslambda.LogRetentionRetryOptions `json:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	LogRetentionRole awsiam.IRole `json:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	// Experimental.
	MemorySize *float64 `json:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	// Experimental.
	Profiling *bool `json:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	// Experimental.
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `json:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	// Experimental.
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
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	// Experimental.
	Tracing awslambda.Tracing `json:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets"`
	// The path to the folder or file that contains the main application entry point files for the project.
	//
	// This accepts either a path to a directory or file.
	//
	// If a directory path is provided then it will assume there is a Go entry file (i.e. `main.go`) and
	// will construct the build command using the directory path.
	//
	// For example, if you provide the entry as:
	//
	//      entry: 'my-lambda-app/cmd/api'
	//
	// Then the `go build` command would be:
	//
	//      `go build ./cmd/api`
	//
	// If a path to a file is provided then it will use the filepath in the build command.
	//
	// For example, if you provide the entry as:
	//
	//      entry: 'my-lambda-app/cmd/api/main.go'
	//
	// Then the `go build` command would be:
	//
	//      `go build ./cmd/api/main.go`
	// Experimental.
	Entry *string `json:"entry"`
	// Bundling options.
	// Experimental.
	Bundling *BundlingOptions `json:"bundling"`
	// Directory containing your go.mod file.
	//
	// This will accept either a directory path containing a `go.mod` file
	// or a filepath to your `go.mod` file (i.e. `path/to/go.mod`).
	//
	// This will be used as the source of the volume mounted in the Docker
	// container and will be the directory where it will run `go build` from.
	// Experimental.
	ModuleDir *string `json:"moduleDir"`
	// The runtime environment.
	//
	// Only runtimes of the Golang family and provided family are supported.
	// Experimental.
	Runtime awslambda.Runtime `json:"runtime"`
}

// Command hooks.
//
// These commands will run in the environment in which bundling occurs: inside
// the container for Docker bundling or on the host OS for local bundling.
//
// Commands are chained with `&&`.
//
// ```text
// {
//    // Run tests prior to bundling
//    beforeBundling(inputDir: string, outputDir: string): string[] {
//      return [`go test -mod=vendor ./...`];
//    }
//    // ...
// }
// ```
// Experimental.
type ICommandHooks interface {
	// Returns commands to run after bundling.
	//
	// Commands are chained with `&&`.
	// Experimental.
	AfterBundling(inputDir *string, outputDir *string) *[]*string
	// Returns commands to run before bundling.
	//
	// Commands are chained with `&&`.
	// Experimental.
	BeforeBundling(inputDir *string, outputDir *string) *[]*string
}

// The jsii proxy for ICommandHooks
type jsiiProxy_ICommandHooks struct {
	_ byte // padding
}

func (i *jsiiProxy_ICommandHooks) AfterBundling(inputDir *string, outputDir *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"afterBundling",
		[]interface{}{inputDir, outputDir},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICommandHooks) BeforeBundling(inputDir *string, outputDir *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"beforeBundling",
		[]interface{}{inputDir, outputDir},
		&returns,
	)

	return returns
}

