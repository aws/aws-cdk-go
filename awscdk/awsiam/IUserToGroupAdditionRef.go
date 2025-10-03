package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserToGroupAddition.
// Experimental.
type IUserToGroupAdditionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IUserToGroupAdditionRef
type jsiiProxy_IUserToGroupAdditionRef struct {
	internal.Type__constructsIConstruct
}

