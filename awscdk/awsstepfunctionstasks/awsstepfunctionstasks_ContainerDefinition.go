package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Describes the container, as part of model definition.
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
	var returns *ContainerDefinitionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

