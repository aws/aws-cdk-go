package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Launch.
// Experimental.
type ILaunchRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILaunchRef
type jsiiProxy_ILaunchRef struct {
	internal.Type__constructsIConstruct
}

