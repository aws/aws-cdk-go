package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

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
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func PipelineBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePipelineBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.PipelineBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object extends {@link PipelineBase}.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func PipelineBase_IsPipeline(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePipelineBase_IsPipelineParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.PipelineBase",
		"isPipeline",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineBase) AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment {
	if err := p.validateAddStageParameters(stage, options); err != nil {
		panic(err)
	}
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
	if err := p.validateAddWaveParameters(id, options); err != nil {
		panic(err)
	}
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

