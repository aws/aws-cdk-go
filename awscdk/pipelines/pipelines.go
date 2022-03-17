package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipelineactions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options to pass to `addStage`.
//
// TODO: EXAMPLE
//
type AddStageOpts struct {
	// Additional steps to run after all of the stacks in the stage.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	Pre *[]Step `json:"pre" yaml:"pre"`
	// Instructions for stack level steps.
	StackSteps *[]*StackSteps `json:"stackSteps" yaml:"stackSteps"`
}

// Translate FileSets to CodePipeline Artifacts.
//
// TODO: EXAMPLE
//
type ArtifactMap interface {
	ToCodePipeline(x FileSet) awscodepipeline.Artifact
}

// The jsii proxy struct for ArtifactMap
type jsiiProxy_ArtifactMap struct {
	_ byte // padding
}

func NewArtifactMap() ArtifactMap {
	_init_.Initialize()

	j := jsiiProxy_ArtifactMap{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ArtifactMap",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewArtifactMap_Override(a ArtifactMap) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ArtifactMap",
		nil, // no parameters
		a,
	)
}

// Return the matching CodePipeline artifact for a FileSet.
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

// Type of the asset that is being published.
type AssetType string

const (
	AssetType_FILE AssetType = "FILE"
	AssetType_DOCKER_IMAGE AssetType = "DOCKER_IMAGE"
)

