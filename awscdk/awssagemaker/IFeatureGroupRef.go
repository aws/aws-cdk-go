package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FeatureGroup.
// Experimental.
type IFeatureGroupRef interface {
	constructs.IConstruct
	// A reference to a FeatureGroup resource.
	// Experimental.
	FeatureGroupRef() *FeatureGroupReference
}

// The jsii proxy for IFeatureGroupRef
type jsiiProxy_IFeatureGroupRef struct {
	internal.Type__constructsIConstruct
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

