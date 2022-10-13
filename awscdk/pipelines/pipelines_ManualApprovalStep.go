package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

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

	if err := validateNewManualApprovalStepParameters(id, props); err != nil {
		panic(err)
	}
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

	if err := validateManualApprovalStep_SequenceParameters(steps); err != nil {
		panic(err)
	}
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
	if err := m.validateAddDependencyFileSetParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (m *jsiiProxy_ManualApprovalStep) AddStepDependency(step Step) {
	if err := m.validateAddStepDependencyParameters(step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (m *jsiiProxy_ManualApprovalStep) ConfigurePrimaryOutput(fs FileSet) {
	if err := m.validateConfigurePrimaryOutputParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (m *jsiiProxy_ManualApprovalStep) DiscoverReferencedOutputs(structure interface{}) {
	if err := m.validateDiscoverReferencedOutputsParameters(structure); err != nil {
		panic(err)
	}
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

