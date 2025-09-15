package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddonSubscription.
// Experimental.
type IMailManagerAddonSubscriptionRef interface {
	constructs.IConstruct
	// A reference to a MailManagerAddonSubscription resource.
	// Experimental.
	MailManagerAddonSubscriptionRef() *MailManagerAddonSubscriptionReference
}

// The jsii proxy for IMailManagerAddonSubscriptionRef
type jsiiProxy_IMailManagerAddonSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMailManagerAddonSubscriptionRef) MailManagerAddonSubscriptionRef() *MailManagerAddonSubscriptionReference {
	var returns *MailManagerAddonSubscriptionReference
	_jsii_.Get(
		j,
		"mailManagerAddonSubscriptionRef",
		&returns,
	)
	return returns
}

