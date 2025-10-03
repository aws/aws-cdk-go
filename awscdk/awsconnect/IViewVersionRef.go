package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ViewVersion.
// Experimental.
type IViewVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IViewVersionRef
type jsiiProxy_IViewVersionRef struct {
	internal.Type__constructsIConstruct
}

