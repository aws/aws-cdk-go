// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Docker image used for asset bundling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   bundlingDockerImage := monocdk.BundlingDockerImage_FromAsset(jsii.String("path"), &DockerBuildOptions{
//   	BuildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	File: jsii.String("file"),
//   	Platform: jsii.String("platform"),
//   })
//
// Deprecated: use DockerImage.
type BundlingDockerImage interface {
	// The Docker image.
	// Deprecated: use DockerImage.
	Image() *string
	// Copies a file or directory out of the Docker image to the local filesystem.
	//
	// If `outputPath` is omitted the destination path is a temporary directory.
	//
	// Returns: the destination path.
	// Deprecated: use DockerImage.
	Cp(imagePath *string, outputPath *string) *string
	// Runs a Docker image.
	// Deprecated: use DockerImage.
	Run(options *DockerRunOptions)
	// Provides a stable representation of this image for JSON serialization.
	//
	// Returns: The overridden image name if set or image hash name in that order.
	// Deprecated: use DockerImage.
	ToJSON() *string
}

// The jsii proxy struct for BundlingDockerImage
type jsiiProxy_BundlingDockerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_BundlingDockerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Deprecated: use DockerImage.
func NewBundlingDockerImage(image *string, _imageHash *string) BundlingDockerImage {
	_init_.Initialize()

	if err := validateNewBundlingDockerImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_BundlingDockerImage{}

	_jsii_.Create(
		"monocdk.BundlingDockerImage",
		[]interface{}{image, _imageHash},
		&j,
	)

	return &j
}

// Deprecated: use DockerImage.
func NewBundlingDockerImage_Override(b BundlingDockerImage, image *string, _imageHash *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.BundlingDockerImage",
		[]interface{}{image, _imageHash},
		b,
	)
}

// Reference an image that's built directly from sources on disk.
// Deprecated: use DockerImage.fromBuild()
func BundlingDockerImage_FromAsset(path *string, options *DockerBuildOptions) BundlingDockerImage {
	_init_.Initialize()

	if err := validateBundlingDockerImage_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns BundlingDockerImage

	_jsii_.StaticInvoke(
		"monocdk.BundlingDockerImage",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
// Deprecated: use DockerImage.
func BundlingDockerImage_FromRegistry(image *string) DockerImage {
	_init_.Initialize()

	if err := validateBundlingDockerImage_FromRegistryParameters(image); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"monocdk.BundlingDockerImage",
		"fromRegistry",
		[]interface{}{image},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BundlingDockerImage) Cp(imagePath *string, outputPath *string) *string {
	if err := b.validateCpParameters(imagePath); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"cp",
		[]interface{}{imagePath, outputPath},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BundlingDockerImage) Run(options *DockerRunOptions) {
	if err := b.validateRunParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"run",
		[]interface{}{options},
	)
}

func (b *jsiiProxy_BundlingDockerImage) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

