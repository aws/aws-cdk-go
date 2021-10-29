package awslambdanodejs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdanodejs/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/constructs-go/constructs/v3"
)

// Bundling options.
// Experimental.
type BundlingOptions struct {
	// Use this to insert an arbitrary string at the beginning of generated JavaScript files.
	//
	// This is similar to footer which inserts at the end instead of the beginning.
	//
	// This is commonly used to insert comments:
	// Experimental.
	Banner *string `json:"banner"`
	// Build arguments to pass when building the bundling image.
	// Experimental.
	BuildArgs *map[string]*string `json:"buildArgs"`
	// The charset to use for esbuild's output.
	//
	// By default esbuild's output is ASCII-only. Any non-ASCII characters are escaped
	// using backslash escape sequences. Using escape sequences makes the generated output
	// slightly bigger, and also makes it harder to read. If you would like for esbuild to print
	// the original characters without using escape sequences, use `Charset.UTF8`.
	// See: https://esbuild.github.io/api/#charset
	//
	// Experimental.
	Charset Charset `json:"charset"`
	// Command hooks.
	// Experimental.
	CommandHooks ICommandHooks `json:"commandHooks"`
	// Replace global identifiers with constant expressions.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Define *map[string]*string `json:"define"`
	// A custom bundling Docker image.
	//
	// This image should have esbuild installed globally. If you plan to use `nodeModules`
	// it should also have `npm` or `yarn` depending on the lock file you're using.
	//
	// See https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk/aws-lambda-nodejs/lib/Dockerfile
	// for the default image provided by @aws-cdk/aws-lambda-nodejs.
	// Experimental.
	DockerImage awscdk.DockerImage `json:"dockerImage"`
	// Environment variables defined when bundling runs.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The version of esbuild to use when running in a Docker container.
	// Experimental.
	EsbuildVersion *string `json:"esbuildVersion"`
	// A list of modules that should be considered as externals (already available in the runtime).
	// Experimental.
	ExternalModules *[]*string `json:"externalModules"`
	// Use this to insert an arbitrary string at the end of generated JavaScript files.
	//
	// This is similar to banner which inserts at the beginning instead of the end.
	//
	// This is commonly used to insert comments
	// Experimental.
	Footer *string `json:"footer"`
	// Force bundling in a Docker container even if local bundling is possible.
	//
	// This is useful if your function relies on node modules
	// that should be installed (`nodeModules`) in a Lambda compatible
	// environment.
	// Experimental.
	ForceDockerBundling *bool `json:"forceDockerBundling"`
	// Whether to preserve the original `name` values even in minified code.
	//
	// In JavaScript the `name` property on functions and classes defaults to a
	// nearby identifier in the source code.
	//
	// However, minification renames symbols to reduce code size and bundling
	// sometimes need to rename symbols to avoid collisions. That changes value of
	// the `name` property for many of these cases. This is usually fine because
	// the `name` property is normally only used for debugging. However, some
	// frameworks rely on the `name` property for registration and binding purposes.
	// If this is the case, you can enable this option to preserve the original
	// `name` values even in minified code.
	// Experimental.
	KeepNames *bool `json:"keepNames"`
	// Use loaders to change how a given input file is interpreted.
	//
	// Configuring a loader for a given file type lets you load that file type with
	// an `import` statement or a `require` call.
	//
	// TODO: EXAMPLE
	//
	// See: https://esbuild.github.io/api/#loader
	//
	// Experimental.
	Loader *map[string]*string `json:"loader"`
	// Log level for esbuild.
	// Experimental.
	LogLevel LogLevel `json:"logLevel"`
	// This option tells esbuild to write out a JSON file relative to output directory with metadata about the build.
	//
	// The metadata in this JSON file follows this schema (specified using TypeScript syntax):
	//
	// ```typescript
	//   {
	//      outputs: {
	//           [path: string]: {
	//             bytes: number
	//             inputs: {
	//               [path: string]: { bytesInOutput: number }
	//             }
	//             imports: { path: string }[]
	//             exports: string[]
	//           }
	//         }
	//      }
	// }
	// ```
	// This data can then be analyzed by other tools. For example,
	// bundle buddy can consume esbuild's metadata format and generates a treemap visualization
	// of the modules in your bundle and how much space each one takes up.
	// See: https://esbuild.github.io/api/#metafile
	//
	// Experimental.
	Metafile *bool `json:"metafile"`
	// Whether to minify files when bundling.
	// Experimental.
	Minify *bool `json:"minify"`
	// A list of modules that should be installed instead of bundled.
	//
	// Modules are
	// installed in a Lambda compatible environment only when bundling runs in
	// Docker.
	// Experimental.
	NodeModules *[]*string `json:"nodeModules"`
	// Run compilation using tsc before running file through bundling step.
	//
	// This usually is not required unless you are using new experimental features that
	// are only supported by typescript's `tsc` compiler.
	// One example of such feature is `emitDecoratorMetadata`.
	// Experimental.
	PreCompilation *bool `json:"preCompilation"`
	// Whether to include source maps when bundling.
	// Experimental.
	SourceMap *bool `json:"sourceMap"`
	// Source map mode to be used when bundling.
	// See: https://esbuild.github.io/api/#sourcemap
	//
	// Experimental.
	SourceMapMode SourceMapMode `json:"sourceMapMode"`
	// Target environment for the generated JavaScript code.
	// See: https://esbuild.github.io/api/#target
	//
	// Experimental.
	Target *string `json:"target"`
	// Normally the esbuild automatically discovers `tsconfig.json` files and reads their contents during a build.
	//
	// However, you can also configure a custom `tsconfig.json` file to use instead.
	//
	// This is similar to entry path, you need to provide path to your custom `tsconfig.json`.
	//
	// This can be useful if you need to do multiple builds of the same code with different settings.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Tsconfig *string `json:"tsconfig"`
}

