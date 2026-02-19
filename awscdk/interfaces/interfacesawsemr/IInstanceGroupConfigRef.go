package interfacesawsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceGroupConfig.
// Experimental.
type IInstanceGroupConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a InstanceGroupConfig resource.
	// Experimental.
	InstanceGroupConfigRef() *InstanceGroupConfigReference
}

// The jsii proxy for IInstanceGroupConfigRef
type jsiiProxy_IInstanceGroupConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IInstanceGroupConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IInstanceGroupConfigRef) InstanceGroupConfigRef() *InstanceGroupConfigReference {
	var returns *InstanceGroupConfigReference
	_jsii_.Get(
		j,
		"instanceGroupConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceGroupConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceGroupConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

