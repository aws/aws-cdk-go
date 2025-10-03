package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerRelay.
// Experimental.
type IMailManagerRelayRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerRelayRef
type jsiiProxy_IMailManagerRelayRef struct {
	internal.Type__constructsIConstruct
}

