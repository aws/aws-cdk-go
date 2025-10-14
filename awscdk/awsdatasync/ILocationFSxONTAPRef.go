package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxONTAP.
// Experimental.
type ILocationFSxONTAPRef interface {
	constructs.IConstruct
	// A reference to a LocationFSxONTAP resource.
	// Experimental.
	LocationFSxOntapRef() *LocationFSxONTAPReference
}

// The jsii proxy for ILocationFSxONTAPRef
type jsiiProxy_ILocationFSxONTAPRef struct {
	internal.Type__constructsIConstruct
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

