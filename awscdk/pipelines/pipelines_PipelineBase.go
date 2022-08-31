package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v3"
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

