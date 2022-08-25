package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
)

// An image from an Amazon ECR repository.
//
// Example:
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//
//   repo := ecr.repository.fromRepositoryName(this, jsii.String("batch-job-repo"), jsii.String("todo-list"))
//
//   batch.NewJobDefinition(this, jsii.String("batch-job-def-from-ecr"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.NewEcrImage(repo, jsii.String("latest")),
//   	},
//   })
//
// Experimental.
type EcrImage interface {
	ContainerImage
	// The image name. Images in Amazon ECR repositories can be specified by either using the full registry/repository:tag or registry/repository@digest.
	//
	// For example, 012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>:latest or
	// 012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE.
	// Experimental.
	ImageName() *string
	// Called when the image is used by a ContainerDefinition.
	// Experimental.
	Bind(_scope awscdk.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for EcrImage
type jsiiProxy_EcrImage struct {
	jsiiProxy_ContainerImage
}

func (j *jsiiProxy_EcrImage) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}


// Constructs a new instance of the EcrImage class.
// Experimental.
func NewEcrImage(repository awsecr.IRepository, tagOrDigest *string) EcrImage {
	_init_.Initialize()

	j := jsiiProxy_EcrImage{}

	_jsii_.Create(
		"monocdk.aws_ecs.EcrImage",
		[]interface{}{repository, tagOrDigest},
		&j,
	)

	return &j
}

// Constructs a new instance of the EcrImage class.
// Experimental.
func NewEcrImage_Override(e EcrImage, repository awsecr.IRepository, tagOrDigest *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.EcrImage",
		[]interface{}{repository, tagOrDigest},
		e,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
// Experimental.
func EcrImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	var returns AssetImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcrImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
// Experimental.
func EcrImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcrImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
// Experimental.
func EcrImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	var returns EcrImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcrImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
// Experimental.
func EcrImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcrImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
// Experimental.
func EcrImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	var returns ContainerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.EcrImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrImage) Bind(_scope awscdk.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope, containerDefinition},
		&returns,
	)

	return returns
}

