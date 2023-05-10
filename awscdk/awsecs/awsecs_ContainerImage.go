package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
)

// Constructs for types of container images.
//
// Example:
//   var vpc vpc
//
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCPCluster"), &ClusterProps{
//   	Vpc: Vpc,
//   	EnableFargateCapacityProviders: jsii.Boolean(true),
//   })
//
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.AddContainer(jsii.String("web"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   })
//
//   ecs.NewFargateService(this, jsii.String("FargateService"), &FargateServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	CapacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			CapacityProvider: jsii.String("FARGATE_SPOT"),
//   			Weight: jsii.Number(2),
//   		},
//   		&capacityProviderStrategy{
//   			CapacityProvider: jsii.String("FARGATE"),
//   			Weight: jsii.Number(1),
//   		},
//   	},
//   })
//
// Experimental.
type ContainerImage interface {
	// Called when the image is used by a ContainerDefinition.
	// Experimental.
	Bind(scope awscdk.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for ContainerImage
type jsiiProxy_ContainerImage struct {
	_ byte // padding
}

// Experimental.
func NewContainerImage_Override(c ContainerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecs.ContainerImage",
		nil, // no parameters
		c,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
// Experimental.
func ContainerImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	if err := validateContainerImage_FromAssetParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ContainerImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
// Experimental.
func ContainerImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	if err := validateContainerImage_FromDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ContainerImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
// Experimental.
func ContainerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	if err := validateContainerImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns EcrImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ContainerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
// Experimental.
func ContainerImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	if err := validateContainerImage_FromRegistryParameters(name, props); err != nil {
		panic(err)
	}
	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ContainerImage",
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
func ContainerImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	if err := validateContainerImage_FromTarballParameters(tarballFile); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_ecs.ContainerImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImage) Bind(scope awscdk.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	if err := c.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

