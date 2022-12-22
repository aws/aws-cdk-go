package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties of the {@link CodeBuildAction CodeBuild build CodePipeline action}.
//
// Example:
//   // Create a Cloudfront Web Distribution
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//   var distribution distribution
//
//
//   // Create the build project that will invalidate the cache
//   invalidateBuildProject := codebuild.NewPipelineProject(this, jsii.String("InvalidateProject"), &pipelineProjectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("aws cloudfront create-invalidation --distribution-id ${CLOUDFRONT_ID} --paths \"/*\""),
//   				},
//   			},
//   		},
//   	}),
//   	environmentVariables: map[string]buildEnvironmentVariable{
//   		"CLOUDFRONT_ID": &buildEnvironmentVariable{
//   			"value": distribution.distributionId,
//   		},
//   	},
//   })
//
//   // Add Cloudfront invalidation permissions to the project
//   distributionArn := fmt.Sprintf("arn:aws:cloudfront::%v:distribution/%v", this.account, distribution.distributionId)
//   invalidateBuildProject.addToRolePolicy(iam.NewPolicyStatement(&policyStatementProps{
//   	resources: []*string{
//   		distributionArn,
//   	},
//   	actions: []*string{
//   		jsii.String("cloudfront:CreateInvalidation"),
//   	},
//   }))
//
//   // Create the pipeline (here only the S3 deploy and Invalidate cache build)
//   deployBucket := s3.NewBucket(this, jsii.String("DeployBucket"))
//   deployInput := codepipeline.NewArtifact()
//   codepipeline.NewPipeline(this, jsii.String("Pipeline"), &pipelineProps{
//   	stages: []stageProps{
//   		&stageProps{
//   			stageName: jsii.String("Deploy"),
//   			actions: []iAction{
//   				codepipeline_actions.NewS3DeployAction(&s3DeployActionProps{
//   					actionName: jsii.String("S3Deploy"),
//   					bucket: deployBucket,
//   					input: deployInput,
//   					runOrder: jsii.Number(1),
//   				}),
//   				codepipeline_actions.NewCodeBuildAction(&codeBuildActionProps{
//   					actionName: jsii.String("InvalidateCache"),
//   					project: invalidateBuildProject,
//   					input: deployInput,
//   					runOrder: jsii.Number(2),
//   				}),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type CodeBuildActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	// Experimental.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	// Experimental.
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	// Experimental.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The source to use as input for this action.
	// Experimental.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The action's Project.
	// Experimental.
	Project awscodebuild.IProject `field:"required" json:"project" yaml:"project"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	// Experimental.
	CheckSecretsInPlainTextEnvVariables *bool `field:"optional" json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Combine the build artifacts for a batch builds.
	//
	// Enabling this will combine the build artifacts into the same location for batch builds.
	// If `executeBatchBuild` is not set to `true`, this property is ignored.
	// Experimental.
	CombineBatchBuildArtifacts *bool `field:"optional" json:"combineBatchBuildArtifacts" yaml:"combineBatchBuildArtifacts"`
	// The environment variables to pass to the CodeBuild project when this action executes.
	//
	// If a variable with the same name was set both on the project level, and here,
	// this value will take precedence.
	// Experimental.
	EnvironmentVariables *map[string]*awscodebuild.BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Trigger a batch build.
	//
	// Enabling this will enable batch builds on the CodeBuild project.
	// Experimental.
	ExecuteBatchBuild *bool `field:"optional" json:"executeBatchBuild" yaml:"executeBatchBuild"`
	// The list of additional input Artifacts for this action.
	//
	// The directories the additional inputs will be available at are available
	// during the project's build in the CODEBUILD_SRC_DIR_<artifact-name> environment variables.
	// The project's build always starts in the directory with the primary input artifact checked out,
	// the one pointed to by the {@link input} property.
	// For more information,
	// see https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html .
	// Experimental.
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The list of output Artifacts for this action.
	//
	// **Note**: if you specify more than one output Artifact here,
	// you cannot use the primary 'artifacts' section of the buildspec;
	// you have to use the 'secondary-artifacts' section instead.
	// See https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	// for details.
	// Experimental.
	Outputs *[]awscodepipeline.Artifact `field:"optional" json:"outputs" yaml:"outputs"`
	// The type of the action that determines its CodePipeline Category - Build, or Test.
	// Experimental.
	Type CodeBuildActionType `field:"optional" json:"type" yaml:"type"`
}

