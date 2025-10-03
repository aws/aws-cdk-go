package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThreatIntelSet.
// Experimental.
type IThreatIntelSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IThreatIntelSetRef
type jsiiProxy_IThreatIntelSetRef struct {
	internal.Type__constructsIConstruct
}

