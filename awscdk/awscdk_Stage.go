// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/cxapi"
	"github.com/aws/constructs-go/constructs/v3"
)

// An abstract application modeling unit consisting of Stacks that should be deployed together.
//
// Derive a subclass of `Stage` and use it to model a single instance of your
// application.
//
// You can then instantiate your subclass multiple times to model multiple
// copies of your application which should be be deployed to different
// environments.
//
// Example:
//   var pipeline codePipeline
//   type myOutputStage struct {
//   	stage
//   	loadBalancerAddress cfnOutput
//   }
//
//   func newMyOutputStage(scope construct, id *string, props stageProps) *myOutputStage {
//   	this := &myOutputStage{}
//   	newStage_Override(this, scope, id, props)
//   	this.loadBalancerAddress = awscdk.NewCfnOutput(this, jsii.String("Output"), &CfnOutputProps{
//   		Value: jsii.String("value"),
//   	})
//   	return this
//   }
//
//   lbApp := NewMyOutputStage(this, jsii.String("MyApp"))
//   pipeline.AddStage(lbApp, &AddStageOpts{
//   	Post: []step{
//   		pipelines.NewShellStep(jsii.String("HitEndpoint"), &ShellStepProps{
//   			EnvFromCfnOutputs: map[string]*cfnOutput{
//   				// Make the load balancer address available as $URL inside the commands
//   				"URL": lbApp.loadBalancerAddress,
//   			},
//   			Commands: []*string{
//   				jsii.String("curl -Ssf $URL"),
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type Stage interface {
	Construct
	// The default account for all resources defined within this stage.
	// Experimental.
	Account() *string
	// Artifact ID of the assembly if it is a nested stage. The root stage (app) will return an empty string.
	//
	// Derived from the construct path.
	// Experimental.
	ArtifactId() *string
	// The cloud assembly asset output directory.
	// Experimental.
	AssetOutdir() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() ConstructNode
	// The cloud assembly output directory.
	// Experimental.
	Outdir() *string
	// The parent stage or `undefined` if this is the app.
	//
	// *.
	// Experimental.
	ParentStage() Stage
	// The default region for all resources defined within this stage.
	// Experimental.
	Region() *string
	// The name of the stage.
	//
	// Based on names of the parent stages separated by
	// hypens.
	// Experimental.
	StageName() *string
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
	// Synthesize this stage into a cloud assembly.
	//
	// Once an assembly has been synthesized, it cannot be modified. Subsequent
	// calls will return the same assembly.
	// Experimental.
	Synth(options *StageSynthesisOptions) cxapi.CloudAssembly
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session ISynthesisSession)
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

// The jsii proxy struct for Stage
type jsiiProxy_Stage struct {
	jsiiProxy_Construct
}

func (j *jsiiProxy_Stage) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) AssetOutdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetOutdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Node() ConstructNode {
	var returns ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) ParentStage() Stage {
	var returns Stage
	_jsii_.Get(
		j,
		"parentStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Stage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}


// Experimental.
func NewStage(scope constructs.Construct, id *string, props *StageProps) Stage {
	_init_.Initialize()

	if err := validateNewStageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Stage{}

	_jsii_.Create(
		"monocdk.Stage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStage_Override(s Stage, scope constructs.Construct, id *string, props *StageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.Stage",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Stage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.Stage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a stage.
// Experimental.
func Stage_IsStage(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStage_IsStageParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.Stage",
		"isStage",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return the stage this construct is contained with, if available.
//
// If called
// on a nested stage, returns its parent.
// Experimental.
func Stage_Of(construct constructs.IConstruct) Stage {
	_init_.Initialize()

	if err := validateStage_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns Stage

	_jsii_.StaticInvoke(
		"monocdk.Stage",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stage) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Stage) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Stage) Synth(options *StageSynthesisOptions) cxapi.CloudAssembly {
	if err := s.validateSynthParameters(options); err != nil {
		panic(err)
	}
	var returns cxapi.CloudAssembly

	_jsii_.Invoke(
		s,
		"synth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) Synthesize(session ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_Stage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Stage) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

