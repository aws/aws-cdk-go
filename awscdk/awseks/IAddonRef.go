package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Addon.
// Experimental.
type IAddonRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAddonRef
type jsiiProxy_IAddonRef struct {
	internal.Type__constructsIConstruct
}

