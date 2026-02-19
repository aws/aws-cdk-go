package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds"
	"github.com/aws/constructs-go/constructs/v10"
)

// An option group.
type IOptionGroup interface {
	interfacesawsrds.IOptionGroupRef
	awscdk.IResource
	// Adds a configuration to this OptionGroup.
	//
	// This method is a no-op for an imported OptionGroup.
	//
	// Returns: true if the OptionConfiguration was successfully added.
	AddConfiguration(configuration *OptionConfiguration) *bool
	// The name of the option group.
	OptionGroupName() *string
}

// The jsii proxy for IOptionGroup
type jsiiProxy_IOptionGroup struct {
	internal.Type__interfacesawsrdsIOptionGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOptionGroup) AddConfiguration(configuration *OptionConfiguration) *bool {
	if err := i.validateAddConfigurationParameters(configuration); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"addConfiguration",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IOptionGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IOptionGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOptionGroup) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroup) OptionGroupRef() *interfacesawsrds.OptionGroupReference {
	var returns *interfacesawsrds.OptionGroupReference
	_jsii_.Get(
		j,
		"optionGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOptionGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

