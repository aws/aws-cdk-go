package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Policy.
// Experimental.
type IPolicyRef interface {
	constructs.IConstruct
	// A reference to a Policy resource.
	// Experimental.
	PolicyRef() *PolicyReference
}

// The jsii proxy for IPolicyRef
type jsiiProxy_IPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPolicyRef) PolicyRef() *PolicyReference {
	var returns *PolicyReference
	_jsii_.Get(
		j,
		"policyRef",
		&returns,
	)
	return returns
}