// Options for customizing a single CodeBuild project.
//
// TODO: EXAMPLE
//
type CodeBuildOptions struct {
	// Partial build environment, will be combined with other build environments that apply.
	BuildEnvironment *awscodebuild.BuildEnvironment `json:"buildEnvironment" yaml:"buildEnvironment"`
	// Partial buildspec, will be combined with other buildspecs that apply.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	PartialBuildSpec awscodebuild.BuildSpec `json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Policy statements to add to role.
	RolePolicy *[]awsiam.PolicyStatement `json:"rolePolicy" yaml:"rolePolicy"`
	// Which security group(s) to associate with the project network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The VPC where to create the CodeBuild network interfaces in.
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
// ```
//
// TODO: EXAMPLE
//
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
	Timeout() awscdk.Duration
	Vpc() awsec2.IVpc
	AddDependencyFileSet(fs FileSet)
	AddOutputDirectory(directory *string) FileSet
	AddStepDependency(step Step)
	ConfigurePrimaryOutput(fs FileSet)
	DiscoverReferencedOutputs(structure interface{})
	ExportedVariable(variableName *string) *string
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


func NewCodeBuildStep(id *string, props *CodeBuildStepProps) CodeBuildStep {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildStep{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.CodeBuildStep",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewCodeBuildStep_Override(c CodeBuildStep, id *string, props *CodeBuildStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.CodeBuildStep",
		[]interface{}{id, props},
		c,
	)
}

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
func CodeBuildStep_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodeBuildStep",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
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

// Add a dependency on another step.
func (c *jsiiProxy_CodeBuildStep) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

// Configure the given FileSet as the primary output of this step.
func (c *jsiiProxy_CodeBuildStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
//
// Should be called by subclasses based on what the user passes in as
// construction properties. The format of the structure passed in here does
// not have to correspond exactly to what gets rendered into the engine, it
// just needs to contain the same amount of data.
func (c *jsiiProxy_CodeBuildStep) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		c,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

// Reference a CodePipeline variable defined by the CodeBuildStep.
//
// The variable must be set in the shell of the CodeBuild step when
// it finishes its `post_build` phase.
//
// TODO: EXAMPLE
//
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

// Configure the given output directory as primary output.
//
// If no primary output has been configured yet, this directory
// will become the primary output of this ShellStep, otherwise this
// method will throw if the given directory is different than the
// currently configured primary output directory.
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
// TODO: EXAMPLE
//
type CodeBuildStepProps struct {
	// Commands to run.
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
	// ```
	AdditionalInputs *map[string]IFileSetProducer `json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	Env *map[string]*string `json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	Input IFileSetProducer `json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	InstallCommands *[]*string `json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	PrimaryOutputDirectory *string `json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
	// Changes to environment.
	//
	// This environment will be combined with the pipeline's default
	// environment.
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
	PartialBuildSpec awscodebuild.BuildSpec `json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Name for the generated CodeBuild project.
	ProjectName *string `json:"projectName" yaml:"projectName"`
	// Custom execution role to be used for the CodeBuild project.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	RolePolicyStatements *[]awsiam.PolicyStatement `json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `json:"timeout" yaml:"timeout"`
	// The VPC where to execute the SimpleSynth.
	Vpc awsec2.IVpc `json:"vpc" yaml:"vpc"`
}

// Configuration options for a CodeCommit source.
//
// TODO: EXAMPLE
//
type CodeCommitSourceOptions struct {
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html
	//
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Role to be used by on commit event rule.
	//
	// Used only when trigger value is CodeCommitTrigger.EVENTS.
	EventRole awsiam.IRole `json:"eventRole" yaml:"eventRole"`
	// How should CodePipeline detect source changes for this Action.
	Trigger awscodepipelineactions.CodeCommitTrigger `json:"trigger" yaml:"trigger"`
}

// A CDK Pipeline that uses CodePipeline to deploy CDK apps.
//
// This is a `Pipeline` with its `engine` property set to
// `CodePipelineEngine`, and exists for nicer ergonomics for
// users that don't need to switch out engines.
//
// TODO: EXAMPLE
//
type CodePipeline interface {
	PipelineBase
	CloudAssemblyFileSet() FileSet
	Node() constructs.Node
	Pipeline() awscodepipeline.Pipeline
	Synth() IFileSetProducer
	SynthProject() awscodebuild.IProject
	Waves() *[]Wave
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
	AddWave(id *string, options *WaveOptions) Wave
	BuildPipeline()
	DoBuildPipeline()
	ToString() *string
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

func (j *jsiiProxy_CodePipeline) Node() constructs.Node {
	var returns constructs.Node
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


func NewCodePipeline(scope constructs.Construct, id *string, props *CodePipelineProps) CodePipeline {
	_init_.Initialize()

	j := jsiiProxy_CodePipeline{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.CodePipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCodePipeline_Override(c CodePipeline, scope constructs.Construct, id *string, props *CodePipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.CodePipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CodePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipeline",
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
// declare const pipeline: pipelines.CodePipeline;
//
// const wave = pipeline.addWave('MyWave');
// wave.addStage(new MyApplicationStage(this, 'Stage1'));
// wave.addStage(new MyApplicationStage(this, 'Stage2'));
// ```
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
func (c *jsiiProxy_CodePipeline) BuildPipeline() {
	_jsii_.InvokeVoid(
		c,
		"buildPipeline",
		nil, // no parameters
	)
}

// Implemented by subclasses to do the actual pipeline construction.
func (c *jsiiProxy_CodePipeline) DoBuildPipeline() {
	_jsii_.InvokeVoid(
		c,
		"doBuildPipeline",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
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

// The result of adding actions to the pipeline.
//
// TODO: EXAMPLE
//
type CodePipelineActionFactoryResult struct {
	// How many RunOrders were consumed.
	//
	// If you add 1 action, return the value 1 here.
	RunOrdersConsumed *float64 `json:"runOrdersConsumed" yaml:"runOrdersConsumed"`
	// If a CodeBuild project got created, the project.
	Project awscodebuild.IProject `json:"project" yaml:"project"`
}

// A FileSet created from a CodePipeline artifact.
//
// You only need to use this if you want to add CDK Pipeline stages
// add the end of an existing CodePipeline, which should be very rare.
//
// TODO: EXAMPLE
//
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
func CodePipelineFileSet_FromArtifact(artifact awscodepipeline.Artifact) CodePipelineFileSet {
	_init_.Initialize()

	var returns CodePipelineFileSet

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineFileSet",
		"fromArtifact",
		[]interface{}{artifact},
		&returns,
	)

	return returns
}

// Mark the given Step as the producer for this FileSet.
//
// This method can only be called once.
func (c *jsiiProxy_CodePipelineFileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		c,
		"producedBy",
		[]interface{}{producer},
	)
}

// Return a string representation of this FileSet.
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
// TODO: EXAMPLE
//
type CodePipelineProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	Synth IFileSetProducer `json:"synth" yaml:"synth"`
	// Additional customizations to apply to the asset publishing CodeBuild projects.
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
	CliVersion *string `json:"cliVersion" yaml:"cliVersion"`
	// Customize the CodeBuild projects created for this pipeline.
	CodeBuildDefaults *CodeBuildOptions `json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An existing Pipeline to be reused and built upon.
	//
	// [disable-awslint:ref-via-interface]
	CodePipeline awscodepipeline.Pipeline `json:"codePipeline" yaml:"codePipeline"`
	// Create KMS keys for the artifact buckets, allowing cross-account deployments.
	//
	// The artifact buckets have to be encrypted to support deploying CDK apps to
	// another account, so if you want to do that or want to have your artifact
	// buckets encrypted, be sure to set this value to `true`.
	//
	// Be aware there is a cost associated with maintaining the KMS keys.
	CrossAccountKeys *bool `json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
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
	DockerEnabledForSynth *bool `json:"dockerEnabledForSynth" yaml:"dockerEnabledForSynth"`
	// The name of the CodePipeline pipeline.
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
	PublishAssetsInParallel *bool `json:"publishAssetsInParallel" yaml:"publishAssetsInParallel"`
	// Reuse the same cross region support stack for all pipelines in the App.
	ReuseCrossRegionSupportStacks *bool `json:"reuseCrossRegionSupportStacks" yaml:"reuseCrossRegionSupportStacks"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	SelfMutation *bool `json:"selfMutation" yaml:"selfMutation"`
	// Additional customizations to apply to the self mutation CodeBuild projects.
	SelfMutationCodeBuildDefaults *CodeBuildOptions `json:"selfMutationCodeBuildDefaults" yaml:"selfMutationCodeBuildDefaults"`
	// Additional customizations to apply to the synthesize CodeBuild projects.
	SynthCodeBuildDefaults *CodeBuildOptions `json:"synthCodeBuildDefaults" yaml:"synthCodeBuildDefaults"`
}

// Factory for CodePipeline source steps.
//
// This class contains a number of factory methods for the different types
// of sources that CodePipeline supports.
//
// TODO: EXAMPLE
//
type CodePipelineSource interface {
	Step
	ICodePipelineActionFactory
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	AddStepDependency(step Step)
	ConfigurePrimaryOutput(fs FileSet)
	DiscoverReferencedOutputs(structure interface{})
	GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64, variablesNamespace *string) awscodepipelineactions.Action
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
	SourceAttribute(name *string) *string
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


func NewCodePipelineSource_Override(c CodePipelineSource, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.CodePipelineSource",
		[]interface{}{id},
		c,
	)
}

// Returns a CodeCommit source.
//
// TODO: EXAMPLE
//
func CodePipelineSource_CodeCommit(repository awscodecommit.IRepository, branch *string, props *CodeCommitSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineSource",
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
// ```
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/welcome-connections.html
//
func CodePipelineSource_Connection(repoString *string, branch *string, props *ConnectionSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineSource",
		"connection",
		[]interface{}{repoString, branch, props},
		&returns,
	)

	return returns
}

// Returns an ECR source.
//
// TODO: EXAMPLE
//
func CodePipelineSource_Ecr(repository awsecr.IRepository, props *ECRSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineSource",
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
// * **admin:repo_hook** - if you plan to use webhooks (true by default)
func CodePipelineSource_GitHub(repoString *string, branch *string, props *GitHubSourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineSource",
		"gitHub",
		[]interface{}{repoString, branch, props},
		&returns,
	)

	return returns
}

// Returns an S3 source.
//
// TODO: EXAMPLE
//
func CodePipelineSource_S3(bucket awss3.IBucket, objectKey *string, props *S3SourceOptions) CodePipelineSource {
	_init_.Initialize()

	var returns CodePipelineSource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineSource",
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
func CodePipelineSource_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipelineSource",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
func (c *jsiiProxy_CodePipelineSource) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Add a dependency on another step.
func (c *jsiiProxy_CodePipelineSource) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

// Configure the given FileSet as the primary output of this step.
func (c *jsiiProxy_CodePipelineSource) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
//
// Should be called by subclasses based on what the user passes in as
// construction properties. The format of the structure passed in here does
// not have to correspond exactly to what gets rendered into the engine, it
// just needs to contain the same amount of data.
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

// Create the desired Action and add it to the pipeline.
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

// Return an attribute of the current source revision.
//
// These values can be passed into the environment variables of pipeline steps,
// so your steps can access information about the source revision.
//
// What attributes are available depends on the type of source. These attributes
// are supported:
//
// - GitHub, CodeCommit, and CodeStarSourceConnection
//    - `AuthorDate`
//    - `BranchName`
//    - `CommitId`
//    - `CommitMessage`
// - GitHub, CodeCommit and ECR
//    - `RepositoryName`
// - GitHub and CodeCommit
//    - `CommitterDate`
// - GitHub
//    - `CommitUrl`
// - CodeStarSourceConnection
//    - `FullRepositoryName`
// - S3
//    - `ETag`
//    - `VersionId`
// - ECR
//    - `ImageDigest`
//    - `ImageTag`
//    - `ImageURI`
//    - `RegistryId`
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-variables.html#reference-variables-list
//
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

// Return a string representation of this Step.
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
// TODO: EXAMPLE
//
type ConfirmPermissionsBroadening interface {
	Step
	ICodePipelineActionFactory
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	AddStepDependency(step Step)
	ConfigurePrimaryOutput(fs FileSet)
	DiscoverReferencedOutputs(structure interface{})
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


func NewConfirmPermissionsBroadening(id *string, props *PermissionsBroadeningCheckProps) ConfirmPermissionsBroadening {
	_init_.Initialize()

	j := jsiiProxy_ConfirmPermissionsBroadening{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ConfirmPermissionsBroadening",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewConfirmPermissionsBroadening_Override(c ConfirmPermissionsBroadening, id *string, props *PermissionsBroadeningCheckProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ConfirmPermissionsBroadening",
		[]interface{}{id, props},
		c,
	)
}

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
func ConfirmPermissionsBroadening_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.ConfirmPermissionsBroadening",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
func (c *jsiiProxy_ConfirmPermissionsBroadening) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Add a dependency on another step.
func (c *jsiiProxy_ConfirmPermissionsBroadening) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

// Configure the given FileSet as the primary output of this step.
func (c *jsiiProxy_ConfirmPermissionsBroadening) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
//
// Should be called by subclasses based on what the user passes in as
// construction properties. The format of the structure passed in here does
// not have to correspond exactly to what gets rendered into the engine, it
// just needs to contain the same amount of data.
func (c *jsiiProxy_ConfirmPermissionsBroadening) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		c,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

// Create the desired Action and add it to the pipeline.
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
// TODO: EXAMPLE
//
type ConnectionSourceOptions struct {
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// TODO: EXAMPLE
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	ConnectionArn *string `json:"connectionArn" yaml:"connectionArn"`
	// Whether the output should be the contents of the repository (which is the default), or a link that allows CodeBuild to clone the repository before building.
	//
	// **Note**: if this option is true,
	// then only CodeBuild actions can use the resulting {@link output}.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	CodeBuildCloneOutput *bool `json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	TriggerOnPush *bool `json:"triggerOnPush" yaml:"triggerOnPush"`
}

// Represents credentials used to access a Docker registry.
//
// TODO: EXAMPLE
//
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


func NewDockerCredential_Override(d DockerCredential, usages *[]DockerCredentialUsage) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.DockerCredential",
		[]interface{}{usages},
		d,
	)
}

// Creates a DockerCredential for a registry, based on its domain name (e.g., 'www.example.com').
func DockerCredential_CustomRegistry(registryDomain *string, secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	var returns DockerCredential

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.DockerCredential",
		"customRegistry",
		[]interface{}{registryDomain, secret, opts},
		&returns,
	)

	return returns
}

// Creates a DockerCredential for DockerHub.
//
// Convenience method for `customRegistry('https://index.docker.io/v1/', opts)`.
func DockerCredential_DockerHub(secret awssecretsmanager.ISecret, opts *ExternalDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	var returns DockerCredential

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.DockerCredential",
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
func DockerCredential_Ecr(repositories *[]awsecr.IRepository, opts *EcrDockerCredentialOptions) DockerCredential {
	_init_.Initialize()

	var returns DockerCredential

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.DockerCredential",
		"ecr",
		[]interface{}{repositories, opts},
		&returns,
	)

	return returns
}

// Grant read-only access to the registry credentials.
//
// This grants read access to any secrets, and pull access to any repositories.
func (d *jsiiProxy_DockerCredential) GrantRead(grantee awsiam.IGrantable, usage DockerCredentialUsage) {
	_jsii_.InvokeVoid(
		d,
		"grantRead",
		[]interface{}{grantee, usage},
	)
}

// Defines which stages of a pipeline require the specified credentials.
//
// TODO: EXAMPLE
//
type DockerCredentialUsage string

const (
	DockerCredentialUsage_SYNTH DockerCredentialUsage = "SYNTH"
	DockerCredentialUsage_SELF_UPDATE DockerCredentialUsage = "SELF_UPDATE"
	DockerCredentialUsage_ASSET_PUBLISHING DockerCredentialUsage = "ASSET_PUBLISHING"
)

// Options for ECR sources.
//
// TODO: EXAMPLE
//
type ECRSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// The image tag that will be checked for changes.
	ImageTag *string `json:"imageTag" yaml:"imageTag"`
}

// Options for defining access for a Docker Credential composed of ECR repos.
//
// TODO: EXAMPLE
//
type EcrDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	AssumeRole awsiam.IRole `json:"assumeRole" yaml:"assumeRole"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	Usages *[]DockerCredentialUsage `json:"usages" yaml:"usages"`
}

// Options for defining credentials for a Docker Credential.
//
// TODO: EXAMPLE
//
type ExternalDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	AssumeRole awsiam.IRole `json:"assumeRole" yaml:"assumeRole"`
	// The name of the JSON field of the secret which contains the secret/password.
	SecretPasswordField *string `json:"secretPasswordField" yaml:"secretPasswordField"`
	// The name of the JSON field of the secret which contains the user/login name.
	SecretUsernameField *string `json:"secretUsernameField" yaml:"secretUsernameField"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	Usages *[]DockerCredentialUsage `json:"usages" yaml:"usages"`
}

// A set of files traveling through the deployment pipeline.
//
// Individual steps in the pipeline produce or consume
// `FileSet`s.
//
// TODO: EXAMPLE
//
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


func NewFileSet(id *string, producer Step) FileSet {
	_init_.Initialize()

	j := jsiiProxy_FileSet{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.FileSet",
		[]interface{}{id, producer},
		&j,
	)

	return &j
}

func NewFileSet_Override(f FileSet, id *string, producer Step) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.FileSet",
		[]interface{}{id, producer},
		f,
	)
}

// Mark the given Step as the producer for this FileSet.
//
// This method can only be called once.
func (f *jsiiProxy_FileSet) ProducedBy(producer Step) {
	_jsii_.InvokeVoid(
		f,
		"producedBy",
		[]interface{}{producer},
	)
}

// Return a string representation of this FileSet.
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
// TODO: EXAMPLE
//
type FileSetLocation struct {
	// The (relative) directory where the FileSet is found.
	Directory *string `json:"directory" yaml:"directory"`
	// The FileSet object.
	FileSet FileSet `json:"fileSet" yaml:"fileSet"`
}

// Options for GitHub sources.
//
// TODO: EXAMPLE
//
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
	// * **admin:repo_hook** - if you plan to use webhooks (true by default)
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/GitHub-create-personal-token-CLI.html
	//
	Authentication awscdk.SecretValue `json:"authentication" yaml:"authentication"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action.
	// With "POLL", CodePipeline periodically checks the source for changes.
	// With "None", the action is not triggered through changes in the source.
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
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
type ICodePipelineActionFactory interface {
	// Create the desired Action and add it to the pipeline.
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
type IFileSetProducer interface {
	// The `FileSet` produced by this file set producer.
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

// A manual approval step.
//
// If this step is added to a Pipeline, the Pipeline will
// be paused waiting for a human to resume it
//
// Only engines that support pausing the deployment will
// support this step type.
//
// TODO: EXAMPLE
//
type ManualApprovalStep interface {
	Step
	Comment() *string
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	AddStepDependency(step Step)
	ConfigurePrimaryOutput(fs FileSet)
	DiscoverReferencedOutputs(structure interface{})
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


func NewManualApprovalStep(id *string, props *ManualApprovalStepProps) ManualApprovalStep {
	_init_.Initialize()

	j := jsiiProxy_ManualApprovalStep{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ManualApprovalStep",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewManualApprovalStep_Override(m ManualApprovalStep, id *string, props *ManualApprovalStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ManualApprovalStep",
		[]interface{}{id, props},
		m,
	)
}

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
func ManualApprovalStep_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.ManualApprovalStep",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
func (m *jsiiProxy_ManualApprovalStep) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		m,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Add a dependency on another step.
func (m *jsiiProxy_ManualApprovalStep) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		m,
		"addStepDependency",
		[]interface{}{step},
	)
}

// Configure the given FileSet as the primary output of this step.
func (m *jsiiProxy_ManualApprovalStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		m,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
//
// Should be called by subclasses based on what the user passes in as
// construction properties. The format of the structure passed in here does
// not have to correspond exactly to what gets rendered into the engine, it
// just needs to contain the same amount of data.
func (m *jsiiProxy_ManualApprovalStep) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		m,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

// Return a string representation of this Step.
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
// TODO: EXAMPLE
//
type ManualApprovalStepProps struct {
	// The comment to display with this manual approval.
	Comment *string `json:"comment" yaml:"comment"`
}

// Properties for a `PermissionsBroadeningCheck`.
//
// TODO: EXAMPLE
//
type PermissionsBroadeningCheckProps struct {
	// The CDK Stage object to check the stacks of.
	//
	// This should be the same Stage object you are passing to `addStage()`.
	Stage awscdk.Stage `json:"stage" yaml:"stage"`
	// Topic to send notifications when a human needs to give manual confirmation.
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
type PipelineBase interface {
	constructs.Construct
	CloudAssemblyFileSet() FileSet
	Node() constructs.Node
	Synth() IFileSetProducer
	Waves() *[]Wave
	AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment
	AddWave(id *string, options *WaveOptions) Wave
	BuildPipeline()
	DoBuildPipeline()
	ToString() *string
}

// The jsii proxy struct for PipelineBase
type jsiiProxy_PipelineBase struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_PipelineBase) Node() constructs.Node {
	var returns constructs.Node
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


func NewPipelineBase_Override(p PipelineBase, scope constructs.Construct, id *string, props *PipelineBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.PipelineBase",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func PipelineBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.PipelineBase",
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
// declare const pipeline: pipelines.CodePipeline;
//
// const wave = pipeline.addWave('MyWave');
// wave.addStage(new MyApplicationStage(this, 'Stage1'));
// wave.addStage(new MyApplicationStage(this, 'Stage2'));
// ```
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
func (p *jsiiProxy_PipelineBase) BuildPipeline() {
	_jsii_.InvokeVoid(
		p,
		"buildPipeline",
		nil, // no parameters
	)
}

// Implemented by subclasses to do the actual pipeline construction.
func (p *jsiiProxy_PipelineBase) DoBuildPipeline() {
	_jsii_.InvokeVoid(
		p,
		"doBuildPipeline",
		nil, // no parameters
	)
}

// Returns a string representation of this construct.
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

// Properties for a `Pipeline`.
//
// TODO: EXAMPLE
//
type PipelineBaseProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	Synth IFileSetProducer `json:"synth" yaml:"synth"`
}

