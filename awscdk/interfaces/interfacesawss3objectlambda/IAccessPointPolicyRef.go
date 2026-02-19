package interfacesawss3objectlambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3objectlambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPointPolicy.
// Experimental.
type IAccessPointPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AccessPointPolicy resource.
	// Experimental.
	AccessPointPolicyRef() *AccessPointPolicyReference
}

// The jsii proxy for IAccessPointPolicyRef
type jsiiProxy_IAccessPointPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAccessPointPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAccessPointPolicyRef) AccessPointPolicyRef() *AccessPointPolicyReference {
	var returns *AccessPointPolicyReference
	_jsii_.Get(
		j,
		"accessPointPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPointPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

