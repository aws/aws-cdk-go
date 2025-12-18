package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for working with Amazon-managed workflows.
//
// Example:
//   containerWorkflowPipeline := imagebuilder.NewImagePipeline(this, jsii.String("ContainerWorkflowPipeline"), &ImagePipelineProps{
//   	Recipe: exampleContainerRecipe,
//   	Workflows: []WorkflowConfiguration{
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AmazonManagedWorkflow_BuildContainer(this, jsii.String("BuildContainer")),
//   		},
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AmazonManagedWorkflow_TestContainer(this, jsii.String("TestContainer")),
//   		},
//   		&WorkflowConfiguration{
//   			Workflow: imagebuilder.AmazonManagedWorkflow_DistributeContainer(this, jsii.String("DistributeContainer")),
//   		},
//   	},
//   })
//
// Experimental.
type AmazonManagedWorkflow interface {
}

// The jsii proxy struct for AmazonManagedWorkflow
type jsiiProxy_AmazonManagedWorkflow struct {
	_ byte // padding
}

// Experimental.
func NewAmazonManagedWorkflow() AmazonManagedWorkflow {
	_init_.Initialize()

	j := jsiiProxy_AmazonManagedWorkflow{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAmazonManagedWorkflow_Override(a AmazonManagedWorkflow) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		nil, // no parameters
		a,
	)
}

// Imports the build-container Amazon-managed workflow.
// Experimental.
func AmazonManagedWorkflow_BuildContainer(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAmazonManagedWorkflow_BuildContainerParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		"buildContainer",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports the build-image AWS-managed workflow.
// Experimental.
func AmazonManagedWorkflow_BuildImage(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAmazonManagedWorkflow_BuildImageParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		"buildImage",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports the distribute-container AWS-managed workflow.
// Experimental.
func AmazonManagedWorkflow_DistributeContainer(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAmazonManagedWorkflow_DistributeContainerParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		"distributeContainer",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports the distribute-image AWS-managed workflow.
// Experimental.
func AmazonManagedWorkflow_DistributeImage(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAmazonManagedWorkflow_DistributeImageParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		"distributeImage",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports an AWS-managed workflow from its attributes.
// Experimental.
func AmazonManagedWorkflow_FromAmazonManagedWorkflowAttributes(scope constructs.Construct, id *string, attrs *AmazonManagedWorkflowAttributes) IWorkflow {
	_init_.Initialize()

	if err := validateAmazonManagedWorkflow_FromAmazonManagedWorkflowAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		"fromAmazonManagedWorkflowAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports the test-container AWS-managed workflow.
// Experimental.
func AmazonManagedWorkflow_TestContainer(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAmazonManagedWorkflow_TestContainerParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		"testContainer",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

// Imports the test-image AWS-managed workflow.
// Experimental.
func AmazonManagedWorkflow_TestImage(scope constructs.Construct, id *string) IWorkflow {
	_init_.Initialize()

	if err := validateAmazonManagedWorkflow_TestImageParameters(scope, id); err != nil {
		panic(err)
	}
	var returns IWorkflow

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AmazonManagedWorkflow",
		"testImage",
		[]interface{}{scope, id},
		&returns,
	)

	return returns
}

