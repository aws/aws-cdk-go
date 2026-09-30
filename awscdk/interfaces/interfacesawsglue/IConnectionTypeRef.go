package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectionType.
// Experimental.
type IConnectionTypeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConnectionType resource.
	// Experimental.
	ConnectionTypeRef() *ConnectionTypeReference
}

// The jsii proxy for IConnectionTypeRef
type jsiiProxy_IConnectionTypeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IConnectionTypeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IConnectionTypeRef) ConnectionTypeRef() *ConnectionTypeReference {
	var returns *ConnectionTypeReference
	_jsii_.Get(
		j,
		"connectionTypeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionTypeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionTypeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

