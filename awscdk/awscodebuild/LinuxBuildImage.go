package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
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
type LinuxBuildImage interface {
	IBuildImage
	// The default `ComputeType` to use with this image, if one was not specified in `BuildEnvironment#computeType` explicitly.
	DefaultComputeType() ComputeType
	// The Docker image identifier that the build environment uses.
	ImageId() *string
	// The type of principal that CodeBuild will use to pull this build Docker image.
	ImagePullPrincipalType() ImagePullPrincipalType
	// An optional ECR repository that the image is hosted in.
	Repository() awsecr.IRepository
	// The secretsManagerCredentials for access to a private registry.
	SecretsManagerCredentials() awssecretsmanager.ISecret
	// The type of build environment.
	Type() *string
	// Make a buildspec to run the indicated script.
	RunScriptBuildspec(entrypoint *string) BuildSpec
	// Allows the image a chance to validate whether the passed configuration is correct.
	Validate(env *BuildEnvironment) *[]*string
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
func LinuxBuildImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
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
func LinuxBuildImage_FromCodeBuildImageId(id *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromCodeBuildImageIdParameters(id); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"fromCodeBuildImageId",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Returns: a x86-64 Linux build image from a Docker Hub image.
func LinuxBuildImage_FromDockerRegistry(name *string, options *DockerImageOptions) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromDockerRegistryParameters(name, options); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
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
func LinuxBuildImage_FromEcrRepository(repository awsecr.IRepository, tagOrDigest *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxBuildImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
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
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_3() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_3",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_4() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_4",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_5() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_5",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_ARM() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_ARM",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_ARM_2() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_ARM_2",
		&returns,
	)
	return returns
}

func LinuxBuildImage_AMAZON_LINUX_2_ARM_3() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"AMAZON_LINUX_2_ARM_3",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_1_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_1_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_2_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_3_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_3_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_4_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_4_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_5_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_5_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_6_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_6_0",
		&returns,
	)
	return returns
}

func LinuxBuildImage_STANDARD_7_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_codebuild.LinuxBuildImage",
		"STANDARD_7_0",
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

func (l *jsiiProxy_LinuxBuildImage) Validate(env *BuildEnvironment) *[]*string {
	if err := l.validateValidateParameters(env); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		[]interface{}{env},
		&returns,
	)

	return returns
}

