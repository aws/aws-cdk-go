package experimental

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/experimental/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Lambda@Edge function.
//
// Convenience resource for requesting a Lambda function in the 'us-east-1' region for use with Lambda@Edge.
// Implements several restrictions enforced by Lambda@Edge.
//
// Note that this construct requires that the 'us-east-1' region has been bootstrapped.
// See https://docs.aws.amazon.com/cdk/latest/guide/bootstrapping.html or 'cdk bootstrap --help' for options.
//
// TODO: EXAMPLE
//
type EdgeFunction interface {
	awscdk.Resource
	awslambda.IVersion
	Architecture() awslambda.Architecture
	Connections() awsec2.Connections
	CurrentVersion() awslambda.IVersion
	EdgeArn() *string
	Env() *awscdk.ResourceEnvironment
	FunctionArn() *string
	FunctionName() *string
	GrantPrincipal() awsiam.IPrincipal
	IsBoundToVpc() *bool
	Lambda() awslambda.IFunction
	LatestVersion() awslambda.IVersion
	Node() constructs.Node
	PermissionsNode() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	Version() *string
	AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias
	AddEventSource(source awslambda.IEventSource)
	AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping
	AddPermission(id *string, permission *awslambda.Permission)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantInvoke(identity awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for EdgeFunction
type jsiiProxy_EdgeFunction struct {
	internal.Type__awscdkResource
	internal.Type__awslambdaIVersion
}

func (j *jsiiProxy_EdgeFunction) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) CurrentVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) EdgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) Lambda() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgeFunction) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewEdgeFunction(scope constructs.Construct, id *string, props *EdgeFunctionProps) EdgeFunction {
	_init_.Initialize()

	j := jsiiProxy_EdgeFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.experimental.EdgeFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewEdgeFunction_Override(e EdgeFunction, scope constructs.Construct, id *string, props *EdgeFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.experimental.EdgeFunction",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func EdgeFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.experimental.EdgeFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func EdgeFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.experimental.EdgeFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Defines an alias for this version.
func (e *jsiiProxy_EdgeFunction) AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias {
	var returns awslambda.Alias

	_jsii_.Invoke(
		e,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

// Adds an event source to this function.
func (e *jsiiProxy_EdgeFunction) AddEventSource(source awslambda.IEventSource) {
	_jsii_.InvokeVoid(
		e,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
func (e *jsiiProxy_EdgeFunction) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		e,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds a permission to the Lambda resource policy.
func (e *jsiiProxy_EdgeFunction) AddPermission(id *string, permission *awslambda.Permission) {
	_jsii_.InvokeVoid(
		e,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (e *jsiiProxy_EdgeFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		e,
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
func (e *jsiiProxy_EdgeFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (e *jsiiProxy_EdgeFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		e,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (e *jsiiProxy_EdgeFunction) GeneratePhysicalName() *string {
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
func (e *jsiiProxy_EdgeFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (e *jsiiProxy_EdgeFunction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
func (e *jsiiProxy_EdgeFunction) GrantInvoke(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		e,
		"grantInvoke",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda Return the given named metric for this Function.
func (e *jsiiProxy_EdgeFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the Duration of this Lambda How long execution of this Lambda takes.
//
// Average over 5 minutes
func (e *jsiiProxy_EdgeFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many invocations of this Lambda fail.
//
// Sum over 5 minutes
func (e *jsiiProxy_EdgeFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of this Lambda How often this Lambda is invoked.
//
// Sum over 5 minutes
func (e *jsiiProxy_EdgeFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of this Lambda How often this Lambda is throttled.
//
// Sum over 5 minutes
func (e *jsiiProxy_EdgeFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (e *jsiiProxy_EdgeFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for creating a Lambda@Edge function.
//
// TODO: EXAMPLE
//
type EdgeFunctionProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours
	MaxEventAge awscdk.Duration `json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure awslambda.IDestination `json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	OnSuccess awslambda.IDestination `json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2
	RetryAttempts *float64 `json:"retryAttempts" yaml:"retryAttempts"`
	// Whether to allow the Lambda to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets.
	AllowAllOutbound *bool `json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Lambda Functions in a public subnet can NOT access the internet.
	//
	// Use this property to acknowledge this limitation and still place the function in a public subnet.
	// See: https://stackoverflow.com/questions/52992085/why-cant-an-aws-lambda-function-inside-a-public-subnet-in-a-vpc-connect-to-the/52994841#52994841
	//
	AllowPublicSubnet *bool `json:"allowPublicSubnet" yaml:"allowPublicSubnet"`
	// The system architectures compatible with this lambda function.
	Architecture awslambda.Architecture `json:"architecture" yaml:"architecture"`
	// Code signing config associated with this function.
	CodeSigningConfig awslambda.ICodeSigningConfig `json:"codeSigningConfig" yaml:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	CurrentVersionOptions *awslambda.VersionOptions `json:"currentVersionOptions" yaml:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `json:"deadLetterQueueEnabled" yaml:"deadLetterQueueEnabled"`
	// A description of the function.
	Description *string `json:"description" yaml:"description"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	EnvironmentEncryption awskms.IKey `json:"environmentEncryption" yaml:"environmentEncryption"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	Events *[]awslambda.IEventSource `json:"events" yaml:"events"`
	// The filesystem configuration for the lambda function.
	Filesystem awslambda.FileSystem `json:"filesystem" yaml:"filesystem"`
	// A name for the function.
	FunctionName *string `json:"functionName" yaml:"functionName"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	InitialPolicy *[]awsiam.PolicyStatement `json:"initialPolicy" yaml:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-docker.html
	//
	InsightsVersion awslambda.LambdaInsightsVersion `json:"insightsVersion" yaml:"insightsVersion"`
	// A list of layers to add to the function's execution environment.
	//
	// You can configure your Lambda function to pull in
	// additional code during initialization in the form of layers. Layers are packages of libraries or other dependencies
	// that can be used by multiple functions.
	Layers *[]awslambda.ILayerVersion `json:"layers" yaml:"layers"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	LogRetention awslogs.RetentionDays `json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	LogRetentionRetryOptions *awslambda.LogRetentionRetryOptions `json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	LogRetentionRole awsiam.IRole `json:"logRetentionRole" yaml:"logRetentionRole"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	MemorySize *float64 `json:"memorySize" yaml:"memorySize"`
	// Enable profiling.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	Profiling *bool `json:"profiling" yaml:"profiling"`
	// Profiling Group.
	// See: https://docs.aws.amazon.com/codeguru/latest/profiler-ug/setting-up-lambda.html
	//
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `json:"profilingGroup" yaml:"profilingGroup"`
	// The maximum of concurrent executions you want to reserve for the function.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html
	//
	ReservedConcurrentExecutions *float64 `json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
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
	Role awsiam.IRole `json:"role" yaml:"role"`
	// The list of security groups to associate with the Lambda's network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// Enable AWS X-Ray Tracing for Lambda Function.
	Tracing awslambda.Tracing `json:"tracing" yaml:"tracing"`
	// VPC network to place Lambda network interfaces.
	//
	// Specify this if the Lambda function needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	VpcSubnets *awsec2.SubnetSelection `json:"vpcSubnets" yaml:"vpcSubnets"`
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	Code awslambda.Code `json:"code" yaml:"code"`
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
	Handler *string `json:"handler" yaml:"handler"`
	// The runtime environment for the Lambda function that you are uploading.
	//
	// For valid values, see the Runtime property in the AWS Lambda Developer
	// Guide.
	//
	// Use `Runtime.FROM_IMAGE` when when defining a function from a Docker image.
	Runtime awslambda.Runtime `json:"runtime" yaml:"runtime"`
	// The stack ID of Lambda@Edge function.
	StackId *string `json:"stackId" yaml:"stackId"`
}

