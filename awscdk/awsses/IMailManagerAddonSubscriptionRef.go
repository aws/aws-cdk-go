package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddonSubscription.
// Experimental.
type IMailManagerAddonSubscriptionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MailManagerAddonSubscription resource.
	// Experimental.
	MailManagerAddonSubscriptionRef() *MailManagerAddonSubscriptionReference
}

// The jsii proxy for IMailManagerAddonSubscriptionRef
type jsiiProxy_IMailManagerAddonSubscriptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IMailManagerAddonSubscriptionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMailManagerAddonSubscriptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

