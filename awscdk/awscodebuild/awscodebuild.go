package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Artifacts definition for a CodeBuild Project.
//
// TODO: EXAMPLE
//
type Artifacts interface {
	IArtifacts
	Identifier() *string
	Type() *string
	Bind(_scope constructs.Construct, _project IProject) *ArtifactsConfig
}

// The jsii proxy struct for Artifacts
type jsiiProxy_Artifacts struct {
	jsiiProxy_IArtifacts
}

func (j *jsiiProxy_Artifacts) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifacts) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewArtifacts_Override(a Artifacts, props *ArtifactsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Artifacts",
		[]interface{}{props},
		a,
	)
}

func Artifacts_S3(props *S3ArtifactsProps) IArtifacts {
	_init_.Initialize()

	var returns IArtifacts

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Artifacts",
		"s3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Callback when an Artifacts class is used in a CodeBuild Project.
func (a *jsiiProxy_Artifacts) Bind(_scope constructs.Construct, _project IProject) *ArtifactsConfig {
	var returns *ArtifactsConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_scope, _project},
		&returns,
	)

	return returns
}

// The type returned from {@link IArtifacts#bind}.
//
// TODO: EXAMPLE
//
type ArtifactsConfig struct {
	// The low-level CloudFormation artifacts property.
	ArtifactsProperty *CfnProject_ArtifactsProperty `json:"artifactsProperty" yaml:"artifactsProperty"`
}

// Properties common to all Artifacts classes.
//
// TODO: EXAMPLE
//
type ArtifactsProps struct {
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	Identifier *string `json:"identifier" yaml:"identifier"`
}

// The type returned from {@link IProject#enableBatchBuilds}.
//
// TODO: EXAMPLE
//
type BatchBuildConfig struct {
	// The IAM batch service Role of this Project.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// The extra options passed to the {@link IProject.bindToCodePipeline} method.
//
// TODO: EXAMPLE
//
type BindToCodePipelineOptions struct {
	// The artifact bucket that will be used by the action that invokes this project.
	ArtifactBucket awss3.IBucket `json:"artifactBucket" yaml:"artifactBucket"`
}

// The source credentials used when contacting the BitBucket API.
//
// **Note**: CodeBuild only allows a single credential for BitBucket
// to be saved in a given AWS account in a given region -
// any attempt to add more than one will result in an error.
//
// TODO: EXAMPLE
//
type BitBucketSourceCredentials interface {
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

// The jsii proxy struct for BitBucketSourceCredentials
type jsiiProxy_BitBucketSourceCredentials struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_BitBucketSourceCredentials) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BitBucketSourceCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BitBucketSourceCredentials) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BitBucketSourceCredentials) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewBitBucketSourceCredentials(scope constructs.Construct, id *string, props *BitBucketSourceCredentialsProps) BitBucketSourceCredentials {
	_init_.Initialize()

	j := jsiiProxy_BitBucketSourceCredentials{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.BitBucketSourceCredentials",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBitBucketSourceCredentials_Override(b BitBucketSourceCredentials, scope constructs.Construct, id *string, props *BitBucketSourceCredentialsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.BitBucketSourceCredentials",
		[]interface{}{scope, id, props},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func BitBucketSourceCredentials_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BitBucketSourceCredentials",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BitBucketSourceCredentials_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BitBucketSourceCredentials",
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
func (b *jsiiProxy_BitBucketSourceCredentials) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BitBucketSourceCredentials) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BitBucketSourceCredentials) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		b,
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
func (b *jsiiProxy_BitBucketSourceCredentials) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (b *jsiiProxy_BitBucketSourceCredentials) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties of {@link BitBucketSourceCredentials}.
//
// TODO: EXAMPLE
//
type BitBucketSourceCredentialsProps struct {
	// Your BitBucket application password.
	Password awscdk.SecretValue `json:"password" yaml:"password"`
	// Your BitBucket username.
	Username awscdk.SecretValue `json:"username" yaml:"username"`
}

// Construction properties for {@link BitBucketSource}.
//
// TODO: EXAMPLE
//
type BitBucketSourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// The BitBucket account/user that owns the repo.
	//
	// TODO: EXAMPLE
	//
	Owner *string `json:"owner" yaml:"owner"`
	// The name of the repo (without the username).
	//
	// TODO: EXAMPLE
	//
	Repo *string `json:"repo" yaml:"repo"`
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build.
	//
	// TODO: EXAMPLE
	//
	BranchOrRef *string `json:"branchOrRef" yaml:"branchOrRef"`
	// The depth of history to download.
	//
	// Minimum value is 0.
	// If this value is 0, greater than 25, or not provided,
	// then the full history is downloaded with each build of the project.
	CloneDepth *float64 `json:"cloneDepth" yaml:"cloneDepth"`
	// Whether to fetch submodules while cloning git repo.
	FetchSubmodules *bool `json:"fetchSubmodules" yaml:"fetchSubmodules"`
	// Whether to send notifications on your build's start and end.
	ReportBuildStatus *bool `json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// Whether to create a webhook that will trigger a build every time an event happens in the repository.
	Webhook *bool `json:"webhook" yaml:"webhook"`
	// A list of webhook filters that can constraint what events in the repository will trigger a build.
	//
	// A build is triggered if any of the provided filter groups match.
	// Only valid if `webhook` was not provided as false.
	WebhookFilters *[]FilterGroup `json:"webhookFilters" yaml:"webhookFilters"`
	// Trigger a batch build from a webhook instead of a standard one.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	WebhookTriggersBatchBuild *bool `json:"webhookTriggersBatchBuild" yaml:"webhookTriggersBatchBuild"`
}

// TODO: EXAMPLE
//
type BucketCacheOptions struct {
	// The prefix to use to store the cache in the bucket.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// TODO: EXAMPLE
//
type BuildEnvironment struct {
	// The image used for the builds.
	BuildImage IBuildImage `json:"buildImage" yaml:"buildImage"`
	// The location of the PEM-encoded certificate for the build project.
	Certificate *BuildEnvironmentCertificate `json:"certificate" yaml:"certificate"`
	// The type of compute to use for this build.
	//
	// See the {@link ComputeType} enum for the possible values.
	ComputeType ComputeType `json:"computeType" yaml:"computeType"`
	// The environment variables that your builds can use.
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// Indicates how the project builds Docker images.
	//
	// Specify true to enable
	// running the Docker daemon inside a Docker container. This value must be
	// set to true only if this build project will be used to build Docker
	// images, and the specified build environment image is not one provided by
	// AWS CodeBuild with Docker support. Otherwise, all associated builds that
	// attempt to interact with the Docker daemon will fail.
	Privileged *bool `json:"privileged" yaml:"privileged"`
}

// Location of a PEM certificate on S3.
//
// TODO: EXAMPLE
//
type BuildEnvironmentCertificate struct {
	// The bucket where the certificate is.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// The full path and name of the key file.
	ObjectKey *string `json:"objectKey" yaml:"objectKey"`
}

// TODO: EXAMPLE
//
type BuildEnvironmentVariable struct {
	// The value of the environment variable.
	//
	// For plain-text variables (the default), this is the literal value of variable.
	// For SSM parameter variables, pass the name of the parameter here (`parameterName` property of `IParameter`).
	// For SecretsManager variables secrets, pass either the secret name (`secretName` property of `ISecret`)
	// or the secret ARN (`secretArn` property of `ISecret`) here,
	// along with optional SecretsManager qualifiers separated by ':', like the JSON key, or the version or stage
	// (see https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.secrets-manager for details).
	Value interface{} `json:"value" yaml:"value"`
	// The type of environment variable.
	Type BuildEnvironmentVariableType `json:"type" yaml:"type"`
}

// TODO: EXAMPLE
//
type BuildEnvironmentVariableType string

const (
	BuildEnvironmentVariableType_PLAINTEXT BuildEnvironmentVariableType = "PLAINTEXT"
	BuildEnvironmentVariableType_PARAMETER_STORE BuildEnvironmentVariableType = "PARAMETER_STORE"
	BuildEnvironmentVariableType_SECRETS_MANAGER BuildEnvironmentVariableType = "SECRETS_MANAGER"
)

// Optional arguments to {@link IBuildImage.binder} - currently empty.
//
// TODO: EXAMPLE
//
type BuildImageBindOptions struct {
}

// The return type from {@link IBuildImage.binder} - currently empty.
//
// TODO: EXAMPLE
//
type BuildImageConfig struct {
}

// BuildSpec for CodeBuild projects.
//
// TODO: EXAMPLE
//
type BuildSpec interface {
	IsImmediate() *bool
	ToBuildSpec() *string
}

// The jsii proxy struct for BuildSpec
type jsiiProxy_BuildSpec struct {
	_ byte // padding
}

func (j *jsiiProxy_BuildSpec) IsImmediate() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isImmediate",
		&returns,
	)
	return returns
}


func NewBuildSpec_Override(b BuildSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		nil, // no parameters
		b,
	)
}

func BuildSpec_FromObject(value *map[string]interface{}) BuildSpec {
	_init_.Initialize()

	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromObject",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Create a buildspec from an object that will be rendered as YAML in the resulting CloudFormation template.
func BuildSpec_FromObjectToYaml(value *map[string]interface{}) BuildSpec {
	_init_.Initialize()

	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromObjectToYaml",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Use a file from the source as buildspec.
//
// Use this if you want to use a file different from 'buildspec.yml'`
func BuildSpec_FromSourceFilename(filename *string) BuildSpec {
	_init_.Initialize()

	var returns BuildSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.BuildSpec",
		"fromSourceFilename",
		[]interface{}{filename},
		&returns,
	)

	return returns
}

// Render the represented BuildSpec.
func (b *jsiiProxy_BuildSpec) ToBuildSpec() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toBuildSpec",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Cache options for CodeBuild Project.
//
// A cache can store reusable pieces of your build environment and use them across multiple builds.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-caching.html
//
type Cache interface {
}

// The jsii proxy struct for Cache
type jsiiProxy_Cache struct {
	_ byte // padding
}

func NewCache_Override(c Cache) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Cache",
		nil, // no parameters
		c,
	)
}

// Create an S3 caching strategy.
func Cache_Bucket(bucket awss3.IBucket, options *BucketCacheOptions) Cache {
	_init_.Initialize()

	var returns Cache

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Cache",
		"bucket",
		[]interface{}{bucket, options},
		&returns,
	)

	return returns
}

// Create a local caching strategy.
func Cache_Local(modes ...LocalCacheMode) Cache {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range modes {
		args = append(args, a)
	}

	var returns Cache

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Cache",
		"local",
		args,
		&returns,
	)

	return returns
}

