package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Link.
// Experimental.
type ILinkRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILinkRef
type jsiiProxy_ILinkRef struct {
	internal.Type__constructsIConstruct
}

