package interfacesawselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ServerlessCache.
// Experimental.
type IServerlessCacheRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ServerlessCache resource.
	// Experimental.
	ServerlessCacheRef() *ServerlessCacheReference
}

// The jsii proxy for IServerlessCacheRef
type jsiiProxy_IServerlessCacheRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IServerlessCacheRef) ServerlessCacheRef() *ServerlessCacheReference {
	var returns *ServerlessCacheReference
	_jsii_.Get(
		j,
		"serverlessCacheRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCacheRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCacheRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

