package awsemrcontainers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsemrcontainers/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualCluster.
// Experimental.
type IVirtualClusterRef interface {
	constructs.IConstruct
	// A reference to a VirtualCluster resource.
	// Experimental.
	VirtualClusterRef() *VirtualClusterReference
}

// The jsii proxy for IVirtualClusterRef
type jsiiProxy_IVirtualClusterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVirtualClusterRef) VirtualClusterRef() *VirtualClusterReference {
	var returns *VirtualClusterReference
	_jsii_.Get(
		j,
		"virtualClusterRef",
		&returns,
	)
	return returns
}

