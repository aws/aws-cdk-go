package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactList.
// Experimental.
type IContactListRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContactListRef
type jsiiProxy_IContactListRef struct {
	internal.Type__constructsIConstruct
}

