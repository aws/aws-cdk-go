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
// Example:
//   var pipeline codePipeline
//
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
type AddStageOpts struct {
	// Additional steps to run after all of the stacks in the stage.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
	// Instructions for stack level steps.
	StackSteps *[]*StackSteps `field:"optional" json:"stackSteps" yaml:"stackSteps"`
}

// Translate FileSets to CodePipeline Artifacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactMap := awscdk.Pipelines.NewArtifactMap()
//
type ArtifactMap interface {
	// Return the matching CodePipeline artifact for a FileSet.
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
	// A file.
	AssetType_FILE AssetType = "FILE"
	// A Docker image.
	AssetType_DOCKER_IMAGE AssetType = "DOCKER_IMAGE"
)

// Options for customizing a single CodeBuild project.
//
// Example:
//   var vpc vpc
//   var mySecurityGroup securityGroup
//
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
//   			subnetType: ec2.subnetType_PRIVATE_WITH_NAT,
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
type CodeBuildOptions struct {
	// Partial build environment, will be combined with other build environments that apply.
	BuildEnvironment *awscodebuild.BuildEnvironment `field:"optional" json:"buildEnvironment" yaml:"buildEnvironment"`
	// Partial buildspec, will be combined with other buildspecs that apply.
	//
	// The BuildSpec must be available inline--it cannot reference a file
	// on disk.
	PartialBuildSpec awscodebuild.BuildSpec `field:"optional" json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Policy statements to add to role.
	RolePolicy *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicy" yaml:"rolePolicy"`
	// Which security group(s) to associate with the project network interfaces.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VPC where to create the CodeBuild network interfaces in.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
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
type CodeBuildStep interface {
	ShellStep
	// Custom execution role to be used for the Code Build Action.
	ActionRole() awsiam.IRole
	// Build environment.
	BuildEnvironment() *awscodebuild.BuildEnvironment
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
type CodeBuildStepProps struct {
	// Commands to run.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
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
	AdditionalInputs *map[string]IFileSetProducer `field:"optional" json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `field:"optional" json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	Input IFileSetProducer `field:"optional" json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	InstallCommands *[]*string `field:"optional" json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	PrimaryOutputDirectory *string `field:"optional" json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
	// Custom execution role to be used for the Code Build Action.
	ActionRole awsiam.IRole `field:"optional" json:"actionRole" yaml:"actionRole"`
	// Changes to environment.
	//
	// This environment will be combined with the pipeline's default
	// environment.
	BuildEnvironment *awscodebuild.BuildEnvironment `field:"optional" json:"buildEnvironment" yaml:"buildEnvironment"`
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
	PartialBuildSpec awscodebuild.BuildSpec `field:"optional" json:"partialBuildSpec" yaml:"partialBuildSpec"`
	// Name for the generated CodeBuild project.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// Custom execution role to be used for the CodeBuild project.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Policy statements to add to role used during the synth.
	//
	// Can be used to add acces to a CodeArtifact repository etc.
	RolePolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"rolePolicyStatements" yaml:"rolePolicyStatements"`
	// Which security group to associate with the script's project network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Which subnets to use.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VPC where to execute the SimpleSynth.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

// Configuration options for a CodeCommit source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   codeCommitSourceOptions := &codeCommitSourceOptions{
//   	codeBuildCloneOutput: jsii.Boolean(false),
//   	eventRole: role,
//   	trigger: awscdk.Aws_codepipeline_actions.codeCommitTrigger_NONE,
//   }
//
type CodeCommitSourceOptions struct {
	// If this is set, the next CodeBuild job clones the repository (instead of CodePipeline downloading the files).
	//
	// This provides access to repository history, and retains symlinks (symlinks would otherwise be
	// removed by CodePipeline).
	//
	// **Note**: if this option is true, only CodeBuild jobs can use the output artifact.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodeCommit.html
	//
	CodeBuildCloneOutput *bool `field:"optional" json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Role to be used by on commit event rule.
	//
	// Used only when trigger value is CodeCommitTrigger.EVENTS.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// How should CodePipeline detect source changes for this Action.
	Trigger awscodepipelineactions.CodeCommitTrigger `field:"optional" json:"trigger" yaml:"trigger"`
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
type CodePipeline interface {
	PipelineBase
	// The FileSet tha contains the cloud assembly.
	//
	// This is the primary output of the synth step.
	CloudAssemblyFileSet() FileSet
	// The tree node.
	Node() constructs.Node
	// The CodePipeline pipeline that deploys the CDK app.
	//
	// Only available after the pipeline has been built.
	Pipeline() awscodepipeline.Pipeline
	// The build step that produces the CDK Cloud Assembly.
	Synth() IFileSetProducer
	// The CodeBuild project that performs the Synth.
	//
	// Only available after the pipeline has been built.
	SynthProject() awscodebuild.IProject
	// The waves in this pipeline.
	Waves() *[]Wave
	// Deploy a single Stage by itself.
	//
	// Add a Stage to the pipeline, to be deployed in sequence with other
	// Stages added to the pipeline. All Stacks in the stage will be deployed
	// in an order automatically determined by their relative dependencies.
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
	AddWave(id *string, options *WaveOptions) Wave
	// Send the current pipeline definition to the engine, and construct the pipeline.
	//
	// It is not possible to modify the pipeline after calling this method.
	BuildPipeline()
	// Implemented by subclasses to do the actual pipeline construction.
	DoBuildPipeline()
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var project project
//
//   codePipelineActionFactoryResult := &codePipelineActionFactoryResult{
//   	runOrdersConsumed: jsii.Number(123),
//
//   	// the properties below are optional
//   	project: project,
//   }
//
type CodePipelineActionFactoryResult struct {
	// How many RunOrders were consumed.
	//
	// If you add 1 action, return the value 1 here.
	RunOrdersConsumed *float64 `field:"required" json:"runOrdersConsumed" yaml:"runOrdersConsumed"`
	// If a CodeBuild project got created, the project.
	Project awscodebuild.IProject `field:"optional" json:"project" yaml:"project"`
}

// A FileSet created from a CodePipeline artifact.
//
// You only need to use this if you want to add CDK Pipeline stages
// add the end of an existing CodePipeline, which should be very rare.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//
//   codePipelineFileSet := awscdk.Pipelines.codePipelineFileSet.fromArtifact(artifact)
//
type CodePipelineFileSet interface {
	FileSet
	// Human-readable descriptor for this file set (does not need to be unique).
	Id() *string
	// The primary output of a file set producer.
	//
	// The primary output of a FileSet is itself.
	PrimaryOutput() FileSet
	// The Step that produces this FileSet.
	Producer() Step
	// Mark the given Step as the producer for this FileSet.
	//
	// This method can only be called once.
	ProducedBy(producer Step)
	// Return a string representation of this FileSet.
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
type CodePipelineProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	Synth IFileSetProducer `field:"required" json:"synth" yaml:"synth"`
	// Additional customizations to apply to the asset publishing CodeBuild projects.
	AssetPublishingCodeBuildDefaults *CodeBuildOptions `field:"optional" json:"assetPublishingCodeBuildDefaults" yaml:"assetPublishingCodeBuildDefaults"`
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
	CliVersion *string `field:"optional" json:"cliVersion" yaml:"cliVersion"`
	// Customize the CodeBuild projects created for this pipeline.
	CodeBuildDefaults *CodeBuildOptions `field:"optional" json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An existing Pipeline to be reused and built upon.
	//
	// [disable-awslint:ref-via-interface].
	CodePipeline awscodepipeline.Pipeline `field:"optional" json:"codePipeline" yaml:"codePipeline"`
	// Create KMS keys for the artifact buckets, allowing cross-account deployments.
	//
	// The artifact buckets have to be encrypted to support deploying CDK apps to
	// another account, so if you want to do that or want to have your artifact
	// buckets encrypted, be sure to set this value to `true`.
	//
	// Be aware there is a cost associated with maintaining the KMS keys.
	CrossAccountKeys *bool `field:"optional" json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
	DockerCredentials *[]DockerCredential `field:"optional" json:"dockerCredentials" yaml:"dockerCredentials"`
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
	DockerEnabledForSelfMutation *bool `field:"optional" json:"dockerEnabledForSelfMutation" yaml:"dockerEnabledForSelfMutation"`
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
	DockerEnabledForSynth *bool `field:"optional" json:"dockerEnabledForSynth" yaml:"dockerEnabledForSynth"`
	// The name of the CodePipeline pipeline.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Publish assets in multiple CodeBuild projects.
	//
	// If set to false, use one Project per type to publish all assets.
	//
	// Publishing in parallel improves concurrency and may reduce publishing
	// latency, but may also increase overall provisioning time of the CodeBuild
	// projects.
	//
	// Experiment and see what value works best for you.
	PublishAssetsInParallel *bool `field:"optional" json:"publishAssetsInParallel" yaml:"publishAssetsInParallel"`
	// Reuse the same cross region support stack for all pipelines in the App.
	ReuseCrossRegionSupportStacks *bool `field:"optional" json:"reuseCrossRegionSupportStacks" yaml:"reuseCrossRegionSupportStacks"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	SelfMutation *bool `field:"optional" json:"selfMutation" yaml:"selfMutation"`
	// Additional customizations to apply to the self mutation CodeBuild projects.
	SelfMutationCodeBuildDefaults *CodeBuildOptions `field:"optional" json:"selfMutationCodeBuildDefaults" yaml:"selfMutationCodeBuildDefaults"`
	// Additional customizations to apply to the synthesize CodeBuild projects.
	SynthCodeBuildDefaults *CodeBuildOptions `field:"optional" json:"synthCodeBuildDefaults" yaml:"synthCodeBuildDefaults"`
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
type CodePipelineSource interface {
	Step
	ICodePipelineActionFactory
	// Return the steps this step depends on, based on the FileSets it requires.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	AddDependencyFileSet(fs FileSet)
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
	GetAction(output awscodepipeline.Artifact, actionName *string, runOrder *float64, variablesNamespace *string) awscodepipelineactions.Action
	// Create the desired Action and add it to the pipeline.
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
	SourceAttribute(name *string) *string
	// Return a string representation of this Step.
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
// If you need access to symlinks or the repository history, be sure to set
// `codeBuildCloneOutput`.
//
// Example:
//   var repository iRepository
//
//   pipelines.codePipelineSource.codeCommit(repository, jsii.String("main"))
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
//
// If you need access to symlinks or the repository history, be sure to set
// `codeBuildCloneOutput`.
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
// Example:
//   var repository iRepository
//
//   pipelines.codePipelineSource.ecr(repository, &eCRSourceOptions{
//   	imageTag: jsii.String("latest"),
//   })
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
//
// If you need access to symlinks or the repository history, use a source of type
// `connection` instead.
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
// Example:
//   var bucket bucket
//
//   pipelines.codePipelineSource.s3(bucket, jsii.String("path/to/file.zip"))
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
//
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.addStage(stage, &addStageOpts{
//   	pre: []step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &permissionsBroadeningCheckProps{
//   			stage: stage,
//   		}),
//   	},
//   })
//
type ConfirmPermissionsBroadening interface {
	Step
	ICodePipelineActionFactory
	// Return the steps this step depends on, based on the FileSets it requires.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	AddDependencyFileSet(fs FileSet)
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
	// Create the desired Action and add it to the pipeline.
	ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult
	// Return a string representation of this Step.
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
type ConnectionSourceOptions struct {
	// The ARN of the CodeStar Connection created in the AWS console that has permissions to access this GitHub or BitBucket repository.
	//
	// Example:
	//   "arn:aws:codestar-connections:us-east-1:123456789012:connection/12345678-abcd-12ab-34cdef5678gh"
	//
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/connections-create.html
	//
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// If this is set, the next CodeBuild job clones the repository (instead of CodePipeline downloading the files).
	//
	// This provides access to repository history, and retains symlinks (symlinks would otherwise be
	// removed by CodePipeline).
	//
	// **Note**: if this option is true, only CodeBuild jobs can use the output artifact.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html#action-reference-CodestarConnectionSource-config
	//
	CodeBuildCloneOutput *bool `field:"optional" json:"codeBuildCloneOutput" yaml:"codeBuildCloneOutput"`
	// Controls automatically starting your pipeline when a new commit is made on the configured repository and branch.
	//
	// If unspecified,
	// the default value is true, and the field does not display by default.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference-CodestarConnectionSource.html
	//
	TriggerOnPush *bool `field:"optional" json:"triggerOnPush" yaml:"triggerOnPush"`
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
type DockerCredential interface {
	Usages() *[]DockerCredentialUsage
	// Grant read-only access to the registry credentials.
	//
	// This grants read access to any secrets, and pull access to any repositories.
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
type DockerCredentialUsage string

const (
	// Synth/Build.
	DockerCredentialUsage_SYNTH DockerCredentialUsage = "SYNTH"
	// Self-update.
	DockerCredentialUsage_SELF_UPDATE DockerCredentialUsage = "SELF_UPDATE"
	// Asset publishing.
	DockerCredentialUsage_ASSET_PUBLISHING DockerCredentialUsage = "ASSET_PUBLISHING"
)

// Options for ECR sources.
//
// Example:
//   var repository iRepository
//
//   pipelines.codePipelineSource.ecr(repository, &eCRSourceOptions{
//   	imageTag: jsii.String("latest"),
//   })
//
type ECRSourceOptions struct {
	// The action name used for this source in the CodePipeline.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// The image tag that will be checked for changes.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

// Options for defining access for a Docker Credential composed of ECR repos.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   ecrDockerCredentialOptions := &ecrDockerCredentialOptions{
//   	assumeRole: role,
//   	usages: []dockerCredentialUsage{
//   		awscdk.Pipelines.*dockerCredentialUsage_SYNTH,
//   	},
//   }
//
type EcrDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	Usages *[]DockerCredentialUsage `field:"optional" json:"usages" yaml:"usages"`
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
type ExternalDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// The name of the JSON field of the secret which contains the secret/password.
	SecretPasswordField *string `field:"optional" json:"secretPasswordField" yaml:"secretPasswordField"`
	// The name of the JSON field of the secret which contains the user/login name.
	SecretUsernameField *string `field:"optional" json:"secretUsernameField" yaml:"secretUsernameField"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	Usages *[]DockerCredentialUsage `field:"optional" json:"usages" yaml:"usages"`
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
type FileSet interface {
	IFileSetProducer
	// Human-readable descriptor for this file set (does not need to be unique).
	Id() *string
	// The primary output of a file set producer.
	//
	// The primary output of a FileSet is itself.
	PrimaryOutput() FileSet
	// The Step that produces this FileSet.
	Producer() Step
	// Mark the given Step as the producer for this FileSet.
	//
	// This method can only be called once.
	ProducedBy(producer Step)
	// Return a string representation of this FileSet.
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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSet fileSet
//
//   fileSetLocation := &fileSetLocation{
//   	directory: jsii.String("directory"),
//   	fileSet: fileSet,
//   }
//
type FileSetLocation struct {
	// The (relative) directory where the FileSet is found.
	Directory *string `field:"required" json:"directory" yaml:"directory"`
	// The FileSet object.
	FileSet FileSet `field:"required" json:"fileSet" yaml:"fileSet"`
}

// Options for GitHub sources.
//
// Example:
//   pipelines.codePipelineSource.gitHub(jsii.String("org/repo"), jsii.String("branch"), &gitHubSourceOptions{
//   	// This is optional
//   	authentication: cdk.secretValue.secretsManager(jsii.String("my-token")),
//   })
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
	// * **admin:repo_hook** - if you plan to use webhooks (true by default).
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/GitHub-create-personal-token-CLI.html
	//
	Authentication awscdk.SecretValue `field:"optional" json:"authentication" yaml:"authentication"`
	// How AWS CodePipeline should be triggered.
	//
	// With the default value "WEBHOOK", a webhook is created in GitHub that triggers the action.
	// With "POLL", CodePipeline periodically checks the source for changes.
	// With "None", the action is not triggered through changes in the source.
	//
	// To use `WEBHOOK`, your GitHub Personal Access Token should have
	// **admin:repo_hook** scope (in addition to the regular **repo** scope).
	Trigger awscodepipelineactions.GitHubTrigger `field:"optional" json:"trigger" yaml:"trigger"`
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
// Example:
//   var pipeline codePipeline
//
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
type ManualApprovalStep interface {
	Step
	// The comment associated with this manual approval.
	Comment() *string
	// Return the steps this step depends on, based on the FileSets it requires.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	AddDependencyFileSet(fs FileSet)
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
	// Return a string representation of this Step.
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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manualApprovalStepProps := &manualApprovalStepProps{
//   	comment: jsii.String("comment"),
//   }
//
type ManualApprovalStepProps struct {
	// The comment to display with this manual approval.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

// Properties for a `PermissionsBroadeningCheck`.
//
// Example:
//   var pipeline codePipeline
//
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.addStage(stage, &addStageOpts{
//   	pre: []step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &permissionsBroadeningCheckProps{
//   			stage: stage,
//   		}),
//   	},
//   })
//
type PermissionsBroadeningCheckProps struct {
	// The CDK Stage object to check the stacks of.
	//
	// This should be the same Stage object you are passing to `addStage()`.
	Stage awscdk.Stage `field:"required" json:"stage" yaml:"stage"`
	// Topic to send notifications when a human needs to give manual confirmation.
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
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
	// The FileSet tha contains the cloud assembly.
	//
	// This is the primary output of the synth step.
	CloudAssemblyFileSet() FileSet
	// The tree node.
	Node() constructs.Node
	// The build step that produces the CDK Cloud Assembly.
	Synth() IFileSetProducer
	// The waves in this pipeline.
	Waves() *[]Wave
	// Deploy a single Stage by itself.
	//
	// Add a Stage to the pipeline, to be deployed in sequence with other
	// Stages added to the pipeline. All Stacks in the stage will be deployed
	// in an order automatically determined by their relative dependencies.
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
	AddWave(id *string, options *WaveOptions) Wave
	// Send the current pipeline definition to the engine, and construct the pipeline.
	//
	// It is not possible to modify the pipeline after calling this method.
	BuildPipeline()
	// Implemented by subclasses to do the actual pipeline construction.
	DoBuildPipeline()
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSetProducer iFileSetProducer
//
//   pipelineBaseProps := &pipelineBaseProps{
//   	synth: fileSetProducer,
//   }
//
type PipelineBaseProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	Synth IFileSetProducer `field:"required" json:"synth" yaml:"synth"`
}

// Options for the `CodePipelineActionFactory.produce()` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var artifact artifact
//   var artifactMap artifactMap
//   var bucket bucket
//   var buildImage iBuildImage
//   var buildSpec buildSpec
//   var codePipeline codePipeline
//   var construct construct
//   var policyStatement policyStatement
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var value interface{}
//   var vpc vpc
//
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
//   			computeType: awscdk.Aws_codebuild.computeType_SMALL,
//   			environmentVariables: map[string]buildEnvironmentVariable{
//   				"environmentVariablesKey": &buildEnvironmentVariable{
//   					"value": value,
//
//   					// the properties below are optional
//   					"type": awscdk.*Aws_codebuild.BuildEnvironmentVariableType_PLAINTEXT,
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
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   		},
//   		timeout: cdk.duration.minutes(jsii.Number(30)),
//   		vpc: vpc,
//   	},
//   	fallbackArtifact: artifact,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   }
//
type ProduceActionOptions struct {
	// Name the action should get.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Helper object to translate FileSets to CodePipeline Artifacts.
	Artifacts ArtifactMap `field:"required" json:"artifacts" yaml:"artifacts"`
	// The pipeline the action is being generated for.
	Pipeline CodePipeline `field:"required" json:"pipeline" yaml:"pipeline"`
	// RunOrder the action should get.
	RunOrder *float64 `field:"required" json:"runOrder" yaml:"runOrder"`
	// Scope in which to create constructs.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
	// Whether or not this action is inserted before self mutation.
	//
	// If it is, the action should take care to reflect some part of
	// its own definition in the pipeline action definition, to
	// trigger a restart after self-mutation (if necessary).
	BeforeSelfMutation *bool `field:"optional" json:"beforeSelfMutation" yaml:"beforeSelfMutation"`
	// If this action factory creates a CodeBuild step, default options to inherit.
	CodeBuildDefaults *CodeBuildOptions `field:"optional" json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An input artifact that CodeBuild projects that don't actually need an input artifact can use.
	//
	// CodeBuild Projects MUST have an input artifact in order to be added to the Pipeline. If
	// the Project doesn't actually care about its input (it can be anything), it can use the
	// Artifact passed here.
	FallbackArtifact awscodepipeline.Artifact `field:"optional" json:"fallbackArtifact" yaml:"fallbackArtifact"`
	// If this step is producing outputs, the variables namespace assigned to it.
	//
	// Pass this on to the Action you are creating.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
}

// Options for S3 sources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourceOptions := &s3SourceOptions{
//   	actionName: jsii.String("actionName"),
//   	trigger: awscdk.Aws_codepipeline_actions.s3Trigger_NONE,
//   }
//
type S3SourceOptions struct {
	// The action name used for this source in the CodePipeline.
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// How should CodePipeline detect source changes for this Action.
	//
	// Note that if this is S3Trigger.EVENTS, you need to make sure to include the source Bucket in a CloudTrail Trail,
	// as otherwise the CloudWatch Events will not be emitted.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/log-s3-data-events.html
	//
	Trigger awscodepipelineactions.S3Trigger `field:"optional" json:"trigger" yaml:"trigger"`
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
type ShellStep interface {
	Step
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
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	PrimaryOutput() FileSet
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
type ShellStepProps struct {
	// Commands to run.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
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
	AdditionalInputs *map[string]IFileSetProducer `field:"optional" json:"additionalInputs" yaml:"additionalInputs"`
	// Environment variables to set.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
	// Set environment variables based on Stack Outputs.
	//
	// `ShellStep`s following stack or stage deployments may
	// access the `CfnOutput`s of those stacks to get access to
	// --for example--automatically generated resource names or
	// endpoint URLs.
	EnvFromCfnOutputs *map[string]awscdk.CfnOutput `field:"optional" json:"envFromCfnOutputs" yaml:"envFromCfnOutputs"`
	// FileSet to run these scripts on.
	//
	// The files in the FileSet will be placed in the working directory when
	// the script is executed. Use `additionalInputs` to download file sets
	// to other directories as well.
	Input IFileSetProducer `field:"optional" json:"input" yaml:"input"`
	// Installation commands to run before the regular commands.
	//
	// For deployment engines that support it, install commands will be classified
	// differently in the job history from the regular `commands`.
	InstallCommands *[]*string `field:"optional" json:"installCommands" yaml:"installCommands"`
	// The directory that will contain the primary output fileset.
	//
	// After running the script, the contents of the given directory
	// will be treated as the primary output of this Step.
	PrimaryOutputDirectory *string `field:"optional" json:"primaryOutputDirectory" yaml:"primaryOutputDirectory"`
}

// An asset used by a Stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackAsset := &stackAsset{
//   	assetId: jsii.String("assetId"),
//   	assetManifestPath: jsii.String("assetManifestPath"),
//   	assetSelector: jsii.String("assetSelector"),
//   	assetType: awscdk.Pipelines.assetType_FILE,
//   	isTemplate: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   }
//
type StackAsset struct {
	// Asset identifier.
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
	// Absolute asset manifest path.
	//
	// This needs to be made relative at a later point in time, but when this
	// information is parsed we don't know about the root cloud assembly yet.
	AssetManifestPath *string `field:"required" json:"assetManifestPath" yaml:"assetManifestPath"`
	// Asset selector to pass to `cdk-assets`.
	AssetSelector *string `field:"required" json:"assetSelector" yaml:"assetSelector"`
	// Type of asset to publish.
	AssetType AssetType `field:"required" json:"assetType" yaml:"assetType"`
	// Does this asset represent the CloudFormation template for the stack.
	IsTemplate *bool `field:"required" json:"isTemplate" yaml:"isTemplate"`
	// Role ARN to assume to publish.
	AssetPublishingRoleArn *string `field:"optional" json:"assetPublishingRoleArn" yaml:"assetPublishingRoleArn"`
}

// Deployment of a single Stack.
//
// You don't need to instantiate this class -- it will
// be automatically instantiated as necessary when you
// add a `Stage` to a pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudFormationStackArtifact cloudFormationStackArtifact
//
//   stackDeployment := awscdk.Pipelines.stackDeployment.fromArtifact(cloudFormationStackArtifact)
//
type StackDeployment interface {
	// Template path on disk to CloudAssembly.
	AbsoluteTemplatePath() *string
	// Account where the stack should be deployed.
	Account() *string
	// Assets referenced by this stack.
	Assets() *[]*StackAsset
	// Role to assume before deploying this stack.
	AssumeRoleArn() *string
	// Steps that take place after stack is prepared but before stack deploys.
	//
	// Your pipeline engine may not disable `prepareStep`.
	ChangeSet() *[]Step
	// Construct path for this stack.
	ConstructPath() *string
	// Execution role to pass to CloudFormation.
	ExecutionRoleArn() *string
	// Steps to execute after stack deploys.
	Post() *[]Step
	// Steps that take place before stack is prepared.
	//
	// If your pipeline engine disables 'prepareStep', then this will happen before stack deploys.
	Pre() *[]Step
	// Region where the stack should be deployed.
	Region() *string
	// Artifact ID for this stack.
	StackArtifactId() *string
	// Other stacks this stack depends on.
	StackDependencies() *[]StackDeployment
	// Name for this stack.
	StackName() *string
	// Tags to apply to the stack.
	Tags() *map[string]*string
	// The asset that represents the CloudFormation template for this stack.
	TemplateAsset() *StackAsset
	// The S3 URL which points to the template asset location in the publishing bucket.
	//
	// This is `undefined` if the stack template is not published. Use the
	// `DefaultStackSynthesizer` to ensure it is.
	//
	// Example value: `https://bucket.s3.amazonaws.com/object/key`
	TemplateUrl() *string
	// Add a dependency on another stack.
	AddStackDependency(stackDeployment StackDeployment)
	// Adds steps to each phase of the stack.
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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
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
//   			assetType: awscdk.Pipelines.assetType_FILE,
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
type StackDeploymentProps struct {
	// Template path on disk to cloud assembly (cdk.out).
	AbsoluteTemplatePath *string `field:"required" json:"absoluteTemplatePath" yaml:"absoluteTemplatePath"`
	// Construct path for this stack.
	ConstructPath *string `field:"required" json:"constructPath" yaml:"constructPath"`
	// Artifact ID for this stack.
	StackArtifactId *string `field:"required" json:"stackArtifactId" yaml:"stackArtifactId"`
	// Name for this stack.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Account where the stack should be deployed.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Assets referenced by this stack.
	Assets *[]*StackAsset `field:"optional" json:"assets" yaml:"assets"`
	// Role to assume before deploying this stack.
	AssumeRoleArn *string `field:"optional" json:"assumeRoleArn" yaml:"assumeRoleArn"`
	// Execution role to pass to CloudFormation.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Region where the stack should be deployed.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Tags to apply to the stack.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The S3 URL which points to the template asset location in the publishing bucket.
	TemplateS3Uri *string `field:"optional" json:"templateS3Uri" yaml:"templateS3Uri"`
}

// A Reference to a Stack Output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnOutput cfnOutput
//
//   stackOutputReference := awscdk.Pipelines.stackOutputReference.fromCfnOutput(cfnOutput)
//
type StackOutputReference interface {
	// Output name of the producing stack.
	OutputName() *string
	// A human-readable description of the producing stack.
	StackDescription() *string
	// Whether or not this stack output is being produced by the given Stack deployment.
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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//   var step step
//
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
type StackSteps struct {
	// The stack you want the steps to run in.
	Stack awscdk.Stack `field:"required" json:"stack" yaml:"stack"`
	// Steps that execute after stack is prepared but before stack is deployed.
	ChangeSet *[]Step `field:"optional" json:"changeSet" yaml:"changeSet"`
	// Steps that execute after stack is deployed.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Steps that execute before stack is prepared.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
}

// Deployment of a single `Stage`.
//
// A `Stage` consists of one or more `Stacks`, which will be
// deployed in dependency order.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//   var stage stage
//   var step step
//
//   stageDeployment := awscdk.Pipelines.stageDeployment.fromStage(stage, &stageDeploymentProps{
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
type StageDeployment interface {
	// Additional steps that are run after all of the stacks in the stage.
	Post() *[]Step
	// Additional steps that are run before any of the stacks in the stage.
	Pre() *[]Step
	// The stacks deployed in this stage.
	Stacks() *[]StackDeployment
	// Instructions for additional steps that are run at stack level.
	StackSteps() *[]*StackSteps
	// The display name of this stage.
	StageName() *string
	// Add an additional step to run after all of the stacks in this stage.
	AddPost(steps ...Step)
	// Add an additional step to run before any of the stacks in this stage.
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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//   var step step
//
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
type StageDeploymentProps struct {
	// Additional steps to run after all of the stacks in the stage.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stacks in the stage.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
	// Instructions for additional steps that are run at the stack level.
	StackSteps *[]*StackSteps `field:"optional" json:"stackSteps" yaml:"stackSteps"`
	// Stage name to use in the pipeline.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
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
type Step interface {
	IFileSetProducer
	// Return the steps this step depends on, based on the FileSets it requires.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	DependencyFileSets() *[]FileSet
	// Identifier for this step.
	Id() *string
	// Whether or not this is a Source step.
	//
	// What it means to be a Source step depends on the engine.
	IsSource() *bool
	// The primary FileSet produced by this Step.
	//
	// Not all steps produce an output FileSet--if they do
	// you can substitute the `Step` object for the `FileSet` object.
	PrimaryOutput() FileSet
	// Add an additional FileSet to the set of file sets required by this step.
	//
	// This will lead to a dependency on the producer of that file set.
	AddDependencyFileSet(fs FileSet)
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
	// Return a string representation of this Step.
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

// Multiple stages that are deployed in parallel.
//
// Example:
//   var pipeline codePipeline
//
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
type Wave interface {
	// Identifier for this Wave.
	Id() *string
	// Additional steps that are run after all of the stages in the wave.
	Post() *[]Step
	// Additional steps that are run before any of the stages in the wave.
	Pre() *[]Step
	// The stages that are deployed in this wave.
	Stages() *[]StageDeployment
	// Add an additional step to run after all of the stages in this wave.
	AddPost(steps ...Step)
	// Add an additional step to run before any of the stages in this wave.
	AddPre(steps ...Step)
	// Add a Stage to this wave.
	//
	// It will be deployed in parallel with all other stages in this
	// wave.
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
type WaveOptions struct {
	// Additional steps to run after all of the stages in the wave.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
}

// Construction properties for a `Wave`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var step step
//
//   waveProps := &waveProps{
//   	post: []*step{
//   		step,
//   	},
//   	pre: []*step{
//   		step,
//   	},
//   }
//
type WaveProps struct {
	// Additional steps to run after all of the stages in the wave.
	Post *[]Step `field:"optional" json:"post" yaml:"post"`
	// Additional steps to run before any of the stages in the wave.
	Pre *[]Step `field:"optional" json:"pre" yaml:"pre"`
}

