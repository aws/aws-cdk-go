package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdNamespace.
// Experimental.
type IIdNamespaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdNamespaceRef
type jsiiProxy_IIdNamespaceRef struct {
	internal.Type__constructsIConstruct
}

