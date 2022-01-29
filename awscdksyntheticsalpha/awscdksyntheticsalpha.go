// The CDK Construct Library for AWS::Synthetics
package awscdksyntheticsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksyntheticsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdksyntheticsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for specifying the s3 location that stores the data of each canary run.
//
// The artifacts bucket location **cannot**
// be updated once the canary is created.
//
// TODO: EXAMPLE
//
// Experimental.
type ArtifactsBucketLocation struct {
	// The s3 location that stores the data of each run.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// The S3 bucket prefix.
	//
	// Specify this if you want a more specific path within the artifacts bucket.
	// Experimental.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// Canary code from an Asset.
//
// TODO: EXAMPLE
//
// Experimental.
type AssetCode interface {
	Code
	Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig
}

// The jsii proxy struct for AssetCode
type jsiiProxy_AssetCode struct {
	jsiiProxy_Code
}

// Experimental.
func NewAssetCode(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	j := jsiiProxy_AssetCode{}

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		[]interface{}{assetPath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetCode_Override(a AssetCode, assetPath *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		[]interface{}{assetPath, options},
		a,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func AssetCode_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		"fromAsset",
		[]interface{}{assetPath, options},
		&returns,
	)

	return returns
}

// Specify code from an s3 bucket.
//
// The object in the s3 bucket must be a .zip file that contains
// the structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `S3Code` associated with the specified S3 object.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func AssetCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
// Experimental.
func AssetCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.AssetCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
// Experimental.
func (a *jsiiProxy_AssetCode) Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, handler, family},
		&returns,
	)

	return returns
}

// Define a new Canary.
//
// TODO: EXAMPLE
//
// Experimental.
type Canary interface {
	awscdk.Resource
	ArtifactsBucket() awss3.IBucket
	CanaryId() *string
	CanaryName() *string
	CanaryState() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	MetricDuration(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSuccessPercent(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
}

// The jsii proxy struct for Canary
type jsiiProxy_Canary struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Canary) ArtifactsBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"artifactsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) CanaryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) CanaryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) CanaryState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canaryState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Canary) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewCanary(scope constructs.Construct, id *string, props *CanaryProps) Canary {
	_init_.Initialize()

	j := jsiiProxy_Canary{}

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.Canary",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCanary_Override(c Canary, scope constructs.Construct, id *string, props *CanaryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.Canary",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Canary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Canary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Canary_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Canary",
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
// Experimental.
func (c *jsiiProxy_Canary) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (c *jsiiProxy_Canary) GeneratePhysicalName() *string {
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
// Experimental.
func (c *jsiiProxy_Canary) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
// Experimental.
func (c *jsiiProxy_Canary) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Measure the Duration of a single canary run, in seconds.
// Experimental.
func (c *jsiiProxy_Canary) MetricDuration(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricDuration",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Measure the number of failed canary runs over a given time period.
//
// Default: sum over 5 minutes
// Experimental.
func (c *jsiiProxy_Canary) MetricFailed(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricFailed",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Measure the percentage of successful canary runs.
// Experimental.
func (c *jsiiProxy_Canary) MetricSuccessPercent(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricSuccessPercent",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_Canary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a canary.
//
// TODO: EXAMPLE
//
// Experimental.
type CanaryProps struct {
	// Specify the runtime version to use for the canary.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html
	//
	// Experimental.
	Runtime Runtime `json:"runtime" yaml:"runtime"`
	// The type of test that you want your canary to run.
	//
	// Use `Test.custom()` to specify the test to run.
	// Experimental.
	Test Test `json:"test" yaml:"test"`
	// The s3 location that stores the data of the canary runs.
	// Experimental.
	ArtifactsBucketLocation *ArtifactsBucketLocation `json:"artifactsBucketLocation" yaml:"artifactsBucketLocation"`
	// The name of the canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from
	// other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary name. The canary name
	// makes up part of the canary ARN, which is included in outbound calls over the internet.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html
	//
	// Experimental.
	CanaryName *string `json:"canaryName" yaml:"canaryName"`
	// Key-value pairs that the Synthetics caches and makes available for your canary scripts.
	//
	// Use environment variables
	// to apply configuration changes, such as test and production environment configurations, without changing your
	// Canary script source code.
	// Experimental.
	EnvironmentVariables *map[string]*string `json:"environmentVariables" yaml:"environmentVariables"`
	// How many days should failed runs be retained.
	// Experimental.
	FailureRetentionPeriod awscdk.Duration `json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// Canary execution role.
	//
	// This is the role that will be assumed by the canary upon execution.
	// It controls the permissions that the canary will have. The role must
	// be assumable by the AWS Lambda service principal.
	//
	// If not supplied, a role will be created with all the required permissions.
	// If you provide a Role, you must add the required permissions.
	// See: required permissions: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-synthetics-canary.html#cfn-synthetics-canary-executionrolearn
	//
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Specify the schedule for how often the canary runs.
	//
	// For example, if you set `schedule` to `rate(10 minutes)`, then the canary will run every 10 minutes.
	// You can set the schedule with `Schedule.rate(Duration)` (recommended) or you can specify an expression using `Schedule.expression()`.
	// Experimental.
	Schedule Schedule `json:"schedule" yaml:"schedule"`
	// Whether or not the canary should start after creation.
	// Experimental.
	StartAfterCreation *bool `json:"startAfterCreation" yaml:"startAfterCreation"`
	// How many days should successful runs be retained.
	// Experimental.
	SuccessRetentionPeriod awscdk.Duration `json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// How long the canary will be in a 'RUNNING' state.
	//
	// For example, if you set `timeToLive` to be 1 hour and `schedule` to be `rate(10 minutes)`,
	// your canary will run at 10 minute intervals for an hour, for a total of 6 times.
	// Experimental.
	TimeToLive awscdk.Duration `json:"timeToLive" yaml:"timeToLive"`
}

// The code the canary should execute.
//
// TODO: EXAMPLE
//
// Experimental.
type Code interface {
	Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig
}

// The jsii proxy struct for Code
type jsiiProxy_Code struct {
	_ byte // padding
}

// Experimental.
func NewCode_Override(c Code) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.Code",
		nil, // no parameters
		c,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func Code_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Code",
		"fromAsset",
		[]interface{}{assetPath, options},
		&returns,
	)

	return returns
}

// Specify code from an s3 bucket.
//
// The object in the s3 bucket must be a .zip file that contains
// the structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `S3Code` associated with the specified S3 object.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func Code_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Code",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
// Experimental.
func Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
//
// Returns: a bound `CodeConfig`.
// Experimental.
func (c *jsiiProxy_Code) Bind(scope constructs.Construct, handler *string, family RuntimeFamily) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, handler, family},
		&returns,
	)

	return returns
}

