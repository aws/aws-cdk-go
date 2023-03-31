package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties of the {@link ServiceCatalogDeployActionBeta1 ServiceCatalog deploy CodePipeline Action}.
//
// Example:
//   cdkBuildOutput := codepipeline.NewArtifact()
//   serviceCatalogDeployAction := codepipeline_actions.NewServiceCatalogDeployActionBeta1(&serviceCatalogDeployActionBeta1Props{
//   	actionName: jsii.String("ServiceCatalogDeploy"),
//   	templatePath: cdkBuildOutput.atPath(jsii.String("Sample.template.json")),
//   	productVersionName: jsii.String("Version - " + date.now.toString),
//   	productVersionDescription: jsii.String("This is a version from the pipeline with a new description."),
//   	productId: jsii.String("prod-XXXXXXXX"),
//   })
//
// Experimental.
type ServiceCatalogDeployActionBeta1Props struct {
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
	// The identifier of the product in the Service Catalog.
	//
	// This product must already exist.
	// Experimental.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The name of the version of the Service Catalog product to be deployed.
	// Experimental.
	ProductVersionName *string `field:"required" json:"productVersionName" yaml:"productVersionName"`
	// The path to the cloudformation artifact.
	// Experimental.
	TemplatePath awscodepipeline.ArtifactPath `field:"required" json:"templatePath" yaml:"templatePath"`
	// The optional description of this version of the Service Catalog product.
	// Experimental.
	ProductVersionDescription *string `field:"optional" json:"productVersionDescription" yaml:"productVersionDescription"`
}

