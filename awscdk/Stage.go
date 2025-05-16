package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//
//   preprod := NewMyApplicationStage(this, jsii.String("PreProd"))
//   prod := NewMyApplicationStage(this, jsii.String("Prod"))
//
//   pipeline.AddStage(preprod, &AddStageOpts{
//   	Post: []step{
//   		pipelines.NewShellStep(jsii.String("Validate Endpoint"), &ShellStepProps{
//   			Commands: []*string{
//   				jsii.String("curl -Ssf https://my.webservice.com/"),
//   			},
//   		}),
//   	},
//   })
//   pipeline.AddStage(prod, &AddStageOpts{
//   	Pre: []*step{
//   		pipelines.NewManualApprovalStep(jsii.String("PromoteToProd")),
//   	},
//   })
//
type Stage interface {
	constructs.Construct
	// The default account for all resources defined within this stage.
	Account() *string
	// Artifact ID of the assembly if it is a nested stage. The root stage (app) will return an empty string.
	//
	// Derived from the construct path.
	ArtifactId() *string
	// The cloud assembly asset output directory.
	AssetOutdir() *string
	// The tree node.
	Node() constructs.Node
	// The cloud assembly output directory.
	Outdir() *string
	// The parent stage or `undefined` if this is the app.
	//
	// *.
	ParentStage() Stage
	// Validation plugins to run during synthesis.
	//
	// If any plugin reports any violation,
	// synthesis will be interrupted and the report displayed to the user.
	// Default: - no validation plugins are used.
	//
	PolicyValidationBeta1() *[]IPolicyValidationPluginBeta1
	// The default region for all resources defined within this stage.
	Region() *string
	// The name of the stage.
	//
	// Based on names of the parent stages separated by
	// hypens.
	StageName() *string
	// Synthesize this stage into a cloud assembly.
	//
	// Once an assembly has been synthesized, it cannot be modified. Subsequent
	// calls will return the same assembly.
	Synth(options *StageSynthesisOptions) cxapi.CloudAssembly
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Stage
type jsiiProxy_Stage struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_Stage) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_Stage) PolicyValidationBeta1() *[]IPolicyValidationPluginBeta1 {
	var returns *[]IPolicyValidationPluginBeta1
	_jsii_.Get(
		j,
		"policyValidationBeta1",
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


func NewStage(scope constructs.Construct, id *string, props *StageProps) Stage {
	_init_.Initialize()

	if err := validateNewStageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Stage{}

	_jsii_.Create(
		"aws-cdk-lib.Stage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStage_Override(s Stage, scope constructs.Construct, id *string, props *StageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.Stage",
		[]interface{}{scope, id, props},
		s,
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
func Stage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Test whether the given construct is a stage.
func Stage_IsStage(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStage_IsStageParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stage",
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
func Stage_Of(construct constructs.IConstruct) Stage {
	_init_.Initialize()

	if err := validateStage_OfParameters(construct); err != nil {
		panic(err)
	}
	var returns Stage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.Stage",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
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

