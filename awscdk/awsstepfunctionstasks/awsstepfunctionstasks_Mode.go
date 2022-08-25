package awsstepfunctionstasks


// Specifies how many models the container hosts.
//
// Example:
//   tasks.NewSageMakerCreateModel(this, jsii.String("Sagemaker"), &sageMakerCreateModelProps{
//   	modelName: jsii.String("MyModel"),
//   	primaryContainer: tasks.NewContainerDefinition(&containerDefinitionOptions{
//   		image: tasks.dockerImage.fromJsonExpression(sfn.jsonPath.stringAt(jsii.String("$.Model.imageName"))),
//   		mode: tasks.mode_SINGLE_MODEL,
//   		modelS3Location: tasks.s3Location.fromJsonExpression(jsii.String("$.TrainingJob.ModelArtifacts.S3ModelArtifacts")),
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

