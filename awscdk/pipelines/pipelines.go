package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v10"
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

// Instructions to publish certain assets.
// Experimental.
type AssetPublishingCommand struct {
	// Asset identifier.
	// Experimental.
	AssetId *string `json:"assetId"`
	// Asset manifest path.
	// Experimental.
	AssetManifestPath *string `json:"assetManifestPath"`
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
	constructs.Construct
	CodePipeline() awscodepipeline.Pipeline
	Node() constructs.Node
	AddApplicationStage(appStage awscdk.Stage, options *AddStageOptions) CdkStage
	AddStage(stageName *string) CdkStage
	StackOutput(cfnOutput awscdk.CfnOutput) StackOutput
	Stage(stageName *string) awscodepipeline.IStage
	ToString() *string
}

// The jsii proxy struct for CdkPipeline
type jsiiProxy_CdkPipeline struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_CdkPipeline) Node() constructs.Node {
	var returns constructs.Node
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
		"aws-cdk-lib.pipelines.CdkPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkPipeline_Override(c CdkPipeline, scope constructs.Construct, id *string, props *CdkPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.CdkPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CdkPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CdkPipeline",
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
func (c *jsiiProxy_CdkPipeline) AddStage(stageName *string) CdkStage {
	var returns CdkStage

	_jsii_.Invoke(
		c,
		"addStage",
		[]interface{}{stageName},
		&returns,
	)

	return returns
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

// Properties for a CdkPipeline.
// Experimental.
type CdkPipelineProps struct {
	// The artifact you have defined to be the artifact to hold the cloudAssemblyArtifact for the synth action.
	// Experimental.
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact"`
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
	// The CodePipeline action used to retrieve the CDK app's source.
	// Experimental.
	SourceAction awscodepipeline.IAction `json:"sourceAction"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
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
	constructs.Construct
	Node() constructs.Node
	AddActions(actions ...awscodepipeline.IAction)
	AddApplication(appStage awscdk.Stage, options *AddStageOptions)
	AddManualApprovalAction(options *AddManualApprovalOptions)
	AddStackArtifactDeployment(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions)
	DeploysStack(artifactId *string) *bool
	NextSequentialRunOrder(count *float64) *float64
	ToString() *string
}

// The jsii proxy struct for CdkStage
type jsiiProxy_CdkStage struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CdkStage) Node() constructs.Node {
	var returns constructs.Node
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
		"aws-cdk-lib.pipelines.CdkStage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCdkStage_Override(c CdkStage, scope constructs.Construct, id *string, props *CdkStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.CdkStage",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CdkStage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CdkStage",
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
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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
		"aws-cdk-lib.pipelines.DeployCdkStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewDeployCdkStackAction_Override(d DeployCdkStackAction, props *DeployCdkStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.DeployCdkStackAction",
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
		"aws-cdk-lib.pipelines.DeployCdkStackAction",
		"fromStackArtifact",
		[]interface{}{scope, artifact, options},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (d *jsiiProxy_DeployCdkStackAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

// Action to publish an asset in the pipeline.
//
// Creates a CodeBuild project which will use the CDK CLI
// to prepare and publish the asset.
//
// You do not need to instantiate this action -- it will automatically
// be added by the pipeline when you add stacks that use assets.
// Experimental.
type PublishAssetsAction interface {
	constructs.Construct
	awscodepipeline.IAction
	ActionProperties() *awscodepipeline.ActionProperties
	Node() constructs.Node
	AddPublishCommand(relativeManifestPath *string, assetSelector *string)
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	ToString() *string
}

// The jsii proxy struct for PublishAssetsAction
type jsiiProxy_PublishAssetsAction struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_PublishAssetsAction) Node() constructs.Node {
	var returns constructs.Node
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
		"aws-cdk-lib.pipelines.PublishAssetsAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPublishAssetsAction_Override(p PublishAssetsAction, scope constructs.Construct, id *string, props *PublishAssetsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.PublishAssetsAction",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PublishAssetsAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.PublishAssetsAction",
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
func (p *jsiiProxy_PublishAssetsAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
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
	// Version of CDK CLI to 'npm install'.
	// Experimental.
	CdkCliVersion *string `json:"cdkCliVersion"`
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

// Validate a revision using shell commands.
// Experimental.
type ShellScriptAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	ActionProperties() *awscodepipeline.ActionProperties
	GrantPrincipal() awsiam.IPrincipal
	Project() awscodebuild.IProject
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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
		"aws-cdk-lib.pipelines.ShellScriptAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewShellScriptAction_Override(s ShellScriptAction, props *ShellScriptActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ShellScriptAction",
		[]interface{}{props},
		s,
	)
}

// Exists to implement IAction.
// Experimental.
func (s *jsiiProxy_ShellScriptAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

// A standard synth with a generated buildspec.
// Experimental.
type SimpleSynthAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	ActionProperties() *awscodepipeline.ActionProperties
	GrantPrincipal() awsiam.IPrincipal
	Project() awscodebuild.IProject
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
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
		"aws-cdk-lib.pipelines.SimpleSynthAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSimpleSynthAction_Override(s SimpleSynthAction, props *SimpleSynthActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.SimpleSynthAction",
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
		"aws-cdk-lib.pipelines.SimpleSynthAction",
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
		"aws-cdk-lib.pipelines.SimpleSynthAction",
		"standardYarnSynth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (s *jsiiProxy_SimpleSynthAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
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
	// The build commands.
	//
	// If your programming language requires a compilation step, put the
	// compilation command here.
	// Experimental.
	BuildCommands *[]*string `json:"buildCommands"`
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
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
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
		"aws-cdk-lib.pipelines.StackOutput",
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
		"aws-cdk-lib.pipelines.StackOutput",
		[]interface{}{artifactFile, outputName},
		s,
	)
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
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
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
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Experimental.
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Experimental.
	Environment *awscodebuild.BuildEnvironment `json:"environment"`
	// Environment variables to send into build.
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
	constructs.Construct
	awscodepipeline.IAction
	ActionProperties() *awscodepipeline.ActionProperties
	Node() constructs.Node
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	ToString() *string
}

// The jsii proxy struct for UpdatePipelineAction
type jsiiProxy_UpdatePipelineAction struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_UpdatePipelineAction) Node() constructs.Node {
	var returns constructs.Node
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
		"aws-cdk-lib.pipelines.UpdatePipelineAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewUpdatePipelineAction_Override(u UpdatePipelineAction, scope constructs.Construct, id *string, props *UpdatePipelineActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.UpdatePipelineAction",
		[]interface{}{scope, id, props},
		u,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func UpdatePipelineAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.UpdatePipelineAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Exists to implement IAction.
// Experimental.
func (u *jsiiProxy_UpdatePipelineAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
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

// Props for the UpdatePipelineAction.
// Experimental.
type UpdatePipelineActionProps struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Experimental.
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput"`
	// Name of the pipeline stack.
	// Experimental.
	PipelineStackName *string `json:"pipelineStackName"`
	// Version of CDK CLI to 'npm install'.
	// Experimental.
	CdkCliVersion *string `json:"cdkCliVersion"`
	// Name of the CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName"`
}

