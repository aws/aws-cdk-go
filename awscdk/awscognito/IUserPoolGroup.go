package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a user pool group.
type IUserPoolGroup interface {
	awscdk.IResource
	interfacesawscognito.IUserPoolGroupRef
	// The user group name.
	GroupName() *string
}

// The jsii proxy for IUserPoolGroup
type jsiiProxy_IUserPoolGroup struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawscognitoIUserPoolGroupRef
}

func (i *jsiiProxy_IUserPoolGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IUserPoolGroup) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolGroup) UserPoolGroupRef() *interfacesawscognito.UserPoolGroupReference {
	var returns *interfacesawscognito.UserPoolGroupReference
	_jsii_.Get(
		j,
		"userPoolGroupRef",
		&returns,
	)
	return returns
}

