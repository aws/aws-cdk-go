package awscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Cognito user pool resource server.
type IUserPoolResourceServer interface {
	awscdk.IResource
	interfacesawscognito.IUserPoolResourceServerRef
	// Resource server id.
	UserPoolResourceServerId() *string
}

// The jsii proxy for IUserPoolResourceServer
type jsiiProxy_IUserPoolResourceServer struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawscognitoIUserPoolResourceServerRef
}

func (i *jsiiProxy_IUserPoolResourceServer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IUserPoolResourceServer) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IUserPoolResourceServer) UserPoolResourceServerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolResourceServerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolResourceServer) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolResourceServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolResourceServer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserPoolResourceServer) UserPoolResourceServerRef() *interfacesawscognito.UserPoolResourceServerReference {
	var returns *interfacesawscognito.UserPoolResourceServerReference
	_jsii_.Get(
		j,
		"userPoolResourceServerRef",
		&returns,
	)
	return returns
}

