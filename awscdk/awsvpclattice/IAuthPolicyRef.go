package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AuthPolicy.
// Experimental.
type IAuthPolicyRef interface {
	constructs.IConstruct
	// A reference to a AuthPolicy resource.
	// Experimental.
	AuthPolicyRef() *AuthPolicyReference
}

// The jsii proxy for IAuthPolicyRef
type jsiiProxy_IAuthPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAuthPolicyRef) AuthPolicyRef() *AuthPolicyReference {
	var returns *AuthPolicyReference
	_jsii_.Get(
		j,
		"authPolicyRef",
		&returns,
	)
	return returns
}

