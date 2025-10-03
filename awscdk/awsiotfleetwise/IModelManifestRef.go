package awsiotfleetwise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelManifest.
// Experimental.
type IModelManifestRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModelManifestRef
type jsiiProxy_IModelManifestRef struct {
	internal.Type__constructsIConstruct
}

