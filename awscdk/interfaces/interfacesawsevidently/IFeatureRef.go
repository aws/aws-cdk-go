package interfacesawsevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Feature.
// Experimental.
type IFeatureRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Feature resource.
	// Experimental.
	FeatureRef() *FeatureReference
}

// The jsii proxy for IFeatureRef
type jsiiProxy_IFeatureRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFeatureRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFeatureRef) FeatureRef() *FeatureReference {
	var returns *FeatureReference
	_jsii_.Get(
		j,
		"featureRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeatureRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeatureRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

