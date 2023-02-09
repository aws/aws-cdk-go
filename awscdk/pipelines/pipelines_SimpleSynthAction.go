package pipelines

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/pipelines/internal"
)

// A standard synth with a generated buildspec.
//
// Example:
//   sourceArtifact := codepipeline.NewArtifact()
//   cloudAssemblyArtifact := codepipeline.NewArtifact()
//   pipeline := pipelines.NewCdkPipeline(this, jsii.String("MyPipeline"), &cdkPipelineProps{
//   	cloudAssemblyArtifact: cloudAssemblyArtifact,
//   	synthAction: pipelines.simpleSynthAction.standardNpmSynth(&standardNpmSynthOptions{
//   		sourceArtifact: sourceArtifact,
//   		cloudAssemblyArtifact: cloudAssemblyArtifact,
//   		environment: &buildEnvironment{
//   			privileged: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
type SimpleSynthAction interface {
	awscodepipeline.IAction
	awsiam.IGrantable
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	ActionProperties() *awscodepipeline.ActionProperties
	// The CodeBuild Project's principal.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	GrantPrincipal() awsiam.IPrincipal
	// Project generated to run the synth command.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Project() awscodebuild.IProject
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Exists to implement IAction.
	// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for SimpleSynthAction
type jsiiProxy_SimpleSynthAction struct {
	internal.Type__awscodepipelineIAction
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_SimpleSynthAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SimpleSynthAction) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SimpleSynthAction) Project() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewSimpleSynthAction(props *SimpleSynthActionProps) SimpleSynthAction {
	_init_.Initialize()

	if err := validateNewSimpleSynthActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SimpleSynthAction{}

	_jsii_.Create(
		"monocdk.pipelines.SimpleSynthAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func NewSimpleSynthAction_Override(s SimpleSynthAction, props *SimpleSynthActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.pipelines.SimpleSynthAction",
		[]interface{}{props},
		s,
	)
}

// Create a standard NPM synth action.
//
// Uses `npm ci` to install dependencies and `npx cdk synth` to synthesize.
//
// If you need a build step, add `buildCommand: 'npm run build'`.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func SimpleSynthAction_StandardNpmSynth(options *StandardNpmSynthOptions) SimpleSynthAction {
	_init_.Initialize()

	if err := validateSimpleSynthAction_StandardNpmSynthParameters(options); err != nil {
		panic(err)
	}
	var returns SimpleSynthAction

	_jsii_.StaticInvoke(
		"monocdk.pipelines.SimpleSynthAction",
		"standardNpmSynth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Create a standard Yarn synth action.
//
// Uses `yarn install --frozen-lockfile` to install dependencies and `npx cdk synth` to synthesize.
//
// If you need a build step, add `buildCommand: 'yarn build'`.
// Deprecated: This class is part of the old API. Use the API based on the `CodePipeline` class instead
func SimpleSynthAction_StandardYarnSynth(options *StandardYarnSynthOptions) SimpleSynthAction {
	_init_.Initialize()

	if err := validateSimpleSynthAction_StandardYarnSynthParameters(options); err != nil {
		panic(err)
	}
	var returns SimpleSynthAction

	_jsii_.StaticInvoke(
		"monocdk.pipelines.SimpleSynthAction",
		"standardYarnSynth",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SimpleSynthAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := s.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SimpleSynthAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := s.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

