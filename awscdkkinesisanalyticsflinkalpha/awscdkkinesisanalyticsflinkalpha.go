// A CDK Construct Library for Kinesis Analytics Flink applications
package awscdkkinesisanalyticsflinkalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisanalytics"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdkkinesisanalyticsflinkalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The L2 construct for Flink Kinesis Data Applications.
// Experimental.
type Application interface {
	awscdk.Resource
	IApplication
	ApplicationArn() *string
	ApplicationName() *string
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() constructs.Node
	PhysicalName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddToRolePolicy(policyStatement awsiam.PolicyStatement) *bool
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Application
type jsiiProxy_Application struct {
	internal.Type__awscdkResource
	jsiiProxy_IApplication
}

func (j *jsiiProxy_Application) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Application) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewApplication(scope constructs.Construct, id *string, props *ApplicationProps) Application {
	_init_.Initialize()

	j := jsiiProxy_Application{}

	_jsii_.Create(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Application",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplication_Override(a Application, scope constructs.Construct, id *string, props *ApplicationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Application",
		[]interface{}{scope, id, props},
		a,
	)
}

// Import an existing application defined outside of CDK code by applicationArn.
// Experimental.
func Application_FromApplicationArn(scope constructs.Construct, id *string, applicationArn *string) IApplication {
	_init_.Initialize()

	var returns IApplication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Application",
		"fromApplicationArn",
		[]interface{}{scope, id, applicationArn},
		&returns,
	)

	return returns
}

// Import an existing Flink application defined outside of CDK code by applicationName.
// Experimental.
func Application_FromApplicationName(scope constructs.Construct, id *string, applicationName *string) IApplication {
	_init_.Initialize()

	var returns IApplication

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Application",
		"fromApplicationName",
		[]interface{}{scope, id, applicationName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Application_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Application",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Application_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Application",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Implement the convenience {@link IApplication.addToPrincipalPolicy} method.
// Experimental.
func (a *jsiiProxy_Application) AddToRolePolicy(policyStatement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"addToRolePolicy",
		[]interface{}{policyStatement},
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
func (a *jsiiProxy_Application) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_Application) GeneratePhysicalName() *string {
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
func (a *jsiiProxy_Application) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (a *jsiiProxy_Application) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_Application) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Code configuration providing the location to a Flink application JAR file.
// Experimental.
type ApplicationCode interface {
	Bind(scope constructs.Construct) *ApplicationCodeConfig
}

// The jsii proxy struct for ApplicationCode
type jsiiProxy_ApplicationCode struct {
	_ byte // padding
}

// Experimental.
func NewApplicationCode_Override(a ApplicationCode) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCode",
		nil, // no parameters
		a,
	)
}

// Reference code from a local directory containing a Flink JAR file.
// Experimental.
func ApplicationCode_FromAsset(path *string, options *awss3assets.AssetOptions) ApplicationCode {
	_init_.Initialize()

	var returns ApplicationCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference code from an S3 bucket.
// Experimental.
func ApplicationCode_FromBucket(bucket awss3.IBucket, fileKey *string, objectVersion *string) ApplicationCode {
	_init_.Initialize()

	var returns ApplicationCode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.ApplicationCode",
		"fromBucket",
		[]interface{}{bucket, fileKey, objectVersion},
		&returns,
	)

	return returns
}

