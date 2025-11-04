package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxOpenZFS.
// Experimental.
type ILocationFSxOpenZFSRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LocationFSxOpenZFS resource.
	// Experimental.
	LocationFSxOpenZfsRef() *LocationFSxOpenZFSReference
}

// The jsii proxy for ILocationFSxOpenZFSRef
type jsiiProxy_ILocationFSxOpenZFSRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILocationFSxOpenZFSRef) LocationFSxOpenZfsRef() *LocationFSxOpenZFSReference {
	var returns *LocationFSxOpenZFSReference
	_jsii_.Get(
		j,
		"locationFSxOpenZfsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxOpenZFSRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationFSxOpenZFSRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

