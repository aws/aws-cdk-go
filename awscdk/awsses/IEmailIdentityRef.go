package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailIdentity.
// Experimental.
type IEmailIdentityRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEmailIdentityRef
type jsiiProxy_IEmailIdentityRef struct {
	internal.Type__constructsIConstruct
}

