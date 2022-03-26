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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   addManualApprovalOptions := &addManualApprovalOptions{
//   	actionName: jsii.String("actionName"),
//   	runOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AddManualApprovalOptions struct {
	// The name of the manual approval action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `json:"actionName" yaml:"actionName"`
	// The runOrder for this action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
}

// Additional options for adding a stack deployment.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   addStackOptions := &addStackOptions{
//   	executeRunOrder: jsii.Number(123),
//   	runOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AddStackOptions struct {
	// Base runorder.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `json:"executeRunOrder" yaml:"executeRunOrder"`
	// Base runorder.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
}

// Options for adding an application stage to a pipeline.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var topic topic
//   addStageOptions := &addStageOptions{
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	extraRunOrderSpace: jsii.Number(123),
//   	manualApprovals: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ConfirmBroadeningPermissions *bool `json:"confirmBroadeningPermissions" yaml:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when the security check registers changes within the application.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SecurityNotificationTopic awssns.ITopic `json:"securityNotificationTopic" yaml:"securityNotificationTopic"`
	// Add room for extra actions.
	//
	// You can use this to make extra room in the runOrder sequence between the
	// changeset 'prepare' and 'execute' actions and insert your own actions there.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExtraRunOrderSpace *float64 `json:"extraRunOrderSpace" yaml:"extraRunOrderSpace"`
	// Add manual approvals before executing change sets.
	//
	// This gives humans the opportunity to confirm the change set looks alright
	// before deploying it.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ManualApprovals *bool `json:"manualApprovals" yaml:"manualApprovals"`
}

// Options to pass to `addStage`.
//
// Example:
//   var pipeline codePipeline
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//
//   pipeline.addStage(preprod, &addStageOpts{
//   	post: []step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &shellStepProps{
//   			commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.addStage(prod, &addStageOpts{
//   	pre: []*step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd")),
//   	},
//   })
//
// Experimental.
type AddStageOpts struct {
	// Additional steps to run after all of the stacks in the stage.
	// Experimental.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Experimental.
	Pre *[]Step `json:"pre" yaml:"pre"`
	// Instructions for stack level steps.
	// Experimental.
	StackSteps *[]*StackSteps `json:"stackSteps" yaml:"stackSteps"`
}

// Specification of an additional artifact to generate.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   additionalArtifact := &additionalArtifact{
//   	artifact: artifact,
//   	directory: jsii.String("directory"),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AdditionalArtifact struct {
	// Artifact to represent the build directory in the pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Artifact awscodepipeline.Artifact `json:"artifact" yaml:"artifact"`
	// Directory to be packaged.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Directory *string `json:"directory" yaml:"directory"`
}

// Translate FileSets to CodePipeline Artifacts.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   artifactMap := pipelines.NewArtifactMap()
//
// Experimental.
type ArtifactMap interface {
	// Return the matching CodePipeline artifact for a FileSet.
	// Experimental.
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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   assetPublishingCommand := &assetPublishingCommand{
//   	assetId: jsii.String("assetId"),
//   	assetManifestPath: jsii.String("assetManifestPath"),
//   	assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   	assetSelector: jsii.String("assetSelector"),
//   	assetType: pipelines.assetType_FILE,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type AssetPublishingCommand struct {
	// Asset identifier.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetId *string `json:"assetId" yaml:"assetId"`
	// Asset manifest path.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetManifestPath *string `json:"assetManifestPath" yaml:"assetManifestPath"`
	// ARN of the IAM Role used to publish this asset.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetPublishingRoleArn *string `json:"assetPublishingRoleArn" yaml:"assetPublishingRoleArn"`
	// Asset selector to pass to `cdk-assets`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetSelector *string `json:"assetSelector" yaml:"assetSelector"`
	// Type of asset to publish.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetType AssetType `json:"assetType" yaml:"assetType"`
}

// Type of the asset that is being published.
// Experimental.
type AssetType string

const (
	// A file.
	// Experimental.
	AssetType_FILE AssetType = "FILE"
	// A Docker image.
	// Experimental.
	AssetType_DOCKER_IMAGE AssetType = "DOCKER_IMAGE"
)

// Base options for a pipelines stage.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var topic topic
//   baseStageOptions := &baseStageOptions{
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ConfirmBroadeningPermissions *bool `json:"confirmBroadeningPermissions" yaml:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when the security check registers changes within the application.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SecurityNotificationTopic awssns.ITopic `json:"securityNotificationTopic" yaml:"securityNotificationTopic"`
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
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &cdkPipelineProps{
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   	synthAction: pipelines.simpleSynthAction.standardNpmSynth(&standardNpmSynthOptions{
//   		sourceArtifact: sourceArtifact,
//   		cloudAssemblyArtifact: cloudAssemblyArtifact,
//   		environment: &buildEnvironment{
//   			privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkPipeline interface {
	awscdk.Construct
	// The underlying CodePipeline object.
	//
	// You can use this to add more Stages to the pipeline, or Actions
	// to Stages.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CodePipeline() awscodepipeline.Pipeline
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
	// Add pipeline stage that will deploy the given application stage.
	//
	// The application construct should subclass `Stage` and can contain any
	// number of `Stacks` inside it that may have dependency relationships
	// on one another.
	//
	// All stacks in the application will be deployed in the appropriate order,
	// and all assets found in the application will be added to the asset
	// publishing stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddApplicationStage(appStage awscdk.Stage, options *AddStageOptions) CdkStage
	// Add a new, empty stage to the pipeline.
	//
	// Prefer to use `addApplicationStage` if you are intended to deploy a CDK
	// application, but you can use this method if you want to add other kinds of
	// Actions to a pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddStage(stageName *string, options *BaseStageOptions) CdkStage
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Prepare()
	// Get the StackOutput object that holds this CfnOutput's value in this pipeline.
	//
	// `StackOutput` can be used in validation actions later in the pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackOutput(cfnOutput awscdk.CfnOutput) StackOutput
	// Access one of the pipeline's stages by stage name.
	//
	// You can use this to add more Actions to a stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Stage(stageName *string) awscodepipeline.IStage
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ToString() *string
	// Validate that we don't have any stacks violating dependency order in the pipeline.
	//
	// Our own convenience methods will never generate a pipeline that does that (although
	// this is a nice verification), but a user can also add the stacks by hand.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewCdkPipeline_Override(c CdkPipeline, scope constructs.Construct, id *string, props *CdkPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CdkPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

func (c *jsiiProxy_CdkPipeline) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkPipeline) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CdkPipeline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

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

func (c *jsiiProxy_CdkPipeline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &cdkPipelineProps{
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   	synthAction: pipelines.simpleSynthAction.standardNpmSynth(&standardNpmSynthOptions{
//   		sourceArtifact: sourceArtifact,
//   		cloudAssemblyArtifact: cloudAssemblyArtifact,
//   		environment: &buildEnvironment{
//   			privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkPipelineProps struct {
	// The artifact you have defined to be the artifact to hold the cloudAssemblyArtifact for the synth action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// Custom BuildSpec that is merged with generated one (for asset publishing actions).
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetBuildSpec awscodebuild.BuildSpec `json:"assetBuildSpec" yaml:"assetBuildSpec"`
	// Additional commands to run before installing cdk-assets during the asset publishing step Use this to setup proxies or npm mirrors.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetPreInstallCommands *[]*string `json:"assetPreInstallCommands" yaml:"assetPreInstallCommands"`
	// CDK CLI version to use in pipeline.
	//
	// Some Actions in the pipeline will download and run a version of the CDK
	// CLI. Specify the version here.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CdkCliVersion *string `json:"cdkCliVersion" yaml:"cdkCliVersion"`
	// Existing CodePipeline to add deployment stages to.
	//
	// Use this if you want more control over the CodePipeline that gets created.
	// You can choose to not pass this value, in which case a new CodePipeline is
	// created with default settings.
	//
	// If you pass an existing CodePipeline, it should have been created
	// with `restartExecutionOnUpdate: true`.
	//
	// [disable-awslint:ref-via-interface].
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CodePipeline awscodepipeline.Pipeline `json:"codePipeline" yaml:"codePipeline"`
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
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CrossAccountKeys *bool `json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DockerCredentials *[]DockerCredential `json:"dockerCredentials" yaml:"dockerCredentials"`
	// Enables KMS key rotation for cross-account keys.
	//
	// Cannot be set if `crossAccountKeys` was set to `false`.
	//
	// Key rotation costs $1/month when enabled.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnableKeyRotation *bool `json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// Name of the pipeline.
	//
	// Can only be set if `codePipeline` is not set.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PipelineName *string `json:"pipelineName" yaml:"pipelineName"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SelfMutating *bool `json:"selfMutating" yaml:"selfMutating"`
	// Custom BuildSpec that is merged with generated one (for self-mutation stage).
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SelfMutationBuildSpec awscodebuild.BuildSpec `json:"selfMutationBuildSpec" yaml:"selfMutationBuildSpec"`
	// Whether this pipeline creates one asset upload action per asset type or one asset upload per asset.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SinglePublisherPerType *bool `json:"singlePublisherPerType" yaml:"singlePublisherPerType"`
	// The CodePipeline action used to retrieve the CDK app's source.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceAction awscodepipeline.IAction `json:"sourceAction" yaml:"sourceAction"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
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
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SupportDockerAssets *bool `json:"supportDockerAssets" yaml:"supportDockerAssets"`
	// The CodePipeline action build and synthesis step of the CDK app.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SynthAction awscodepipeline.IAction `json:"synthAction" yaml:"synthAction"`
	// The VPC where to execute the CdkPipeline actions.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Options for the 'fromStackArtifact' operation.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   cdkStackActionFromArtifactOptions := &cdkStackActionFromArtifactOptions{
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	baseActionName: jsii.String("baseActionName"),
//   	changeSetName: jsii.String("changeSetName"),
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   	stackName: jsii.String("stackName"),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkStackActionFromArtifactOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Base name of the action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BaseActionName *string `json:"baseActionName" yaml:"baseActionName"`
	// Name of the change set to create and deploy.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ChangeSetName *string `json:"changeSetName" yaml:"changeSetName"`
	// Run order for the Execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `json:"executeRunOrder" yaml:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Output awscodepipeline.Artifact `json:"output" yaml:"output"`
	// Filename in output to write Stack outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OutputFileName *string `json:"outputFileName" yaml:"outputFileName"`
	// Run order for the Prepare action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder *float64 `json:"prepareRunOrder" yaml:"prepareRunOrder"`
	// The name of the stack that should be created/updated.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackName *string `json:"stackName" yaml:"stackName"`
}

// Stage in a CdkPipeline.
//
// You don't need to instantiate this class directly. Use
// `cdkPipeline.addStage()` instead.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var stage iStage
//   var stageHost iStageHost
//   var topic topic
//   cdkStage := pipelines.NewCdkStage(this, jsii.String("MyCdkStage"), &cdkStageProps{
//   	cloudAssemblyArtifact: artifact,
//   	host: stageHost,
//   	pipelineStage: stage,
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkStage interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
	// Add one or more CodePipeline Actions.
	//
	// You need to make sure it is created with the right runOrder. Call `nextSequentialRunOrder()`
	// for every action to get actions to execute in sequence.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddActions(actions ...awscodepipeline.IAction)
	// Add all stacks in the application Stage to this stage.
	//
	// The application construct should subclass `Stage` and can contain any
	// number of `Stacks` inside it that may have dependency relationships
	// on one another.
	//
	// All stacks in the application will be deployed in the appropriate order,
	// and all assets found in the application will be added to the asset
	// publishing stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddApplication(appStage awscdk.Stage, options *AddStageOptions)
	// Add a manual approval action.
	//
	// If you need more flexibility than what this method offers,
	// use `addAction` with a `ManualApprovalAction`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddManualApprovalAction(options *AddManualApprovalOptions)
	// Add a deployment action based on a stack artifact.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddStackArtifactDeployment(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions)
	// Whether this Stage contains an action to deploy the given stack, identified by its artifact ID.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DeploysStack(artifactId *string) *bool
	// Return the runOrder number necessary to run the next Action in sequence with the rest.
	//
	// FIXME: This is here because Actions are immutable and can't be reordered
	// after creation, nor is there a way to specify relative priorities, which
	// is a limitation that we should take away in the base library.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	NextSequentialRunOrder(count *float64) *float64
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewCdkStage_Override(c CdkStage, scope constructs.Construct, id *string, props *CdkStageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CdkStage",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

func (c *jsiiProxy_CdkStage) AddApplication(appStage awscdk.Stage, options *AddStageOptions) {
	_jsii_.InvokeVoid(
		c,
		"addApplication",
		[]interface{}{appStage, options},
	)
}

func (c *jsiiProxy_CdkStage) AddManualApprovalAction(options *AddManualApprovalOptions) {
	_jsii_.InvokeVoid(
		c,
		"addManualApprovalAction",
		[]interface{}{options},
	)
}

func (c *jsiiProxy_CdkStage) AddStackArtifactDeployment(stackArtifact cxapi.CloudFormationStackArtifact, options *AddStackOptions) {
	_jsii_.InvokeVoid(
		c,
		"addStackArtifactDeployment",
		[]interface{}{stackArtifact, options},
	)
}

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

func (c *jsiiProxy_CdkStage) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkStage) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CdkStage) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkStage) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import sns "github.com/aws/aws-cdk-go/awscdk/aws_sns"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var stage iStage
//   var stageHost iStageHost
//   var topic topic
//   cdkStageProps := &cdkStageProps{
//   	cloudAssemblyArtifact: artifact,
//   	host: stageHost,
//   	pipelineStage: stage,
//   	stageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	confirmBroadeningPermissions: jsii.Boolean(false),
//   	securityNotificationTopic: topic,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkStageProps struct {
	// The CodePipeline Artifact with the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// Features the Stage needs from its environment.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Host IStageHost `json:"host" yaml:"host"`
	// The underlying Pipeline Stage associated with thisCdkStage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PipelineStage awscodepipeline.IStage `json:"pipelineStage" yaml:"pipelineStage"`
	// Name of the stage that should be created.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StageName *string `json:"stageName" yaml:"stageName"`
	// Run a security check before every application prepare/deploy actions.
	//
	// Note: Stage level security check can be overriden per application as follows:
	//    `stage.addApplication(app, { confirmBroadeningPermissions: false })`
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ConfirmBroadeningPermissions *bool `json:"confirmBroadeningPermissions" yaml:"confirmBroadeningPermissions"`
	// Optional SNS topic to send notifications to when any security check registers changes within a application.
	//
	// Note: The Stage Notification Topic can be overriden per application as follows:
	//    `stage.addApplication(app, { securityNotificationTopic: newTopic })`
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SecurityNotificationTopic awssns.ITopic `json:"securityNotificationTopic" yaml:"securityNotificationTopic"`
}

