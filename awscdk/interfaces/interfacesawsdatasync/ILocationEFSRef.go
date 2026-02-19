package interfacesawsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationEFS.
// Experimental.
type ILocationEFSRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocationEFS resource.
	// Experimental.
	LocationEfsRef() *LocationEFSReference
}

// The jsii proxy for ILocationEFSRef
type jsiiProxy_ILocationEFSRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILocationEFSRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILocationEFSRef) LocationEfsRef() *LocationEFSReference {
	var returns *LocationEFSReference
	_jsii_.Get(
		j,
		"locationEfsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationEFSRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationEFSRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

