package awss3deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a source for bucket deployments.
type ISource interface {
	// Binds the source to a bucket deployment.
	Bind(scope constructs.Construct, context *DeploymentSourceContext) *SourceConfig
}

// The jsii proxy for ISource
type jsiiProxy_ISource struct {
	_ byte // padding
}

func (i *jsiiProxy_ISource) Bind(scope constructs.Construct, context *DeploymentSourceContext) *SourceConfig {
	if err := i.validateBindParameters(scope, context); err != nil {
		panic(err)
	}
	var returns *SourceConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, context},
		&returns,
	)

	return returns
}

