package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline Action to deploy EC2 instances.
//
// Example:
//   sourceOutput := codepipeline.NewArtifact()
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
//   	PipelineType: codepipeline.PipelineType_V2,
//   })
//   deployAction := codepipeline_actions.NewEc2DeployAction(&Ec2DeployActionProps{
//   	ActionName: jsii.String("Ec2Deploy"),
//   	Input: sourceOutput,
//   	InstanceType: codepipeline_actions.Ec2InstanceType_EC2,
//   	InstanceTagKey: jsii.String("Name"),
//   	InstanceTagValue: jsii.String("MyInstance"),
//   	DeploySpecifications: codepipeline_actions.Ec2DeploySpecifications_Inline(&Ec2DeploySpecificationsInlineProps{
//   		TargetDirectory: jsii.String("/home/ec2-user/deploy"),
//   		PreScript: jsii.String("scripts/pre-deploy.sh"),
//   		PostScript: jsii.String("scripts/post-deploy.sh"),
//   	}),
//   })
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []iAction{
//   		deployAction,
//   	},
//   })
//
type Ec2DeployAction interface {
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
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for Ec2DeployAction
type jsiiProxy_Ec2DeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_Ec2DeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2DeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewEc2DeployAction(props *Ec2DeployActionProps) Ec2DeployAction {
	_init_.Initialize()

	if err := validateNewEc2DeployActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2DeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.Ec2DeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEc2DeployAction_Override(e Ec2DeployAction, props *Ec2DeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.Ec2DeployAction",
		[]interface{}{props},
		e,
	)
}

func (e *jsiiProxy_Ec2DeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (e *jsiiProxy_Ec2DeployAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := e.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		e,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2DeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (e *jsiiProxy_Ec2DeployAction) VariableExpression(variableName *string) *string {
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

