package interfacesawscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a UserPoolDomain.
// Experimental.
type IUserPoolDomainRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a UserPoolDomain resource.
	// Experimental.
	UserPoolDomainRef() *UserPoolDomainReference
}

// The jsii proxy for IUserPoolDomainRef
type jsiiProxy_IUserPoolDomainRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IUserPoolDomainRef) UserPoolDomainRef() *UserPoolDomainReference {
	var returns *UserPoolDomainReference
	_jsii_.Get(
		j,
		"userPoolDomainRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolDomainRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolDomainRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

