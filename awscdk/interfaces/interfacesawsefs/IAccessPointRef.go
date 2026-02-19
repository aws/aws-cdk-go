package interfacesawsefs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsefs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPoint.
// Experimental.
type IAccessPointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccessPoint resource.
	// Experimental.
	AccessPointRef() *AccessPointReference
}

// The jsii proxy for IAccessPointRef
type jsiiProxy_IAccessPointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAccessPointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAccessPointRef) AccessPointRef() *AccessPointReference {
	var returns *AccessPointReference
	_jsii_.Get(
		j,
		"accessPointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

