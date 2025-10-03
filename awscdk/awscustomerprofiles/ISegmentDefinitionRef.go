package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SegmentDefinition.
// Experimental.
type ISegmentDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISegmentDefinitionRef
type jsiiProxy_ISegmentDefinitionRef struct {
	internal.Type__constructsIConstruct
}

