package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddonInstance.
// Experimental.
type IMailManagerAddonInstanceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerAddonInstanceRef
type jsiiProxy_IMailManagerAddonInstanceRef struct {
	internal.Type__constructsIConstruct
}

