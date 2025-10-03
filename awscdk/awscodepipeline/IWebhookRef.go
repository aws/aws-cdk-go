package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodepipeline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Webhook.
// Experimental.
type IWebhookRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWebhookRef
type jsiiProxy_IWebhookRef struct {
	internal.Type__constructsIConstruct
}

