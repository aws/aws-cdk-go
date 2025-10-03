package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Feature.
// Experimental.
type IFeatureRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFeatureRef
type jsiiProxy_IFeatureRef struct {
	internal.Type__constructsIConstruct
}

