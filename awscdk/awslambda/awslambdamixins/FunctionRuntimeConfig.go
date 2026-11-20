package awslambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/awslambdamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Lambda-specific Mixin to run a function on a customer-managed runtime (CMR).
//
// Sets the `runtimeConfig` property on the function, replacing any existing
// `runtimeConfig`. The function's runtime must be `custom`, and the function
// must not be a container image function; both are checked at synthesis, so a
// runtime or package type set after the mixin is applied is still checked. A
// runtime that is only known at deployment, such as a parameter value, is not
// checked.
//
// Which runtime versions and architectures the runtime supports, and whether
// the calling account may use it, are validated by the service at deployment.
//
// Example:
//   lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.NewRuntime(jsii.String("custom"), lambda.RuntimeFamily_OTHER),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   }).With(mixins.NewFunctionRuntimeConfig(&FunctionRuntimeConfigProps{
//   	RuntimeArn: jsii.String("arn:aws:lambda:us-east-1:123456789012:runtime/my-python-runtime"),
//   	RuntimeAlias: jsii.String("stable"),
//   }))
//
type FunctionRuntimeConfig interface {
	awscdk.Mixin
	// Applies the mixin functionality to the target construct.
	ApplyTo(construct constructs.IConstruct)
	// Determines whether this mixin can be applied to the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for FunctionRuntimeConfig
type jsiiProxy_FunctionRuntimeConfig struct {
	internal.Type__awscdkMixin
}

func NewFunctionRuntimeConfig(props *FunctionRuntimeConfigProps) FunctionRuntimeConfig {
	_init_.Initialize()

	if err := validateNewFunctionRuntimeConfigParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionRuntimeConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.mixins.FunctionRuntimeConfig",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewFunctionRuntimeConfig_Override(f FunctionRuntimeConfig, props *FunctionRuntimeConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.mixins.FunctionRuntimeConfig",
		[]interface{}{props},
		f,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func FunctionRuntimeConfig_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFunctionRuntimeConfig_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.mixins.FunctionRuntimeConfig",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionRuntimeConfig) ApplyTo(construct constructs.IConstruct) {
	if err := f.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyTo",
		[]interface{}{construct},
	)
}

func (f *jsiiProxy_FunctionRuntimeConfig) Supports(construct constructs.IConstruct) *bool {
	if err := f.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		f,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

