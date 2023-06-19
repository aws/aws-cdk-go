package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A Pipeline to deploy CDK apps.
//
// Defines an AWS CodePipeline-based Pipeline to deploy CDK applications.
//
// Automatically manages the following:
//
// - Stack dependency order.
// - Asset publishing.
// - Keeping the pipeline up-to-date as the CDK apps change.
// - Using stack outputs later on in the pipeline.
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &CdkPipelineProps{
//   	CloudAssemblyArtifact: CloudAssemblyArtifact,
//   	SynthAction: pipelines.SimpleSynthAction_StandardNpmSynth(&StandardNpmSynthOptions{
//   		SourceArtifact: *SourceArtifact,
//   		CloudAssemblyArtifact: *CloudAssemblyArtifact,
//   		Environment: &BuildEnvironment{
//   			Privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type CdkPipeline interface {
	awscdk.Construct
	// The underlying CodePipeline object.
	//
	// You can use this to add more Stages to the pipeline, or Actions
	// to Stages.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	CodePipeline() awscodepipeline.Pipeline
	// The construct tree node associated with this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Node() awscdk.ConstructNode
	// Add pipeline stage that will deploy the given application stage.
	//
	// The application construct should subclass `Stage` and can contain any
	// number of `Stacks` inside it that may have dependency relationships
	// on one another.
	//
	// All stacks in the application will be deployed in the appropriate order,
	// and all assets found in the application will be added to the asset
	// publishing stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddApplicationStage(appStage awscdk.Stage, options *AddStageOptions) CdkStage
	// Add a new, empty stage to the pipeline.
	//
	// Prefer to use `addApplicationStage` if you are intended to deploy a CDK
	// application, but you can use this method if you want to add other kinds of
	// Actions to a pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	AddStage(stageName *string, options *BaseStageOptions) CdkStage
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Prepare()
	// Get the StackOutput object that holds this CfnOutput's value in this pipeline.
	//
	// `StackOutput` can be used in validation actions later in the pipeline.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	StackOutput(cfnOutput awscdk.CfnOutput) StackOutput
	// Access one of the pipeline's stages by stage name.
	//
	// You can use this to add more Actions to a stage.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Stage(stageName *string) awscodepipeline.IStage
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ToString() *string
	// Validate that we don't have any stacks violating dependency order in the pipeline.
	//
	// Our own convenience methods will never generate a pipeline that does that (although
	// this is a nice verification), but a user can also add the stacks by hand.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Validate() *[]*string
}

// The jsii proxy struct for CdkPipeline
type jsiiProxy_CdkPipeline struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_CdkPipeline) CodePipeline() awscodepipeline.Pipeline {
	var returns awscodepipeline.Pipeline
	_jsii_.Get(
		j,
		"codePipeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdkPipeline) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewCdkPipeline(scope constructs.Construct, id *string, props *CdkPipelineProps) CdkPipeline {
	_init_.Initialize()

	if err := validateNewCdkPipelineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdkPipeline{}

	_jsii_.Create(
		"monocdk.pipelines.CdkPipeline",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewCdkPipeline_Override(c CdkPipeline, scope constructs.Construct, id *string, props *CdkPipelineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.CdkPipeline",
		[]interface{}{scope, id, props},
		c,
	)
}

// Return whether the given object is a Construct.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func CdkPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCdkPipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.pipelines.CdkPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkPipeline) AddApplicationStage(appStage awscdk.Stage, options *AddStageOptions) CdkStage {
	if err := c.validateAddApplicationStageParameters(appStage, options); err != nil {
		panic(err)
	}
	var returns CdkStage

	_jsii_.Invoke(
		c,
		"addApplicationStage",
		[]interface{}{appStage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkPipeline) AddStage(stageName *string, options *BaseStageOptions) CdkStage {
	if err := c.validateAddStageParameters(stageName, options); err != nil {
		panic(err)
	}
	var returns CdkStage

	_jsii_.Invoke(
		c,
		"addStage",
		[]interface{}{stageName, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkPipeline) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkPipeline) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CdkPipeline) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkPipeline) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdkPipeline) StackOutput(cfnOutput awscdk.CfnOutput) StackOutput {
	if err := c.validateStackOutputParameters(cfnOutput); err != nil {
		panic(err)
	}
	var returns StackOutput

	_jsii_.Invoke(
		c,
		"stackOutput",
		[]interface{}{cfnOutput},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkPipeline) Stage(stageName *string) awscodepipeline.IStage {
	if err := c.validateStageParameters(stageName); err != nil {
		panic(err)
	}
	var returns awscodepipeline.IStage

	_jsii_.Invoke(
		c,
		"stage",
		[]interface{}{stageName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkPipeline) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CdkPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdkPipeline) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

