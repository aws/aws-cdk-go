package interfacesawsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OAuthClientApplication.
// Experimental.
type IOAuthClientApplicationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OAuthClientApplication resource.
	// Experimental.
	OAuthClientApplicationRef() *OAuthClientApplicationReference
}

// The jsii proxy for IOAuthClientApplicationRef
type jsiiProxy_IOAuthClientApplicationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOAuthClientApplicationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOAuthClientApplicationRef) OAuthClientApplicationRef() *OAuthClientApplicationReference {
	var returns *OAuthClientApplicationReference
	_jsii_.Get(
		j,
		"oAuthClientApplicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuthClientApplicationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOAuthClientApplicationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

