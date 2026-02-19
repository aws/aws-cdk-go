package interfacesawsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerlessCluster.
// Experimental.
type IServerlessClusterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServerlessCluster resource.
	// Experimental.
	ServerlessClusterRef() *ServerlessClusterReference
}

// The jsii proxy for IServerlessClusterRef
type jsiiProxy_IServerlessClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IServerlessClusterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IServerlessClusterRef) ServerlessClusterRef() *ServerlessClusterReference {
	var returns *ServerlessClusterReference
	_jsii_.Get(
		j,
		"serverlessClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessClusterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessClusterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

