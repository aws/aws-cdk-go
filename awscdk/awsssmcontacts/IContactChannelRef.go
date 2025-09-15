package awsssmcontacts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactChannel.
// Experimental.
type IContactChannelRef interface {
	constructs.IConstruct
	// A reference to a ContactChannel resource.
	// Experimental.
	ContactChannelRef() *ContactChannelReference
}

// The jsii proxy for IContactChannelRef
type jsiiProxy_IContactChannelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IContactChannelRef) ContactChannelRef() *ContactChannelReference {
	var returns *ContactChannelReference
	_jsii_.Get(
		j,
		"contactChannelRef",
		&returns,
	)
	return returns
}

