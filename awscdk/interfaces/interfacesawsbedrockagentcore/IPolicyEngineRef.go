package interfacesawsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyEngine.
// Experimental.
type IPolicyEngineRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PolicyEngine resource.
	// Experimental.
	PolicyEngineRef() *PolicyEngineReference
}

// The jsii proxy for IPolicyEngineRef
type jsiiProxy_IPolicyEngineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPolicyEngineRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPolicyEngineRef) PolicyEngineRef() *PolicyEngineReference {
	var returns *PolicyEngineReference
	_jsii_.Get(
		j,
		"policyEngineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngineRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyEngineRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

