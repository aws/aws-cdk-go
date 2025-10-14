package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrivateDnsNamespace.
// Experimental.
type IPrivateDnsNamespaceRef interface {
	constructs.IConstruct
	// A reference to a PrivateDnsNamespace resource.
	// Experimental.
	PrivateDnsNamespaceRef() *PrivateDnsNamespaceReference
}

// The jsii proxy for IPrivateDnsNamespaceRef
type jsiiProxy_IPrivateDnsNamespaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPrivateDnsNamespaceRef) PrivateDnsNamespaceRef() *PrivateDnsNamespaceReference {
	var returns *PrivateDnsNamespaceReference
	_jsii_.Get(
		j,
		"privateDnsNamespaceRef",
		&returns,
	)
	return returns
}

