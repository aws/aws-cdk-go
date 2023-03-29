package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Pause the pipeline if a deployment would add IAM permissions or Security Group rules.
//
// This step is only supported in CodePipeline pipelines.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var pipeline codePipeline
//
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.AddStage(stage, &AddStageOpts{
//   	Pre: []step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &PermissionsBroadeningCheckProps{
//   			Stage: *Stage,
//   		}),
//   	},
//   })
//
type ConfirmPermissionsBroadening interface {
	Step
	ICodePipelineActionFactory
	// StackOutputReferences this step consumes.
	ConsumedStackOutputs() *[]StackOutputReference
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

func (j *jsiiProxy_ConfirmPermissionsBroadening) ConsumedStackOutputs() *[]StackOutputReference {
	var returns *[]StackOutputReference
	_jsii_.Get(
		j,
		"consumedStackOutputs",
		&returns,
	)
	return returns
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

	if err := validateNewConfirmPermissionsBroadeningParameters(id, props); err != nil {
		panic(err)
	}
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

	if err := validateConfirmPermissionsBroadening_SequenceParameters(steps); err != nil {
		panic(err)
	}
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
	if err := c.validateAddDependencyFileSetParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) AddStepDependency(step Step) {
	if err := c.validateAddStepDependencyParameters(step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) ConfigurePrimaryOutput(fs FileSet) {
	if err := c.validateConfigurePrimaryOutputParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) DiscoverReferencedOutputs(structure interface{}) {
	if err := c.validateDiscoverReferencedOutputsParameters(structure); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"discoverReferencedOutputs",
		[]interface{}{structure},
	)
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) ProduceAction(stage awscodepipeline.IStage, options *ProduceActionOptions) *CodePipelineActionFactoryResult {
	if err := c.validateProduceActionParameters(stage, options); err != nil {
		panic(err)
	}
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

