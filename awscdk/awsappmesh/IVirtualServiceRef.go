package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VirtualService.
// Experimental.
type IVirtualServiceRef interface {
	constructs.IConstruct
	// A reference to a VirtualService resource.
	// Experimental.
	VirtualServiceRef() *VirtualServiceReference
}

// The jsii proxy for IVirtualServiceRef
type jsiiProxy_IVirtualServiceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVirtualServiceRef) VirtualServiceRef() *VirtualServiceReference {
	var returns *VirtualServiceReference
	_jsii_.Get(
		j,
		"virtualServiceRef",
		&returns,
	)
	return returns
}

