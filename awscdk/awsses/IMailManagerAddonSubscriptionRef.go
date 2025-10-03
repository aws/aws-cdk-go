package awsses

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddonSubscription.
// Experimental.
type IMailManagerAddonSubscriptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMailManagerAddonSubscriptionRef
type jsiiProxy_IMailManagerAddonSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

