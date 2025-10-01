package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAM.
// Experimental.
type IIPAMRef interface {
	constructs.IConstruct
	// A reference to a IPAM resource.
	// Experimental.
	IpamRef() *IPAMReference
}

// The jsii proxy for IIPAMRef
type jsiiProxy_IIPAMRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIPAMRef) IpamRef() *IPAMReference {
	var returns *IPAMReference
	_jsii_.Get(
		j,
		"ipamRef",
		&returns,
	)
	return returns
}

