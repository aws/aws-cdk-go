package appdelivery

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/appdelivery/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// A class to deploy a stack that is part of a CDK App, using CodePipeline.
//
// This composite Action takes care of preparing and executing a CloudFormation ChangeSet.
//
// It currently does *not* support stacks that make use of ``Asset``s, and
// requires the deployed stack is in the same account and region where the
// CodePipeline is hosted.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//   import codepipeline "github.com/aws/aws-cdk-go/awscdk"
//   import codepipeline_actions "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import cicd "github.com/aws/aws-cdk-go/awscdk"
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//
//
//   type myServiceStackA struct {
//   	stack
//   }
//   type myServiceStackB struct {
//   	stack
//   }
//
//   app := cdk.NewApp()
//
//   // We define a stack that contains the CodePipeline
//   pipelineStack := cdk.NewStack(app, jsii.String("PipelineStack"))
//   pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("CodePipeline"), &pipelineProps{
//   	// Mutating a CodePipeline can cause the currently propagating state to be
//   	// "lost". Ensure we re-run the latest change through the pipeline after it's
//   	// been mutated so we're sure the latest state is fully deployed through.
//   	restartExecutionOnUpdate: jsii.Boolean(true),
//   })
//
//   // Configure the CodePipeline source - where your CDK App's source code is hosted
//   sourceOutput := codepipeline.NewArtifact()
//   source := codepipeline_actions.NewGitHubSourceAction(&gitHubSourceActionProps{
//   	actionName: jsii.String("GitHub"),
//   	output: sourceOutput,
//   	owner: jsii.String("myName"),
//   	repo: jsii.String("myRepo"),
//   	oauthToken: cdk.secretValue.unsafePlainText(jsii.String("secret")),
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("source"),
//   	actions: []iAction{
//   		source,
//   	},
//   })
//
//   project := codebuild.NewPipelineProject(pipelineStack, jsii.String("CodeBuild"), &pipelineProjectProps{
//   })
//   synthesizedApp := codepipeline.NewArtifact()
//   buildAction := codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   	actionName: jsii.String("CodeBuild"),
//   	project: project,
//   	input: sourceOutput,
//   	outputs: []artifact{
//   		synthesizedApp,
//   	},
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("build"),
//   	actions: []*iAction{
//   		buildAction,
//   	},
//   })
//
//   // Optionally, self-update the pipeline stack
//   selfUpdateStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("SelfUpdate"),
//   })
//   selfUpdateStage.addAction(cicd.NewPipelineDeployStackAction(&pipelineDeployStackActionProps{
//   	stack: pipelineStack,
//   	input: synthesizedApp,
//   	adminPermissions: jsii.Boolean(true),
//   }))
//
//   // Now add our service stacks
//   deployStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Deploy"),
//   })
//   serviceStackA := NewMyServiceStackA(app, jsii.String("ServiceStackA"), &stackProps{
//   })
//   // Add actions to deploy the stacks in the deploy stage:
//   deployServiceAAction := cicd.NewPipelineDeployStackAction(&pipelineDeployStackActionProps{
//   	stack: serviceStackA,
//   	input: synthesizedApp,
//   	// See the note below for details about this option.
//   	adminPermissions: jsii.Boolean(false),
//   })
//   deployStage.addAction(deployServiceAAction)
//   // Add the necessary permissions for you service deploy action. This role is
//   // is passed to CloudFormation and needs the permissions necessary to deploy
//   // stack. Alternatively you can enable [Administrator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_job-functions.html#jf_administrator) permissions above,
//   // users should understand the privileged nature of this role.
//   myResourceArn := "arn:partition:service:region:account-id:resource-id"
//   deployServiceAAction.addToDeploymentRolePolicy(iam.NewPolicyStatement(&policyStatementProps{
//   	actions: []*string{
//   		jsii.String("service:SomeAction"),
//   	},
//   	resources: []*string{
//   		myResourceArn,
//   	},
//   }))
//
//   serviceStackB := NewMyServiceStackB(app, jsii.String("ServiceStackB"), &stackProps{
//   })
//   deployStage.addAction(cicd.NewPipelineDeployStackAction(&pipelineDeployStackActionProps{
//   	stack: serviceStackB,
//   	input: synthesizedApp,
//   	createChangeSetRunOrder: jsii.Number(998),
//   	adminPermissions: jsii.Boolean(true),
//   }))
//
// Experimental.
type PipelineDeployStackAction interface {
	awscodepipeline.IAction
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	// Experimental.
	ActionProperties() *awscodepipeline.ActionProperties
	// Experimental.
	DeploymentRole() awsiam.IRole
	// Add policy statements to the role deploying the stack.
	//
	// This role is passed to CloudFormation and must have the IAM permissions
	// necessary to deploy the stack or you can grant this role `adminPermissions`
	// by using that option during creation. If you do not grant
	// `adminPermissions` you need to identify the proper statements to add to
	// this role based on the CloudFormation Resources in your stack.
	// Experimental.
	AddToDeploymentRolePolicy(statement awsiam.PolicyStatement)
	// The callback invoked when this Action is added to a Pipeline.
	// Experimental.
	Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// Creates an Event that will be triggered whenever the state of this Action changes.
	// Experimental.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
}

// The jsii proxy struct for PipelineDeployStackAction
type jsiiProxy_PipelineDeployStackAction struct {
	internal.Type__awscodepipelineIAction
}

func (j *jsiiProxy_PipelineDeployStackAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineDeployStackAction) DeploymentRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"deploymentRole",
		&returns,
	)
	return returns
}


// Experimental.
func NewPipelineDeployStackAction(props *PipelineDeployStackActionProps) PipelineDeployStackAction {
	_init_.Initialize()

	j := jsiiProxy_PipelineDeployStackAction{}

	_jsii_.Create(
		"monocdk.app_delivery.PipelineDeployStackAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewPipelineDeployStackAction_Override(p PipelineDeployStackAction, props *PipelineDeployStackActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.app_delivery.PipelineDeployStackAction",
		[]interface{}{props},
		p,
	)
}

func (p *jsiiProxy_PipelineDeployStackAction) AddToDeploymentRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		p,
		"addToDeploymentRolePolicy",
		[]interface{}{statement},
	)
}

func (p *jsiiProxy_PipelineDeployStackAction) Bind(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineDeployStackAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	var returns awsevents.Rule

	_jsii_.Invoke(
		p,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

