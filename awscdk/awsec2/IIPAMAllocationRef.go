package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMAllocation.
// Experimental.
type IIPAMAllocationRef interface {
	constructs.IConstruct
	// A reference to a IPAMAllocation resource.
	// Experimental.
	IpamAllocationRef() *IPAMAllocationReference
}

// The jsii proxy for IIPAMAllocationRef
type jsiiProxy_IIPAMAllocationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIPAMAllocationRef) IpamAllocationRef() *IPAMAllocationReference {
	var returns *IPAMAllocationReference
	_jsii_.Get(
		j,
		"ipamAllocationRef",
		&returns,
	)
	return returns
}

