package awspinpointemail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Identity.
// Experimental.
type IIdentityRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdentityRef
type jsiiProxy_IIdentityRef struct {
	internal.Type__constructsIConstruct
}

