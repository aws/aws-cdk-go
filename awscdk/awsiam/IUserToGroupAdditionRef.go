package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserToGroupAddition.
// Experimental.
type IUserToGroupAdditionRef interface {
	constructs.IConstruct
	// A reference to a UserToGroupAddition resource.
	// Experimental.
	UserToGroupAdditionRef() *UserToGroupAdditionReference
}

// The jsii proxy for IUserToGroupAdditionRef
type jsiiProxy_IUserToGroupAdditionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserToGroupAdditionRef) UserToGroupAdditionRef() *UserToGroupAdditionReference {
	var returns *UserToGroupAdditionReference
	_jsii_.Get(
		j,
		"userToGroupAdditionRef",
		&returns,
	)
	return returns
}

