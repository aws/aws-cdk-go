package awsbillingconductor

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BillingGroup.
// Experimental.
type IBillingGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IBillingGroupRef
type jsiiProxy_IBillingGroupRef struct {
	internal.Type__constructsIConstruct
}

