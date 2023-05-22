package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Describes the container, as part of model definition.
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
// Experimental.
type ContainerDefinition interface {
	IContainerDefinition
	// Called when the ContainerDefinition type configured on Sagemaker Task.
	// Experimental.
	Bind(task ISageMakerTask) *ContainerDefinitionConfig
}

// The jsii proxy struct for ContainerDefinition
type jsiiProxy_ContainerDefinition struct {
	jsiiProxy_IContainerDefinition
}

// Experimental.
func NewContainerDefinition(options *ContainerDefinitionOptions) ContainerDefinition {
	_init_.Initialize()

	if err := validateNewContainerDefinitionParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerDefinition{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.ContainerDefinition",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewContainerDefinition_Override(c ContainerDefinition, options *ContainerDefinitionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.ContainerDefinition",
		[]interface{}{options},
		c,
	)
}

func (c *jsiiProxy_ContainerDefinition) Bind(task ISageMakerTask) *ContainerDefinitionConfig {
	if err := c.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *ContainerDefinitionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

