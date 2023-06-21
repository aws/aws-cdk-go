package awsstepfunctionstasks


// Specifies how many models the container hosts.
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
type Mode string

const (
	// Container hosts a single model.
	Mode_SINGLE_MODEL Mode = "SINGLE_MODEL"
	// Container hosts multiple models.
	// See: https://docs.aws.amazon.com/sagemaker/latest/dg/multi-model-endpoints.html
	//
	Mode_MULTI_MODEL Mode = "MULTI_MODEL"
)

