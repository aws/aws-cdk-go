package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FeatureGroup.
// Experimental.
type IFeatureGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FeatureGroup resource.
	// Experimental.
	FeatureGroupRef() *FeatureGroupReference
}

// The jsii proxy for IFeatureGroupRef
type jsiiProxy_IFeatureGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFeatureGroupRef) FeatureGroupRef() *FeatureGroupReference {
	var returns *FeatureGroupReference
	_jsii_.Get(
		j,
		"featureGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeatureGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFeatureGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

