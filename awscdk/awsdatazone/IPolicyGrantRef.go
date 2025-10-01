package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyGrant.
// Experimental.
type IPolicyGrantRef interface {
	constructs.IConstruct
	// A reference to a PolicyGrant resource.
	// Experimental.
	PolicyGrantRef() *PolicyGrantReference
}

// The jsii proxy for IPolicyGrantRef
type jsiiProxy_IPolicyGrantRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPolicyGrantRef) PolicyGrantRef() *PolicyGrantReference {
	var returns *PolicyGrantReference
	_jsii_.Get(
		j,
		"policyGrantRef",
		&returns,
	)
	return returns
}

