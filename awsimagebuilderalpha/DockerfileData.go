package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for referencing and uploading dockerfile data for the container recipe.
//
// Example:
//   containerRecipe := imagebuilder.NewContainerRecipe(this, jsii.String("CustomDockerfileContainerRecipe"), &ContainerRecipeProps{
//   	BaseImage: imagebuilder.BaseContainerImage_*FromDockerHub(jsii.String("amazonlinux"), jsii.String("latest")),
//   	TargetRepository: imagebuilder.Repository_*FromEcr(ecr.Repository_*FromRepositoryName(this, jsii.String("Repository"), jsii.String("my-container-repo"))),
//   	Dockerfile: imagebuilder.DockerfileData_FromInline(jsii.String(`
//   	FROM {{{ imagebuilder:parentImage }}}
//   	CMD ["echo", "Hello, world!"]
//   	{{{ imagebuilder:environments }}}
//   	{{{ imagebuilder:components }}}
//   	`)),
//   })
//
// Experimental.
type DockerfileData interface {
	// The rendered Dockerfile value, for use in CloudFormation.
	//
	// - For inline dockerfiles, dockerfileTemplateData is the Dockerfile template text
	// - For S3-backed dockerfiles, dockerfileTemplateUri is the S3 URL.
	// Experimental.
	Render() *DockerfileTemplateConfig
}

// The jsii proxy struct for DockerfileData
type jsiiProxy_DockerfileData struct {
	_ byte // padding
}

// Experimental.
func NewDockerfileData_Override(d DockerfileData) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.DockerfileData",
		nil, // no parameters
		d,
	)
}

// Uploads dockerfile data from a local file to S3 to use as the dockerfile data.
// Experimental.
func DockerfileData_FromAsset(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) S3DockerfileData {
	_init_.Initialize()

	if err := validateDockerfileData_FromAssetParameters(scope, id, path, options); err != nil {
		panic(err)
	}
	var returns S3DockerfileData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DockerfileData",
		"fromAsset",
		[]interface{}{scope, id, path, options},
		&returns,
	)

	return returns
}

// Uses an inline string as the dockerfile data.
// Experimental.
func DockerfileData_FromInline(data *string) DockerfileData {
	_init_.Initialize()

	if err := validateDockerfileData_FromInlineParameters(data); err != nil {
		panic(err)
	}
	var returns DockerfileData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DockerfileData",
		"fromInline",
		[]interface{}{data},
		&returns,
	)

	return returns
}

// References dockerfile data from a pre-existing S3 object.
// Experimental.
func DockerfileData_FromS3(bucket awss3.IBucket, key *string) S3DockerfileData {
	_init_.Initialize()

	if err := validateDockerfileData_FromS3Parameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3DockerfileData

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.DockerfileData",
		"fromS3",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerfileData) Render() *DockerfileTemplateConfig {
	var returns *DockerfileTemplateConfig

	_jsii_.Invoke(
		d,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

