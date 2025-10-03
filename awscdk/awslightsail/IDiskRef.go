package awslightsail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Disk.
// Experimental.
type IDiskRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDiskRef
type jsiiProxy_IDiskRef struct {
	internal.Type__constructsIConstruct
}

