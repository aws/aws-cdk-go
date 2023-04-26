package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery/internal"
)

// Experimental.
type INamespace interface {
	awscdk.IResource
	// Namespace ARN for the Namespace.
	// Experimental.
	NamespaceArn() *string
	// Namespace Id for the Namespace.
	// Experimental.
	NamespaceId() *string
	// A name for the Namespace.
	// Experimental.
	NamespaceName() *string
	// Type of Namespace.
	// Experimental.
	Type() NamespaceType
}

// The jsii proxy for INamespace
type jsiiProxy_INamespace struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_INamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

