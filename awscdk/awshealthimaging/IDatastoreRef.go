package awshealthimaging

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awshealthimaging/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Datastore.
// Experimental.
type IDatastoreRef interface {
	constructs.IConstruct
	// A reference to a Datastore resource.
	// Experimental.
	DatastoreRef() *DatastoreReference
}

// The jsii proxy for IDatastoreRef
type jsiiProxy_IDatastoreRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDatastoreRef) DatastoreRef() *DatastoreReference {
	var returns *DatastoreReference
	_jsii_.Get(
		j,
		"datastoreRef",
		&returns,
	)
	return returns
}