// Options for the `CodePipelineActionFactory.produce()` method.
//
// TODO: EXAMPLE
//
type ProduceActionOptions struct {
	// Name the action should get.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// Helper object to translate FileSets to CodePipeline Artifacts.
	Artifacts ArtifactMap `json:"artifacts" yaml:"artifacts"`
	// The pipeline the action is being generated for.
	Pipeline CodePipeline `json:"pipeline" yaml:"pipeline"`
	// RunOrder the action should get.
	RunOrder *float64 `json:"runOrder" yaml:"runOrder"`
	// Scope in which to create constructs.
	Scope constructs.Construct `json:"scope" yaml:"scope"`
	// Whether or not this action is inserted before self mutation.
	//
	// If it is, the action should take care to reflect some part of
	// its own definition in the pipeline action definition, to
	// trigger a restart after self-mutation (if necessary).
	BeforeSelfMutation *bool `json:"beforeSelfMutation" yaml:"beforeSelfMutation"`
	// If this action factory creates a CodeBuild step, default options to inherit.
	CodeBuildDefaults *CodeBuildOptions `json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An input artifact that CodeBuild projects that don't actually need an input artifact can use.
	//
	// CodeBuild Projects MUST have an input artifact in order to be added to the Pipeline. If
	// the Project doesn't actually care about its input (it can be anything), it can use the
	// Artifact passed here.
	FallbackArtifact awscodepipeline.Artifact `json:"fallbackArtifact" yaml:"fallbackArtifact"`
	// If this step is producing outputs, the variables namespace assigned to it.
	//
	// Pass this on to the Action you are creating.
	VariablesNamespace *string `json:"variablesNamespace" yaml:"variablesNamespace"`
}

// Options for S3 sources.
//
// TODO: EXAMPLE
//
type S3SourceOptions struct {
	// The action name used for this source in the CodePipeline.
	ActionName *string `json:"actionName" yaml:"actionName"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	Trigger awscodepipelineactions.S3Trigger `json:"trigger" yaml:"trigger"`
}

