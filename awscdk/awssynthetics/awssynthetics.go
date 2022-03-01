package awssynthetics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/awssynthetics/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
		"monocdk.aws_synthetics.AssetCode",
		[]interface{}{assetPath, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetCode_Override(a AssetCode, assetPath *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_synthetics.AssetCode",
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
		"monocdk.aws_synthetics.AssetCode",
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
		"monocdk.aws_synthetics.AssetCode",
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
		"monocdk.aws_synthetics.AssetCode",
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
	Node() awscdk.ConstructNode
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_Canary) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.aws_synthetics.Canary",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCanary_Override(c Canary, scope constructs.Construct, id *string, props *CanaryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_synthetics.Canary",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Canary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_synthetics.Canary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Canary_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_synthetics.Canary",
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_Canary) OnPrepare() {
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
func (c *jsiiProxy_Canary) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_Canary) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
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
func (c *jsiiProxy_Canary) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_Canary) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_Canary) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
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

// A CloudFormation `AWS::Synthetics::Canary`.
//
// Creates or updates a canary. Canaries are scripts that monitor your endpoints and APIs from the outside-in. Canaries help you check the availability and latency of your web services and troubleshoot anomalies by investigating load time data, screenshots of the UI, logs, and metrics. You can set up a canary to run continuously or just once.
//
// To create canaries, you must have the `CloudWatchSyntheticsFullAccess` policy. If you are creating a new IAM role for the canary, you also need the the `iam:CreateRole` , `iam:CreatePolicy` and `iam:AttachRolePolicy` permissions. For more information, see [Necessary Roles and Permissions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Roles) .
//
// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the Amazon Resource Name (ARN) for the canary, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
//
// TODO: EXAMPLE
//
type CfnCanary interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ArtifactConfig() interface{}
	SetArtifactConfig(val interface{})
	ArtifactS3Location() *string
	SetArtifactS3Location(val *string)
	AttrId() *string
	AttrState() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Code() interface{}
	SetCode(val interface{})
	CreationStack() *[]*string
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	FailureRetentionPeriod() *float64
	SetFailureRetentionPeriod(val *float64)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() awscdk.ConstructNode
	Ref() *string
	RunConfig() interface{}
	SetRunConfig(val interface{})
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	Schedule() interface{}
	SetSchedule(val interface{})
	Stack() awscdk.Stack
	StartCanaryAfterCreation() interface{}
	SetStartCanaryAfterCreation(val interface{})
	SuccessRetentionPeriod() *float64
	SetSuccessRetentionPeriod(val *float64)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VisualReference() interface{}
	SetVisualReference(val interface{})
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

// The jsii proxy struct for CfnCanary
type jsiiProxy_CfnCanary struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCanary) ArtifactConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifactConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) ArtifactS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) AttrState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Code() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) FailureRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) RunConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) StartCanaryAfterCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryAfterCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) SuccessRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) VisualReference() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visualReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCanary) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::Synthetics::Canary`.
func NewCfnCanary(scope awscdk.Construct, id *string, props *CfnCanaryProps) CfnCanary {
	_init_.Initialize()

	j := jsiiProxy_CfnCanary{}

	_jsii_.Create(
		"monocdk.aws_synthetics.CfnCanary",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Synthetics::Canary`.
