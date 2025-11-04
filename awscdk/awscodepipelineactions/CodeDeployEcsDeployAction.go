package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact Artifact
//   var artifactPath ArtifactPath
//   var ecsDeploymentGroup EcsDeploymentGroup
//   var role Role
//
//   codeDeployEcsDeployAction := awscdk.Aws_codepipeline_actions.NewCodeDeployEcsDeployAction(&CodeDeployEcsDeployActionProps{
//   	ActionName: jsii.String("actionName"),
//   	DeploymentGroup: ecsDeploymentGroup,
//
//   	// the properties below are optional
//   	AppSpecTemplateFile: artifactPath,
//   	AppSpecTemplateInput: artifact,
//   	ContainerImageInputs: []CodeDeployEcsContainerImageInput{
//   		&CodeDeployEcsContainerImageInput{
//   			Input: artifact,
//
//   			// the properties below are optional
//   			TaskDefinitionPlaceholder: jsii.String("taskDefinitionPlaceholder"),
//   		},
//   	},
//   	Role: role,
//   	RunOrder: jsii.Number(123),
//   	TaskDefinitionTemplateFile: artifactPath,
//   	TaskDefinitionTemplateInput: artifact,
//   	VariablesNamespace: jsii.String("variablesNamespace"),
//   })
//
type CodeDeployEcsDeployAction interface {
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

// The jsii proxy struct for CodeDeployEcsDeployAction
type jsiiProxy_CodeDeployEcsDeployAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CodeDeployEcsDeployAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeDeployEcsDeployAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCodeDeployEcsDeployAction(props *CodeDeployEcsDeployActionProps) CodeDeployEcsDeployAction {
	_init_.Initialize()

	if err := validateNewCodeDeployEcsDeployActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeDeployEcsDeployAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployEcsDeployAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCodeDeployEcsDeployAction_Override(c CodeDeployEcsDeployAction, props *CodeDeployEcsDeployActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CodeDeployEcsDeployAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := c.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := c.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		c,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := c.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) VariableExpression(variableName *string) *string {
	if err := c.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

