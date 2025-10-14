package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EmailIdentity.
// Experimental.
type IEmailIdentityRef interface {
	constructs.IConstruct
	// A reference to a EmailIdentity resource.
	// Experimental.
	EmailIdentityRef() *EmailIdentityReference
}

// The jsii proxy for IEmailIdentityRef
type jsiiProxy_IEmailIdentityRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEmailIdentityRef) EmailIdentityRef() *EmailIdentityReference {
	var returns *EmailIdentityReference
	_jsii_.Get(
		j,
		"emailIdentityRef",
		&returns,
	)
	return returns
}