func NewCfnCanary_Override(c CfnCanary, scope awscdk.Construct, id *string, props *CfnCanaryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_synthetics.CfnCanary",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCanary) SetArtifactConfig(val interface{}) {
	_jsii_.Set(
		j,
		"artifactConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetArtifactS3Location(val *string) {
	_jsii_.Set(
		j,
		"artifactS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetCode(val interface{}) {
	_jsii_.Set(
		j,
		"code",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetFailureRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"failureRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetRunConfig(val interface{}) {
	_jsii_.Set(
		j,
		"runConfig",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetRuntimeVersion(val *string) {
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetSchedule(val interface{}) {
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetStartCanaryAfterCreation(val interface{}) {
	_jsii_.Set(
		j,
		"startCanaryAfterCreation",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetSuccessRetentionPeriod(val *float64) {
	_jsii_.Set(
		j,
		"successRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetVisualReference(val interface{}) {
	_jsii_.Set(
		j,
		"visualReference",
		val,
	)
}

func (j *jsiiProxy_CfnCanary) SetVpcConfig(val interface{}) {
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
// Experimental.
func CfnCanary_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_synthetics.CfnCanary",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCanary_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_synthetics.CfnCanary",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnCanary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_synthetics.CfnCanary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCanary_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.CfnCanary",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCanary) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnCanary) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnCanary) AddMetadata(key *string, value interface{}) {
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
// Experimental.
func (c *jsiiProxy_CfnCanary) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCanary) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnCanary) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCanary) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnCanary) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnCanary) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCanary) Inspect(inspector awscdk.TreeInspector) {
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
func (c *jsiiProxy_CfnCanary) OnPrepare() {
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
func (c *jsiiProxy_CfnCanary) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CfnCanary) OnValidate() *[]*string {
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
func (c *jsiiProxy_CfnCanary) OverrideLogicalId(newLogicalId *string) {
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
func (c *jsiiProxy_CfnCanary) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnCanary) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnCanary) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnCanary) Synthesize(session awscdk.ISynthesisSession) {
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
func (c *jsiiProxy_CfnCanary) ToString() *string {
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
func (c *jsiiProxy_CfnCanary) Validate() *[]*string {
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
func (c *jsiiProxy_CfnCanary) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
//
// TODO: EXAMPLE
//
type CfnCanary_ArtifactConfigProperty struct {
	// A structure that contains the configuration of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
	//
	// Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3 or later. For more information, see [Encrypting canary artifacts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_artifact_encryption.html) .
	S3Encryption interface{} `json:"s3Encryption" yaml:"s3Encryption"`
}

// A structure representing a screenshot that is used as a baseline during visual monitoring comparisons made by the canary.
//
// TODO: EXAMPLE
//
type CfnCanary_BaseScreenshotProperty struct {
	// The name of the screenshot.
	//
	// This is generated the first time the canary is run after the `UpdateCanary` operation that specified for this canary to perform visual monitoring.
	ScreenshotName *string `json:"screenshotName" yaml:"screenshotName"`
	// Coordinates that define the part of a screen to ignore during screenshot comparisons.
	//
	// To obtain the coordinates to use here, use the CloudWatch Logs console to draw the boundaries on the screen. For more information, see {LINK}
	IgnoreCoordinates *[]*string `json:"ignoreCoordinates" yaml:"ignoreCoordinates"`
}

// Use this structure to input your script code for the canary.
//
// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
//
// TODO: EXAMPLE
//
type CfnCanary_CodeProperty struct {
	// The entry point to use for the source code when running the canary.
	//
	// This value must end with the string `.handler` . The string is limited to 29 characters or fewer.
	Handler *string `json:"handler" yaml:"handler"`
	// If your canary script is located in S3, specify the bucket name here.
	//
	// The bucket must already exist.
	S3Bucket *string `json:"s3Bucket" yaml:"s3Bucket"`
	// The S3 key of your script.
	//
	// For more information, see [Working with Amazon S3 Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingObjects.html) .
	S3Key *string `json:"s3Key" yaml:"s3Key"`
	// The S3 version ID of your script.
	S3ObjectVersion *string `json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
	// If you input your canary script directly into the canary instead of referring to an S3 location, the value of this parameter is the script in plain text.
	//
	// It can be up to 5 MB.
	Script *string `json:"script" yaml:"script"`
}

// A structure that contains input information for a canary run.
//
// This structure is required.
//
// TODO: EXAMPLE
//
type CfnCanary_RunConfigProperty struct {
	// Specifies whether this canary is to use active AWS X-Ray tracing when it runs.
	//
	// Active tracing enables this canary run to be displayed in the ServiceLens and X-Ray service maps even if the canary does not hit an endpoint that has X-Ray tracing enabled. Using X-Ray tracing incurs charges. For more information, see [Canaries and X-Ray tracing](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_tracing.html) .
	//
	// You can enable active tracing only for canaries that use version `syn-nodejs-2.0` or later for their canary runtime.
	ActiveTracing interface{} `json:"activeTracing" yaml:"activeTracing"`
	// Specifies the keys and values to use for any environment variables used in the canary script.
	//
	// Use the following format:
	//
	// { "key1" : "value1", "key2" : "value2", ...}
	//
	// Keys must start with a letter and be at least two characters. The total size of your environment variables cannot exceed 4 KB. You can't specify any Lambda reserved environment variables as the keys for your environment variables. For more information about reserved keys, see [Runtime environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime) .
	EnvironmentVariables interface{} `json:"environmentVariables" yaml:"environmentVariables"`
	// The maximum amount of memory that the canary can use while running.
	//
	// This value must be a multiple of 64. The range is 960 to 3008.
	MemoryInMb *float64 `json:"memoryInMb" yaml:"memoryInMb"`
	// How long the canary is allowed to run before it must stop.
	//
	// You can't set this time to be longer than the frequency of the runs of this canary.
	//
	// If you omit this field, the frequency of the canary is used as this value, up to a maximum of 900 seconds.
	TimeoutInSeconds *float64 `json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}

// A structure that contains the configuration of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3 .
//
// Artifact encryption functionality is available only for canaries that use Synthetics runtime version syn-nodejs-puppeteer-3.3 or later. For more information, see [Encrypting canary artifacts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_artifact_encryption.html) .
//
// TODO: EXAMPLE
//
type CfnCanary_S3EncryptionProperty struct {
	// The encryption method to use for artifacts created by this canary.
	//
	// Specify `SSE_S3` to use server-side encryption (SSE) with an Amazon S3-managed key. Specify `SSE-KMS` to use server-side encryption with a customer-managed AWS KMS key.
	//
	// If you omit this parameter, an AWS -managed AWS KMS key is used.
	EncryptionMode *string `json:"encryptionMode" yaml:"encryptionMode"`
	// The ARN of the customer-managed AWS KMS key to use, if you specify `SSE-KMS` for `EncryptionMode`.
	KmsKeyArn *string `json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

// This structure specifies how often a canary is to make runs and the date and time when it should stop making runs.
//
// TODO: EXAMPLE
//
type CfnCanary_ScheduleProperty struct {
	// A `rate` expression or a `cron` expression that defines how often the canary is to run.
	//
	// For a rate expression, The syntax is `rate( *number unit* )` . *unit* can be `minute` , `minutes` , or `hour` .
	//
	// For example, `rate(1 minute)` runs the canary once a minute, `rate(10 minutes)` runs it once every 10 minutes, and `rate(1 hour)` runs it once every hour. You can specify a frequency between `rate(1 minute)` and `rate(1 hour)` .
	//
	// Specifying `rate(0 minute)` or `rate(0 hour)` is a special value that causes the canary to run only once when it is started.
	//
	// Use `cron( *expression* )` to specify a cron expression. You can't schedule a canary to wait for more than a year before running. For information about the syntax for cron expressions, see [Scheduling canary runs using cron](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_cron.html) .
	Expression *string `json:"expression" yaml:"expression"`
	// How long, in seconds, for the canary to continue making regular runs according to the schedule in the `Expression` value.
	//
	// If you specify 0, the canary continues making runs until you stop it. If you omit this field, the default of 0 is used.
	DurationInSeconds *string `json:"durationInSeconds" yaml:"durationInSeconds"`
}

// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
//
// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
//
// TODO: EXAMPLE
//
type CfnCanary_VPCConfigProperty struct {
	// The IDs of the security groups for this canary.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the subnets where this canary is to run.
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC where this canary is to run.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Defines the screenshots to use as the baseline for comparisons during visual monitoring comparisons during future runs of this canary.
//
// If you omit this parameter, no changes are made to any baseline screenshots that the canary might be using already.
//
// Visual monitoring is supported only on canaries running the *syn-puppeteer-node-3.2* runtime or later. For more information, see [Visual monitoring](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Library_SyntheticsLogger_VisualTesting.html) and [Visual monitoring blueprint](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Blueprints_VisualTesting.html)
//
// TODO: EXAMPLE
//
type CfnCanary_VisualReferenceProperty struct {
	// Specifies which canary run to use the screenshots from as the baseline for future visual monitoring with this canary.
	//
	// Valid values are `nextrun` to use the screenshots from the next run after this update is made, `lastrun` to use the screenshots from the most recent run before this update was made, or the value of `Id` in the [CanaryRun](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CanaryRun.html) from any past run of this canary.
	BaseCanaryRunId *string `json:"baseCanaryRunId" yaml:"baseCanaryRunId"`
	// An array of screenshots that are used as the baseline for comparisons during visual monitoring.
	BaseScreenshots interface{} `json:"baseScreenshots" yaml:"baseScreenshots"`
}

// Properties for defining a `CfnCanary`.
//
// TODO: EXAMPLE
//
type CfnCanaryProps struct {
	// The location in Amazon S3 where Synthetics stores artifacts from the runs of this canary.
	//
	// Artifacts include the log file, screenshots, and HAR files. Specify the full location path, including `s3://` at the beginning of the path.
	ArtifactS3Location *string `json:"artifactS3Location" yaml:"artifactS3Location"`
	// Use this structure to input your script code for the canary.
	//
	// This structure contains the Lambda handler with the location where the canary should start running the script. If the script is stored in an S3 bucket, the bucket name, key, and version are also included. If the script is passed into the canary directly, the script code is contained in the value of `Script` .
	Code interface{} `json:"code" yaml:"code"`
	// The ARN of the IAM role to be used to run the canary.
	//
	// This role must already exist, and must include `lambda.amazonaws.com` as a principal in the trust policy. The role must also have the following permissions:
	//
	// - `s3:PutObject`
	// - `s3:GetBucketLocation`
	// - `s3:ListAllMyBuckets`
	// - `cloudwatch:PutMetricData`
	// - `logs:CreateLogGroup`
	// - `logs:CreateLogStream`
	// - `logs:PutLogEvents`
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name for this canary.
	//
	// Be sure to give it a descriptive name that distinguishes it from other canaries in your account.
	//
	// Do not include secrets or proprietary information in your canary names. The canary name makes up part of the canary ARN, and the ARN is included in outbound calls over the internet. For more information, see [Security Considerations for Synthetics Canaries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/servicelens_canaries_security.html) .
	Name *string `json:"name" yaml:"name"`
	// Specifies the runtime version to use for the canary.
	//
	// For more information about runtime versions, see [Canary Runtime Versions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) .
	RuntimeVersion *string `json:"runtimeVersion" yaml:"runtimeVersion"`
	// A structure that contains information about how often the canary is to run, and when these runs are to stop.
	Schedule interface{} `json:"schedule" yaml:"schedule"`
	// Specify TRUE to have the canary start making runs immediately after it is created.
	//
	// A canary that you create using CloudFormation can't be used to monitor the CloudFormation stack that creates the canary or to roll back that stack if there is a failure.
	StartCanaryAfterCreation interface{} `json:"startCanaryAfterCreation" yaml:"startCanaryAfterCreation"`
	// A structure that contains the configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3.
	ArtifactConfig interface{} `json:"artifactConfig" yaml:"artifactConfig"`
	// The number of days to retain data about failed runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod *float64 `json:"failureRetentionPeriod" yaml:"failureRetentionPeriod"`
	// A structure that contains input information for a canary run.
	//
	// If you omit this structure, the frequency of the canary is used as canary's timeout value, up to a maximum of 900 seconds.
	RunConfig interface{} `json:"runConfig" yaml:"runConfig"`
	// The number of days to retain data about successful runs of this canary.
	//
	// If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod *float64 `json:"successRetentionPeriod" yaml:"successRetentionPeriod"`
	// The list of key-value pairs that are associated with the canary.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// If this canary performs visual monitoring by comparing screenshots, this structure contains the ID of the canary run to use as the baseline for screenshots, and the coordinates of any parts of the screen to ignore during the visual monitoring comparison.
	VisualReference interface{} `json:"visualReference" yaml:"visualReference"`
	// If this canary is to test an endpoint in a VPC, this structure contains information about the subnet and security groups of the VPC endpoint.
	//
	// For more information, see [Running a Canary in a VPC](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html) .
	VpcConfig interface{} `json:"vpcConfig" yaml:"vpcConfig"`
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
		"monocdk.aws_synthetics.Code",
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
		"monocdk.aws_synthetics.Code",
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
		"monocdk.aws_synthetics.Code",
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
		"monocdk.aws_synthetics.Code",
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
		"monocdk.aws_synthetics.InlineCode",
		[]interface{}{code},
		&j,
	)

	return &j
}

// Experimental.
func NewInlineCode_Override(i InlineCode, code *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_synthetics.InlineCode",
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
		"monocdk.aws_synthetics.InlineCode",
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
		"monocdk.aws_synthetics.InlineCode",
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
		"monocdk.aws_synthetics.InlineCode",
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
		"monocdk.aws_synthetics.Runtime",
		[]interface{}{name, family},
		&j,
	)

	return &j
}

// Experimental.
func NewRuntime_Override(r Runtime, name *string, family RuntimeFamily) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_synthetics.Runtime",
		[]interface{}{name, family},
		r,
	)
}

func Runtime_SYNTHETICS_1_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_1_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_2_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_2_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_2_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_2_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_2_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_2_2",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_0",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_1",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_2",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_NODEJS_PUPPETEER_3_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
		"SYNTHETICS_NODEJS_PUPPETEER_3_3",
		&returns,
	)
	return returns
}

func Runtime_SYNTHETICS_PYTHON_SELENIUM_1_0() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"monocdk.aws_synthetics.Runtime",
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
		"monocdk.aws_synthetics.S3Code",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

// Experimental.
func NewS3Code_Override(s S3Code, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_synthetics.S3Code",
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
		"monocdk.aws_synthetics.S3Code",
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
		"monocdk.aws_synthetics.S3Code",
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
		"monocdk.aws_synthetics.S3Code",
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
		"monocdk.aws_synthetics.Schedule",
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
		"monocdk.aws_synthetics.Schedule",
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
		"monocdk.aws_synthetics.Schedule",
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
		"monocdk.aws_synthetics.Schedule",
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
		"monocdk.aws_synthetics.Test",
		"custom",
		[]interface{}{options},
		&returns,
	)

	return returns
}

