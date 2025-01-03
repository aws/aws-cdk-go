package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The function's runtime environment version.
//
// Example:
//   store := cloudfront.NewKeyValueStore(this, jsii.String("KeyValueStore"))
//   cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
//   	// Note that JS_2_0 must be used for Key Value Store support
//   	Runtime: cloudfront.FunctionRuntime_JS_2_0(),
//   	KeyValueStore: store,
//   })
//
type FunctionRuntime interface {
	Value() *string
}

// The jsii proxy struct for FunctionRuntime
type jsiiProxy_FunctionRuntime struct {
	_ byte // padding
}

func (j *jsiiProxy_FunctionRuntime) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A custom runtime string.
//
// Gives full control over the runtime string fragment.
func FunctionRuntime_Custom(runtimeString *string) FunctionRuntime {
	_init_.Initialize()

	if err := validateFunctionRuntime_CustomParameters(runtimeString); err != nil {
		panic(err)
	}
	var returns FunctionRuntime

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.FunctionRuntime",
		"custom",
		[]interface{}{runtimeString},
		&returns,
	)

	return returns
}

func FunctionRuntime_JS_1_0() FunctionRuntime {
	_init_.Initialize()
	var returns FunctionRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.FunctionRuntime",
		"JS_1_0",
		&returns,
	)
	return returns
}

func FunctionRuntime_JS_2_0() FunctionRuntime {
	_init_.Initialize()
	var returns FunctionRuntime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudfront.FunctionRuntime",
		"JS_2_0",
		&returns,
	)
	return returns
}

