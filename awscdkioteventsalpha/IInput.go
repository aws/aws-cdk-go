package awscdkioteventsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2/internal"
)

// Represents an AWS IoT Events input.
// Experimental.
type IInput interface {
	awscdk.IResource
	// Grant the indicated permissions on this input to the given IAM principal (Role/Group/User).
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant write permissions on this input and its contents to an IAM principal (Role/Group/User).
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the input.
	// Experimental.
	InputArn() *string
	// The name of the input.
	// Experimental.
	InputName() *string
}

// The jsii proxy for IInput
type jsiiProxy_IInput struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IInput) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IInput) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IInput) InputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInput) InputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputName",
		&returns,
	)
	return returns
}

