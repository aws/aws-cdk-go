package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CodeBuild image running x86-64 Linux.
//
// This class has a bunch of public constants that represent the most popular images.
//
// You can also specify a custom image using one of the static methods:
//
// - LinuxBuildImage.fromDockerRegistry(image[, { secretsManagerCredentials }])
// - LinuxBuildImage.fromEcrRepository(repo[, tag])
// - LinuxBuildImage.fromAsset(parent, id, props)
//
// Example:
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &CodePipelineProps{
//   	Synth: pipelines.NewShellStep(jsii.String("Synth"), &ShellStepProps{
//   		Input: pipelines.CodePipelineSource_Connection(jsii.String("my-org/my-app"), jsii.String("main"), &ConnectionSourceOptions{
//   			ConnectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		Commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//
//   	// Turn this on because the pipeline uses Docker image assets
//   	DockerEnabledForSelfMutation: jsii.Boolean(true),
//   })
//
//   pipeline.AddWave(jsii.String("MyWave"), &WaveOptions{
//   	Post: []step{
//   		pipelines.NewCodeBuildStep(jsii.String("RunApproval"), &CodeBuildStepProps{
//   			Commands: []*string{
//   				jsii.String("command-from-image"),
//   			},
//   			BuildEnvironment: &BuildEnvironment{
//   				// The user of a Docker image asset in the pipeline requires turning on
//   				// 'dockerEnabledForSelfMutation'.
//   				BuildImage: codebuild.LinuxBuildImage_FromAsset(this, jsii.String("Image"), &DockerImageAssetProps{
//   					Directory: jsii.String("./docker-image"),
//   				}),
//   			},
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
// Experimental.
type LinuxBuildImage interface {
	IBuildImage
	// The default {@link ComputeType} to use with this image, if one was not specified in {@link BuildEnvironment#computeType} explicitly.
	// Experimental.
	DefaultComputeType() ComputeType
	// The Docker image identifier that the build environment uses.
	// Experimental.
	ImageId() *string
	// The type of principal that CodeBuild will use to pull this build Docker image.
	// Experimental.
	ImagePullPrincipalType() ImagePullPrincipalType
	// An optional ECR repository that the image is hosted in.
	// Experimental.
	Repository() awsecr.IRepository
	// The secretsManagerCredentials for access to a private registry.
	// Experimental.
	SecretsManagerCredentials() awssecretsmanager.ISecret
	// The type of build environment.
	// Experimental.
	Type() *string
	// Make a buildspec to run the indicated script.
	// Experimental.
	RunScriptBuildspec(entrypoint *string) BuildSpec
	// Allows the image a chance to validate whether the passed configuration is correct.
	// Experimental.
	Validate(_arg *BuildEnvironment) *[]*string
}

// The jsii proxy struct for LinuxBuildImage
type jsiiProxy_LinuxBuildImage struct {
	jsiiProxy_IBuildImage
}

func (j *jsiiProxy_LinuxBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) SecretsManagerCredentials() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secretsManagerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Uses an Docker image asset as a x86-64 Linux build image.
// Experimental.
func LinuxBuildImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"fromAsset",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Uses a Docker image provided by CodeBuild.
//
// Returns: A Docker image provided by CodeBuild.
//
// Example:
//   "aws/codebuild/standard:4.0"
//
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-env-ref-available.html
//
// Experimental.
func LinuxBuildImage_FromCodeBuildImageId(id *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromCodeBuildImageIdParameters(id); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"fromCodeBuildImageId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Returns: a x86-64 Linux build image from a Docker Hub image.
// Experimental.
func LinuxBuildImage_FromDockerRegistry(name *string, options *DockerImageOptions) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromDockerRegistryParameters(name, options); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"fromDockerRegistry",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

// Returns: A x86-64 Linux build image from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
// Experimental.
func LinuxBuildImage_FromEcrRepository(repository awsecr.IRepository, tagOrDigest *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tagOrDigest},
		&returns,
	)

	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_3() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_3",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_ARM() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_ARM",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_ARM_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_ARM_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_1_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"STANDARD_1_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"STANDARD_2_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_3_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"STANDARD_3_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_4_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"STANDARD_4_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_5_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"STANDARD_5_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_ANDROID_JAVA8_24_4_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_ANDROID_JAVA8_24_4_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_ANDROID_JAVA8_26_1_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_ANDROID_JAVA8_26_1_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_BASE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_BASE",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_DOCKER_17_09_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_DOCKER_17_09_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_DOCKER_18_09_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_DOCKER_18_09_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_DOTNET_CORE_1_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_DOTNET_CORE_1_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_DOTNET_CORE_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_DOTNET_CORE_2_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_DOTNET_CORE_2_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_DOTNET_CORE_2_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_GOLANG_1_10() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_GOLANG_1_10",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_GOLANG_1_11() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_GOLANG_1_11",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_NODEJS_10_1_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_NODEJS_10_1_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_NODEJS_10_14_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_NODEJS_10_14_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_NODEJS_6_3_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_NODEJS_6_3_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_NODEJS_8_11_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_NODEJS_8_11_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_OPEN_JDK_11() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_OPEN_JDK_11",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_OPEN_JDK_8() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_OPEN_JDK_8",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_OPEN_JDK_9() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_OPEN_JDK_9",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PHP_5_6() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PHP_5_6",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PHP_7_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PHP_7_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PHP_7_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PHP_7_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PYTHON_2_7_12() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PYTHON_2_7_12",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PYTHON_3_3_6() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PYTHON_3_3_6",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PYTHON_3_4_5() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PYTHON_3_4_5",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PYTHON_3_5_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PYTHON_3_5_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PYTHON_3_6_5() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PYTHON_3_6_5",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_PYTHON_3_7_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_PYTHON_3_7_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_RUBY_2_2_5() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_RUBY_2_2_5",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_RUBY_2_3_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_RUBY_2_3_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_RUBY_2_5_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_RUBY_2_5_1",
		&returns,
	)
	return returns
}

func LinuxBuildImage_UBUNTU_14_04_RUBY_2_5_3() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxBuildImage",
		"UBUNTU_14_04_RUBY_2_5_3",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
	if err := l.validateRunScriptBuildspecParameters(entrypoint); err != nil {
		panic(err)
	}
	var returns BuildSpec

	_jsii_.Invoke(
		l,
		"runScriptBuildspec",
		[]interface{}{entrypoint},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxBuildImage) Validate(_arg *BuildEnvironment) *[]*string {
	if err := l.validateValidateParameters(_arg); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

