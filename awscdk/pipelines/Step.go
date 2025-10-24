package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A generic Step which can be added to a Pipeline.
//
// Steps can be used to add Sources, Build Actions and Validations
// to your pipeline.
//
// This class is abstract. See specific subclasses of Step for
// useful steps to add to your Pipeline.
//
// Example:
//   type myJenkinsStep struct {
//   	Step
//   }
//
//   func newMyJenkinsStep(provider JenkinsProvider, input FileSet) *myJenkinsStep {
//   	this := &myJenkinsStep{}
//   	pipelines.NewStep_Override(this, jsii.String("MyJenkinsStep"))
//
//   	// This is necessary if your step accepts parameters, like environment variables,
//   	// that may contain outputs from other steps. It doesn't matter what the
//   	// structure is, as long as it contains the values that may contain outputs.
//   	this.DiscoverReferencedOutputs(map[string]map[string]interface{}{
//   		"env": map[string]interface{}{
//   		},
//   	})
//   	return this
//   }
//
//   func (this *myJenkinsStep) produceAction(stage IStage, options ProduceActionOptions) CodePipelineActionFactoryResult {
//   	// This is where you control what type of Action gets added to the
//   	// CodePipeline
//   	*stage.AddAction(
//   	cpactions.NewJenkinsAction(&JenkinsActionProps{
//   		// Copy 'actionName' and 'runOrder' from the options
//   		ActionName: options.ActionName,
//   		RunOrder: options.RunOrder,
//
//   		// Jenkins-specific configuration
//   		Type: cpactions.JenkinsActionType_TEST,
//   		JenkinsProvider: this.provider,
//   		ProjectName: jsii.String("MyJenkinsProject"),
//
//   		// Translate the FileSet into a codepipeline.Artifact
//   		Inputs: []Artifact{
//   			options.Artifacts.ToCodePipeline(this.input),
//   		},
//   	}))
//
//   	return &CodePipelineActionFactoryResult{
//   		RunOrdersConsumed: jsii.Number(1),
//   	}
//   }
//
type Step interface {
	IFileSetProducer
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
	// Return a string representation of this Step.
	ToString() *string
}

// The jsii proxy struct for Step
type jsiiProxy_Step struct {
	jsiiProxy_IFileSetProducer
}

func (j *jsiiProxy_Step) ConsumedStackOutputs() *[]StackOutputReference {
	var returns *[]StackOutputReference
	_jsii_.Get(
		j,
		"consumedStackOutputs",
		&returns,
	)
	return returns
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

	if err := validateStep_SequenceParameters(steps); err != nil {
		panic(err)
	}
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
	if err := s.validateAddDependencyFileSetParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addDependencyFileSet",
		[]interface{}{fs},
	)
}

func (s *jsiiProxy_Step) AddStepDependency(step Step) {
	if err := s.validateAddStepDependencyParameters(step); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addStepDependency",
		[]interface{}{step},
	)
}

func (s *jsiiProxy_Step) ConfigurePrimaryOutput(fs FileSet) {
	if err := s.validateConfigurePrimaryOutputParameters(fs); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"configurePrimaryOutput",
		[]interface{}{fs},
	)
}

func (s *jsiiProxy_Step) DiscoverReferencedOutputs(structure interface{}) {
	if err := s.validateDiscoverReferencedOutputsParameters(structure); err != nil {
		panic(err)
	}
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

