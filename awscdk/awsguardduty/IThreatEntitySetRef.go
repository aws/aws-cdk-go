package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThreatEntitySet.
// Experimental.
type IThreatEntitySetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IThreatEntitySetRef
type jsiiProxy_IThreatEntitySetRef struct {
	internal.Type__constructsIConstruct
}

