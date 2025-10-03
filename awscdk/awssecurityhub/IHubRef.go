package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Hub.
// Experimental.
type IHubRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHubRef
type jsiiProxy_IHubRef struct {
	internal.Type__constructsIConstruct
}