// Configuration of the code class.
//
// TODO: EXAMPLE
//
// Experimental.
type CodeConfig struct {
	// Inline code (mutually exclusive with `s3Location`).
	// Experimental.
	InlineCode *string `json:"inlineCode" yaml:"inlineCode"`
	// The location of the code in S3 (mutually exclusive with `inlineCode`).
	// Experimental.
	S3Location *awss3.Location `json:"s3Location" yaml:"s3Location"`
}

// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_cron.html
//
// Experimental.
type CronOptions struct {
	// The day of the month to run this rule at.
	// Experimental.
	Day *string `json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Experimental.
	Hour *string `json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	// Experimental.
	Minute *string `json:"minute" yaml:"minute"`
	// The month to run this rule at.
	// Experimental.
	Month *string `json:"month" yaml:"month"`
	// The day of the week to run this rule at.
	// Experimental.
	WeekDay *string `json:"weekDay" yaml:"weekDay"`
}

// Properties for specifying a test.
//
// TODO: EXAMPLE
//
// Experimental.
type CustomTestOptions struct {
	// The code of the canary script.
	// Experimental.
	Code Code `json:"code" yaml:"code"`
	// The handler for the code.
	//
	// Must end with `.handler`.
	// Experimental.
	Handler *string `json:"handler" yaml:"handler"`
}

// Canary code from an inline string.
//
// TODO: EXAMPLE
//
// Experimental.
type InlineCode interface {
	Code
	Bind(_scope constructs.Construct, handler *string, _family RuntimeFamily) *CodeConfig
}

// The jsii proxy struct for InlineCode
type jsiiProxy_InlineCode struct {
	jsiiProxy_Code
}

// Experimental.
func NewInlineCode(code *string) InlineCode {
	_init_.Initialize()

	j := jsiiProxy_InlineCode{}

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		[]interface{}{code},
		&j,
	)

	return &j
}

// Experimental.
func NewInlineCode_Override(i InlineCode, code *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		[]interface{}{code},
		i,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func InlineCode_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		"fromAsset",
		[]interface{}{assetPath, options},
		&returns,
	)

	return returns
}

// Specify code from an s3 bucket.
//
// The object in the s3 bucket must be a .zip file that contains
// the structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `S3Code` associated with the specified S3 object.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func InlineCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
// Experimental.
func InlineCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.InlineCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
// Experimental.
func (i *jsiiProxy_InlineCode) Bind(_scope constructs.Construct, handler *string, _family RuntimeFamily) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope, handler, _family},
		&returns,
	)

	return returns
}

