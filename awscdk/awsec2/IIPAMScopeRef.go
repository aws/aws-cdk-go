package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMScope.
// Experimental.
type IIPAMScopeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIPAMScopeRef
type jsiiProxy_IIPAMScopeRef struct {
	internal.Type__constructsIConstruct
}

