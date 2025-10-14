package awsmemorydb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MultiRegionCluster.
// Experimental.
type IMultiRegionClusterRef interface {
	constructs.IConstruct
	// A reference to a MultiRegionCluster resource.
	// Experimental.
	MultiRegionClusterRef() *MultiRegionClusterReference
}

// The jsii proxy for IMultiRegionClusterRef
type jsiiProxy_IMultiRegionClusterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMultiRegionClusterRef) MultiRegionClusterRef() *MultiRegionClusterReference {
	var returns *MultiRegionClusterReference
	_jsii_.Get(
		j,
		"multiRegionClusterRef",
		&returns,
	)
	return returns
}

