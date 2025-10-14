package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxOpenZFS.
// Experimental.
type ILocationFSxOpenZFSRef interface {
	constructs.IConstruct
	// A reference to a LocationFSxOpenZFS resource.
	// Experimental.
	LocationFSxOpenZfsRef() *LocationFSxOpenZFSReference
}

// The jsii proxy for ILocationFSxOpenZFSRef
type jsiiProxy_ILocationFSxOpenZFSRef struct {
	internal.Type__constructsIConstruct
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