func Cache_None() Cache {
	_init_.Initialize()

	var returns Cache

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Cache",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

// A CloudFormation `AWS::CodeBuild::Project`.
//
// The `AWS::CodeBuild::Project` resource configures how AWS CodeBuild builds your source code. For example, it tells CodeBuild where to get the source code and which build environment to use.
//
// TODO: EXAMPLE
//
type CfnProject interface {
	awscdk.CfnResource
	awscdk.IInspectable
	Artifacts() interface{}
	SetArtifacts(val interface{})
	AttrArn() *string
	BadgeEnabled() interface{}
	SetBadgeEnabled(val interface{})
	BuildBatchConfig() interface{}
	SetBuildBatchConfig(val interface{})
	Cache() interface{}
	SetCache(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ConcurrentBuildLimit() *float64
	SetConcurrentBuildLimit(val *float64)
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	EncryptionKey() *string
	SetEncryptionKey(val *string)
	Environment() interface{}
	SetEnvironment(val interface{})
	FileSystemLocations() interface{}
	SetFileSystemLocations(val interface{})
	LogicalId() *string
	LogsConfig() interface{}
	SetLogsConfig(val interface{})
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	QueuedTimeoutInMinutes() *float64
	SetQueuedTimeoutInMinutes(val *float64)
	Ref() *string
	ResourceAccessRole() *string
	SetResourceAccessRole(val *string)
	SecondaryArtifacts() interface{}
	SetSecondaryArtifacts(val interface{})
	SecondarySources() interface{}
	SetSecondarySources(val interface{})
	SecondarySourceVersions() interface{}
	SetSecondarySourceVersions(val interface{})
	ServiceRole() *string
	SetServiceRole(val *string)
	Source() interface{}
	SetSource(val interface{})
	SourceVersion() *string
	SetSourceVersion(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	Triggers() interface{}
	SetTriggers(val interface{})
	UpdatedProperites() *map[string]interface{}
	Visibility() *string
	SetVisibility(val *string)
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

// The jsii proxy struct for CfnProject
type jsiiProxy_CfnProject struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnProject) Artifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) BadgeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) BuildBatchConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"buildBatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Cache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ConcurrentBuildLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) FileSystemLocations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) LogsConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) QueuedTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ResourceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SecondaryArtifacts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SecondarySources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SecondarySourceVersions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourceVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Source() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) SourceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Triggers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProject) VpcConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeBuild::Project`.
func NewCfnProject(scope constructs.Construct, id *string, props *CfnProjectProps) CfnProject {
	_init_.Initialize()

	j := jsiiProxy_CfnProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeBuild::Project`.
func NewCfnProject_Override(c CfnProject, scope constructs.Construct, id *string, props *CfnProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnProject) SetArtifacts(val interface{}) {
	_jsii_.Set(
		j,
		"artifacts",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetBadgeEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"badgeEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetBuildBatchConfig(val interface{}) {
	_jsii_.Set(
		j,
		"buildBatchConfig",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetCache(val interface{}) {
	_jsii_.Set(
		j,
		"cache",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetConcurrentBuildLimit(val *float64) {
	_jsii_.Set(
		j,
		"concurrentBuildLimit",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetEncryptionKey(val *string) {
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetEnvironment(val interface{}) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetFileSystemLocations(val interface{}) {
	_jsii_.Set(
		j,
		"fileSystemLocations",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetLogsConfig(val interface{}) {
	_jsii_.Set(
		j,
		"logsConfig",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetQueuedTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"queuedTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetResourceAccessRole(val *string) {
	_jsii_.Set(
		j,
		"resourceAccessRole",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetSecondaryArtifacts(val interface{}) {
	_jsii_.Set(
		j,
		"secondaryArtifacts",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetSecondarySources(val interface{}) {
	_jsii_.Set(
		j,
		"secondarySources",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetSecondarySourceVersions(val interface{}) {
	_jsii_.Set(
		j,
		"secondarySourceVersions",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetSource(val interface{}) {
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetSourceVersion(val *string) {
	_jsii_.Set(
		j,
		"sourceVersion",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetTimeoutInMinutes(val *float64) {
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetTriggers(val interface{}) {
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetVisibility(val *string) {
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_CfnProject) SetVpcConfig(val interface{}) {
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
func CfnProject_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnProject_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnProject",
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
func CfnProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProject_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.CfnProject",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnProject) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnProject) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnProject) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnProject) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnProject) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnProject) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnProject) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnProject) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnProject) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnProject) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnProject) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnProject) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnProject) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `Artifacts` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies output settings for artifacts generated by an AWS CodeBuild build.
//
// TODO: EXAMPLE
//
type CfnProject_ArtifactsProperty struct {
	// The type of build output artifact. Valid values include:.
	//
	// - `CODEPIPELINE` : The build project has build output generated through CodePipeline.
	//
	// > The `CODEPIPELINE` type is not supported for `secondaryArtifacts` .
	// - `NO_ARTIFACTS` : The build project does not produce any build output.
	// - `S3` : The build project stores build output in Amazon S3.
	Type *string `json:"type" yaml:"type"`
	// An identifier for this artifact definition.
	ArtifactIdentifier *string `json:"artifactIdentifier" yaml:"artifactIdentifier"`
	// Set to true if you do not want your output artifacts encrypted.
	//
	// This option is valid only if your artifacts type is Amazon Simple Storage Service (Amazon S3). If this is set with another artifacts type, an `invalidInputException` is thrown.
	EncryptionDisabled interface{} `json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Information about the build output artifact location:.
	//
	// - If `type` is set to `CODEPIPELINE` , AWS CodePipeline ignores this value if specified. This is because CodePipeline manages its build output locations instead of CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , this is the name of the output bucket.
	//
	// If you specify `CODEPIPELINE` or `NO_ARTIFACTS` for the `Type` property, don't specify this property. For all of the other types, you must specify this property.
	Location *string `json:"location" yaml:"location"`
	// Along with `path` and `namespaceType` , the pattern that AWS CodeBuild uses to name and store the output artifact:.
	//
	// - If `type` is set to `CODEPIPELINE` , AWS CodePipeline ignores this value if specified. This is because CodePipeline manages its build output names instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , this is the name of the output artifact object. If you set the name to be a forward slash ("/"), the artifact is stored in the root of the output bucket.
	//
	// For example:
	//
	// - If `path` is set to `MyArtifacts` , `namespaceType` is set to `BUILD_ID` , and `name` is set to `MyArtifact.zip` , then the output artifact is stored in `MyArtifacts/ *build-ID* /MyArtifact.zip` .
	// - If `path` is empty, `namespaceType` is set to `NONE` , and `name` is set to " `/` ", the output artifact is stored in the root of the output bucket.
	// - If `path` is set to `MyArtifacts` , `namespaceType` is set to `BUILD_ID` , and `name` is set to " `/` ", the output artifact is stored in `MyArtifacts/ *build-ID*` .
	//
	// If you specify `CODEPIPELINE` or `NO_ARTIFACTS` for the `Type` property, don't specify this property. For all of the other types, you must specify this property.
	Name *string `json:"name" yaml:"name"`
	// Along with `path` and `name` , the pattern that AWS CodeBuild uses to determine the name and location to store the output artifact:  - If `type` is set to `CODEPIPELINE` , CodePipeline ignores this value if specified.
	//
	// This is because CodePipeline manages its build output names instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , valid values include:
	//
	// - `BUILD_ID` : Include the build ID in the location of the build output artifact.
	// - `NONE` : Do not include the build ID. This is the default if `namespaceType` is not specified.
	//
	// For example, if `path` is set to `MyArtifacts` , `namespaceType` is set to `BUILD_ID` , and `name` is set to `MyArtifact.zip` , the output artifact is stored in `MyArtifacts/<build-ID>/MyArtifact.zip` .
	NamespaceType *string `json:"namespaceType" yaml:"namespaceType"`
	// If set to true a name specified in the buildspec file overrides the artifact name.
	//
	// The name specified in a buildspec file is calculated at build time and uses the Shell command language. For example, you can append a date and time to your artifact name so that it is always unique.
	OverrideArtifactName interface{} `json:"overrideArtifactName" yaml:"overrideArtifactName"`
	// The type of build output artifact to create:.
	//
	// - If `type` is set to `CODEPIPELINE` , CodePipeline ignores this value if specified. This is because CodePipeline manages its build output artifacts instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , valid values include:
	//
	// - `NONE` : AWS CodeBuild creates in the output bucket a folder that contains the build output. This is the default if `packaging` is not specified.
	// - `ZIP` : AWS CodeBuild creates in the output bucket a ZIP file that contains the build output.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// Along with `namespaceType` and `name` , the pattern that AWS CodeBuild uses to name and store the output artifact:.
	//
	// - If `type` is set to `CODEPIPELINE` , CodePipeline ignores this value if specified. This is because CodePipeline manages its build output names instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , this is the path to the output artifact. If `path` is not specified, `path` is not used.
	//
	// For example, if `path` is set to `MyArtifacts` , `namespaceType` is set to `NONE` , and `name` is set to `MyArtifact.zip` , the output artifact is stored in the output bucket at `MyArtifacts/MyArtifact.zip` .
	Path *string `json:"path" yaml:"path"`
}

// Specifies restrictions for the batch build.
//
// TODO: EXAMPLE
//
type CfnProject_BatchRestrictionsProperty struct {
	// An array of strings that specify the compute types that are allowed for the batch build.
	//
	// See [Build environment compute types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild User Guide* for these values.
	ComputeTypesAllowed *[]*string `json:"computeTypesAllowed" yaml:"computeTypesAllowed"`
	// Specifies the maximum number of builds allowed.
	MaximumBuildsAllowed *float64 `json:"maximumBuildsAllowed" yaml:"maximumBuildsAllowed"`
}

// Contains information that defines how the AWS CodeBuild build project reports the build status to the source provider.
//
// TODO: EXAMPLE
//
type CfnProject_BuildStatusConfigProperty struct {
	// Specifies the context of the build status CodeBuild sends to the source provider.
	//
	// The usage of this parameter depends on the source provider.
	//
	// - **Bitbucket** - This parameter is used for the `name` parameter in the Bitbucket commit status. For more information, see [build](https://docs.aws.amazon.com/https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build) in the Bitbucket API documentation.
	// - **GitHub/GitHub Enterprise Server** - This parameter is used for the `context` parameter in the GitHub commit status. For more information, see [Create a commit status](https://docs.aws.amazon.com/https://developer.github.com/v3/repos/statuses/#create-a-commit-status) in the GitHub developer guide.
	Context *string `json:"context" yaml:"context"`
	// Specifies the target url of the build status CodeBuild sends to the source provider.
	//
	// The usage of this parameter depends on the source provider.
	//
	// - **Bitbucket** - This parameter is used for the `url` parameter in the Bitbucket commit status. For more information, see [build](https://docs.aws.amazon.com/https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build) in the Bitbucket API documentation.
	// - **GitHub/GitHub Enterprise Server** - This parameter is used for the `target_url` parameter in the GitHub commit status. For more information, see [Create a commit status](https://docs.aws.amazon.com/https://developer.github.com/v3/repos/statuses/#create-a-commit-status) in the GitHub developer guide.
	TargetUrl *string `json:"targetUrl" yaml:"targetUrl"`
}

// `CloudWatchLogs` is a property of the [AWS CodeBuild Project LogsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html) property type that specifies settings for CloudWatch logs generated by an AWS CodeBuild build.
//
// TODO: EXAMPLE
//
type CfnProject_CloudWatchLogsConfigProperty struct {
	// The current status of the logs in CloudWatch Logs for a build project. Valid values are:.
	//
	// - `ENABLED` : CloudWatch Logs are enabled for this build project.
	// - `DISABLED` : CloudWatch Logs are not enabled for this build project.
	Status *string `json:"status" yaml:"status"`
	// The group name of the logs in CloudWatch Logs.
	//
	// For more information, see [Working with Log Groups and Log Streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html) .
	GroupName *string `json:"groupName" yaml:"groupName"`
	// The prefix of the stream name of the CloudWatch Logs.
	//
	// For more information, see [Working with Log Groups and Log Streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html) .
	StreamName *string `json:"streamName" yaml:"streamName"`
}

// `Environment` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies the environment for an AWS CodeBuild project.
//
// TODO: EXAMPLE
//
type CfnProject_EnvironmentProperty struct {
	// The type of compute environment.
	//
	// This determines the number of CPU cores and memory the build environment uses. Available values include:
	//
	// - `BUILD_GENERAL1_SMALL` : Use up to 3 GB memory and 2 vCPUs for builds.
	// - `BUILD_GENERAL1_MEDIUM` : Use up to 7 GB memory and 4 vCPUs for builds.
	// - `BUILD_GENERAL1_LARGE` : Use up to 15 GB memory and 8 vCPUs for builds.
	//
	// For more information, see [Build Environment Compute Types](https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild User Guide.*
	ComputeType *string `json:"computeType" yaml:"computeType"`
	// The image tag or image digest that identifies the Docker image to use for this build project.
	//
	// Use the following formats:
	//
	// - For an image tag: `<registry>/<repository>:<tag>` . For example, in the Docker repository that CodeBuild uses to manage its Docker images, this would be `aws/codebuild/standard:4.0` .
	// - For an image digest: `<registry>/<repository>@<digest>` . For example, to specify an image with the digest "sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0382cfbdbf," use `<registry>/<repository>@sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0382cfbdbf` .
	//
	// For more information, see [Docker images provided by CodeBuild](https://docs.aws.amazon.com//codebuild/latest/userguide/build-env-ref-available.html) in the *AWS CodeBuild user guide* .
	Image *string `json:"image" yaml:"image"`
	// The type of build environment to use for related builds.
	//
	// - The environment type `ARM_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), EU (Ireland), Asia Pacific (Mumbai), Asia Pacific (Tokyo), Asia Pacific (Sydney), and EU (Frankfurt).
	// - The environment type `LINUX_CONTAINER` with compute type `build.general1.2xlarge` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), Canada (Central), EU (Ireland), EU (London), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Seoul), Asia Pacific (Singapore), Asia Pacific (Sydney), China (Beijing), and China (Ningxia).
	// - The environment type `LINUX_GPU_CONTAINER` is available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), Canada (Central), EU (Ireland), EU (London), EU (Frankfurt), Asia Pacific (Tokyo), Asia Pacific (Seoul), Asia Pacific (Singapore), Asia Pacific (Sydney) , China (Beijing), and China (Ningxia).
	//
	// - The environment types `WINDOWS_CONTAINER` and `WINDOWS_SERVER_2019_CONTAINER` are available only in regions US East (N. Virginia), US East (Ohio), US West (Oregon), and EU (Ireland).
	//
	// For more information, see [Build environment compute types](https://docs.aws.amazon.com//codebuild/latest/userguide/build-env-ref-compute-types.html) in the *AWS CodeBuild user guide* .
	Type *string `json:"type" yaml:"type"`
	// The ARN of the Amazon S3 bucket, path prefix, and object key that contains the PEM-encoded certificate for the build project.
	//
	// For more information, see [certificate](https://docs.aws.amazon.com/codebuild/latest/userguide/create-project-cli.html#cli.environment.certificate) in the *AWS CodeBuild User Guide* .
	Certificate *string `json:"certificate" yaml:"certificate"`
	// A set of environment variables to make available to builds for this build project.
	EnvironmentVariables interface{} `json:"environmentVariables" yaml:"environmentVariables"`
	// The type of credentials AWS CodeBuild uses to pull images in your build. There are two valid values:.
	//
	// - `CODEBUILD` specifies that AWS CodeBuild uses its own credentials. This requires that you modify your ECR repository policy to trust AWS CodeBuild service principal.
	// - `SERVICE_ROLE` specifies that AWS CodeBuild uses your build project's service role.
	//
	// When you use a cross-account or private registry image, you must use SERVICE_ROLE credentials. When you use an AWS CodeBuild curated image, you must use CODEBUILD credentials.
	ImagePullCredentialsType *string `json:"imagePullCredentialsType" yaml:"imagePullCredentialsType"`
	// Enables running the Docker daemon inside a Docker container.
	//
	// Set to true only if the build project is used to build Docker images. Otherwise, a build that attempts to interact with the Docker daemon fails. The default setting is `false` .
	//
	// You can initialize the Docker daemon during the install phase of your build by adding one of the following sets of commands to the install phase of your buildspec file:
	//
	// If the operating system's base image is Ubuntu Linux:
	//
	// `- nohup /usr/local/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay&`
	//
	// `- timeout 15 sh -c "until docker info; do echo .; sleep 1; done"`
	//
	// If the operating system's base image is Alpine Linux and the previous command does not work, add the `-t` argument to `timeout` :
	//
	// `- nohup /usr/local/bin/dockerd --host=unix:///var/run/docker.sock --host=tcp://0.0.0.0:2375 --storage-driver=overlay&`
	//
	// `- timeout -t 15 sh -c "until docker info; do echo .; sleep 1; done"`
	PrivilegedMode interface{} `json:"privilegedMode" yaml:"privilegedMode"`
	// `RegistryCredential` is a property of the [AWS::CodeBuild::Project Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-environment) property that specifies information about credentials that provide access to a private Docker registry. When this is set:.
	//
	// - `imagePullCredentialsType` must be set to `SERVICE_ROLE` .
	// - images cannot be curated or an Amazon ECR image.
	RegistryCredential interface{} `json:"registryCredential" yaml:"registryCredential"`
}

// `EnvironmentVariable` is a property of the [AWS CodeBuild Project Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html) property type that specifies the name and value of an environment variable for an AWS CodeBuild project environment. When you use the environment to run a build, these variables are available for your builds to use. `EnvironmentVariable` contains a list of `EnvironmentVariable` property types.
//
// TODO: EXAMPLE
//
type CfnProject_EnvironmentVariableProperty struct {
	// The name or key of the environment variable.
	Name *string `json:"name" yaml:"name"`
	// The value of the environment variable.
	//
	// > We strongly discourage the use of `PLAINTEXT` environment variables to store sensitive values, especially AWS secret key IDs and secret access keys. `PLAINTEXT` environment variables can be displayed in plain text using the AWS CodeBuild console and the AWS CLI . For sensitive values, we recommend you use an environment variable of type `PARAMETER_STORE` or `SECRETS_MANAGER` .
	Value *string `json:"value" yaml:"value"`
	// The type of environment variable. Valid values include:.
	//
	// - `PARAMETER_STORE` : An environment variable stored in Systems Manager Parameter Store. To learn how to specify a parameter store environment variable, see [env/parameter-store](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.parameter-store) in the *AWS CodeBuild User Guide* .
	// - `PLAINTEXT` : An environment variable in plain text format. This is the default value.
	// - `SECRETS_MANAGER` : An environment variable stored in AWS Secrets Manager . To learn how to specify a secrets manager environment variable, see [env/secrets-manager](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec.env.secrets-manager) in the *AWS CodeBuild User Guide* .
	Type *string `json:"type" yaml:"type"`
}

// `GitSubmodulesConfig` is a property of the [AWS CodeBuild Project Source](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html) property type that specifies information about the Git submodules configuration for the build project.
//
// TODO: EXAMPLE
//
type CfnProject_GitSubmodulesConfigProperty struct {
	// Set to true to fetch Git submodules for your AWS CodeBuild build project.
	FetchSubmodules interface{} `json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

// `LogsConfig` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies information about logs for a build project. These can be logs in Amazon CloudWatch Logs, built in a specified S3 bucket, or both.
//
// TODO: EXAMPLE
//
type CfnProject_LogsConfigProperty struct {
	// Information about CloudWatch Logs for a build project.
	//
	// CloudWatch Logs are enabled by default.
	CloudWatchLogs interface{} `json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Information about logs built to an S3 bucket for a build project.
	//
	// S3 logs are not enabled by default.
	S3Logs interface{} `json:"s3Logs" yaml:"s3Logs"`
}

// Contains configuration information about a batch build project.
//
// TODO: EXAMPLE
//
type CfnProject_ProjectBuildBatchConfigProperty struct {
	// Specifies how build status reports are sent to the source provider for the batch build.
	//
	// This property is only used when the source provider for your project is Bitbucket, GitHub, or GitHub Enterprise, and your project is configured to report build statuses to the source provider.
	//
	// - **REPORT_AGGREGATED_BATCH** - (Default) Aggregate all of the build statuses into a single status report.
	// - **REPORT_INDIVIDUAL_BUILDS** - Send a separate status report for each individual build.
	BatchReportMode *string `json:"batchReportMode" yaml:"batchReportMode"`
	// Specifies if the build artifacts for the batch build should be combined into a single artifact location.
	CombineArtifacts interface{} `json:"combineArtifacts" yaml:"combineArtifacts"`
	// A `BatchRestrictions` object that specifies the restrictions for the batch build.
	Restrictions interface{} `json:"restrictions" yaml:"restrictions"`
	// Specifies the service role ARN for the batch build project.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// Specifies the maximum amount of time, in minutes, that the batch build must be completed in.
	TimeoutInMins *float64 `json:"timeoutInMins" yaml:"timeoutInMins"`
}

// `ProjectCache` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies information about the cache for the build project. If `ProjectCache` is not specified, then both of its properties default to `NO_CACHE` .
//
// TODO: EXAMPLE
//
type CfnProject_ProjectCacheProperty struct {
	// The type of cache used by the build project. Valid values include:.
	//
	// - `NO_CACHE` : The build project does not use any cache.
	// - `S3` : The build project reads and writes from and to S3.
	// - `LOCAL` : The build project stores a cache locally on a build host that is only available to that build host.
	Type *string `json:"type" yaml:"type"`
	// Information about the cache location:.
	//
	// - `NO_CACHE` or `LOCAL` : This value is ignored.
	// - `S3` : This is the S3 bucket name/prefix.
	Location *string `json:"location" yaml:"location"`
	// An array of strings that specify the local cache modes.
	//
	// You can use one or more local cache modes at the same time. This is only used for `LOCAL` cache types.
	//
	// Possible values are:
	//
	// - **LOCAL_SOURCE_CACHE** - Caches Git metadata for primary and secondary sources. After the cache is created, subsequent builds pull only the change between commits. This mode is a good choice for projects with a clean working directory and a source that is a large Git repository. If you choose this option and your project does not use a Git repository (GitHub, GitHub Enterprise, or Bitbucket), the option is ignored.
	// - **LOCAL_DOCKER_LAYER_CACHE** - Caches existing Docker layers. This mode is a good choice for projects that build or pull large Docker images. It can prevent the performance issues caused by pulling large Docker images down from the network.
	//
	// > - You can use a Docker layer cache in the Linux environment only.
	// > - The `privileged` flag must be set so that your project has the required Docker permissions.
	// > - You should consider the security implications before you use a Docker layer cache.
	// - **LOCAL_CUSTOM_CACHE** - Caches directories you specify in the buildspec file. This mode is a good choice if your build scenario is not suited to one of the other three local cache modes. If you use a custom cache:
	//
	// - Only directories can be specified for caching. You cannot specify individual files.
	// - Symlinks are used to reference cached directories.
	// - Cached directories are linked to your build before it downloads its project sources. Cached items are overridden if a source item has the same name. Directories are specified using cache paths in the buildspec file.
	Modes *[]*string `json:"modes" yaml:"modes"`
}

// Information about a file system created by Amazon Elastic File System (EFS).
//
// For more information, see [What Is Amazon Elastic File System?](https://docs.aws.amazon.com/efs/latest/ug/whatisefs.html)
//
// TODO: EXAMPLE
//
type CfnProject_ProjectFileSystemLocationProperty struct {
	// The name used to access a file system created by Amazon EFS.
	//
	// CodeBuild creates an environment variable by appending the `identifier` in all capital letters to `CODEBUILD_` . For example, if you specify `my_efs` for `identifier` , a new environment variable is create named `CODEBUILD_MY_EFS` .
	//
	// The `identifier` is used to mount your file system.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// A string that specifies the location of the file system created by Amazon EFS.
	//
	// Its format is `efs-dns-name:/directory-path` . You can find the DNS name of file system when you view it in the Amazon EFS console. The directory path is a path to a directory in the file system that CodeBuild mounts. For example, if the DNS name of a file system is `fs-abcd1234.efs.us-west-2.amazonaws.com` , and its mount directory is `my-efs-mount-directory` , then the `location` is `fs-abcd1234.efs.us-west-2.amazonaws.com:/my-efs-mount-directory` .
	//
	// The directory path in the format `efs-dns-name:/directory-path` is optional. If you do not specify a directory path, the location is only the DNS name and CodeBuild mounts the entire file system.
	Location *string `json:"location" yaml:"location"`
	// The location in the container where you mount the file system.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// The type of the file system.
	//
	// The one supported type is `EFS` .
	Type *string `json:"type" yaml:"type"`
	// The mount options for a file system created by Amazon EFS.
	//
	// The default mount options used by CodeBuild are `nfsvers=4.1,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2` . For more information, see [Recommended NFS Mount Options](https://docs.aws.amazon.com/efs/latest/ug/mounting-fs-nfs-mount-settings.html) .
	MountOptions *string `json:"mountOptions" yaml:"mountOptions"`
}

// A source identifier and its corresponding version.
//
// TODO: EXAMPLE
//
type CfnProject_ProjectSourceVersionProperty struct {
	// An identifier for a source in the build project.
	//
	// The identifier can only contain alphanumeric characters and underscores, and must be less than 128 characters in length.
	SourceIdentifier *string `json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// The source version for the corresponding source identifier. If specified, must be one of:.
	//
	// - For CodeCommit: the commit ID, branch, or Git tag to use.
	// - For GitHub: the commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build. If a pull request ID is specified, it must use the format `pr/pull-request-ID` (for example, `pr/25` ). If a branch name is specified, the branch's HEAD commit ID is used. If not specified, the default branch's HEAD commit ID is used.
	// - For Bitbucket: the commit ID, branch name, or tag name that corresponds to the version of the source code you want to build. If a branch name is specified, the branch's HEAD commit ID is used. If not specified, the default branch's HEAD commit ID is used.
	// - For Amazon S3: the version ID of the object that represents the build input ZIP file to use.
	//
	// For more information, see [Source Version Sample with CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html) in the *AWS CodeBuild User Guide* .
	SourceVersion *string `json:"sourceVersion" yaml:"sourceVersion"`
}

// `ProjectTriggers` is a property of the [AWS CodeBuild Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies webhooks that trigger an AWS CodeBuild build.
//
// TODO: EXAMPLE
//
type CfnProject_ProjectTriggersProperty struct {
	// Specifies the type of build this webhook will trigger. Allowed values are:.
	//
	// - **BUILD** - A single build
	// - **BUILD_BATCH** - A batch build
	BuildType *string `json:"buildType" yaml:"buildType"`
	// A list of lists of `WebhookFilter` objects used to determine which webhook events are triggered.
	//
	// At least one `WebhookFilter` in the array must specify `EVENT` as its type.
	FilterGroups interface{} `json:"filterGroups" yaml:"filterGroups"`
	// Specifies whether or not to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	Webhook interface{} `json:"webhook" yaml:"webhook"`
}

// `RegistryCredential` is a property of the [AWS CodeBuild Project Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html) property type that specifies information about credentials that provide access to a private Docker registry. When this is set:.
//
// - `imagePullCredentialsType` must be set to `SERVICE_ROLE` .
// - images cannot be curated or an Amazon ECR image.
//
// For more information, see [Private Registry with AWS Secrets Manager Sample for AWS CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-private-registry.html) .
//
// TODO: EXAMPLE
//
type CfnProject_RegistryCredentialProperty struct {
	// The Amazon Resource Name (ARN) or name of credentials created using AWS Secrets Manager .
	//
	// > The `credential` can use the name of the credentials only if they exist in your current AWS Region .
	Credential *string `json:"credential" yaml:"credential"`
	// The service that created the credentials to access a private Docker registry.
	//
	// The valid value, SECRETS_MANAGER, is for AWS Secrets Manager .
	CredentialProvider *string `json:"credentialProvider" yaml:"credentialProvider"`
}

// `S3Logs` is a property of the [AWS CodeBuild Project LogsConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html) property type that specifies settings for logs generated by an AWS CodeBuild build in an S3 bucket.
//
// TODO: EXAMPLE
//
type CfnProject_S3LogsConfigProperty struct {
	// The current status of the S3 build logs. Valid values are:.
	//
	// - `ENABLED` : S3 build logs are enabled for this build project.
	// - `DISABLED` : S3 build logs are not enabled for this build project.
	Status *string `json:"status" yaml:"status"`
	// Set to true if you do not want your S3 build log output encrypted.
	//
	// By default S3 build logs are encrypted.
	EncryptionDisabled interface{} `json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// The ARN of an S3 bucket and the path prefix for S3 logs.
	//
	// If your Amazon S3 bucket name is `my-bucket` , and your path prefix is `build-log` , then acceptable formats are `my-bucket/build-log` or `arn:aws:s3:::my-bucket/build-log` .
	Location *string `json:"location" yaml:"location"`
}

// `SourceAuth` is a property of the [AWS CodeBuild Project Source](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-source.html) property type that specifies authorization settings for AWS CodeBuild to access the source code to be built.
//
// `SourceAuth` is for use by the CodeBuild console only. Do not get or set it directly.
//
// TODO: EXAMPLE
//
type CfnProject_SourceAuthProperty struct {
	// The authorization type to use. The only valid value is `OAUTH` , which represents the OAuth authorization type.
	//
	// > This data type is used by the AWS CodeBuild console only.
	Type *string `json:"type" yaml:"type"`
	// The resource value that applies to the specified authorization type.
	//
	// > This data type is used by the AWS CodeBuild console only.
	Resource *string `json:"resource" yaml:"resource"`
}

// `Source` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies the source code settings for the project, such as the source code's repository type and location.
//
// TODO: EXAMPLE
//
type CfnProject_SourceProperty struct {
	// The type of repository that contains the source code to be built. Valid values include:.
	//
	// - `BITBUCKET` : The source code is in a Bitbucket repository.
	// - `CODECOMMIT` : The source code is in an CodeCommit repository.
	// - `CODEPIPELINE` : The source code settings are specified in the source action of a pipeline in CodePipeline.
	// - `GITHUB` : The source code is in a GitHub or GitHub Enterprise Cloud repository.
	// - `GITHUB_ENTERPRISE` : The source code is in a GitHub Enterprise Server repository.
	// - `NO_SOURCE` : The project does not have input source code.
	// - `S3` : The source code is in an Amazon S3 bucket.
	Type *string `json:"type" yaml:"type"`
	// Information about the authorization settings for AWS CodeBuild to access the source code to be built.
	//
	// This information is for the AWS CodeBuild console's use only. Your code should not get or set `Auth` directly.
	Auth interface{} `json:"auth" yaml:"auth"`
	// The build specification for the project.
	//
	// If this value is not provided, then the source code must contain a buildspec file named `buildspec.yml` at the root level. If this value is provided, it can be either a single string containing the entire build specification, or the path to an alternate buildspec file relative to the value of the built-in environment variable `CODEBUILD_SRC_DIR` . The alternate buildspec file can have a name other than `buildspec.yml` , for example `myspec.yml` or `build_spec_qa.yml` or similar. For more information, see the [Build Spec Reference](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-example) in the *AWS CodeBuild User Guide* .
	BuildSpec *string `json:"buildSpec" yaml:"buildSpec"`
	// Contains information that defines how the build project reports the build status to the source provider.
	//
	// This option is only used when the source provider is `GITHUB` , `GITHUB_ENTERPRISE` , or `BITBUCKET` .
	BuildStatusConfig interface{} `json:"buildStatusConfig" yaml:"buildStatusConfig"`
	// The depth of history to download.
	//
	// Minimum value is 0. If this value is 0, greater than 25, or not provided, then the full history is downloaded with each build project. If your source type is Amazon S3, this value is not supported.
	GitCloneDepth *float64 `json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// Information about the Git submodules configuration for the build project.
	GitSubmodulesConfig interface{} `json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
	// This is used with GitHub Enterprise only.
	//
	// Set to true to ignore SSL warnings while connecting to your GitHub Enterprise project repository. The default value is `false` . `InsecureSsl` should be used for testing purposes only. It should not be used in a production environment.
	InsecureSsl interface{} `json:"insecureSsl" yaml:"insecureSsl"`
	// Information about the location of the source code to be built. Valid values include:.
	//
	// - For source code settings that are specified in the source action of a pipeline in CodePipeline, `location` should not be specified. If it is specified, CodePipeline ignores it. This is because CodePipeline uses the settings in a pipeline's source action instead of this value.
	// - For source code in an CodeCommit repository, the HTTPS clone URL to the repository that contains the source code and the buildspec file (for example, `https://git-codecommit.<region-ID>.amazonaws.com/v1/repos/<repo-name>` ).
	// - For source code in an Amazon S3 input bucket, one of the following.
	//
	// - The path to the ZIP file that contains the source code (for example, `<bucket-name>/<path>/<object-name>.zip` ).
	// - The path to the folder that contains the source code (for example, `<bucket-name>/<path-to-source-code>/<folder>/` ).
	// - For source code in a GitHub repository, the HTTPS clone URL to the repository that contains the source and the buildspec file. You must connect your AWS account to your GitHub account. Use the AWS CodeBuild console to start creating a build project. When you use the console to connect (or reconnect) with GitHub, on the GitHub *Authorize application* page, for *Organization access* , choose *Request access* next to each repository you want to allow AWS CodeBuild to have access to, and then choose *Authorize application* . (After you have connected to your GitHub account, you do not need to finish creating the build project. You can leave the AWS CodeBuild console.) To instruct AWS CodeBuild to use this connection, in the `source` object, set the `auth` object's `type` value to `OAUTH` .
	// - For source code in a Bitbucket repository, the HTTPS clone URL to the repository that contains the source and the buildspec file. You must connect your AWS account to your Bitbucket account. Use the AWS CodeBuild console to start creating a build project. When you use the console to connect (or reconnect) with Bitbucket, on the Bitbucket *Confirm access to your account* page, choose *Grant access* . (After you have connected to your Bitbucket account, you do not need to finish creating the build project. You can leave the AWS CodeBuild console.) To instruct AWS CodeBuild to use this connection, in the `source` object, set the `auth` object's `type` value to `OAUTH` .
	//
	// If you specify `CODEPIPELINE` for the `Type` property, don't specify this property. For all of the other types, you must specify `Location` .
	Location *string `json:"location" yaml:"location"`
	// Set to true to report the status of a build's start and finish to your source provider.
	//
	// This option is valid only when your source provider is GitHub, GitHub Enterprise, or Bitbucket. If this is set and you use a different source provider, an `invalidInputException` is thrown.
	ReportBuildStatus interface{} `json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// An identifier for this project source.
	//
	// The identifier can only contain alphanumeric characters and underscores, and must be less than 128 characters in length.
	SourceIdentifier *string `json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

// `VpcConfig` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that enable AWS CodeBuild to access resources in an Amazon VPC. For more information, see [Use AWS CodeBuild with Amazon Virtual Private Cloud](https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html) in the *AWS CodeBuild User Guide* .
//
// TODO: EXAMPLE
//
type CfnProject_VpcConfigProperty struct {
	// A list of one or more security groups IDs in your Amazon VPC.
	//
	// The maximum count is 5.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of one or more subnet IDs in your Amazon VPC.
	//
	// The maximum count is 16.
	Subnets *[]*string `json:"subnets" yaml:"subnets"`
	// The ID of the Amazon VPC.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// `WebhookFilter` is a structure of the `FilterGroups` property on the [AWS CodeBuild Project ProjectTriggers](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projecttriggers.html) property type that specifies which webhooks trigger an AWS CodeBuild build.
//
// TODO: EXAMPLE
//
type CfnProject_WebhookFilterProperty struct {
	// For a `WebHookFilter` that uses `EVENT` type, a comma-separated string that specifies one or more events.
	//
	// For example, the webhook filter `PUSH, PULL_REQUEST_CREATED, PULL_REQUEST_UPDATED` allows all push, pull request created, and pull request updated events to trigger a build.
	//
	// For a `WebHookFilter` that uses any of the other filter types, a regular expression pattern. For example, a `WebHookFilter` that uses `HEAD_REF` for its `type` and the pattern `^refs/heads/` triggers a build when the head reference is a branch with a reference name `refs/heads/branch-name` .
	Pattern *string `json:"pattern" yaml:"pattern"`
	// The type of webhook filter.
	//
	// There are six webhook filter types: `EVENT` , `ACTOR_ACCOUNT_ID` , `HEAD_REF` , `BASE_REF` , `FILE_PATH` , and `COMMIT_MESSAGE` .
	//
	// - **EVENT** - A webhook event triggers a build when the provided `pattern` matches one of five event types: `PUSH` , `PULL_REQUEST_CREATED` , `PULL_REQUEST_UPDATED` , `PULL_REQUEST_REOPENED` , and `PULL_REQUEST_MERGED` . The `EVENT` patterns are specified as a comma-separated string. For example, `PUSH, PULL_REQUEST_CREATED, PULL_REQUEST_UPDATED` filters all push, pull request created, and pull request updated events.
	//
	// > The `PULL_REQUEST_REOPENED` works with GitHub and GitHub Enterprise only.
	// - **ACTOR_ACCOUNT_ID** - A webhook event triggers a build when a GitHub, GitHub Enterprise, or Bitbucket account ID matches the regular expression `pattern` .
	// - **HEAD_REF** - A webhook event triggers a build when the head reference matches the regular expression `pattern` . For example, `refs/heads/branch-name` and `refs/tags/tag-name` .
	//
	// Works with GitHub and GitHub Enterprise push, GitHub and GitHub Enterprise pull request, Bitbucket push, and Bitbucket pull request events.
	// - **BASE_REF** - A webhook event triggers a build when the base reference matches the regular expression `pattern` . For example, `refs/heads/branch-name` .
	//
	// > Works with pull request events only.
	// - **FILE_PATH** - A webhook triggers a build when the path of a changed file matches the regular expression `pattern` .
	//
	// > Works with GitHub and Bitbucket events push and pull requests events. Also works with GitHub Enterprise push events, but does not work with GitHub Enterprise pull request events.
	// - **COMMIT_MESSAGE** - A webhook triggers a build when the head commit message matches the regular expression `pattern` .
	//
	// > Works with GitHub and Bitbucket events push and pull requests events. Also works with GitHub Enterprise push events, but does not work with GitHub Enterprise pull request events.
	Type *string `json:"type" yaml:"type"`
	// Used to indicate that the `pattern` determines which webhook events do not trigger a build.
	//
	// If true, then a webhook event that does not match the `pattern` triggers a build. If false, then a webhook event that matches the `pattern` triggers a build.
	ExcludeMatchedPattern interface{} `json:"excludeMatchedPattern" yaml:"excludeMatchedPattern"`
}

// Properties for defining a `CfnProject`.
//
// TODO: EXAMPLE
//
type CfnProjectProps struct {
	// `Artifacts` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies output settings for artifacts generated by an AWS CodeBuild build.
	Artifacts interface{} `json:"artifacts" yaml:"artifacts"`
	// The build environment settings for the project, such as the environment type or the environment variables to use for the build environment.
	Environment interface{} `json:"environment" yaml:"environment"`
	// The ARN of the IAM role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
	ServiceRole *string `json:"serviceRole" yaml:"serviceRole"`
	// The source code settings for the project, such as the source code's repository type and location.
	Source interface{} `json:"source" yaml:"source"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see [Build Badges Sample](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-badges.html) in the *AWS CodeBuild User Guide* .
	//
	// > Including build badges with your project is currently not supported if the source type is CodePipeline. If you specify `CODEPIPELINE` for the `Source` property, do not specify the `BadgeEnabled` property.
	BadgeEnabled interface{} `json:"badgeEnabled" yaml:"badgeEnabled"`
	// A `ProjectBuildBatchConfig` object that defines the batch build options for the project.
	BuildBatchConfig interface{} `json:"buildBatchConfig" yaml:"buildBatchConfig"`
	// Settings that AWS CodeBuild uses to store and reuse build dependencies.
	Cache interface{} `json:"cache" yaml:"cache"`
	// The maximum number of concurrent builds that are allowed for this project.
	//
	// New builds are only started if the current number of builds is less than or equal to this limit. If the current build count meets this limit, new builds are throttled and are not run.
	ConcurrentBuildLimit *float64 `json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description that makes the build project easy to identify.
	Description *string `json:"description" yaml:"description"`
	// The AWS Key Management Service customer master key (CMK) to be used for encrypting the build output artifacts.
	//
	// > You can use a cross-account KMS key to encrypt the build output artifacts if your service role has permission to that key.
	//
	// You can specify either the Amazon Resource Name (ARN) of the CMK or, if available, the CMK's alias (using the format `alias/<alias-name>` ). If you don't specify a value, CodeBuild uses the managed CMK for Amazon Simple Storage Service (Amazon S3).
	EncryptionKey *string `json:"encryptionKey" yaml:"encryptionKey"`
	// An array of `ProjectFileSystemLocation` objects for a CodeBuild build project.
	//
	// A `ProjectFileSystemLocation` object specifies the `identifier` , `location` , `mountOptions` , `mountPoint` , and `type` of a file system created using Amazon Elastic File System.
	FileSystemLocations interface{} `json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Information about logs for the build project.
	//
	// A project can create logs in CloudWatch Logs, an S3 bucket, or both.
	LogsConfig interface{} `json:"logsConfig" yaml:"logsConfig"`
	// The name of the build project.
	//
	// The name must be unique across all of the projects in your AWS account .
	Name *string `json:"name" yaml:"name"`
	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes *float64 `json:"queuedTimeoutInMinutes" yaml:"queuedTimeoutInMinutes"`
	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole *string `json:"resourceAccessRole" yaml:"resourceAccessRole"`
	// A list of `Artifacts` objects.
	//
	// Each artifacts object specifies output settings that the project generates during a build.
	SecondaryArtifacts interface{} `json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// An array of `ProjectSource` objects.
	SecondarySources interface{} `json:"secondarySources" yaml:"secondarySources"`
	// An array of `ProjectSourceVersion` objects.
	//
	// If `secondarySourceVersions` is specified at the build level, then they take over these `secondarySourceVersions` (at the project level).
	SecondarySourceVersions interface{} `json:"secondarySourceVersions" yaml:"secondarySourceVersions"`
	// A version of the build input to be built for this project.
	//
	// If not specified, the latest version is used. If specified, it must be one of:
	//
	// - For CodeCommit: the commit ID, branch, or Git tag to use.
	// - For GitHub: the commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build. If a pull request ID is specified, it must use the format `pr/pull-request-ID` (for example `pr/25` ). If a branch name is specified, the branch's HEAD commit ID is used. If not specified, the default branch's HEAD commit ID is used.
	// - For Bitbucket: the commit ID, branch name, or tag name that corresponds to the version of the source code you want to build. If a branch name is specified, the branch's HEAD commit ID is used. If not specified, the default branch's HEAD commit ID is used.
	// - For Amazon S3: the version ID of the object that represents the build input ZIP file to use.
	//
	// If `sourceVersion` is specified at the build level, then that version takes precedence over this `sourceVersion` (at the project level).
	//
	// For more information, see [Source Version Sample with CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html) in the *AWS CodeBuild User Guide* .
	SourceVersion *string `json:"sourceVersion" yaml:"sourceVersion"`
	// An arbitrary set of tags (key-value pairs) for the AWS CodeBuild project.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild build project tags.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// How long, in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait before timing out any related build that did not get marked as completed.
	//
	// The default is 60 minutes.
	TimeoutInMinutes *float64 `json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// For an existing AWS CodeBuild build project that has its source code stored in a GitHub repository, enables AWS CodeBuild to begin automatically rebuilding the source code every time a code change is pushed to the repository.
	Triggers interface{} `json:"triggers" yaml:"triggers"`
	// Specifies the visibility of the project's builds. Possible values are:.
	//
	// - **PUBLIC_READ** - The project builds are visible to the public.
	// - **PRIVATE** - The project builds are not visible to the public.
	Visibility *string `json:"visibility" yaml:"visibility"`
	// `VpcConfig` specifies settings that enable AWS CodeBuild to access resources in an Amazon VPC.
	//
	// For more information, see [Use AWS CodeBuild with Amazon Virtual Private Cloud](https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html) in the *AWS CodeBuild User Guide* .
	VpcConfig interface{} `json:"vpcConfig" yaml:"vpcConfig"`
}

// A CloudFormation `AWS::CodeBuild::ReportGroup`.
//
// Represents a report group. A report group contains a collection of reports.
//
// TODO: EXAMPLE
//
type CfnReportGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DeleteReports() interface{}
	SetDeleteReports(val interface{})
	ExportConfig() interface{}
	SetExportConfig(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnReportGroup
type jsiiProxy_CfnReportGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnReportGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) DeleteReports() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) ExportConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnReportGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeBuild::ReportGroup`.
func NewCfnReportGroup(scope constructs.Construct, id *string, props *CfnReportGroupProps) CfnReportGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnReportGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnReportGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeBuild::ReportGroup`.
func NewCfnReportGroup_Override(c CfnReportGroup, scope constructs.Construct, id *string, props *CfnReportGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnReportGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnReportGroup) SetDeleteReports(val interface{}) {
	_jsii_.Set(
		j,
		"deleteReports",
		val,
	)
}

func (j *jsiiProxy_CfnReportGroup) SetExportConfig(val interface{}) {
	_jsii_.Set(
		j,
		"exportConfig",
		val,
	)
}

func (j *jsiiProxy_CfnReportGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnReportGroup) SetType(val *string) {
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
func CfnReportGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnReportGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnReportGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnReportGroup",
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
func CfnReportGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnReportGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnReportGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.CfnReportGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnReportGroup) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnReportGroup) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnReportGroup) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnReportGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnReportGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnReportGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnReportGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnReportGroup) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnReportGroup) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnReportGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnReportGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnReportGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnReportGroup) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnReportGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnReportGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Information about the location where the run of a report is exported.
//
// TODO: EXAMPLE
//
type CfnReportGroup_ReportExportConfigProperty struct {
	// The export configuration type. Valid values are:.
	//
	// - `S3` : The report results are exported to an S3 bucket.
	// - `NO_EXPORT` : The report results are not exported.
	ExportConfigType *string `json:"exportConfigType" yaml:"exportConfigType"`
	// A `S3ReportExportConfig` object that contains information about the S3 bucket where the run of a report is exported.
	S3Destination interface{} `json:"s3Destination" yaml:"s3Destination"`
}

// Information about the S3 bucket where the raw data of a report are exported.
//
// TODO: EXAMPLE
//
type CfnReportGroup_S3ReportExportConfigProperty struct {
	// The name of the S3 bucket where the raw data of a report are exported.
	Bucket *string `json:"bucket" yaml:"bucket"`
	// The AWS account identifier of the owner of the Amazon S3 bucket.
	//
	// This allows report data to be exported to an Amazon S3 bucket that is owned by an account other than the account running the build.
	BucketOwner *string `json:"bucketOwner" yaml:"bucketOwner"`
	// A boolean value that specifies if the results of a report are encrypted.
	EncryptionDisabled interface{} `json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// The encryption key for the report's encrypted raw data.
	EncryptionKey *string `json:"encryptionKey" yaml:"encryptionKey"`
	// The type of build output artifact to create. Valid values include:.
	//
	// - `NONE` : CodeBuild creates the raw data in the output bucket. This is the default if packaging is not specified.
	// - `ZIP` : CodeBuild creates a ZIP file with the raw data in the output bucket.
	Packaging *string `json:"packaging" yaml:"packaging"`
	// The path to the exported report's raw data results.
	Path *string `json:"path" yaml:"path"`
}

// Properties for defining a `CfnReportGroup`.
//
// TODO: EXAMPLE
//
type CfnReportGroupProps struct {
	// Information about the destination where the raw data of this `ReportGroup` is exported.
	ExportConfig interface{} `json:"exportConfig" yaml:"exportConfig"`
	// The type of the `ReportGroup` . This can be one of the following values:.
	//
	// - **CODE_COVERAGE** - The report group contains code coverage reports.
	// - **TEST** - The report group contains test reports.
	Type *string `json:"type" yaml:"type"`
	// When deleting a report group, specifies if reports within the report group should be deleted.
	//
	// - **true** - Deletes any reports that belong to the report group before deleting the report group.
	// - **false** - You must delete any reports in the report group. This is the default value. If you delete a report group that contains one or more reports, an exception is thrown.
	DeleteReports interface{} `json:"deleteReports" yaml:"deleteReports"`
	// The name of the `ReportGroup` .
	Name *string `json:"name" yaml:"name"`
	// A list of tag key and value pairs associated with this report group.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild report group tags.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::CodeBuild::SourceCredential`.
//
// Information about the credentials for a GitHub, GitHub Enterprise, or Bitbucket repository. We strongly recommend that you use AWS Secrets Manager to store your credentials. If you use Secrets Manager , you must have secrets in your secrets manager. For more information, see [Using Dynamic References to Specify Template Values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/dynamic-references.html#dynamic-references-secretsmanager) .
//
// > For security purposes, do not use plain text in your AWS CloudFormation template to store your credentials.
//
// TODO: EXAMPLE
//
type CfnSourceCredential interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AuthType() *string
	SetAuthType(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ServerType() *string
	SetServerType(val *string)
	Stack() awscdk.Stack
	Token() *string
	SetToken(val *string)
	UpdatedProperites() *map[string]interface{}
	Username() *string
	SetUsername(val *string)
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

// The jsii proxy struct for CfnSourceCredential
type jsiiProxy_CfnSourceCredential struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSourceCredential) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) ServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceCredential) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


// Create a new `AWS::CodeBuild::SourceCredential`.
func NewCfnSourceCredential(scope constructs.Construct, id *string, props *CfnSourceCredentialProps) CfnSourceCredential {
	_init_.Initialize()

	j := jsiiProxy_CfnSourceCredential{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnSourceCredential",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::CodeBuild::SourceCredential`.
func NewCfnSourceCredential_Override(c CfnSourceCredential, scope constructs.Construct, id *string, props *CfnSourceCredentialProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.CfnSourceCredential",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSourceCredential) SetAuthType(val *string) {
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_CfnSourceCredential) SetServerType(val *string) {
	_jsii_.Set(
		j,
		"serverType",
		val,
	)
}

func (j *jsiiProxy_CfnSourceCredential) SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_CfnSourceCredential) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnSourceCredential_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnSourceCredential",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnSourceCredential_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnSourceCredential",
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
func CfnSourceCredential_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.CfnSourceCredential",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSourceCredential_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.CfnSourceCredential",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnSourceCredential) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnSourceCredential) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnSourceCredential) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnSourceCredential) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnSourceCredential) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnSourceCredential) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnSourceCredential) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnSourceCredential) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnSourceCredential) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSourceCredential) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnSourceCredential) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSourceCredential) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnSourceCredential) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnSourceCredential) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnSourceCredential) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnSourceCredential`.
//
// TODO: EXAMPLE
//
type CfnSourceCredentialProps struct {
	// The type of authentication used by the credentials.
	//
	// Valid options are OAUTH, BASIC_AUTH, or PERSONAL_ACCESS_TOKEN.
	AuthType *string `json:"authType" yaml:"authType"`
	// The type of source provider.
	//
	// The valid options are GITHUB, GITHUB_ENTERPRISE, or BITBUCKET.
	ServerType *string `json:"serverType" yaml:"serverType"`
	// For GitHub or GitHub Enterprise, this is the personal access token.
	//
	// For Bitbucket, this is the app password.
	Token *string `json:"token" yaml:"token"`
	// The Bitbucket username when the `authType` is BASIC_AUTH.
	//
	// This parameter is not valid for other types of source providers or connections.
	Username *string `json:"username" yaml:"username"`
}

// Information about logs built to a CloudWatch Log Group for a build project.
//
// TODO: EXAMPLE
//
type CloudWatchLoggingOptions struct {
	// The current status of the logs in Amazon CloudWatch Logs for a build project.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The Log Group to send logs to.
	LogGroup awslogs.ILogGroup `json:"logGroup" yaml:"logGroup"`
	// The prefix of the stream name of the Amazon CloudWatch Logs.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// Construction properties for {@link CodeCommitSource}.
//
// TODO: EXAMPLE
//
type CodeCommitSourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `json:"identifier" yaml:"identifier"`
	Repository awscodecommit.IRepository `json:"repository" yaml:"repository"`
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build.
	//
	// TODO: EXAMPLE
	//
	BranchOrRef *string `json:"branchOrRef" yaml:"branchOrRef"`
	// The depth of history to download.
	//
	// Minimum value is 0.
	// If this value is 0, greater than 25, or not provided,
	// then the full history is downloaded with each build of the project.
	CloneDepth *float64 `json:"cloneDepth" yaml:"cloneDepth"`
	// Whether to fetch submodules while cloning git repo.
	FetchSubmodules *bool `json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

