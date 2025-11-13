package interfacesawsaiops

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaiops/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InvestigationGroup.
// Experimental.
type IInvestigationGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InvestigationGroup resource.
	// Experimental.
	InvestigationGroupRef() *InvestigationGroupReference
}

// The jsii proxy for IInvestigationGroupRef
type jsiiProxy_IInvestigationGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IInvestigationGroupRef) InvestigationGroupRef() *InvestigationGroupReference {
	var returns *InvestigationGroupReference
	_jsii_.Get(
		j,
		"investigationGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInvestigationGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInvestigationGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

