package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerArchive.
// Experimental.
type IMailManagerArchiveRef interface {
	constructs.IConstruct
	// A reference to a MailManagerArchive resource.
	// Experimental.
	MailManagerArchiveRef() *MailManagerArchiveReference
}

// The jsii proxy for IMailManagerArchiveRef
type jsiiProxy_IMailManagerArchiveRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMailManagerArchiveRef) MailManagerArchiveRef() *MailManagerArchiveReference {
	var returns *MailManagerArchiveReference
	_jsii_.Get(
		j,
		"mailManagerArchiveRef",
		&returns,
	)
	return returns
}

