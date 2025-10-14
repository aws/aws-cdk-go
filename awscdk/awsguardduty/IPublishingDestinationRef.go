package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublishingDestination.
// Experimental.
type IPublishingDestinationRef interface {
	constructs.IConstruct
	// A reference to a PublishingDestination resource.
	// Experimental.
	PublishingDestinationRef() *PublishingDestinationReference
}

// The jsii proxy for IPublishingDestinationRef
type jsiiProxy_IPublishingDestinationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPublishingDestinationRef) PublishingDestinationRef() *PublishingDestinationReference {
	var returns *PublishingDestinationReference
	_jsii_.Get(
		j,
		"publishingDestinationRef",
		&returns,
	)
	return returns
}

