package interfacesawsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SAMLProvider.
// Experimental.
type ISAMLProviderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SAMLProvider resource.
	// Experimental.
	SamlProviderRef() *SAMLProviderReference
}

// The jsii proxy for ISAMLProviderRef
type jsiiProxy_ISAMLProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISAMLProviderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISAMLProviderRef) SamlProviderRef() *SAMLProviderReference {
	var returns *SAMLProviderReference
	_jsii_.Get(
		j,
		"samlProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISAMLProviderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISAMLProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

