package interfacesawsapplicationsignals

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapplicationsignals/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstrumentationConfig.
// Experimental.
type IInstrumentationConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InstrumentationConfig resource.
	// Experimental.
	InstrumentationConfigRef() *InstrumentationConfigReference
}

// The jsii proxy for IInstrumentationConfigRef
type jsiiProxy_IInstrumentationConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IInstrumentationConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IInstrumentationConfigRef) InstrumentationConfigRef() *InstrumentationConfigReference {
	var returns *InstrumentationConfigReference
	_jsii_.Get(
		j,
		"instrumentationConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstrumentationConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstrumentationConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

