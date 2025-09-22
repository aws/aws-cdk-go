package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddonInstance.
// Experimental.
type IMailManagerAddonInstanceRef interface {
	constructs.IConstruct
	// A reference to a MailManagerAddonInstance resource.
	// Experimental.
	MailManagerAddonInstanceRef() *MailManagerAddonInstanceReference
}

// The jsii proxy for IMailManagerAddonInstanceRef
type jsiiProxy_IMailManagerAddonInstanceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMailManagerAddonInstanceRef) MailManagerAddonInstanceRef() *MailManagerAddonInstanceReference {
	var returns *MailManagerAddonInstanceReference
	_jsii_.Get(
		j,
		"mailManagerAddonInstanceRef",
		&returns,
	)
	return returns
}

