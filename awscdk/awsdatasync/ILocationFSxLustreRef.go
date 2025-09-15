package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationFSxLustre.
// Experimental.
type ILocationFSxLustreRef interface {
	constructs.IConstruct
	// A reference to a LocationFSxLustre resource.
	// Experimental.
	LocationFSxLustreRef() *LocationFSxLustreReference
}

// The jsii proxy for ILocationFSxLustreRef
type jsiiProxy_ILocationFSxLustreRef struct {
	internal.Type__constructsIConstruct
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

