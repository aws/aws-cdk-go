package awslambdanodejs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdanodejs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Bundling options.
//
// TODO: EXAMPLE
//
type BundlingOptions struct {
	// Specify a custom hash for this asset.
	//
	// For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	AssetHash *string `json:"assetHash" yaml:"assetHash"`
	// Use this to insert an arbitrary string at the beginning of generated JavaScript files.
	//
	// This is similar to footer which inserts at the end instead of the beginning.
	//
	// This is commonly used to insert comments:
	Banner *string `json:"banner" yaml:"banner"`
	// Build arguments to pass when building the bundling image.
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// The charset to use for esbuild's output.
	//
	// By default esbuild's output is ASCII-only. Any non-ASCII characters are escaped
	// using backslash escape sequences. Using escape sequences makes the generated output
	// slightly bigger, and also makes it harder to read. If you would like for esbuild to print
	// the original characters without using escape sequences, use `Charset.UTF8`.
	// See: https://esbuild.github.io/api/#charset
	//
	Charset Charset `json:"charset" yaml:"charset"`
	// Command hooks.
	CommandHooks ICommandHooks `json:"commandHooks" yaml:"commandHooks"`
	// Replace global identifiers with constant expressions.
	//
	// For example, `{ 'process.env.DEBUG': 'true' }`.
	//
	// Another example, `{ 'process.env.API_KEY': JSON.stringify('xxx-xxxx-xxx') }`.
	Define *map[string]*string `json:"define" yaml:"define"`
	// A custom bundling Docker image.
	//
	// This image should have esbuild installed globally. If you plan to use `nodeModules`
	// it should also have `npm` or `yarn` depending on the lock file you're using.
	//
	// See https://github.com/aws/aws-cdk/blob/master/packages/%40aws-cdk/aws-lambda-nodejs/lib/Dockerfile
	// for the default image provided by @aws-cdk/aws-lambda-nodejs.
	DockerImage awscdk.DockerImage `json:"dockerImage" yaml:"dockerImage"`
	// Environment variables defined when bundling runs.
	Environment *map[string]*string `json:"environment" yaml:"environment"`
	// The version of esbuild to use when running in a Docker container.
	EsbuildVersion *string `json:"esbuildVersion" yaml:"esbuildVersion"`
	// A list of modules that should be considered as externals (already available in the runtime).
	ExternalModules *[]*string `json:"externalModules" yaml:"externalModules"`
	// Use this to insert an arbitrary string at the end of generated JavaScript files.
	//
	// This is similar to banner which inserts at the beginning instead of the end.
	//
	// This is commonly used to insert comments
	Footer *string `json:"footer" yaml:"footer"`
	// Force bundling in a Docker container even if local bundling is possible.
	//
	// This is useful if your function relies on node modules
	// that should be installed (`nodeModules`) in a Lambda compatible
	// environment.
	ForceDockerBundling *bool `json:"forceDockerBundling" yaml:"forceDockerBundling"`
	// Output format for the generated JavaScript files.
	Format OutputFormat `json:"format" yaml:"format"`
	// This option allows you to automatically replace a global variable with an import from another file.
	// See: https://esbuild.github.io/api/#inject
	//
	Inject *[]*string `json:"inject" yaml:"inject"`
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
	KeepNames *bool `json:"keepNames" yaml:"keepNames"`
	// Use loaders to change how a given input file is interpreted.
	//
	// Configuring a loader for a given file type lets you load that file type with
	// an `import` statement or a `require` call.
	// See: https://esbuild.github.io/api/#loader
	//
	// For example, `{ '.png': 'dataurl' }`.
	//
	Loader *map[string]*string `json:"loader" yaml:"loader"`
	// Log level for esbuild.
	//
	// This is also propagated to the package manager and
	// applies to its specific install command.
	LogLevel LogLevel `json:"logLevel" yaml:"logLevel"`
	// How to determine the entry point for modules.
	//
	// Try ['module', 'main'] to default to ES module versions.
	MainFields *[]*string `json:"mainFields" yaml:"mainFields"`
	// This option tells esbuild to write out a JSON file relative to output directory with metadata about the build.
	//
	// The metadata in this JSON file follows this schema (specified using TypeScript syntax):
	//
	// ```text
	// {
	//    outputs: {
	//      [path: string]: {
	//        bytes: number
	//        inputs: {
	//          [path: string]: { bytesInOutput: number }
	//        }
	//        imports: { path: string }[]
	//        exports: string[]
	//      }
	//    }
	// }
	// ```
	// This data can then be analyzed by other tools. For example,
	// bundle buddy can consume esbuild's metadata format and generates a treemap visualization
	// of the modules in your bundle and how much space each one takes up.
	// See: https://esbuild.github.io/api/#metafile
	//
	Metafile *bool `json:"metafile" yaml:"metafile"`
	// Whether to minify files when bundling.
	Minify *bool `json:"minify" yaml:"minify"`
	// A list of modules that should be installed instead of bundled.
	//
	// Modules are
	// installed in a Lambda compatible environment only when bundling runs in
	// Docker.
	NodeModules *[]*string `json:"nodeModules" yaml:"nodeModules"`
	// Run compilation using tsc before running file through bundling step.
	//
	// This usually is not required unless you are using new experimental features that
	// are only supported by typescript's `tsc` compiler.
	// One example of such feature is `emitDecoratorMetadata`.
	PreCompilation *bool `json:"preCompilation" yaml:"preCompilation"`
	// Whether to include source maps when bundling.
	SourceMap *bool `json:"sourceMap" yaml:"sourceMap"`
	// Source map mode to be used when bundling.
	// See: https://esbuild.github.io/api/#sourcemap
	//
	SourceMapMode SourceMapMode `json:"sourceMapMode" yaml:"sourceMapMode"`
	// Whether to include original source code in source maps when bundling.
	// See: https://esbuild.github.io/api/#sources-content
	//
	SourcesContent *bool `json:"sourcesContent" yaml:"sourcesContent"`
	// Target environment for the generated JavaScript code.
	// See: https://esbuild.github.io/api/#target
	//
	Target *string `json:"target" yaml:"target"`
	// Normally the esbuild automatically discovers `tsconfig.json` files and reads their contents during a build.
	//
	// However, you can also configure a custom `tsconfig.json` file to use instead.
	//
	// This is similar to entry path, you need to provide path to your custom `tsconfig.json`.
	//
	// This can be useful if you need to do multiple builds of the same code with different settings.
	//
	// For example, `{ 'tsconfig': 'path/custom.tsconfig.json' }`.
	Tsconfig *string `json:"tsconfig" yaml:"tsconfig"`
}

