package awscassandra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscassandra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Type.
// Experimental.
type ITypeRef interface {
	constructs.IConstruct
	// A reference to a Type resource.
	// Experimental.
	TypeRef() *TypeReference
}

// The jsii proxy for ITypeRef
type jsiiProxy_ITypeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITypeRef) TypeRef() *TypeReference {
	var returns *TypeReference
	_jsii_.Get(
		j,
		"typeRef",
		&returns,
	)
	return returns
}

