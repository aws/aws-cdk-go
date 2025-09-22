package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrantsLocation.
// Experimental.
type IAccessGrantsLocationRef interface {
	constructs.IConstruct
	// A reference to a AccessGrantsLocation resource.
	// Experimental.
	AccessGrantsLocationRef() *AccessGrantsLocationReference
}

// The jsii proxy for IAccessGrantsLocationRef
type jsiiProxy_IAccessGrantsLocationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessGrantsLocationRef) AccessGrantsLocationRef() *AccessGrantsLocationReference {
	var returns *AccessGrantsLocationReference
	_jsii_.Get(
		j,
		"accessGrantsLocationRef",
		&returns,
	)
	return returns
}

