package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a User.
// Experimental.
type IUserRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserRef
type jsiiProxy_IUserRef struct {
	internal.Type__constructsIConstruct
}

