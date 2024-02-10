package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A specific file within an output artifact.
//
// The most common use case for this is specifying the template file
// for a CloudFormation action.
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
//   	Stages: []stageProps{
//   		sourceStage,
//   		prodStage,
//   	},
//   })
//
type ArtifactPath interface {
	Artifact() Artifact
	FileName() *string
	Location() *string
}

// The jsii proxy struct for ArtifactPath
type jsiiProxy_ArtifactPath struct {
	_ byte // padding
}

func (j *jsiiProxy_ArtifactPath) Artifact() Artifact {
	var returns Artifact
	_jsii_.Get(
		j,
		"artifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactPath) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactPath) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}


func NewArtifactPath(artifact Artifact, fileName *string) ArtifactPath {
	_init_.Initialize()

	if err := validateNewArtifactPathParameters(artifact, fileName); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArtifactPath{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.ArtifactPath",
		[]interface{}{artifact, fileName},
		&j,
	)

	return &j
}

func NewArtifactPath_Override(a ArtifactPath, artifact Artifact, fileName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.ArtifactPath",
		[]interface{}{artifact, fileName},
		a,
	)
}

func ArtifactPath_ArtifactPath(artifactName *string, fileName *string) ArtifactPath {
	_init_.Initialize()

	if err := validateArtifactPath_ArtifactPathParameters(artifactName, fileName); err != nil {
		panic(err)
	}
	var returns ArtifactPath

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.ArtifactPath",
		"artifactPath",
		[]interface{}{artifactName, fileName},
		&returns,
	)

	return returns
}

