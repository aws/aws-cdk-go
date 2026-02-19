package awscognitoidentitypool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognitoidentitypool/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Cognito Identity Pool.
type IIdentityPool interface {
	interfacesawscognito.IIdentityPoolRef
	awscdk.IResource
	// The ARN of the Identity Pool.
	IdentityPoolArn() *string
	// The ID of the Identity Pool in the format REGION:GUID.
	IdentityPoolId() *string
	// Name of the Identity Pool.
	IdentityPoolName() *string
}

// The jsii proxy for IIdentityPool
type jsiiProxy_IIdentityPool struct {
	internal.Type__interfacesawscognitoIIdentityPoolRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IIdentityPool) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IIdentityPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIdentityPool) IdentityPoolArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) IdentityPoolRef() *interfacesawscognito.IdentityPoolReference {
	var returns *interfacesawscognito.IdentityPoolReference
	_jsii_.Get(
		j,
		"identityPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPool) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

