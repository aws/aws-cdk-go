package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HttpNamespace.
// Experimental.
type IHttpNamespaceRef interface {
	constructs.IConstruct
	// A reference to a HttpNamespace resource.
	// Experimental.
	HttpNamespaceRef() *HttpNamespaceReference
}

// The jsii proxy for IHttpNamespaceRef
type jsiiProxy_IHttpNamespaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHttpNamespaceRef) HttpNamespaceRef() *HttpNamespaceReference {
	var returns *HttpNamespaceReference
	_jsii_.Get(
		j,
		"httpNamespaceRef",
		&returns,
	)
	return returns
}

