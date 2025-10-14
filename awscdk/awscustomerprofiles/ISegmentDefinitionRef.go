package awscustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SegmentDefinition.
// Experimental.
type ISegmentDefinitionRef interface {
	constructs.IConstruct
	// A reference to a SegmentDefinition resource.
	// Experimental.
	SegmentDefinitionRef() *SegmentDefinitionReference
}

// The jsii proxy for ISegmentDefinitionRef
type jsiiProxy_ISegmentDefinitionRef struct {
	internal.Type__constructsIConstruct
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