// Run shell script commands in the pipeline.
//
// This is a generic step designed
// to be deployment engine agnostic.
//
// TODO: EXAMPLE
//
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
	AddStepDependency(step Step)
	ConfigurePrimaryOutput(fs FileSet)
	DiscoverReferencedOutputs(structure interface{})
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


func NewShellStep(id *string, props *ShellStepProps) ShellStep {
	_init_.Initialize()

	j := jsiiProxy_ShellStep{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ShellStep",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewShellStep_Override(s ShellStep, id *string, props *ShellStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.ShellStep",
		[]interface{}{id, props},
		s,
	)
}

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
func ShellStep_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.ShellStep",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
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

// Add a dependency on another step.
func (s *jsiiProxy_ShellStep) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		s,
		"addStepDependency",
		[]interface{}{step},
	)
}

// Configure the given FileSet as the primary output of this step.
func (s *jsiiProxy_ShellStep) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
//
// Should be called by subclasses based on what the user passes in as
// construction properties. The format of the structure passed in here does
// not have to correspond exactly to what gets rendered into the engine, it
// just needs to contain the same amount of data.
func (s *jsiiProxy_ShellStep) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		s,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

// Configure the given output directory as primary output.
//
// If no primary output has been configured yet, this directory
// will become the primary output of this ShellStep, otherwise this
// method will throw if the given directory is different than the
// currently configured primary output directory.
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
// TODO: EXAMPLE
//
type ShellStepProps struct {
	// Commands to run.
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
	// ```
	AdditionalInputs *map[string]IFileSetProducer `json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	Env *map[string]*string `json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	Input IFileSetProducer `json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	InstallCommands *[]*string `json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	PrimaryOutputDirectory *string `json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
}

// An asset used by a Stack.
//
// TODO: EXAMPLE
//
type StackAsset struct {
	// Asset identifier.
	AssetId *string `json:"assetId" yaml:"assetId"`
	// Absolute asset manifest path.
	//
	// This needs to be made relative at a later point in time, but when this
	// information is parsed we don't know about the root cloud assembly yet.
	AssetManifestPath *string `json:"assetManifestPath" yaml:"assetManifestPath"`
	// Asset selector to pass to `cdk-assets`.
	AssetSelector *string `json:"assetSelector" yaml:"assetSelector"`
	// Type of asset to publish.
	AssetType AssetType `json:"assetType" yaml:"assetType"`
	// Does this asset represent the CloudFormation template for the stack.
	IsTemplate *bool `json:"isTemplate" yaml:"isTemplate"`
	// Role ARN to assume to publish.
	AssetPublishingRoleArn *string `json:"assetPublishingRoleArn" yaml:"assetPublishingRoleArn"`
}

// Deployment of a single Stack.
//
// You don't need to instantiate this class -- it will
// be automatically instantiated as necessary when you
// add a `Stage` to a pipeline.
//
// TODO: EXAMPLE
//
type StackDeployment interface {
	AbsoluteTemplatePath() *string
	Account() *string
	Assets() *[]*StackAsset
	AssumeRoleArn() *string
	ChangeSet() *[]Step
	ConstructPath() *string
	ExecutionRoleArn() *string
	Post() *[]Step
	Pre() *[]Step
	Region() *string
	StackArtifactId() *string
	StackDependencies() *[]StackDeployment
	StackName() *string
	Tags() *map[string]*string
	TemplateAsset() *StackAsset
	TemplateUrl() *string
	AddStackDependency(stackDeployment StackDeployment)
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
func StackDeployment_FromArtifact(stackArtifact cxapi.CloudFormationStackArtifact) StackDeployment {
	_init_.Initialize()

	var returns StackDeployment

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.StackDeployment",
		"fromArtifact",
		[]interface{}{stackArtifact},
		&returns,
	)

	return returns
}

// Add a dependency on another stack.
func (s *jsiiProxy_StackDeployment) AddStackDependency(stackDeployment StackDeployment) {
	_jsii_.InvokeVoid(
		s,
		"addStackDependency",
		[]interface{}{stackDeployment},
	)
}

// Adds steps to each phase of the stack.
func (s *jsiiProxy_StackDeployment) AddStackSteps(pre *[]Step, changeSet *[]Step, post *[]Step) {
	_jsii_.InvokeVoid(
		s,
		"addStackSteps",
		[]interface{}{pre, changeSet, post},
	)
}

// Properties for a `StackDeployment`.
//
// TODO: EXAMPLE
//
type StackDeploymentProps struct {
	// Template path on disk to cloud assembly (cdk.out).
	AbsoluteTemplatePath *string `json:"absoluteTemplatePath" yaml:"absoluteTemplatePath"`
	// Construct path for this stack.
	ConstructPath *string `json:"constructPath" yaml:"constructPath"`
	// Artifact ID for this stack.
	StackArtifactId *string `json:"stackArtifactId" yaml:"stackArtifactId"`
	// Name for this stack.
	StackName *string `json:"stackName" yaml:"stackName"`
	// Account where the stack should be deployed.
	Account *string `json:"account" yaml:"account"`
	// Assets referenced by this stack.
	Assets *[]*StackAsset `json:"assets" yaml:"assets"`
	// Role to assume before deploying this stack.
	AssumeRoleArn *string `json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Execution role to pass to CloudFormation.
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// Region where the stack should be deployed.
	Region *string `json:"region" yaml:"region"`
	// Tags to apply to the stack.
	Tags *map[string]*string `json:"tags" yaml:"tags"`
	// The S3 URL which points to the template asset location in the publishing bucket.
	TemplateS3Uri *string `json:"templateS3Uri" yaml:"templateS3Uri"`
}

