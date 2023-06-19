package awscodebuild

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
)

// A CodeBuild GPU image running Linux.
//
// This class has public constants that represent the most popular GPU images from AWS Deep Learning Containers.
//
// Example:
//   codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	Environment: &BuildEnvironment{
//   		BuildImage: codebuild.LinuxGpuBuildImage_DLC_TENSORFLOW_2_1_0_INFERENCE(),
//   	},
//   })
//
// See: https://aws.amazon.com/releasenotes/available-deep-learning-containers-images
//
// Experimental.
type LinuxGpuBuildImage interface {
	IBindableBuildImage
	// The default {@link ComputeType} to use with this image, if one was not specified in {@link BuildEnvironment#computeType} explicitly.
	// Experimental.
	DefaultComputeType() ComputeType
	// The Docker image identifier that the build environment uses.
	// Experimental.
	ImageId() *string
	// The type of principal that CodeBuild will use to pull this build Docker image.
	// Experimental.
	ImagePullPrincipalType() ImagePullPrincipalType
	// The type of build environment.
	// Experimental.
	Type() *string
	// Function that allows the build image access to the construct tree.
	// Experimental.
	Bind(scope awscdk.Construct, project IProject, _options *BuildImageBindOptions) *BuildImageConfig
	// Make a buildspec to run the indicated script.
	// Experimental.
	RunScriptBuildspec(entrypoint *string) BuildSpec
	// Allows the image a chance to validate whether the passed configuration is correct.
	// Experimental.
	Validate(buildEnvironment *BuildEnvironment) *[]*string
}

// The jsii proxy struct for LinuxGpuBuildImage
type jsiiProxy_LinuxGpuBuildImage struct {
	jsiiProxy_IBindableBuildImage
}

func (j *jsiiProxy_LinuxGpuBuildImage) DefaultComputeType() ComputeType {
	var returns ComputeType
	_jsii_.Get(
		j,
		"defaultComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxGpuBuildImage) ImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxGpuBuildImage) ImagePullPrincipalType() ImagePullPrincipalType {
	var returns ImagePullPrincipalType
	_jsii_.Get(
		j,
		"imagePullPrincipalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxGpuBuildImage) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Returns a Linux GPU build image from AWS Deep Learning Containers.
// See: https://aws.amazon.com/releasenotes/available-deep-learning-containers-images
//
// Experimental.
func LinuxGpuBuildImage_AwsDeepLearningContainersImage(repositoryName *string, tag *string, account *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxGpuBuildImage_AwsDeepLearningContainersImageParameters(repositoryName, tag); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"awsDeepLearningContainersImage",
		[]interface{}{repositoryName, tag, account},
		&returns,
	)

	return returns
}

// Returns a GPU image running Linux from an ECR repository.
//
// NOTE: if the repository is external (i.e. imported), then we won't be able to add
// a resource policy statement for it so CodeBuild can pull the image.
// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-ecr.html
//
// Experimental.
func LinuxGpuBuildImage_FromEcrRepository(repository awsecr.IRepository, tag *string) IBuildImage {
	_init_.Initialize()

	if err := validateLinuxGpuBuildImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns IBuildImage

	_jsii_.StaticInvoke(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

func LinuxGpuBuildImage_DLC_MXNET_1_4_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_MXNET_1_4_1",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_MXNET_1_6_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_MXNET_1_6_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_2_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_2_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_3_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_3_1",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_4_0_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_4_0_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_4_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_4_0_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_5_0_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_5_0_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_PYTORCH_1_5_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_PYTORCH_1_5_0_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_14_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_14_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_15_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_15_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_15_2_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_15_2_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_1_15_2_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_1_15_2_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_0_0() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_0_0",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_0_1() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_0_1",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_1_0_INFERENCE() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_1_0_INFERENCE",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_1_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_1_0_TRAINING",
		&returns,
	)
	return returns
}

func LinuxGpuBuildImage_DLC_TENSORFLOW_2_2_0_TRAINING() IBuildImage {
	_init_.Initialize()
	var returns IBuildImage
	_jsii_.StaticGet(
		"monocdk.aws_codebuild.LinuxGpuBuildImage",
		"DLC_TENSORFLOW_2_2_0_TRAINING",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxGpuBuildImage) Bind(scope awscdk.Construct, project IProject, _options *BuildImageBindOptions) *BuildImageConfig {
	if err := l.validateBindParameters(scope, project, _options); err != nil {
		panic(err)
	}
	var returns *BuildImageConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{scope, project, _options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxGpuBuildImage) RunScriptBuildspec(entrypoint *string) BuildSpec {
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

func (l *jsiiProxy_LinuxGpuBuildImage) Validate(buildEnvironment *BuildEnvironment) *[]*string {
	if err := l.validateValidateParameters(buildEnvironment); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		[]interface{}{buildEnvironment},
		&returns,
	)

	return returns
}

