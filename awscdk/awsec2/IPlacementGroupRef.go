package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlacementGroup.
// Experimental.
type IPlacementGroupRef interface {
	constructs.IConstruct
	// A reference to a PlacementGroup resource.
	// Experimental.
	PlacementGroupRef() *PlacementGroupReference
}

// The jsii proxy for IPlacementGroupRef
type jsiiProxy_IPlacementGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPlacementGroupRef) PlacementGroupRef() *PlacementGroupReference {
	var returns *PlacementGroupReference
	_jsii_.Get(
		j,
		"placementGroupRef",
		&returns,
	)
	return returns
}

