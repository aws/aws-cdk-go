package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Offering.
// Experimental.
type IOfferingRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Offering resource.
	// Experimental.
	OfferingRef() *OfferingReference
}

// The jsii proxy for IOfferingRef
type jsiiProxy_IOfferingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOfferingRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOfferingRef) OfferingRef() *OfferingReference {
	var returns *OfferingReference
	_jsii_.Get(
		j,
		"offeringRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOfferingRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOfferingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

