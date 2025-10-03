package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ACL.
// Experimental.
type IACLRef interface {
	constructs.IConstruct
}

// The jsii proxy for IACLRef
type jsiiProxy_IACLRef struct {
	internal.Type__constructsIConstruct
}

