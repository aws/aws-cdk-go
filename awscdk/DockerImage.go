package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Docker image.
//
// Example:
//   entry := "/path/to/function"
//   image := awscdk.DockerImage_FromBuild(entry)
//
//   python.NewPythonFunction(this, jsii.String("function"), &PythonFunctionProps{
//   	Entry: jsii.String(Entry),
//   	Runtime: awscdk.Runtime_PYTHON_3_8(),
//   	Bundling: &BundlingOptions{
//   		BuildArgs: map[string]*string{
//   			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
//   			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
//   		},
//   	},
//   })
//
type DockerImage interface {
	// The Docker image.
	Image() *string
	// Build the image, if it is built from a Docker file and hasn't been built yet.
	//
	// Does nothing for other images, such as images from a registry. Running the
	// image or copying files out of it calls this automatically. Call it yourself
	// before you pass `image` to your own `docker` commands.
	BuildImage()
	// Copies a file or directory out of the Docker image to the local filesystem.
	//
	// If `outputPath` is omitted the destination path is a temporary directory.
	//
	// Builds the image first if it is built on demand and hasn't been built yet.
	//
	// Returns: the destination path.
	Cp(imagePath *string, outputPath *string) *string
	// Runs a Docker image.
	//
	// Builds the image first if it is built on demand and hasn't been built yet.
	Run(options *DockerRunOptions)
	// Provides a stable representation of this image for JSON serialization.
	//
	// Returns: The overridden image name if set or image hash name in that order.
	ToJSON() *string
}

// The jsii proxy struct for DockerImage
type jsiiProxy_DockerImage struct {
	_ byte // padding
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


func NewDockerImage(image *string, _imageHash *string) DockerImage {
	_init_.Initialize()

	if err := validateNewDockerImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_DockerImage{}

	_jsii_.Create(
		"aws-cdk-lib.DockerImage",
		[]interface{}{image, _imageHash},
		&j,
	)

	return &j
}

func NewDockerImage_Override(d DockerImage, image *string, _imageHash *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DockerImage",
		[]interface{}{image, _imageHash},
		d,
	)
}

// Builds a Docker image.
//
// The image is built right away, even if it is never used.
//
// Use `DockerImage.fromBuildV2()` instead. It builds the same image under the
// same tag, but only when the image is first needed, so assets whose bundled
// output already exists, or that skip bundling, never pay for the build. If
// you need the image right away, call `buildImage()` on it.
func DockerImage_FromBuild(path *string, options *DockerBuildOptions) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerImage",
		"fromBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference a Docker image that is built from a Docker file when it is first needed.
//
// The image is built when `buildImage()` is called, when it first runs, or when
// files are first copied out of it. Asset bundling builds the image right
// before it runs the bundling container, so assets whose bundled output
// already exists, or that skip bundling, never build it.
//
// Until the image is built, `image` names a tag that may not exist in the
// local Docker daemon yet. Call `buildImage()` before you pass the tag to
// your own `docker` commands. Errors in the Docker file are reported when the
// image is built, for example while an asset is bundled, not when this
// method is called.
//
// The tag and the `docker build` command are computed from the Docker file
// directory and the build options when this method is called, so later
// changes to the options have no effect. Don't change the Docker file
// directory afterwards, because the image would no longer match its tag.
//
// Images with the same tag are only built or checked once until the next
// `App` is created, unless they use `buildContexts`. After that, the local
// Docker daemon is checked again.
//
// The image has the same tag and hash as one built with `DockerImage.fromBuild()`
// from the same arguments, so switching between the two doesn't change asset hashes.
func DockerImage_FromBuildV2(path *string, options *DockerBuildOptions) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromBuildV2Parameters(path, options); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerImage",
		"fromBuildV2",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func DockerImage_FromRegistry(image *string) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromRegistryParameters(image); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerImage",
		"fromRegistry",
		[]interface{}{image},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) BuildImage() {
	_jsii_.InvokeVoid(
		d,
		"buildImage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DockerImage) Cp(imagePath *string, outputPath *string) *string {
	if err := d.validateCpParameters(imagePath); err != nil {
		panic(err)
	}
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
	if err := d.validateRunParameters(options); err != nil {
		panic(err)
	}
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

