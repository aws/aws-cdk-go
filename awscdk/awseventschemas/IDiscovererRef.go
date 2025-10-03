package awseventschemas

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Discoverer.
// Experimental.
type IDiscovererRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDiscovererRef
type jsiiProxy_IDiscovererRef struct {
	internal.Type__constructsIConstruct
}

