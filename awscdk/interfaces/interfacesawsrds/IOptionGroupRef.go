package interfacesawsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OptionGroup.
// Experimental.
type IOptionGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OptionGroup resource.
	// Experimental.
	OptionGroupRef() *OptionGroupReference
}

// The jsii proxy for IOptionGroupRef
type jsiiProxy_IOptionGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOptionGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOptionGroupRef) OptionGroupRef() *OptionGroupReference {
	var returns *OptionGroupReference
	_jsii_.Get(
		j,
		"optionGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

