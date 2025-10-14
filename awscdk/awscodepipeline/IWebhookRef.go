package awscodepipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Webhook.
// Experimental.
type IWebhookRef interface {
	constructs.IConstruct
	// A reference to a Webhook resource.
	// Experimental.
	WebhookRef() *WebhookReference
}

// The jsii proxy for IWebhookRef
type jsiiProxy_IWebhookRef struct {
	internal.Type__constructsIConstruct
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

