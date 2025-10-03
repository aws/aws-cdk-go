package awsssmcontacts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmcontacts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Contact.
// Experimental.
type IContactRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContactRef
type jsiiProxy_IContactRef struct {
	internal.Type__constructsIConstruct
}

