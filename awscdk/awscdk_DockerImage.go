// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Docker image.
//
// Example:
//   entry := "/path/to/function"
//   image := awscdk.DockerImage.fromBuild(entry)
//
//   lambda.NewPythonFunction(this, jsii.String("function"), &pythonFunctionProps{
//   	entry: jsii.String(entry),
//   	runtime: awscdk.Runtime_PYTHON_3_8(),
//   	bundling: &bundlingOptions{
//   		buildArgs: map[string]*string{
//   			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
//   			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
//   		},
//   	},
//   })
//
// Experimental.
type DockerImage interface {
	BundlingDockerImage
	// The Docker image.
	// Experimental.
	Image() *string
	// Copies a file or directory out of the Docker image to the local filesystem.
	//
	// If `outputPath` is omitted the destination path is a temporary directory.
	//
	// Returns: the destination path.
	// Experimental.
	Cp(imagePath *string, outputPath *string) *string
	// Runs a Docker image.
	// Experimental.
	Run(options *DockerRunOptions)
	// Provides a stable representation of this image for JSON serialization.
	//
	// Returns: The overridden image name if set or image hash name in that order.
	// Experimental.
	ToJSON() *string
}

// The jsii proxy struct for DockerImage
type jsiiProxy_DockerImage struct {
	jsiiProxy_BundlingDockerImage
}

func (j *jsiiProxy_DockerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


// Experimental.
func NewDockerImage(image *string, _imageHash *string) DockerImage {
	_init_.Initialize()

	j := jsiiProxy_DockerImage{}

	_jsii_.Create(
		"monocdk.DockerImage",
		[]interface{}{image, _imageHash},
		&j,
	)

	return &j
}

// Experimental.
func NewDockerImage_Override(d DockerImage, image *string, _imageHash *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.DockerImage",
		[]interface{}{image, _imageHash},
		d,
	)
}

// Reference an image that's built directly from sources on disk.
// Deprecated: use DockerImage.fromBuild()
func DockerImage_FromAsset(path *string, options *DockerBuildOptions) BundlingDockerImage {
	_init_.Initialize()

	var returns BundlingDockerImage

	_jsii_.StaticInvoke(
		"monocdk.DockerImage",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Builds a Docker image.
// Experimental.
func DockerImage_FromBuild(path *string, options *DockerBuildOptions) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"monocdk.DockerImage",
		"fromBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
// Experimental.
func DockerImage_FromRegistry(image *string) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"monocdk.DockerImage",
		"fromRegistry",
		[]interface{}{image},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Cp(imagePath *string, outputPath *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"cp",
		[]interface{}{imagePath, outputPath},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Run(options *DockerRunOptions) {
	_jsii_.InvokeVoid(
		d,
		"run",
		[]interface{}{options},
	)
}

func (d *jsiiProxy_DockerImage) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

