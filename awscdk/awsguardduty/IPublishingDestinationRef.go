package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublishingDestination.
// Experimental.
type IPublishingDestinationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPublishingDestinationRef
type jsiiProxy_IPublishingDestinationRef struct {
	internal.Type__constructsIConstruct
}

