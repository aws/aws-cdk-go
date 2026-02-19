package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Cognito user pool client.
type IUserPoolClient interface {
	awscdk.IResource
	interfacesawscognito.IUserPoolClientRef
	// Name of the application client.
	UserPoolClientId() *string
	// The generated client secret.
	//
	// Only available if the "generateSecret" props is set to true.
	UserPoolClientSecret() awscdk.SecretValue
}

// The jsii proxy for IUserPoolClient
type jsiiProxy_IUserPoolClient struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawscognitoIUserPoolClientRef
}

func (i *jsiiProxy_IUserPoolClient) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IUserPoolClient) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IUserPoolClient) UserPoolClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClient) UserPoolClientSecret() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"userPoolClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClient) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClient) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClient) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolClient) UserPoolClientRef() *interfacesawscognito.UserPoolClientReference {
	var returns *interfacesawscognito.UserPoolClientReference
	_jsii_.Get(
		j,
		"userPoolClientRef",
		&returns,
	)
	return returns
}