// TODO: EXAMPLE
//
type CommonProjectProps struct {
	// Whether to allow the CodeBuild to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// CodeBuild project to connect to network targets.
	//
	// Only used if 'vpc' is supplied.
	AllowAllOutbound *bool `json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see Build Badges Sample
	// in the AWS CodeBuild User Guide.
	Badge *bool `json:"badge" yaml:"badge"`
	// Filename or contents of buildspec in JSON format.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-example
	//
	BuildSpec BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Caching strategy to use.
	Cache Cache `json:"cache" yaml:"cache"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	CheckSecretsInPlainTextEnvVariables *bool `json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Maximum number of concurrent builds.
	//
	// Minimum value is 1 and maximum is account build limit.
	ConcurrentBuildLimit *float64 `json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description of the project.
	//
	// Use the description to identify the purpose
	// of the project.
	Description *string `json:"description" yaml:"description"`
	// Encryption key to use to read and write artifacts.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Build environment to use for the build.
	Environment *BuildEnvironment `json:"environment" yaml:"environment"`
	// Additional environment variables to add to the build environment.
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// An  ProjectFileSystemLocation objects for a CodeBuild build project.
	//
	// A ProjectFileSystemLocation object specifies the identifier, location, mountOptions, mountPoint,
	// and type of a file system created using Amazon Elastic File System.
	FileSystemLocations *[]IFileSystemLocation `json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Add permissions to this project's role to create and use test report groups with name starting with the name of this project.
	//
	// That is the standard report group that gets created when a simple name
	// (in contrast to an ARN)
	// is used in the 'reports' section of the buildspec of this project.
	// This is usually harmless, but you can turn these off if you don't plan on using test
	// reports in this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/test-report-group-naming.html
	//
	GrantReportGroupPermissions *bool `json:"grantReportGroupPermissions" yaml:"grantReportGroupPermissions"`
	// Information about logs for the build project.
	//
	// A project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
	Logging *LoggingOptions `json:"logging" yaml:"logging"`
	// The physical, human-readable name of the CodeBuild Project.
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// The number of minutes after which AWS CodeBuild stops the build if it's still in queue.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	QueuedTimeout awscdk.Duration `json:"queuedTimeout" yaml:"queuedTimeout"`
	// Service Role to assume while running the build.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// What security group to associate with the codebuild project's network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// VPC network to place codebuild network interfaces.
	//
	// Specify this if the codebuild project needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Build machine compute type.
//
// TODO: EXAMPLE
//
type ComputeType string

const (
	ComputeType_SMALL ComputeType = "SMALL"
	ComputeType_MEDIUM ComputeType = "MEDIUM"
	ComputeType_LARGE ComputeType = "LARGE"
	ComputeType_X2_LARGE ComputeType = "X2_LARGE"
)

// The options when creating a CodeBuild Docker build image using {@link LinuxBuildImage.fromDockerRegistry} or {@link WindowsBuildImage.fromDockerRegistry}.
//
// TODO: EXAMPLE
//
type DockerImageOptions struct {
	// The credentials, stored in Secrets Manager, used for accessing the repository holding the image, if the repository is private.
	SecretsManagerCredentials awssecretsmanager.ISecret `json:"secretsManagerCredentials" yaml:"secretsManagerCredentials"`
}

// Construction properties for {@link EfsFileSystemLocation}.
//
// TODO: EXAMPLE
//
type EfsFileSystemLocationProps struct {
	// The name used to access a file system created by Amazon EFS.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// A string that specifies the location of the file system, like Amazon EFS.
	//
	// This value looks like `fs-abcd1234.efs.us-west-2.amazonaws.com:/my-efs-mount-directory`.
	Location *string `json:"location" yaml:"location"`
	// The location in the container where you mount the file system.
	MountPoint *string `json:"mountPoint" yaml:"mountPoint"`
	// The mount options for a file system such as Amazon EFS.
	MountOptions *string `json:"mountOptions" yaml:"mountOptions"`
}

// The types of webhook event actions.
//
// TODO: EXAMPLE
//
type EventAction string

const (
	EventAction_PUSH EventAction = "PUSH"
	EventAction_PULL_REQUEST_CREATED EventAction = "PULL_REQUEST_CREATED"
	EventAction_PULL_REQUEST_UPDATED EventAction = "PULL_REQUEST_UPDATED"
	EventAction_PULL_REQUEST_MERGED EventAction = "PULL_REQUEST_MERGED"
	EventAction_PULL_REQUEST_REOPENED EventAction = "PULL_REQUEST_REOPENED"
)

// The type returned from {@link IFileSystemLocation#bind}.
//
// TODO: EXAMPLE
//
type FileSystemConfig struct {
	// File system location wrapper property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-projectfilesystemlocation.html
	//
	Location *CfnProject_ProjectFileSystemLocationProperty `json:"location" yaml:"location"`
}

// FileSystemLocation provider definition for a CodeBuild Project.
//
// TODO: EXAMPLE
//
type FileSystemLocation interface {
}

// The jsii proxy struct for FileSystemLocation
type jsiiProxy_FileSystemLocation struct {
	_ byte // padding
}

func NewFileSystemLocation() FileSystemLocation {
	_init_.Initialize()

	j := jsiiProxy_FileSystemLocation{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.FileSystemLocation",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewFileSystemLocation_Override(f FileSystemLocation) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.FileSystemLocation",
		nil, // no parameters
		f,
	)
}

// EFS file system provider.
func FileSystemLocation_Efs(props *EfsFileSystemLocationProps) IFileSystemLocation {
	_init_.Initialize()

	var returns IFileSystemLocation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.FileSystemLocation",
		"efs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// An object that represents a group of filter conditions for a webhook.
//
// Every condition in a given FilterGroup must be true in order for the whole group to be true.
// You construct instances of it by calling the {@link #inEventOf} static factory method,
// and then calling various `andXyz` instance methods to create modified instances of it
// (this class is immutable).
//
// You pass instances of this class to the `webhookFilters` property when constructing a source.
//
// TODO: EXAMPLE
//
type FilterGroup interface {
	AndActorAccountIs(pattern *string) FilterGroup
	AndActorAccountIsNot(pattern *string) FilterGroup
	AndBaseBranchIs(branchName *string) FilterGroup
	AndBaseBranchIsNot(branchName *string) FilterGroup
	AndBaseRefIs(pattern *string) FilterGroup
	AndBaseRefIsNot(pattern *string) FilterGroup
	AndBranchIs(branchName *string) FilterGroup
	AndBranchIsNot(branchName *string) FilterGroup
	AndCommitMessageIs(commitMessage *string) FilterGroup
	AndCommitMessageIsNot(commitMessage *string) FilterGroup
	AndFilePathIs(pattern *string) FilterGroup
	AndFilePathIsNot(pattern *string) FilterGroup
	AndHeadRefIs(pattern *string) FilterGroup
	AndHeadRefIsNot(pattern *string) FilterGroup
	AndTagIs(tagName *string) FilterGroup
	AndTagIsNot(tagName *string) FilterGroup
}

// The jsii proxy struct for FilterGroup
type jsiiProxy_FilterGroup struct {
	_ byte // padding
}

// Creates a new event FilterGroup that triggers on any of the provided actions.
func FilterGroup_InEventOf(actions ...EventAction) FilterGroup {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns FilterGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.FilterGroup",
		"inEventOf",
		args,
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the account ID of the actor initiating the event must match the given pattern.
func (f *jsiiProxy_FilterGroup) AndActorAccountIs(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andActorAccountIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the account ID of the actor initiating the event must not match the given pattern.
func (f *jsiiProxy_FilterGroup) AndActorAccountIsNot(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andActorAccountIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must target the given base branch.
//
// Note that you cannot use this method if this Group contains the `PUSH` event action.
func (f *jsiiProxy_FilterGroup) AndBaseBranchIs(branchName *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseBranchIs",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must not target the given base branch.
//
// Note that you cannot use this method if this Group contains the `PUSH` event action.
func (f *jsiiProxy_FilterGroup) AndBaseBranchIsNot(branchName *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseBranchIsNot",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must target the given Git reference.
//
// Note that you cannot use this method if this Group contains the `PUSH` event action.
func (f *jsiiProxy_FilterGroup) AndBaseRefIs(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseRefIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the Pull Request that is the source of the event must not target the given Git reference.
//
// Note that you cannot use this method if this Group contains the `PUSH` event action.
func (f *jsiiProxy_FilterGroup) AndBaseRefIsNot(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBaseRefIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must affect the given branch.
func (f *jsiiProxy_FilterGroup) AndBranchIs(branchName *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBranchIs",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must not affect the given branch.
func (f *jsiiProxy_FilterGroup) AndBranchIsNot(branchName *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andBranchIsNot",
		[]interface{}{branchName},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must affect a head commit with the given message.
func (f *jsiiProxy_FilterGroup) AndCommitMessageIs(commitMessage *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andCommitMessageIs",
		[]interface{}{commitMessage},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must not affect a head commit with the given message.
func (f *jsiiProxy_FilterGroup) AndCommitMessageIsNot(commitMessage *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andCommitMessageIsNot",
		[]interface{}{commitMessage},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the push that is the source of the event must affect a file that matches the given pattern.
//
// Note that you can only use this method if this Group contains only the `PUSH` event action,
// and only for GitHub, Bitbucket and GitHubEnterprise sources.
func (f *jsiiProxy_FilterGroup) AndFilePathIs(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andFilePathIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the push that is the source of the event must not affect a file that matches the given pattern.
//
// Note that you can only use this method if this Group contains only the `PUSH` event action,
// and only for GitHub, Bitbucket and GitHubEnterprise sources.
func (f *jsiiProxy_FilterGroup) AndFilePathIsNot(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andFilePathIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must affect a Git reference (ie., a branch or a tag) that matches the given pattern.
func (f *jsiiProxy_FilterGroup) AndHeadRefIs(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andHeadRefIs",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must not affect a Git reference (ie., a branch or a tag) that matches the given pattern.
func (f *jsiiProxy_FilterGroup) AndHeadRefIsNot(pattern *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andHeadRefIsNot",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must affect the given tag.
func (f *jsiiProxy_FilterGroup) AndTagIs(tagName *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andTagIs",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

// Create a new FilterGroup with an added condition: the event must not affect the given tag.
func (f *jsiiProxy_FilterGroup) AndTagIsNot(tagName *string) FilterGroup {
	var returns FilterGroup

	_jsii_.Invoke(
		f,
		"andTagIsNot",
		[]interface{}{tagName},
		&returns,
	)

	return returns
}

// The source credentials used when contacting the GitHub Enterprise API.
//
// **Note**: CodeBuild only allows a single credential for GitHub Enterprise
// to be saved in a given AWS account in a given region -
// any attempt to add more than one will result in an error.
//
// TODO: EXAMPLE
//
type GitHubEnterpriseSourceCredentials interface {
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

// The jsii proxy struct for GitHubEnterpriseSourceCredentials
type jsiiProxy_GitHubEnterpriseSourceCredentials struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_GitHubEnterpriseSourceCredentials) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubEnterpriseSourceCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubEnterpriseSourceCredentials) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubEnterpriseSourceCredentials) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewGitHubEnterpriseSourceCredentials(scope constructs.Construct, id *string, props *GitHubEnterpriseSourceCredentialsProps) GitHubEnterpriseSourceCredentials {
	_init_.Initialize()

	j := jsiiProxy_GitHubEnterpriseSourceCredentials{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.GitHubEnterpriseSourceCredentials",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGitHubEnterpriseSourceCredentials_Override(g GitHubEnterpriseSourceCredentials, scope constructs.Construct, id *string, props *GitHubEnterpriseSourceCredentialsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.GitHubEnterpriseSourceCredentials",
		[]interface{}{scope, id, props},
		g,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GitHubEnterpriseSourceCredentials_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.GitHubEnterpriseSourceCredentials",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func GitHubEnterpriseSourceCredentials_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.GitHubEnterpriseSourceCredentials",
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
func (g *jsiiProxy_GitHubEnterpriseSourceCredentials) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GitHubEnterpriseSourceCredentials) GeneratePhysicalName() *string {
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
func (g *jsiiProxy_GitHubEnterpriseSourceCredentials) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (g *jsiiProxy_GitHubEnterpriseSourceCredentials) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (g *jsiiProxy_GitHubEnterpriseSourceCredentials) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creation properties for {@link GitHubEnterpriseSourceCredentials}.
//
// TODO: EXAMPLE
//
type GitHubEnterpriseSourceCredentialsProps struct {
	// The personal access token to use when contacting the instance of the GitHub Enterprise API.
	AccessToken awscdk.SecretValue `json:"accessToken" yaml:"accessToken"`
}

// Construction properties for {@link GitHubEnterpriseSource}.
//
// TODO: EXAMPLE
//
type GitHubEnterpriseSourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// The HTTPS URL of the repository in your GitHub Enterprise installation.
	HttpsCloneUrl *string `json:"httpsCloneUrl" yaml:"httpsCloneUrl"`
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build.
	//
	// TODO: EXAMPLE
	//
	BranchOrRef *string `json:"branchOrRef" yaml:"branchOrRef"`
	// The depth of history to download.
	//
	// Minimum value is 0.
	// If this value is 0, greater than 25, or not provided,
	// then the full history is downloaded with each build of the project.
	CloneDepth *float64 `json:"cloneDepth" yaml:"cloneDepth"`
	// Whether to fetch submodules while cloning git repo.
	FetchSubmodules *bool `json:"fetchSubmodules" yaml:"fetchSubmodules"`
	// Whether to ignore SSL errors when connecting to the repository.
	IgnoreSslErrors *bool `json:"ignoreSslErrors" yaml:"ignoreSslErrors"`
	// Whether to send notifications on your build's start and end.
	ReportBuildStatus *bool `json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// Whether to create a webhook that will trigger a build every time an event happens in the repository.
	Webhook *bool `json:"webhook" yaml:"webhook"`
	// A list of webhook filters that can constraint what events in the repository will trigger a build.
	//
	// A build is triggered if any of the provided filter groups match.
	// Only valid if `webhook` was not provided as false.
	WebhookFilters *[]FilterGroup `json:"webhookFilters" yaml:"webhookFilters"`
	// Trigger a batch build from a webhook instead of a standard one.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	WebhookTriggersBatchBuild *bool `json:"webhookTriggersBatchBuild" yaml:"webhookTriggersBatchBuild"`
}

