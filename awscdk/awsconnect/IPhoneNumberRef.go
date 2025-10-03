package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PhoneNumber.
// Experimental.
type IPhoneNumberRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPhoneNumberRef
type jsiiProxy_IPhoneNumberRef struct {
	internal.Type__constructsIConstruct
}

