package awslambdapython

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
	"github.com/aws/aws-cdk-go/awscdk/awslambdapython/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/constructs-go/constructs/v3"
)

// Options for bundling.
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
	// Determines how asset hash is calculated. Assets will get rebuild and uploaded only if their hash has changed.
	//
	// If asset hash is set to `SOURCE` (default), then only changes to the source
	// directory will cause the asset to rebuild. This means, for example, that in
	// order to pick up a new dependency version, a change must be made to the
	// source tree. Ideally, this can be implemented by including a dependency
	// lockfile in your source tree or using fixed dependencies.
	//
	// If the asset hash is set to `OUTPUT`, the hash is calculated after
	// bundling. This means that any change in the output will cause the asset to
	// be invalidated and uploaded. Bear in mind that `pip` adds timestamps to
	// dependencies it installs, which implies that in this mode Python bundles
	// will _always_ get rebuild and uploaded. Normally this is an anti-pattern
	// since build
	// Experimental.
	AssetHashType awscdk.AssetHashType `json:"assetHashType"`
	// Optional build arguments to pass to the default container.
	//
	// This can be used to customize
	// the index URLs used for installing dependencies.
	// This is not used if a custom image is provided.
	// Experimental.
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Docker image to use for bundling.
	//
	// If no options are provided, the default bundling image
	// will be used. Dependencies will be installed using the default packaging commands
	// and copied over from into the Lambda asset.
	// Experimental.
	Image awscdk.DockerImage `json:"image"`
	// Output path suffix: the suffix for the directory into which the bundled output is written.
	// Experimental.
	OutputPathSuffix *string `json:"outputPathSuffix"`
}

// A Python Lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type PythonFunction interface {
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
	Node() awscdk.ConstructNode
	PermissionsNode() awscdk.ConstructNode
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

