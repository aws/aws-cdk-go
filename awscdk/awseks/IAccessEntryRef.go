package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessEntry.
// Experimental.
type IAccessEntryRef interface {
	constructs.IConstruct
	// A reference to a AccessEntry resource.
	// Experimental.
	AccessEntryRef() *AccessEntryReference
}

// The jsii proxy for IAccessEntryRef
type jsiiProxy_IAccessEntryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessEntryRef) AccessEntryRef() *AccessEntryReference {
	var returns *AccessEntryReference
	_jsii_.Get(
		j,
		"accessEntryRef",
		&returns,
	)
	return returns
}

