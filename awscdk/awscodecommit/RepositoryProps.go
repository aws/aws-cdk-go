package awscodecommit

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Example:
//   var project pipelineProject
//
//   repository := codecommit.NewRepository(this, jsii.String("MyRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("MyRepository"),
//   })
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   	ActionName: jsii.String("CodeCommit"),
//   	Repository: Repository,
//   	Output: sourceOutput,
//   })
//   buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CodeBuild"),
//   	Project: Project,
//   	Input: sourceOutput,
//   	Outputs: []artifact{
//   		codepipeline.NewArtifact(),
//   	},
//   	 // optional
//   	ExecuteBatchBuild: jsii.Boolean(true),
//   	 // optional, defaults to false
//   	CombineBatchBuildArtifacts: jsii.Boolean(true),
//   })
//
//   codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
//   	Stages: []stageProps{
//   		&stageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []iAction{
//   				sourceAction,
//   			},
//   		},
//   		&stageProps{
//   			StageName: jsii.String("Build"),
//   			Actions: []*iAction{
//   				buildAction,
//   			},
//   		},
//   	},
//   })
//
type RepositoryProps struct {
	// Name of the repository.
	//
	// This property is required for all CodeCommit repositories.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The contents with which to initialize the repository after it has been created.
	// Default: - No initialization (create empty repo).
	//
	Code Code `field:"optional" json:"code" yaml:"code"`
	// A description of the repository.
	//
	// Use the description to identify the
	// purpose of the repository.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The customer managed key used to encrypt and decrypt the data in repository.
	// Default: - Use an AWS managed key.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

