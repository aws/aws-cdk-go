package interfacesawsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxLustre.
// Experimental.
type ILocationFSxLustreRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocationFSxLustre resource.
	// Experimental.
	LocationFSxLustreRef() *LocationFSxLustreReference
}

// The jsii proxy for ILocationFSxLustreRef
type jsiiProxy_ILocationFSxLustreRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILocationFSxLustreRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILocationFSxLustreRef) LocationFSxLustreRef() *LocationFSxLustreReference {
	var returns *LocationFSxLustreReference
	_jsii_.Get(
		j,
		"locationFSxLustreRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxLustreRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxLustreRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

