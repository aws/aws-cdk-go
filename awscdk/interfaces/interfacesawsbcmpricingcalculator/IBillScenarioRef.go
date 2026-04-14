package interfacesawsbcmpricingcalculator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbcmpricingcalculator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BillScenario.
// Experimental.
type IBillScenarioRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BillScenario resource.
	// Experimental.
	BillScenarioRef() *BillScenarioReference
}

// The jsii proxy for IBillScenarioRef
type jsiiProxy_IBillScenarioRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBillScenarioRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBillScenarioRef) BillScenarioRef() *BillScenarioReference {
	var returns *BillScenarioReference
	_jsii_.Get(
		j,
		"billScenarioRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBillScenarioRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBillScenarioRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

