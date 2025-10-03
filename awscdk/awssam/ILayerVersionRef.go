package awssam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LayerVersion.
// Experimental.
type ILayerVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILayerVersionRef
type jsiiProxy_ILayerVersionRef struct {
	internal.Type__constructsIConstruct
}

