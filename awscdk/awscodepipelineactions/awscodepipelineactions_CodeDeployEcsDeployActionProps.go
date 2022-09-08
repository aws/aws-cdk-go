package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties of the {@link CodeDeployEcsDeployAction CodeDeploy ECS deploy CodePipeline Action}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactPath artifactPath
//   var ecsDeploymentGroup iEcsDeploymentGroup
//   var role role
//
//   codeDeployEcsDeployActionProps := &codeDeployEcsDeployActionProps{
//   	actionName: jsii.String("actionName"),
//   	deploymentGroup: ecsDeploymentGroup,
//
//   	// the properties below are optional
//   	appSpecTemplateFile: artifactPath,
//   	appSpecTemplateInput: artifact,
//   	containerImageInputs: []codeDeployEcsContainerImageInput{
//   		&codeDeployEcsContainerImageInput{
//   			input: artifact,
//
//   			// the properties below are optional
//   			taskDefinitionPlaceholder: jsii.String("taskDefinitionPlaceholder"),
//   		},
//   	},
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	taskDefinitionTemplateFile: artifactPath,
//   	taskDefinitionTemplateInput: artifact,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   }
//
// Experimental.
type CodeDeployEcsDeployActionProps struct {
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
	// The CodeDeploy ECS Deployment Group to deploy to.
	// Experimental.
	DeploymentGroup awscodedeploy.IEcsDeploymentGroup `field:"required" json:"deploymentGroup" yaml:"deploymentGroup"`
	// The name of the CodeDeploy AppSpec file.
	//
	// During deployment, a new task definition will be registered
	// with ECS, and the new task definition ID will be inserted into
	// the CodeDeploy AppSpec file.  The AppSpec file contents will be
	// provided to CodeDeploy for the deployment.
	//
	// Use this property if you want to use a different name for this file than the default 'appspec.yaml'.
	// If you use this property, you don't need to specify the `appSpecTemplateInput` property.
	// Experimental.
	AppSpecTemplateFile awscodepipeline.ArtifactPath `field:"optional" json:"appSpecTemplateFile" yaml:"appSpecTemplateFile"`
	// The artifact containing the CodeDeploy AppSpec file.
	//
	// During deployment, a new task definition will be registered
	// with ECS, and the new task definition ID will be inserted into
	// the CodeDeploy AppSpec file.  The AppSpec file contents will be
	// provided to CodeDeploy for the deployment.
	//
	// If you use this property, it's assumed the file is called 'appspec.yaml'.
	// If your AppSpec file uses a different filename, leave this property empty,
	// and use the `appSpecTemplateFile` property instead.
	// Experimental.
	AppSpecTemplateInput awscodepipeline.Artifact `field:"optional" json:"appSpecTemplateInput" yaml:"appSpecTemplateInput"`
	// Configuration for dynamically updated images in the task definition.
	//
	// Provide pairs of an image details input artifact and a placeholder string
	// that will be used to dynamically update the ECS task definition template
	// file prior to deployment. A maximum of 4 images can be given.
	// Experimental.
	ContainerImageInputs *[]*CodeDeployEcsContainerImageInput `field:"optional" json:"containerImageInputs" yaml:"containerImageInputs"`
	// The name of the ECS task definition template file.
	//
	// During deployment, the task definition template file contents
	// will be registered with ECS.
	//
	// Use this property if you want to use a different name for this file than the default 'taskdef.json'.
	// If you use this property, you don't need to specify the `taskDefinitionTemplateInput` property.
	// Experimental.
	TaskDefinitionTemplateFile awscodepipeline.ArtifactPath `field:"optional" json:"taskDefinitionTemplateFile" yaml:"taskDefinitionTemplateFile"`
	// The artifact containing the ECS task definition template file.
	//
	// During deployment, the task definition template file contents
	// will be registered with ECS.
	//
	// If you use this property, it's assumed the file is called 'taskdef.json'.
	// If your task definition template uses a different filename, leave this property empty,
	// and use the `taskDefinitionTemplateFile` property instead.
	// Experimental.
	TaskDefinitionTemplateInput awscodepipeline.Artifact `field:"optional" json:"taskDefinitionTemplateInput" yaml:"taskDefinitionTemplateInput"`
}