// The source credentials used when contacting the GitHub API.
//
// **Note**: CodeBuild only allows a single credential for GitHub
// to be saved in a given AWS account in a given region -
// any attempt to add more than one will result in an error.
//
// TODO: EXAMPLE
//
type GitHubSourceCredentials interface {
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

// The jsii proxy struct for GitHubSourceCredentials
type jsiiProxy_GitHubSourceCredentials struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_GitHubSourceCredentials) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubSourceCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubSourceCredentials) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GitHubSourceCredentials) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewGitHubSourceCredentials(scope constructs.Construct, id *string, props *GitHubSourceCredentialsProps) GitHubSourceCredentials {
	_init_.Initialize()

	j := jsiiProxy_GitHubSourceCredentials{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.GitHubSourceCredentials",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGitHubSourceCredentials_Override(g GitHubSourceCredentials, scope constructs.Construct, id *string, props *GitHubSourceCredentialsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.GitHubSourceCredentials",
		[]interface{}{scope, id, props},
		g,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func GitHubSourceCredentials_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.GitHubSourceCredentials",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func GitHubSourceCredentials_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.GitHubSourceCredentials",
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
func (g *jsiiProxy_GitHubSourceCredentials) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GitHubSourceCredentials) GeneratePhysicalName() *string {
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
func (g *jsiiProxy_GitHubSourceCredentials) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (g *jsiiProxy_GitHubSourceCredentials) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (g *jsiiProxy_GitHubSourceCredentials) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Creation properties for {@link GitHubSourceCredentials}.
//
// TODO: EXAMPLE
//
type GitHubSourceCredentialsProps struct {
	// The personal access token to use when contacting the GitHub API.
	AccessToken awscdk.SecretValue `json:"accessToken" yaml:"accessToken"`
}

// Construction properties for {@link GitHubSource} and {@link GitHubEnterpriseSource}.
//
// TODO: EXAMPLE
//
type GitHubSourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// The GitHub account/user that owns the repo.
	//
	// TODO: EXAMPLE
	//
	Owner *string `json:"owner" yaml:"owner"`
	// The name of the repo (without the username).
	//
	// TODO: EXAMPLE
	//
	Repo *string `json:"repo" yaml:"repo"`
	// The commit ID, pull request ID, branch name, or tag name that corresponds to the version of the source code you want to build.
	//
	// TODO: EXAMPLE
	//
	BranchOrRef *string `json:"branchOrRef" yaml:"branchOrRef"`
	// The depth of history to download.
	//
	// Minimum value is 0.
	// If this value is 0, greater than 25, or not provided,
	// then the full history is downloaded with each build of the project.
	CloneDepth *float64 `json:"cloneDepth" yaml:"cloneDepth"`
	// Whether to fetch submodules while cloning git repo.
	FetchSubmodules *bool `json:"fetchSubmodules" yaml:"fetchSubmodules"`
	// Whether to send notifications on your build's start and end.
	ReportBuildStatus *bool `json:"reportBuildStatus" yaml:"reportBuildStatus"`
	// Whether to create a webhook that will trigger a build every time an event happens in the repository.
	Webhook *bool `json:"webhook" yaml:"webhook"`
	// A list of webhook filters that can constraint what events in the repository will trigger a build.
	//
	// A build is triggered if any of the provided filter groups match.
	// Only valid if `webhook` was not provided as false.
	WebhookFilters *[]FilterGroup `json:"webhookFilters" yaml:"webhookFilters"`
	// Trigger a batch build from a webhook instead of a standard one.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	WebhookTriggersBatchBuild *bool `json:"webhookTriggersBatchBuild" yaml:"webhookTriggersBatchBuild"`
}

// The abstract interface of a CodeBuild build output.
//
// Implemented by {@link Artifacts}.
type IArtifacts interface {
	// Callback when an Artifacts class is used in a CodeBuild Project.
	Bind(scope constructs.Construct, project IProject) *ArtifactsConfig
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	Identifier() *string
	// The CodeBuild type of this artifact.
	Type() *string
}

// The jsii proxy for IArtifacts
type jsiiProxy_IArtifacts struct {
	_ byte // padding
}

func (i *jsiiProxy_IArtifacts) Bind(scope constructs.Construct, project IProject) *ArtifactsConfig {
	var returns *ArtifactsConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IArtifacts) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IArtifacts) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// A variant of {@link IBuildImage} that allows binding to the project.
