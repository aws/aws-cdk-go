package interfacesawssigner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SigningProfile.
// Experimental.
type ISigningProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SigningProfile resource.
	// Experimental.
	SigningProfileRef() *SigningProfileReference
}

// The jsii proxy for ISigningProfileRef
type jsiiProxy_ISigningProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISigningProfileRef) SigningProfileRef() *SigningProfileReference {
	var returns *SigningProfileReference
	_jsii_.Get(
		j,
		"signingProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISigningProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

