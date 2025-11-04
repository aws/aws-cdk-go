package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxONTAP.
// Experimental.
type ILocationFSxONTAPRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LocationFSxONTAP resource.
	// Experimental.
	LocationFSxOntapRef() *LocationFSxONTAPReference
}

// The jsii proxy for ILocationFSxONTAPRef
type jsiiProxy_ILocationFSxONTAPRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILocationFSxONTAPRef) LocationFSxOntapRef() *LocationFSxONTAPReference {
	var returns *LocationFSxONTAPReference
	_jsii_.Get(
		j,
		"locationFSxOntapRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxONTAPRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxONTAPRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

