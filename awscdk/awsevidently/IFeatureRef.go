package awsevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevidently/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Feature.
// Experimental.
type IFeatureRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Feature resource.
	// Experimental.
	FeatureRef() *FeatureReference
}

// The jsii proxy for IFeatureRef
type jsiiProxy_IFeatureRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IFeatureRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

