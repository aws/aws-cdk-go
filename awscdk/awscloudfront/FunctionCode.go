package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the function's source code.
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
type FunctionCode interface {
	// renders the function code.
	Render() *string
}

// The jsii proxy struct for FunctionCode
type jsiiProxy_FunctionCode struct {
	_ byte // padding
}

func NewFunctionCode_Override(f FunctionCode) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		nil, // no parameters
		f,
	)
}

// Code from external file for function.
//
// Returns: code object with contents from file.
func FunctionCode_FromFile(options *FileCodeOptions) FunctionCode {
	_init_.Initialize()

	if err := validateFunctionCode_FromFileParameters(options); err != nil {
		panic(err)
	}
	var returns FunctionCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		"fromFile",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Inline code for function.
//
// Returns: code object with inline code.
func FunctionCode_FromInline(code *string) FunctionCode {
	_init_.Initialize()

	if err := validateFunctionCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns FunctionCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.FunctionCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionCode) Render() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

