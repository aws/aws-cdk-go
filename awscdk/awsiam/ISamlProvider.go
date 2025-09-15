package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A SAML provider.
type ISamlProvider interface {
	awscdk.IResource
	ISAMLProviderRef
	// The Amazon Resource Name (ARN) of the provider.
	SamlProviderArn() *string
}

// The jsii proxy for ISamlProvider
type jsiiProxy_ISamlProvider struct {
	internal.Type__awscdkIResource
	jsiiProxy_ISAMLProviderRef
}

func (i *jsiiProxy_ISamlProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ISamlProvider) SamlProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISamlProvider) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISamlProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISamlProvider) SamlProviderRef() *SAMLProviderReference {
	var returns *SAMLProviderReference
	_jsii_.Get(
		j,
		"samlProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISamlProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