// Options for customizing a single CodeBuild project.
//
// Example:
//   var vpc vpc
//   var mySecurityGroup securityGroup
//   pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	// Standard CodePipeline properties
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//
//   	// Defaults for all CodeBuild projects
//   	codeBuildDefaults: &codeBuildOptions{
//   		// Prepend commands and configuration to all projects
//   		partialBuildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   			"version": jsii.String("0.2"),
//   		}),
//
//   		// Control the build environment
//   		buildEnvironment: &buildEnvironment{
//   			computeType: codebuild.computeType_LARGE,
//   		},
//
//   		// Control Elastic Network Interface creation
//   		vpc: vpc,
//   		subnetSelection: &subnetSelection{
//   			subnetType: ec2.subnetType_PRIVATE,
//   		},
//   		securityGroups: []iSecurityGroup{
//   			mySecurityGroup,
//   		},
//
//   		// Additional policy statements for the execution role
//   		rolePolicy: []policyStatement{
//   			iam.NewPolicyStatement(&policyStatementProps{
//   			}),
//   		},
//   	},
//
//   	synthCodeBuildDefaults: &codeBuildOptions{
//   	},
//   	assetPublishingCodeBuildDefaults: &codeBuildOptions{
//   	},
//   	selfMutationCodeBuildDefaults: &codeBuildOptions{
//   	},
//   })
//
// Experimental.
type CodeBuildOptions struct {
	// Partial build environment, will be combined with other build environments that apply.
	// Experimental.
	BuildEnvironment *awscodebuild.BuildEnvironment `json:"buildEnvironment" yaml:"buildEnvironment"`
	// Partial buildspec, will be combined with other buildspecs that apply.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	// Experimental.
	PartialBuildSpec awscodebuild.BuildSpec `json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Policy statements to add to role.
	// Experimental.
	RolePolicy *[]awsiam.PolicyStatement `json:"rolePolicy" yaml:"rolePolicy"`
	// Which security group(s) to associate with the project network interfaces.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The VPC where to create the CodeBuild network interfaces in.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Run a script as a CodeBuild Project.
//
// The BuildSpec must be available inline--it cannot reference a file
// on disk. If your current build instructions are in a file like
// `buildspec.yml` in your repository, extract them to a script
// (say, `build.sh`) and invoke that script as part of the build:
//
// ```ts
// new pipelines.CodeBuildStep('Synth', {
//    commands: ['./build.sh'],
// });
// ```.
//
// Example:
//   pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	synth: pipelines.NewCodeBuildStep(jsii.String("Synth"), &codeBuildStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("..."),
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   			jsii.String("..."),
//   		},
//   		rolePolicyStatements: []policyStatement{
//   			iam.NewPolicyStatement(&policyStatementProps{
//   				actions: []*string{
//   					jsii.String("sts:AssumeRole"),
//   				},
//   				resources: []*string{
//   					jsii.String("*"),
//   				},
//   				conditions: map[string]interface{}{
//   					"StringEquals": map[string]*string{
//   						"iam:ResourceTag/aws-cdk:bootstrap-role": jsii.String("lookup"),
//   					},
//   				},
//   			}),
//   		},
//   	}),
//   })
//
// Experimental.
type CodeBuildStep interface {
	ShellStep
	// Build environment.
	// Experimental.
	BuildEnvironment() *awscodebuild.BuildEnvironment
	// Commands to run.
	// Experimental.
	Commands() *[]*string
	// Return the steps this step depends on, based on the FileSets it requires.
	// Experimental.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	// Experimental.
	DependencyFileSets() *[]FileSet
	// Environment variables to set.
	// Experimental.
	Env() *map[string]*string
	// Set environment variables based on Stack Outputs.
	// Experimental.
	EnvFromCfnOutputs() *map[string]StackOutputReference
	// The CodeBuild Project's principal.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Identifier for this step.
	// Experimental.
	Id() *string
	// Input FileSets.
	//
	// A list of `(FileSet, directory)` pairs, which are a copy of the
	// input properties. This list should not be modified directly.
	// Experimental.
	Inputs() *[]*FileSetLocation
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Experimental.
	InstallCommands() *[]*string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	// Experimental.
	IsSource() *bool
	// Output FileSets.
	//
	// A list of `(FileSet, directory)` pairs, which are a copy of the
	// input properties. This list should not be modified directly.
	// Experimental.
	Outputs() *[]*FileSetLocation
	// Additional configuration that can only be configured via BuildSpec.
	//
	// Contains exported variables.
	// Experimental.
	PartialBuildSpec() awscodebuild.BuildSpec
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	// Experimental.
	PrimaryOutput() FileSet
	// CodeBuild Project generated for the pipeline.
	//
	// Will only be available after the pipeline has been built.
	// Experimental.
	Project() awscodebuild.IProject
	// Name for the generated CodeBuild project.
	// Experimental.
	ProjectName() *string
	// Custom execution role to be used for the CodeBuild project.
	// Experimental.
	Role() awsiam.IRole
	// Policy statements to add to role used during the synth.
	// Experimental.
	RolePolicyStatements() *[]awsiam.PolicyStatement
	// Which security group to associate with the script's project network interfaces.
	// Experimental.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// Which subnets to use.
	// Experimental.
	SubnetSelection() *awsec2.SubnetSelection
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout() awscdk.Duration
	// The VPC where to execute the SimpleSynth.
	// Experimental.
	Vpc() awsec2.IVpc
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	// Experimental.
	AddDependencyFileSet(fs FileSet)
	// Add an additional output FileSet based on a directory.
	//
	// After running the script, the contents of the given directory
	// will be exported as a `FileSet`. Use the `FileSet` as the
	// input to another step.
	//
	// Multiple calls with the exact same directory name string (not normalized)
	// will return the same FileSet.
	// Experimental.
	AddOutputDirectory(directory *string) FileSet
	// Add a dependency on another step.
	// Experimental.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	// Experimental.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
	// Experimental.
	DiscoverReferencedOutputs(structure interface{})
	// Reference a CodePipeline variable defined by the CodeBuildStep.
	//
	// The variable must be set in the shell of the CodeBuild step when
	// it finishes its `post_build` phase.
	//
	// Example:
	//   // Access the output of one CodeBuildStep in another CodeBuildStep
	//   var pipeline codePipeline
	//
	//   step1 := pipelines.NewCodeBuildStep(jsii.String("Step1"), &codeBuildStepProps{
	//   	commands: []*string{
	//   		jsii.String("export MY_VAR=hello"),
	//   	},
	//   })
	//
	//   step2 := pipelines.NewCodeBuildStep(jsii.String("Step2"), &codeBuildStepProps{
	//   	env: map[string]*string{
	//   		"IMPORTED_VAR": step1.exportedVariable(jsii.String("MY_VAR")),
	//   	},
	//   	commands: []*string{
	//   		jsii.String("echo $IMPORTED_VAR"),
	//   	},
	//   })
	//
	// Experimental.
	ExportedVariable(variableName *string) *string
	// Configure the given output directory as primary output.
	//
	// If no primary output has been configured yet, this directory
	// will become the primary output of this ShellStep, otherwise this
	// method will throw if the given directory is different than the
	// currently configured primary output directory.
	// Experimental.
	PrimaryOutputDirectory(directory *string) FileSet
	// Return a string representation of this Step.
	// Experimental.
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

func (j *jsiiProxy_CodeBuildStep) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
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

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
// Experimental.
func CodeBuildStep_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodeBuildStep",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStep) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

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

func (c *jsiiProxy_CodeBuildStep) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (c *jsiiProxy_CodeBuildStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_CodeBuildStep) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		c,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

func (c *jsiiProxy_CodeBuildStep) ExportedVariable(variableName *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"exportedVariable",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

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
//
// Example:
//   pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	synth: pipelines.NewCodeBuildStep(jsii.String("Synth"), &codeBuildStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("..."),
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   			jsii.String("..."),
//   		},
//   		rolePolicyStatements: []policyStatement{
//   			iam.NewPolicyStatement(&policyStatementProps{
//   				actions: []*string{
//   					jsii.String("sts:AssumeRole"),
//   				},
//   				resources: []*string{
//   					jsii.String("*"),
//   				},
//   				conditions: map[string]interface{}{
//   					"StringEquals": map[string]*string{
//   						"iam:ResourceTag/aws-cdk:bootstrap-role": jsii.String("lookup"),
//   					},
//   				},
//   			}),
//   		},
//   	}),
//   })
//
// Experimental.
type CodeBuildStepProps struct {
	// Commands to run.
	// Experimental.
	Commands *[]*string `json:"commands" yaml:"commands"`
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
	// const script = new pipelines.ShellStep('MainScript', {
	//    commands: ['npm ci','npm run build','npx cdk synth'],
	//    input: pipelines.CodePipelineSource.gitHub('org/source1', 'main'),
	//    additionalInputs: {
	//      '../siblingdir': pipelines.CodePipelineSource.gitHub('org/source2', 'main'),
	//    }
	// });
	// ```.
	// Experimental.
	AdditionalInputs *map[string]IFileSetProducer `json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	// Experimental.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	// Experimental.
	Input IFileSetProducer `json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Experimental.
	InstallCommands *[]*string `json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	// Experimental.
	PrimaryOutputDirectory *string `json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
	// Changes to environment.
	//
	// This environment will be combined with the pipeline's default
	// environment.
	// Experimental.
	BuildEnvironment *awscodebuild.BuildEnvironment `json:"buildEnvironment" yaml:"buildEnvironment"`
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
	PartialBuildSpec awscodebuild.BuildSpec `json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Name for the generated CodeBuild project.
	// Experimental.
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Custom execution role to be used for the CodeBuild project.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Experimental.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The VPC where to execute the SimpleSynth.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Configuration options for a CodeCommit source.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline_actions "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline_actions"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var role role
//   codeCommitSourceOptions := &codeCommitSourceOptions{
//   	codeBuildCloneOutput: jsii.Boolean(false),
//   	eventRole: role,
//   	trigger: codepipeline_actions.codeCommitTrigger_NONE,
//   }
//
// Experimental.
type CodeCommitSourceOptions struct {
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html
	//
	// Experimental.
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Role to be used by on commit event rule.
	//
	// Used only when trigger value is CodeCommitTrigger.EVENTS.
	// Experimental.
	EventRole awsiam.IRole `json:"eventRole" yaml:"eventRole"`
	// How should CodePipeline detect source changes for this Action.
	// Experimental.
	Trigger awscodepipelineactions.CodeCommitTrigger `json:"trigger" yaml:"trigger"`
}

// A CDK Pipeline that uses CodePipeline to deploy CDK apps.
//
// This is a `Pipeline` with its `engine` property set to
// `CodePipelineEngine`, and exists for nicer ergonomics for
// users that don't need to switch out engines.
//
// Example:
//   // Modern API
//   modernPipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	selfMutation: jsii.Boolean(false),
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   })
//
//   // Original API
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   originalPipeline := pipelines.NewCdkPipeline(this, jsii.String("Pipeline"), &cdkPipelineProps{
//   	selfMutating: jsii.Boolean(false),
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   })
//
// Experimental.
type CodePipeline interface {
	PipelineBase
	// The FileSet tha contains the cloud assembly.
	//
	// This is the primary output of the synth step.
	// Experimental.
	CloudAssemblyFileSet() FileSet
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The CodePipeline pipeline that deploys the CDK app.
	//
	// Only available after the pipeline has been built.
	// Experimental.
	Pipeline() awscodepipeline.Pipeline
	// The build step that produces the CDK Cloud Assembly.
	// Experimental.
	Synth() IFileSetProducer
	// The CodeBuild project that performs the Synth.
	//
	// Only available after the pipeline has been built.
	// Experimental.
	SynthProject() awscodebuild.IProject
	// The waves in this pipeline.
	// Experimental.
	Waves() *[]Wave
	// Deploy a single Stage by itself.
	//
	// Add a Stage to the pipeline, to be deployed in sequence with other
	// Stages added to the pipeline. All Stacks in the stage will be deployed
	// in an order automatically determined by their relative dependencies.
	// Experimental.
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
	// Add a Wave to the pipeline, for deploying multiple Stages in parallel.
	//
	// Use the return object of this method to deploy multiple stages in parallel.
	//
	// Example:
	//
	// ```ts
	// declare const pipeline: pipelines.CodePipeline;
	//
	// const wave = pipeline.addWave('MyWave');
	// wave.addStage(new MyApplicationStage(this, 'Stage1'));
	// wave.addStage(new MyApplicationStage(this, 'Stage2'));
	// ```.
	// Experimental.
	AddWave(id *string, options *WaveOptions) Wave
	// Send the current pipeline definition to the engine, and construct the pipeline.
	//
	// It is not possible to modify the pipeline after calling this method.
	// Experimental.
	BuildPipeline()
	// Implemented by subclasses to do the actual pipeline construction.
	// Experimental.
	DoBuildPipeline()
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
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

func (c *jsiiProxy_CodePipeline) BuildPipeline() {
	_jsii_.InvokeVoid(
		c,
		"buildPipeline",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodePipeline) DoBuildPipeline() {
	_jsii_.InvokeVoid(
		c,
		"doBuildPipeline",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodePipeline) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodePipeline) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CodePipeline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodePipeline) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var project project
//   codePipelineActionFactoryResult := &codePipelineActionFactoryResult{
//   	runOrdersConsumed: jsii.Number(123),
//
//   	// the properties below are optional
//   	project: project,
//   }
//
// Experimental.
type CodePipelineActionFactoryResult struct {
	// How many RunOrders were consumed.
	//
	// If you add 1 action, return the value 1 here.
	// Experimental.
	RunOrdersConsumed *float64 `json:"runOrdersConsumed" yaml:"runOrdersConsumed"`
	// If a CodeBuild project got created, the project.
	// Experimental.
	Project awscodebuild.IProject `json:"project" yaml:"project"`
}

// A FileSet created from a CodePipeline artifact.
//
// You only need to use this if you want to add CDK Pipeline stages
// add the end of an existing CodePipeline, which should be very rare.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   codePipelineFileSet := pipelines.codePipelineFileSet.fromArtifact(artifact)
//
// Experimental.
type CodePipelineFileSet interface {
	FileSet
	// Human-readable descriptor for this file set (does not need to be unique).
	// Experimental.
	Id() *string
	// The primary output of a file set producer.
	//
	// The primary output of a FileSet is itself.
	// Experimental.
	PrimaryOutput() FileSet
	// The Step that produces this FileSet.
	// Experimental.
	Producer() Step
	// Mark the given Step as the producer for this FileSet.
	//
	// This method can only be called once.
	// Experimental.
	ProducedBy(producer Step)
	// Return a string representation of this FileSet.
	// Experimental.
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

func (c *jsiiProxy_CodePipelineFileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		c,
		"producedBy",
		[]interface{}{producer},
	)
}

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
//
// Example:
//   // Modern API
//   modernPipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	selfMutation: jsii.Boolean(false),
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   })
//
//   // Original API
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   originalPipeline := pipelines.NewCdkPipeline(this, jsii.String("Pipeline"), &cdkPipelineProps{
//   	selfMutating: jsii.Boolean(false),
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   })
//
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
	Synth IFileSetProducer `json:"synth" yaml:"synth"`
	// Additional customizations to apply to the asset publishing CodeBuild projects.
	// Experimental.
	AssetPublishingCodeBuildDefaults *CodeBuildOptions `json:"assetPublishingCodeBuildDefaults" yaml:"assetPublishingCodeBuildDefaults"`
	// CDK CLI version to use in self-mutation and asset publishing steps.
	//
	// If you want to lock the CDK CLI version used in the pipeline, by steps
	// that are automatically generated for you, specify the version here.
	//
	// We recommend you do not specify this value, as not specifying it always
	// uses the latest CLI version which is backwards compatible with old versions.
	//
	// If you do specify it, be aware that this version should always be equal to or higher than the
	// version of the CDK framework used by the CDK app, when the CDK commands are
	// run during your pipeline execution. When you change this version, the *next
	// time* the `SelfMutate` step runs it will still be using the CLI of the the
	// *previous* version that was in this property: it will only start using the
	// new version after `SelfMutate` completes successfully. That means that if
	// you want to update both framework and CLI version, you should update the
	// CLI version first, commit, push and deploy, and only then update the
	// framework version.
	// Experimental.
	CliVersion *string `json:"cliVersion" yaml:"cliVersion"`
	// Customize the CodeBuild projects created for this pipeline.
	// Experimental.
	CodeBuildDefaults *CodeBuildOptions `json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An existing Pipeline to be reused and built upon.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	CodePipeline awscodepipeline.Pipeline `json:"codePipeline" yaml:"codePipeline"`
	// Create KMS keys for the artifact buckets, allowing cross-account deployments.
	//
	// The artifact buckets have to be encrypted to support deploying CDK apps to
	// another account, so if you want to do that or want to have your artifact
	// buckets encrypted, be sure to set this value to `true`.
	//
	// Be aware there is a cost associated with maintaining the KMS keys.
	// Experimental.
	CrossAccountKeys *bool `json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
	// Experimental.
	DockerCredentials *[]DockerCredential `json:"dockerCredentials" yaml:"dockerCredentials"`
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
	DockerEnabledForSelfMutation *bool `json:"dockerEnabledForSelfMutation" yaml:"dockerEnabledForSelfMutation"`
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
	DockerEnabledForSynth *bool `json:"dockerEnabledForSynth" yaml:"dockerEnabledForSynth"`
	// The name of the CodePipeline pipeline.
	// Experimental.
	PipelineName *string `json:"pipelineName" yaml:"pipelineName"`
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
	PublishAssetsInParallel *bool `json:"publishAssetsInParallel" yaml:"publishAssetsInParallel"`
	// Reuse the same cross region support stack for all pipelines in the App.
	// Experimental.
	ReuseCrossRegionSupportStacks *bool `json:"reuseCrossRegionSupportStacks" yaml:"reuseCrossRegionSupportStacks"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	// Experimental.
	SelfMutation *bool `json:"selfMutation" yaml:"selfMutation"`
	// Additional customizations to apply to the self mutation CodeBuild projects.
	// Experimental.
	SelfMutationCodeBuildDefaults *CodeBuildOptions `json:"selfMutationCodeBuildDefaults" yaml:"selfMutationCodeBuildDefaults"`
	// Additional customizations to apply to the synthesize CodeBuild projects.
	// Experimental.
	SynthCodeBuildDefaults *CodeBuildOptions `json:"synthCodeBuildDefaults" yaml:"synthCodeBuildDefaults"`
}

// Factory for CodePipeline source steps.
//
// This class contains a number of factory methods for the different types
// of sources that CodePipeline supports.
//
// Example:
//   // Access the CommitId of a GitHub source in the synth
//   source := pipelines.codePipelineSource.gitHub(jsii.String("owner/repo"), jsii.String("main"))
//
//   pipeline := pipelines.NewCodePipeline(*scope, jsii.String("MyPipeline"), &codePipelineProps{
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: source,
//   		commands: []*string{
//   		},
//   		env: map[string]*string{
//   			"COMMIT_ID": source.sourceAttribute(jsii.String("CommitId")),
//   		},
//   	}),
//   })
//
// Experimental.
type CodePipelineSource interface {
	Step
	ICodePipelineActionFactory
	// Return the steps this step depends on, based on the FileSets it requires.
	// Experimental.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	// Experimental.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	// Experimental.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	// Experimental.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	// Experimental.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	// Experimental.
	AddDependencyFileSet(fs FileSet)
	// Add a dependency on another step.
	// Experimental.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	// Experimental.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
	// Experimental.
	DiscoverReferencedOutputs(structure interface{})
	// Experimental.
	GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64, variablesNamespace *string) awscodepipelineactions.Action
	// Create the desired Action and add it to the pipeline.
	// Experimental.
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
	// Return an attribute of the current source revision.
	//
	// These values can be passed into the environment variables of pipeline steps,
	// so your steps can access information about the source revision.
	//
	// Pipeline synth step has some source attributes predefined in the environment.
	// If these suffice, you don't need to use this method for the synth step.
	//
	// Example:
	//   // Access the CommitId of a GitHub source in the synth
	//   source := pipelines.codePipelineSource.gitHub(jsii.String("owner/repo"), jsii.String("main"))
	//
	//   pipeline := pipelines.NewCodePipeline(*scope, jsii.String("MyPipeline"), &codePipelineProps{
	//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
	//   		input: source,
	//   		commands: []*string{
	//   		},
	//   		env: map[string]*string{
	//   			"COMMIT_ID": source.sourceAttribute(jsii.String("CommitId")),
	//   		},
	//   	}),
	//   })
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-variables.html#reference-variables-list
	//
	// Experimental.
	SourceAttribute(name *string) *string
	// Return a string representation of this Step.
	// Experimental.
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
//
// Example:
//   var repository iRepository
//   pipelines.codePipelineSource.codeCommit(repository, jsii.String("main"))
//
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
// pipelines.CodePipelineSource.connection('owner/repo', 'main', {
//    connectionArn: 'arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41', // Created using the AWS console
// });
// ```.
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

// Returns an ECR source.
//
// Example:
//   var repository iRepository
//   pipelines.codePipelineSource.ecr(repository, &eCRSourceOptions{
//   	imageTag: jsii.String("latest"),
//   })
//
// Experimental.
func CodePipelineSource_Ecr(repository awsecr.IRepository, props *ECRSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"ecr",
		[]interface{}{repository, props},
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
// pipelines.CodePipelineSource.gitHub('owner/repo', 'main');
// ```
//
// Authentication will be done by a secret called `github-token` in AWS
// Secrets Manager (unless specified otherwise).
//
// The token should have these permissions:
//
// * **repo** - to read the repository
// * **admin:repo_hook** - if you plan to use webhooks (true by default).
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
//
// Example:
//   var bucket bucket
//   pipelines.codePipelineSource.s3(bucket, jsii.String("path/to/file.zip"))
//
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

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
// Experimental.
func CodePipelineSource_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CodePipelineSource",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipelineSource) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_CodePipelineSource) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (c *jsiiProxy_CodePipelineSource) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_CodePipelineSource) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		c,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

func (c *jsiiProxy_CodePipelineSource) GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64, variablesNamespace *string) awscodepipelineactions.Action {
	var returns awscodepipelineactions.Action

	_jsii_.Invoke(
		c,
		"getAction",
		[]interface{}{output, actionName, runOrder, variablesNamespace},
		&returns,
	)

	return returns
}

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

func (c *jsiiProxy_CodePipelineSource) SourceAttribute(name *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"sourceAttribute",
		[]interface{}{name},
		&returns,
	)

	return returns
}

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
//
// Example:
//   var pipeline codePipeline
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.addStage(stage, &addStageOpts{
//   	pre: []step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &permissionsBroadeningCheckProps{
//   			stage: stage,
//   		}),
//   	},
//   })
//
// Experimental.
type ConfirmPermissionsBroadening interface {
	Step
	ICodePipelineActionFactory
	// Return the steps this step depends on, based on the FileSets it requires.
	// Experimental.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	// Experimental.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	// Experimental.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	// Experimental.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	// Experimental.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	// Experimental.
	AddDependencyFileSet(fs FileSet)
	// Add a dependency on another step.
	// Experimental.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	// Experimental.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
	// Experimental.
	DiscoverReferencedOutputs(structure interface{})
	// Create the desired Action and add it to the pipeline.
	// Experimental.
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
	// Return a string representation of this Step.
	// Experimental.
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

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
// Experimental.
func ConfirmPermissionsBroadening_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"monocdk.pipelines.ConfirmPermissionsBroadening",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		c,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

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
//
// Example:
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//
//   	// Turn this on because the pipeline uses Docker image assets
//   	dockerEnabledForSelfMutation: jsii.Boolean(true),
//   })
//
//   pipeline.addWave(jsii.String("MyWave"), &waveOptions{
//   	post: []step{
//   		pipelines.NewCodeBuildStep(jsii.String("RunApproval"), &codeBuildStepProps{
//   			commands: []*string{
//   				jsii.String("command-from-image"),
//   			},
//   			buildEnvironment: &buildEnvironment{
//   				// The user of a Docker image asset in the pipeline requires turning on
//   				// 'dockerEnabledForSelfMutation'.
//   				buildImage: codebuild.linuxBuildImage.fromAsset(this, jsii.String("Image"), &dockerImageAssetProps{
//   					directory: jsii.String("./docker-image"),
//   				}),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type ConnectionSourceOptions struct {
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// Example:
	//   "arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	// Experimental.
	ConnectionArn *string `json:"connectionArn" yaml:"connectionArn"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	// Experimental.
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	// Experimental.
	TriggerOnPush *bool `json:"triggerOnPush" yaml:"triggerOnPush"`
}

// Action to deploy a CDK Stack.
//
// Adds two CodePipeline Actions to the pipeline: one to create a ChangeSet
// and one to execute it.
//
// You do not need to instantiate this action yourself -- it will automatically
// be added by the pipeline when you add stack artifacts or entire stages.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var role role
//   deployCdkStackAction := pipelines.NewDeployCdkStackAction(&deployCdkStackActionProps{
//   	actionRole: role,
//   	cloudAssemblyInput: artifact,
//   	stackName: jsii.String("stackName"),
//   	templatePath: jsii.String("templatePath"),
//
//   	// the properties below are optional
//   	baseActionName: jsii.String("baseActionName"),
//   	changeSetName: jsii.String("changeSetName"),
//   	cloudFormationExecutionRole: role,
//   	dependencyStackArtifactIds: []*string{
//   		jsii.String("dependencyStackArtifactIds"),
//   	},
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   	region: jsii.String("region"),
//   	stackArtifactId: jsii.String("stackArtifactId"),
//   	templateConfigurationPath: jsii.String("templateConfigurationPath"),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type DeployCdkStackAction interface {
	awscodepipeline.IAction
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// Artifact ids of the artifact this stack artifact depends on.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DependencyStackArtifactIds() *[]*string
	// The runorder for the execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder() *float64
	// The runorder for the prepare action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder() *float64
	// Artifact id of the artifact this action was based on.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackArtifactId() *string
	// Name of the deployed stack.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackName() *string
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewDeployCdkStackAction_Override(d DeployCdkStackAction, props *DeployCdkStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.DeployCdkStackAction",
		[]interface{}{props},
		d,
	)
}

// Construct a DeployCdkStackAction from a Stack artifact.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   deployCdkStackActionOptions := &deployCdkStackActionOptions{
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	baseActionName: jsii.String("baseActionName"),
//   	changeSetName: jsii.String("changeSetName"),
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type DeployCdkStackActionOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Base name of the action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BaseActionName *string `json:"baseActionName" yaml:"baseActionName"`
	// Name of the change set to create and deploy.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ChangeSetName *string `json:"changeSetName" yaml:"changeSetName"`
	// Run order for the Execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `json:"executeRunOrder" yaml:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Output awscodepipeline.Artifact `json:"output" yaml:"output"`
	// Filename in output to write Stack outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OutputFileName *string `json:"outputFileName" yaml:"outputFileName"`
	// Run order for the Prepare action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder *float64 `json:"prepareRunOrder" yaml:"prepareRunOrder"`
}

// Properties for a DeployCdkStackAction.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var role role
//   deployCdkStackActionProps := &deployCdkStackActionProps{
//   	actionRole: role,
//   	cloudAssemblyInput: artifact,
//   	stackName: jsii.String("stackName"),
//   	templatePath: jsii.String("templatePath"),
//
//   	// the properties below are optional
//   	baseActionName: jsii.String("baseActionName"),
//   	changeSetName: jsii.String("changeSetName"),
//   	cloudFormationExecutionRole: role,
//   	dependencyStackArtifactIds: []*string{
//   		jsii.String("dependencyStackArtifactIds"),
//   	},
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   	region: jsii.String("region"),
//   	stackArtifactId: jsii.String("stackArtifactId"),
//   	templateConfigurationPath: jsii.String("templateConfigurationPath"),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type DeployCdkStackActionProps struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Base name of the action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BaseActionName *string `json:"baseActionName" yaml:"baseActionName"`
	// Name of the change set to create and deploy.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ChangeSetName *string `json:"changeSetName" yaml:"changeSetName"`
	// Run order for the Execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `json:"executeRunOrder" yaml:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Output awscodepipeline.Artifact `json:"output" yaml:"output"`
	// Filename in output to write Stack outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OutputFileName *string `json:"outputFileName" yaml:"outputFileName"`
	// Run order for the Prepare action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder *float64 `json:"prepareRunOrder" yaml:"prepareRunOrder"`
	// Role for the action to assume.
	//
	// This controls the account to deploy into.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionRole awsiam.IRole `json:"actionRole" yaml:"actionRole"`
	// The name of the stack that should be created/updated.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackName *string `json:"stackName" yaml:"stackName"`
	// Relative path of template in the input artifact.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TemplatePath *string `json:"templatePath" yaml:"templatePath"`
	// Role to execute CloudFormation under.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudFormationExecutionRole awsiam.IRole `json:"cloudFormationExecutionRole" yaml:"cloudFormationExecutionRole"`
	// Artifact ID for the stacks this stack depends on.
	//
	// Used for pipeline order checking.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DependencyStackArtifactIds *[]*string `json:"dependencyStackArtifactIds" yaml:"dependencyStackArtifactIds"`
	// Region to deploy into.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Region *string `json:"region" yaml:"region"`
	// Artifact ID for the stack deployed here.
	//
	// Used for pipeline order checking.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackArtifactId *string `json:"stackArtifactId" yaml:"stackArtifactId"`
	// Template configuration path relative to the input artifact.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TemplateConfigurationPath *string `json:"templateConfigurationPath" yaml:"templateConfigurationPath"`
}

// Represents credentials used to access a Docker registry.
//
// Example:
//   dockerHubSecret := secretsmanager.secret.fromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
//   customRegSecret := secretsmanager.secret.fromSecretCompleteArn(this, jsii.String("CRSecret"), jsii.String("arn:aws:..."))
//   repo1 := ecr.repository.fromRepositoryArn(this, jsii.String("Repo"), jsii.String("arn:aws:ecr:eu-west-1:0123456789012:repository/Repo1"))
//   repo2 := ecr.repository.fromRepositoryArn(this, jsii.String("Repo"), jsii.String("arn:aws:ecr:eu-west-1:0123456789012:repository/Repo2"))
//
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	dockerCredentials: []dockerCredential{
//   		pipelines.*dockerCredential.dockerHub(dockerHubSecret),
//   		pipelines.*dockerCredential.customRegistry(jsii.String("dockerregistry.example.com"), customRegSecret),
//   		pipelines.*dockerCredential.ecr([]iRepository{
//   			repo1,
//   			repo2,
//   		}),
//   	},
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   })
//
// Experimental.
type DockerCredential interface {
	// Experimental.
	Usages() *[]DockerCredentialUsage
	// Grant read-only access to the registry credentials.
	//
	// This grants read access to any secrets, and pull access to any repositories.
	// Experimental.
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
// Convenience method for `customRegistry('https://index.docker.io/v1/', opts)`.
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

func (d *jsiiProxy_DockerCredential) GrantRead(grantee awsiam.IGrantable, usage DockerCredentialUsage) {
	_jsii_.InvokeVoid(
		d,
		"grantRead",
		[]interface{}{grantee, usage},
	)
}

// Defines which stages of a pipeline require the specified credentials.
//
// Example:
//   dockerHubSecret := secretsmanager.secret.fromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
//   // Only the image asset publishing actions will be granted read access to the secret.
//   creds := pipelines.dockerCredential.dockerHub(dockerHubSecret, &externalDockerCredentialOptions{
//   	usages: []dockerCredentialUsage{
//   		pipelines.*dockerCredentialUsage_ASSET_PUBLISHING,
//   	},
//   })
//
// Experimental.
type DockerCredentialUsage string

const (
	// Synth/Build.
	// Experimental.
	DockerCredentialUsage_SYNTH DockerCredentialUsage = "SYNTH"
	// Self-update.
	// Experimental.
	DockerCredentialUsage_SELF_UPDATE DockerCredentialUsage = "SELF_UPDATE"
	// Asset publishing.
	// Experimental.
	DockerCredentialUsage_ASSET_PUBLISHING DockerCredentialUsage = "ASSET_PUBLISHING"
)

// Options for ECR sources.
//
// Example:
//   var repository iRepository
//   pipelines.codePipelineSource.ecr(repository, &eCRSourceOptions{
//   	imageTag: jsii.String("latest"),
//   })
//
// Experimental.
type ECRSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	// Experimental.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// The image tag that will be checked for changes.
	// Experimental.
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
}

// Options for defining access for a Docker Credential composed of ECR repos.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var role role
//   ecrDockerCredentialOptions := &ecrDockerCredentialOptions{
//   	assumeRole: role,
//   	usages: []dockerCredentialUsage{
//   		pipelines.*dockerCredentialUsage_SYNTH,
//   	},
//   }
//
// Experimental.
type EcrDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	// Experimental.
	AssumeRole awsiam.IRole `json:"assumeRole" yaml:"assumeRole"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	// Experimental.
	Usages *[]DockerCredentialUsage `json:"usages" yaml:"usages"`
}

// Options for defining credentials for a Docker Credential.
//
// Example:
//   dockerHubSecret := secretsmanager.secret.fromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
//   // Only the image asset publishing actions will be granted read access to the secret.
//   creds := pipelines.dockerCredential.dockerHub(dockerHubSecret, &externalDockerCredentialOptions{
//   	usages: []dockerCredentialUsage{
//   		pipelines.*dockerCredentialUsage_ASSET_PUBLISHING,
//   	},
//   })
//
// Experimental.
type ExternalDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	// Experimental.
	AssumeRole awsiam.IRole `json:"assumeRole" yaml:"assumeRole"`
	// The name of the JSON field of the secret which contains the secret/password.
	// Experimental.
	SecretPasswordField *string `json:"secretPasswordField" yaml:"secretPasswordField"`
	// The name of the JSON field of the secret which contains the user/login name.
	// Experimental.
	SecretUsernameField *string `json:"secretUsernameField" yaml:"secretUsernameField"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	// Experimental.
	Usages *[]DockerCredentialUsage `json:"usages" yaml:"usages"`
}

