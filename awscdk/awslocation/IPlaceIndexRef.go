package awslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PlaceIndex.
// Experimental.
type IPlaceIndexRef interface {
	constructs.IConstruct
	// A reference to a PlaceIndex resource.
	// Experimental.
	PlaceIndexRef() *PlaceIndexReference
}

// The jsii proxy for IPlaceIndexRef
type jsiiProxy_IPlaceIndexRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPlaceIndexRef) PlaceIndexRef() *PlaceIndexReference {
	var returns *PlaceIndexReference
	_jsii_.Get(
		j,
		"placeIndexRef",
		&returns,
	)
	return returns
}