// Runtime options for a canary.
//
// TODO: EXAMPLE
//
// Experimental.
type Runtime interface {
	Family() RuntimeFamily
	Name() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
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


// Experimental.
func NewRuntime(name *string, family RuntimeFamily) Runtime {
	_init_.Initialize()

	j := jsiiProxy_Runtime{}

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		[]interface{}{name, family},
		&j,
	)

	return &j
}

// Experimental.
func NewRuntime_Override(r Runtime, name *string, family RuntimeFamily) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		[]interface{}{name, family},
		r,
	)
}

func Runtime_SYNTHETICS_1_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_1_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_2_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_NODEJS_2_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_2_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_NODEJS_2_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_2_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_NODEJS_2_2",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_2",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_3",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_1_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-synthetics-alpha.Runtime",
		"SYNTHETICS_PYTHON_SELENIUM_1_0",
		&returns,
	)
	return returns
}

// All known Lambda runtime families.
// Experimental.
type RuntimeFamily string

const (
	RuntimeFamily_NODEJS RuntimeFamily = "NODEJS"
	RuntimeFamily_PYTHON RuntimeFamily = "PYTHON"
	RuntimeFamily_OTHER RuntimeFamily = "OTHER"
)

// S3 bucket path to the code zip file.
//
// TODO: EXAMPLE
//
// Experimental.
type S3Code interface {
	Code
	Bind(_scope constructs.Construct, _handler *string, _family RuntimeFamily) *CodeConfig
}

// The jsii proxy struct for S3Code
type jsiiProxy_S3Code struct {
	jsiiProxy_Code
}

// Experimental.
func NewS3Code(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	j := jsiiProxy_S3Code{}

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.S3Code",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Code_Override(s S3Code, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-synthetics-alpha.S3Code",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Specify code from a local path.
//
// Path must include the folder structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `AssetCode` associated with the specified path.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func S3Code_FromAsset(assetPath *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	var returns AssetCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.S3Code",
		"fromAsset",
		[]interface{}{assetPath, options},
		&returns,
	)

	return returns
}

// Specify code from an s3 bucket.
//
// The object in the s3 bucket must be a .zip file that contains
// the structure `nodejs/node_modules/myCanaryFilename.js`.
//
// Returns: `S3Code` associated with the specified S3 object.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_WritingCanary.html#CloudWatch_Synthetics_Canaries_write_from_scratch
//
// Experimental.
func S3Code_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	var returns S3Code

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.S3Code",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Specify code inline.
//
// Returns: `InlineCode` with inline code.
// Experimental.
func S3Code_FromInline(code *string) InlineCode {
	_init_.Initialize()

	var returns InlineCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.S3Code",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

// Called when the canary is initialized to allow this object to bind to the stack, add resources and have fun.
// Experimental.
func (s *jsiiProxy_S3Code) Bind(_scope constructs.Construct, _handler *string, _family RuntimeFamily) *CodeConfig {
	var returns *CodeConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _handler, _family},
		&returns,
	)

	return returns
}

// Schedule for canary runs.
//
// TODO: EXAMPLE
//
// Experimental.
type Schedule interface {
	ExpressionString() *string
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


// Create a schedule from a set of cron fields.
// Experimental.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
//
// The expression must be in a `rate(number units)` format.
// For example, `Schedule.expression('rate(10 minutes)')`
// Experimental.
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// The canary will be executed once.
// Experimental.
func Schedule_Once() Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Schedule",
		"once",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construct a schedule from an interval.
//
// Allowed values: 0 (for a single run) or between 1 and 60 minutes.
// To specify a single run, you can use `Schedule.once()`.
// Experimental.
func Schedule_Rate(interval awscdk.Duration) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Schedule",
		"rate",
		[]interface{}{interval},
		&returns,
	)

	return returns
}

// Specify a test that the canary should run.
//
// TODO: EXAMPLE
//
// Experimental.
type Test interface {
	Code() Code
	Handler() *string
}

// The jsii proxy struct for Test
type jsiiProxy_Test struct {
	_ byte // padding
}

func (j *jsiiProxy_Test) Code() Code {
	var returns Code
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Test) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}


// Specify a custom test with your own code.
//
// Returns: `Test` associated with the specified Code object
// Experimental.
func Test_Custom(options *CustomTestOptions) Test {
	_init_.Initialize()

	var returns Test

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-synthetics-alpha.Test",
		"custom",
		[]interface{}{options},
		&returns,
	)

	return returns
}

