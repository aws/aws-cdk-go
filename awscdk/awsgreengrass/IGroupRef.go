package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Group.
// Experimental.
type IGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGroupRef
type jsiiProxy_IGroupRef struct {
	internal.Type__constructsIConstruct
}