// Charset for esbuild's output.
//
// TODO: EXAMPLE
//
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
// The following example (specified in TypeScript) copies a file from the input
// directory to the output directory to include it in the bundled asset:
//
// ```text
// afterBundling(inputDir: string, outputDir: string): string[]{
//    return [`cp ${inputDir}/my-binary.node ${outputDir}`];
// }
// ```
type ICommandHooks interface {
	// Returns commands to run after bundling.
	//
	// Commands are chained with `&&`.
	AfterBundling(inputDir *string, outputDir *string) *[]*string
	// Returns commands to run before bundling.
	//
	// Commands are chained with `&&`.
	BeforeBundling(inputDir *string, outputDir *string) *[]*string
	// Returns commands to run before installing node modules.
	//
	// This hook only runs when node modules are installed.
	//
	// Commands are chained with `&&`.
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

// Log levels for esbuild and package managers' install commands.
//
// TODO: EXAMPLE
//
type LogLevel string

const (
	LogLevel_INFO LogLevel = "INFO"
	LogLevel_WARNING LogLevel = "WARNING"
	LogLevel_ERROR LogLevel = "ERROR"
	LogLevel_SILENT LogLevel = "SILENT"
)

// A Node.js Lambda function bundled using esbuild.
//
// TODO: EXAMPLE
//
type NodejsFunction interface {
	awslambda.Function
	Architecture() awslambda.Architecture
	CanCreatePermissions() *bool
	Connections() awsec2.Connections
	CurrentVersion() awslambda.Version
	DeadLetterQueue() awssqs.IQueue
	DeadLetterTopic() awssns.ITopic
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
	Timeout() awscdk.Duration
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

// The jsii proxy struct for NodejsFunction
type jsiiProxy_NodejsFunction struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_NodejsFunction) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_NodejsFunction) DeadLetterTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"deadLetterTopic",
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

