package interfacesawsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationsList.
// Experimental.
type IApplicationsListRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApplicationsList resource.
	// Experimental.
	ApplicationsListRef() *ApplicationsListReference
}

// The jsii proxy for IApplicationsListRef
type jsiiProxy_IApplicationsListRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IApplicationsListRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IApplicationsListRef) ApplicationsListRef() *ApplicationsListReference {
	var returns *ApplicationsListReference
	_jsii_.Get(
		j,
		"applicationsListRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationsListRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationsListRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

