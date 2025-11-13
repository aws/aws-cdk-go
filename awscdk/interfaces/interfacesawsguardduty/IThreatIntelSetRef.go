package interfacesawsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ThreatIntelSet.
// Experimental.
type IThreatIntelSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ThreatIntelSet resource.
	// Experimental.
	ThreatIntelSetRef() *ThreatIntelSetReference
}

// The jsii proxy for IThreatIntelSetRef
type jsiiProxy_IThreatIntelSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IThreatIntelSetRef) ThreatIntelSetRef() *ThreatIntelSetReference {
	var returns *ThreatIntelSetReference
	_jsii_.Get(
		j,
		"threatIntelSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThreatIntelSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IThreatIntelSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

