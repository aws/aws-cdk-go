package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CDK Pipeline that uses CodePipeline to deploy CDK apps.
//
// This is a `Pipeline` with its `engine` property set to
// `CodePipelineEngine`, and exists for nicer ergonomics for
// users that don't need to switch out engines.
//
// Example:
//   var codePipeline Pipeline
//
//
//   sourceArtifact := codepipeline.NewArtifact(jsii.String("MySourceArtifact"))
//
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
//   	CodePipeline: codePipeline,
//   	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
//   		Input: pipelines.CodePipelineFileSet_FromArtifact(sourceArtifact),
//   		Commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
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
	// Whether SelfMutation is enabled for this CDK Pipeline.
	SelfMutationEnabled() *bool
	// The CodeBuild project that performs the SelfMutation.
	//
	// Will throw an error if this is accessed before `buildPipeline()`
	// is called, or if selfMutation has been disabled.
	SelfMutationProject() awscodebuild.IProject
	// The build step that produces the CDK Cloud Assembly.
	Synth() IFileSetProducer
	// The CodeBuild project that performs the Synth.
	//
	// Only available after the pipeline has been built.
	SynthProject() awscodebuild.IProject
	// Allow pipeline service role used for actions if no action role configured instead of creating a new role for each action.
	UsePipelineRoleForActions() *bool
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

func (j *jsiiProxy_CodePipeline) SelfMutationEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"selfMutationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodePipeline) SelfMutationProject() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"selfMutationProject",
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

func (j *jsiiProxy_CodePipeline) UsePipelineRoleForActions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"usePipelineRoleForActions",
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

	if err := validateNewCodePipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
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
func CodePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object extends `PipelineBase`.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func CodePipeline_IsPipeline(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodePipeline_IsPipelineParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.pipelines.CodePipeline",
		"isPipeline",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodePipeline) AddStage(stage awscdk.Stage, options *AddStageOpts) StageDeployment {
	if err := c.validateAddStageParameters(stage, options); err != nil {
		panic(err)
	}
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
	if err := c.validateAddWaveParameters(id, options); err != nil {
		panic(err)
	}
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

