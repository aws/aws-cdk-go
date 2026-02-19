package interfacesawsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DomainUnit.
// Experimental.
type IDomainUnitRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DomainUnit resource.
	// Experimental.
	DomainUnitRef() *DomainUnitReference
}

// The jsii proxy for IDomainUnitRef
type jsiiProxy_IDomainUnitRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDomainUnitRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDomainUnitRef) DomainUnitRef() *DomainUnitReference {
	var returns *DomainUnitReference
	_jsii_.Get(
		j,
		"domainUnitRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainUnitRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDomainUnitRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

