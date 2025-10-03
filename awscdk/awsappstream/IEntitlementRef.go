package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Entitlement.
// Experimental.
type IEntitlementRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEntitlementRef
type jsiiProxy_IEntitlementRef struct {
	internal.Type__constructsIConstruct
}

