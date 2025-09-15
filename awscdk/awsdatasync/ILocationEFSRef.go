package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationEFS.
// Experimental.
type ILocationEFSRef interface {
	constructs.IConstruct
	// A reference to a LocationEFS resource.
	// Experimental.
	LocationEfsRef() *LocationEFSReference
}

// The jsii proxy for ILocationEFSRef
type jsiiProxy_ILocationEFSRef struct {
	internal.Type__constructsIConstruct
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

