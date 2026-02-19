package interfacesawsiotcoredeviceadvisor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotcoredeviceadvisor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SuiteDefinition.
// Experimental.
type ISuiteDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SuiteDefinition resource.
	// Experimental.
	SuiteDefinitionRef() *SuiteDefinitionReference
}

// The jsii proxy for ISuiteDefinitionRef
type jsiiProxy_ISuiteDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISuiteDefinitionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISuiteDefinitionRef) SuiteDefinitionRef() *SuiteDefinitionReference {
	var returns *SuiteDefinitionReference
	_jsii_.Get(
		j,
		"suiteDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuiteDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISuiteDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

