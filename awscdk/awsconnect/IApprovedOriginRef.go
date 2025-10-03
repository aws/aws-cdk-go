package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApprovedOrigin.
// Experimental.
type IApprovedOriginRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApprovedOriginRef
type jsiiProxy_IApprovedOriginRef struct {
	internal.Type__constructsIConstruct
}

