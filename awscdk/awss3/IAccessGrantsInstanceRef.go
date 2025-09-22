package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrantsInstance.
// Experimental.
type IAccessGrantsInstanceRef interface {
	constructs.IConstruct
	// A reference to a AccessGrantsInstance resource.
	// Experimental.
	AccessGrantsInstanceRef() *AccessGrantsInstanceReference
}

// The jsii proxy for IAccessGrantsInstanceRef
type jsiiProxy_IAccessGrantsInstanceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessGrantsInstanceRef) AccessGrantsInstanceRef() *AccessGrantsInstanceReference {
	var returns *AccessGrantsInstanceReference
	_jsii_.Get(
		j,
		"accessGrantsInstanceRef",
		&returns,
	)
	return returns
}

