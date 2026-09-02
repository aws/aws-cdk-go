package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/aws-cdk-go/awsmedialivealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaLive Input Security Group.
// Experimental.
type IInputSecurityGroup interface {
	interfacesawsmedialive.IInputSecurityGroupRef
	awscdk.IResource
	// The ARN of the input security group.
	// Experimental.
	InputSecurityGroupArn() *string
	// The ID of the input security group.
	// Experimental.
	InputSecurityGroupId() *string
}

// The jsii proxy for IInputSecurityGroup
type jsiiProxy_IInputSecurityGroup struct {
	internal.Type__interfacesawsmedialiveIInputSecurityGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IInputSecurityGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IInputSecurityGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IInputSecurityGroup) InputSecurityGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputSecurityGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputSecurityGroup) InputSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputSecurityGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputSecurityGroup) InputSecurityGroupRef() *interfacesawsmedialive.InputSecurityGroupReference {
	var returns *interfacesawsmedialive.InputSecurityGroupReference
	_jsii_.Get(
		j,
		"inputSecurityGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputSecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputSecurityGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

