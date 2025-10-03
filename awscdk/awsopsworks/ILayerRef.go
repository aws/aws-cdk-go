package awsopsworks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Layer.
// Experimental.
type ILayerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILayerRef
type jsiiProxy_ILayerRef struct {
	internal.Type__constructsIConstruct
}

