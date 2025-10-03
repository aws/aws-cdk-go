package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SchemaMapping.
// Experimental.
type ISchemaMappingRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISchemaMappingRef
type jsiiProxy_ISchemaMappingRef struct {
	internal.Type__constructsIConstruct
}

