package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PhoneNumber.
// Experimental.
type IPhoneNumberRef interface {
	constructs.IConstruct
	// A reference to a PhoneNumber resource.
	// Experimental.
	PhoneNumberRef() *PhoneNumberReference
}

// The jsii proxy for IPhoneNumberRef
type jsiiProxy_IPhoneNumberRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPhoneNumberRef) PhoneNumberRef() *PhoneNumberReference {
	var returns *PhoneNumberReference
	_jsii_.Get(
		j,
		"phoneNumberRef",
		&returns,
	)
	return returns
}

