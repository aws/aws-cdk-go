package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

// An image from an Amazon ECR repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImageAsset dockerImageAsset
//
//   ecrImage := awscdk.Aws_ecs.EcrImage_FromDockerImageAsset(dockerImageAsset)
//
type EcrImage interface {
	ContainerImage
	// The image name. Images in Amazon ECR repositories can be specified by either using the full registry/repository:tag or registry/repository@digest.
	//
	// For example, 012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>:latest or
	// 012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE.
	ImageName() *string
	// Called when the image is used by a ContainerDefinition.
	Bind(_scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
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
func NewEcrImage(repository awsecr.IRepository, tagOrDigest *string) EcrImage {
	_init_.Initialize()

	if err := validateNewEcrImageParameters(repository, tagOrDigest); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.EcrImage",
		[]interface{}{repository, tagOrDigest},
		&j,
	)

	return &j
}

// Constructs a new instance of the EcrImage class.
func NewEcrImage_Override(e EcrImage, repository awsecr.IRepository, tagOrDigest *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.EcrImage",
		[]interface{}{repository, tagOrDigest},
		e,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func EcrImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	if err := validateEcrImage_FromAssetParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func EcrImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	if err := validateEcrImage_FromDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func EcrImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	if err := validateEcrImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func EcrImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	if err := validateEcrImage_FromRegistryParameters(name, props); err != nil {
		panic(err)
	}
	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
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
func EcrImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	if err := validateEcrImage_FromTarballParameters(tarballFile); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.EcrImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrImage) Bind(_scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	if err := e.validateBindParameters(_scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_scope, containerDefinition},
		&returns,
	)

	return returns
}

