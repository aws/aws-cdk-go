package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents an Amazon Bedrock abstraction on which you can run the `Invoke` API.
//
// This can be a Foundational Model,
// a Custom Model, or an Inference Profile.
// Experimental.
type IBedrockInvokable interface {
	// Gives the appropriate policies to invoke and use the invokable abstraction.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the Bedrock invokable abstraction.
	// Experimental.
	InvokableArn() *string
}

// The jsii proxy for IBedrockInvokable
type jsiiProxy_IBedrockInvokable struct {
	_ byte // padding
}

func (i *jsiiProxy_IBedrockInvokable) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IBedrockInvokable) InvokableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokableArn",
		&returns,
	)
	return returns
}

