package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for working with AWS-managed workflows.
//
// Example:
//   containerWorkflowPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ContainerWorkflowPipeline"), &ImagePipelineProps{
//   	Recipe: exampleContainerRecipe,
//   	Workflows: []WorkflowConfiguration{
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AwsManagedWorkflow_BuildContainer(this, jsii.String("BuildContainer")),
//   		},
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AwsManagedWorkflow_TestContainer(this, jsii.String("TestContainer")),
//   		},
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AwsManagedWorkflow_DistributeContainer(this, jsii.String("DistributeContainer")),
//   		},
//   	},
//   })
//
// Experimental.
type AwsManagedWorkflow interface {
}

// The jsii proxy struct for AwsManagedWorkflow
type jsiiProxy_AwsManagedWorkflow struct {
	_ byte // padding
}

// Experimental.
func NewAwsManagedWorkflow() AwsManagedWorkflow {
	_init_.Initialize()

	j := jsiiProxy_AwsManagedWorkflow{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAwsManagedWorkflow_Override(a AwsManagedWorkflow) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		nil, // no parameters
		a,
	)
}

// Imports the build-container AWS-managed workflow.
// Experimental.
func AwsManagedWorkflow_BuildContainer(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAwsManagedWorkflow_BuildContainerParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		"buildContainer",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports the build-image AWS-managed workflow.
// Experimental.
func AwsManagedWorkflow_BuildImage(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAwsManagedWorkflow_BuildImageParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		"buildImage",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports the distribute-container AWS-managed workflow.
// Experimental.
func AwsManagedWorkflow_DistributeContainer(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAwsManagedWorkflow_DistributeContainerParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		"distributeContainer",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports an AWS-managed workflow from its attributes.
// Experimental.
func AwsManagedWorkflow_FromAwsManagedWorkflowAttributes(scope constructs.Construct, id *string, attrs *AwsManagedWorkflowAttributes) IWorkflow {
	_init_.Initialize()

	if err := validateAwsManagedWorkflow_FromAwsManagedWorkflowAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		"fromAwsManagedWorkflowAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports the test-container AWS-managed workflow.
// Experimental.
func AwsManagedWorkflow_TestContainer(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAwsManagedWorkflow_TestContainerParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		"testContainer",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports the test-image AWS-managed workflow.
// Experimental.
func AwsManagedWorkflow_TestImage(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAwsManagedWorkflow_TestImageParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsManagedWorkflow",
		"testImage",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

