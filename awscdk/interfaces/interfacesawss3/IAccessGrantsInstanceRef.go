package interfacesawss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessGrantsInstance.
// Experimental.
type IAccessGrantsInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccessGrantsInstance resource.
	// Experimental.
	AccessGrantsInstanceRef() *AccessGrantsInstanceReference
}

// The jsii proxy for IAccessGrantsInstanceRef
type jsiiProxy_IAccessGrantsInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAccessGrantsInstanceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAccessGrantsInstanceRef) AccessGrantsInstanceRef() *AccessGrantsInstanceReference {
	var returns *AccessGrantsInstanceReference
	_jsii_.Get(
		j,
		"accessGrantsInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessGrantsInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessGrantsInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

