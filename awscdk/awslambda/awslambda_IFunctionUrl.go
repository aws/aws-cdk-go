package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda/internal"
)

// A Lambda function Url.
// Experimental.
type IFunctionUrl interface {
	awscdk.IResource
	// Grant the given identity permissions to invoke this Lambda Function URL.
	// Experimental.
	GrantInvokeUrl(identity awsiam.IGrantable) awsiam.Grant
	// The ARN of the function this URL refers to.
	// Experimental.
	FunctionArn() *string
	// The url of the Lambda function.
	// Experimental.
	Url() *string
}

// The jsii proxy for IFunctionUrl
type jsiiProxy_IFunctionUrl struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFunctionUrl) GrantInvokeUrl(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeUrlParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeUrl",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IFunctionUrl) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionUrl) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

