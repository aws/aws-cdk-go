package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties of the `InspectorEcrImageScanAction`.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline pipeline
//   var repository iRepository
//
//
//   scanOutput := codepipeline.NewArtifact()
//   scanAction := codepipeline_actions.NewInspectorEcrImageScanAction(&InspectorEcrImageScanActionProps{
//   	ActionName: jsii.String("InspectorEcrImageScanAction"),
//   	Output: scanOutput,
//   	Repository: repository,
//   })
//
//   pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Scan"),
//   	Actions: []iAction{
//   		scanAction,
//   	},
//   })
//
type InspectorEcrImageScanActionProps struct {
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
	// Vulnerability details of your source in the form of a Software Bill of Materials (SBOM) file.
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The number of critical severity vulnerabilities found in your source beyond which CodePipeline should fail the action.
	// Default: - no threshold.
	//
	CriticalThreshold *float64 `field:"optional" json:"criticalThreshold" yaml:"criticalThreshold"`
	// The number of high severity vulnerabilities found in your source beyond which CodePipeline should fail the action.
	// Default: - no threshold.
	//
	HighThreshold *float64 `field:"optional" json:"highThreshold" yaml:"highThreshold"`
	// The number of low severity vulnerabilities found in your source beyond which CodePipeline should fail the action.
	// Default: - no threshold.
	//
	LowThreshold *float64 `field:"optional" json:"lowThreshold" yaml:"lowThreshold"`
	// The number of medium severity vulnerabilities found in your source beyond which CodePipeline should fail the action.
	// Default: - no threshold.
	//
	MediumThreshold *float64 `field:"optional" json:"mediumThreshold" yaml:"mediumThreshold"`
	// The Amazon ECR repository where the image is pushed.
	Repository awsecr.IRepository `field:"required" json:"repository" yaml:"repository"`
	// The tag used for the image.
	// Default: 'latest'.
	//
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