type IBindableBuildImage interface {
	IBuildImage
	// Function that allows the build image access to the construct tree.
	Bind(scope constructs.Construct, project IProject, options *BuildImageBindOptions) *BuildImageConfig
}

// The jsii proxy for IBindableBuildImage
type jsiiProxy_IBindableBuildImage struct {
	jsiiProxy_IBuildImage
}

func (i *jsiiProxy_IBindableBuildImage) Bind(scope constructs.Construct, project IProject, options *BuildImageBindOptions) *BuildImageConfig {
	var returns *BuildImageConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project, options},
		&returns,
	)

	return returns
}

// Represents a Docker image used for the CodeBuild Project builds.
//
// Use the concrete subclasses, either:
// {@link LinuxBuildImage} or {@link WindowsBuildImage}.
type IBuildImage interface {
	// Make a buildspec to run the indicated script.
	RunScriptBuildspec(entrypoint *string) BuildSpec
	// Allows the image a chance to validate whether the passed configuration is correct.
	Validate(buildEnvironment *BuildEnvironment) *[]*string
	// The default {@link ComputeType} to use with this image, if one was not specified in {@link BuildEnvironment#computeType} explicitly.
	DefaultComputeType() ComputeType
	// The Docker image identifier that the build environment uses.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
	//
	ImageId() *string
	// The type of principal that CodeBuild will use to pull this build Docker image.
	ImagePullPrincipalType() ImagePullPrincipalType
	// An optional ECR repository that the image is hosted in.
	Repository() awsecr.IRepository
	// The secretsManagerCredentials for access to a private registry.
	SecretsManagerCredentials() awssecretsmanager.ISecret
	// The type of build environment.
	Type() *string
}

// The jsii proxy for IBuildImage
type jsiiProxy_IBuildImage struct {
	_ byte // padding
}

func (i *jsiiProxy_IBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	var returns BuildSpec

	_jsii_.Invoke(
		i,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// The interface of a CodeBuild FileSystemLocation.
//
// Implemented by {@link EfsFileSystemLocation}.
type IFileSystemLocation interface {
	// Called by the project when a file system is added so it can perform binding operations on this file system location.
	Bind(scope constructs.Construct, project IProject) *FileSystemConfig
}

// The jsii proxy for IFileSystemLocation
type jsiiProxy_IFileSystemLocation struct {
	_ byte // padding
}

func (i *jsiiProxy_IFileSystemLocation) Bind(scope constructs.Construct, project IProject) *FileSystemConfig {
	var returns *FileSystemConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project},
		&returns,
	)

	return returns
}

