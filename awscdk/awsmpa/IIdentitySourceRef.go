package awsmpa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmpa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentitySource.
// Experimental.
type IIdentitySourceRef interface {
	constructs.IConstruct
	// A reference to a IdentitySource resource.
	// Experimental.
	IdentitySourceRef() *IdentitySourceReference
}

// The jsii proxy for IIdentitySourceRef
type jsiiProxy_IIdentitySourceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIdentitySourceRef) IdentitySourceRef() *IdentitySourceReference {
	var returns *IdentitySourceReference
	_jsii_.Get(
		j,
		"identitySourceRef",
		&returns,
	)
	return returns
}

