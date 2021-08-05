package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipelineactions"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Options for addManualApproval.
// Experimental.
type AddManualApprovalOptions struct {
	// The name of the manual approval action.
	// Experimental.
	ActionName *string `json:"actionName"`
	// The runOrder for this action.
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
}

// Additional options for adding a stack deployment.
// Experimental.
type AddStackOptions struct {
	// Base runorder.
	// Experimental.
	ExecuteRunOrder *float64 `json:"executeRunOrder"`
	// Base runorder.
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
}

// Options for adding an application stage to a pipeline.
// Experimental.
type AddStageOptions struct {
	// Runs a `cdk diff --security-only --fail` to pause the pipeline if there are any security changes.
	//
	// If the stage is configured with `confirmBroadeningPermissions` enabled, you can use this
	// property to override the stage configuration. For example, Pipeline Stage
	// "Prod" has confirmBroadeningPermissions enabled, with applications "A", "B", "C". All three
	// applications will run a security check, but if we want to disable the one for "C",
	// we run `stage.addApplication(C, { confirmBroadeningPermissions: false })` to override the pipeline
	// stage behavior.
	//
	// Adds 1 to the run order space.
	// Experimental.
	ConfirmBroadeningPermissions *bool `json:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when the security check registers changes within the application.
	// Experimental.
	SecurityNotificationTopic awssns.ITopic `json:"securityNotificationTopic"`
	// Add room for extra actions.
	//
	// You can use this to make extra room in the runOrder sequence between the
	// changeset 'prepare' and 'execute' actions and insert your own actions there.
	// Experimental.
	ExtraRunOrderSpace *float64 `json:"extraRunOrderSpace"`
	// Add manual approvals before executing change sets.
	//
	// This gives humans the opportunity to confirm the change set looks alright
	// before deploying it.
	// Experimental.
	ManualApprovals *bool `json:"manualApprovals"`
}

// Options to pass to `addStage`.
// Experimental.
type AddStageOpts struct {
	// Additional steps to run after all of the stacks in the stage.
	// Experimental.
	Post *[]Step `json:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Experimental.
	Pre *[]Step `json:"pre"`
}

// Specification of an additional artifact to generate.
// Experimental.
type AdditionalArtifact struct {
	// Artifact to represent the build directory in the pipeline.
	// Experimental.
	Artifact awscodepipeline.Artifact `json:"artifact"`
	// Directory to be packaged.
	// Experimental.
	Directory *string `json:"directory"`
}

// Translate FileSets to CodePipeline Artifacts.
// Experimental.
type ArtifactMap interface {
	ToCodePipeline(x FileSet) awscodepipeline.Artifact
}

// The jsii proxy struct for ArtifactMap
type jsiiProxy_ArtifactMap struct {
	_ byte // padding
}

// Experimental.
func NewArtifactMap() ArtifactMap {
	_init_.Initialize()

	j := jsiiProxy_ArtifactMap{}

	_jsii_.Create(
		"monocdk.pipelines.ArtifactMap",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewArtifactMap_Override(a ArtifactMap) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.ArtifactMap",
		nil, // no parameters
		a,
	)
}

