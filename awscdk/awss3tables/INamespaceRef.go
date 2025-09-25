package awss3tables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Namespace.
// Experimental.
type INamespaceRef interface {
	constructs.IConstruct
	// A reference to a Namespace resource.
	// Experimental.
	NamespaceRef() *NamespaceReference
}

// The jsii proxy for INamespaceRef
type jsiiProxy_INamespaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INamespaceRef) NamespaceRef() *NamespaceReference {
	var returns *NamespaceReference
	_jsii_.Get(
		j,
		"namespaceRef",
		&returns,
	)
	return returns
}

