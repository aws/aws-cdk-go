package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SegmentDefinition.
// Experimental.
type ISegmentDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SegmentDefinition resource.
	// Experimental.
	SegmentDefinitionRef() *SegmentDefinitionReference
}

// The jsii proxy for ISegmentDefinitionRef
type jsiiProxy_ISegmentDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISegmentDefinitionRef) SegmentDefinitionRef() *SegmentDefinitionReference {
	var returns *SegmentDefinitionReference
	_jsii_.Get(
		j,
		"segmentDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISegmentDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISegmentDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

