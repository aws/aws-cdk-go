package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the function's source code.
//
// Example:
//   var s3Bucket bucket
//   // Add a cloudfront Function to a Distribution
//   cfFunction := cloudfront.NewFunction(this, jsii.String("Function"), &FunctionProps{
//   	Code: cloudfront.FunctionCode_FromInline(jsii.String("function handler(event) { return event.request }")),
//   })
//   cloudfront.NewDistribution(this, jsii.String("distro"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(s3Bucket),
//   		FunctionAssociations: []functionAssociation{
//   			&functionAssociation{
//   				Function: cfFunction,
//   				EventType: cloudfront.FunctionEventType_VIEWER_REQUEST,
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type FunctionCode interface {
	// renders the function code.
	// Experimental.
	Render() *string
}

// The jsii proxy struct for FunctionCode
type jsiiProxy_FunctionCode struct {
	_ byte // padding
}

// Experimental.
func NewFunctionCode_Override(f FunctionCode) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_cloudfront.FunctionCode",
		nil, // no parameters
		f,
	)
}

// Code from external file for function.
//
// Returns: code object with contents from file.
// Experimental.
func FunctionCode_FromFile(options *FileCodeOptions) FunctionCode {
	_init_.Initialize()

	if err := validateFunctionCode_FromFileParameters(options); err != nil {
		panic(err)
	}
	var returns FunctionCode

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.FunctionCode",
		"fromFile",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Inline code for function.
//
// Returns: code object with inline code.
// Experimental.
func FunctionCode_FromInline(code *string) FunctionCode {
	_init_.Initialize()

	if err := validateFunctionCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns FunctionCode

	_jsii_.StaticInvoke(
		"monocdk.aws_cloudfront.FunctionCode",
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

