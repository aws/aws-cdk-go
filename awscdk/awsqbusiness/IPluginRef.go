package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Plugin.
// Experimental.
type IPluginRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPluginRef
type jsiiProxy_IPluginRef struct {
	internal.Type__constructsIConstruct
}

