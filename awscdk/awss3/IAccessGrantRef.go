package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrant.
// Experimental.
type IAccessGrantRef interface {
	constructs.IConstruct
	// A reference to a AccessGrant resource.
	// Experimental.
	AccessGrantRef() *AccessGrantReference
}

// The jsii proxy for IAccessGrantRef
type jsiiProxy_IAccessGrantRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessGrantRef) AccessGrantRef() *AccessGrantReference {
	var returns *AccessGrantReference
	_jsii_.Get(
		j,
		"accessGrantRef",
		&returns,
	)
	return returns
}

