package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Extension.
// Experimental.
type IExtensionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IExtensionRef
type jsiiProxy_IExtensionRef struct {
	internal.Type__constructsIConstruct
}

