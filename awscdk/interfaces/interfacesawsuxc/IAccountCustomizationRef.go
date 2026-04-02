package interfacesawsuxc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsuxc/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccountCustomization.
// Experimental.
type IAccountCustomizationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccountCustomization resource.
	// Experimental.
	AccountCustomizationRef() *AccountCustomizationReference
}

// The jsii proxy for IAccountCustomizationRef
type jsiiProxy_IAccountCustomizationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAccountCustomizationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAccountCustomizationRef) AccountCustomizationRef() *AccountCustomizationReference {
	var returns *AccountCustomizationReference
	_jsii_.Get(
		j,
		"accountCustomizationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountCustomizationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccountCustomizationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

