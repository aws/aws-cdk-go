package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrustedEntitySet.
// Experimental.
type ITrustedEntitySetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrustedEntitySetRef
type jsiiProxy_ITrustedEntitySetRef struct {
	internal.Type__constructsIConstruct
}

