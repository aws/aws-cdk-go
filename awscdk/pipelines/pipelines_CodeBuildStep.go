package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

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
type CodeBuildStep interface {
	ShellStep
	// Custom execution role to be used for the Code Build Action.
	ActionRole() awsiam.IRole
	// Build environment.
	BuildEnvironment() *awscodebuild.BuildEnvironment
	// Caching strategy to use.
	Cache() awscodebuild.Cache
	// Commands to run.
	Commands() *[]*string
	// Return the steps this step depends on, based on the FileSets it requires.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	DependencyFileSets() *[]FileSet
	// Environment variables to set.
	Env() *map[string]*string
	// Set environment variables based on Stack Outputs.
	EnvFromCfnOutputs() *map[string]StackOutputReference
	// The CodeBuild Project's principal.
	GrantPrincipal() awsiam.IPrincipal
	// Identifier for this step.
	Id() *string
	// Input FileSets.
	//
	// A list of `(FileSet, directory)` pairs, which are a copy of the
	// input properties. This list should not be modified directly.
	Inputs() *[]*FileSetLocation
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	InstallCommands() *[]*string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	IsSource() *bool
	// Output FileSets.
	//
	// A list of `(FileSet, directory)` pairs, which are a copy of the
	// input properties. This list should not be modified directly.
	Outputs() *[]*FileSetLocation
	// Additional configuration that can only be configured via BuildSpec.
	//
	// Contains exported variables.
	PartialBuildSpec() awscodebuild.BuildSpec
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	PrimaryOutput() FileSet
	// CodeBuild Project generated for the pipeline.
	//
	// Will only be available after the pipeline has been built.
	Project() awscodebuild.IProject
	// Name for the generated CodeBuild project.
	ProjectName() *string
	// Custom execution role to be used for the CodeBuild project.
	Role() awsiam.IRole
	// Policy statements to add to role used during the synth.
	RolePolicyStatements() *[]awsiam.PolicyStatement
	// Which security group to associate with the script's project network interfaces.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// Which subnets to use.
	SubnetSelection() *awsec2.SubnetSelection
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout() awscdk.Duration
	// The VPC where to execute the SimpleSynth.
	Vpc() awsec2.IVpc
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	AddDependencyFileSet(fs FileSet)
	// Add an additional output FileSet based on a directory.
	//
	// After running the script, the contents of the given directory
	// will be exported as a `FileSet`. Use the `FileSet` as the
	// input to another step.
	//
	// Multiple calls with the exact same directory name string (not normalized)
	// will return the same FileSet.
	AddOutputDirectory(directory *string) FileSet
	// Add a dependency on another step.
	AddStepDependency(step Step)
	// Configure the given FileSet as the primary output of this step.
	ConfigurePrimaryOutput(fs FileSet)
	// Crawl the given structure for references to StepOutputs and add dependencies on all steps found.
	//
	// Should be called in the constructor of subclasses based on what the user
	// passes in as construction properties. The format of the structure passed in
	// here does not have to correspond exactly to what gets rendered into the
	// engine, it just needs to contain the same data.
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
	ExportedVariable(variableName *string) *string
	// Configure the given output directory as primary output.
	//
	// If no primary output has been configured yet, this directory
	// will become the primary output of this ShellStep, otherwise this
	// method will throw if the given directory is different than the
	// currently configured primary output directory.
	PrimaryOutputDirectory(directory *string) FileSet
	// Return a string representation of this Step.
	ToString() *string
}

// The jsii proxy struct for CodeBuildStep
type jsiiProxy_CodeBuildStep struct {
	jsiiProxy_ShellStep
}

func (j *jsiiProxy_CodeBuildStep) ActionRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"actionRole",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CodeBuildStep) Cache() awscodebuild.Cache {
	var returns awscodebuild.Cache
	_jsii_.Get(
		j,
		"cache",
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

