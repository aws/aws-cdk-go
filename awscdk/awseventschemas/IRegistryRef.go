package awseventschemas

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Registry.
// Experimental.
type IRegistryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRegistryRef
type jsiiProxy_IRegistryRef struct {
	internal.Type__constructsIConstruct
}

