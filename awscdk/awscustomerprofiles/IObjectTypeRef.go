package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ObjectType.
// Experimental.
type IObjectTypeRef interface {
	constructs.IConstruct
	// A reference to a ObjectType resource.
	// Experimental.
	ObjectTypeRef() *ObjectTypeReference
}

// The jsii proxy for IObjectTypeRef
type jsiiProxy_IObjectTypeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IObjectTypeRef) ObjectTypeRef() *ObjectTypeReference {
	var returns *ObjectTypeReference
	_jsii_.Get(
		j,
		"objectTypeRef",
		&returns,
	)
	return returns
}

