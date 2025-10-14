package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MultiRegionAccessPoint.
// Experimental.
type IMultiRegionAccessPointRef interface {
	constructs.IConstruct
	// A reference to a MultiRegionAccessPoint resource.
	// Experimental.
	MultiRegionAccessPointRef() *MultiRegionAccessPointReference
}

// The jsii proxy for IMultiRegionAccessPointRef
type jsiiProxy_IMultiRegionAccessPointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMultiRegionAccessPointRef) MultiRegionAccessPointRef() *MultiRegionAccessPointReference {
	var returns *MultiRegionAccessPointReference
	_jsii_.Get(
		j,
		"multiRegionAccessPointRef",
		&returns,
	)
	return returns
}

