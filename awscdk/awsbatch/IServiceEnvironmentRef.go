package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceEnvironment.
// Experimental.
type IServiceEnvironmentRef interface {
	constructs.IConstruct
	// A reference to a ServiceEnvironment resource.
	// Experimental.
	ServiceEnvironmentRef() *ServiceEnvironmentReference
}

// The jsii proxy for IServiceEnvironmentRef
type jsiiProxy_IServiceEnvironmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IServiceEnvironmentRef) ServiceEnvironmentRef() *ServiceEnvironmentReference {
	var returns *ServiceEnvironmentReference
	_jsii_.Get(
		j,
		"serviceEnvironmentRef",
		&returns,
	)
	return returns
}