type IProject interface {
	awsec2.IConnectable
	awsiam.IGrantable
	awscodestarnotifications.INotificationRuleSource
	awscdk.IResource
	AddToRolePolicy(policyStatement awsiam.PolicyStatement)
	// Enable batch builds.
	//
	// Returns an object contining the batch service role if batch builds
	// could be enabled.
	EnableBatchBuilds() *BatchBuildConfig
	// Returns: a CloudWatch metric associated with this build project.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of builds triggered.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum
	MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the duration of all builds over time.
	//
	// Units: Seconds
	//
	// Valid CloudWatch statistics: Average (recommended), Maximum, Minimum
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of builds that failed because of client error or because of a timeout.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum
	MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Measures the number of successful builds.
	//
	// Units: Count
	//
	// Valid CloudWatch statistics: Sum
	MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Defines a CodeStar Notification rule triggered when the project events emitted by you specified, it very similar to `onEvent` API.
	//
	// You can also use the methods `notifyOnBuildSucceeded` and
	// `notifyOnBuildFailed` to define rules for these specific event emitted.
	//
	// Returns: CodeStar Notifications rule associated with this build project.
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar notification rule which triggers when a build fails.
	NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines a CodeStar notification rule which triggers when a build completes successfully.
	NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	// Defines an event rule which triggers when a build fails.
	OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an event rule which triggers when a build starts.
	OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an event rule which triggers when a build completes successfully.
	OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when something happens with this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule that triggers upon phase change of this build project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines a CloudWatch event rule triggered when the build project state changes.
	//
	// You can filter specific build status events using an event
	// pattern filter on the `build-status` detail field:
	//
	//     const rule = project.onStateChange('OnBuildStarted', { target });
	//     rule.addEventPattern({
	//       detail: {
	//         'build-status': [
	//           "IN_PROGRESS",
	//           "SUCCEEDED",
	//           "FAILED",
	//           "STOPPED"
	//         ]
	//       }
	//     });
	//
	// You can also use the methods `onBuildFailed` and `onBuildSucceeded` to define rules for
	// these specific state changes.
	//
	// To access fields from the event in the event target input,
	// use the static fields on the `StateChangeEvent` class.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
	//
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// The ARN of this Project.
	ProjectArn() *string
	// The human-visible name of this Project.
	ProjectName() *string
	// The IAM service Role of this Project.
	//
	// Undefined for imported Projects.
	Role() awsiam.IRole
}

// The jsii proxy for IProject
type jsiiProxy_IProject struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscodestarnotificationsINotificationRuleSource
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IProject) AddToRolePolicy(policyStatement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{policyStatement},
	)
}

func (i *jsiiProxy_IProject) EnableBatchBuilds() *BatchBuildConfig {
	var returns *BatchBuildConfig

	_jsii_.Invoke(
		i,
		"enableBatchBuilds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFailedBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSucceededBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBuildFailed",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		i,
		"notifyOnBuildSucceeded",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onBuildFailed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onBuildStarted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onBuildSucceeded",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onPhaseChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		i,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IProject) BindAsNotificationRuleSource(scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleSource",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IProject) ProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

// The interface representing the ReportGroup resource - either an existing one, imported using the {@link ReportGroup.fromReportGroupName} method, or a new one, created with the {@link ReportGroup} class.
type IReportGroup interface {
	awscdk.IResource
	// Grants the given entity permissions to write (that is, upload reports to) this report group.
	GrantWrite(identity awsiam.IGrantable) awsiam.Grant
	// The ARN of the ReportGroup.
	ReportGroupArn() *string
	// The name of the ReportGroup.
	ReportGroupName() *string
}

// The jsii proxy for IReportGroup
type jsiiProxy_IReportGroup struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IReportGroup) GrantWrite(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IReportGroup) ReportGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReportGroup) ReportGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportGroupName",
		&returns,
	)
	return returns
}

// The abstract interface of a CodeBuild source.
//
// Implemented by {@link Source}.
type ISource interface {
	Bind(scope constructs.Construct, project IProject) *SourceConfig
	BadgeSupported() *bool
	Identifier() *string
	Type() *string
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISource) Bind(scope constructs.Construct, project IProject) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, project},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ISource) BadgeSupported() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"badgeSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISource) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISource) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

// The type of principal CodeBuild will use to pull your build Docker image.
type ImagePullPrincipalType string

const (
	ImagePullPrincipalType_CODEBUILD ImagePullPrincipalType = "CODEBUILD"
	ImagePullPrincipalType_SERVICE_ROLE ImagePullPrincipalType = "SERVICE_ROLE"
)

// A CodeBuild image running aarch64 Linux.
//
// This class has a bunch of public constants that represent the CodeBuild ARM images.
//
// You can also specify a custom image using the static method:
//
// - LinuxBuildImage.fromEcrRepository(repo[, tag])
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
type LinuxArmBuildImage interface {
	IBuildImage
	DefaultComputeType() ComputeType
	ImageId() *string
	ImagePullPrincipalType() ImagePullPrincipalType
	Repository() awsecr.IRepository
	SecretsManagerCredentials() awssecretsmanager.ISecret
	Type() *string
	RunScriptBuildspec(entrypoint *string) BuildSpec
	Validate(buildEnvironment *BuildEnvironment) *[]*string
}

// The jsii proxy struct for LinuxArmBuildImage
type jsiiProxy_LinuxArmBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_LinuxArmBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxArmBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Uses a Docker image provided by CodeBuild.
//
// Returns: A Docker image provided by CodeBuild.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
func LinuxArmBuildImage_FromCodeBuildImageId(id *string) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"fromCodeBuildImageId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Returns an ARM image running Linux from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
//
// Returns: An aarch64 Linux build image from an ECR repository.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
func LinuxArmBuildImage_FromEcrRepository(repository awsecr.IRepository, tag *string) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

func LinuxArmBuildImage_AMAZON_LINUX_2_STANDARD_1_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"AMAZON_LINUX_2_STANDARD_1_0",
		&returns,
	)
	return returns
}

func LinuxArmBuildImage_AMAZON_LINUX_2_STANDARD_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxArmBuildImage",
		"AMAZON_LINUX_2_STANDARD_2_0",
		&returns,
	)
	return returns
}

// Make a buildspec to run the indicated script.
func (l *jsiiProxy_LinuxArmBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	var returns BuildSpec

	_jsii_.Invoke(
		l,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

// Validates by checking the BuildEnvironment computeType as aarch64 images only support ComputeType.SMALL and ComputeType.LARGE.
func (l *jsiiProxy_LinuxArmBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

// A CodeBuild image running x86-64 Linux.
//
// This class has a bunch of public constants that represent the most popular images.
//
// You can also specify a custom image using one of the static methods:
//
// - LinuxBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }])
// - LinuxBuildImage.fromEcrRepository(repo[, tag])
// - LinuxBuildImage.fromAsset(parent, id, props)
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
type LinuxBuildImage interface {
	IBuildImage
	DefaultComputeType() ComputeType
	ImageId() *string
	ImagePullPrincipalType() ImagePullPrincipalType
	Repository() awsecr.IRepository
	SecretsManagerCredentials() awssecretsmanager.ISecret
	Type() *string
	RunScriptBuildspec(entrypoint *string) BuildSpec
	Validate(_arg *BuildEnvironment) *[]*string
}

// The jsii proxy struct for LinuxBuildImage
type jsiiProxy_LinuxBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_LinuxBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Uses an Docker image asset as a x86-64 Linux build image.
func LinuxBuildImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"fromAsset",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Uses a Docker image provided by CodeBuild.
//
// Returns: A Docker image provided by CodeBuild.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
func LinuxBuildImage_FromCodeBuildImageId(id *string) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"fromCodeBuildImageId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Returns: a x86-64 Linux build image from a Docker Hub image.
func LinuxBuildImage_FromDockerRegistry(name *string, options *DockerImageOptions) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"fromDockerRegistry",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Returns: A x86-64 Linux build image from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
func LinuxBuildImage_FromEcrRepository(repository awsecr.IRepository, tag *string) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_3() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_3",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_ARM() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_ARM",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_ARM_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_ARM_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_1_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_1_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_2_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_3_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_3_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_4_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_4_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_5_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_5_0",
		&returns,
	)
	return returns
}

// Make a buildspec to run the indicated script.
func (l *jsiiProxy_LinuxBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	var returns BuildSpec

	_jsii_.Invoke(
		l,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

// Allows the image a chance to validate whether the passed configuration is correct.
func (l *jsiiProxy_LinuxBuildImage) Validate(_arg *BuildEnvironment) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

// A CodeBuild GPU image running Linux.
//
// This class has public constants that represent the most popular GPU images from AWS Deep Learning Containers.
//
// TODO: EXAMPLE
//
// See: https://aws.amazon.com/releasenotes/available-deep-learning-containers-images
//
type LinuxGpuBuildImage interface {
	IBindableBuildImage
	DefaultComputeType() ComputeType
	ImageId() *string
	ImagePullPrincipalType() ImagePullPrincipalType
	Type() *string
	Bind(scope constructs.Construct, project IProject, _options *BuildImageBindOptions) *BuildImageConfig
	RunScriptBuildspec(entrypoint *string) BuildSpec
	Validate(buildEnvironment *BuildEnvironment) *[]*string
}

// The jsii proxy struct for LinuxGpuBuildImage
type jsiiProxy_LinuxGpuBuildImage struct {
	jsiiProxy_IBindableBuildImage
}

func (j *jsiiProxy_LinuxGpuBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxGpuBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxGpuBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxGpuBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Returns a Linux GPU build image from AWS Deep Learning Containers.
// See: https://aws.amazon.com/releasenotes/available-deep-learning-containers-images
//
func LinuxGpuBuildImage_AwsDeepLearningContainersImage(repositoryName *string, tag *string, account *string) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"awsDeepLearningContainersImage",
		[]interface{}{repositoryName, tag, account},
		&returns,
	)

	return returns
}

// Returns a GPU image running Linux from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
func LinuxGpuBuildImage_FromEcrRepository(repository awsecr.IRepository, tag *string) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

func LinuxGpuBuildImage_DLC_MXNET_1_4_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_MXNET_1_4_1",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_MXNET_1_6_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_MXNET_1_6_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_2_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_3_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_3_1",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_4_0_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_4_0_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_4_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_4_0_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_5_0_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_5_0_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_5_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_5_0_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_14_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_14_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_15_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_15_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_15_2_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_15_2_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_15_2_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_15_2_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_0_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_0_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_0_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_0_1",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_1_0_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_1_0_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_1_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_1_0_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_2_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_2_0_TRAINING",
		&returns,
	)
	return returns
}

// Function that allows the build image access to the construct tree.
func (l *jsiiProxy_LinuxGpuBuildImage) Bind(scope constructs.Construct, project IProject, _options *BuildImageBindOptions) *BuildImageConfig {
	var returns *BuildImageConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, project, _options},
		&returns,
	)

	return returns
}

// Make a buildspec to run the indicated script.
func (l *jsiiProxy_LinuxGpuBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	var returns BuildSpec

	_jsii_.Invoke(
		l,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

// Allows the image a chance to validate whether the passed configuration is correct.
func (l *jsiiProxy_LinuxGpuBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

// Local cache modes to enable for the CodeBuild Project.
//
// TODO: EXAMPLE
//
type LocalCacheMode string

const (
	LocalCacheMode_SOURCE LocalCacheMode = "SOURCE"
	LocalCacheMode_DOCKER_LAYER LocalCacheMode = "DOCKER_LAYER"
	LocalCacheMode_CUSTOM LocalCacheMode = "CUSTOM"
)

// Information about logs for the build project.
//
// A project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
//
// TODO: EXAMPLE
//
type LoggingOptions struct {
	// Information about Amazon CloudWatch Logs for a build project.
	CloudWatch *CloudWatchLoggingOptions `json:"cloudWatch" yaml:"cloudWatch"`
	// Information about logs built to an S3 bucket for a build project.
	S3 *S3LoggingOptions `json:"s3" yaml:"s3"`
}

// Event fields for the CodeBuild "phase change" event.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html#sample-build-notifications-ref
//
type PhaseChangeEvent interface {
}

// The jsii proxy struct for PhaseChangeEvent
type jsiiProxy_PhaseChangeEvent struct {
	_ byte // padding
}

func PhaseChangeEvent_BuildComplete() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"buildComplete",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_BuildId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"buildId",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_CompletedPhase() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"completedPhase",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_CompletedPhaseDurationSeconds() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"completedPhaseDurationSeconds",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_CompletedPhaseStatus() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"completedPhaseStatus",
		&returns,
	)
	return returns
}

func PhaseChangeEvent_ProjectName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.PhaseChangeEvent",
		"projectName",
		&returns,
	)
	return returns
}

// A convenience class for CodeBuild Projects that are used in CodePipeline.
//
// TODO: EXAMPLE
//
type PipelineProject interface {
	Project
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() constructs.Node
	PhysicalName() *string
	ProjectArn() *string
	ProjectName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddFileSystemLocation(fileSystemLocation IFileSystemLocation)
	AddSecondaryArtifact(secondaryArtifact IArtifacts)
	AddSecondarySource(secondarySource ISource)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig
	BindToCodePipeline(_scope constructs.Construct, options *BindToCodePipelineOptions)
	EnableBatchBuilds() *BatchBuildConfig
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule
	NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	ToString() *string
}

// The jsii proxy struct for PipelineProject
type jsiiProxy_PipelineProject struct {
	jsiiProxy_Project
}

func (j *jsiiProxy_PipelineProject) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) ProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineProject) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewPipelineProject(scope constructs.Construct, id *string, props *PipelineProjectProps) PipelineProject {
	_init_.Initialize()

	j := jsiiProxy_PipelineProject{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.PipelineProject",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPipelineProject_Override(p PipelineProject, scope constructs.Construct, id *string, props *PipelineProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.PipelineProject",
		[]interface{}{scope, id, props},
		p,
	)
}

func PipelineProject_FromProjectArn(scope constructs.Construct, id *string, projectArn *string) IProject {
	_init_.Initialize()

	var returns IProject

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.PipelineProject",
		"fromProjectArn",
		[]interface{}{scope, id, projectArn},
		&returns,
	)

	return returns
}

// Import a Project defined either outside the CDK, or in a different CDK Stack (and exported using the {@link export} method).
//
// Returns: a reference to the existing Project
func PipelineProject_FromProjectName(scope constructs.Construct, id *string, projectName *string) IProject {
	_init_.Initialize()

	var returns IProject

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.PipelineProject",
		"fromProjectName",
		[]interface{}{scope, id, projectName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PipelineProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.PipelineProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func PipelineProject_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.PipelineProject",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Convert the environment variables map of string to {@link BuildEnvironmentVariable}, which is the customer-facing type, to a list of {@link CfnProject.EnvironmentVariableProperty}, which is the representation of environment variables in CloudFormation.
//
// Returns: an array of {@link CfnProject.EnvironmentVariableProperty} instances
func PipelineProject_SerializeEnvVariables(environmentVariables *map[string]*BuildEnvironmentVariable, validateNoPlainTextSecrets *bool, principal awsiam.IGrantable) *[]*CfnProject_EnvironmentVariableProperty {
	_init_.Initialize()

	var returns *[]*CfnProject_EnvironmentVariableProperty

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.PipelineProject",
		"serializeEnvVariables",
		[]interface{}{environmentVariables, validateNoPlainTextSecrets, principal},
		&returns,
	)

	return returns
}

// Adds a fileSystemLocation to the Project.
func (p *jsiiProxy_PipelineProject) AddFileSystemLocation(fileSystemLocation IFileSystemLocation) {
	_jsii_.InvokeVoid(
		p,
		"addFileSystemLocation",
		[]interface{}{fileSystemLocation},
	)
}

// Adds a secondary artifact to the Project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
//
func (p *jsiiProxy_PipelineProject) AddSecondaryArtifact(secondaryArtifact IArtifacts) {
	_jsii_.InvokeVoid(
		p,
		"addSecondaryArtifact",
		[]interface{}{secondaryArtifact},
	)
}

// Adds a secondary source to the Project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
//
func (p *jsiiProxy_PipelineProject) AddSecondarySource(secondarySource ISource) {
	_jsii_.InvokeVoid(
		p,
		"addSecondarySource",
		[]interface{}{secondarySource},
	)
}

// Add a permission only if there's a policy attached.
func (p *jsiiProxy_PipelineProject) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_PipelineProject) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns a source configuration for notification rule.
func (p *jsiiProxy_PipelineProject) BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		p,
		"bindAsNotificationRuleSource",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// A callback invoked when the given project is added to a CodePipeline.
func (p *jsiiProxy_PipelineProject) BindToCodePipeline(_scope constructs.Construct, options *BindToCodePipelineOptions) {
	_jsii_.InvokeVoid(
		p,
		"bindToCodePipeline",
		[]interface{}{_scope, options},
	)
}

