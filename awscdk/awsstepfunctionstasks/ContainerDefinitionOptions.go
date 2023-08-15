package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties to define a ContainerDefinition.
//
// Example:
//   tasks.NewSageMakerCreateModel(this, jsii.String("Sagemaker"), &SageMakerCreateModelProps{
//   	ModelName: jsii.String("MyModel"),
//   	PrimaryContainer: tasks.NewContainerDefinition(&ContainerDefinitionOptions{
//   		Image: tasks.DockerImage_FromJsonExpression(sfn.JsonPath_StringAt(jsii.String("$.Model.imageName"))),
//   		Mode: tasks.Mode_SINGLE_MODEL,
//   		ModelS3Location: tasks.S3Location_FromJsonExpression(jsii.String("$.TrainingJob.ModelArtifacts.S3ModelArtifacts")),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ContainerDefinition.html
//
type ContainerDefinitionOptions struct {
	// This parameter is ignored for models that contain only a PrimaryContainer.
	//
	// When a ContainerDefinition is part of an inference pipeline,
	// the value of the parameter uniquely identifies the container for the purposes of logging and metrics.
	// Default: - None.
	//
	ContainerHostName *string `field:"optional" json:"containerHostName" yaml:"containerHostName"`
	// The environment variables to set in the Docker container.
	// Default: - No variables.
	//
	EnvironmentVariables awsstepfunctions.TaskInput `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
	// Default: - None.
	//
	Image DockerImage `field:"optional" json:"image" yaml:"image"`
	// Defines how many models the container hosts.
	// Default: - Mode.SINGLE_MODEL
	//
	Mode Mode `field:"optional" json:"mode" yaml:"mode"`
	// The name or Amazon Resource Name (ARN) of the model package to use to create the model.
	// Default: - None.
	//
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// The S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	// The S3 path is required for Amazon SageMaker built-in algorithms, but not if you use your own algorithms.
	// Default: - None.
	//
	ModelS3Location S3Location `field:"optional" json:"modelS3Location" yaml:"modelS3Location"`
}