// A method to lazily bind asset resources to the parent FlinkApplication.
// Experimental.
func (a *jsiiProxy_ApplicationCode) Bind(scope constructs.Construct) *ApplicationCodeConfig {
	var returns *ApplicationCodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// The return type of {@link ApplicationCode.bind}. This represents CloudFormation configuration and an s3 bucket holding the Flink application JAR file.
// Experimental.
type ApplicationCodeConfig struct {
	// Low-level Cloudformation ApplicationConfigurationProperty.
	// Experimental.
	ApplicationCodeConfigurationProperty *awskinesisanalytics.CfnApplicationV2_ApplicationConfigurationProperty `json:"applicationCodeConfigurationProperty"`
	// S3 Bucket that stores the Flink application code.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket"`
}

// Props for creating an Application construct.
// Experimental.
type ApplicationProps struct {
	// The Flink code asset to run.
	// Experimental.
	Code ApplicationCode `json:"code"`
	// The Flink version to use for this application.
	// Experimental.
	Runtime Runtime `json:"runtime"`
	// A name for your Application that is unique to an AWS account.
	// Experimental.
	ApplicationName *string `json:"applicationName"`
	// Whether the Kinesis Data Analytics service can increase the parallelism of the application in response to resource usage.
	// Experimental.
	AutoScalingEnabled *bool `json:"autoScalingEnabled"`
	// Whether checkpointing is enabled while your application runs.
	// Experimental.
	CheckpointingEnabled *bool `json:"checkpointingEnabled"`
	// The interval between checkpoints.
	// Experimental.
	CheckpointInterval awscdk.Duration `json:"checkpointInterval"`
	// The log group to send log entries to.
	// Experimental.
	LogGroup awslogs.ILogGroup `json:"logGroup"`
	// The level of log verbosity from the Flink application.
	// Experimental.
	LogLevel LogLevel `json:"logLevel"`
	// Describes the granularity of the CloudWatch metrics for an application.
	//
	// Use caution with Parallelism level metrics. Parallelism granularity logs
	// metrics for each parallel thread and can quickly become expensive when
	// parallelism is high (e.g. > 64).
	// Experimental.
	MetricsLevel MetricsLevel `json:"metricsLevel"`
	// The minimum amount of time in to wait after a checkpoint finishes to start a new checkpoint.
	// Experimental.
	MinPauseBetweenCheckpoints awscdk.Duration `json:"minPauseBetweenCheckpoints"`
	// The initial parallelism for the application.
	//
	// Kinesis Data Analytics can
	// stop the app, increase the parallelism, and start the app again if
	// autoScalingEnabled is true (the default value).
	// Experimental.
	Parallelism *float64 `json:"parallelism"`
	// The Flink parallelism allowed per Kinesis Processing Unit (KPU).
	// Experimental.
	ParallelismPerKpu *float64 `json:"parallelismPerKpu"`
	// Configuration PropertyGroups.
	//
	// You can use these property groups to pass
	// arbitrary runtime configuration values to your Flink app.
	// Experimental.
	PropertyGroups *PropertyGroups `json:"propertyGroups"`
	// Provide a RemovalPolicy to override the default.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// A role to use to grant permissions to your application.
	//
	// Prefer omitting
	// this property and using the default role.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Determines if Flink snapshots are enabled.
	// Experimental.
	SnapshotsEnabled *bool `json:"snapshotsEnabled"`
}

// An interface expressing the public properties on both an imported and CDK-created Flink application.
// Experimental.
type IApplication interface {
	awsiam.IGrantable
	awscdk.IResource
	// Convenience method for adding a policy statement to the application role.
	// Experimental.
	AddToRolePolicy(policyStatement awsiam.PolicyStatement) *bool
	// The application ARN.
	// Experimental.
	ApplicationArn() *string
	// The name of the Flink application.
	// Experimental.
	ApplicationName() *string
	// The application IAM role.
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IApplication
type jsiiProxy_IApplication struct {
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApplication) AddToRolePolicy(policyStatement awsiam.PolicyStatement) *bool {
	var returns *bool

	_jsii_.Invoke(
		i,
		"addToRolePolicy",
		[]interface{}{policyStatement},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IApplication) ApplicationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) ApplicationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplication) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// Available log levels for Flink applications.
// Experimental.
type LogLevel string

const (
	LogLevel_DEBUG LogLevel = "DEBUG"
	LogLevel_INFO LogLevel = "INFO"
	LogLevel_WARN LogLevel = "WARN"
	LogLevel_ERROR LogLevel = "ERROR"
)

// Granularity of metrics sent to CloudWatch.
// Experimental.
type MetricsLevel string

const (
	MetricsLevel_APPLICATION MetricsLevel = "APPLICATION"
	MetricsLevel_TASK MetricsLevel = "TASK"
	MetricsLevel_OPERATOR MetricsLevel = "OPERATOR"
	MetricsLevel_PARALLELISM MetricsLevel = "PARALLELISM"
)

// Interface for building AWS::KinesisAnalyticsV2::Application PropertyGroup configuration.
// Experimental.
type PropertyGroups struct {
}

// Available Flink runtimes for Kinesis Analytics.
// Experimental.
type Runtime interface {
	Value() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Create a new Runtime with with an arbitrary Flink version string.
// Experimental.
func Runtime_Of(value *string) Runtime {
	_init_.Initialize()

	var returns Runtime

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Runtime_FLINK_1_11() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_11",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_13() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_13",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_6",
		&returns,
	)
	return returns
}

func Runtime_FLINK_1_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"@aws-cdk/aws-kinesisanalytics-flink-alpha.Runtime",
		"FLINK_1_8",
		&returns,
	)
	return returns
}

