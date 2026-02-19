package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseks"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an Amazon EKS Add-On.
type IAddon interface {
	interfacesawseks.IAddonRef
	awscdk.IResource
	// ARN of the Add-On.
	AddonArn() *string
	// Name of the Add-On.
	AddonName() *string
}

// The jsii proxy for IAddon
type jsiiProxy_IAddon struct {
	internal.Type__interfacesawseksIAddonRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAddon) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IAddon) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAddon) AddonArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddon) AddonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddon) AddonRef() *interfacesawseks.AddonReference {
	var returns *interfacesawseks.AddonReference
	_jsii_.Get(
		j,
		"addonRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddon) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddon) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddon) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

