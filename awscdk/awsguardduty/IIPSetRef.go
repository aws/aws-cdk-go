package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPSet.
// Experimental.
type IIPSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIPSetRef
type jsiiProxy_IIPSetRef struct {
	internal.Type__constructsIConstruct
}

