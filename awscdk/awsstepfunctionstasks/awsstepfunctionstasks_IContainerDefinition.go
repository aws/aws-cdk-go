package awsstepfunctionstasks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration of the container used to host the model.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ContainerDefinition.html
//
type IContainerDefinition interface {
	// Called when the ContainerDefinition is used by a SageMaker task.
	Bind(task ISageMakerTask) *ContainerDefinitionConfig
}

// The jsii proxy for IContainerDefinition
type jsiiProxy_IContainerDefinition struct {
	_ byte // padding
}

func (i *jsiiProxy_IContainerDefinition) Bind(task ISageMakerTask) *ContainerDefinitionConfig {
	if err := i.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *ContainerDefinitionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

