package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline action to deploy an AWS ElasticBeanstalk Application.
//
// Example:
//   sourceOutput := codepipeline.NewArtifact()
//   targetBucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   deployAction := codepipeline_actions.NewElasticBeanstalkDeployAction(&ElasticBeanstalkDeployActionProps{
//   	ActionName: jsii.String("ElasticBeanstalkDeploy"),
//   	Input: sourceOutput,
//   	EnvironmentName: jsii.String("envName"),
//   	ApplicationName: jsii.String("appName"),
//   })
//
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type ElasticBeanstalkDeployAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for ElasticBeanstalkDeployAction
type jsiiProxy_ElasticBeanstalkDeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_ElasticBeanstalkDeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticBeanstalkDeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewElasticBeanstalkDeployAction(props *ElasticBeanstalkDeployActionProps) ElasticBeanstalkDeployAction {
	_init_.Initialize()

	if err := validateNewElasticBeanstalkDeployActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticBeanstalkDeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ElasticBeanstalkDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewElasticBeanstalkDeployAction_Override(e ElasticBeanstalkDeployAction, props *ElasticBeanstalkDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ElasticBeanstalkDeployAction",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_ElasticBeanstalkDeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := e.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticBeanstalkDeployAction) Bound(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := e.validateBoundParameters(_scope, _stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bound",
		[]interface{}{_scope, _stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticBeanstalkDeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := e.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		e,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticBeanstalkDeployAction) VariableExpression(variableName *string) *string {
	if err := e.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

