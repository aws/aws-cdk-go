package awscodepipelineactions


// How should the CodeCommit Action detect changes.
//
// This is the type of the {@link CodeCommitSourceAction.trigger} property.
//
// Example:
//   // Source stage: read from repository
//   repo := codecommit.NewRepository(stack, jsii.String("TemplateRepo"), &repositoryProps{
//   	repositoryName: jsii.String("template-repo"),
//   })
//   sourceOutput := codepipeline.NewArtifact(jsii.String("SourceArtifact"))
//   source := cpactions.NewCodeCommitSourceAction(&codeCommitSourceActionProps{
//   	actionName: jsii.String("Source"),
//   	repository: repo,
//   	output: sourceOutput,
//   	trigger: cpactions.codeCommitTrigger_POLL,
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
//   codepipeline.NewPipeline(stack, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type CodeCommitTrigger string

const (
	// The Action will never detect changes - the Pipeline it's part of will only begin a run when explicitly started.
	CodeCommitTrigger_NONE CodeCommitTrigger = "NONE"
	// CodePipeline will poll the repository to detect changes.
	CodeCommitTrigger_POLL CodeCommitTrigger = "POLL"
	// CodePipeline will use CloudWatch Events to be notified of changes.
	//
	// This is the default method of detecting changes.
	CodeCommitTrigger_EVENTS CodeCommitTrigger = "EVENTS"
)

