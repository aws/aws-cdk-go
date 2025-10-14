package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a User.
// Experimental.
type IUserRef interface {
	constructs.IConstruct
	// A reference to a User resource.
	// Experimental.
	UserRef() *UserReference
}

// The jsii proxy for IUserRef
type jsiiProxy_IUserRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserRef) UserRef() *UserReference {
	var returns *UserReference
	_jsii_.Get(
		j,
		"userRef",
		&returns,
	)
	return returns
}

