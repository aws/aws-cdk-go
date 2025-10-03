package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddressList.
// Experimental.
type IMailManagerAddressListRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerAddressListRef
type jsiiProxy_IMailManagerAddressListRef struct {
	internal.Type__constructsIConstruct
}

