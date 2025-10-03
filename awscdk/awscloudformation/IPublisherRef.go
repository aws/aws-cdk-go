package awscloudformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Publisher.
// Experimental.
type IPublisherRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPublisherRef
type jsiiProxy_IPublisherRef struct {
	internal.Type__constructsIConstruct
}