// A set of files traveling through the deployment pipeline.
//
// Individual steps in the pipeline produce or consume
// `FileSet`s.
//
// Example:
//   type myJenkinsStep struct {
//   	step
//   }
//
//   func newMyJenkinsStep(provider jenkinsProvider, input fileSet) *myJenkinsStep {
//   	this := &myJenkinsStep{}
//   	pipelines.NewStep_Override(this, jsii.String("MyJenkinsStep"))
//
//   	// This is necessary if your step accepts parametres, like environment variables,
//   	// that may contain outputs from other steps. It doesn't matter what the
//   	// structure is, as long as it contains the values that may contain outputs.
//   	this.discoverReferencedOutputs(map[string]map[string]interface{}{
//   		"env": map[string]interface{}{
//   		},
//   	})
//   	return this
//   }
//
//   func (this *myJenkinsStep) produceAction(stage iStage, options produceActionOptions) codePipelineActionFactoryResult {
//
//   	// This is where you control what type of Action gets added to the
//   	// CodePipeline
//   	*stage.addAction(cpactions.NewJenkinsAction(&jenkinsActionProps{
//   		// Copy 'actionName' and 'runOrder' from the options
//   		actionName: options.actionName,
//   		runOrder: options.runOrder,
//
//   		// Jenkins-specific configuration
//   		type: cpactions.jenkinsActionType_TEST,
//   		jenkinsProvider: this.provider,
//   		projectName: jsii.String("MyJenkinsProject"),
//
//   		// Translate the FileSet into a codepipeline.Artifact
//   		inputs: []artifact{
//   			options.artifacts.toCodePipeline(this.input),
//   		},
//   	}))
//
//   	return &codePipelineActionFactoryResult{
//   		runOrdersConsumed: jsii.Number(1),
//   	}
//   }
//
// Experimental.
type FileSet interface {
	IFileSetProducer
	// Human-readable descriptor for this file set (does not need to be unique).
	// Experimental.
	Id() *string
	// The primary output of a file set producer.
	//
	// The primary output of a FileSet is itself.
	// Experimental.
	PrimaryOutput() FileSet
	// The Step that produces this FileSet.
	// Experimental.
	Producer() Step
	// Mark the given Step as the producer for this FileSet.
	//
	// This method can only be called once.
	// Experimental.
	ProducedBy(producer Step)
	// Return a string representation of this FileSet.
	// Experimental.
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

func (f *jsiiProxy_FileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		f,
		"producedBy",
		[]interface{}{producer},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var fileSet fileSet
//   fileSetLocation := &fileSetLocation{
//   	directory: jsii.String("directory"),
//   	fileSet: fileSet,
//   }
//
// Experimental.
type FileSetLocation struct {
	// The (relative) directory where the FileSet is found.
	// Experimental.
	Directory *string `json:"directory" yaml:"directory"`
	// The FileSet object.
	// Experimental.
	FileSet FileSet `json:"fileSet" yaml:"fileSet"`
}

// Options for CdkDeployAction.fromStackArtifact.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   fromStackArtifactOptions := &fromStackArtifactOptions{
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	executeRunOrder: jsii.Number(123),
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	prepareRunOrder: jsii.Number(123),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type FromStackArtifactOptions struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Run order for the Execute action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ExecuteRunOrder *float64 `json:"executeRunOrder" yaml:"executeRunOrder"`
	// Artifact to write Stack Outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Output awscodepipeline.Artifact `json:"output" yaml:"output"`
	// Filename in output to write Stack outputs to.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OutputFileName *string `json:"outputFileName" yaml:"outputFileName"`
	// Run order for the 2 actions that will be created.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PrepareRunOrder *float64 `json:"prepareRunOrder" yaml:"prepareRunOrder"`
}

// Options for GitHub sources.
//
// Example:
//   pipelines.codePipelineSource.gitHub(jsii.String("org/repo"), jsii.String("branch"), &gitHubSourceOptions{
//   	// This is optional
//   	authentication: cdk.secretValue.secretsManager(jsii.String("my-token")),
//   })
//
// Experimental.
type GitHubSourceOptions struct {
	// A GitHub OAuth token to use for authentication.
	//
	// It is recommended to use a Secrets Manager `Secret` to obtain the token:
	//
	// ```ts
	// const oauth = cdk.SecretValue.secretsManager('my-github-token');
	// ```
	//
	// The GitHub Personal Access Token should have these scopes:
	//
	// * **repo** - to read the repository
	// * **admin:repo_hook** - if you plan to use webhooks (true by default).
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/GitHub-create-personal-token-CLI.html
	//
	// Experimental.
	Authentication awscdk.SecretValue `json:"authentication" yaml:"authentication"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action.
	// With "POLL", CodePipeline periodically checks the source for changes.
	// With "None", the action is not triggered through changes in the source.
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	// Experimental.
	Trigger awscodepipelineactions.GitHubTrigger `json:"trigger" yaml:"trigger"`
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
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type IStageHost interface {
	// Make sure all the assets from the given manifest are published.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PublishAsset(command *AssetPublishingCommand)
	// Return the Artifact the given stack has to emit its outputs into, if any.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
//
// Example:
//   var pipeline codePipeline
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//
//   pipeline.addStage(preprod, &addStageOpts{
//   	post: []step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &shellStepProps{
//   			commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.addStage(prod, &addStageOpts{
//   	pre: []*step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd")),
//   	},
//   })
//
// Experimental.
type ManualApprovalStep interface {
	Step
	// The comment associated with this manual approval.
	// Experimental.
	Comment() *string
	// Return the steps this step depends on, based on the FileSets it requires.
	// Experimental.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	// Experimental.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	// Experimental.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	// Experimental.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	// Experimental.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	// Experimental.
	AddDependencyFileSet(fs FileSet)
	// Add a dependency on another step.
	// Experimental.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	// Experimental.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
	// Experimental.
	DiscoverReferencedOutputs(structure interface{})
	// Return a string representation of this Step.
	// Experimental.
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

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
// Experimental.
func ManualApprovalStep_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"monocdk.pipelines.ManualApprovalStep",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManualApprovalStep) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		m,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (m *jsiiProxy_ManualApprovalStep) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		m,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (m *jsiiProxy_ManualApprovalStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		m,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (m *jsiiProxy_ManualApprovalStep) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		m,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   manualApprovalStepProps := &manualApprovalStepProps{
//   	comment: jsii.String("comment"),
//   }
//
// Experimental.
type ManualApprovalStepProps struct {
	// The comment to display with this manual approval.
	// Experimental.
	Comment *string `json:"comment" yaml:"comment"`
}

// Properties for a `PermissionsBroadeningCheck`.
//
// Example:
//   var pipeline codePipeline
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.addStage(stage, &addStageOpts{
//   	pre: []step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &permissionsBroadeningCheckProps{
//   			stage: stage,
//   		}),
//   	},
//   })
//
// Experimental.
type PermissionsBroadeningCheckProps struct {
	// The CDK Stage object to check the stacks of.
	//
	// This should be the same Stage object you are passing to `addStage()`.
	// Experimental.
	Stage awscdk.Stage `json:"stage" yaml:"stage"`
	// Topic to send notifications when a human needs to give manual confirmation.
	// Experimental.
	NotificationTopic awssns.ITopic `json:"notificationTopic" yaml:"notificationTopic"`
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
	// The FileSet tha contains the cloud assembly.
	//
	// This is the primary output of the synth step.
	// Experimental.
	CloudAssemblyFileSet() FileSet
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The build step that produces the CDK Cloud Assembly.
	// Experimental.
	Synth() IFileSetProducer
	// The waves in this pipeline.
	// Experimental.
	Waves() *[]Wave
	// Deploy a single Stage by itself.
	//
	// Add a Stage to the pipeline, to be deployed in sequence with other
	// Stages added to the pipeline. All Stacks in the stage will be deployed
	// in an order automatically determined by their relative dependencies.
	// Experimental.
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
	// Add a Wave to the pipeline, for deploying multiple Stages in parallel.
	//
	// Use the return object of this method to deploy multiple stages in parallel.
	//
	// Example:
	//
	// ```ts
	// declare const pipeline: pipelines.CodePipeline;
	//
	// const wave = pipeline.addWave('MyWave');
	// wave.addStage(new MyApplicationStage(this, 'Stage1'));
	// wave.addStage(new MyApplicationStage(this, 'Stage2'));
	// ```.
	// Experimental.
	AddWave(id *string, options *WaveOptions) Wave
	// Send the current pipeline definition to the engine, and construct the pipeline.
	//
	// It is not possible to modify the pipeline after calling this method.
	// Experimental.
	BuildPipeline()
	// Implemented by subclasses to do the actual pipeline construction.
	// Experimental.
	DoBuildPipeline()
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
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

func (p *jsiiProxy_PipelineBase) BuildPipeline() {
	_jsii_.InvokeVoid(
		p,
		"buildPipeline",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineBase) DoBuildPipeline() {
	_jsii_.InvokeVoid(
		p,
		"doBuildPipeline",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineBase) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineBase) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (p *jsiiProxy_PipelineBase) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineBase) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var fileSetProducer iFileSetProducer
//   pipelineBaseProps := &pipelineBaseProps{
//   	synth: fileSetProducer,
//   }
//
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
	Synth IFileSetProducer `json:"synth" yaml:"synth"`
}

// Options for the `CodePipelineActionFactory.produce()` method.
//
// Example:
//   import constructs "github.com/aws/constructs-go/constructs"import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3 "github.com/aws/aws-cdk-go/awscdk/aws_s3"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var artifactMap artifactMap
//   var bucket bucket
//   var buildImage iBuildImage
//   var buildSpec buildSpec
//   var codePipeline codePipeline
//   var construct construct
//   var duration duration
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//   produceActionOptions := &produceActionOptions{
//   	actionName: jsii.String("actionName"),
//   	artifacts: artifactMap,
//   	pipeline: codePipeline,
//   	runOrder: jsii.Number(123),
//   	scope: construct,
//
//   	// the properties below are optional
//   	beforeSelfMutation: jsii.Boolean(false),
//   	codeBuildDefaults: &codeBuildOptions{
//   		buildEnvironment: &buildEnvironment{
//   			buildImage: buildImage,
//   			certificate: &buildEnvironmentCertificate{
//   				bucket: bucket,
//   				objectKey: jsii.String("objectKey"),
//   			},
//   			computeType: codebuild.computeType_SMALL,
//   			environmentVariables: map[string]buildEnvironmentVariable{
//   				"environmentVariablesKey": &buildEnvironmentVariable{
//   					"value": value,
//
//   					// the properties below are optional
//   					"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   				},
//   			},
//   			privileged: jsii.Boolean(false),
//   		},
//   		partialBuildSpec: buildSpec,
//   		rolePolicy: []*policyStatement{
//   			policyStatement,
//   		},
//   		securityGroups: []iSecurityGroup{
//   			securityGroup,
//   		},
//   		subnetSelection: &subnetSelection{
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			onePerAz: jsii.Boolean(false),
//   			subnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			subnetGroupName: jsii.String("subnetGroupName"),
//   			subnetName: jsii.String("subnetName"),
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: ec2.subnetType_ISOLATED,
//   		},
//   		timeout: duration,
//   		vpc: vpc,
//   	},
//   	fallbackArtifact: artifact,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   }
//
// Experimental.
type ProduceActionOptions struct {
	// Name the action should get.
	// Experimental.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Helper object to translate FileSets to CodePipeline Artifacts.
	// Experimental.
	Artifacts ArtifactMap `json:"artifacts" yaml:"artifacts"`
	// The pipeline the action is being generated for.
	// Experimental.
	Pipeline CodePipeline `json:"pipeline" yaml:"pipeline"`
	// RunOrder the action should get.
	// Experimental.
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
	// Scope in which to create constructs.
	// Experimental.
	Scope constructs.Construct `json:"scope" yaml:"scope"`
	// Whether or not this action is inserted before self mutation.
	//
	// If it is, the action should take care to reflect some part of
	// its own definition in the pipeline action definition, to
	// trigger a restart after self-mutation (if necessary).
	// Experimental.
	BeforeSelfMutation *bool `json:"beforeSelfMutation" yaml:"beforeSelfMutation"`
	// If this action factory creates a CodeBuild step, default options to inherit.
	// Experimental.
	CodeBuildDefaults *CodeBuildOptions `json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An input artifact that CodeBuild projects that don't actually need an input artifact can use.
	//
	// CodeBuild Projects MUST have an input artifact in order to be added to the Pipeline. If
	// the Project doesn't actually care about its input (it can be anything), it can use the
	// Artifact passed here.
	// Experimental.
	FallbackArtifact awscodepipeline.Artifact `json:"fallbackArtifact" yaml:"fallbackArtifact"`
	// If this step is producing outputs, the variables namespace assigned to it.
	//
	// Pass this on to the Action you are creating.
	// Experimental.
	VariablesNamespace *string `json:"variablesNamespace" yaml:"variablesNamespace"`
}

// Action to publish an asset in the pipeline.
//
// Creates a CodeBuild project which will use the CDK CLI
// to prepare and publish the asset.
//
// You do not need to instantiate this action -- it will automatically
// be added by the pipeline when you add stacks that use assets.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dependable iDependable
//   var role role
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//   publishAssetsAction := pipelines.NewPublishAssetsAction(this, jsii.String("MyPublishAssetsAction"), &publishAssetsActionProps{
//   	actionName: jsii.String("actionName"),
//   	assetType: pipelines.assetType_FILE,
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	buildSpec: buildSpec,
//   	cdkCliVersion: jsii.String("cdkCliVersion"),
//   	createBuildspecFile: jsii.Boolean(false),
//   	dependable: dependable,
//   	preInstallCommands: []*string{
//   		jsii.String("preInstallCommands"),
//   	},
//   	projectName: jsii.String("projectName"),
//   	role: role,
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: ec2.subnetType_ISOLATED,
//   	},
//   	vpc: vpc,
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type PublishAssetsAction interface {
	awscdk.Construct
	awscodepipeline.IAction
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
	// Add a single publishing command.
	//
	// Manifest path should be relative to the root Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddPublishCommand(relativeManifestPath *string, assetSelector *string)
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnPrepare()
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewPublishAssetsAction_Override(p PublishAssetsAction, scope constructs.Construct, id *string, props *PublishAssetsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.PublishAssetsAction",
		[]interface{}{scope, id, props},
		p,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

func (p *jsiiProxy_PublishAssetsAction) AddPublishCommand(relativeManifestPath *string, assetSelector *string) {
	_jsii_.InvokeVoid(
		p,
		"addPublishCommand",
		[]interface{}{relativeManifestPath, assetSelector},
	)
}

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

func (p *jsiiProxy_PublishAssetsAction) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

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

func (p *jsiiProxy_PublishAssetsAction) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (p *jsiiProxy_PublishAssetsAction) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PublishAssetsAction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dependable iDependable
//   var role role
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//   publishAssetsActionProps := &publishAssetsActionProps{
//   	actionName: jsii.String("actionName"),
//   	assetType: pipelines.assetType_FILE,
//   	cloudAssemblyInput: artifact,
//
//   	// the properties below are optional
//   	buildSpec: buildSpec,
//   	cdkCliVersion: jsii.String("cdkCliVersion"),
//   	createBuildspecFile: jsii.Boolean(false),
//   	dependable: dependable,
//   	preInstallCommands: []*string{
//   		jsii.String("preInstallCommands"),
//   	},
//   	projectName: jsii.String("projectName"),
//   	role: role,
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: ec2.subnetType_ISOLATED,
//   	},
//   	vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type PublishAssetsActionProps struct {
	// Name of publishing action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `json:"actionName" yaml:"actionName"`
	// AssetType we're publishing.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AssetType AssetType `json:"assetType" yaml:"assetType"`
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Custom BuildSpec that is merged with generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Version of CDK CLI to 'npm install'.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CdkCliVersion *string `json:"cdkCliVersion" yaml:"cdkCliVersion"`
	// Use a file buildspec written to the cloud assembly instead of an inline buildspec.
	//
	// This prevents size limitation errors as inline specs have a max length of 25600 characters.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CreateBuildspecFile *bool `json:"createBuildspecFile" yaml:"createBuildspecFile"`
	// Any Dependable construct that the CodeBuild project needs to take a dependency on.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Dependable awscdk.IDependable `json:"dependable" yaml:"dependable"`
	// Additional commands to run before installing cdk-assert Use this to setup proxies or npm mirrors.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PreInstallCommands *[]*string `json:"preInstallCommands" yaml:"preInstallCommands"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Role to use for CodePipeline and CodeBuild to build and publish the assets.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the PublishAssetsAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Options for S3 sources.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline_actions "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline_actions"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   s3SourceOptions := &s3SourceOptions{
//   	actionName: jsii.String("actionName"),
//   	trigger: codepipeline_actions.s3Trigger_NONE,
//   }
//
// Experimental.
type S3SourceOptions struct {
	// The action name used for this source in the CodePipeline.
	// Experimental.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	// Experimental.
	Trigger awscodepipelineactions.S3Trigger `json:"trigger" yaml:"trigger"`
}

// Validate a revision using shell commands.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3 "github.com/aws/aws-cdk-go/awscdk/aws_s3"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var bucket bucket
//   var buildImage iBuildImage
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var stackOutput stackOutput
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//   shellScriptAction := pipelines.NewShellScriptAction(&shellScriptActionProps{
//   	actionName: jsii.String("actionName"),
//   	commands: []*string{
//   		jsii.String("commands"),
//   	},
//
//   	// the properties below are optional
//   	additionalArtifacts: []*artifact{
//   		artifact,
//   	},
//   	bashOptions: jsii.String("bashOptions"),
//   	environment: &buildEnvironment{
//   		buildImage: buildImage,
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: bucket,
//   			objectKey: jsii.String("objectKey"),
//   		},
//   		computeType: codebuild.computeType_SMALL,
//   		environmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		privileged: jsii.Boolean(false),
//   	},
//   	environmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	rolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	runOrder: jsii.Number(123),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: ec2.subnetType_ISOLATED,
//   	},
//   	useOutputs: map[string]*stackOutput{
//   		"useOutputsKey": stackOutput,
//   	},
//   	vpc: vpc,
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type ShellScriptAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The CodeBuild Project's principal.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	GrantPrincipal() awsiam.IPrincipal
	// Project generated to run the shell script in.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Project() awscodebuild.IProject
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewShellScriptAction_Override(s ShellScriptAction, props *ShellScriptActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.ShellScriptAction",
		[]interface{}{props},
		s,
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3 "github.com/aws/aws-cdk-go/awscdk/aws_s3"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var bucket bucket
//   var buildImage iBuildImage
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var stackOutput stackOutput
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//   shellScriptActionProps := &shellScriptActionProps{
//   	actionName: jsii.String("actionName"),
//   	commands: []*string{
//   		jsii.String("commands"),
//   	},
//
//   	// the properties below are optional
//   	additionalArtifacts: []*artifact{
//   		artifact,
//   	},
//   	bashOptions: jsii.String("bashOptions"),
//   	environment: &buildEnvironment{
//   		buildImage: buildImage,
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: bucket,
//   			objectKey: jsii.String("objectKey"),
//   		},
//   		computeType: codebuild.computeType_SMALL,
//   		environmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		privileged: jsii.Boolean(false),
//   	},
//   	environmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	rolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	runOrder: jsii.Number(123),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: ec2.subnetType_ISOLATED,
//   	},
//   	useOutputs: map[string]*stackOutput{
//   		"useOutputsKey": stackOutput,
//   	},
//   	vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type ShellScriptActionProps struct {
	// Name of the validation action in the pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Commands to run.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Commands *[]*string `json:"commands" yaml:"commands"`
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
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]awscodepipeline.Artifact `json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// Bash options to set at the start of the script.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BashOptions *string `json:"bashOptions" yaml:"bashOptions"`
	// The CodeBuild environment where scripts are executed.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// Additional policy statements to add to the execution role.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// RunOrder for this action.
	//
	// Use this to sequence the shell script after the deployments.
	//
	// The default value is 100 so you don't have to supply the value if you just
	// want to run this after the application stacks have been deployed, and you
	// don't have more than 100 stacks.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// Stack outputs to make available as environment variables.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	UseOutputs *map[string]StackOutput `json:"useOutputs" yaml:"useOutputs"`
	// The VPC where to execute the specified script.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Run shell script commands in the pipeline.
//
// This is a generic step designed
// to be deployment engine agnostic.
//
// Example:
//   // Modern API
//   modernPipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	selfMutation: jsii.Boolean(false),
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   })
//
//   // Original API
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   originalPipeline := pipelines.NewCdkPipeline(this, jsii.String("Pipeline"), &cdkPipelineProps{
//   	selfMutating: jsii.Boolean(false),
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   })
//
// Experimental.
type ShellStep interface {
	Step
	// Commands to run.
	// Experimental.
	Commands() *[]*string
	// Return the steps this step depends on, based on the FileSets it requires.
	// Experimental.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	// Experimental.
	DependencyFileSets() *[]FileSet
	// Environment variables to set.
	// Experimental.
	Env() *map[string]*string
	// Set environment variables based on Stack Outputs.
	// Experimental.
	EnvFromCfnOutputs() *map[string]StackOutputReference
	// Identifier for this step.
	// Experimental.
	Id() *string
	// Input FileSets.
	//
	// A list of `(FileSet, directory)` pairs, which are a copy of the
	// input properties. This list should not be modified directly.
	// Experimental.
	Inputs() *[]*FileSetLocation
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Experimental.
	InstallCommands() *[]*string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	// Experimental.
	IsSource() *bool
	// Output FileSets.
	//
	// A list of `(FileSet, directory)` pairs, which are a copy of the
	// input properties. This list should not be modified directly.
	// Experimental.
	Outputs() *[]*FileSetLocation
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	// Experimental.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	// Experimental.
	AddDependencyFileSet(fs FileSet)
	// Add an additional output FileSet based on a directory.
	//
	// After running the script, the contents of the given directory
	// will be exported as a `FileSet`. Use the `FileSet` as the
	// input to another step.
	//
	// Multiple calls with the exact same directory name string (not normalized)
	// will return the same FileSet.
	// Experimental.
	AddOutputDirectory(directory *string) FileSet
	// Add a dependency on another step.
	// Experimental.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	// Experimental.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
	// Experimental.
	DiscoverReferencedOutputs(structure interface{})
	// Configure the given output directory as primary output.
	//
	// If no primary output has been configured yet, this directory
	// will become the primary output of this ShellStep, otherwise this
	// method will throw if the given directory is different than the
	// currently configured primary output directory.
	// Experimental.
	PrimaryOutputDirectory(directory *string) FileSet
	// Return a string representation of this Step.
	// Experimental.
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

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
// Experimental.
func ShellStep_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"monocdk.pipelines.ShellStep",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ShellStep) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

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

func (s *jsiiProxy_ShellStep) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		s,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (s *jsiiProxy_ShellStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (s *jsiiProxy_ShellStep) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		s,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

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
//
// Example:
//   // Modern API
//   modernPipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	selfMutation: jsii.Boolean(false),
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   })
//
//   // Original API
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   originalPipeline := pipelines.NewCdkPipeline(this, jsii.String("Pipeline"), &cdkPipelineProps{
//   	selfMutating: jsii.Boolean(false),
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   })
//
// Experimental.
type ShellStepProps struct {
	// Commands to run.
	// Experimental.
	Commands *[]*string `json:"commands" yaml:"commands"`
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
	// const script = new pipelines.ShellStep('MainScript', {
	//    commands: ['npm ci','npm run build','npx cdk synth'],
	//    input: pipelines.CodePipelineSource.gitHub('org/source1', 'main'),
	//    additionalInputs: {
	//      '../siblingdir': pipelines.CodePipelineSource.gitHub('org/source2', 'main'),
	//    }
	// });
	// ```.
	// Experimental.
	AdditionalInputs *map[string]IFileSetProducer `json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	// Experimental.
	Env *map[string]*string `json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	// Experimental.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	// Experimental.
	Input IFileSetProducer `json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	// Experimental.
	InstallCommands *[]*string `json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	// Experimental.
	PrimaryOutputDirectory *string `json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
}

// A standard synth with a generated buildspec.
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &cdkPipelineProps{
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   	synthAction: pipelines.simpleSynthAction.standardNpmSynth(&standardNpmSynthOptions{
//   		sourceArtifact: sourceArtifact,
//   		cloudAssemblyArtifact: cloudAssemblyArtifact,
//   		environment: &buildEnvironment{
//   			privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type SimpleSynthAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The CodeBuild Project's principal.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	GrantPrincipal() awsiam.IPrincipal
	// Project generated to run the synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Project() awscodebuild.IProject
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3 "github.com/aws/aws-cdk-go/awscdk/aws_s3"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var bucket bucket
//   var buildImage iBuildImage
//   var buildSpec buildSpec
//   var policyStatement policyStatement
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//   simpleSynthActionProps := &simpleSynthActionProps{
//   	cloudAssemblyArtifact: artifact,
//   	sourceArtifact: artifact,
//   	synthCommand: jsii.String("synthCommand"),
//
//   	// the properties below are optional
//   	actionName: jsii.String("actionName"),
//   	additionalArtifacts: []additionalArtifact{
//   		&additionalArtifact{
//   			artifact: artifact,
//   			directory: jsii.String("directory"),
//   		},
//   	},
//   	buildCommand: jsii.String("buildCommand"),
//   	buildCommands: []*string{
//   		jsii.String("buildCommands"),
//   	},
//   	buildSpec: buildSpec,
//   	copyEnvironmentVariables: []*string{
//   		jsii.String("copyEnvironmentVariables"),
//   	},
//   	environment: &buildEnvironment{
//   		buildImage: buildImage,
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: bucket,
//   			objectKey: jsii.String("objectKey"),
//   		},
//   		computeType: codebuild.computeType_SMALL,
//   		environmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		privileged: jsii.Boolean(false),
//   	},
//   	environmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	installCommand: jsii.String("installCommand"),
//   	installCommands: []*string{
//   		jsii.String("installCommands"),
//   	},
//   	projectName: jsii.String("projectName"),
//   	rolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	subdirectory: jsii.String("subdirectory"),
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: ec2.subnetType_ISOLATED,
//   	},
//   	testCommands: []*string{
//   		jsii.String("testCommands"),
//   	},
//   	vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type SimpleSynthActionProps struct {
	// The artifact where the CloudAssembly should be emitted.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact" yaml:"sourceArtifact"`
	// Name of the build action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables" yaml:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SynthCommand *string `json:"synthCommand" yaml:"synthCommand"`
	// The build command.
	//
	// If your programming language requires a compilation step, put the
	// compilation command here.
	// Deprecated: Use `buildCommands` instead.
	BuildCommand *string `json:"buildCommand" yaml:"buildCommand"`
	// The build commands.
	//
	// If your programming language requires a compilation step, put the
	// compilation command here.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildCommands *[]*string `json:"buildCommands" yaml:"buildCommands"`
	// The install command.
	//
	// If not provided by the build image or another dependency
	// management tool, at least install the CDK CLI here using
	// `npm install -g aws-cdk`.
	// Deprecated: Use `installCommands` instead.
	InstallCommand *string `json:"installCommand" yaml:"installCommand"`
	// Install commands.
	//
	// If not provided by the build image or another dependency
	// management tool, at least install the CDK CLI here using
	// `npm install -g aws-cdk`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	InstallCommands *[]*string `json:"installCommands" yaml:"installCommands"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TestCommands *[]*string `json:"testCommands" yaml:"testCommands"`
}

// Configuration options for a SimpleSynth.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3 "github.com/aws/aws-cdk-go/awscdk/aws_s3"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var bucket bucket
//   var buildImage iBuildImage
//   var buildSpec buildSpec
//   var policyStatement policyStatement
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//   simpleSynthOptions := &simpleSynthOptions{
//   	cloudAssemblyArtifact: artifact,
//   	sourceArtifact: artifact,
//
//   	// the properties below are optional
//   	actionName: jsii.String("actionName"),
//   	additionalArtifacts: []additionalArtifact{
//   		&additionalArtifact{
//   			artifact: artifact,
//   			directory: jsii.String("directory"),
//   		},
//   	},
//   	buildSpec: buildSpec,
//   	copyEnvironmentVariables: []*string{
//   		jsii.String("copyEnvironmentVariables"),
//   	},
//   	environment: &buildEnvironment{
//   		buildImage: buildImage,
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: bucket,
//   			objectKey: jsii.String("objectKey"),
//   		},
//   		computeType: codebuild.computeType_SMALL,
//   		environmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		privileged: jsii.Boolean(false),
//   	},
//   	environmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	projectName: jsii.String("projectName"),
//   	rolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	subdirectory: jsii.String("subdirectory"),
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: ec2.subnetType_ISOLATED,
//   	},
//   	vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type SimpleSynthOptions struct {
	// The artifact where the CloudAssembly should be emitted.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact" yaml:"sourceArtifact"`
	// Name of the build action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables" yaml:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// An asset used by a Stack.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   stackAsset := &stackAsset{
//   	assetId: jsii.String("assetId"),
//   	assetManifestPath: jsii.String("assetManifestPath"),
//   	assetSelector: jsii.String("assetSelector"),
//   	assetType: pipelines.assetType_FILE,
//   	isTemplate: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   }
//
// Experimental.
type StackAsset struct {
	// Asset identifier.
	// Experimental.
	AssetId *string `json:"assetId" yaml:"assetId"`
	// Absolute asset manifest path.
	//
	// This needs to be made relative at a later point in time, but when this
	// information is parsed we don't know about the root cloud assembly yet.
	// Experimental.
	AssetManifestPath *string `json:"assetManifestPath" yaml:"assetManifestPath"`
	// Asset selector to pass to `cdk-assets`.
	// Experimental.
	AssetSelector *string `json:"assetSelector" yaml:"assetSelector"`
	// Type of asset to publish.
	// Experimental.
	AssetType AssetType `json:"assetType" yaml:"assetType"`
	// Does this asset represent the CloudFormation template for the stack.
	// Experimental.
	IsTemplate *bool `json:"isTemplate" yaml:"isTemplate"`
	// Role ARN to assume to publish.
	// Experimental.
	AssetPublishingRoleArn *string `json:"assetPublishingRoleArn" yaml:"assetPublishingRoleArn"`
}

// Deployment of a single Stack.
//
// You don't need to instantiate this class -- it will
// be automatically instantiated as necessary when you
// add a `Stage` to a pipeline.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var cloudFormationStackArtifact cloudFormationStackArtifact
//   stackDeployment := pipelines.stackDeployment.fromArtifact(cloudFormationStackArtifact)
//
// Experimental.
type StackDeployment interface {
	// Template path on disk to CloudAssembly.
	// Experimental.
	AbsoluteTemplatePath() *string
	// Account where the stack should be deployed.
	// Experimental.
	Account() *string
	// Assets referenced by this stack.
	// Experimental.
	Assets() *[]*StackAsset
	// Role to assume before deploying this stack.
	// Experimental.
	AssumeRoleArn() *string
	// Steps that take place after stack is prepared but before stack deploys.
	//
	// Your pipeline engine may not disable `prepareStep`.
	// Experimental.
	ChangeSet() *[]Step
	// Construct path for this stack.
	// Experimental.
	ConstructPath() *string
	// Execution role to pass to CloudFormation.
	// Experimental.
	ExecutionRoleArn() *string
	// Steps to execute after stack deploys.
	// Experimental.
	Post() *[]Step
	// Steps that take place before stack is prepared.
	//
	// If your pipeline engine disables 'prepareStep', then this will happen before stack deploys.
	// Experimental.
	Pre() *[]Step
	// Region where the stack should be deployed.
	// Experimental.
	Region() *string
	// Artifact ID for this stack.
	// Experimental.
	StackArtifactId() *string
	// Other stacks this stack depends on.
	// Experimental.
	StackDependencies() *[]StackDeployment
	// Name for this stack.
	// Experimental.
	StackName() *string
	// Tags to apply to the stack.
	// Experimental.
	Tags() *map[string]*string
	// The asset that represents the CloudFormation template for this stack.
	// Experimental.
	TemplateAsset() *StackAsset
	// The S3 URL which points to the template asset location in the publishing bucket.
	//
	// This is `undefined` if the stack template is not published. Use the
	// `DefaultStackSynthesizer` to ensure it is.
	//
	// Example value: `https://bucket.s3.amazonaws.com/object/key`
	// Experimental.
	TemplateUrl() *string
	// Add a dependency on another stack.
	// Experimental.
	AddStackDependency(stackDeployment StackDeployment)
	// Adds steps to each phase of the stack.
	// Experimental.
	AddStackSteps(pre *[]Step, changeSet *[]Step, post *[]Step)
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

func (j *jsiiProxy_StackDeployment) ChangeSet() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"changeSet",
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

func (j *jsiiProxy_StackDeployment) Post() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"post",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackDeployment) Pre() *[]Step {
	var returns *[]Step
	_jsii_.Get(
		j,
		"pre",
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

func (s *jsiiProxy_StackDeployment) AddStackDependency(stackDeployment StackDeployment) {
	_jsii_.InvokeVoid(
		s,
		"addStackDependency",
		[]interface{}{stackDeployment},
	)
}

func (s *jsiiProxy_StackDeployment) AddStackSteps(pre *[]Step, changeSet *[]Step, post *[]Step) {
	_jsii_.InvokeVoid(
		s,
		"addStackSteps",
		[]interface{}{pre, changeSet, post},
	)
}

// Properties for a `StackDeployment`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//   stackDeploymentProps := &stackDeploymentProps{
//   	absoluteTemplatePath: jsii.String("absoluteTemplatePath"),
//   	constructPath: jsii.String("constructPath"),
//   	stackArtifactId: jsii.String("stackArtifactId"),
//   	stackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	assets: []stackAsset{
//   		&stackAsset{
//   			assetId: jsii.String("assetId"),
//   			assetManifestPath: jsii.String("assetManifestPath"),
//   			assetSelector: jsii.String("assetSelector"),
//   			assetType: pipelines.assetType_FILE,
//   			isTemplate: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   		},
//   	},
//   	assumeRoleArn: jsii.String("assumeRoleArn"),
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	region: jsii.String("region"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	templateS3Uri: jsii.String("templateS3Uri"),
//   }
//
// Experimental.
type StackDeploymentProps struct {
	// Template path on disk to cloud assembly (cdk.out).
	// Experimental.
	AbsoluteTemplatePath *string `json:"absoluteTemplatePath" yaml:"absoluteTemplatePath"`
	// Construct path for this stack.
	// Experimental.
	ConstructPath *string `json:"constructPath" yaml:"constructPath"`
	// Artifact ID for this stack.
	// Experimental.
	StackArtifactId *string `json:"stackArtifactId" yaml:"stackArtifactId"`
	// Name for this stack.
	// Experimental.
	StackName *string `json:"stackName" yaml:"stackName"`
	// Account where the stack should be deployed.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// Assets referenced by this stack.
	// Experimental.
	Assets *[]*StackAsset `json:"assets" yaml:"assets"`
	// Role to assume before deploying this stack.
	// Experimental.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Execution role to pass to CloudFormation.
	// Experimental.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// Region where the stack should be deployed.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// Tags to apply to the stack.
	// Experimental.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// The S3 URL which points to the template asset location in the publishing bucket.
	// Experimental.
	TemplateS3Uri *string `json:"templateS3Uri" yaml:"templateS3Uri"`
}

// A single output of a Stack.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifactPath artifactPath
//   stackOutput := pipelines.NewStackOutput(artifactPath, jsii.String("outputName"))
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type StackOutput interface {
	// The artifact and file the output is stored in.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ArtifactFile() awscodepipeline.ArtifactPath
	// The name of the output in the JSON object in the file.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewStackOutput_Override(s StackOutput, artifactFile awscodepipeline.ArtifactPath, outputName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.StackOutput",
		[]interface{}{artifactFile, outputName},
		s,
	)
}

// A Reference to a Stack Output.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var cfnOutput cfnOutput
//   stackOutputReference := pipelines.stackOutputReference.fromCfnOutput(cfnOutput)
//
// Experimental.
type StackOutputReference interface {
	// Output name of the producing stack.
	// Experimental.
	OutputName() *string
	// A human-readable description of the producing stack.
	// Experimental.
	StackDescription() *string
	// Whether or not this stack output is being produced by the given Stack deployment.
	// Experimental.
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

// Instructions for additional steps that are run at stack level.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var stack stack
//   var step step
//   stackSteps := &stackSteps{
//   	stack: stack,
//
//   	// the properties below are optional
//   	changeSet: []*step{
//   		step,
//   	},
//   	post: []*step{
//   		step,
//   	},
//   	pre: []*step{
//   		step,
//   	},
//   }
//
// Experimental.
type StackSteps struct {
	// The stack you want the steps to run in.
	// Experimental.
	Stack awscdk.Stack `json:"stack" yaml:"stack"`
	// Steps that execute after stack is prepared but before stack is deployed.
	// Experimental.
	ChangeSet *[]Step `json:"changeSet" yaml:"changeSet"`
	// Steps that execute after stack is deployed.
	// Experimental.
	Post *[]Step `json:"post" yaml:"post"`
	// Steps that execute before stack is prepared.
	// Experimental.
	Pre *[]Step `json:"pre" yaml:"pre"`
}

// Deployment of a single `Stage`.
//
// A `Stage` consists of one or more `Stacks`, which will be
// deployed in dependency order.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var stack stack
//   var stage stage
//   var step step
//   stageDeployment := pipelines.stageDeployment.fromStage(stage, &stageDeploymentProps{
//   	post: []*step{
//   		step,
//   	},
//   	pre: []*step{
//   		step,
//   	},
//   	stackSteps: []stackSteps{
//   		&stackSteps{
//   			stack: stack,
//
//   			// the properties below are optional
//   			changeSet: []*step{
//   				step,
//   			},
//   			post: []*step{
//   				step,
//   			},
//   			pre: []*step{
//   				step,
//   			},
//   		},
//   	},
//   	stageName: jsii.String("stageName"),
//   })
//
// Experimental.
type StageDeployment interface {
	// Additional steps that are run after all of the stacks in the stage.
	// Experimental.
	Post() *[]Step
	// Additional steps that are run before any of the stacks in the stage.
	// Experimental.
	Pre() *[]Step
	// The stacks deployed in this stage.
	// Experimental.
	Stacks() *[]StackDeployment
	// Instructions for additional steps that are run at stack level.
	// Experimental.
	StackSteps() *[]*StackSteps
	// The display name of this stage.
	// Experimental.
	StageName() *string
	// Add an additional step to run after all of the stacks in this stage.
	// Experimental.
	AddPost(steps ...Step)
	// Add an additional step to run before any of the stacks in this stage.
	// Experimental.
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

func (j *jsiiProxy_StageDeployment) StackSteps() *[]*StackSteps {
	var returns *[]*StackSteps
	_jsii_.Get(
		j,
		"stackSteps",
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
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var stack stack
//   var step step
//   stageDeploymentProps := &stageDeploymentProps{
//   	post: []*step{
//   		step,
//   	},
//   	pre: []*step{
//   		step,
//   	},
//   	stackSteps: []stackSteps{
//   		&stackSteps{
//   			stack: stack,
//
//   			// the properties below are optional
//   			changeSet: []*step{
//   				step,
//   			},
//   			post: []*step{
//   				step,
//   			},
//   			pre: []*step{
//   				step,
//   			},
//   		},
//   	},
//   	stageName: jsii.String("stageName"),
//   }
//
// Experimental.
type StageDeploymentProps struct {
	// Additional steps to run after all of the stacks in the stage.
	// Experimental.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	// Experimental.
	Pre *[]Step `json:"pre" yaml:"pre"`
	// Instructions for additional steps that are run at the stack level.
	// Experimental.
	StackSteps *[]*StackSteps `json:"stackSteps" yaml:"stackSteps"`
	// Stage name to use in the pipeline.
	// Experimental.
	StageName *string `json:"stageName" yaml:"stageName"`
}

// Options for a convention-based synth using NPM.
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &cdkPipelineProps{
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   	synthAction: pipelines.simpleSynthAction.standardNpmSynth(&standardNpmSynthOptions{
//   		sourceArtifact: sourceArtifact,
//   		cloudAssemblyArtifact: cloudAssemblyArtifact,
//   		environment: &buildEnvironment{
//   			privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type StandardNpmSynthOptions struct {
	// The artifact where the CloudAssembly should be emitted.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact" yaml:"sourceArtifact"`
	// Name of the build action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables" yaml:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The build command.
	//
	// By default, we assume NPM projects are either written in JavaScript or are
	// using `ts-node`, so don't need a build command.
	//
	// Otherwise, put the build command here, for example `npm run build`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildCommand *string `json:"buildCommand" yaml:"buildCommand"`
	// The install command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	InstallCommand *string `json:"installCommand" yaml:"installCommand"`
	// The synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SynthCommand *string `json:"synthCommand" yaml:"synthCommand"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TestCommands *[]*string `json:"testCommands" yaml:"testCommands"`
}

// Options for a convention-based synth using Yarn.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ec2 "github.com/aws/aws-cdk-go/awscdk/aws_ec2"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"import awscdk "github.com/aws/aws-cdk-go/awscdk"import s3 "github.com/aws/aws-cdk-go/awscdk/aws_s3"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var bucket bucket
//   var buildImage iBuildImage
//   var buildSpec buildSpec
//   var policyStatement policyStatement
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//   standardYarnSynthOptions := &standardYarnSynthOptions{
//   	cloudAssemblyArtifact: artifact,
//   	sourceArtifact: artifact,
//
//   	// the properties below are optional
//   	actionName: jsii.String("actionName"),
//   	additionalArtifacts: []additionalArtifact{
//   		&additionalArtifact{
//   			artifact: artifact,
//   			directory: jsii.String("directory"),
//   		},
//   	},
//   	buildCommand: jsii.String("buildCommand"),
//   	buildSpec: buildSpec,
//   	copyEnvironmentVariables: []*string{
//   		jsii.String("copyEnvironmentVariables"),
//   	},
//   	environment: &buildEnvironment{
//   		buildImage: buildImage,
//   		certificate: &buildEnvironmentCertificate{
//   			bucket: bucket,
//   			objectKey: jsii.String("objectKey"),
//   		},
//   		computeType: codebuild.computeType_SMALL,
//   		environmentVariables: map[string]buildEnvironmentVariable{
//   			"environmentVariablesKey": &buildEnvironmentVariable{
//   				"value": value,
//
//   				// the properties below are optional
//   				"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			},
//   		},
//   		privileged: jsii.Boolean(false),
//   	},
//   	environmentVariables: map[string]*buildEnvironmentVariable{
//   		"environmentVariablesKey": &buildEnvironmentVariable{
//   			"value": value,
//
//   			// the properties below are optional
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   		},
//   	},
//   	installCommand: jsii.String("installCommand"),
//   	projectName: jsii.String("projectName"),
//   	rolePolicyStatements: []*policyStatement{
//   		policyStatement,
//   	},
//   	subdirectory: jsii.String("subdirectory"),
//   	subnetSelection: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnetName: jsii.String("subnetName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: ec2.subnetType_ISOLATED,
//   	},
//   	synthCommand: jsii.String("synthCommand"),
//   	testCommands: []*string{
//   		jsii.String("testCommands"),
//   	},
//   	vpc: vpc,
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type StandardYarnSynthOptions struct {
	// The artifact where the CloudAssembly should be emitted.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyArtifact awscodepipeline.Artifact `json:"cloudAssemblyArtifact" yaml:"cloudAssemblyArtifact"`
	// The source artifact of the CodePipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SourceArtifact awscodepipeline.Artifact `json:"sourceArtifact" yaml:"sourceArtifact"`
	// Name of the build action.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Produce additional output artifacts after the build based on the given directories.
	//
	// Can be used to produce additional artifacts during the build step,
	// separate from the cloud assembly, which can be used further on in the
	// pipeline.
	//
	// Directories are evaluated with respect to `subdirectory`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AdditionalArtifacts *[]*AdditionalArtifact `json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// custom BuildSpec that is merged with the generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Environment variables to copy over from parent env.
	//
	// These are environment variables that are being used by the build.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CopyEnvironmentVariables *[]*string `json:"copyEnvironmentVariables" yaml:"copyEnvironmentVariables"`
	// Build environment to use for CodeBuild job.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Environment *awscodebuild.BuildEnvironment `json:"environment" yaml:"environment"`
	// Environment variables to send into build.
	//
	// NOTE: You may run into the 1000-character limit for the Action configuration if you have a large
	// number of variables or if their names or values are very long.
	// If you do, pass them to the underlying CodeBuild project directly in `environment` instead.
	// However, you will not be able to use CodePipeline Variables in this case.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariables" yaml:"environmentVariables"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Directory inside the source where package.json and cdk.json are located.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Subdirectory *string `json:"subdirectory" yaml:"subdirectory"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The VPC where to execute the SimpleSynth.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
	// The build command.
	//
	// By default, we assume NPM projects are either written in JavaScript or are
	// using `ts-node`, so don't need a build command.
	//
	// Otherwise, put the build command here, for example `npm run build`.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildCommand *string `json:"buildCommand" yaml:"buildCommand"`
	// The install command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	InstallCommand *string `json:"installCommand" yaml:"installCommand"`
	// The synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	SynthCommand *string `json:"synthCommand" yaml:"synthCommand"`
	// Test commands.
	//
	// These commands are run after the build commands but before the
	// synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	TestCommands *[]*string `json:"testCommands" yaml:"testCommands"`
}

// A generic Step which can be added to a Pipeline.
//
// Steps can be used to add Sources, Build Actions and Validations
// to your pipeline.
//
// This class is abstract. See specific subclasses of Step for
// useful steps to add to your Pipeline.
//
// Example:
//   // Step A will depend on step B and step B will depend on step C
//   orderedSteps := pipelines.step.sequence([]step{
//   	pipelines.NewManualApprovalStep(jsii.String("A")),
//   	pipelines.NewManualApprovalStep(jsii.String("B")),
//   	pipelines.NewManualApprovalStep(jsii.String("C")),
//   })
//
// Experimental.
type Step interface {
	IFileSetProducer
	// Return the steps this step depends on, based on the FileSets it requires.
	// Experimental.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	// Experimental.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	// Experimental.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	// Experimental.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	// Experimental.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	// Experimental.
	AddDependencyFileSet(fs FileSet)
	// Add a dependency on another step.
	// Experimental.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	// Experimental.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
	// Experimental.
	DiscoverReferencedOutputs(structure interface{})
	// Return a string representation of this Step.
	// Experimental.
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

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
// Experimental.
func Step_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"monocdk.pipelines.Step",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Step) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (s *jsiiProxy_Step) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		s,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (s *jsiiProxy_Step) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (s *jsiiProxy_Step) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		s,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dockerCredential dockerCredential
//   updatePipelineAction := pipelines.NewUpdatePipelineAction(this, jsii.String("MyUpdatePipelineAction"), &updatePipelineActionProps{
//   	cloudAssemblyInput: artifact,
//   	pipelineStackHierarchicalId: jsii.String("pipelineStackHierarchicalId"),
//
//   	// the properties below are optional
//   	buildSpec: buildSpec,
//   	cdkCliVersion: jsii.String("cdkCliVersion"),
//   	dockerCredentials: []*dockerCredential{
//   		dockerCredential,
//   	},
//   	pipelineStackName: jsii.String("pipelineStackName"),
//   	privileged: jsii.Boolean(false),
//   	projectName: jsii.String("projectName"),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type UpdatePipelineAction interface {
	awscdk.Construct
	awscodepipeline.IAction
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnPrepare()
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewUpdatePipelineAction_Override(u UpdatePipelineAction, scope constructs.Construct, id *string, props *UpdatePipelineActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.UpdatePipelineAction",
		[]interface{}{scope, id, props},
		u,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
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

func (u *jsiiProxy_UpdatePipelineAction) OnPrepare() {
	_jsii_.InvokeVoid(
		u,
		"onPrepare",
		nil, // no parameters
	)
}

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

func (u *jsiiProxy_UpdatePipelineAction) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (u *jsiiProxy_UpdatePipelineAction) Prepare() {
	_jsii_.InvokeVoid(
		u,
		"prepare",
		nil, // no parameters
	)
}

func (u *jsiiProxy_UpdatePipelineAction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		u,
		"synthesize",
		[]interface{}{session},
	)
}

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
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import codebuild "github.com/aws/aws-cdk-go/awscdk/aws_codebuild"import awscdk "github.com/aws/aws-cdk-go/awscdk"import codepipeline "github.com/aws/aws-cdk-go/awscdk/aws_codepipeline"import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var artifact artifact
//   var buildSpec buildSpec
//   var dockerCredential dockerCredential
//   updatePipelineActionProps := &updatePipelineActionProps{
//   	cloudAssemblyInput: artifact,
//   	pipelineStackHierarchicalId: jsii.String("pipelineStackHierarchicalId"),
//
//   	// the properties below are optional
//   	buildSpec: buildSpec,
//   	cdkCliVersion: jsii.String("cdkCliVersion"),
//   	dockerCredentials: []*dockerCredential{
//   		dockerCredential,
//   	},
//   	pipelineStackName: jsii.String("pipelineStackName"),
//   	privileged: jsii.Boolean(false),
//   	projectName: jsii.String("projectName"),
//   }
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type UpdatePipelineActionProps struct {
	// The CodePipeline artifact that holds the Cloud Assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CloudAssemblyInput awscodepipeline.Artifact `json:"cloudAssemblyInput" yaml:"cloudAssemblyInput"`
	// Hierarchical id of the pipeline stack.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	PipelineStackHierarchicalId *string `json:"pipelineStackHierarchicalId" yaml:"pipelineStackHierarchicalId"`
	// Custom BuildSpec that is merged with generated one.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	BuildSpec awscodebuild.BuildSpec `json:"buildSpec" yaml:"buildSpec"`
	// Version of CDK CLI to 'npm install'.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CdkCliVersion *string `json:"cdkCliVersion" yaml:"cdkCliVersion"`
	// Docker registries and associated credentials necessary during the pipeline self-update stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	DockerCredentials *[]DockerCredential `json:"dockerCredentials" yaml:"dockerCredentials"`
	// Name of the pipeline stack.
	// Deprecated: - Use `pipelineStackHierarchicalId` instead.
	PipelineStackName *string `json:"pipelineStackName" yaml:"pipelineStackName"`
	// Whether the build step should run in privileged mode.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Privileged *bool `json:"privileged" yaml:"privileged"`
	// Name of the CodeBuild project.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ProjectName *string `json:"projectName" yaml:"projectName"`
}

// Multiple stages that are deployed in parallel.
//
// Example:
//   var pipeline codePipeline
//   europeWave := pipeline.addWave(jsii.String("Europe"))
//   europeWave.addStage(NewMyApplicationStage(this, jsii.String("Ireland"), &stageProps{
//   	env: &environment{
//   		region: jsii.String("eu-west-1"),
//   	},
//   }))
//   europeWave.addStage(NewMyApplicationStage(this, jsii.String("Germany"), &stageProps{
//   	env: &environment{
//   		region: jsii.String("eu-central-1"),
//   	},
//   }))
//
// Experimental.
type Wave interface {
	// Identifier for this Wave.
	// Experimental.
	Id() *string
	// Additional steps that are run after all of the stages in the wave.
	// Experimental.
	Post() *[]Step
	// Additional steps that are run before any of the stages in the wave.
	// Experimental.
	Pre() *[]Step
	// The stages that are deployed in this wave.
	// Experimental.
	Stages() *[]StageDeployment
	// Add an additional step to run after all of the stages in this wave.
	// Experimental.
	AddPost(steps ...Step)
	// Add an additional step to run before any of the stages in this wave.
	// Experimental.
	AddPre(steps ...Step)
	// Add a Stage to this wave.
	//
	// It will be deployed in parallel with all other stages in this
	// wave.
	// Experimental.
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
//
// Example:
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//
//   	// Turn this on because the pipeline uses Docker image assets
//   	dockerEnabledForSelfMutation: jsii.Boolean(true),
//   })
//
//   pipeline.addWave(jsii.String("MyWave"), &waveOptions{
//   	post: []step{
//   		pipelines.NewCodeBuildStep(jsii.String("RunApproval"), &codeBuildStepProps{
//   			commands: []*string{
//   				jsii.String("command-from-image"),
//   			},
//   			buildEnvironment: &buildEnvironment{
//   				// The user of a Docker image asset in the pipeline requires turning on
//   				// 'dockerEnabledForSelfMutation'.
//   				buildImage: codebuild.linuxBuildImage.fromAsset(this, jsii.String("Image"), &dockerImageAssetProps{
//   					directory: jsii.String("./docker-image"),
//   				}),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type WaveOptions struct {
	// Additional steps to run after all of the stages in the wave.
	// Experimental.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	// Experimental.
	Pre *[]Step `json:"pre" yaml:"pre"`
}

// Construction properties for a `Wave`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import pipelines "github.com/aws/aws-cdk-go/awscdk/pipelines"
//
//   var step step
//   waveProps := &waveProps{
//   	post: []*step{
//   		step,
//   	},
//   	pre: []*step{
//   		step,
//   	},
//   }
//
// Experimental.
type WaveProps struct {
	// Additional steps to run after all of the stages in the wave.
	// Experimental.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	// Experimental.
	Pre *[]Step `json:"pre" yaml:"pre"`
}

