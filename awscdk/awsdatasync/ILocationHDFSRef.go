package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationHDFS.
// Experimental.
type ILocationHDFSRef interface {
	constructs.IConstruct
	// A reference to a LocationHDFS resource.
	// Experimental.
	LocationHdfsRef() *LocationHDFSReference
}

// The jsii proxy for ILocationHDFSRef
type jsiiProxy_ILocationHDFSRef struct {
	internal.Type__constructsIConstruct
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

