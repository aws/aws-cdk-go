package interfacesawsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MailManagerAddonSubscription.
// Experimental.
type IMailManagerAddonSubscriptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MailManagerAddonSubscription resource.
	// Experimental.
	MailManagerAddonSubscriptionRef() *MailManagerAddonSubscriptionReference
}

// The jsii proxy for IMailManagerAddonSubscriptionRef
type jsiiProxy_IMailManagerAddonSubscriptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMailManagerAddonSubscriptionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IMailManagerAddonSubscriptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

