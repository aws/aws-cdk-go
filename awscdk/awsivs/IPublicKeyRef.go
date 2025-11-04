package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicKey.
// Experimental.
type IPublicKeyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PublicKey resource.
	// Experimental.
	PublicKeyRef() *PublicKeyReference
}

// The jsii proxy for IPublicKeyRef
type jsiiProxy_IPublicKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPublicKeyRef) PublicKeyRef() *PublicKeyReference {
	var returns *PublicKeyReference
	_jsii_.Get(
		j,
		"publicKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicKeyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPublicKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