// Return the matching CodePipeline artifact for a FileSet.
// Experimental.
func (a *jsiiProxy_ArtifactMap) ToCodePipeline(x FileSet) awscodepipeline.Artifact {
	var returns awscodepipeline.Artifact

	_jsii_.Invoke(
		a,
		"toCodePipeline",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Instructions to publish certain assets.
// Experimental.
type AssetPublishingCommand struct {
	// Asset identifier.
	// Experimental.
	AssetId *string `json:"assetId"`
	// Asset manifest path.
	// Experimental.
	AssetManifestPath *string `json:"assetManifestPath"`
	// ARN of the IAM Role used to publish this asset.
	// Experimental.
	AssetPublishingRoleArn *string `json:"assetPublishingRoleArn"`
	// Asset selector to pass to `cdk-assets`.
	// Experimental.
	AssetSelector *string `json:"assetSelector"`
	// Type of asset to publish.
	// Experimental.
	AssetType AssetType `json:"assetType"`
}

// Type of the asset that is being published.
// Experimental.
type AssetType string

const (
	AssetType_FILE AssetType = "FILE"
	AssetType_DOCKER_IMAGE AssetType = "DOCKER_IMAGE"
)

// Base options for a pipelines stage.
// Experimental.
type BaseStageOptions struct {
	// Runs a `cdk diff --security-only --fail` to pause the pipeline if there are any security changes.
	//
	// If the stage is configured with `confirmBroadeningPermissions` enabled, you can use this
	// property to override the stage configuration. For example, Pipeline Stage
	// "Prod" has confirmBroadeningPermissions enabled, with applications "A", "B", "C". All three
	// applications will run a security check, but if we want to disable the one for "C",
	// we run `stage.addApplication(C, { confirmBroadeningPermissions: false })` to override the pipeline
	// stage behavior.
	//
	// Adds 1 to the run order space.
	// Experimental.
	ConfirmBroadeningPermissions *bool `json:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when the security check registers changes within the application.
	// Experimental.
	SecurityNotificationTopic awssns.ITopic `json:"securityNotificationTopic"`
}

// A Pipeline to deploy CDK apps.
//
// Defines an AWS CodePipeline-based Pipeline to deploy CDK applications.
//
// Automatically manages the following:
//
// - Stack dependency order.
// - Asset publishing.
// - Keeping the pipeline up-to-date as the CDK apps change.
// - Using stack outputs later on in the pipeline.
// Experimental.
type CdkPipeline interface {
	awscdk.Construct
	CodePipeline() awscodepipeline.Pipeline
	Node() awscdk.ConstructNode
	AddApplicationStage(appStage awscdk.Stage, options *AddStageOptions) CdkStage
	AddStage(stageName *string, options *BaseStageOptions) CdkStage
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	StackOutput(cfnOutput awscdk.CfnOutput) StackOutput
	Stage(stageName *string) awscodepipeline.IStage
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for CdkPipeline
type jsiiProxy_CdkPipeline struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_CdkPipeline) CodePipeline() awscodepipeline.Pipeline {
	var returns awscodepipeline.Pipeline
	_jsii_.Get(
		j,
		"codePipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkPipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdkPipeline(scope constructs.Construct, id *string, props *CdkPipelineProps) CdkPipeline {
	_init_.Initialize()

	j := jsiiProxy_CdkPipeline{}

	_jsii_.Create(
		"monocdk.pipelines.CdkPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkPipeline_Override(c CdkPipeline, scope constructs.Construct, id *string, props *CdkPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CdkPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CdkPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CdkPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add pipeline stage that will deploy the given application stage.
//
// The application construct should subclass `Stage` and can contain any
// number of `Stacks` inside it that may have dependency relationships
// on one another.
//
// All stacks in the application will be deployed in the appropriate order,
// and all assets found in the application will be added to the asset
// publishing stage.
// Experimental.
func (c *jsiiProxy_CdkPipeline) AddApplicationStage(appStage awscdk.Stage, options *AddStageOptions) CdkStage {
	var returns CdkStage

	_jsii_.Invoke(
		c,
		"addApplicationStage",
		[]interface{}{appStage, options},
		&returns,
	)

	return returns
}

// Add a new, empty stage to the pipeline.
//
// Prefer to use `addApplicationStage` if you are intended to deploy a CDK
// application, but you can use this method if you want to add other kinds of
// Actions to a pipeline.
// Experimental.
func (c *jsiiProxy_CdkPipeline) AddStage(stageName *string, options *BaseStageOptions) CdkStage {
	var returns CdkStage

	_jsii_.Invoke(
		c,
		"addStage",
		[]interface{}{stageName, options},
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
func (c *jsiiProxy_CdkPipeline) OnPrepare() {
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
func (c *jsiiProxy_CdkPipeline) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CdkPipeline) OnValidate() *[]*string {
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
func (c *jsiiProxy_CdkPipeline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Get the StackOutput object that holds this CfnOutput's value in this pipeline.
//
// `StackOutput` can be used in validation actions later in the pipeline.
// Experimental.
func (c *jsiiProxy_CdkPipeline) StackOutput(cfnOutput awscdk.CfnOutput) StackOutput {
	var returns StackOutput

	_jsii_.Invoke(
		c,
		"stackOutput",
		[]interface{}{cfnOutput},
		&returns,
	)

	return returns
}

// Access one of the pipeline's stages by stage name.
//
// You can use this to add more Actions to a stage.
// Experimental.
func (c *jsiiProxy_CdkPipeline) Stage(stageName *string) awscodepipeline.IStage {
	var returns awscodepipeline.IStage

	_jsii_.Invoke(
		c,
		"stage",
		[]interface{}{stageName},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CdkPipeline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CdkPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate that we don't have any stacks violating dependency order in the pipeline.
//
// Our own convenience methods will never generate a pipeline that does that (although
// this is a nice verification), but a user can also add the stacks by hand.
// Experimental.
func (c *jsiiProxy_CdkPipeline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a CdkPipeline.
// Experimental.
type CdkPipelineProps struct {
	// The artifact you have defined to be the artifact to hold the cloudAssemblyArtifact for the synth action.
	// Experimental.
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact"`
	// Custom BuildSpec that is merged with generated one (for asset publishing actions).
	// Experimental.
	AssetBuildSpec awscodebuild.BuildSpec `json:"assetBuildSpec"`
	// Additional commands to run before installing cdk-assets during the asset publishing step Use this to setup proxies or npm mirrors.
	// Experimental.
	AssetPreInstallCommands *[]*string `json:"assetPreInstallCommands"`
	// CDK CLI version to use in pipeline.
	//
	// Some Actions in the pipeline will download and run a version of the CDK
	// CLI. Specify the version here.
	// Experimental.
	CdkCliVersion *string `json:"cdkCliVersion"`
	// Existing CodePipeline to add deployment stages to.
	//
	// Use this if you want more control over the CodePipeline that gets created.
	// You can choose to not pass this value, in which case a new CodePipeline is
	// created with default settings.
	//
	// If you pass an existing CodePipeline, it should should have been created
	// with `restartExecutionOnUpdate: true`.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	CodePipeline awscodepipeline.Pipeline `json:"codePipeline"`
	// Create KMS keys for cross-account deployments.
	//
	// This controls whether the pipeline is enabled for cross-account deployments.
	//
	// Can only be set if `codePipeline` is not set.
	//
	// By default cross-account deployments are enabled, but this feature requires
	// that KMS Customer Master Keys are created which have a cost of $1/month.
	//
	// If you do not need cross-account deployments, you can set this to `false` to
	// not create those keys and save on that cost (the artifact bucket will be
	// encrypted with an AWS-managed key). However, cross-account deployments will
	// no longer be possible.
	// Experimental.
	CrossAccountKeys *bool `json:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
	// Experimental.
	DockerCredentials *[]DockerCredential `json:"dockerCredentials"`
	// Name of the pipeline.
	//
	// Can only be set if `codePipeline` is not set.
	// Experimental.
	PipelineName *string `json:"pipelineName"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	// Experimental.
	SelfMutating *bool `json:"selfMutating"`
	// Custom BuildSpec that is merged with generated one (for self-mutation stage).
	// Experimental.
	SelfMutationBuildSpec awscodebuild.BuildSpec `json:"selfMutationBuildSpec"`
	// Whether this pipeline creates one asset upload action per asset type or one asset upload per asset.
	// Experimental.
	SinglePublisherPerType *bool `json:"singlePublisherPerType"`
	// The CodePipeline action used to retrieve the CDK app's source.
	// Experimental.
	SourceAction awscodepipeline.IAction `json:"sourceAction"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// Whether the pipeline needs to build Docker images in the UpdatePipeline stage.
	//
	// If the UpdatePipeline stage tries to build a Docker image and this flag is not
	// set to `true`, the build step will run in non-privileged mode and consequently
	// will fail with a message like:
	//
	// > Cannot connect to the Docker daemon at unix:///var/run/docker.sock.
	// > Is the docker daemon running?
	//
	// This flag has an effect only if `selfMutating` is also `true`.
	// Experimental.
	SupportDockerAssets *bool `json:"supportDockerAssets"`
	// The CodePipeline action build and synthesis step of the CDK app.
	// Experimental.
	SynthAction awscodepipeline.IAction `json:"synthAction"`
	// The VPC where to execute the CdkPipeline actions.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Options for the 'fromStackArtifact' operation.
// Experimental.
type CdkStackActionFromArtifactOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Experimental.
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput"`
	// Base name of the action.
	// Experimental.
	BaseActionName *string `json:"baseActionName"`
	// Name of the change set to create and deploy.
	// Experimental.
	ChangeSetName *string `json:"changeSetName"`
	// Run order for the Execute action.
	// Experimental.
	ExecuteRunOrder *float64 `json:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// Filename in output to write Stack outputs to.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
	// Run order for the Prepare action.
	// Experimental.
	PrepareRunOrder *float64 `json:"prepareRunOrder"`
	// The name of the stack that should be created/updated.
	// Experimental.
	StackName *string `json:"stackName"`
}

// Stage in a CdkPipeline.
//
// You don't need to instantiate this class directly. Use
// `cdkPipeline.addStage()` instead.
// Experimental.
type CdkStage interface {
	awscdk.Construct
	Node() awscdk.ConstructNode
	AddActions(actions ...awscodepipeline.IAction)
	AddApplication(appStage awscdk.Stage, options *AddStageOptions)
	AddManualApprovalAction(options *AddManualApprovalOptions)
	AddStackArtifactDeployment(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions)
	DeploysStack(artifactId *string) *bool
	NextSequentialRunOrder(count *float64) *float64
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for CdkStage
type jsiiProxy_CdkStage struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_CdkStage) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewCdkStage(scope constructs.Construct, id *string, props *CdkStageProps) CdkStage {
	_init_.Initialize()

	j := jsiiProxy_CdkStage{}

	_jsii_.Create(
		"monocdk.pipelines.CdkStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkStage_Override(c CdkStage, scope constructs.Construct, id *string, props *CdkStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CdkStage",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CdkStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CdkStage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add one or more CodePipeline Actions.
//
// You need to make sure it is created with the right runOrder. Call `nextSequentialRunOrder()`
// for every action to get actions to execute in sequence.
// Experimental.
func (c *jsiiProxy_CdkStage) AddActions(actions ...awscodepipeline.IAction) {
	args := []interface{}{}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addActions",
		args,
	)
}

// Add all stacks in the application Stage to this stage.
//
// The application construct should subclass `Stage` and can contain any
// number of `Stacks` inside it that may have dependency relationships
// on one another.
//
// All stacks in the application will be deployed in the appropriate order,
// and all assets found in the application will be added to the asset
// publishing stage.
// Experimental.
func (c *jsiiProxy_CdkStage) AddApplication(appStage awscdk.Stage, options *AddStageOptions) {
	_jsii_.InvokeVoid(
		c,
		"addApplication",
		[]interface{}{appStage, options},
	)
}

// Add a manual approval action.
//
// If you need more flexibility than what this method offers,
// use `addAction` with a `ManualApprovalAction`.
// Experimental.
func (c *jsiiProxy_CdkStage) AddManualApprovalAction(options *AddManualApprovalOptions) {
	_jsii_.InvokeVoid(
		c,
		"addManualApprovalAction",
		[]interface{}{options},
	)
}

// Add a deployment action based on a stack artifact.
// Experimental.
func (c *jsiiProxy_CdkStage) AddStackArtifactDeployment(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions) {
	_jsii_.InvokeVoid(
		c,
		"addStackArtifactDeployment",
		[]interface{}{stackArtifact, options},
	)
}

// Whether this Stage contains an action to deploy the given stack, identified by its artifact ID.
// Experimental.
func (c *jsiiProxy_CdkStage) DeploysStack(artifactId *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"deploysStack",
		[]interface{}{artifactId},
		&returns,
	)

	return returns
}

// Return the runOrder number necessary to run the next Action in sequence with the rest.
//
// FIXME: This is here because Actions are immutable and can't be reordered
// after creation, nor is there a way to specify relative priorities, which
// is a limitation that we should take away in the base library.
// Experimental.
func (c *jsiiProxy_CdkStage) NextSequentialRunOrder(count *float64) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"nextSequentialRunOrder",
		[]interface{}{count},
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
func (c *jsiiProxy_CdkStage) OnPrepare() {
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
func (c *jsiiProxy_CdkStage) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CdkStage) OnValidate() *[]*string {
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
func (c *jsiiProxy_CdkStage) Prepare() {
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
func (c *jsiiProxy_CdkStage) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CdkStage) ToString() *string {
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
func (c *jsiiProxy_CdkStage) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a CdkStage.
// Experimental.
type CdkStageProps struct {
	// The CodePipeline Artifact with the Cloud Assembly.
	// Experimental.
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact"`
	// Features the Stage needs from its environment.
	// Experimental.
	Host IStageHost `json:"host"`
	// The underlying Pipeline Stage associated with thisCdkStage.
	// Experimental.
	PipelineStage awscodepipeline.IStage `json:"pipelineStage"`
	// Name of the stage that should be created.
	// Experimental.
	StageName *string `json:"stageName"`
	// Run a security check before every application prepare/deploy actions.
	//
	// Note: Stage level security check can be overriden per application as follows:
	//    `stage.addApplication(app, { confirmBroadeningPermissions: false })`
	// Experimental.
	ConfirmBroadeningPermissions *bool `json:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when any security check registers changes within a application.
	//
	// Note: The Stage Notification Topic can be overriden per application as follows:
	//    `stage.addApplication(app, { securityNotificationTopic: newTopic })`
	// Experimental.
	SecurityNotificationTopic awssns.ITopic `json:"securityNotificationTopic"`
}

// Options for customizing a single CodeBuild project.
// Experimental.
type CodeBuildOptions struct {
	// Partial build environment, will be combined with other build environments that apply.
	// Experimental.
	BuildEnvironment *awscodebuild.BuildEnvironment `json:"buildEnvironment"`
	// Partial buildspec, will be combined with other buildspecs that apply.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	// Experimental.
	PartialBuildSpec awscodebuild.BuildSpec `json:"partialBuildSpec"`
	// Policy statements to add to role.
	// Experimental.
	RolePolicy *[]awsiam.PolicyStatement `json:"rolePolicy"`
	// Which security group(s) to associate with the project network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where to create the CodeBuild network interfaces in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Run a script as a CodeBuild Project.
// Experimental.
type CodeBuildStep interface {
	ShellStep
	BuildEnvironment() *awscodebuild.BuildEnvironment
	Commands() *[]*string
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Env() *map[string]*string
	EnvFromCfnOutputs() *map[string]StackOutputReference
	GrantPrincipal() awsiam.IPrincipal
	Id() *string
	Inputs() *[]*FileSetLocation
	InstallCommands() *[]*string
	IsSource() *bool
	Outputs() *[]*FileSetLocation
	PartialBuildSpec() awscodebuild.BuildSpec
	PrimaryOutput() FileSet
	Project() awscodebuild.IProject
	ProjectName() *string
	Role() awsiam.IRole
	RolePolicyStatements() *[]awsiam.PolicyStatement
	SecurityGroups() *[]awsec2.ISecurityGroup
	SubnetSelection() *awsec2.SubnetSelection
	Vpc() awsec2.IVpc
	AddDependencyFileSet(fs FileSet)
	AddOutputDirectory(directory *string) FileSet
	ConfigurePrimaryOutput(fs FileSet)
	PrimaryOutputDirectory(directory *string) FileSet
	ToString() *string
}

// The jsii proxy struct for CodeBuildStep
type jsiiProxy_CodeBuildStep struct {
	jsiiProxy_ShellStep
}

func (j *jsiiProxy_CodeBuildStep) BuildEnvironment() *awscodebuild.BuildEnvironment {
	var returns *awscodebuild.BuildEnvironment
	_jsii_.Get(
		j,
		"buildEnvironment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Commands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Dependencies() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) DependencyFileSets() *[]FileSet {
	var returns *[]FileSet
	_jsii_.Get(
		j,
		"dependencyFileSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) EnvFromCfnOutputs() *map[string]StackOutputReference {
	var returns *map[string]StackOutputReference
	_jsii_.Get(
		j,
		"envFromCfnOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Inputs() *[]*FileSetLocation {
	var returns *[]*FileSetLocation
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) InstallCommands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"installCommands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) IsSource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Outputs() *[]*FileSetLocation {
	var returns *[]*FileSetLocation
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) PartialBuildSpec() awscodebuild.BuildSpec {
	var returns awscodebuild.BuildSpec
	_jsii_.Get(
		j,
		"partialBuildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Project() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) RolePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"rolePolicyStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) SubnetSelection() *awsec2.SubnetSelection {
	var returns *awsec2.SubnetSelection
	_jsii_.Get(
		j,
		"subnetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStep) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodeBuildStep(id *string, props *CodeBuildStepProps) CodeBuildStep {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildStep{}

	_jsii_.Create(
		"monocdk.pipelines.CodeBuildStep",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildStep_Override(c CodeBuildStep, id *string, props *CodeBuildStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CodeBuildStep",
		[]interface{}{id, props},
		c,
	)
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
// Experimental.
func (c *jsiiProxy_CodeBuildStep) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Add an additional output FileSet based on a directory.
//
// After running the script, the contents of the given directory
// will be exported as a `FileSet`. Use the `FileSet` as the
// input to another step.
//
// Multiple calls with the exact same directory name string (not normalized)
// will return the same FileSet.
// Experimental.
func (c *jsiiProxy_CodeBuildStep) AddOutputDirectory(directory *string) FileSet {
	var returns FileSet

	_jsii_.Invoke(
		c,
		"addOutputDirectory",
		[]interface{}{directory},
		&returns,
	)

	return returns
}

// Configure the given FileSet as the primary output of this step.
// Experimental.
func (c *jsiiProxy_CodeBuildStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Configure the given output directory as primary output.
//
// If no primary output has been configured yet, this directory
// will become the primary output of this ShellStep, otherwise this
// method will throw if the given directory is different than the
// currently configured primary output directory.
// Experimental.
func (c *jsiiProxy_CodeBuildStep) PrimaryOutputDirectory(directory *string) FileSet {
	var returns FileSet

	_jsii_.Invoke(
		c,
		"primaryOutputDirectory",
		[]interface{}{directory},
		&returns,
	)

	return returns
}

// Return a string representation of this Step.
// Experimental.
func (c *jsiiProxy_CodeBuildStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction props for a CodeBuildStep.
// Experimental.
type CodeBuildStepProps struct {
	// Commands to run.
	// Experimental.
	Commands *[]*string `json:"commands"`
	// Additional FileSets to put in other directories.
	//
	// Specifies a mapping from directory name to FileSets. During the
	// script execution, the FileSets will be available in the directories
	// indicated.
	//
	// The directory names may be relative. For example, you can put
	// the main input and an additional input side-by-side with the
	// following configuration:
	//
	// ```ts
	// const script = new ShellStep('MainScript', {
	//    // ...
	//    input: MyEngineSource.gitHub('org/source1'),
	//    additionalInputs: {
	//      '../siblingdir': MyEngineSource.gitHub('org/source2'),
	//    }
	// });
	// ```
	// Experimental.
	AdditionalInputs *map[string]IFileSetProducer `json:"additionalInputs"`
	// Environment variables to set.
	// Experimental.
	Env *map[string]*string `json:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	// Experimental.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `json:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	// Experimental.
	Input IFileSetProducer `json:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Experimental.
	InstallCommands *[]*string `json:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	// Experimental.
	PrimaryOutputDirectory *string `json:"primaryOutputDirectory"`
	// Changes to environment.
	//
	// This environment will be combined with the pipeline's default
	// environment.
	// Experimental.
	BuildEnvironment *awscodebuild.BuildEnvironment `json:"buildEnvironment"`
	// Additional configuration that can only be configured via BuildSpec.
	//
	// You should not use this to specify output artifacts; those
	// should be supplied via the other properties of this class, otherwise
	// CDK Pipelines won't be able to inspect the artifacts.
	//
	// Set the `commands` to an empty array if you want to fully specify
	// the BuildSpec using this field.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	// Experimental.
	PartialBuildSpec awscodebuild.BuildSpec `json:"partialBuildSpec"`
	// Name for the generated CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
	// Custom execution role to be used for the CodeBuild project.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Experimental.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Configuration options for a CodeCommit source.
// Experimental.
type CodeCommitSourceOptions struct {
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html
	//
	// Experimental.
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput"`
	// Role to be used by on commit event rule.
	//
	// Used only when trigger value is CodeCommitTrigger.EVENTS.
	// Experimental.
	EventRole awsiam.IRole `json:"eventRole"`
	// How should CodePipeline detect source changes for this Action.
	// Experimental.
	Trigger awscodepipelineactions.CodeCommitTrigger `json:"trigger"`
}

// A CDK Pipeline that uses CodePipeline to deploy CDK apps.
//
// This is a `Pipeline` with its `engine` property set to
// `CodePipelineEngine`, and exists for nicer ergonomics for
// users that don't need to switch out engines.
// Experimental.
type CodePipeline interface {
	PipelineBase
	CloudAssemblyFileSet() FileSet
	Node() awscdk.ConstructNode
	Pipeline() awscodepipeline.Pipeline
	Synth() IFileSetProducer
	SynthProject() awscodebuild.IProject
	Waves() *[]Wave
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
	AddWave(id *string, options *WaveOptions) Wave
	BuildPipeline()
	DoBuildPipeline()
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for CodePipeline
type jsiiProxy_CodePipeline struct {
	jsiiProxy_PipelineBase
}

func (j *jsiiProxy_CodePipeline) CloudAssemblyFileSet() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"cloudAssemblyFileSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipeline) Pipeline() awscodepipeline.Pipeline {
	var returns awscodepipeline.Pipeline
	_jsii_.Get(
		j,
		"pipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipeline) Synth() IFileSetProducer {
	var returns IFileSetProducer
	_jsii_.Get(
		j,
		"synth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipeline) SynthProject() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"synthProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipeline) Waves() *[]Wave {
	var returns *[]Wave
	_jsii_.Get(
		j,
		"waves",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodePipeline(scope constructs.Construct, id *string, props *CodePipelineProps) CodePipeline {
	_init_.Initialize()

	j := jsiiProxy_CodePipeline{}

	_jsii_.Create(
		"monocdk.pipelines.CodePipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodePipeline_Override(c CodePipeline, scope constructs.Construct, id *string, props *CodePipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CodePipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func CodePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Deploy a single Stage by itself.
//
// Add a Stage to the pipeline, to be deployed in sequence with other
// Stages added to the pipeline. All Stacks in the stage will be deployed
// in an order automatically determined by their relative dependencies.
// Experimental.
func (c *jsiiProxy_CodePipeline) AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment {
	var returns StageDeployment

	_jsii_.Invoke(
		c,
		"addStage",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

// Add a Wave to the pipeline, for deploying multiple Stages in parallel.
//
// Use the return object of this method to deploy multiple stages in parallel.
//
// Example:
//
// ```ts
// const wave = pipeline.addWave('MyWave');
// wave.addStage(new MyStage('Stage1', ...));
// wave.addStage(new MyStage('Stage2', ...));
// ```
// Experimental.
func (c *jsiiProxy_CodePipeline) AddWave(id *string, options *WaveOptions) Wave {
	var returns Wave

	_jsii_.Invoke(
		c,
		"addWave",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Send the current pipeline definition to the engine, and construct the pipeline.
//
// It is not possible to modify the pipeline after calling this method.
// Experimental.
func (c *jsiiProxy_CodePipeline) BuildPipeline() {
	_jsii_.InvokeVoid(
		c,
		"buildPipeline",
		nil, // no parameters
	)
}

// Implemented by subclasses to do the actual pipeline construction.
// Experimental.
func (c *jsiiProxy_CodePipeline) DoBuildPipeline() {
	_jsii_.InvokeVoid(
		c,
		"doBuildPipeline",
		nil, // no parameters
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
func (c *jsiiProxy_CodePipeline) OnPrepare() {
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
func (c *jsiiProxy_CodePipeline) OnSynthesize(session constructs.ISynthesisSession) {
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
func (c *jsiiProxy_CodePipeline) OnValidate() *[]*string {
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
func (c *jsiiProxy_CodePipeline) Prepare() {
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
func (c *jsiiProxy_CodePipeline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CodePipeline) ToString() *string {
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
func (c *jsiiProxy_CodePipeline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The result of adding actions to the pipeline.
// Experimental.
type CodePipelineActionFactoryResult struct {
	// How many RunOrders were consumed.
	// Experimental.
	RunOrdersConsumed *float64 `json:"runOrdersConsumed"`
	// If a CodeBuild project got created, the project.
	// Experimental.
	Project awscodebuild.IProject `json:"project"`
}

// A FileSet created from a CodePipeline artifact.
//
// You only need to use this if you want to add CDK Pipeline stages
// add the end of an existing CodePipeline, which should be very rare.
// Experimental.
type CodePipelineFileSet interface {
	FileSet
	Id() *string
	PrimaryOutput() FileSet
	Producer() Step
	ProducedBy(producer Step)
	ToString() *string
}

// The jsii proxy struct for CodePipelineFileSet
type jsiiProxy_CodePipelineFileSet struct {
	jsiiProxy_FileSet
}

func (j *jsiiProxy_CodePipelineFileSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineFileSet) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineFileSet) Producer() Step {
	var returns Step
	_jsii_.Get(
		j,
		"producer",
		&returns,
	)
	return returns
}


// Turn a CodePipeline Artifact into a FileSet.
// Experimental.
func CodePipelineFileSet_FromArtifact(artifact awscodepipeline.Artifact) CodePipelineFileSet {
	_init_.Initialize()

	var returns CodePipelineFileSet

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineFileSet",
		"fromArtifact",
		[]interface{}{artifact},
		&returns,
	)

	return returns
}

// Mark the given Step as the producer for this FileSet.
//
// This method can only be called once.
// Experimental.
func (c *jsiiProxy_CodePipelineFileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		c,
		"producedBy",
		[]interface{}{producer},
	)
}

// Return a string representation of this FileSet.
// Experimental.
func (c *jsiiProxy_CodePipelineFileSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a `CodePipeline`.
// Experimental.
type CodePipelineProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	// Experimental.
	Synth IFileSetProducer `json:"synth"`
	// Additional customizations to apply to the asset publishing CodeBuild projects.
	// Experimental.
	AssetPublishingCodeBuildDefaults *CodeBuildOptions `json:"assetPublishingCodeBuildDefaults"`
	// CDK CLI version to use in self-mutation and asset publishing steps.
	//
	// If you want to lock the CDK CLI version used in the pipeline, by steps
	// that are automatically generated for you, specify the version here.
	//
	// You should not typically need to specify this value.
	// Experimental.
	CliVersion *string `json:"cliVersion"`
	// Customize the CodeBuild projects created for this pipeline.
	// Experimental.
	CodeBuildDefaults *CodeBuildOptions `json:"codeBuildDefaults"`
	// An existing Pipeline to be reused and built upon.
	//
	// [disable-awslint:ref-via-interface]
	// Experimental.
	CodePipeline awscodepipeline.Pipeline `json:"codePipeline"`
	// Create KMS keys for the artifact buckets, allowing cross-account deployments.
	//
	// The artifact buckets have to be encrypted to support deploying CDK apps to
	// another account, so if you want to do that or want to have your artifact
	// buckets encrypted, be sure to set this value to `true`.
	//
	// Be aware there is a cost associated with maintaining the KMS keys.
	// Experimental.
	CrossAccountKeys *bool `json:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
	// Experimental.
	DockerCredentials *[]DockerCredential `json:"dockerCredentials"`
	// Enable Docker for the self-mutate step.
	//
	// Set this to true if the pipeline itself uses Docker container assets
	// (for example, if you use `LinuxBuildImage.fromAsset()` as the build
	// image of a CodeBuild step in the pipeline).
	//
	// You do not need to set it if you build Docker image assets in the
	// application Stages and Stacks that are *deployed* by this pipeline.
	//
	// Configures privileged mode for the self-mutation CodeBuild action.
	//
	// If you are about to turn this on in an already-deployed Pipeline,
	// set the value to `true` first, commit and allow the pipeline to
	// self-update, and only then use the Docker asset in the pipeline.
	// Experimental.
	DockerEnabledForSelfMutation *bool `json:"dockerEnabledForSelfMutation"`
	// Enable Docker for the 'synth' step.
	//
	// Set this to true if you are using file assets that require
	// "bundling" anywhere in your application (meaning an asset
	// compilation step will be run with the tools provided by
	// a Docker image), both for the Pipeline stack as well as the
	// application stacks.
	//
	// A common way to use bundling assets in your application is by
	// using the `@aws-cdk/aws-lambda-nodejs` library.
	//
	// Configures privileged mode for the synth CodeBuild action.
	//
	// If you are about to turn this on in an already-deployed Pipeline,
	// set the value to `true` first, commit and allow the pipeline to
	// self-update, and only then use the bundled asset.
	// Experimental.
	DockerEnabledForSynth *bool `json:"dockerEnabledForSynth"`
	// The name of the CodePipeline pipeline.
	// Experimental.
	PipelineName *string `json:"pipelineName"`
	// Publish assets in multiple CodeBuild projects.
	//
	// If set to false, use one Project per type to publish all assets.
	//
	// Publishing in parallel improves concurrency and may reduce publishing
	// latency, but may also increase overall provisioning time of the CodeBuild
	// projects.
	//
	// Experiment and see what value works best for you.
	// Experimental.
	PublishAssetsInParallel *bool `json:"publishAssetsInParallel"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	// Experimental.
	SelfMutation *bool `json:"selfMutation"`
	// Additional customizations to apply to the self mutation CodeBuild projects.
	// Experimental.
	SelfMutationCodeBuildDefaults *CodeBuildOptions `json:"selfMutationCodeBuildDefaults"`
}

// CodePipeline source steps.
//
// This class contains a number of factory methods for the different types
// of sources that CodePipeline supports.
// Experimental.
type CodePipelineSource interface {
	Step
	ICodePipelineActionFactory
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	ConfigurePrimaryOutput(fs FileSet)
	GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64) awscodepipelineactions.Action
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
	ToString() *string
}

// The jsii proxy struct for CodePipelineSource
type jsiiProxy_CodePipelineSource struct {
	jsiiProxy_Step
	jsiiProxy_ICodePipelineActionFactory
}

func (j *jsiiProxy_CodePipelineSource) Dependencies() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) DependencyFileSets() *[]FileSet {
	var returns *[]FileSet
	_jsii_.Get(
		j,
		"dependencyFileSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) IsSource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipelineSource) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodePipelineSource_Override(c CodePipelineSource, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CodePipelineSource",
		[]interface{}{id},
		c,
	)
}

// Returns a CodeCommit source.
// Experimental.
func CodePipelineSource_CodeCommit(repository awscodecommit.IRepository, branch *string, props *CodeCommitSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"codeCommit",
		[]interface{}{repository, branch, props},
		&returns,
	)

	return returns
}

// Returns a CodeStar connection source.
//
// A CodeStar connection allows AWS CodePipeline to
// access external resources, such as repositories in GitHub, GitHub Enterprise or
// BitBucket.
//
// To use this method, you first need to create a CodeStar connection
// using the AWS console. In the process, you may have to sign in to the external provider
// -- GitHub, for example -- to authorize AWS to read and modify your repository.
// Once you have done this, copy the connection ARN and use it to create the source.
//
// Example:
//
// ```ts
// CodePipelineSource.connection('owner/repo', 'main', {
//    connectionArn: 'arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41', // Created using the AWS console
// });
// ```
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/welcome-connections.html
//
// Experimental.
func CodePipelineSource_Connection(repoString *string, branch *string, props *ConnectionSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"connection",
		[]interface{}{repoString, branch, props},
		&returns,
	)

	return returns
}

// Returns a GitHub source, using OAuth tokens to authenticate with GitHub and a separate webhook to detect changes.
//
// This is no longer
// the recommended method. Please consider using `connection()`
// instead.
//
// Pass in the owner and repository in a single string, like this:
//
// ```ts
// CodePipelineSource.gitHub('owner/repo', 'main');
// ```
//
// Authentication will be done by a secret called `github-token` in AWS
// Secrets Manager (unless specified otherwise).
//
// The token should have these permissions:
//
// * **repo** - to read the repository
// * **admin:repo_hook** - if you plan to use webhooks (true by default)
// Experimental.
func CodePipelineSource_GitHub(repoString *string, branch *string, props *GitHubSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"gitHub",
		[]interface{}{repoString, branch, props},
		&returns,
	)

	return returns
}

// Returns an S3 source.
// Experimental.
func CodePipelineSource_S3(bucket awss3.IBucket, objectKey *string, props *S3SourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"s3",
		[]interface{}{bucket, objectKey, props},
		&returns,
	)

	return returns
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
// Experimental.
func (c *jsiiProxy_CodePipelineSource) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Configure the given FileSet as the primary output of this step.
// Experimental.
func (c *jsiiProxy_CodePipelineSource) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Experimental.
func (c *jsiiProxy_CodePipelineSource) GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64) awscodepipelineactions.Action {
	var returns awscodepipelineactions.Action

	_jsii_.Invoke(
		c,
		"getAction",
		[]interface{}{output, actionName, runOrder},
		&returns,
	)

	return returns
}

// Create the desired Action and add it to the pipeline.
// Experimental.
func (c *jsiiProxy_CodePipelineSource) ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult {
	var returns *CodePipelineActionFactoryResult

	_jsii_.Invoke(
		c,
		"produceAction",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

// Return a string representation of this Step.
// Experimental.
func (c *jsiiProxy_CodePipelineSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Pause the pipeline if a deployment would add IAM permissions or Security Group rules.
//
// This step is only supported in CodePipeline pipelines.
// Experimental.
type ConfirmPermissionsBroadening interface {
	Step
	ICodePipelineActionFactory
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	ConfigurePrimaryOutput(fs FileSet)
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
	ToString() *string
}

// The jsii proxy struct for ConfirmPermissionsBroadening
type jsiiProxy_ConfirmPermissionsBroadening struct {
	jsiiProxy_Step
	jsiiProxy_ICodePipelineActionFactory
}

func (j *jsiiProxy_ConfirmPermissionsBroadening) Dependencies() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfirmPermissionsBroadening) DependencyFileSets() *[]FileSet {
	var returns *[]FileSet
	_jsii_.Get(
		j,
		"dependencyFileSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfirmPermissionsBroadening) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfirmPermissionsBroadening) IsSource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfirmPermissionsBroadening) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}


// Experimental.
func NewConfirmPermissionsBroadening(id *string, props *PermissionsBroadeningCheckProps) ConfirmPermissionsBroadening {
	_init_.Initialize()

	j := jsiiProxy_ConfirmPermissionsBroadening{}

	_jsii_.Create(
		"monocdk.pipelines.ConfirmPermissionsBroadening",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewConfirmPermissionsBroadening_Override(c ConfirmPermissionsBroadening, id *string, props *PermissionsBroadeningCheckProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.ConfirmPermissionsBroadening",
		[]interface{}{id, props},
		c,
	)
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
// Experimental.
func (c *jsiiProxy_ConfirmPermissionsBroadening) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Configure the given FileSet as the primary output of this step.
// Experimental.
func (c *jsiiProxy_ConfirmPermissionsBroadening) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Create the desired Action and add it to the pipeline.
// Experimental.
func (c *jsiiProxy_ConfirmPermissionsBroadening) ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult {
	var returns *CodePipelineActionFactoryResult

	_jsii_.Invoke(
		c,
		"produceAction",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

// Return a string representation of this Step.
// Experimental.
func (c *jsiiProxy_ConfirmPermissionsBroadening) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Configuration options for CodeStar source.
// Experimental.
type ConnectionSourceOptions struct {
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// TODO: EXAMPLE
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	// Experimental.
	ConnectionArn *string `json:"connectionArn"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	// Experimental.
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	// Experimental.
	TriggerOnPush *bool `json:"triggerOnPush"`
}

// Action to deploy a CDK Stack.
//
// Adds two CodePipeline Actions to the pipeline: one to create a ChangeSet
// and one to execute it.
//
// You do not need to instantiate this action yourself -- it will automatically
// be added by the pipeline when you add stack artifacts or entire stages.
// Experimental.
type DeployCdkStackAction interface {
	awscodepipeline.IAction
	ActionProperties() *awscodepipeline.ActionProperties
	DependencyStackArtifactIds() *[]*string
	ExecuteRunOrder() *float64
	PrepareRunOrder() *float64
	StackArtifactId() *string
	StackName() *string
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for DeployCdkStackAction
type jsiiProxy_DeployCdkStackAction struct {
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_DeployCdkStackAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) DependencyStackArtifactIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependencyStackArtifactIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) ExecuteRunOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"executeRunOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) PrepareRunOrder() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"prepareRunOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) StackArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeployCdkStackAction) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}


// Experimental.
func NewDeployCdkStackAction(props *DeployCdkStackActionProps) DeployCdkStackAction {
	_init_.Initialize()

	j := jsiiProxy_DeployCdkStackAction{}

	_jsii_.Create(
		"monocdk.pipelines.DeployCdkStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewDeployCdkStackAction_Override(d DeployCdkStackAction, props *DeployCdkStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.DeployCdkStackAction",
		[]interface{}{props},
		d,
	)
}

// Construct a DeployCdkStackAction from a Stack artifact.
// Experimental.
func DeployCdkStackAction_FromStackArtifact(scope constructs.Construct, artifact cxapi.CloudFormationStackArtifact, options *CdkStackActionFromArtifactOptions) DeployCdkStackAction {
	_init_.Initialize()

	var returns DeployCdkStackAction

	_jsii_.StaticInvoke(
		"monocdk.pipelines.DeployCdkStackAction",
		"fromStackArtifact",
		[]interface{}{scope, artifact, options},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (d *jsiiProxy_DeployCdkStackAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (d *jsiiProxy_DeployCdkStackAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		d,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

// Customization options for a DeployCdkStackAction.
// Experimental.
type DeployCdkStackActionOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Experimental.
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput"`
	// Base name of the action.
	// Experimental.
	BaseActionName *string `json:"baseActionName"`
	// Name of the change set to create and deploy.
	// Experimental.
	ChangeSetName *string `json:"changeSetName"`
	// Run order for the Execute action.
	// Experimental.
	ExecuteRunOrder *float64 `json:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// Filename in output to write Stack outputs to.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
	// Run order for the Prepare action.
	// Experimental.
	PrepareRunOrder *float64 `json:"prepareRunOrder"`
}

// Properties for a DeployCdkStackAction.
// Experimental.
type DeployCdkStackActionProps struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Experimental.
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput"`
	// Base name of the action.
	// Experimental.
	BaseActionName *string `json:"baseActionName"`
	// Name of the change set to create and deploy.
	// Experimental.
	ChangeSetName *string `json:"changeSetName"`
	// Run order for the Execute action.
	// Experimental.
	ExecuteRunOrder *float64 `json:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// Filename in output to write Stack outputs to.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
	// Run order for the Prepare action.
	// Experimental.
	PrepareRunOrder *float64 `json:"prepareRunOrder"`
	// Role for the action to assume.
	//
	// This controls the account to deploy into
	// Experimental.
	ActionRole awsiam.IRole `json:"actionRole"`
	// The name of the stack that should be created/updated.
	// Experimental.
	StackName *string `json:"stackName"`
	// Relative path of template in the input artifact.
	// Experimental.
	TemplatePath *string `json:"templatePath"`
	// Role to execute CloudFormation under.
	// Experimental.
	CloudFormationExecutionRole awsiam.IRole `json:"cloudFormationExecutionRole"`
	// Artifact ID for the stacks this stack depends on.
	//
	// Used for pipeline order checking.
	// Experimental.
	DependencyStackArtifactIds *[]*string `json:"dependencyStackArtifactIds"`
	// Region to deploy into.
	// Experimental.
	Region *string `json:"region"`
	// Artifact ID for the stack deployed here.
	//
	// Used for pipeline order checking.
	// Experimental.
	StackArtifactId *string `json:"stackArtifactId"`
	// Template configuration path relative to the input artifact.
	// Experimental.
	TemplateConfigurationPath *string `json:"templateConfigurationPath"`
}

// Represents credentials used to access a Docker registry.
// Experimental.
type DockerCredential interface {
	Usages() *[]DockerCredentialUsage
	GrantRead(grantee awsiam.IGrantable, usage DockerCredentialUsage)
}

// The jsii proxy struct for DockerCredential
type jsiiProxy_DockerCredential struct {
	_ byte // padding
}

func (j *jsiiProxy_DockerCredential) Usages() *[]DockerCredentialUsage {
	var returns *[]DockerCredentialUsage
	_jsii_.Get(
		j,
		"usages",
		&returns,
	)
	return returns
}


// Experimental.
func NewDockerCredential_Override(d DockerCredential, usages *[]DockerCredentialUsage) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.DockerCredential",
		[]interface{}{usages},
		d,
	)
}

// Creates a DockerCredential for a registry, based on its domain name (e.g., 'www.example.com').
// Experimental.
func DockerCredential_CustomRegistry(registryDomain *string, secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	var returns DockerCredential

	_jsii_.StaticInvoke(
		"monocdk.pipelines.DockerCredential",
		"customRegistry",
		[]interface{}{registryDomain, secret, opts},
		&returns,
	)

	return returns
}

// Creates a DockerCredential for DockerHub.
//
// Convenience method for `fromCustomRegistry('index.docker.io', opts)`.
// Experimental.
func DockerCredential_DockerHub(secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	var returns DockerCredential

	_jsii_.StaticInvoke(
		"monocdk.pipelines.DockerCredential",
		"dockerHub",
		[]interface{}{secret, opts},
		&returns,
	)

	return returns
}

// Creates a DockerCredential for one or more ECR repositories.
//
// NOTE - All ECR repositories in the same account and region share a domain name
// (e.g., 0123456789012.dkr.ecr.eu-west-1.amazonaws.com), and can only have one associated
// set of credentials (and DockerCredential). Attempting to associate one set of credentials
// with one ECR repo and another with another ECR repo in the same account and region will
// result in failures when using these credentials in the pipeline.
// Experimental.
func DockerCredential_Ecr(repositories *[]awsecr.IRepository, opts *EcrDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	var returns DockerCredential

	_jsii_.StaticInvoke(
		"monocdk.pipelines.DockerCredential",
		"ecr",
		[]interface{}{repositories, opts},
		&returns,
	)

	return returns
}

// Grant read-only access to the registry credentials.
//
// This grants read access to any secrets, and pull access to any repositories.
// Experimental.
func (d *jsiiProxy_DockerCredential) GrantRead(grantee awsiam.IGrantable, usage DockerCredentialUsage) {
	_jsii_.InvokeVoid(
		d,
		"grantRead",
		[]interface{}{grantee, usage},
	)
}

// Defines which stages of a pipeline require the specified credentials.
// Experimental.
type DockerCredentialUsage string

const (
	DockerCredentialUsage_SYNTH DockerCredentialUsage = "SYNTH"
	DockerCredentialUsage_SELF_UPDATE DockerCredentialUsage = "SELF_UPDATE"
	DockerCredentialUsage_ASSET_PUBLISHING DockerCredentialUsage = "ASSET_PUBLISHING"
)

// Options for defining access for a Docker Credential composed of ECR repos.
// Experimental.
type EcrDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	// Experimental.
	AssumeRole awsiam.IRole `json:"assumeRole"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	// Experimental.
	Usages *[]DockerCredentialUsage `json:"usages"`
}

// Options for defining credentials for a Docker Credential.
// Experimental.
type ExternalDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	// Experimental.
	AssumeRole awsiam.IRole `json:"assumeRole"`
	// The name of the JSON field of the secret which contains the secret/password.
	// Experimental.
	SecretPasswordField *string `json:"secretPasswordField"`
	// The name of the JSON field of the secret which contains the user/login name.
	// Experimental.
	SecretUsernameField *string `json:"secretUsernameField"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	// Experimental.
	Usages *[]DockerCredentialUsage `json:"usages"`
}

// A set of files traveling through the deployment pipeline.
//
// Individual steps in the pipeline produce or consume
// `FileSet`s.
// Experimental.
type FileSet interface {
	IFileSetProducer
	Id() *string
	PrimaryOutput() FileSet
	Producer() Step
	ProducedBy(producer Step)
	ToString() *string
}

// The jsii proxy struct for FileSet
type jsiiProxy_FileSet struct {
	jsiiProxy_IFileSetProducer
}

func (j *jsiiProxy_FileSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSet) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSet) Producer() Step {
	var returns Step
	_jsii_.Get(
		j,
		"producer",
		&returns,
	)
	return returns
}


// Experimental.
func NewFileSet(id *string, producer Step) FileSet {
	_init_.Initialize()

	j := jsiiProxy_FileSet{}

	_jsii_.Create(
		"monocdk.pipelines.FileSet",
		[]interface{}{id, producer},
		&j,
	)

	return &j
}

// Experimental.
func NewFileSet_Override(f FileSet, id *string, producer Step) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.FileSet",
		[]interface{}{id, producer},
		f,
	)
}

// Mark the given Step as the producer for this FileSet.
//
// This method can only be called once.
// Experimental.
func (f *jsiiProxy_FileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		f,
		"producedBy",
		[]interface{}{producer},
	)
}

// Return a string representation of this FileSet.
// Experimental.
func (f *jsiiProxy_FileSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Location of a FileSet consumed or produced by a ShellStep.
// Experimental.
type FileSetLocation struct {
	// The (relative) directory where the FileSet is found.
	// Experimental.
	Directory *string `json:"directory"`
	// The FileSet object.
	// Experimental.
	FileSet FileSet `json:"fileSet"`
}

// Options for CdkDeployAction.fromStackArtifact.
// Experimental.
type FromStackArtifactOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Experimental.
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput"`
	// Run order for the Execute action.
	// Experimental.
	ExecuteRunOrder *float64 `json:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Experimental.
	Output awscodepipeline.Artifact `json:"output"`
	// Filename in output to write Stack outputs to.
	// Experimental.
	OutputFileName *string `json:"outputFileName"`
	// Run order for the 2 actions that will be created.
	// Experimental.
	PrepareRunOrder *float64 `json:"prepareRunOrder"`
}

// Options for GitHub sources.
// Experimental.
type GitHubSourceOptions struct {
	// A GitHub OAuth token to use for authentication.
	//
	// It is recommended to use a Secrets Manager `Secret` to obtain the token:
	//
	// ```ts
	// const oauth = cdk.SecretValue.secretsManager('my-github-token');
	// new GitHubSource(this, 'GitHubSource', { authentication: oauth, ... });
	// ```
	//
	// The GitHub Personal Access Token should have these scopes:
	//
	// * **repo** - to read the repository
	// * **admin:repo_hook** - if you plan to use webhooks (true by default)
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/GitHub-create-personal-token-CLI.html
	//
	// Experimental.
	Authentication awscdk.SecretValue `json:"authentication"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action.
	// With "POLL", CodePipeline periodically checks the source for changes.
	// With "None", the action is not triggered through changes in the source.
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	// Experimental.
	Trigger awscodepipelineactions.GitHubTrigger `json:"trigger"`
}

// Factory for explicit CodePipeline Actions.
//
// If you have specific types of Actions you want to add to a
// CodePipeline, write a subclass of `Step` that implements this
// interface, and add the action or actions you want in the `produce` method.
//
// There needs to be a level of indirection here, because some aspects of the
// Action creation need to be controlled by the workflow engine (name and
// runOrder). All the rest of the properties are controlled by the factory.
// Experimental.
type ICodePipelineActionFactory interface {
	// Create the desired Action and add it to the pipeline.
	// Experimental.
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
}

// The jsii proxy for ICodePipelineActionFactory
type jsiiProxy_ICodePipelineActionFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_ICodePipelineActionFactory) ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult {
	var returns *CodePipelineActionFactoryResult

	_jsii_.Invoke(
		i,
		"produceAction",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

// Any class that produces, or is itself, a `FileSet`.
//
// Steps implicitly produce a primary FileSet as an output.
// Experimental.
type IFileSetProducer interface {
	// The `FileSet` produced by this file set producer.
	// Experimental.
	PrimaryOutput() FileSet
}

// The jsii proxy for IFileSetProducer
type jsiiProxy_IFileSetProducer struct {
	_ byte // padding
}

func (j *jsiiProxy_IFileSetProducer) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}

// Features that the Stage needs from its environment.
// Experimental.
type IStageHost interface {
	// Make sure all the assets from the given manifest are published.
	// Experimental.
	PublishAsset(command *AssetPublishingCommand)
	// Return the Artifact the given stack has to emit its outputs into, if any.
	// Experimental.
	StackOutputArtifact(stackArtifactId *string) awscodepipeline.Artifact
}

// The jsii proxy for IStageHost
type jsiiProxy_IStageHost struct {
	_ byte // padding
}

func (i *jsiiProxy_IStageHost) PublishAsset(command *AssetPublishingCommand) {
	_jsii_.InvokeVoid(
		i,
		"publishAsset",
		[]interface{}{command},
	)
}

func (i *jsiiProxy_IStageHost) StackOutputArtifact(stackArtifactId *string) awscodepipeline.Artifact {
	var returns awscodepipeline.Artifact

	_jsii_.Invoke(
		i,
		"stackOutputArtifact",
		[]interface{}{stackArtifactId},
		&returns,
	)

	return returns
}

// A manual approval step.
//
// If this step is added to a Pipeline, the Pipeline will
// be paused waiting for a human to resume it
//
// Only engines that support pausing the deployment will
// support this step type.
// Experimental.
type ManualApprovalStep interface {
	Step
	Comment() *string
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	ConfigurePrimaryOutput(fs FileSet)
	ToString() *string
}

// The jsii proxy struct for ManualApprovalStep
type jsiiProxy_ManualApprovalStep struct {
	jsiiProxy_Step
}

func (j *jsiiProxy_ManualApprovalStep) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalStep) Dependencies() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalStep) DependencyFileSets() *[]FileSet {
	var returns *[]FileSet
	_jsii_.Get(
		j,
		"dependencyFileSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalStep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalStep) IsSource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalStep) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}


// Experimental.
func NewManualApprovalStep(id *string, props *ManualApprovalStepProps) ManualApprovalStep {
	_init_.Initialize()

	j := jsiiProxy_ManualApprovalStep{}

	_jsii_.Create(
		"monocdk.pipelines.ManualApprovalStep",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewManualApprovalStep_Override(m ManualApprovalStep, id *string, props *ManualApprovalStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.ManualApprovalStep",
		[]interface{}{id, props},
		m,
	)
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
// Experimental.
func (m *jsiiProxy_ManualApprovalStep) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		m,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Configure the given FileSet as the primary output of this step.
// Experimental.
func (m *jsiiProxy_ManualApprovalStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		m,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Return a string representation of this Step.
// Experimental.
func (m *jsiiProxy_ManualApprovalStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a `ManualApprovalStep`.
// Experimental.
type ManualApprovalStepProps struct {
	// The comment to display with this manual approval.
	// Experimental.
	Comment *string `json:"comment"`
}

// Properties for a `PermissionsBroadeningCheck`.
// Experimental.
type PermissionsBroadeningCheckProps struct {
	// The CDK Stage object to check the stacks of.
	//
	// This should be the same Stage object you are passing to `addStage()`.
	// Experimental.
	Stage awscdk.Stage `json:"stage"`
	// Topic to send notifications when a human needs to give manual confirmation.
	// Experimental.
	NotificationTopic awssns.ITopic `json:"notificationTopic"`
}

// A generic CDK Pipelines pipeline.
//
// Different deployment systems will provide subclasses of `Pipeline` that generate
// the deployment infrastructure necessary to deploy CDK apps, specific to that system.
//
// This library comes with the `CodePipeline` class, which uses AWS CodePipeline
// to deploy CDK apps.
//
// The actual pipeline infrastructure is constructed (by invoking the engine)
// when `buildPipeline()` is called, or when `app.synth()` is called (whichever
// happens first).
// Experimental.
type PipelineBase interface {
	awscdk.Construct
	CloudAssemblyFileSet() FileSet
	Node() awscdk.ConstructNode
	Synth() IFileSetProducer
	Waves() *[]Wave
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
	AddWave(id *string, options *WaveOptions) Wave
	BuildPipeline()
	DoBuildPipeline()
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for PipelineBase
type jsiiProxy_PipelineBase struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_PipelineBase) CloudAssemblyFileSet() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"cloudAssemblyFileSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineBase) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineBase) Synth() IFileSetProducer {
	var returns IFileSetProducer
	_jsii_.Get(
		j,
		"synth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineBase) Waves() *[]Wave {
	var returns *[]Wave
	_jsii_.Get(
		j,
		"waves",
		&returns,
	)
	return returns
}


// Experimental.
func NewPipelineBase_Override(p PipelineBase, scope constructs.Construct, id *string, props *PipelineBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.PipelineBase",
		[]interface{}{scope, id, props},
		p,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func PipelineBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.PipelineBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Deploy a single Stage by itself.
//
// Add a Stage to the pipeline, to be deployed in sequence with other
// Stages added to the pipeline. All Stacks in the stage will be deployed
// in an order automatically determined by their relative dependencies.
// Experimental.
func (p *jsiiProxy_PipelineBase) AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment {
	var returns StageDeployment

	_jsii_.Invoke(
		p,
		"addStage",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

// Add a Wave to the pipeline, for deploying multiple Stages in parallel.
//
// Use the return object of this method to deploy multiple stages in parallel.
//
// Example:
//
// ```ts
// const wave = pipeline.addWave('MyWave');
// wave.addStage(new MyStage('Stage1', ...));
// wave.addStage(new MyStage('Stage2', ...));
// ```
// Experimental.
func (p *jsiiProxy_PipelineBase) AddWave(id *string, options *WaveOptions) Wave {
	var returns Wave

	_jsii_.Invoke(
		p,
		"addWave",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

// Send the current pipeline definition to the engine, and construct the pipeline.
//
// It is not possible to modify the pipeline after calling this method.
// Experimental.
func (p *jsiiProxy_PipelineBase) BuildPipeline() {
	_jsii_.InvokeVoid(
		p,
		"buildPipeline",
		nil, // no parameters
	)
}

// Implemented by subclasses to do the actual pipeline construction.
// Experimental.
func (p *jsiiProxy_PipelineBase) DoBuildPipeline() {
	_jsii_.InvokeVoid(
		p,
		"doBuildPipeline",
		nil, // no parameters
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
func (p *jsiiProxy_PipelineBase) OnPrepare() {
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
func (p *jsiiProxy_PipelineBase) OnSynthesize(session constructs.ISynthesisSession) {
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
func (p *jsiiProxy_PipelineBase) OnValidate() *[]*string {
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
func (p *jsiiProxy_PipelineBase) Prepare() {
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
func (p *jsiiProxy_PipelineBase) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_PipelineBase) ToString() *string {
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
func (p *jsiiProxy_PipelineBase) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a `Pipeline`.
// Experimental.
type PipelineBaseProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	// Experimental.
	Synth IFileSetProducer `json:"synth"`
}

// Options for the `CodePipelineActionFactory.produce()` method.
// Experimental.
type ProduceActionOptions struct {
	// Name the action should get.
	// Experimental.
	ActionName *string `json:"actionName"`
	// Helper object to translate FileSets to CodePipeline Artifacts.
	// Experimental.
	Artifacts ArtifactMap `json:"artifacts"`
	// The pipeline the action is being generated for.
	// Experimental.
	Pipeline CodePipeline `json:"pipeline"`
	// RunOrder the action should get.
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// Scope in which to create constructs.
	// Experimental.
	Scope constructs.Construct `json:"scope"`
	// Whether or not this action is inserted before self mutation.
	//
	// If it is, the action should take care to reflect some part of
	// its own definition in the pipeline action definition, to
	// trigger a restart after self-mutation (if necessary).
	// Experimental.
	BeforeSelfMutation *bool `json:"beforeSelfMutation"`
	// If this action factory creates a CodeBuild step, default options to inherit.
	// Experimental.
	CodeBuildDefaults *CodeBuildOptions `json:"codeBuildDefaults"`
	// An input artifact that CodeBuild projects that don't actually need an input artifact can use.
	//
	// CodeBuild Projects MUST have an input artifact in order to be added to the Pipeline. If
	// the Project doesn't actually care about its input (it can be anything), it can use the
	// Artifact passed here.
	// Experimental.
	FallbackArtifact awscodepipeline.Artifact `json:"fallbackArtifact"`
}

// Action to publish an asset in the pipeline.
//
// Creates a CodeBuild project which will use the CDK CLI
// to prepare and publish the asset.
//
// You do not need to instantiate this action -- it will automatically
// be added by the pipeline when you add stacks that use assets.
// Experimental.
type PublishAssetsAction interface {
	awscdk.Construct
	awscodepipeline.IAction
	ActionProperties() *awscodepipeline.ActionProperties
	Node() awscdk.ConstructNode
	AddPublishCommand(relativeManifestPath *string, assetSelector *string)
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnPrepare()
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for PublishAssetsAction
type jsiiProxy_PublishAssetsAction struct {
	internal.Type__awscdkConstruct
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_PublishAssetsAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PublishAssetsAction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewPublishAssetsAction(scope constructs.Construct, id *string, props *PublishAssetsActionProps) PublishAssetsAction {
	_init_.Initialize()

	j := jsiiProxy_PublishAssetsAction{}

	_jsii_.Create(
		"monocdk.pipelines.PublishAssetsAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPublishAssetsAction_Override(p PublishAssetsAction, scope constructs.Construct, id *string, props *PublishAssetsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.PublishAssetsAction",
		[]interface{}{scope, id, props},
		p,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func PublishAssetsAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.PublishAssetsAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a single publishing command.
//
// Manifest path should be relative to the root Cloud Assembly.
// Experimental.
func (p *jsiiProxy_PublishAssetsAction) AddPublishCommand(relativeManifestPath *string, assetSelector *string) {
	_jsii_.InvokeVoid(
		p,
		"addPublishCommand",
		[]interface{}{relativeManifestPath, assetSelector},
	)
}

// Exists to implement IAction.
// Experimental.
func (p *jsiiProxy_PublishAssetsAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{scope, stage, options},
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
func (p *jsiiProxy_PublishAssetsAction) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

// Exists to implement IAction.
// Experimental.
func (p *jsiiProxy_PublishAssetsAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (p *jsiiProxy_PublishAssetsAction) OnSynthesize(session constructs.ISynthesisSession) {
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
func (p *jsiiProxy_PublishAssetsAction) OnValidate() *[]*string {
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
func (p *jsiiProxy_PublishAssetsAction) Prepare() {
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
func (p *jsiiProxy_PublishAssetsAction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (p *jsiiProxy_PublishAssetsAction) ToString() *string {
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
func (p *jsiiProxy_PublishAssetsAction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Props for a PublishAssetsAction.
// Experimental.
type PublishAssetsActionProps struct {
	// Name of publishing action.
	// Experimental.
	ActionName *string `json:"actionName"`
	// AssetType we're publishing.
	// Experimental.
	AssetType AssetType `json:"assetType"`
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Experimental.
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput"`
	// Custom BuildSpec that is merged with generated one.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// Version of CDK CLI to 'npm install'.
	// Experimental.
	CdkCliVersion *string `json:"cdkCliVersion"`
	// Use a file buildspec written to the cloud assembly instead of an inline buildspec.
	//
	// This prevents size limitation errors as inline specs have a max length of 25600 characters
	// Experimental.
	CreateBuildspecFile *bool `json:"createBuildspecFile"`
	// Any Dependable construct that the CodeBuild project needs to take a dependency on.
	// Experimental.
	Dependable awscdk.IDependable `json:"dependable"`
	// Additional commands to run before installing cdk-assert Use this to setup proxies or npm mirrors.
	// Experimental.
	PreInstallCommands *[]*string `json:"preInstallCommands"`
	// Name of the CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
	// Role to use for CodePipeline and CodeBuild to build and publish the assets.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where to execute the PublishAssetsAction.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Options for S3 sources.
// Experimental.
type S3SourceOptions struct {
	// The action name used for this source in the CodePipeline.
	// Experimental.
	ActionName *string `json:"actionName"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	// Experimental.
	Trigger awscodepipelineactions.S3Trigger `json:"trigger"`
}

// Validate a revision using shell commands.
// Experimental.
type ShellScriptAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	ActionProperties() *awscodepipeline.ActionProperties
	GrantPrincipal() awsiam.IPrincipal
	Project() awscodebuild.IProject
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for ShellScriptAction
type jsiiProxy_ShellScriptAction struct {
	internal.Type__awscodepipelineIAction
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_ShellScriptAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellScriptAction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellScriptAction) Project() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewShellScriptAction(props *ShellScriptActionProps) ShellScriptAction {
	_init_.Initialize()

	j := jsiiProxy_ShellScriptAction{}

	_jsii_.Create(
		"monocdk.pipelines.ShellScriptAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewShellScriptAction_Override(s ShellScriptAction, props *ShellScriptActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.ShellScriptAction",
		[]interface{}{props},
		s,
	)
}

// Exists to implement IAction.
// Experimental.
func (s *jsiiProxy_ShellScriptAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (s *jsiiProxy_ShellScriptAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

// Properties for ShellScriptAction.
// Experimental.
type ShellScriptActionProps struct {
	// Name of the validation action in the pipeline.
	// Experimental.
	ActionName *string `json:"actionName"`
	// Commands to run.
	// Experimental.
	Commands *[]*string `json:"commands"`
	// Additional artifacts to use as input for the CodeBuild project.
	//
	// You can use these files to load more complex test sets into the
	// shellscript build environment.
	//
	// The files artifact given here will be unpacked into the current
	// working directory, the other ones will be unpacked into directories
	// which are available through the environment variables
	// $CODEBUILD_SRC_DIR_<artifactName>.
	//
	// The CodeBuild job must have at least one input artifact, so you
	// must provide either at least one additional artifact here or one
	// stack output using `useOutput`.
	// Experimental.
	AdditionalArtifacts *[]awscodepipeline.Artifact `json:"additionalArtifacts"`
	// Bash options to set at the start of the script.
	// Experimental.
	BashOptions *string `json:"bashOptions"`
	// The CodeBuild environment where scripts are executed.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
	// Experimental.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables"`
	// Additional policy statements to add to the execution role.
	// Experimental.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements"`
	// RunOrder for this action.
	//
	// Use this to sequence the shell script after the deployments.
	//
	// The default value is 100 so you don't have to supply the value if you just
	// want to run this after the application stacks have been deployed, and you
	// don't have more than 100 stacks.
	// Experimental.
	RunOrder *float64 `json:"runOrder"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// Stack outputs to make available as environment variables.
	// Experimental.
	UseOutputs *map[string]StackOutput `json:"useOutputs"`
	// The VPC where to execute the specified script.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Run shell script commands in the pipeline.
// Experimental.
type ShellStep interface {
	Step
	Commands() *[]*string
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Env() *map[string]*string
	EnvFromCfnOutputs() *map[string]StackOutputReference
	Id() *string
	Inputs() *[]*FileSetLocation
	InstallCommands() *[]*string
	IsSource() *bool
	Outputs() *[]*FileSetLocation
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	AddOutputDirectory(directory *string) FileSet
	ConfigurePrimaryOutput(fs FileSet)
	PrimaryOutputDirectory(directory *string) FileSet
	ToString() *string
}

// The jsii proxy struct for ShellStep
type jsiiProxy_ShellStep struct {
	jsiiProxy_Step
}

func (j *jsiiProxy_ShellStep) Commands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) Dependencies() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) DependencyFileSets() *[]FileSet {
	var returns *[]FileSet
	_jsii_.Get(
		j,
		"dependencyFileSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) Env() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) EnvFromCfnOutputs() *map[string]StackOutputReference {
	var returns *map[string]StackOutputReference
	_jsii_.Get(
		j,
		"envFromCfnOutputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) Inputs() *[]*FileSetLocation {
	var returns *[]*FileSetLocation
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) InstallCommands() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"installCommands",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) IsSource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) Outputs() *[]*FileSetLocation {
	var returns *[]*FileSetLocation
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ShellStep) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}


// Experimental.
func NewShellStep(id *string, props *ShellStepProps) ShellStep {
	_init_.Initialize()

	j := jsiiProxy_ShellStep{}

	_jsii_.Create(
		"monocdk.pipelines.ShellStep",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewShellStep_Override(s ShellStep, id *string, props *ShellStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.ShellStep",
		[]interface{}{id, props},
		s,
	)
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
// Experimental.
func (s *jsiiProxy_ShellStep) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Add an additional output FileSet based on a directory.
//
// After running the script, the contents of the given directory
// will be exported as a `FileSet`. Use the `FileSet` as the
// input to another step.
//
// Multiple calls with the exact same directory name string (not normalized)
// will return the same FileSet.
// Experimental.
func (s *jsiiProxy_ShellStep) AddOutputDirectory(directory *string) FileSet {
	var returns FileSet

	_jsii_.Invoke(
		s,
		"addOutputDirectory",
		[]interface{}{directory},
		&returns,
	)

	return returns
}

// Configure the given FileSet as the primary output of this step.
// Experimental.
func (s *jsiiProxy_ShellStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Configure the given output directory as primary output.
//
// If no primary output has been configured yet, this directory
// will become the primary output of this ShellStep, otherwise this
// method will throw if the given directory is different than the
// currently configured primary output directory.
// Experimental.
func (s *jsiiProxy_ShellStep) PrimaryOutputDirectory(directory *string) FileSet {
	var returns FileSet

	_jsii_.Invoke(
		s,
		"primaryOutputDirectory",
		[]interface{}{directory},
		&returns,
	)

	return returns
}

// Return a string representation of this Step.
// Experimental.
func (s *jsiiProxy_ShellStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Construction properties for a `ShellStep`.
// Experimental.
type ShellStepProps struct {
	// Commands to run.
	// Experimental.
	Commands *[]*string `json:"commands"`
	// Additional FileSets to put in other directories.
	//
	// Specifies a mapping from directory name to FileSets. During the
	// script execution, the FileSets will be available in the directories
	// indicated.
	//
	// The directory names may be relative. For example, you can put
	// the main input and an additional input side-by-side with the
	// following configuration:
	//
	// ```ts
	// const script = new ShellStep('MainScript', {
	//    // ...
	//    input: MyEngineSource.gitHub('org/source1'),
	//    additionalInputs: {
	//      '../siblingdir': MyEngineSource.gitHub('org/source2'),
	//    }
	// });
	// ```
	// Experimental.
	AdditionalInputs *map[string]IFileSetProducer `json:"additionalInputs"`
	// Environment variables to set.
	// Experimental.
	Env *map[string]*string `json:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	// Experimental.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `json:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	// Experimental.
	Input IFileSetProducer `json:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Experimental.
	InstallCommands *[]*string `json:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	// Experimental.
	PrimaryOutputDirectory *string `json:"primaryOutputDirectory"`
}

// A standard synth with a generated buildspec.
// Experimental.
type SimpleSynthAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	ActionProperties() *awscodepipeline.ActionProperties
	GrantPrincipal() awsiam.IPrincipal
	Project() awscodebuild.IProject
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for SimpleSynthAction
type jsiiProxy_SimpleSynthAction struct {
	internal.Type__awscodepipelineIAction
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_SimpleSynthAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SimpleSynthAction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SimpleSynthAction) Project() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewSimpleSynthAction(props *SimpleSynthActionProps) SimpleSynthAction {
	_init_.Initialize()

	j := jsiiProxy_SimpleSynthAction{}

	_jsii_.Create(
		"monocdk.pipelines.SimpleSynthAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSimpleSynthAction_Override(s SimpleSynthAction, props *SimpleSynthActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.SimpleSynthAction",
		[]interface{}{props},
		s,
	)
}

// Create a standard NPM synth action.
//
// Uses `npm ci` to install dependencies and `npx cdk synth` to synthesize.
//
// If you need a build step, add `buildCommand: 'npm run build'`.
// Experimental.
func SimpleSynthAction_StandardNpmSynth(options *StandardNpmSynthOptions) SimpleSynthAction {
	_init_.Initialize()

	var returns SimpleSynthAction

	_jsii_.StaticInvoke(
		"monocdk.pipelines.SimpleSynthAction",
		"standardNpmSynth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Create a standard Yarn synth action.
//
// Uses `yarn install --frozen-lockfile` to install dependencies and `npx cdk synth` to synthesize.
//
// If you need a build step, add `buildCommand: 'yarn build'`.
// Experimental.
func SimpleSynthAction_StandardYarnSynth(options *StandardYarnSynthOptions) SimpleSynthAction {
	_init_.Initialize()

	var returns SimpleSynthAction

	_jsii_.StaticInvoke(
		"monocdk.pipelines.SimpleSynthAction",
		"standardYarnSynth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (s *jsiiProxy_SimpleSynthAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (s *jsiiProxy_SimpleSynthAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

// Construction props for SimpleSynthAction.
// Experimental.
type SimpleSynthActionProps struct {
	// The artifact where the CloudAssembly should be emitted.
	// Experimental.
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Experimental.
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact"`
	// Name of the build action.
	// Experimental.
	ActionName *string `json:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Experimental.
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Experimental.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables"`
	// Name of the CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Experimental.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Experimental.
	Subdirectory *string `json:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The synth command.
	// Experimental.
	SynthCommand *string `json:"synthCommand"`
	// The build command.
	//
	// If your programming language requires a compilation step, put the
	// compilation command here.
	// Deprecated: Use `buildCommands` instead
	BuildCommand *string `json:"buildCommand"`
	// The build commands.
	//
	// If your programming language requires a compilation step, put the
	// compilation command here.
	// Experimental.
	BuildCommands *[]*string `json:"buildCommands"`
	// The install command.
	//
	// If not provided by the build image or another dependency
	// management tool, at least install the CDK CLI here using
	// `npm install -g aws-cdk`.
	// Deprecated: Use `installCommands` instead
	InstallCommand *string `json:"installCommand"`
	// Install commands.
	//
	// If not provided by the build image or another dependency
	// management tool, at least install the CDK CLI here using
	// `npm install -g aws-cdk`.
	// Experimental.
	InstallCommands *[]*string `json:"installCommands"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Experimental.
	TestCommands *[]*string `json:"testCommands"`
}

// Configuration options for a SimpleSynth.
// Experimental.
type SimpleSynthOptions struct {
	// The artifact where the CloudAssembly should be emitted.
	// Experimental.
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Experimental.
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact"`
	// Name of the build action.
	// Experimental.
	ActionName *string `json:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Experimental.
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Experimental.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables"`
	// Name of the CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Experimental.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Experimental.
	Subdirectory *string `json:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// An asset used by a Stack.
// Experimental.
type StackAsset struct {
	// Asset identifier.
	// Experimental.
	AssetId *string `json:"assetId"`
	// Absolute asset manifest path.
	//
	// This needs to be made relative at a later point in time, but when this
	// information is parsed we don't know about the root cloud assembly yet.
	// Experimental.
	AssetManifestPath *string `json:"assetManifestPath"`
	// Asset selector to pass to `cdk-assets`.
	// Experimental.
	AssetSelector *string `json:"assetSelector"`
	// Type of asset to publish.
	// Experimental.
	AssetType AssetType `json:"assetType"`
	// Does this asset represent the CloudFormation template for the stack.
	// Experimental.
	IsTemplate *bool `json:"isTemplate"`
	// Role ARN to assume to publish.
	// Experimental.
	AssetPublishingRoleArn *string `json:"assetPublishingRoleArn"`
}

// Deployment of a single Stack.
//
// You don't need to instantiate this class -- it will
// be automatically instantiated as necessary when you
// add a `Stage` to a pipeline.
// Experimental.
type StackDeployment interface {
	AbsoluteTemplatePath() *string
	Account() *string
	Assets() *[]*StackAsset
	AssumeRoleArn() *string
	ConstructPath() *string
	ExecutionRoleArn() *string
	Region() *string
	StackArtifactId() *string
	StackDependencies() *[]StackDeployment
	StackName() *string
	Tags() *map[string]*string
	TemplateAsset() *StackAsset
	TemplateUrl() *string
	AddStackDependency(stackDeployment StackDeployment)
}

// The jsii proxy struct for StackDeployment
type jsiiProxy_StackDeployment struct {
	_ byte // padding
}

func (j *jsiiProxy_StackDeployment) AbsoluteTemplatePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteTemplatePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Assets() *[]*StackAsset {
	var returns *[]*StackAsset
	_jsii_.Get(
		j,
		"assets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) AssumeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) ConstructPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) StackArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackArtifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) StackDependencies() *[]StackDeployment {
	var returns *[]StackDeployment
	_jsii_.Get(
		j,
		"stackDependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) TemplateAsset() *StackAsset {
	var returns *StackAsset
	_jsii_.Get(
		j,
		"templateAsset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) TemplateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateUrl",
		&returns,
	)
	return returns
}


// Build a `StackDeployment` from a Stack Artifact in a Cloud Assembly.
// Experimental.
func StackDeployment_FromArtifact(stackArtifact cxapi.CloudFormationStackArtifact) StackDeployment {
	_init_.Initialize()

	var returns StackDeployment

	_jsii_.StaticInvoke(
		"monocdk.pipelines.StackDeployment",
		"fromArtifact",
		[]interface{}{stackArtifact},
		&returns,
	)

	return returns
}

// Add a dependency on another stack.
// Experimental.
func (s *jsiiProxy_StackDeployment) AddStackDependency(stackDeployment StackDeployment) {
	_jsii_.InvokeVoid(
		s,
		"addStackDependency",
		[]interface{}{stackDeployment},
	)
}

// Properties for a `StackDeployment`.
// Experimental.
type StackDeploymentProps struct {
	// Template path on disk to cloud assembly (cdk.out).
	// Experimental.
	AbsoluteTemplatePath *string `json:"absoluteTemplatePath"`
	// Construct path for this stack.
	// Experimental.
	ConstructPath *string `json:"constructPath"`
	// Artifact ID for this stack.
	// Experimental.
	StackArtifactId *string `json:"stackArtifactId"`
	// Name for this stack.
	// Experimental.
	StackName *string `json:"stackName"`
	// Account where the stack should be deployed.
	// Experimental.
	Account *string `json:"account"`
	// Assets referenced by this stack.
	// Experimental.
	Assets *[]*StackAsset `json:"assets"`
	// Role to assume before deploying this stack.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn"`
	// Execution role to pass to CloudFormation.
	// Experimental.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// Region where the stack should be deployed.
	// Experimental.
	Region *string `json:"region"`
	// Tags to apply to the stack.
	// Experimental.
	Tags *map[string]*string `json:"tags"`
	// The S3 URL which points to the template asset location in the publishing bucket.
	// Experimental.
	TemplateS3Uri *string `json:"templateS3Uri"`
}

// A single output of a Stack.
// Experimental.
type StackOutput interface {
	ArtifactFile() awscodepipeline.ArtifactPath
	OutputName() *string
}

// The jsii proxy struct for StackOutput
type jsiiProxy_StackOutput struct {
	_ byte // padding
}

func (j *jsiiProxy_StackOutput) ArtifactFile() awscodepipeline.ArtifactPath {
	var returns awscodepipeline.ArtifactPath
	_jsii_.Get(
		j,
		"artifactFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackOutput) OutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputName",
		&returns,
	)
	return returns
}


// Build a StackOutput from a known artifact and an output name.
// Experimental.
func NewStackOutput(artifactFile awscodepipeline.ArtifactPath, outputName *string) StackOutput {
	_init_.Initialize()

	j := jsiiProxy_StackOutput{}

	_jsii_.Create(
		"monocdk.pipelines.StackOutput",
		[]interface{}{artifactFile, outputName},
		&j,
	)

	return &j
}

// Build a StackOutput from a known artifact and an output name.
// Experimental.
func NewStackOutput_Override(s StackOutput, artifactFile awscodepipeline.ArtifactPath, outputName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.StackOutput",
		[]interface{}{artifactFile, outputName},
		s,
	)
}

// A Reference to a Stack Output.
// Experimental.
type StackOutputReference interface {
	OutputName() *string
	StackDescription() *string
	IsProducedBy(stack StackDeployment) *bool
}

// The jsii proxy struct for StackOutputReference
type jsiiProxy_StackOutputReference struct {
	_ byte // padding
}

func (j *jsiiProxy_StackOutputReference) OutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackOutputReference) StackDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackDescription",
		&returns,
	)
	return returns
}


// Create a StackOutputReference that references the given CfnOutput.
// Experimental.
func StackOutputReference_FromCfnOutput(output awscdk.CfnOutput) StackOutputReference {
	_init_.Initialize()

	var returns StackOutputReference

	_jsii_.StaticInvoke(
		"monocdk.pipelines.StackOutputReference",
		"fromCfnOutput",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Whether or not this stack output is being produced by the given Stack deployment.
// Experimental.
func (s *jsiiProxy_StackOutputReference) IsProducedBy(stack StackDeployment) *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"isProducedBy",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

// Deployment of a single `Stage`.
//
// A `Stage` consists of one or more `Stacks`, which will be
// deployed in dependency order.
// Experimental.
type StageDeployment interface {
	Post() *[]Step
	Pre() *[]Step
	Stacks() *[]StackDeployment
	StageName() *string
	AddPost(steps ...Step)
	AddPre(steps ...Step)
}

// The jsii proxy struct for StageDeployment
type jsiiProxy_StageDeployment struct {
	_ byte // padding
}

func (j *jsiiProxy_StageDeployment) Post() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"post",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageDeployment) Pre() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"pre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageDeployment) Stacks() *[]StackDeployment {
	var returns *[]StackDeployment
	_jsii_.Get(
		j,
		"stacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StageDeployment) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}


// Create a new `StageDeployment` from a `Stage`.
//
// Synthesizes the target stage, and deployes the stacks found inside
// in dependency order.
// Experimental.
func StageDeployment_FromStage(stage awscdk.Stage, props *StageDeploymentProps) StageDeployment {
	_init_.Initialize()

	var returns StageDeployment

	_jsii_.StaticInvoke(
		"monocdk.pipelines.StageDeployment",
		"fromStage",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

// Add an additional step to run after all of the stacks in this stage.
// Experimental.
func (s *jsiiProxy_StageDeployment) AddPost(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addPost",
		args,
	)
}

// Add an additional step to run before any of the stacks in this stage.
// Experimental.
func (s *jsiiProxy_StageDeployment) AddPre(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addPre",
		args,
	)
}

// Properties for a `StageDeployment`.
// Experimental.
type StageDeploymentProps struct {
	// Additional steps to run after all of the stacks in the stage.
	// Experimental.
	Post *[]Step `json:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Experimental.
	Pre *[]Step `json:"pre"`
	// Stage name to use in the pipeline.
	// Experimental.
	StageName *string `json:"stageName"`
}

// Options for a convention-based synth using NPM.
// Experimental.
type StandardNpmSynthOptions struct {
	// The artifact where the CloudAssembly should be emitted.
	// Experimental.
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Experimental.
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact"`
	// Name of the build action.
	// Experimental.
	ActionName *string `json:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Experimental.
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Experimental.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables"`
	// Name of the CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Experimental.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Experimental.
	Subdirectory *string `json:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The build command.
	//
	// By default, we assume NPM projects are either written in JavaScript or are
	// using `ts-node`, so don't need a build command.
	//
	// Otherwise, put the build command here, for example `npm run build`.
	// Experimental.
	BuildCommand *string `json:"buildCommand"`
	// The install command.
	// Experimental.
	InstallCommand *string `json:"installCommand"`
	// The synth command.
	// Experimental.
	SynthCommand *string `json:"synthCommand"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Experimental.
	TestCommands *[]*string `json:"testCommands"`
}

// Options for a convention-based synth using Yarn.
// Experimental.
type StandardYarnSynthOptions struct {
	// The artifact where the CloudAssembly should be emitted.
	// Experimental.
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Experimental.
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact"`
	// Name of the build action.
	// Experimental.
	ActionName *string `json:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Experimental.
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Experimental.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables"`
	// Name of the CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Experimental.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Experimental.
	Subdirectory *string `json:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The build command.
	//
	// By default, we assume NPM projects are either written in JavaScript or are
	// using `ts-node`, so don't need a build command.
	//
	// Otherwise, put the build command here, for example `npm run build`.
	// Experimental.
	BuildCommand *string `json:"buildCommand"`
	// The install command.
	// Experimental.
	InstallCommand *string `json:"installCommand"`
	// The synth command.
	// Experimental.
	SynthCommand *string `json:"synthCommand"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Experimental.
	TestCommands *[]*string `json:"testCommands"`
}

// A generic Step which can be added to a Pipeline.
//
// Steps can be used to add Sources, Build Actions and Validations
// to your pipeline.
//
// This class is abstract. See specific subclasses of Step for
// useful steps to add to your Pipeline
// Experimental.
type Step interface {
	IFileSetProducer
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	ConfigurePrimaryOutput(fs FileSet)
	ToString() *string
}

// The jsii proxy struct for Step
type jsiiProxy_Step struct {
	jsiiProxy_IFileSetProducer
}

func (j *jsiiProxy_Step) Dependencies() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Step) DependencyFileSets() *[]FileSet {
	var returns *[]FileSet
	_jsii_.Get(
		j,
		"dependencyFileSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Step) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Step) IsSource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Step) PrimaryOutput() FileSet {
	var returns FileSet
	_jsii_.Get(
		j,
		"primaryOutput",
		&returns,
	)
	return returns
}


// Experimental.
func NewStep_Override(s Step, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.Step",
		[]interface{}{id},
		s,
	)
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
// Experimental.
func (s *jsiiProxy_Step) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Configure the given FileSet as the primary output of this step.
// Experimental.
func (s *jsiiProxy_Step) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Return a string representation of this Step.
// Experimental.
func (s *jsiiProxy_Step) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Action to self-mutate the pipeline.
//
// Creates a CodeBuild project which will use the CDK CLI
// to deploy the pipeline stack.
//
// You do not need to instantiate this action -- it will automatically
// be added by the pipeline.
// Experimental.
type UpdatePipelineAction interface {
	awscdk.Construct
	awscodepipeline.IAction
	ActionProperties() *awscodepipeline.ActionProperties
	Node() awscdk.ConstructNode
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnPrepare()
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for UpdatePipelineAction
type jsiiProxy_UpdatePipelineAction struct {
	internal.Type__awscdkConstruct
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_UpdatePipelineAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UpdatePipelineAction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewUpdatePipelineAction(scope constructs.Construct, id *string, props *UpdatePipelineActionProps) UpdatePipelineAction {
	_init_.Initialize()

	j := jsiiProxy_UpdatePipelineAction{}

	_jsii_.Create(
		"monocdk.pipelines.UpdatePipelineAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUpdatePipelineAction_Override(u UpdatePipelineAction, scope constructs.Construct, id *string, props *UpdatePipelineActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.UpdatePipelineAction",
		[]interface{}{scope, id, props},
		u,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func UpdatePipelineAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.UpdatePipelineAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (u *jsiiProxy_UpdatePipelineAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{scope, stage, options},
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
func (u *jsiiProxy_UpdatePipelineAction) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

// Exists to implement IAction.
// Experimental.
func (u *jsiiProxy_UpdatePipelineAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		u,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (u *jsiiProxy_UpdatePipelineAction) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
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
func (u *jsiiProxy_UpdatePipelineAction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
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
func (u *jsiiProxy_UpdatePipelineAction) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (u *jsiiProxy_UpdatePipelineAction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (u *jsiiProxy_UpdatePipelineAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
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
func (u *jsiiProxy_UpdatePipelineAction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		u,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Props for the UpdatePipelineAction.
// Experimental.
type UpdatePipelineActionProps struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Experimental.
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput"`
	// Hierarchical id of the pipeline stack.
	// Experimental.
	PipelineStackHierarchicalId *string `json:"pipelineStackHierarchicalId"`
	// Custom BuildSpec that is merged with generated one.
	// Experimental.
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec"`
	// Version of CDK CLI to 'npm install'.
	// Experimental.
	CdkCliVersion *string `json:"cdkCliVersion"`
	// Docker registries and associated credentials necessary during the pipeline self-update stage.
	// Experimental.
	DockerCredentials *[]DockerCredential `json:"dockerCredentials"`
	// Name of the pipeline stack.
	// Deprecated: - Use `pipelineStackHierarchicalId` instead.
	PipelineStackName *string `json:"pipelineStackName"`
	// Whether the build step should run in privileged mode.
	// Experimental.
	Privileged *bool `json:"privileged"`
	// Name of the CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
}

// Multiple stages that are deployed in parallel.
// Experimental.
type Wave interface {
	Id() *string
	Post() *[]Step
	Pre() *[]Step
	Stages() *[]StageDeployment
	AddPost(steps ...Step)
	AddPre(steps ...Step)
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
}

// The jsii proxy struct for Wave
type jsiiProxy_Wave struct {
	_ byte // padding
}

func (j *jsiiProxy_Wave) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wave) Post() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"post",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wave) Pre() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"pre",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wave) Stages() *[]StageDeployment {
	var returns *[]StageDeployment
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}


// Experimental.
func NewWave(id *string, props *WaveProps) Wave {
	_init_.Initialize()

	j := jsiiProxy_Wave{}

	_jsii_.Create(
		"monocdk.pipelines.Wave",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWave_Override(w Wave, id *string, props *WaveProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.Wave",
		[]interface{}{id, props},
		w,
	)
}

// Add an additional step to run after all of the stages in this wave.
// Experimental.
func (w *jsiiProxy_Wave) AddPost(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addPost",
		args,
	)
}

// Add an additional step to run before any of the stages in this wave.
// Experimental.
func (w *jsiiProxy_Wave) AddPre(steps ...Step) {
	args := []interface{}{}
	for _, a := range steps {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		w,
		"addPre",
		args,
	)
}

// Add a Stage to this wave.
//
// It will be deployed in parallel with all other stages in this
// wave.
// Experimental.
func (w *jsiiProxy_Wave) AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment {
	var returns StageDeployment

	_jsii_.Invoke(
		w,
		"addStage",
		[]interface{}{stage, options},
		&returns,
	)

	return returns
}

// Options to pass to `addWave`.
// Experimental.
type WaveOptions struct {
	// Additional steps to run after all of the stages in the wave.
	// Experimental.
	Post *[]Step `json:"post"`
	// Additional steps to run before any of the stages in the wave.
	// Experimental.
	Pre *[]Step `json:"pre"`
}

// Construction properties for a `Wave`.
// Experimental.
type WaveProps struct {
	// Additional steps to run after all of the stages in the wave.
	// Experimental.
	Post *[]Step `json:"post"`
	// Additional steps to run before any of the stages in the wave.
	// Experimental.
	Pre *[]Step `json:"pre"`
}

