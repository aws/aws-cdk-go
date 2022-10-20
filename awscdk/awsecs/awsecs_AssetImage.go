package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
)

// An image that will be built from a local directory with a Dockerfile.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import ec2 "github.com/aws/aws-cdk-go/awscdk"
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//   import ecsPatterns "github.com/aws/aws-cdk-go/awscdk"
//   import cxapi "github.com/aws/aws-cdk-go/awscdk"
//   import path "github.com/aws-samples/dummy/path"
//
//   app := awscdk.NewApp()
//
//   stack := awscdk.NewStack(app, jsii.String("aws-ecs-patterns-queue"))
//   stack.node.setContext(cxapi.eCS_REMOVE_DEFAULT_DESIRED_COUNT, jsii.Boolean(true))
//
//   vpc := ec2.NewVpc(stack, jsii.String("VPC"), &vpcProps{
//   	maxAzs: jsii.Number(2),
//   })
//
//   ecsPatterns.NewQueueProcessingFargateService(stack, jsii.String("QueueProcessingService"), &queueProcessingFargateServiceProps{
//   	vpc: vpc,
//   	memoryLimitMiB: jsii.Number(512),
//   	image: ecs.NewAssetImage(path.join(__dirname, jsii.String(".."), jsii.String("sqs-reader"))),
//   })
//
// Experimental.
type AssetImage interface {
	ContainerImage
	// Called when the image is used by a ContainerDefinition.
	// Experimental.
	Bind(scope awscdk.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for AssetImage
type jsiiProxy_AssetImage struct {
	jsiiProxy_ContainerImage
}

// Constructs a new instance of the AssetImage class.
// Experimental.
func NewAssetImage(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	if err := validateNewAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetImage{}

	_jsii_.Create(
		"monocdk.aws_ecs.AssetImage",
		[]interface{}{directory, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AssetImage class.
// Experimental.
func NewAssetImage_Override(a AssetImage, directory *string, props *AssetImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.AssetImage",
		[]interface{}{directory, props},
		a,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
// Experimental.
func AssetImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	if err := validateAssetImage_FromAssetParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AssetImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
// Experimental.
func AssetImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	if err := validateAssetImage_FromDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AssetImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
// Experimental.
func AssetImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	if err := validateAssetImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns EcrImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AssetImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
// Experimental.
func AssetImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	if err := validateAssetImage_FromRegistryParameters(name, props); err != nil {
		panic(err)
	}
	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AssetImage",
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
func AssetImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	if err := validateAssetImage_FromTarballParameters(tarballFile); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.AssetImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetImage) Bind(scope awscdk.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	if err := a.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