// Enable batch builds.
//
// Returns an object contining the batch service role if batch builds
// could be enabled.
func (p *jsiiProxy_PipelineProject) EnableBatchBuilds() *BatchBuildConfig {
	var returns *BatchBuildConfig

	_jsii_.Invoke(
		p,
		"enableBatchBuilds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineProject) GeneratePhysicalName() *string {
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
func (p *jsiiProxy_PipelineProject) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (p *jsiiProxy_PipelineProject) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns: a CloudWatch metric associated with this build project.
func (p *jsiiProxy_PipelineProject) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Measures the number of builds triggered.
//
// Units: Count
//
// Valid CloudWatch statistics: Sum
func (p *jsiiProxy_PipelineProject) MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Measures the duration of all builds over time.
//
// Units: Seconds
//
// Valid CloudWatch statistics: Average (recommended), Maximum, Minimum
func (p *jsiiProxy_PipelineProject) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Measures the number of builds that failed because of client error or because of a timeout.
//
// Units: Count
//
// Valid CloudWatch statistics: Sum
func (p *jsiiProxy_PipelineProject) MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricFailedBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Measures the number of successful builds.
//
// Units: Count
//
// Valid CloudWatch statistics: Sum
func (p *jsiiProxy_PipelineProject) MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricSucceededBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Defines a CodeStar Notification rule triggered when the project events emitted by you specified, it very similar to `onEvent` API.
//
// You can also use the methods `notifyOnBuildSucceeded` and
// `notifyOnBuildFailed` to define rules for these specific event emitted.
func (p *jsiiProxy_PipelineProject) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Defines a CodeStar notification rule which triggers when a build fails.
func (p *jsiiProxy_PipelineProject) NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnBuildFailed",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Defines a CodeStar notification rule which triggers when a build completes successfully.
func (p *jsiiProxy_PipelineProject) NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnBuildSucceeded",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Defines an event rule which triggers when a build fails.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
func (p *jsiiProxy_PipelineProject) OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildFailed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an event rule which triggers when a build starts.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
func (p *jsiiProxy_PipelineProject) OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildStarted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an event rule which triggers when a build completes successfully.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
func (p *jsiiProxy_PipelineProject) OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildSucceeded",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule triggered when something happens with this project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
//
func (p *jsiiProxy_PipelineProject) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule that triggers upon phase change of this build project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
//
func (p *jsiiProxy_PipelineProject) OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onPhaseChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule triggered when the build project state changes.
//
// You can filter specific build status events using an event
// pattern filter on the `build-status` detail field:
//
//     const rule = project.onStateChange('OnBuildStarted', { target });
//     rule.addEventPattern({
//       detail: {
//         'build-status': [
//           "IN_PROGRESS",
//           "SUCCEEDED",
//           "FAILED",
//           "STOPPED"
//         ]
//       }
//     });
//
// You can also use the methods `onBuildFailed` and `onBuildSucceeded` to define rules for
// these specific state changes.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
//
func (p *jsiiProxy_PipelineProject) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_PipelineProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// TODO: EXAMPLE
//
type PipelineProjectProps struct {
	// Whether to allow the CodeBuild to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// CodeBuild project to connect to network targets.
	//
	// Only used if 'vpc' is supplied.
	AllowAllOutbound *bool `json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see Build Badges Sample
	// in the AWS CodeBuild User Guide.
	Badge *bool `json:"badge" yaml:"badge"`
	// Filename or contents of buildspec in JSON format.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-example
	//
	BuildSpec BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Caching strategy to use.
	Cache Cache `json:"cache" yaml:"cache"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	CheckSecretsInPlainTextEnvVariables *bool `json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Maximum number of concurrent builds.
	//
	// Minimum value is 1 and maximum is account build limit.
	ConcurrentBuildLimit *float64 `json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description of the project.
	//
	// Use the description to identify the purpose
	// of the project.
	Description *string `json:"description" yaml:"description"`
	// Encryption key to use to read and write artifacts.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Build environment to use for the build.
	Environment *BuildEnvironment `json:"environment" yaml:"environment"`
	// Additional environment variables to add to the build environment.
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// An  ProjectFileSystemLocation objects for a CodeBuild build project.
	//
	// A ProjectFileSystemLocation object specifies the identifier, location, mountOptions, mountPoint,
	// and type of a file system created using Amazon Elastic File System.
	FileSystemLocations *[]IFileSystemLocation `json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Add permissions to this project's role to create and use test report groups with name starting with the name of this project.
	//
	// That is the standard report group that gets created when a simple name
	// (in contrast to an ARN)
	// is used in the 'reports' section of the buildspec of this project.
	// This is usually harmless, but you can turn these off if you don't plan on using test
	// reports in this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/test-report-group-naming.html
	//
	GrantReportGroupPermissions *bool `json:"grantReportGroupPermissions" yaml:"grantReportGroupPermissions"`
	// Information about logs for the build project.
	//
	// A project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
	Logging *LoggingOptions `json:"logging" yaml:"logging"`
	// The physical, human-readable name of the CodeBuild Project.
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// The number of minutes after which AWS CodeBuild stops the build if it's still in queue.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	QueuedTimeout awscdk.Duration `json:"queuedTimeout" yaml:"queuedTimeout"`
	// Service Role to assume while running the build.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// What security group to associate with the codebuild project's network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// VPC network to place codebuild network interfaces.
	//
	// Specify this if the codebuild project needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// A representation of a CodeBuild Project.
//
// TODO: EXAMPLE
//
type Project interface {
	awscdk.Resource
	IProject
	Connections() awsec2.Connections
	Env() *awscdk.ResourceEnvironment
	GrantPrincipal() awsiam.IPrincipal
	Node() constructs.Node
	PhysicalName() *string
	ProjectArn() *string
	ProjectName() *string
	Role() awsiam.IRole
	Stack() awscdk.Stack
	AddFileSystemLocation(fileSystemLocation IFileSystemLocation)
	AddSecondaryArtifact(secondaryArtifact IArtifacts)
	AddSecondarySource(secondarySource ISource)
	AddToRolePolicy(statement awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig
	BindToCodePipeline(_scope constructs.Construct, options *BindToCodePipelineOptions)
	EnableBatchBuilds() *BatchBuildConfig
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule
	NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule
	OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	ToString() *string
}

// The jsii proxy struct for Project
type jsiiProxy_Project struct {
	internal.Type__awscdkResource
	jsiiProxy_IProject
}

func (j *jsiiProxy_Project) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Project) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewProject(scope constructs.Construct, id *string, props *ProjectProps) Project {
	_init_.Initialize()

	j := jsiiProxy_Project{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Project",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewProject_Override(p Project, scope constructs.Construct, id *string, props *ProjectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Project",
		[]interface{}{scope, id, props},
		p,
	)
}

func Project_FromProjectArn(scope constructs.Construct, id *string, projectArn *string) IProject {
	_init_.Initialize()

	var returns IProject

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"fromProjectArn",
		[]interface{}{scope, id, projectArn},
		&returns,
	)

	return returns
}

// Import a Project defined either outside the CDK, or in a different CDK Stack (and exported using the {@link export} method).
//
// Returns: a reference to the existing Project
func Project_FromProjectName(scope constructs.Construct, id *string, projectName *string) IProject {
	_init_.Initialize()

	var returns IProject

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"fromProjectName",
		[]interface{}{scope, id, projectName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Project_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Project_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Convert the environment variables map of string to {@link BuildEnvironmentVariable}, which is the customer-facing type, to a list of {@link CfnProject.EnvironmentVariableProperty}, which is the representation of environment variables in CloudFormation.
//
// Returns: an array of {@link CfnProject.EnvironmentVariableProperty} instances
func Project_SerializeEnvVariables(environmentVariables *map[string]*BuildEnvironmentVariable, validateNoPlainTextSecrets *bool, principal awsiam.IGrantable) *[]*CfnProject_EnvironmentVariableProperty {
	_init_.Initialize()

	var returns *[]*CfnProject_EnvironmentVariableProperty

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Project",
		"serializeEnvVariables",
		[]interface{}{environmentVariables, validateNoPlainTextSecrets, principal},
		&returns,
	)

	return returns
}

// Adds a fileSystemLocation to the Project.
func (p *jsiiProxy_Project) AddFileSystemLocation(fileSystemLocation IFileSystemLocation) {
	_jsii_.InvokeVoid(
		p,
		"addFileSystemLocation",
		[]interface{}{fileSystemLocation},
	)
}

// Adds a secondary artifact to the Project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
//
func (p *jsiiProxy_Project) AddSecondaryArtifact(secondaryArtifact IArtifacts) {
	_jsii_.InvokeVoid(
		p,
		"addSecondaryArtifact",
		[]interface{}{secondaryArtifact},
	)
}

// Adds a secondary source to the Project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
//
func (p *jsiiProxy_Project) AddSecondarySource(secondarySource ISource) {
	_jsii_.InvokeVoid(
		p,
		"addSecondarySource",
		[]interface{}{secondarySource},
	)
}

// Add a permission only if there's a policy attached.
func (p *jsiiProxy_Project) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		p,
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
func (p *jsiiProxy_Project) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Returns a source configuration for notification rule.
func (p *jsiiProxy_Project) BindAsNotificationRuleSource(_scope constructs.Construct) *awscodestarnotifications.NotificationRuleSourceConfig {
	var returns *awscodestarnotifications.NotificationRuleSourceConfig

	_jsii_.Invoke(
		p,
		"bindAsNotificationRuleSource",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

// A callback invoked when the given project is added to a CodePipeline.
func (p *jsiiProxy_Project) BindToCodePipeline(_scope constructs.Construct, options *BindToCodePipelineOptions) {
	_jsii_.InvokeVoid(
		p,
		"bindToCodePipeline",
		[]interface{}{_scope, options},
	)
}

// Enable batch builds.
//
// Returns an object contining the batch service role if batch builds
// could be enabled.
func (p *jsiiProxy_Project) EnableBatchBuilds() *BatchBuildConfig {
	var returns *BatchBuildConfig

	_jsii_.Invoke(
		p,
		"enableBatchBuilds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Project) GeneratePhysicalName() *string {
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
func (p *jsiiProxy_Project) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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
func (p *jsiiProxy_Project) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns: a CloudWatch metric associated with this build project.
func (p *jsiiProxy_Project) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Measures the number of builds triggered.
//
// Units: Count
//
// Valid CloudWatch statistics: Sum
func (p *jsiiProxy_Project) MetricBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Measures the duration of all builds over time.
//
// Units: Seconds
//
// Valid CloudWatch statistics: Average (recommended), Maximum, Minimum
func (p *jsiiProxy_Project) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Measures the number of builds that failed because of client error or because of a timeout.
//
// Units: Count
//
// Valid CloudWatch statistics: Sum
func (p *jsiiProxy_Project) MetricFailedBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricFailedBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Measures the number of successful builds.
//
// Units: Count
//
// Valid CloudWatch statistics: Sum
func (p *jsiiProxy_Project) MetricSucceededBuilds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		p,
		"metricSucceededBuilds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Defines a CodeStar Notification rule triggered when the project events emitted by you specified, it very similar to `onEvent` API.
//
// You can also use the methods `notifyOnBuildSucceeded` and
// `notifyOnBuildFailed` to define rules for these specific event emitted.
func (p *jsiiProxy_Project) NotifyOn(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOn",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Defines a CodeStar notification rule which triggers when a build fails.
func (p *jsiiProxy_Project) NotifyOnBuildFailed(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnBuildFailed",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Defines a CodeStar notification rule which triggers when a build completes successfully.
func (p *jsiiProxy_Project) NotifyOnBuildSucceeded(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) awscodestarnotifications.INotificationRule {
	var returns awscodestarnotifications.INotificationRule

	_jsii_.Invoke(
		p,
		"notifyOnBuildSucceeded",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

// Defines an event rule which triggers when a build fails.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
func (p *jsiiProxy_Project) OnBuildFailed(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildFailed",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an event rule which triggers when a build starts.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
func (p *jsiiProxy_Project) OnBuildStarted(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildStarted",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines an event rule which triggers when a build completes successfully.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
func (p *jsiiProxy_Project) OnBuildSucceeded(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onBuildSucceeded",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule triggered when something happens with this project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
//
func (p *jsiiProxy_Project) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule that triggers upon phase change of this build project.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
//
func (p *jsiiProxy_Project) OnPhaseChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onPhaseChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Defines a CloudWatch event rule triggered when the build project state changes.
//
// You can filter specific build status events using an event
// pattern filter on the `build-status` detail field:
//
//     const rule = project.onStateChange('OnBuildStarted', { target });
//     rule.addEventPattern({
//       detail: {
//         'build-status': [
//           "IN_PROGRESS",
//           "SUCCEEDED",
//           "FAILED",
//           "STOPPED"
//         ]
//       }
//     });
//
// You can also use the methods `onBuildFailed` and `onBuildSucceeded` to define rules for
// these specific state changes.
//
// To access fields from the event in the event target input,
// use the static fields on the `StateChangeEvent` class.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html
//
func (p *jsiiProxy_Project) OnStateChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (p *jsiiProxy_Project) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The list of event types for AWS Codebuild.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-buildproject
//
type ProjectNotificationEvents string

const (
	ProjectNotificationEvents_BUILD_FAILED ProjectNotificationEvents = "BUILD_FAILED"
	ProjectNotificationEvents_BUILD_SUCCEEDED ProjectNotificationEvents = "BUILD_SUCCEEDED"
	ProjectNotificationEvents_BUILD_IN_PROGRESS ProjectNotificationEvents = "BUILD_IN_PROGRESS"
	ProjectNotificationEvents_BUILD_STOPPED ProjectNotificationEvents = "BUILD_STOPPED"
	ProjectNotificationEvents_BUILD_PHASE_FAILED ProjectNotificationEvents = "BUILD_PHASE_FAILED"
	ProjectNotificationEvents_BUILD_PHASE_SUCCEEDED ProjectNotificationEvents = "BUILD_PHASE_SUCCEEDED"
)

// Additional options to pass to the notification rule.
//
// TODO: EXAMPLE
//
type ProjectNotifyOnOptions struct {
	// The level of detail to include in the notifications for this resource.
	//
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType awscodestarnotifications.DetailType `json:"detailType" yaml:"detailType"`
	// The status of the notification rule.
	//
	// If the enabled is set to DISABLED, notifications aren't sent for the notification rule.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// The name for the notification rule.
	//
	// Notification rule names must be unique in your AWS account.
	NotificationRuleName *string `json:"notificationRuleName" yaml:"notificationRuleName"`
	// A list of event types associated with this notification rule for CodeBuild Project.
	//
	// For a complete list of event types and IDs, see Notification concepts in the Developer Tools Console User Guide.
	// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#concepts-api
	//
	Events *[]ProjectNotificationEvents `json:"events" yaml:"events"`
}

// TODO: EXAMPLE
//
type ProjectProps struct {
	// Whether to allow the CodeBuild to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// CodeBuild project to connect to network targets.
	//
	// Only used if 'vpc' is supplied.
	AllowAllOutbound *bool `json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see Build Badges Sample
	// in the AWS CodeBuild User Guide.
	Badge *bool `json:"badge" yaml:"badge"`
	// Filename or contents of buildspec in JSON format.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-example
	//
	BuildSpec BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Caching strategy to use.
	Cache Cache `json:"cache" yaml:"cache"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	CheckSecretsInPlainTextEnvVariables *bool `json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Maximum number of concurrent builds.
	//
	// Minimum value is 1 and maximum is account build limit.
	ConcurrentBuildLimit *float64 `json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description of the project.
	//
	// Use the description to identify the purpose
	// of the project.
	Description *string `json:"description" yaml:"description"`
	// Encryption key to use to read and write artifacts.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Build environment to use for the build.
	Environment *BuildEnvironment `json:"environment" yaml:"environment"`
	// Additional environment variables to add to the build environment.
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// An  ProjectFileSystemLocation objects for a CodeBuild build project.
	//
	// A ProjectFileSystemLocation object specifies the identifier, location, mountOptions, mountPoint,
	// and type of a file system created using Amazon Elastic File System.
	FileSystemLocations *[]IFileSystemLocation `json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Add permissions to this project's role to create and use test report groups with name starting with the name of this project.
	//
	// That is the standard report group that gets created when a simple name
	// (in contrast to an ARN)
	// is used in the 'reports' section of the buildspec of this project.
	// This is usually harmless, but you can turn these off if you don't plan on using test
	// reports in this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/test-report-group-naming.html
	//
	GrantReportGroupPermissions *bool `json:"grantReportGroupPermissions" yaml:"grantReportGroupPermissions"`
	// Information about logs for the build project.
	//
	// A project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
	Logging *LoggingOptions `json:"logging" yaml:"logging"`
	// The physical, human-readable name of the CodeBuild Project.
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// The number of minutes after which AWS CodeBuild stops the build if it's still in queue.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	QueuedTimeout awscdk.Duration `json:"queuedTimeout" yaml:"queuedTimeout"`
	// Service Role to assume while running the build.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// What security group to associate with the codebuild project's network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// VPC network to place codebuild network interfaces.
	//
	// Specify this if the codebuild project needs to access resources in a VPC.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// Defines where build artifacts will be stored.
	//
	// Could be: PipelineBuildArtifacts, NoArtifacts and S3Artifacts.
	Artifacts IArtifacts `json:"artifacts" yaml:"artifacts"`
	// The secondary artifacts for the Project.
	//
	// Can also be added after the Project has been created by using the {@link Project#addSecondaryArtifact} method.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	SecondaryArtifacts *[]IArtifacts `json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// The secondary sources for the Project.
	//
	// Can be also added after the Project has been created by using the {@link Project#addSecondarySource} method.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	SecondarySources *[]ISource `json:"secondarySources" yaml:"secondarySources"`
	// The source of the build.
	//
	// *Note*: if {@link NoSource} is given as the source,
	// then you need to provide an explicit `buildSpec`.
	Source ISource `json:"source" yaml:"source"`
}

// The ReportGroup resource class.
//
// TODO: EXAMPLE
//
type ReportGroup interface {
	awscdk.Resource
	IReportGroup
	Env() *awscdk.ResourceEnvironment
	ExportBucket() awss3.IBucket
	Node() constructs.Node
	PhysicalName() *string
	ReportGroupArn() *string
	ReportGroupName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	GrantWrite(identity awsiam.IGrantable) awsiam.Grant
	ToString() *string
}

// The jsii proxy struct for ReportGroup
type jsiiProxy_ReportGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_IReportGroup
}

func (j *jsiiProxy_ReportGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportGroup) ExportBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"exportBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportGroup) ReportGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportGroup) ReportGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReportGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewReportGroup(scope constructs.Construct, id *string, props *ReportGroupProps) ReportGroup {
	_init_.Initialize()

	j := jsiiProxy_ReportGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.ReportGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewReportGroup_Override(r ReportGroup, scope constructs.Construct, id *string, props *ReportGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.ReportGroup",
		[]interface{}{scope, id, props},
		r,
	)
}

