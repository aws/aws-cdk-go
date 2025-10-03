package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Schema.
// Experimental.
type ISchemaRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISchemaRef
type jsiiProxy_ISchemaRef struct {
	internal.Type__constructsIConstruct
}