// Charset for esbuild's output.
// Experimental.
type Charset string

const (
	Charset_ASCII Charset = "ASCII"
	Charset_UTF8 Charset = "UTF8"
)

// Command hooks.
//
// These commands will run in the environment in which bundling occurs: inside
// the container for Docker bundling or on the host OS for local bundling.
//
// Commands are chained with `&&`.
//
// TODO: EXAMPLE
//
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
	// Returns commands to run before installing node modules.
	//
	// This hook only runs when node modules are installed.
	//
	// Commands are chained with `&&`.
	// Experimental.
	BeforeInstall(inputDir *string, outputDir *string) *[]*string
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

func (i *jsiiProxy_ICommandHooks) BeforeInstall(inputDir *string, outputDir *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"beforeInstall",
		[]interface{}{inputDir, outputDir},
		&returns,
	)

	return returns
}

// Log level for esbuild.
// Experimental.
type LogLevel string

const (
	LogLevel_INFO LogLevel = "INFO"
	LogLevel_WARNING LogLevel = "WARNING"
	LogLevel_ERROR LogLevel = "ERROR"
	LogLevel_SILENT LogLevel = "SILENT"
)

// A Node.js Lambda function bundled using esbuild.
// Experimental.
type NodejsFunction interface {
	awslambda.Function
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
	Node() awscdk.ConstructNode
	PermissionsNode() awscdk.ConstructNode
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
	AddVersion(name *string, codeSha256 *string, description *string, provisionedExecutions *float64, asyncInvokeConfig *awslambda.EventInvokeConfigOptions) awslambda.Version
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for NodejsFunction
type jsiiProxy_NodejsFunction struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_NodejsFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) CurrentVersion() awslambda.Version {
	var returns awslambda.Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) PermissionsNode() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) Runtime() awslambda.Runtime {
	var returns awslambda.Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodejsFunction(scope awscdk.Construct, id *string, props *NodejsFunctionProps) NodejsFunction {
	_init_.Initialize()

	j := jsiiProxy_NodejsFunction{}

	_jsii_.Create(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNodejsFunction_Override(n NodejsFunction, scope awscdk.Construct, id *string, props *NodejsFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		[]interface{}{scope, id, props},
		n,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
// Experimental.
func NodejsFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
// Experimental.
func NodejsFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
// Experimental.
func NodejsFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func NodejsFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func NodejsFunction_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
// Experimental.
func NodejsFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
// Experimental.
func NodejsFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
// Experimental.
func NodejsFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
// Experimental.
func NodejsFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
// Experimental.
func NodejsFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
// Experimental.
func NodejsFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
// Experimental.
func NodejsFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_nodejs.NodejsFunction",
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
func (n *jsiiProxy_NodejsFunction) AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function {
	var returns awslambda.Function

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NodejsFunction) AddEventSource(source awslambda.IEventSource) {
	_jsii_.InvokeVoid(
		n,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
// Experimental.
func (n *jsiiProxy_NodejsFunction) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		n,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds one or more Lambda Layers to this Lambda function.
// Experimental.
func (n *jsiiProxy_NodejsFunction) AddLayers(layers ...awslambda.ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addLayers",
		args,
	)
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
// Experimental.
func (n *jsiiProxy_NodejsFunction) AddPermission(id *string, permission *awslambda.Permission) {
	_jsii_.InvokeVoid(
		n,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
// Experimental.
func (n *jsiiProxy_NodejsFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		n,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

// Add a new version for this Lambda.
//
// If you want to deploy through CloudFormation and use aliases, you need to
// add a new version (with a new name) to your Lambda every time you want to
// deploy an update. An alias can then refer to the newly created Version.
//
// All versions should have distinct names, and you should not delete versions
// as long as your Alias needs to refer to them.
//
// Returns: A new Version object.
// Deprecated: This method will create an AWS::Lambda::Version resource which
// snapshots the AWS Lambda function *at the time of its creation* and it
// won't get updated when the function changes. Instead, use
// `this.currentVersion` to obtain a reference to a version resource that gets
// automatically recreated when the function configuration (or code) changes.
func (n *jsiiProxy_NodejsFunction) AddVersion(name *string, codeSha256 *string, description *string, provisionedExecutions *float64, asyncInvokeConfig *awslambda.EventInvokeConfigOptions) awslambda.Version {
	var returns awslambda.Version

	_jsii_.Invoke(
		n,
		"addVersion",
		[]interface{}{name, codeSha256, description, provisionedExecutions, asyncInvokeConfig},
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
// Experimental.
func (n *jsiiProxy_NodejsFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
// Experimental.
func (n *jsiiProxy_NodejsFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		n,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

// Experimental.
func (n *jsiiProxy_NodejsFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NodejsFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NodejsFunction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
// Experimental.
func (n *jsiiProxy_NodejsFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		n,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
// Experimental.
func (n *jsiiProxy_NodejsFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NodejsFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NodejsFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NodejsFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
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
func (n *jsiiProxy_NodejsFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		n,
		"metricThrottles",
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
func (n *jsiiProxy_NodejsFunction) OnPrepare() {
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
func (n *jsiiProxy_NodejsFunction) OnSynthesize(session constructs.ISynthesisSession) {
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
func (n *jsiiProxy_NodejsFunction) OnValidate() *[]*string {
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
func (n *jsiiProxy_NodejsFunction) Prepare() {
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
func (n *jsiiProxy_NodejsFunction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		n,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (n *jsiiProxy_NodejsFunction) ToString() *string {
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
func (n *jsiiProxy_NodejsFunction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a NodejsFunction.
// Experimental.
type NodejsFunctionProps struct {
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
	// DEPRECATED.
	// Deprecated: use `architecture`
	Architectures *[]awslambda.Architecture `json:"architectures"`
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
	// What security group to associate with the Lambda's network interfaces. This property is being deprecated, consider using securityGroups instead.
	//
	// Only used if 'vpc' is supplied.
	//
	// Use securityGroups property instead.
	// Function constructor will throw an error if both are specified.
	// Deprecated: - This property is deprecated, use securityGroups instead
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
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
	// Whether to automatically reuse TCP connections when working with the AWS SDK for JavaScript.
	//
	// This sets the `AWS_NODEJS_CONNECTION_REUSE_ENABLED` environment variable
	// to `1`.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/node-reusing-connections.html
	//
	// Experimental.
	AwsSdkConnectionReuse *bool `json:"awsSdkConnectionReuse"`
	// Bundling options.
	// Experimental.
	Bundling *BundlingOptions `json:"bundling"`
	// The path to the dependencies lock file (`yarn.lock` or `package-lock.json`).
	//
	// This will be used as the source for the volume mounted in the Docker
	// container.
	//
	// Modules specified in `nodeModules` will be installed using the right
	// installer (`npm` or `yarn`) along with this lock file.
	// Experimental.
	DepsLockFilePath *string `json:"depsLockFilePath"`
	// Path to the entry file (JavaScript or TypeScript).
	// Experimental.
	Entry *string `json:"entry"`
	// The name of the exported handler in the entry file.
	// Experimental.
	Handler *string `json:"handler"`
	// The path to the directory containing project config files (`package.json` or `tsconfig.json`).
	// Experimental.
	ProjectRoot *string `json:"projectRoot"`
	// The runtime environment.
	//
	// Only runtimes of the Node.js family are
	// supported.
	// Experimental.
	Runtime awslambda.Runtime `json:"runtime"`
}

// SourceMap mode for esbuild.
// See: https://esbuild.github.io/api/#sourcemap
//
// Experimental.
type SourceMapMode string

const (
	SourceMapMode_DEFAULT SourceMapMode = "DEFAULT"
	SourceMapMode_EXTERNAL SourceMapMode = "EXTERNAL"
	SourceMapMode_INLINE SourceMapMode = "INLINE"
	SourceMapMode_BOTH SourceMapMode = "BOTH"
)

