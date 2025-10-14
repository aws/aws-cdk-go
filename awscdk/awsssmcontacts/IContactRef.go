package awsssmcontacts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Contact.
// Experimental.
type IContactRef interface {
	constructs.IConstruct
	// A reference to a Contact resource.
	// Experimental.
	ContactRef() *ContactReference
}

// The jsii proxy for IContactRef
type jsiiProxy_IContactRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IContactRef) ContactRef() *ContactReference {
	var returns *ContactReference
	_jsii_.Get(
		j,
		"contactRef",
		&returns,
	)
	return returns
}

