package interfacesawscognito

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentityPool.
// Experimental.
type IIdentityPoolRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IdentityPool resource.
	// Experimental.
	IdentityPoolRef() *IdentityPoolReference
}

// The jsii proxy for IIdentityPoolRef
type jsiiProxy_IIdentityPoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIdentityPoolRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIdentityPoolRef) IdentityPoolRef() *IdentityPoolReference {
	var returns *IdentityPoolReference
	_jsii_.Get(
		j,
		"identityPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPoolRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIdentityPoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

