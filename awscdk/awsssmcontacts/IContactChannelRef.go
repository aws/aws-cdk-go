package awsssmcontacts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactChannel.
// Experimental.
type IContactChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContactChannelRef
type jsiiProxy_IContactChannelRef struct {
	internal.Type__constructsIConstruct
}

