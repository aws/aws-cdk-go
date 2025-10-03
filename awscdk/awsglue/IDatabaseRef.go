package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Database.
// Experimental.
type IDatabaseRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDatabaseRef
type jsiiProxy_IDatabaseRef struct {
	internal.Type__constructsIConstruct
}

