package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailAddress.
// Experimental.
type IEmailAddressRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEmailAddressRef
type jsiiProxy_IEmailAddressRef struct {
	internal.Type__constructsIConstruct
}

