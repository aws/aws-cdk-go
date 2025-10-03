package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerArchive.
// Experimental.
type IMailManagerArchiveRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerArchiveRef
type jsiiProxy_IMailManagerArchiveRef struct {
	internal.Type__constructsIConstruct
}

