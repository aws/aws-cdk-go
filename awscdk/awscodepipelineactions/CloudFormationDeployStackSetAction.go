package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// CodePipeline action to deploy a stackset.
//
// CodePipeline offers the ability to perform AWS CloudFormation StackSets
// operations as part of your CI/CD process. You use a stack set to create
// stacks in AWS accounts across AWS Regions by using a single AWS
// CloudFormation template. All the resources included in each stack are defined
// by the stack set’s AWS CloudFormation template. When you create the stack
// set, you specify the template to use, as well as any parameters and
// capabilities that the template requires.
//
// For more information about concepts for AWS CloudFormation StackSets, see
// [StackSets
// concepts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html)
// in the AWS CloudFormation User Guide.
//
// If you use this action to make an update that includes adding stack
// instances, the new instances are deployed first and the update is completed
// last. The new instances first receive the old version, and then the update is
// applied to all instances.
//
// As a best practice, you should construct your pipeline so that the stack set
// is created and initially deploys to a subset or a single instance. After you
// test your deployment and view the generated stack set, then add the
// CloudFormationStackInstances action so that the remaining instances are
// created and updated.
//
// Example:
//   var pipeline Pipeline
//   var sourceOutput Artifact
//
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("DeployStackSets"),
//   	Actions: []IAction{
//   		// First, update the StackSet itself with the newest template
//   		codepipeline_actions.NewCloudFormationDeployStackSetAction(&CloudFormationDeployStackSetActionProps{
//   			ActionName: jsii.String("UpdateStackSet"),
//   			RunOrder: jsii.Number(1),
//   			StackSetName: jsii.String("MyStackSet"),
//   			Template: codepipeline_actions.StackSetTemplate_FromArtifactPath(sourceOutput.AtPath(jsii.String("template.yaml"))),
//
//   			// Change this to 'StackSetDeploymentModel.organizations()' if you want to deploy to OUs
//   			DeploymentModel: codepipeline_actions.StackSetDeploymentModel_SelfManaged(),
//   			// This deploys to a set of accounts
//   			StackInstances: codepipeline_actions.StackInstances_InAccounts([]*string{
//   				jsii.String("111111111111"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//
//   		// Afterwards, update/create additional instances in other accounts
//   		codepipeline_actions.NewCloudFormationDeployStackInstancesAction(&CloudFormationDeployStackInstancesActionProps{
//   			ActionName: jsii.String("AddMoreInstances"),
//   			RunOrder: jsii.Number(2),
//   			StackSetName: jsii.String("MyStackSet"),
//   			StackInstances: codepipeline_actions.StackInstances_*InAccounts([]*string{
//   				jsii.String("222222222222"),
//   				jsii.String("333333333333"),
//   			}, []*string{
//   				jsii.String("us-east-1"),
//   				jsii.String("eu-west-1"),
//   			}),
//   		}),
//   	},
//   })
//
type CloudFormationDeployStackSetAction interface {
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

// The jsii proxy struct for CloudFormationDeployStackSetAction
type jsiiProxy_CloudFormationDeployStackSetAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_CloudFormationDeployStackSetAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationDeployStackSetAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewCloudFormationDeployStackSetAction(props *CloudFormationDeployStackSetActionProps) CloudFormationDeployStackSetAction {
	_init_.Initialize()

	if err := validateNewCloudFormationDeployStackSetActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationDeployStackSetAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeployStackSetAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewCloudFormationDeployStackSetAction_Override(c CloudFormationDeployStackSetAction, props *CloudFormationDeployStackSetActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.CloudFormationDeployStackSetAction",
		[]interface{}{props},
		c,
	)
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CloudFormationDeployStackSetAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
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

func (c *jsiiProxy_CloudFormationDeployStackSetAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
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

func (c *jsiiProxy_CloudFormationDeployStackSetAction) VariableExpression(variableName *string) *string {
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

