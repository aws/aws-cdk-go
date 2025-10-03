package awsnetworkmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Site.
// Experimental.
type ISiteRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISiteRef
type jsiiProxy_ISiteRef struct {
	internal.Type__constructsIConstruct
}

