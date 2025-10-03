package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyGrant.
// Experimental.
type IPolicyGrantRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPolicyGrantRef
type jsiiProxy_IPolicyGrantRef struct {
	internal.Type__constructsIConstruct
}

