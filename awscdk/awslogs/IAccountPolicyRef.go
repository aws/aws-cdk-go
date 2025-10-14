package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccountPolicy.
// Experimental.
type IAccountPolicyRef interface {
	constructs.IConstruct
	// A reference to a AccountPolicy resource.
	// Experimental.
	AccountPolicyRef() *AccountPolicyReference
}

// The jsii proxy for IAccountPolicyRef
type jsiiProxy_IAccountPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccountPolicyRef) AccountPolicyRef() *AccountPolicyReference {
	var returns *AccountPolicyReference
	_jsii_.Get(
		j,
		"accountPolicyRef",
		&returns,
	)
	return returns
}