// A Reference to a Stack Output.
//
// TODO: EXAMPLE
//
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
func StackOutputReference_FromCfnOutput(output awscdk.CfnOutput) StackOutputReference {
	_init_.Initialize()

	var returns StackOutputReference

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.StackOutputReference",
		"fromCfnOutput",
		[]interface{}{output},
		&returns,
	)

	return returns
}

// Whether or not this stack output is being produced by the given Stack deployment.
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
// TODO: EXAMPLE
//
type StackSteps struct {
	// The stack you want the steps to run in.
	Stack awscdk.Stack `json:"stack" yaml:"stack"`
	// Steps that execute after stack is prepared but before stack is deployed.
	ChangeSet *[]Step `json:"changeSet" yaml:"changeSet"`
	// Steps that execute after stack is deployed.
	Post *[]Step `json:"post" yaml:"post"`
	// Steps that execute before stack is prepared.
	Pre *[]Step `json:"pre" yaml:"pre"`
}

// Deployment of a single `Stage`.
//
// A `Stage` consists of one or more `Stacks`, which will be
// deployed in dependency order.
//
// TODO: EXAMPLE
//
type StageDeployment interface {
	Post() *[]Step
	Pre() *[]Step
	Stacks() *[]StackDeployment
	StackSteps() *[]*StackSteps
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
func StageDeployment_FromStage(stage awscdk.Stage, props *StageDeploymentProps) StageDeployment {
	_init_.Initialize()

	var returns StageDeployment

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.StageDeployment",
		"fromStage",
		[]interface{}{stage, props},
		&returns,
	)

	return returns
}