// The jsii proxy struct for PythonFunction
type jsiiProxy_PythonFunction struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_PythonFunction) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) CurrentVersion() awslambda.Version {
	var returns awslambda.Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) PermissionsNode() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Runtime() awslambda.Runtime {
	var returns awslambda.Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonFunction) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonFunction(scope awscdk.Construct, id *string, props *PythonFunctionProps) PythonFunction {
	_init_.Initialize()

	j := jsiiProxy_PythonFunction{}

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonFunction_Override(p PythonFunction, scope awscdk.Construct, id *string, props *PythonFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonFunction",
		[]interface{}{scope, id, props},
		p,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
// Experimental.
func PythonFunction_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_lambda_python.PythonFunction",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
// Experimental.
func PythonFunction_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
// Experimental.
func PythonFunction_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func PythonFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func PythonFunction_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
// Experimental.
func PythonFunction_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
// Experimental.
func PythonFunction_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
// Experimental.
func PythonFunction_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
// Experimental.
func PythonFunction_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
// Experimental.
func PythonFunction_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
// Experimental.
func PythonFunction_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
// Experimental.
func PythonFunction_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonFunction",
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
func (p *jsiiProxy_PythonFunction) AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function {
	var returns awslambda.Function

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) AddEventSource(source awslambda.IEventSource) {
	_jsii_.InvokeVoid(
		p,
		"addEventSource",
		[]interface{}{source},
	)
}

// Adds an event source that maps to this AWS Lambda function.
// Experimental.
func (p *jsiiProxy_PythonFunction) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		p,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Adds one or more Lambda Layers to this Lambda function.
// Experimental.
func (p *jsiiProxy_PythonFunction) AddLayers(layers ...awslambda.ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addLayers",
		args,
	)
}

// Adds a permission to the Lambda resource policy.
// See: Permission for details.
//
// Experimental.
func (p *jsiiProxy_PythonFunction) AddPermission(id *string, permission *awslambda.Permission) {
	_jsii_.InvokeVoid(
		p,
		"addPermission",
		[]interface{}{id, permission},
	)
}

// Adds a statement to the IAM role assumed by the instance.
// Experimental.
func (p *jsiiProxy_PythonFunction) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PythonFunction) AddVersion(name *string, codeSha256 *string, description *string, provisionedExecutions *float64, asyncInvokeConfig *awslambda.EventInvokeConfigOptions) awslambda.Version {
	var returns awslambda.Version

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Configures options for asynchronous invocation.
// Experimental.
func (p *jsiiProxy_PythonFunction) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	_jsii_.InvokeVoid(
		p,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

// Experimental.
func (p *jsiiProxy_PythonFunction) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grant the given identity permissions to invoke this Lambda.
// Experimental.
func (p *jsiiProxy_PythonFunction) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Function.
// Experimental.
func (p *jsiiProxy_PythonFunction) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_PythonFunction) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PythonFunction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_PythonFunction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_PythonFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonFunction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a PythonFunction.
//
// TODO: EXAMPLE
//
// Experimental.
type PythonFunctionProps struct {
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
	// Path to the source of the function or the location for dependencies.
	// Experimental.
	Entry *string `json:"entry"`
	// The runtime environment.
	//
	// Only runtimes of the Python family are
	// supported.
	// Experimental.
	Runtime awslambda.Runtime `json:"runtime"`
	// Bundling options to use for this function.
	//
	// Use this to specify custom bundling options like
	// the bundling Docker image, asset hash type, custom hash, architecture, etc.
	// Experimental.
	Bundling *BundlingOptions `json:"bundling"`
	// The name of the exported handler in the index file.
	// Experimental.
	Handler *string `json:"handler"`
	// The path (relative to entry) to the index file containing the exported handler.
	// Experimental.
	Index *string `json:"index"`
}

// A lambda layer version.
//
// TODO: EXAMPLE
//
// Experimental.
type PythonLayerVersion interface {
	awslambda.LayerVersion
	CompatibleRuntimes() *[]awslambda.Runtime
	Env() *awscdk.ResourceEnvironment
	LayerVersionArn() *string
	Node() awscdk.ConstructNode
	PhysicalName() *string
	Stack() awscdk.Stack
	AddPermission(id *string, permission *awslambda.LayerVersionPermission)
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

// The jsii proxy struct for PythonLayerVersion
type jsiiProxy_PythonLayerVersion struct {
	internal.Type__awslambdaLayerVersion
}

func (j *jsiiProxy_PythonLayerVersion) CompatibleRuntimes() *[]awslambda.Runtime {
	var returns *[]awslambda.Runtime
	_jsii_.Get(
		j,
		"compatibleRuntimes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonLayerVersion) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonLayerVersion(scope awscdk.Construct, id *string, props *PythonLayerVersionProps) PythonLayerVersion {
	_init_.Initialize()

	j := jsiiProxy_PythonLayerVersion{}

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonLayerVersion_Override(p PythonLayerVersion, scope awscdk.Construct, id *string, props *PythonLayerVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		[]interface{}{scope, id, props},
		p,
	)
}

// Imports a layer version by ARN.
//
// Assumes it is compatible with all Lambda runtimes.
// Experimental.
func PythonLayerVersion_FromLayerVersionArn(scope constructs.Construct, id *string, layerVersionArn *string) awslambda.ILayerVersion {
	_init_.Initialize()

	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		"fromLayerVersionArn",
		[]interface{}{scope, id, layerVersionArn},
		&returns,
	)

	return returns
}

// Imports a Layer that has been defined externally.
// Experimental.
func PythonLayerVersion_FromLayerVersionAttributes(scope constructs.Construct, id *string, attrs *awslambda.LayerVersionAttributes) awslambda.ILayerVersion {
	_init_.Initialize()

	var returns awslambda.ILayerVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		"fromLayerVersionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func PythonLayerVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func PythonLayerVersion_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda_python.PythonLayerVersion",
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
// Experimental.
func (p *jsiiProxy_PythonLayerVersion) AddPermission(id *string, permission *awslambda.LayerVersionPermission) {
	_jsii_.InvokeVoid(
		p,
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
// Experimental.
func (p *jsiiProxy_PythonLayerVersion) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (p *jsiiProxy_PythonLayerVersion) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonLayerVersion) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonLayerVersion) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonLayerVersion) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_PythonLayerVersion) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PythonLayerVersion) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonLayerVersion) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_PythonLayerVersion) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_PythonLayerVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
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
func (p *jsiiProxy_PythonLayerVersion) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for PythonLayerVersion.
//
// TODO: EXAMPLE
//
// Experimental.
type PythonLayerVersionProps struct {
	// The description the this Lambda Layer.
	// Experimental.
	Description *string `json:"description"`
	// The name of the layer.
	// Experimental.
	LayerVersionName *string `json:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	// Experimental.
	License *string `json:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// The path to the root directory of the lambda layer.
	// Experimental.
	Entry *string `json:"entry"`
	// Bundling options to use for this function.
	//
	// Use this to specify custom bundling options like
	// the bundling Docker image, asset hash type, custom hash, architecture, etc.
	// Experimental.
	Bundling *BundlingOptions `json:"bundling"`
	// The system architectures compatible with this layer.
	// Experimental.
	CompatibleArchitectures *[]awslambda.Architecture `json:"compatibleArchitectures"`
	// The runtimes compatible with the python layer.
	// Experimental.
	CompatibleRuntimes *[]awslambda.Runtime `json:"compatibleRuntimes"`
}

