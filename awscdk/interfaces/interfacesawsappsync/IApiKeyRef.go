package interfacesawsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiKey.
// Experimental.
type IApiKeyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApiKey resource.
	// Experimental.
	ApiKeyRef() *ApiKeyReference
}

// The jsii proxy for IApiKeyRef
type jsiiProxy_IApiKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IApiKeyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IApiKeyRef) ApiKeyRef() *ApiKeyReference {
	var returns *ApiKeyReference
	_jsii_.Get(
		j,
		"apiKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

