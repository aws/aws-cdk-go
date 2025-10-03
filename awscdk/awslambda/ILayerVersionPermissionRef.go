package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LayerVersionPermission.
// Experimental.
type ILayerVersionPermissionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILayerVersionPermissionRef
type jsiiProxy_ILayerVersionPermissionRef struct {
	internal.Type__constructsIConstruct
}

