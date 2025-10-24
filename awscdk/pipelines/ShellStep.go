package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Run shell script commands in the pipeline.
//
// This is a generic step designed
// to be deployment engine agnostic.
//
// Example:
//   var pipeline CodePipeline
//
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//   topic := sns.NewTopic(this, jsii.String("ChangeApprovalTopic"))
//
//   pipeline.AddStage(preprod, &AddStageOpts{
//   	Post: []Step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &ShellStepProps{
//   			Commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.AddStage(prod, &AddStageOpts{
//   	Pre: []Step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd"), &ManualApprovalStepProps{
//   			//All options below are optional
//   			Comment: jsii.String("Please validate changes"),
//   			ReviewUrl: jsii.String("https://my.webservice.com/"),
//   			NotificationTopic: topic,
//   		}),
//   	},
//   })
//
type ShellStep interface {
	Step
	// Commands to run.
	Commands() *[]*string
	// StackOutputReferences this step consumes.
	ConsumedStackOutputs() *[]StackOutputReference
	// Return the steps this step depends on, based on the FileSets it requires.
	Dependencies() *[]Step
	// The list of FileSets consumed by this Step.
	DependencyFileSets() *[]FileSet
	// Environment variables to set.
	// Default: - No environment variables.
	//
	Env() *map[string]*string
	// Set environment variables based on Stack Outputs.
	// Default: - No environment variables created from stack outputs.
	//
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
	// Default: - No installation commands.
	//
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

func (j *jsiiProxy_ShellStep) ConsumedStackOutputs() *[]StackOutputReference {
	var returns *[]StackOutputReference
	_jsii_.Get(
		j,
		"consumedStackOutputs",
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

	if err := validateNewShellStepParameters(id, props); err != nil {
		panic(err)
	}
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

	if err := validateShellStep_SequenceParameters(steps); err != nil {
		panic(err)
	}
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
	if err := s.validateAddDependencyFileSetParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (s *jsiiProxy_ShellStep) AddOutputDirectory(directory *string) FileSet {
	if err := s.validateAddOutputDirectoryParameters(directory); err != nil {
		panic(err)
	}
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
	if err := s.validateAddStepDependencyParameters(step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (s *jsiiProxy_ShellStep) ConfigurePrimaryOutput(fs FileSet) {
	if err := s.validateConfigurePrimaryOutputParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (s *jsiiProxy_ShellStep) DiscoverReferencedOutputs(structure interface{}) {
	if err := s.validateDiscoverReferencedOutputsParameters(structure); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

func (s *jsiiProxy_ShellStep) PrimaryOutputDirectory(directory *string) FileSet {
	if err := s.validatePrimaryOutputDirectoryParameters(directory); err != nil {
		panic(err)
	}
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

