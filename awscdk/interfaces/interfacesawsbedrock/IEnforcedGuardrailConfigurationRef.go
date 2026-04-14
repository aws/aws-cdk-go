package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnforcedGuardrailConfiguration.
// Experimental.
type IEnforcedGuardrailConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnforcedGuardrailConfiguration resource.
	// Experimental.
	EnforcedGuardrailConfigurationRef() *EnforcedGuardrailConfigurationReference
}

// The jsii proxy for IEnforcedGuardrailConfigurationRef
type jsiiProxy_IEnforcedGuardrailConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEnforcedGuardrailConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEnforcedGuardrailConfigurationRef) EnforcedGuardrailConfigurationRef() *EnforcedGuardrailConfigurationReference {
	var returns *EnforcedGuardrailConfigurationReference
	_jsii_.Get(
		j,
		"enforcedGuardrailConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnforcedGuardrailConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnforcedGuardrailConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

