package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline action to prepare a change set.
//
// Creates the change set if it doesn't exist based on the stack name and template that you submit.
// If the change set exists, AWS CloudFormation deletes it, and then creates a new one.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &RepositoryProps{
//   	RepositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   	ActionName: jsii.String("Source"),
//   	Repository: repo,
//   	Output: sourceOutput,
//   	Trigger: cpactions.CodeCommitTrigger_POLL,
//   })
//   sourceStage := map[string]interface{}{
//   	"stageName": jsii.String("Source"),
//   	"actions": []CodeCommitSourceAction{
//   		source,
//   	},
//   }
//
//   // Deployment stage: create and deploy changeset with manual approval
//   stackName := "OurStack"
//   changeSetName := "StagedChangeSet"
//
//   prodStage := map[string]interface{}{
//   	"stageName": jsii.String("Deploy"),
//   	"actions": []interface{}{
//   		cpactions.NewCloudFormationCreateReplaceChangeSetAction(&CloudFormationCreateReplaceChangeSetActionProps{
//   			"actionName": jsii.String("PrepareChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"adminPermissions": jsii.Boolean(true),
//   			"templatePath": sourceOutput.atPath(jsii.String("template.yaml")),
//   			"runOrder": jsii.Number(1),
//   		}),
//   		cpactions.NewManualApprovalAction(&ManualApprovalActionProps{
//   			"actionName": jsii.String("ApproveChanges"),
//   			"runOrder": jsii.Number(2),
//   		}),
//   		cpactions.NewCloudFormationExecuteChangeSetAction(&CloudFormationExecuteChangeSetActionProps{
//   			"actionName": jsii.String("ExecuteChanges"),
//   			"stackName": jsii.String(stackName),
//   			"changeSetName": jsii.String(changeSetName),
//   			"runOrder": jsii.Number(3),
//   		}),
//   	},
//   }
//
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &PipelineProps{
//   	Stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CloudFormationCreateReplaceChangeSetAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the `bind` callback.
	ActionProperties() *awscodepipeline.ActionProperties
	DeploymentRole() awsiam.IRole
	// This is a renamed version of the `IAction.actionProperties` property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// Add statement to the service role assumed by CloudFormation while executing this action.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the `IAction.bind` method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for CloudFormationCreateReplaceChangeSetAction
type jsiiProxy_CloudFormationCreateReplaceChangeSetAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) DeploymentRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationCreateReplaceChangeSetAction(props *CloudFormationCreateReplaceChangeSetActionProps) CloudFormationCreateReplaceChangeSetAction {
	_init_.Initialize()

	if err := validateNewCloudFormationCreateReplaceChangeSetActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationCreateReplaceChangeSetAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateReplaceChangeSetAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationCreateReplaceChangeSetAction_Override(c CloudFormationCreateReplaceChangeSetAction, props *CloudFormationCreateReplaceChangeSetActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationCreateReplaceChangeSetAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) *bool {
	if err := c.validateAddToDeploymentRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) VariableExpression(variableName *string) *string {
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

