package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Target configuration.
// Experimental.
type ITarget interface {
	// Bind this target to a pipe.
	// Experimental.
	Bind(pipe IPipe) *TargetConfig
	// Grant the pipe role to push to the target.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
	// The ARN of the target resource.
	// Experimental.
	TargetArn() *string
}

// The jsii proxy for ITarget
type jsiiProxy_ITarget struct {
	_ byte // padding
}

func (i *jsiiProxy_ITarget) Bind(pipe IPipe) *TargetConfig {
	if err := i.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *TargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITarget) GrantPush(grantee awsiam.IRole) {
	if err := i.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantPush",
		[]interface{}{grantee},
	)
}

func (j *jsiiProxy_ITarget) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