func (j *jsiiProxy_NodejsFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodejsFunction) PermissionsNode() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_NodejsFunction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewNodejsFunction(scope constructs.Construct, id *string, props *NodejsFunctionProps) NodejsFunction {
	_init_.Initialize()

	j := jsiiProxy_NodejsFunction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewNodejsFunction_Override(n NodejsFunction, scope constructs.Construct, id *string, props *NodejsFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		[]interface{}{scope, id, props},
		n,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
func NodejsFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
func NodejsFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
func NodejsFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a lambda function into the CDK using its name.
func NodejsFunction_FromFunctionName(scope constructs.Construct, id *string, functionName *string) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"fromFunctionName",
		[]interface{}{scope, id, functionName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func NodejsFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func NodejsFunction_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
func NodejsFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
func NodejsFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
func NodejsFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
func NodejsFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
func NodejsFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
func NodejsFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
func NodejsFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda_nodejs.NodejsFunction",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Adds an environment variable to this Lambda function.
//
// If this is a ref to a Lambda function, this operation results in a no-op.
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
func (n *jsiiProxy_NodejsFunction) AddEventSource(source awslambda.IEventSource) {
	_jsii_.InvokeVoid(
		n,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
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
func (n *jsiiProxy_NodejsFunction) AddPermission(id *string, permission *awslambda.Permission) {
	_jsii_.InvokeVoid(
		n,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
func (n *jsiiProxy_NodejsFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		n,
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
func (n *jsiiProxy_NodejsFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		n,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
func (n *jsiiProxy_NodejsFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		n,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

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

// Returns a string representation of this construct.
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

// Properties for a NodejsFunction.
//
// TODO: EXAMPLE
//
type NodejsFunctionProps struct {
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
	//
	// If SNS topic is desired, specify `deadLetterTopic` property instead.
	DeadLetterQueue awssqs.IQueue `json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	DeadLetterQueueEnabled *bool `json:"deadLetterQueueEnabled" yaml:"deadLetterQueueEnabled"`
	// The SNS topic to use as a DLQ.
	//
	// Note that if `deadLetterQueueEnabled` is set to `true`, an SQS queue will be created
	// rather than an SNS topic. Using an SNS topic as a DLQ requires this property to be set explicitly.
	DeadLetterTopic awssns.ITopic `json:"deadLetterTopic" yaml:"deadLetterTopic"`
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
	// Whether to automatically reuse TCP connections when working with the AWS SDK for JavaScript.
	//
	// This sets the `AWS_NODEJS_CONNECTION_REUSE_ENABLED` environment variable
	// to `1`.
	// See: https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/node-reusing-connections.html
	//
	AwsSdkConnectionReuse *bool `json:"awsSdkConnectionReuse" yaml:"awsSdkConnectionReuse"`
	// Bundling options.
	Bundling *BundlingOptions `json:"bundling" yaml:"bundling"`
	// The path to the dependencies lock file (`yarn.lock` or `package-lock.json`).
	//
	// This will be used as the source for the volume mounted in the Docker
	// container.
	//
	// Modules specified in `nodeModules` will be installed using the right
	// installer (`npm` or `yarn`) along with this lock file.
	DepsLockFilePath *string `json:"depsLockFilePath" yaml:"depsLockFilePath"`
	// Path to the entry file (JavaScript or TypeScript).
	Entry *string `json:"entry" yaml:"entry"`
	// The name of the exported handler in the entry file.
	Handler *string `json:"handler" yaml:"handler"`
	// The path to the directory containing project config files (`package.json` or `tsconfig.json`).
	ProjectRoot *string `json:"projectRoot" yaml:"projectRoot"`
	// The runtime environment.
	//
	// Only runtimes of the Node.js family are
	// supported.
	Runtime awslambda.Runtime `json:"runtime" yaml:"runtime"`
}

// Output format for the generated JavaScript files.
//
// TODO: EXAMPLE
//
type OutputFormat string

const (
	OutputFormat_CJS OutputFormat = "CJS"
	OutputFormat_ESM OutputFormat = "ESM"
)

// SourceMap mode for esbuild.
//
// TODO: EXAMPLE
//
// See: https://esbuild.github.io/api/#sourcemap
//
type SourceMapMode string

const (
	SourceMapMode_DEFAULT SourceMapMode = "DEFAULT"
	SourceMapMode_EXTERNAL SourceMapMode = "EXTERNAL"
	SourceMapMode_INLINE SourceMapMode = "INLINE"
	SourceMapMode_BOTH SourceMapMode = "BOTH"
)

