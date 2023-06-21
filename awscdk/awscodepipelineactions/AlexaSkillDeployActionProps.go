package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
)

// Construction properties of the `AlexaSkillDeployAction Alexa deploy Action`.
//
// Example:
//   // Read the secrets from ParameterStore
//   clientId := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientId"))
//   clientSecret := awscdk.SecretValue_SecretsManager(jsii.String("AlexaClientSecret"))
//   refreshToken := awscdk.SecretValue_SecretsManager(jsii.String("AlexaRefreshToken"))
//
//   // Add deploy action
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewAlexaSkillDeployAction(&AlexaSkillDeployActionProps{
//   	ActionName: jsii.String("DeploySkill"),
//   	RunOrder: jsii.Number(1),
//   	Input: sourceOutput,
//   	ClientId: clientId.ToString(),
//   	ClientSecret: clientSecret,
//   	RefreshToken: refreshToken,
//   	SkillId: jsii.String("amzn1.ask.skill.12345678-1234-1234-1234-123456789012"),
//   })
//
type AlexaSkillDeployActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The client id of the developer console token.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret of the developer console token.
	ClientSecret awscdk.SecretValue `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The source artifact containing the voice model and skill manifest.
	Input awscodepipeline.Artifact `field:"required" json:"input" yaml:"input"`
	// The refresh token of the developer console token.
	RefreshToken awscdk.SecretValue `field:"required" json:"refreshToken" yaml:"refreshToken"`
	// The Alexa skill id.
	SkillId *string `field:"required" json:"skillId" yaml:"skillId"`
	// An optional artifact containing overrides for the skill manifest.
	ParameterOverridesArtifact awscodepipeline.Artifact `field:"optional" json:"parameterOverridesArtifact" yaml:"parameterOverridesArtifact"`
}

