package interfacesawscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodepipeline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Webhook.
// Experimental.
type IWebhookRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Webhook resource.
	// Experimental.
	WebhookRef() *WebhookReference
}

// The jsii proxy for IWebhookRef
type jsiiProxy_IWebhookRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IWebhookRef) WebhookRef() *WebhookReference {
	var returns *WebhookReference
	_jsii_.Get(
		j,
		"webhookRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebhookRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebhookRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

