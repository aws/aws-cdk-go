package interfacesawsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocationAzureBlob.
// Experimental.
type ILocationAzureBlobRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocationAzureBlob resource.
	// Experimental.
	LocationAzureBlobRef() *LocationAzureBlobReference
}

// The jsii proxy for ILocationAzureBlobRef
type jsiiProxy_ILocationAzureBlobRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILocationAzureBlobRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILocationAzureBlobRef) LocationAzureBlobRef() *LocationAzureBlobReference {
	var returns *LocationAzureBlobReference
	_jsii_.Get(
		j,
		"locationAzureBlobRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationAzureBlobRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocationAzureBlobRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

