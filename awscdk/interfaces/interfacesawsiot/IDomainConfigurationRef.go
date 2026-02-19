package interfacesawsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainConfiguration.
// Experimental.
type IDomainConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DomainConfiguration resource.
	// Experimental.
	DomainConfigurationRef() *DomainConfigurationReference
}

// The jsii proxy for IDomainConfigurationRef
type jsiiProxy_IDomainConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDomainConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDomainConfigurationRef) DomainConfigurationRef() *DomainConfigurationReference {
	var returns *DomainConfigurationReference
	_jsii_.Get(
		j,
		"domainConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

