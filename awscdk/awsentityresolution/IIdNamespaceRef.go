package awsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdNamespace.
// Experimental.
type IIdNamespaceRef interface {
	constructs.IConstruct
	// A reference to a IdNamespace resource.
	// Experimental.
	IdNamespaceRef() *IdNamespaceReference
}

// The jsii proxy for IIdNamespaceRef
type jsiiProxy_IIdNamespaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIdNamespaceRef) IdNamespaceRef() *IdNamespaceReference {
	var returns *IdNamespaceReference
	_jsii_.Get(
		j,
		"idNamespaceRef",
		&returns,
	)
	return returns
}

