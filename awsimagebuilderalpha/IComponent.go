package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/internal"
)

// An EC2 Image Builder Component.
// Experimental.
type IComponent interface {
	awscdk.IResource
	// Grant custom actions to the given grantee for the component.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions to the given grantee for the component.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the component.
	// Experimental.
	ComponentArn() *string
	// The name of the component.
	// Experimental.
	ComponentName() *string
	// The version of the component.
	// Experimental.
	ComponentVersion() *string
}

// The jsii proxy for IComponent
type jsiiProxy_IComponent struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IComponent) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IComponent) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IComponent) ComponentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponent) ComponentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponent) ComponentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentVersion",
		&returns,
	)
	return returns
}

