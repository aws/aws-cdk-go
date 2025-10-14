package awsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StoredQuery.
// Experimental.
type IStoredQueryRef interface {
	constructs.IConstruct
	// A reference to a StoredQuery resource.
	// Experimental.
	StoredQueryRef() *StoredQueryReference
}

// The jsii proxy for IStoredQueryRef
type jsiiProxy_IStoredQueryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStoredQueryRef) StoredQueryRef() *StoredQueryReference {
	var returns *StoredQueryReference
	_jsii_.Get(
		j,
		"storedQueryRef",
		&returns,
	)
	return returns
}

