package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddressList.
// Experimental.
type IMailManagerAddressListRef interface {
	constructs.IConstruct
	// A reference to a MailManagerAddressList resource.
	// Experimental.
	MailManagerAddressListRef() *MailManagerAddressListReference
}

// The jsii proxy for IMailManagerAddressListRef
type jsiiProxy_IMailManagerAddressListRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMailManagerAddressListRef) MailManagerAddressListRef() *MailManagerAddressListReference {
	var returns *MailManagerAddressListReference
	_jsii_.Get(
		j,
		"mailManagerAddressListRef",
		&returns,
	)
	return returns
}

