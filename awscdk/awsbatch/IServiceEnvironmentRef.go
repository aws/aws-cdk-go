package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServiceEnvironment.
// Experimental.
type IServiceEnvironmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ServiceEnvironment resource.
	// Experimental.
	ServiceEnvironmentRef() *ServiceEnvironmentReference
}

// The jsii proxy for IServiceEnvironmentRef
type jsiiProxy_IServiceEnvironmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IServiceEnvironmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServiceEnvironmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

