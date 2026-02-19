package interfacesawsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationHDFS.
// Experimental.
type ILocationHDFSRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocationHDFS resource.
	// Experimental.
	LocationHdfsRef() *LocationHDFSReference
}

// The jsii proxy for ILocationHDFSRef
type jsiiProxy_ILocationHDFSRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILocationHDFSRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILocationHDFSRef) LocationHdfsRef() *LocationHDFSReference {
	var returns *LocationHDFSReference
	_jsii_.Get(
		j,
		"locationHdfsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationHDFSRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationHDFSRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