// Add an additional step to run after all of the stacks in this stage.
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
// TODO: EXAMPLE
//
type StageDeploymentProps struct {
	// Additional steps to run after all of the stacks in the stage.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	Pre *[]Step `json:"pre" yaml:"pre"`
	// Instructions for additional steps that are run at the stack level.
	StackSteps *[]*StackSteps `json:"stackSteps" yaml:"stackSteps"`
	// Stage name to use in the pipeline.
	StageName *string `json:"stageName" yaml:"stageName"`
}

// A generic Step which can be added to a Pipeline.
//
// Steps can be used to add Sources, Build Actions and Validations
// to your pipeline.
//
// This class is abstract. See specific subclasses of Step for
// useful steps to add to your Pipeline
//
// TODO: EXAMPLE
//
type Step interface {
	IFileSetProducer
	Dependencies() *[]Step
	DependencyFileSets() *[]FileSet
	Id() *string
	IsSource() *bool
	PrimaryOutput() FileSet
	AddDependencyFileSet(fs FileSet)
	AddStepDependency(step Step)
	ConfigurePrimaryOutput(fs FileSet)
	DiscoverReferencedOutputs(structure interface{})
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


func NewStep_Override(s Step, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.Step",
		[]interface{}{id},
		s,
	)
}

// Define a sequence of steps to be executed in order.
//
// If you need more fine-grained step ordering, use the `addStepDependency()`
// API. For example, if you want `secondStep` to occur after `firstStep`, call
// `secondStep.addStepDependency(firstStep)`.
func Step_Sequence(steps *[]Step) *[]Step {
	_init_.Initialize()

	var returns *[]Step

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.Step",
		"sequence",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Add an additional FileSet to the set of file sets required by this step.
//
// This will lead to a dependency on the producer of that file set.
func (s *jsiiProxy_Step) AddDependencyFileSet(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

// Add a dependency on another step.
func (s *jsiiProxy_Step) AddStepDependency(step Step) {
	_jsii_.InvokeVoid(
		s,
		"addStepDependency",
		[]interface{}{step},
	)
}

// Configure the given FileSet as the primary output of this step.
func (s *jsiiProxy_Step) ConfigurePrimaryOutput(fs FileSet) {
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
//
// Should be called by subclasses based on what the user passes in as
// construction properties. The format of the structure passed in here does
// not have to correspond exactly to what gets rendered into the engine, it
// just needs to contain the same amount of data.
func (s *jsiiProxy_Step) DiscoverReferencedOutputs(structure interface{}) {
	_jsii_.InvokeVoid(
		s,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

// Return a string representation of this Step.
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

// Multiple stages that are deployed in parallel.
//
// TODO: EXAMPLE
//
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


func NewWave(id *string, props *WaveProps) Wave {
	_init_.Initialize()

	j := jsiiProxy_Wave{}

	_jsii_.Create(
		"aws-cdk-lib.pipelines.Wave",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewWave_Override(w Wave, id *string, props *WaveProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.pipelines.Wave",
		[]interface{}{id, props},
		w,
	)
}

// Add an additional step to run after all of the stages in this wave.
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
// TODO: EXAMPLE
//
type WaveOptions struct {
	// Additional steps to run after all of the stages in the wave.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	Pre *[]Step `json:"pre" yaml:"pre"`
}

// Construction properties for a `Wave`.
//
// TODO: EXAMPLE
//
type WaveProps struct {
	// Additional steps to run after all of the stages in the wave.
	Post *[]Step `json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	Pre *[]Step `json:"pre" yaml:"pre"`
}

