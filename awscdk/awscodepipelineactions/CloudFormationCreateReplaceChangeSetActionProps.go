package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for the CloudFormationCreateReplaceChangeSetAction.
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
//   	CrossAccountKeys: jsii.Boolean(true),
//   	Stages: []StageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CloudFormationCreateReplaceChangeSetActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Default: 1.
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Default: - a name will be generated, based on the stage and action names,
	// if any of the action's variables were referenced - otherwise,
	// no namespace will be set.
	//
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your `IAction.bind`
	// method in the `ActionBindOptions.role` property.
	// Default: a new Role will be generated.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to grant full permissions to CloudFormation while deploying this template.
	//
	// Setting this to `true` affects the defaults for `role` and `capabilities`, if you
	// don't specify any alternatives.
	//
	// The default role that will be created for you will have full (i.e., `*`)
	// permissions on all resources, and the deployment will have named IAM
	// capabilities (i.e., able to create all IAM resources).
	//
	// This is a shorthand that you can use if you fully trust the templates that
	// are deployed in this pipeline. If you want more fine-grained permissions,
	// use `addToRolePolicy` and `capabilities` to control what the CloudFormation
	// deployment is allowed to do.
	AdminPermissions *bool `field:"required" json:"adminPermissions" yaml:"adminPermissions"`
	// Name of the change set to create or update.
	ChangeSetName *string `field:"required" json:"changeSetName" yaml:"changeSetName"`
	// The name of the stack to apply this action to.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Input artifact with the ChangeSet's CloudFormation template.
	TemplatePath awscodepipeline.ArtifactPath `field:"required" json:"templatePath" yaml:"templatePath"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	// Default: - action resides in the same account as the pipeline.
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	// Default: None, unless `adminPermissions` is true.
	//
	CfnCapabilities *[]awscdk.CfnCapabilities `field:"optional" json:"cfnCapabilities" yaml:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	// Default: A fresh role with full or no permissions (depending on the value of `adminPermissions`).
	//
	DeploymentRole awsiam.IRole `field:"optional" json:"deploymentRole" yaml:"deploymentRole"`
	// The list of additional input Artifacts for this Action.
	//
	// This is especially useful when used in conjunction with the `parameterOverrides` property.
	// For example, if you have:
	//
	//   parameterOverrides: {
	//     'Param1': action1.outputArtifact.bucketName,
	//     'Param2': action2.outputArtifact.objectKey,
	//   }
	//
	// , if the output Artifacts of `action1` and `action2` were not used to
	// set either the `templateConfiguration` or the `templatePath` properties,
	// you need to make sure to include them in the `extraInputs` -
	// otherwise, you'll get an "unrecognized Artifact" error during your Pipeline's execution.
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	// Default: Automatically generated artifact name.
	//
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	// Default: No output artifact generated.
	//
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// Additional template parameters.
	//
	// Template parameters specified here take precedence over template parameters
	// found in the artifact specified by the `templateConfiguration` property.
	//
	// We recommend that you use the template configuration file to specify
	// most of your parameter values. Use parameter overrides to specify only
	// dynamic parameter values (values that are unknown until you run the
	// pipeline).
	//
	// All parameter names must be present in the stack template.
	//
	// Note: the entire object cannot be more than 1kB.
	// Default: No overrides.
	//
	ParameterOverrides *map[string]interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps#crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	// Default: the Action resides in the same region as the Pipeline.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	// Default: No template configuration based on input artifacts.
	//
	TemplateConfiguration awscodepipeline.ArtifactPath `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

