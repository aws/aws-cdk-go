package awspinpointemail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpointemail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Identity.
// Experimental.
type IIdentityRef interface {
	constructs.IConstruct
	// A reference to a Identity resource.
	// Experimental.
	IdentityRef() *IdentityReference
}

// The jsii proxy for IIdentityRef
type jsiiProxy_IIdentityRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIdentityRef) IdentityRef() *IdentityReference {
	var returns *IdentityReference
	_jsii_.Get(
		j,
		"identityRef",
		&returns,
	)
	return returns
}

