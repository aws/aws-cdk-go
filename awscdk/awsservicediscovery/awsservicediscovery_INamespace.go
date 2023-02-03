package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
)

type INamespace interface {
	awscdk.IResource
	// Namespace ARN for the Namespace.
	NamespaceArn() *string
	// Namespace Id for the Namespace.
	NamespaceId() *string
	// A name for the Namespace.
	NamespaceName() *string
	// Type of Namespace.
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

