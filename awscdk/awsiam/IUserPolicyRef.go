package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPolicy.
// Experimental.
type IUserPolicyRef interface {
	constructs.IConstruct
	// A reference to a UserPolicy resource.
	// Experimental.
	UserPolicyRef() *UserPolicyReference
}

// The jsii proxy for IUserPolicyRef
type jsiiProxy_IUserPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IUserPolicyRef) UserPolicyRef() *UserPolicyReference {
	var returns *UserPolicyReference
	_jsii_.Get(
		j,
		"userPolicyRef",
		&returns,
	)
	return returns
}

