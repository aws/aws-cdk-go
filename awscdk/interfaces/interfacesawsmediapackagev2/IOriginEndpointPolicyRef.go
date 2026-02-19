package interfacesawsmediapackagev2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a OriginEndpointPolicy.
// Experimental.
type IOriginEndpointPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a OriginEndpointPolicy resource.
	// Experimental.
	OriginEndpointPolicyRef() *OriginEndpointPolicyReference
}

// The jsii proxy for IOriginEndpointPolicyRef
type jsiiProxy_IOriginEndpointPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IOriginEndpointPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOriginEndpointPolicyRef) OriginEndpointPolicyRef() *OriginEndpointPolicyReference {
	var returns *OriginEndpointPolicyReference
	_jsii_.Get(
		j,
		"originEndpointPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpointPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpointPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

