package awskafkaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomPlugin.
// Experimental.
type ICustomPluginRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomPluginRef
type jsiiProxy_ICustomPluginRef struct {
	internal.Type__constructsIConstruct
}

