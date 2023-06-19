package appdelivery

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudformation"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import codebuild "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   pipeline := codepipeline.NewPipeline(pipelineStack, jsii.String("CodePipeline"), &PipelineProps{
//   	// Mutating a CodePipeline can cause the currently propagating state to be
//   	// "lost". Ensure we re-run the latest change through the pipeline after it's
//   	// been mutated so we're sure the latest state is fully deployed through.
//   	RestartExecutionOnUpdate: jsii.Boolean(true),
//   })
//
//   // Configure the CodePipeline source - where your CDK App's source code is hosted
//   sourceOutput := codepipeline.NewArtifact()
//   source := codepipeline_actions.NewGitHubSourceAction(&GitHubSourceActionProps{
//   	ActionName: jsii.String("GitHub"),
//   	Output: sourceOutput,
//   	Owner: jsii.String("myName"),
//   	Repo: jsii.String("myRepo"),
//   	OauthToken: cdk.SecretValue_UnsafePlainText(jsii.String("secret")),
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("source"),
//   	Actions: []iAction{
//   		source,
//   	},
//   })
//
//   project := codebuild.NewPipelineProject(pipelineStack, jsii.String("CodeBuild"), &PipelineProjectProps{
//   })
//   synthesizedApp := codepipeline.NewArtifact()
//   buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CodeBuild"),
//   	Project: Project,
//   	Input: sourceOutput,
//   	Outputs: []artifact{
//   		synthesizedApp,
//   	},
//   })
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("build"),
//   	Actions: []*iAction{
//   		buildAction,
//   	},
//   })
//
//   // Optionally, self-update the pipeline stack
//   selfUpdateStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("SelfUpdate"),
//   })
//   selfUpdateStage.AddAction(cicd.NewPipelineDeployStackAction(&PipelineDeployStackActionProps{
//   	Stack: pipelineStack,
//   	Input: synthesizedApp,
//   	AdminPermissions: jsii.Boolean(true),
//   }))
//
//   // Now add our service stacks
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   })
//   serviceStackA := NewMyServiceStackA(app, jsii.String("ServiceStackA"), &stackProps{
//   })
//   // Add actions to deploy the stacks in the deploy stage:
//   deployServiceAAction := cicd.NewPipelineDeployStackAction(&PipelineDeployStackActionProps{
//   	Stack: serviceStackA,
//   	Input: synthesizedApp,
//   	// See the note below for details about this option.
//   	AdminPermissions: jsii.Boolean(false),
//   })
//   deployStage.AddAction(deployServiceAAction)
//   // Add the necessary permissions for you service deploy action. This role is
//   // is passed to CloudFormation and needs the permissions necessary to deploy
//   // stack. Alternatively you can enable [Administrator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_job-functions.html#jf_administrator) permissions above,
//   // users should understand the privileged nature of this role.
//   myResourceArn := "arn:partition:service:region:account-id:resource-id"
//   deployServiceAAction.AddToDeploymentRolePolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("service:SomeAction"),
//   	},
//   	Resources: []*string{
//   		myResourceArn,
//   	},
//   }))
//
//   serviceStackB := NewMyServiceStackB(app, jsii.String("ServiceStackB"), &stackProps{
//   })
//   deployStage.AddAction(cicd.NewPipelineDeployStackAction(&PipelineDeployStackActionProps{
//   	Stack: serviceStackB,
//   	Input: synthesizedApp,
//   	CreateChangeSetRunOrder: jsii.Number(998),
//   	AdminPermissions: jsii.Boolean(true),
//   }))
//
// Experimental.
type PipelineDeployStackActionProps struct {
	// Whether to grant admin permissions to CloudFormation while deploying this template.
	//
	// Setting this to `true` affects the defaults for `role` and `capabilities`, if you
	// don't specify any alternatives.
	//
	// The default role that will be created for you will have admin (i.e., `*`)
	// permissions on all resources, and the deployment will have named IAM
	// capabilities (i.e., able to create all IAM resources).
	//
	// This is a shorthand that you can use if you fully trust the templates that
	// are deployed in this pipeline. If you want more fine-grained permissions,
	// use `addToRolePolicy` and `capabilities` to control what the CloudFormation
	// deployment is allowed to do.
	// Experimental.
	AdminPermissions *bool `field:"required" json:"adminPermissions" yaml:"adminPermissions"`
	// The CodePipeline artifact that holds the synthesized app, which is the contents of the ``<directory>`` when running ``cdk synth -o <directory>``.
	// Experimental.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The CDK stack to be deployed.
	// Experimental.
	Stack awscdk.Stack `field:"required" json:"stack" yaml:"stack"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources, explicit acknowledgement that AWS CloudFormation
	// might create or update those resources. For example, you must specify AnonymousIAM if your
	// stack template contains AWS Identity and Access Management (IAM) resources. For more
	// information.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Experimental.
	Capabilities *[]awscloudformation.CloudFormationCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The name to use when creating a ChangeSet for the stack.
	// Experimental.
	ChangeSetName *string `field:"optional" json:"changeSetName" yaml:"changeSetName"`
	// The name of the CodePipeline action creating the ChangeSet.
	// Experimental.
	CreateChangeSetActionName *string `field:"optional" json:"createChangeSetActionName" yaml:"createChangeSetActionName"`
	// The runOrder for the CodePipeline action creating the ChangeSet.
	// Experimental.
	CreateChangeSetRunOrder *float64 `field:"optional" json:"createChangeSetRunOrder" yaml:"createChangeSetRunOrder"`
	// The name of the CodePipeline action creating the ChangeSet.
	// Experimental.
	ExecuteChangeSetActionName *string `field:"optional" json:"executeChangeSetActionName" yaml:"executeChangeSetActionName"`
	// The runOrder for the CodePipeline action executing the ChangeSet.
	// Experimental.
	ExecuteChangeSetRunOrder *float64 `field:"optional" json:"executeChangeSetRunOrder" yaml:"executeChangeSetRunOrder"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// admin permissions.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

