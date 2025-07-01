package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates `IDockerImage` instances.
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
type DockerImage interface {
	// Called when the image is used by a SageMaker task.
	Bind(task ISageMakerTask) *DockerImageConfig
}

// The jsii proxy struct for DockerImage
type jsiiProxy_DockerImage struct {
	_ byte // padding
}

func NewDockerImage_Override(d DockerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.DockerImage",
		nil, // no parameters
		d,
	)
}

// Reference a Docker image that is provided as an Asset in the current app.
func DockerImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DockerImage",
		"fromAsset",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Reference a Docker image stored in an ECR repository.
func DockerImage_FromEcrRepository(repository awsecr.IRepository, tagOrDigest *string) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DockerImage",
		"fromEcrRepository",
		[]interface{}{repository, tagOrDigest},
		&returns,
	)

	return returns
}

// Reference a Docker image which URI is obtained from the task's input.
func DockerImage_FromJsonExpression(expression *string, allowAnyEcrImagePull *bool) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromJsonExpressionParameters(expression); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DockerImage",
		"fromJsonExpression",
		[]interface{}{expression, allowAnyEcrImagePull},
		&returns,
	)

	return returns
}

// Reference a Docker image by it's URI.
//
// When referencing ECR images, prefer using `inEcr`.
func DockerImage_FromRegistry(imageUri *string) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromRegistryParameters(imageUri); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DockerImage",
		"fromRegistry",
		[]interface{}{imageUri},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Bind(task ISageMakerTask) *DockerImageConfig {
	if err := d.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *DockerImageConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

