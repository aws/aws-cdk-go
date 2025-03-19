package awscdkpipesalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Source interface.
// Experimental.
type ISource interface {
	// Bind the source to a pipe.
	// Experimental.
	Bind(pipe IPipe) *SourceConfig
	// Grant the pipe role read access to the source.
	// Experimental.
	GrantRead(grantee awsiam.IRole)
	// The ARN of the source resource.
	// Experimental.
	SourceArn() *string
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISource) Bind(pipe IPipe) *SourceConfig {
	if err := i.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ISource) GrantRead(grantee awsiam.IRole) {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantRead",
		[]interface{}{grantee},
	)
}

func (j *jsiiProxy_ISource) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

