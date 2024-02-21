package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Lambda function runtime environment.
//
// If you need to use a runtime name that doesn't exist as a static member, you
// can instantiate a `Runtime` object, e.g: `new Runtime('nodejs99.99')`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   signingProfile := signer.NewSigningProfile(this, jsii.String("SigningProfile"), &SigningProfileProps{
//   	Platform: signer.Platform_AWS_LAMBDA_SHA384_ECDSA(),
//   })
//
//   codeSigningConfig := lambda.NewCodeSigningConfig(this, jsii.String("CodeSigningConfig"), &CodeSigningConfigProps{
//   	SigningProfiles: []iSigningProfile{
//   		signingProfile,
//   	},
//   })
//
//   lambda.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	CodeSigningConfig: CodeSigningConfig,
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   })
//
type Runtime interface {
	// The bundling Docker image for this runtime.
	BundlingImage() awscdk.DockerImage
	// The runtime family.
	Family() RuntimeFamily
	// Enabled for runtime enums that always target the latest available.
	IsVariable() *bool
	// The name of this runtime, as expected by the Lambda resource.
	Name() *string
	// Whether this runtime is integrated with and supported for profiling using Amazon CodeGuru Profiler.
	SupportsCodeGuruProfiling() *bool
	// Whether the ``ZipFile`` (aka inline code) property can be used with this runtime.
	SupportsInlineCode() *bool
	// Whether this runtime supports snapstart.
	SupportsSnapStart() *bool
	RuntimeEquals(other Runtime) *bool
	ToString() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) BundlingImage() awscdk.DockerImage {
	var returns awscdk.DockerImage
	_jsii_.Get(
		j,
		"bundlingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) Family() RuntimeFamily {
	var returns RuntimeFamily
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) IsVariable() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) SupportsCodeGuruProfiling() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsCodeGuruProfiling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) SupportsInlineCode() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsInlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) SupportsSnapStart() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsSnapStart",
		&returns,
	)
	return returns
}


func NewRuntime(name *string, family RuntimeFamily, props *LambdaRuntimeProps) Runtime {
	_init_.Initialize()

	if err := validateNewRuntimeParameters(name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Runtime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Runtime",
		[]interface{}{name, family, props},
		&j,
	)

	return &j
}

func NewRuntime_Override(r Runtime, name *string, family RuntimeFamily, props *LambdaRuntimeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Runtime",
		[]interface{}{name, family, props},
		r,
	)
}

func Runtime_ALL() *[]Runtime {
	_init_.Initialize()
	var returns *[]Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"ALL",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_6",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_8",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_1",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_2",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_2_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_2_1",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_3_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_3_1",
		&returns,
	)
	return returns
}

func Runtime_FROM_IMAGE() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"FROM_IMAGE",
		&returns,
	)
	return returns
}

func Runtime_GO_1_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"GO_1_X",
		&returns,
	)
	return returns
}

func Runtime_JAVA_11() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_11",
		&returns,
	)
	return returns
}

func Runtime_JAVA_17() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_17",
		&returns,
	)
	return returns
}

func Runtime_JAVA_21() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_21",
		&returns,
	)
	return returns
}

func Runtime_JAVA_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_8",
		&returns,
	)
	return returns
}

func Runtime_JAVA_8_CORRETTO() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_8_CORRETTO",
		&returns,
	)
	return returns
}

func Runtime_NODEJS() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_10_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_10_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_12_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_12_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_14_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_14_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_16_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_16_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_18_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_18_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_20_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_20_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_4_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_4_3",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_6_10() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_6_10",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_8_10() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_8_10",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_LATEST() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_LATEST",
		&returns,
	)
	return returns
}

func Runtime_PROVIDED() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PROVIDED",
		&returns,
	)
	return returns
}

func Runtime_PROVIDED_AL2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PROVIDED_AL2",
		&returns,
	)
	return returns
}

func Runtime_PROVIDED_AL2023() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PROVIDED_AL2023",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_2_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_2_7",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_10() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_10",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_11() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_11",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_12() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_12",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_6",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_7",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_8",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_9() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_9",
		&returns,
	)
	return returns
}

func Runtime_RUBY_2_5() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"RUBY_2_5",
		&returns,
	)
	return returns
}

func Runtime_RUBY_2_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"RUBY_2_7",
		&returns,
	)
	return returns
}

func Runtime_RUBY_3_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"RUBY_3_2",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Runtime) RuntimeEquals(other Runtime) *bool {
	if err := r.validateRuntimeEqualsParameters(other); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		r,
		"runtimeEquals",
		[]interface{}{other},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Runtime) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