// Reference an existing ReportGroup, defined outside of the CDK code, by name.
func ReportGroup_FromReportGroupName(scope constructs.Construct, id *string, reportGroupName *string) IReportGroup {
	_init_.Initialize()

	var returns IReportGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.ReportGroup",
		"fromReportGroupName",
		[]interface{}{scope, id, reportGroupName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func ReportGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.ReportGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ReportGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.ReportGroup",
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
func (r *jsiiProxy_ReportGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		r,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (r *jsiiProxy_ReportGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ReportGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		r,
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
func (r *jsiiProxy_ReportGroup) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Grants the given entity permissions to write (that is, upload reports to) this report group.
func (r *jsiiProxy_ReportGroup) GrantWrite(identity awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		r,
		"grantWrite",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (r *jsiiProxy_ReportGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for {@link ReportGroup}.
//
// TODO: EXAMPLE
//
type ReportGroupProps struct {
	// An optional S3 bucket to export the reports to.
	ExportBucket awss3.IBucket `json:"exportBucket" yaml:"exportBucket"`
	// What to do when this resource is deleted from a stack.
	//
	// As CodeBuild does not allow deleting a ResourceGroup that has reports inside of it,
	// this is set to retain the resource by default.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// The physical name of the report group.
	ReportGroupName *string `json:"reportGroupName" yaml:"reportGroupName"`
	// Whether to output the report files into the export bucket as-is, or create a ZIP from them before doing the export.
	//
	// Ignored if {@link exportBucket} has not been provided.
	ZipExport *bool `json:"zipExport" yaml:"zipExport"`
}

// Construction properties for {@link S3Artifacts}.
//
// TODO: EXAMPLE
//
type S3ArtifactsProps struct {
	// The artifact identifier.
	//
	// This property is required on secondary artifacts.
	Identifier *string `json:"identifier" yaml:"identifier"`
	// The name of the output bucket.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// If this is false, build output will not be encrypted.
	//
	// This is useful if the artifact to publish a static website or sharing content with others
	Encryption *bool `json:"encryption" yaml:"encryption"`
	// Indicates if the build ID should be included in the path.
	//
	// If this is set to true,
	// then the build artifact will be stored in "<path>/<build-id>/<name>".
	IncludeBuildId *bool `json:"includeBuildId" yaml:"includeBuildId"`
	// The name of the build output ZIP file or folder inside the bucket.
	//
	// The full S3 object key will be "<path>/<build-id>/<name>" or
	// "<path>/<name>" depending on whether `includeBuildId` is set to true.
	//
	// If not set, `overrideArtifactName` will be set and the name from the
	// buildspec will be used instead.
	Name *string `json:"name" yaml:"name"`
	// If this is true, all build output will be packaged into a single .zip file. Otherwise, all files will be uploaded to <path>/<name>.
	PackageZip *bool `json:"packageZip" yaml:"packageZip"`
	// The path inside of the bucket for the build output .zip file or folder. If a value is not specified, then build output will be stored at the root of the bucket (or under the <build-id> directory if `includeBuildId` is set to true).
	Path *string `json:"path" yaml:"path"`
}

// Information about logs built to an S3 bucket for a build project.
//
// TODO: EXAMPLE
//
type S3LoggingOptions struct {
	// The S3 Bucket to send logs to.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// The current status of the logs in Amazon CloudWatch Logs for a build project.
	Enabled *bool `json:"enabled" yaml:"enabled"`
	// Encrypt the S3 build log output.
	Encrypted *bool `json:"encrypted" yaml:"encrypted"`
	// The path prefix for S3 logs.
	Prefix *string `json:"prefix" yaml:"prefix"`
}

// Construction properties for {@link S3Source}.
//
// TODO: EXAMPLE
//
type S3SourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `json:"identifier" yaml:"identifier"`
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	Path *string `json:"path" yaml:"path"`
	// The version ID of the object that represents the build input ZIP file to use.
	Version *string `json:"version" yaml:"version"`
}

// Source provider definition for a CodeBuild Project.
//
// TODO: EXAMPLE
//
type Source interface {
	ISource
	BadgeSupported() *bool
	Identifier() *string
	Type() *string
	Bind(_scope constructs.Construct, _project IProject) *SourceConfig
}

// The jsii proxy struct for Source
type jsiiProxy_Source struct {
	jsiiProxy_ISource
}

func (j *jsiiProxy_Source) BadgeSupported() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"badgeSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Source) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Source) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewSource_Override(s Source, props *SourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.Source",
		[]interface{}{props},
		s,
	)
}

func Source_BitBucket(props *BitBucketSourceProps) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"bitBucket",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_CodeCommit(props *CodeCommitSourceProps) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"codeCommit",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_GitHub(props *GitHubSourceProps) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"gitHub",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_GitHubEnterprise(props *GitHubEnterpriseSourceProps) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"gitHubEnterprise",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Source_S3(props *S3SourceProps) ISource {
	_init_.Initialize()

	var returns ISource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.Source",
		"s3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Called by the project when the source is added so that the source can perform binding operations on the source.
//
// For example, it can grant permissions to the
// code build project to read from the S3 bucket.
func (s *jsiiProxy_Source) Bind(_scope constructs.Construct, _project IProject) *SourceConfig {
	var returns *SourceConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _project},
		&returns,
	)

	return returns
}

// The type returned from {@link ISource#bind}.
//
// TODO: EXAMPLE
//
type SourceConfig struct {
	SourceProperty *CfnProject_SourceProperty `json:"sourceProperty" yaml:"sourceProperty"`
	BuildTriggers *CfnProject_ProjectTriggersProperty `json:"buildTriggers" yaml:"buildTriggers"`
	// `AWS::CodeBuild::Project.SourceVersion`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html#cfn-codebuild-project-sourceversion
	//
	SourceVersion *string `json:"sourceVersion" yaml:"sourceVersion"`
}

// Properties common to all Source classes.
//
// TODO: EXAMPLE
//
type SourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `json:"identifier" yaml:"identifier"`
}

// Event fields for the CodeBuild "state change" event.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-build-notifications.html#sample-build-notifications-ref
//
type StateChangeEvent interface {
}

// The jsii proxy struct for StateChangeEvent
type jsiiProxy_StateChangeEvent struct {
	_ byte // padding
}

func StateChangeEvent_BuildId() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"buildId",
		&returns,
	)
	return returns
}

func StateChangeEvent_BuildStatus() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"buildStatus",
		&returns,
	)
	return returns
}

func StateChangeEvent_CurrentPhase() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"currentPhase",
		&returns,
	)
	return returns
}

func StateChangeEvent_ProjectName() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.StateChangeEvent",
		"projectName",
		&returns,
	)
	return returns
}

// Permissions Boundary for a CodeBuild Project running untrusted code.
//
// This class is a Policy, intended to be used as a Permissions Boundary
// for a CodeBuild project. It allows most of the actions necessary to run
// the CodeBuild project, but disallows reading from Parameter Store
// and Secrets Manager.
//
// Use this when your CodeBuild project is running untrusted code (for
// example, if you are using one to automatically build Pull Requests
// that anyone can submit), and you want to prevent your future self
// from accidentally exposing Secrets to this build.
//
// (The reason you might want to do this is because otherwise anyone
// who can submit a Pull Request to your project can write a script
// to email those secrets to themselves).
//
// TODO: EXAMPLE
//
type UntrustedCodeBoundaryPolicy interface {
	awsiam.ManagedPolicy
	Description() *string
	Document() awsiam.PolicyDocument
	Env() *awscdk.ResourceEnvironment
	ManagedPolicyArn() *string
	ManagedPolicyName() *string
	Node() constructs.Node
	Path() *string
	PhysicalName() *string
	Stack() awscdk.Stack
	AddStatements(statement ...awsiam.PolicyStatement)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AttachToGroup(group awsiam.IGroup)
	AttachToRole(role awsiam.IRole)
	AttachToUser(user awsiam.IUser)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for UntrustedCodeBoundaryPolicy
type jsiiProxy_UntrustedCodeBoundaryPolicy struct {
	internal.Type__awsiamManagedPolicy
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) Document() awsiam.PolicyDocument {
	var returns awsiam.PolicyDocument
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) ManagedPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) ManagedPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UntrustedCodeBoundaryPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewUntrustedCodeBoundaryPolicy(scope constructs.Construct, id *string, props *UntrustedCodeBoundaryPolicyProps) UntrustedCodeBoundaryPolicy {
	_init_.Initialize()

	j := jsiiProxy_UntrustedCodeBoundaryPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.UntrustedCodeBoundaryPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewUntrustedCodeBoundaryPolicy_Override(u UntrustedCodeBoundaryPolicy, scope constructs.Construct, id *string, props *UntrustedCodeBoundaryPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codebuild.UntrustedCodeBoundaryPolicy",
		[]interface{}{scope, id, props},
		u,
	)
}

// Import a managed policy from one of the policies that AWS manages.
//
// For this managed policy, you only need to know the name to be able to use it.
//
// Some managed policy names start with "service-role/", some start with
// "job-function/", and some don't start with anything. Include the
// prefix when constructing this object.
func UntrustedCodeBoundaryPolicy_FromAwsManagedPolicyName(managedPolicyName *string) awsiam.IManagedPolicy {
	_init_.Initialize()

	var returns awsiam.IManagedPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.UntrustedCodeBoundaryPolicy",
		"fromAwsManagedPolicyName",
		[]interface{}{managedPolicyName},
		&returns,
	)

	return returns
}

// Import an external managed policy by ARN.
//
// For this managed policy, you only need to know the ARN to be able to use it.
// This can be useful if you got the ARN from a CloudFormation Export.
//
// If the imported Managed Policy ARN is a Token (such as a
// `CfnParameter.valueAsString` or a `Fn.importValue()`) *and* the referenced
// managed policy has a `path` (like `arn:...:policy/AdminPolicy/AdminAllow`), the
// `managedPolicyName` property will not resolve to the correct value. Instead it
// will resolve to the first path component. We unfortunately cannot express
// the correct calculation of the full path name as a CloudFormation
// expression. In this scenario the Managed Policy ARN should be supplied without the
// `path` in order to resolve the correct managed policy resource.
func UntrustedCodeBoundaryPolicy_FromManagedPolicyArn(scope constructs.Construct, id *string, managedPolicyArn *string) awsiam.IManagedPolicy {
	_init_.Initialize()

	var returns awsiam.IManagedPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.UntrustedCodeBoundaryPolicy",
		"fromManagedPolicyArn",
		[]interface{}{scope, id, managedPolicyArn},
		&returns,
	)

	return returns
}

// Import a customer managed policy from the managedPolicyName.
//
// For this managed policy, you only need to know the name to be able to use it.
func UntrustedCodeBoundaryPolicy_FromManagedPolicyName(scope constructs.Construct, id *string, managedPolicyName *string) awsiam.IManagedPolicy {
	_init_.Initialize()

	var returns awsiam.IManagedPolicy

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.UntrustedCodeBoundaryPolicy",
		"fromManagedPolicyName",
		[]interface{}{scope, id, managedPolicyName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func UntrustedCodeBoundaryPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.UntrustedCodeBoundaryPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func UntrustedCodeBoundaryPolicy_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.UntrustedCodeBoundaryPolicy",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Adds a statement to the policy document.
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) AddStatements(statement ...awsiam.PolicyStatement) {
	args := []interface{}{}
	for _, a := range statement {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		u,
		"addStatements",
		args,
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
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		u,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Attaches this policy to a group.
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) AttachToGroup(group awsiam.IGroup) {
	_jsii_.InvokeVoid(
		u,
		"attachToGroup",
		[]interface{}{group},
	)
}

// Attaches this policy to a role.
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) AttachToRole(role awsiam.IRole) {
	_jsii_.InvokeVoid(
		u,
		"attachToRole",
		[]interface{}{role},
	)
}

// Attaches this policy to a user.
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) AttachToUser(user awsiam.IUser) {
	_jsii_.InvokeVoid(
		u,
		"attachToUser",
		[]interface{}{user},
	)
}

func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		u,
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
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		u,
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
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (u *jsiiProxy_UntrustedCodeBoundaryPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for UntrustedCodeBoundaryPolicy.
//
// TODO: EXAMPLE
//
type UntrustedCodeBoundaryPolicyProps struct {
	// Additional statements to add to the default set of statements.
	AdditionalStatements *[]awsiam.PolicyStatement `json:"additionalStatements" yaml:"additionalStatements"`
	// The name of the managed policy.
	ManagedPolicyName *string `json:"managedPolicyName" yaml:"managedPolicyName"`
}

// A CodeBuild image running Windows.
//
// This class has a bunch of public constants that represent the most popular images.
//
// You can also specify a custom image using one of the static methods:
//
// - WindowsBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }, imageType])
// - WindowsBuildImage.fromEcrRepository(repo[, tag, imageType])
// - WindowsBuildImage.fromAsset(parent, id, props, [, imageType])
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
type WindowsBuildImage interface {
	IBuildImage
	DefaultComputeType() ComputeType
	ImageId() *string
	ImagePullPrincipalType() ImagePullPrincipalType
	Repository() awsecr.IRepository
	SecretsManagerCredentials() awssecretsmanager.ISecret
	Type() *string
	RunScriptBuildspec(entrypoint *string) BuildSpec
	Validate(buildEnvironment *BuildEnvironment) *[]*string
}

// The jsii proxy struct for WindowsBuildImage
type jsiiProxy_WindowsBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_WindowsBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Uses an Docker image asset as a Windows build image.
func WindowsBuildImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps, imageType WindowsImageType) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"fromAsset",
		[]interface{}{scope, id, props, imageType},
		&returns,
	)

	return returns
}

// Returns: a Windows build image from a Docker Hub image.
func WindowsBuildImage_FromDockerRegistry(name *string, options *DockerImageOptions, imageType WindowsImageType) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"fromDockerRegistry",
		[]interface{}{name, options, imageType},
		&returns,
	)

	return returns
}

// Returns: A Windows build image from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
func WindowsBuildImage_FromEcrRepository(repository awsecr.IRepository, tag *string, imageType WindowsImageType) IBuildImage {
	_init_.Initialize()

	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tag, imageType},
		&returns,
	)

	return returns
}

func WindowsBuildImage_WIN_SERVER_CORE_2019_BASE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"WIN_SERVER_CORE_2019_BASE",
		&returns,
	)
	return returns
}

func WindowsBuildImage_WINDOWS_BASE_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.WindowsBuildImage",
		"WINDOWS_BASE_2_0",
		&returns,
	)
	return returns
}

// Make a buildspec to run the indicated script.
func (w *jsiiProxy_WindowsBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	var returns BuildSpec

	_jsii_.Invoke(
		w,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

// Allows the image a chance to validate whether the passed configuration is correct.
func (w *jsiiProxy_WindowsBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

// Environment type for Windows Docker images.
//
// TODO: EXAMPLE
//
type WindowsImageType string

const (
	WindowsImageType_STANDARD WindowsImageType = "STANDARD"
	WindowsImageType_SERVER_2019 WindowsImageType = "SERVER_2019"
)

