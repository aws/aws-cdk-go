package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the function's source code.
//
// Example:
//   var s3Bucket bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &functionProps{
//   	code: cloudfront.functionCode.fromInline(jsii.String("function handler(event) { return event.request }")),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: origins.NewS3Origin(s3Bucket),
//   		functionAssociations: []functionAssociation{
//   			&functionAssociation{
//   				function: cfFunction,
//   				eventType: cloudfront.functionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
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

