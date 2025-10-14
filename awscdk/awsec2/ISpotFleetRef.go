package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SpotFleet.
// Experimental.
type ISpotFleetRef interface {
	constructs.IConstruct
	// A reference to a SpotFleet resource.
	// Experimental.
	SpotFleetRef() *SpotFleetReference
}

// The jsii proxy for ISpotFleetRef
type jsiiProxy_ISpotFleetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISpotFleetRef) SpotFleetRef() *SpotFleetReference {
	var returns *SpotFleetReference
	_jsii_.Get(
		j,
		"spotFleetRef",
		&returns,
	)
	return returns
}

